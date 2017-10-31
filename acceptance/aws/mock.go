package awsat

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
	"github.com/wallix/awless/aws/spec"
	"github.com/wallix/awless/logger"
)

var MockCommands map[string]func(interface{}) interface{}

func init() {
	MockCommands = make(map[string]func(interface{}) interface{})
	MockCommands["createinstance"] = func(api interface{}) interface{} {
		cmd := new(awsspec.CreateInstance)
		cmd.SetApi(api.(ec2iface.EC2API))
		cmd.SetLogger(logger.DiscardLogger)
		return cmd
	}
	MockCommands["createpolicy"] = func(api interface{}) interface{} {
		cmd := new(awsspec.CreatePolicy)
		cmd.SetApi(api.(iamiface.IAMAPI))
		cmd.SetLogger(logger.DiscardLogger)
		return cmd
	}
}

type mock interface {
	Calls() map[string]int
	SetInputs(map[string]interface{})
	SetTesting(*testing.T)
}

// type ec2Mock struct {
// 	basicMock
// 	ec2iface.EC2API
// 	RunInstancesFunc      func(input *ec2.RunInstancesInput) (*ec2.Reservation, error)
// 	CreateTagsRequestFunc func(input *ec2.CreateTagsInput) (req *request.Request, output *ec2.CreateTagsOutput)
// }
//
// func (m *ec2Mock) RunInstances(input *ec2.RunInstancesInput) (*ec2.Reservation, error) {
// 	m.addCall("RunInstances")
// 	m.verifyInput("RunInstances", input)
// 	return m.RunInstancesFunc(input)
// }
//
// func (m *ec2Mock) CreateTagsRequest(input *ec2.CreateTagsInput) (*request.Request, *ec2.CreateTagsOutput) {
// 	m.addCall("CreateTagsRequest")
// 	m.verifyInput("CreateTagsRequest", input)
// 	return m.CreateTagsRequestFunc(input)
// }
//
// type iamMock struct {
// 	basicMock
// 	iamiface.IAMAPI
// 	CreatePolicyFunc func(input *iam.CreatePolicyInput) (*iam.CreatePolicyOutput, error)
// }
//
// func (m *iamMock) CreatePolicy(input *iam.CreatePolicyInput) (*iam.CreatePolicyOutput, error) {
// 	m.addCall("CreatePolicy")
// 	m.verifyInput("CreatePolicy", input)
// 	return m.CreatePolicyFunc(input)
// }

type basicMock struct {
	t         *testing.T
	calls     map[string]int
	expInputs map[string]interface{}
}

func (m *basicMock) addCall(call string) {
	if m.calls == nil {
		m.calls = make(map[string]int)
	}
	m.calls[call]++
}

func (m *basicMock) Calls() map[string]int {
	return m.calls
}

func (m *basicMock) SetTesting(t *testing.T) {
	m.t = t
}

func (m *basicMock) SetInputs(inputs map[string]interface{}) {
	m.expInputs = inputs
}

func (m *basicMock) verifyInput(call string, got interface{}) {
	m.t.Helper()
	if m.expInputs == nil {
		return
	}
	if want := m.expInputs[call]; !reflect.DeepEqual(want, got) {
		m.t.Fatalf("got %#v, want %#v", got, want)
	}
}
