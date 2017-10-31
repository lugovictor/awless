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
	template      string
	inputs        map[string]interface{}
	inputToVerify map[string]interface{}
	tplResult     *string
	mock          interface{}
	t             *testing.T
}

func Command(template string) *ATBuilder {
	return &ATBuilder{template: template, inputs: make(map[string]interface{})}
}

func (b *ATBuilder) Input(key string, input interface{}) *ATBuilder {
	b.inputs[key] = input
	return b
}

func (b *ATBuilder) VerifyInput(key string, i interface{}) {
	b.t.Helper()
	v, ok := b.inputs[key]
	if !ok {
		b.t.Fatalf("no input to verify for key '%s'", key)
	}
	if got, want := i, v; !reflect.DeepEqual(got, want) {
		b.t.Fatalf("%s: got %#v, want %#v", key, got, want)
	}
}

func (b *ATBuilder) VerifyTemplateResult(key string) *ATBuilder {
	b.tplResult = &key
	return b
}

func (b *ATBuilder) Run(t *testing.T) {
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
	b.t = t

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
	if b.tplResult != nil {
		if got, want := fmt.Sprint(ran.CommandNodesIterator()[0].Result()), StringValue(b.tplResult); got != want {
			b.t.Fatalf("got %s, want %s", got, want)
		}
	}
}

func (b *ATBuilder) Mock(i interface{}) *ATBuilder {
	b.mock = i
	return b
}

func StringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}
