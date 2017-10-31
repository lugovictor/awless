package awsat

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/wallix/awless/aws/spec"
	"github.com/wallix/awless/logger"
	"github.com/wallix/awless/template"
)

type ATBuilder struct {
	template    string
	cmdResult   *string
	expectCalls map[string]int
	expectInput map[string]interface{}
	mock        mock
}

func Template(template string) *ATBuilder {
	return &ATBuilder{template: template, expectCalls: make(map[string]int), expectInput: make(map[string]interface{})}
}

func (b *ATBuilder) ExpectCommandResult(key string) *ATBuilder {
	b.cmdResult = &key
	return b
}

func (b *ATBuilder) ExpectCalls(expects ...string) *ATBuilder {
	for _, expect := range expects {
		b.expectCalls[expect]++
	}
	return b
}

func (b *ATBuilder) ExpectInput(call string, input interface{}) *ATBuilder {
	b.expectInput[call] = input
	return b
}

func (b *ATBuilder) Run(t *testing.T) {
	b.mock.SetInputs(b.expectInput)
	b.mock.SetTesting(t)
	awsspec.NewCommandFuncs["createtag"] = func() interface{} {
		cmd := new(awsspec.CreateTag)
		cmd.SetApi(b.mock.(ec2iface.EC2API))
		cmd.SetLogger(logger.DiscardLogger)
		return cmd
	}

	tpl, err := template.Parse(b.template)
	if err != nil {
		t.Fatal(err)
	}

	env := template.NewEnv()
	env.Lookuper = func(tokens ...string) interface{} {
		newCommandFunc, ok := MockCommands[strings.Join(tokens, "")]
		if !ok {
			return nil
		}
		return newCommandFunc(b.mock)
	}
	env.IsNewRunner = true
	compiled, env, err := template.Compile(tpl, env, template.NewRunnerCompileMode)
	if err != nil {
		t.Fatal(err)
	}

	ran, err := compiled.DoRun(env)
	if err != nil {
		t.Fatal(err)
	}
	if len(b.expectCalls) > 0 {
		if got, want := b.mock.Calls(), b.expectCalls; !reflect.DeepEqual(got, want) {
			t.Fatalf("got %#v, want %#v", got, want)
		}
	}
	if b.cmdResult != nil {
		if got, want := fmt.Sprint(ran.CommandNodesIterator()[0].Result()), StringValue(b.cmdResult); got != want {
			t.Fatalf("got %s, want %s", got, want)
		}
	}
}

func (b *ATBuilder) Mock(i mock) *ATBuilder {
	b.mock = i
	return b
}

func StringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}
