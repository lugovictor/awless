package awsat

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
	"github.com/wallix/awless/aws/spec"
	"github.com/wallix/awless/logger"
)

type AcceptanceFactory struct {
	Mock interface{}
}

func NewAcceptanceFactory(mock interface{}) *AcceptanceFactory {
	return &AcceptanceFactory{Mock: mock}
}

func (f *AcceptanceFactory) Build(key string) func() interface{} {
	switch key {
	case "createinstance":
		return func() interface{} {
			cmd := new(awsspec.CreateInstance)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			cmd.SetLogger(logger.DiscardLogger)
			return cmd
		}
	case "createtag":
		return func() interface{} {
			cmd := new(awsspec.CreateTag)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			cmd.SetLogger(logger.DiscardLogger)
			return cmd
		}
	case "deleteinstance":
		return func() interface{} {
			cmd := new(awsspec.DeleteInstance)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			cmd.SetLogger(logger.DiscardLogger)
			return cmd
		}
	case "checkinstance":
		return func() interface{} {
			cmd := new(awsspec.CheckInstance)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			cmd.SetLogger(logger.DiscardLogger)
			return cmd
		}
	case "createpolicy":
		return func() interface{} {
			cmd := new(awsspec.CreatePolicy)
			cmd.SetApi(f.Mock.(iamiface.IAMAPI))
			cmd.SetLogger(logger.DiscardLogger)
			return cmd
		}
	}
	return nil
}

type mock interface {
	Calls() map[string]int
	SetInputs(map[string]interface{})
	SetTesting(*testing.T)
}

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
