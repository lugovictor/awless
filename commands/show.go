/*
Copyright 2017 WALLIX

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package commands

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/wallix/awless/cloud"
	"github.com/wallix/awless/config"
	"github.com/wallix/awless/console"
	"github.com/wallix/awless/graph"
	"github.com/wallix/awless/logger"
	"github.com/wallix/awless/sync"
)

var (
	listAllSiblingsFlag          bool
	noAliasFlag                  bool
	showPropertiesValuesOnlyFlag []string
)

func init() {
	RootCmd.AddCommand(showCmd)
	showCmd.Flags().BoolVar(&listAllSiblingsFlag, "siblings", false, "List all the resource's siblings")
	showCmd.Flags().BoolVar(&noAliasFlag, "no-alias", false, "Disable the resolution of ID to alias")
	showCmd.Flags().StringSliceVar(&showPropertiesValuesOnlyFlag, "values-for", []string{}, "Output values only for given properties keys")
}

var showCmd = &cobra.Command{
	Use:   "show REFERENCE",
	Short: "Show resources lineage and dependencies given a REFERENCE: name, id, arn, etc...",
	Example: `  awless show i-8d43b21b            # show an instance via its ref
  awless show AIDAJ3Z24GOKHTZO4OIX6 # show a user via its ref
  awless show jsmith                # show a user via its ref,
  awless show @jsmith               # forcing search by name`,
	PersistentPreRun:  applyHooks(initLoggerHook, initAwlessEnvHook, initCloudServicesHook, initSyncerHook, firstInstallDoneHook),
	PersistentPostRun: applyHooks(verifyNewVersionHook, onVersionUpgrade, networkMonitorHook),

	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("REFERENCE required. See examples.")
		}

		ref := args[0]
		notFound := fmt.Errorf("resource with reference '%s' not found", deprefix(ref))

		var resource *graph.Resource
		var gph *graph.Graph

		resource, gph = findResourceInLocalGraphs(ref)

		if resource == nil && localGlobalFlag {
			exitOn(decorateWithSuggestion(notFound, ref))
		} else if resource == nil {
			runFullSync()

			if resource, gph = findResourceInLocalGraphs(ref); resource == nil {
				exitOn(decorateWithSuggestion(notFound, ref))
			}
		}

		if !localGlobalFlag && config.GetAutosync() {
			var services []cloud.Service
			if resource.Type() == cloud.Region {
				services = append(services, cloud.AllServices()...)
			} else {
				srv, err := cloud.GetServiceForType(resource.Type())
				exitOn(err)
				services = append(services, srv)
			}

			logger.Verbosef("syncing services for %s type", resource.Type())
			if _, err := sync.DefaultSyncer.Sync(services...); err != nil {
				logger.Verbose(err)
			}
			resource, gph = findResourceInLocalGraphs(ref)
		}

		if resource != nil {
			if len(showPropertiesValuesOnlyFlag) > 0 {
				showResourceValuesOnlyFor(resource, showPropertiesValuesOnlyFlag)
			} else {
				showResource(resource, gph)
			}
		}

		return nil
	},
}

func showResourceValuesOnlyFor(resource *graph.Resource, propKeys []string) {
	var normalized []string
	for _, p := range propKeys {
		normalized = append(normalized, strings.ToLower(strings.Replace(p, " ", "", -1)))
	}

	valuesForKeys := map[string]string{}
	isIncluded := func(s string) (bool, string) {
		for _, n := range normalized {
			if n == strings.ToLower(s) {
				return true, n
			}
		}
		return false, ""
	}
	for k, v := range resource.Properties {
		if ok, p := isIncluded(k); ok {
			valuesForKeys[p] = fmt.Sprint(v)
		}
	}

	var values []string
	for _, n := range normalized {
		if v, ok := valuesForKeys[n]; ok {
			values = append(values, v)
		}
	}

	if len(values) > 0 {
		fmt.Println(strings.Join(values, ","))
	} else {
		exitOn(fmt.Errorf("no values for %q", propKeys))
	}
}

func showResource(resource *graph.Resource, gph *graph.Graph) {
	displayer, err := console.BuildOptions(
		console.WithColumnDefinitions(console.DefaultsColumnDefinitions[resource.Type()]),
		console.WithFormat(listingFormat),
		console.WithMaxWidth(console.GetTerminalWidth()),
	).SetSource(resource).Build()
	exitOn(err)

	exitOn(displayer.Print(os.Stdout))

	var parents []*graph.Resource
	err = gph.Accept(&graph.ParentsVisitor{From: resource, Each: graph.VisitorCollectFunc(&parents)})
	exitOn(err)

	var parentsW bytes.Buffer
	var count int
	for i := len(parents) - 1; i >= 0; i-- {
		if count == 0 {
			fmt.Fprintf(&parentsW, "%s\n", printResourceRef(parents[i]))
		} else {
			fmt.Fprintf(&parentsW, "%s↳ %s\n", strings.Repeat("\t", count), printResourceRef(parents[i]))
		}
		count++
	}

	var childrenW bytes.Buffer
	var hasChildren bool
	printWithTabs := func(r *graph.Resource, distance int) error {
		var tabs bytes.Buffer
		tabs.WriteString(strings.Repeat("\t", count))
		for i := 0; i < distance; i++ {
			tabs.WriteByte('\t')
		}

		display := printResourceRef(r)
		if r.Same(resource) {
			display = printResourceRef(resource, renderGreenFn)
		} else {
			hasChildren = true
		}
		fmt.Fprintf(&childrenW, "%s↳ %s\n", tabs.String(), display)
		return nil
	}
	err = gph.Accept(&graph.ChildrenVisitor{From: resource, Each: printWithTabs, IncludeFrom: true})
	exitOn(err)

	if len(parents) > 0 || hasChildren {
		fmt.Println(renderCyanBoldFn("\nLineage:"))
		fmt.Printf(parentsW.String())
		fmt.Printf(childrenW.String())
	}

	appliedOn, err := gph.ListResourcesAppliedOn(resource)
	exitOn(err)
	printResourceList(renderCyanBoldFn("Applied on"), appliedOn)

	dependingOn, err := gph.ListResourcesDependingOn(resource)
	exitOn(err)
	printResourceList(renderCyanBoldFn("Depending on"), dependingOn)

	var siblings []*graph.Resource
	err = gph.Accept(&graph.SiblingsVisitor{From: resource, Each: graph.VisitorCollectFunc(&siblings)})
	exitOn(err)
	printResourceList(renderCyanBoldFn("Siblings"), siblings, "display all with flag --siblings")
}

func runFullSync() {
	if !config.GetAutosync() {
		logger.Info("autosync disabled")
		return
	}

	logger.Infof("cannot find resource in existing data synced locally")
	logger.Infof("running sync for current region '%s'", config.GetAWSRegion())

	var services []cloud.Service
	for _, srv := range cloud.ServiceRegistry {
		services = append(services, srv)
	}

	if _, err := sync.DefaultSyncer.Sync(services...); err != nil {
		logger.Verbose(err)
	}
}

func findResourceInLocalGraphs(ref string) (*graph.Resource, *graph.Graph) {
	g, resources := resolveResourceFromRefInCurrentRegion(ref)
	switch len(resources) {
	case 0:
		return nil, nil
	case 1:
		return resources[0], g
	default:
		logger.Infof("%d resources found with name '%s' in region '%s'. Show a specific resource with:", len(resources), deprefix(ref), config.GetAWSRegion())
		for _, res := range resources {
			var buf bytes.Buffer
			buf.WriteString(fmt.Sprintf("\t`awless show %s` to show the %s", res.Id(), res.Type()))
			if state, ok := res.Properties["State"].(string); ok {
				buf.WriteString(fmt.Sprintf(" (state: '%s')", state))
			}
			logger.Infof(buf.String())
		}

		os.Exit(0)
	}

	return nil, nil
}

func resolveResourceFromRefInCurrentRegion(ref string) (*graph.Graph, []*graph.Resource) {
	g, err := sync.LoadLocalGraphs(config.GetAWSRegion())
	exitOn(err)
	return resolveResourceFromRef(g, ref)
}

func resolveResourceFromRefInAllLocalRegion(ref string) (*graph.Graph, []*graph.Resource) {
	g, err := sync.LoadAllLocalGraphs()
	exitOn(err)
	return resolveResourceFromRef(g, ref)
}

func resolveResourceFromRef(g *graph.Graph, ref string) (*graph.Graph, []*graph.Resource) {
	name := deprefix(ref)
	byName := &graph.ByProperty{Key: "Name", Value: name}

	if strings.HasPrefix(ref, "@") {
		logger.Verbosef("prefixed with @: forcing research by name '%s'", name)
		rs, err := g.ResolveResources(byName)
		exitOn(err)
		return g, rs
	}
	rs, err := g.ResolveResources(&graph.ById{Id: name})
	exitOn(err)

	if len(rs) > 0 {
		return g, rs
	}

	rs, err = g.ResolveResources(
		byName,
		&graph.ByProperty{Key: "Arn", Value: name},
	)
	exitOn(err)

	return g, rs
}

func deprefix(s string) string {
	return strings.TrimPrefix(s, "@")
}

func decorateWithSuggestion(err error, ref string) error {
	buf := bytes.NewBufferString(fmt.Sprintf("%s in region %s", err.Error(), config.GetAWSRegion()))
	g, resources := resolveResourceFromRefInAllLocalRegion(ref)
	for _, res := range resources {
		if found := g.FindAncestor(res, cloud.Region); found != nil {
			buf.WriteString(fmt.Sprintf("\n\tfound previously synced under region %s as %s. Show it with `awless show %s -r %s --local`", found.Id(), res, res.Id(), found.Id()))
		}
	}
	return errors.New(buf.String())
}

func printResourceList(title string, list []*graph.Resource, shortenListMsg ...string) {
	sort.Sort(byTypeAndString{list})
	all := graph.Resources(list).Map(func(r *graph.Resource) string { return printResourceRef(r) })
	count := len(all)
	max := 3
	if count > 0 {
		if !listAllSiblingsFlag && len(shortenListMsg) > 0 && count > max {
			fmt.Printf("\n%s: %s, ... (%s)\n", title, strings.Join(all[0:max], ", "), shortenListMsg[0])
		} else {
			fmt.Printf("\n%s: %s\n", title, strings.Join(all, ", "))
		}
	}
}

func printResourceRef(r *graph.Resource, idRenderFunc ...func(a ...interface{}) string) string {
	render := fmt.Sprint
	if len(idRenderFunc) > 0 {
		render = idRenderFunc[0]
	}
	if noAliasFlag {
		return r.Format(render("%i") + "[" + color.New(color.FgBlue, color.Bold).SprintFunc()("%t") + "]")
	}
	return r.Format(render("%n") + "[" + color.New(color.FgBlue, color.Bold).SprintFunc()("%t") + "]")
}

type byTypeAndString struct {
	res []*graph.Resource
}

func (b byTypeAndString) Len() int { return len(b.res) }
func (b byTypeAndString) Swap(i, j int) {
	b.res[i], b.res[j] = b.res[j], b.res[i]
}
func (b byTypeAndString) Less(i, j int) bool {
	if b.res[i].Type() != b.res[j].Type() {
		return b.res[i].Type() < b.res[j].Type()
	}
	return b.res[i].String() <= b.res[j].String()
}
