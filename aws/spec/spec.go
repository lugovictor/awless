package awsspec

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/wallix/awless/cloud"
)

const (
	dryRunOperation = "DryRunOperation"
	notFound        = "NotFound"
)

type BeforeRunner interface {
	BeforeRun(ctx, params map[string]interface{}) error
}

type AfterRunner interface {
	AfterRun(ctx map[string]interface{}, output interface{}) error
}

type ManualValidator interface {
	ManualValidateCommand(params map[string]interface{}) []error
}

type command interface {
	ValidateParams([]string) ([]string, error)
	ValidateCommand(map[string]interface{}) []error
	inject(params map[string]interface{}) error
	Run(ctx map[string]interface{}, params map[string]interface{}) (interface{}, error)
}

func implementsBeforeRun(i interface{}) (BeforeRunner, bool) {
	v, ok := i.(BeforeRunner)
	return v, ok
}

func implementsAfterRun(i interface{}) (AfterRunner, bool) {
	v, ok := i.(AfterRunner)
	return v, ok
}

func implementsManualValidator(i interface{}) (ManualValidator, bool) {
	v, ok := i.(ManualValidator)
	return v, ok
}

func fakeDryRunId(entity string) string {
	suffix := rand.Intn(1e6)
	switch entity {
	case cloud.Instance:
		return fmt.Sprintf("i-%d", suffix)
	case cloud.Subnet:
		return fmt.Sprintf("subnet-%d", suffix)
	case cloud.Vpc:
		return fmt.Sprintf("vpc-%d", suffix)
	case cloud.Volume:
		return fmt.Sprintf("vol-%d", suffix)
	case cloud.SecurityGroup:
		return fmt.Sprintf("sg-%d", suffix)
	case cloud.InternetGateway:
		return fmt.Sprintf("igw-%d", suffix)
	case cloud.NatGateway:
		return fmt.Sprintf("nat-%d", suffix)
	case cloud.RouteTable:
		return fmt.Sprintf("rtb-%d", suffix)
	default:
		return fmt.Sprintf("dryrunid-%d", suffix)
	}
}

type paramRule struct {
	tree   ruleNode
	extras []string
}

func (p paramRule) help() string {
	if len(p.extras) == 0 {
		return p.tree.help()
	}

	return fmt.Sprintf("%s or extra params: \"%s\"", p.tree.help(), strings.Join(p.extras, "\", \""))
}

func (p paramRule) verify(keys []string) ([]string, error) {
	if p.tree == nil {
		return nil, nil
	}
	var unexpected []string
	for _, key := range keys {
		if p.tree.unexpected(key) && !contains(p.extras, key) {
			unexpected = append(unexpected, key)
		}
	}
	var err error
	if len(unexpected) > 0 {
		err = fmt.Errorf("unexpected param(s) '%s': expecting %s", strings.Join(unexpected, "', '"), p.help())
	}
	missings, _ := p.tree.missings(keys)
	return missings, err
}

type ruleNode interface {
	help() string
	unexpected(string) bool
	missings([]string) ([]string, int)
}

type oneOfNode struct {
	children []ruleNode
}

func (o oneOfNode) help() string {
	var childrenHint []string
	for _, c := range o.children {
		childrenHint = append(childrenHint, c.help())
	}
	return fmt.Sprintf("(%s)", strings.Join(childrenHint, " or "))
}

func (o oneOfNode) unexpected(s string) bool {
	for _, c := range o.children {
		if !c.unexpected(s) {
			return false
		}
	}
	return true
}

func (o oneOfNode) missings(keys []string) ([]string, int) {
	maxFound := -1
	var missings []string
	for _, child := range o.children {
		cMissings, nbFound := child.missings(keys)
		if nbFound > maxFound {
			missings = cMissings
			maxFound = nbFound
		}
	}
	return missings, maxFound
}

type allOfNode struct {
	children []ruleNode
}

func (a allOfNode) help() string {
	var childrenHint []string
	for _, c := range a.children {
		childrenHint = append(childrenHint, c.help())
	}
	return fmt.Sprintf("(%s)", strings.Join(childrenHint, " and "))
}

func (a allOfNode) unexpected(s string) bool {
	for _, c := range a.children {
		if !c.unexpected(s) {
			return false
		}
	}
	return true
}

func (a allOfNode) missings(keys []string) (missings []string, nbFound int) {
	for _, child := range a.children {
		cMissings, cFound := child.missings(keys)
		if len(cMissings) > 0 {
			missings = append(missings, cMissings...)
		} else {
			nbFound += cFound
		}
	}
	return
}

type stringNode struct {
	key string
}

func (k stringNode) help() string {
	return fmt.Sprintf("\"%s\"", k.key)
}

func (k stringNode) unexpected(s string) bool {
	return k.key != s
}

func (k stringNode) missings(keys []string) (missings []string, nbFound int) {
	if contains(keys, k.key) {
		nbFound++
		return
	}
	missings = append(missings, k.key)
	return
}

func oneOf(nodes ...ruleNode) ruleNode {
	return oneOfNode{children: nodes}
}

func allOf(nodes ...ruleNode) ruleNode {
	return allOfNode{children: nodes}
}

func node(key string) ruleNode {
	return stringNode{key}
}

func String(v string) *string {
	return &v
}

func StringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}
