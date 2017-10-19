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
	rule   ruleNode
	extras []string
}

func (p *paramRule) hint() string {
	if len(p.extras) == 0 {
		return p.rule.hint()
	}

	return fmt.Sprintf("(%s OR %s)", p.rule.hint(), strings.Join(p.extras, " OR "))
}

func (p *paramRule) unexpected(keys []string) []string {
	var unexpected []string
	for _, key := range keys {
		if p.rule.unexpected(key) && !contains(p.extras, key) {
			unexpected = append(unexpected, key)
		}
	}
	return unexpected
}

func (p *paramRule) verify(keys []string) ([]string, error) {
	var err error
	if unexpected := p.unexpected(keys); len(unexpected) > 0 {
		err = fmt.Errorf("unexpected param key(s) '%s' - expected params: %s", strings.Join(unexpected, "', '"), p.hint())
	}
	missings, _ := p.rule.eval(keys)
	return missings, err
}

type ruleNode interface {
	hint() string
	unexpected(string) bool
	eval([]string) ([]string, int)
}

type oneOfNode struct {
	children []ruleNode
}

func (o oneOfNode) hint() string {
	var childrenHint []string
	for _, c := range o.children {
		childrenHint = append(childrenHint, c.hint())
	}
	return fmt.Sprintf("(%s)", strings.Join(childrenHint, " OR "))
}

func (o oneOfNode) unexpected(s string) bool {
	for _, c := range o.children {
		if !c.unexpected(s) {
			return false
		}
	}
	return true
}

func (o oneOfNode) eval(keys []string) ([]string, int) {
	var maxProcessed int
	minSize := len(keys)
	var minMissings []string
	for _, child := range o.children {
		missings, nbFound := child.eval(keys)
		if nbFound > maxProcessed && len(missings) < minSize ||
			nbFound == maxProcessed && len(missings) < minSize ||
			nbFound > maxProcessed && len(missings) == minSize {
			minMissings = missings
			minSize = len(missings)
			maxProcessed = nbFound
		}
	}
	return minMissings, maxProcessed
}

type allOfNode struct {
	children []ruleNode
}

func (a allOfNode) hint() string {
	var childrenHint []string
	for _, c := range a.children {
		childrenHint = append(childrenHint, c.hint())
	}
	return fmt.Sprintf("(%s)", strings.Join(childrenHint, " AND "))
}

func (a allOfNode) unexpected(s string) bool {
	for _, c := range a.children {
		if !c.unexpected(s) {
			return false
		}
	}
	return true
}

func (a allOfNode) eval(keys []string) (missings []string, nbFound int) {
	for _, child := range a.children {
		cMissings, cFound := child.eval(keys)
		if len(cMissings) > 0 {
			missings = append(missings, cMissings...)
		}
		nbFound += cFound
	}
	return
}

type stringNode struct {
	key string
}

func (k stringNode) hint() string {
	return fmt.Sprintf("\"%s\"", k.key)
}

func (k stringNode) unexpected(s string) bool {
	return k.key != s
}

func (k stringNode) eval(keys []string) (missings []string, nbFound int) {
	if contains(keys, k.key) {
		nbFound++
		return
	}
	missings = append(missings, k.key)
	return
}

func oneOf(ifaces ...*paramRule) *paramRule {
	var children []ruleNode
	for _, i := range ifaces {
		children = append(children, i.rule)
	}
	return &paramRule{rule: oneOfNode{children: children}}
}

func allOf(ifaces ...*paramRule) *paramRule {
	var children []ruleNode
	for _, i := range ifaces {
		children = append(children, i.rule)
	}
	return &paramRule{rule: allOfNode{children: children}}
}

func node(key string) *paramRule {
	return &paramRule{rule: stringNode{key}}
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
