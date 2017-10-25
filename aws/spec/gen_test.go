/* Copyright 2017 WALLIX

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
package awsspec

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
	"github.com/wallix/awless/logger"
)

var (
	genTestsParams            = make(map[string]map[string]interface{})
	genTestsContext           = make(map[string]map[string]interface{})
	genTestsExpected          = make(map[string]interface{})
	genTestsOutput            = make(map[string]interface{})
	genTestsCleanupFunc       = make(map[string]func())
	genTestsOutputExtractFunc = make(map[string]func() interface{})

	newCreateInstance = func() *CreateInstance { return &CreateInstance{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newCreateSubnet   = func() *CreateSubnet { return &CreateSubnet{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newCreateTag      = func() *CreateTag { return &CreateTag{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newCreateVpc      = func() *CreateVpc { return &CreateVpc{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newDeleteInstance = func() *DeleteInstance { return &DeleteInstance{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newDeleteVpc      = func() *DeleteVpc { return &DeleteVpc{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newAttachPolicy   = func() *AttachPolicy { return &AttachPolicy{api: &mockIam{}, logger: logger.DiscardLogger} }
)

func init() {

	NewCommandFuncs["createinstance"] = func() interface{} { return newCreateInstance() }
	NewCommandFuncs["createsubnet"] = func() interface{} { return newCreateSubnet() }
	NewCommandFuncs["createtag"] = func() interface{} { return newCreateTag() }
	NewCommandFuncs["createvpc"] = func() interface{} { return newCreateVpc() }
	NewCommandFuncs["deleteinstance"] = func() interface{} { return newDeleteInstance() }
	NewCommandFuncs["deletevpc"] = func() interface{} { return newDeleteVpc() }
	NewCommandFuncs["attachpolicy"] = func() interface{} { return newAttachPolicy() }
}

type mockEc2 struct {
	ec2iface.EC2API
}

func TestGenCreateInstance(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["createinstance"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["createinstance"]
	missings, err := newCreateInstance().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newCreateInstance(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newCreateInstance().ValidateCommand(genTestsParams["createinstance"]); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newCreateInstance().Run(genTestsContext["createinstance"], genTestsParams["createinstance"])
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["createinstance"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenCreateSubnet(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["createsubnet"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["createsubnet"]
	missings, err := newCreateSubnet().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newCreateSubnet(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newCreateSubnet().ValidateCommand(genTestsParams["createsubnet"]); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newCreateSubnet().Run(genTestsContext["createsubnet"], genTestsParams["createsubnet"])
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["createsubnet"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenCreateTag(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["createtag"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["createtag"]
	missings, err := newCreateTag().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newCreateTag(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newCreateTag().ValidateCommand(genTestsParams["createtag"]); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newCreateTag().Run(genTestsContext["createtag"], genTestsParams["createtag"])
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["createtag"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenCreateVpc(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["createvpc"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["createvpc"]
	missings, err := newCreateVpc().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newCreateVpc(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newCreateVpc().ValidateCommand(genTestsParams["createvpc"]); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newCreateVpc().Run(genTestsContext["createvpc"], genTestsParams["createvpc"])
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["createvpc"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenDeleteInstance(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["deleteinstance"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["deleteinstance"]
	missings, err := newDeleteInstance().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newDeleteInstance(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newDeleteInstance().ValidateCommand(genTestsParams["deleteinstance"]); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newDeleteInstance().Run(genTestsContext["deleteinstance"], genTestsParams["deleteinstance"])
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["deleteinstance"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenDeleteVpc(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["deletevpc"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["deletevpc"]
	missings, err := newDeleteVpc().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newDeleteVpc(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newDeleteVpc().ValidateCommand(genTestsParams["deletevpc"]); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newDeleteVpc().Run(genTestsContext["deletevpc"], genTestsParams["deletevpc"])
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["deletevpc"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}

type mockIam struct {
	iamiface.IAMAPI
}

func TestGenAttachPolicy(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["attachpolicy"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["attachpolicy"]
	missings, err := newAttachPolicy().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newAttachPolicy(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newAttachPolicy().ValidateCommand(genTestsParams["attachpolicy"]); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newAttachPolicy().Run(genTestsContext["attachpolicy"], genTestsParams["attachpolicy"])
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["attachpolicy"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}

func (m *mockEc2) RunInstances(input *ec2.RunInstancesInput) (*ec2.Reservation, error) {
	if got, want := input, genTestsExpected["createinstance"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["createinstance"]; ok {
		return outFunc().(*ec2.Reservation), nil
	}
	return nil, nil
}

func (m *mockEc2) CreateSubnet(input *ec2.CreateSubnetInput) (*ec2.CreateSubnetOutput, error) {
	if got, want := input, genTestsExpected["createsubnet"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["createsubnet"]; ok {
		return outFunc().(*ec2.CreateSubnetOutput), nil
	}
	return nil, nil
}

func (m *mockEc2) CreateVpc(input *ec2.CreateVpcInput) (*ec2.CreateVpcOutput, error) {
	if got, want := input, genTestsExpected["createvpc"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["createvpc"]; ok {
		return outFunc().(*ec2.CreateVpcOutput), nil
	}
	return nil, nil
}

func (m *mockEc2) TerminateInstances(input *ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error) {
	if got, want := input, genTestsExpected["deleteinstance"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["deleteinstance"]; ok {
		return outFunc().(*ec2.TerminateInstancesOutput), nil
	}
	return nil, nil
}

func (m *mockEc2) DeleteVpc(input *ec2.DeleteVpcInput) (*ec2.DeleteVpcOutput, error) {
	if got, want := input, genTestsExpected["deletevpc"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["deletevpc"]; ok {
		return outFunc().(*ec2.DeleteVpcOutput), nil
	}
	return nil, nil
}
