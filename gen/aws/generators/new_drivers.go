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

package main

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func generateNewDrivers() {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, DRIVERS_DIR, func(os.FileInfo) bool { return true }, 0)
	if err != nil {
		panic(err)
	}

	finder := &findStructs{}
	for _, pkg := range pkgs {
		for _, f := range pkg.Files {
			ast.Walk(finder, f)
		}
	}

	templ, err := template.New("cmdRuns").Parse(cmdRuns)
	if err != nil {
		panic(err)
	}

	var buff bytes.Buffer
	err = templ.Execute(&buff, finder.result)
	if err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(filepath.Join(DRIVERS_DIR, "gen_cmd_runs.go"), buff.Bytes(), 0666); err != nil {
		panic(err)
	}

	templ, err = template.New("cmdInits").Funcs(template.FuncMap{
		"ToLower": strings.ToLower,
	}).Parse(cmdInits)
	if err != nil {
		panic(err)
	}

	buff.Reset()
	err = templ.Execute(&buff, finder.result)
	if err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(filepath.Join(DRIVERS_DIR, "gen_cmd_inits.go"), buff.Bytes(), 0666); err != nil {
		panic(err)
	}
}

type cmdData struct {
	API, Call, Input, Output string
	HasDryRun                bool
}

type findStructs struct {
	result map[string]cmdData
}

func (v *findStructs) Visit(node ast.Node) (w ast.Visitor) {
	if v.result == nil {
		v.result = make(map[string]cmdData)
	}
	if typ, ok := node.(*ast.TypeSpec); ok {
		if s, isStruct := typ.Type.(*ast.StructType); isStruct {
			for _, f := range s.Fields.List {
				if tag := f.Tag; tag != nil && strings.Contains(tag.Value, "awsAPI") {
					key := typ.Name.Name
					v.result[key] = extractTag(tag.Value)
					break
				}
			}
		}
	}
	return v
}

func extractTag(s string) (t cmdData) {
	splits := strings.Split(s[1:len(s)-1], " ")
	for i, e := range splits {
		el := strings.Split(e, ":")
		ell := el[1][1 : len(el[1])-1]
		switch {
		case i == 0:
			t.API = ell
		case i == 1:
			t.Call = ell
		case i == 2:
			t.Input = ell
		case i == 3:
			t.Output = ell
		case i == 4:
			t.HasDryRun = true
		}
	}
	return
}

const cmdRuns = `/* Copyright 2017 WALLIX

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

// DO NOT EDIT
// This file was automatically generated with go generate
package awsdriver

{{ range $cmdName, $tag := . }}
func New{{ $cmdName }}(l *logger.Logger, sess *session.Session) *{{ $cmdName }}{
	cmd := new({{ $cmdName }})
	cmd.api = {{ $tag.API }}.New(sess)
	cmd.logger = l
	return cmd
}
{{ if $tag.Call }}
func (cmd *{{ $cmdName }}) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("{{ $cmdName }}: BeforeRun: %s", brErr)
		}
	}

	if err := structSetter(cmd, params); err != nil {
		return nil, fmt.Errorf("{{ $cmdName }}: cannot set params on command struct: %s", err)
	}

	input := &{{ $tag.Input }}{}
	if err := structInjector(cmd, input) ; err != nil {
		return nil, fmt.Errorf("{{ $cmdName }}: cannot inject in {{ $tag.Input }}: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.{{ $tag.Call }}(input)
	cmd.logger.ExtraVerbosef("{{ $tag.API }}.{{ $tag.Call }} call took %s", time.Since(start))
	if err != nil {
		return nil, fmt.Errorf("{{ $cmdName }}: %s", err)
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("{{ $cmdName }}: AfterRun: %s", brErr)
		}
	}

	return cmd.ExtractResultString(output), nil
}
{{- end }}
{{ if $tag.HasDryRun }}
func (cmd *{{ $cmdName }}) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := structSetter(cmd, params); err != nil {
		return nil, fmt.Errorf("dry run: {{ $cmdName }}: cannot set params on command struct: %s", err)
	}

	input := &{{ $tag.Input }}{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input) ; err != nil {
		return nil, fmt.Errorf("dry run: {{ $cmdName }}: cannot inject in {{ $tag.Input }}: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.{{ $tag.Call }}(input);
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: {{ $tag.API }}.{{ $tag.Call }} call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: {{ $cmdName }} ok")
			return fakeDryRunId(cmd.Entity()), nil
		}
	}

	return nil, fmt.Errorf("dry run: {{ $cmdName }} : %s", err) 
}
{{- end }}
{{ end }}
`

const cmdInits = `/* Copyright 2017 WALLIX

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

// DO NOT EDIT
// This file was automatically generated with go generate
package awsdriver

var NewCommandFuncs = map[string]func() interface{}{}

func InitCommands(l *logger.Logger, sess *session.Session) {
{{- range $cmdName, $tag := . }}
	NewCommandFuncs["{{ ToLower $cmdName }}"] = func() interface{} { return New{{ $cmdName }}(l, sess) }
{{- end }}
}
`
