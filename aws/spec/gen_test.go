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
	"github.com/aws/aws-sdk-go/service/iam"
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

	newAttachInternetgateway = func() *AttachInternetgateway {
		return &AttachInternetgateway{api: &mockEc2{}, logger: logger.DiscardLogger}
	}
	newAttachRoutetable    = func() *AttachRoutetable { return &AttachRoutetable{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newAttachSecuritygroup = func() *AttachSecuritygroup {
		return &AttachSecuritygroup{api: &mockEc2{}, logger: logger.DiscardLogger}
	}
	newCheckInstance         = func() *CheckInstance { return &CheckInstance{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newCheckSecuritygroup    = func() *CheckSecuritygroup { return &CheckSecuritygroup{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newCreateInstance        = func() *CreateInstance { return &CreateInstance{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newCreateInternetgateway = func() *CreateInternetgateway {
		return &CreateInternetgateway{api: &mockEc2{}, logger: logger.DiscardLogger}
	}
	newCreateKeypair       = func() *CreateKeypair { return &CreateKeypair{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newCreateRoute         = func() *CreateRoute { return &CreateRoute{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newCreateRoutetable    = func() *CreateRoutetable { return &CreateRoutetable{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newCreateSecuritygroup = func() *CreateSecuritygroup {
		return &CreateSecuritygroup{api: &mockEc2{}, logger: logger.DiscardLogger}
	}
	newCreateSubnet          = func() *CreateSubnet { return &CreateSubnet{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newCreateTag             = func() *CreateTag { return &CreateTag{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newCreateVpc             = func() *CreateVpc { return &CreateVpc{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newDeleteInstance        = func() *DeleteInstance { return &DeleteInstance{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newDeleteInternetgateway = func() *DeleteInternetgateway {
		return &DeleteInternetgateway{api: &mockEc2{}, logger: logger.DiscardLogger}
	}
	newDeleteKeypair       = func() *DeleteKeypair { return &DeleteKeypair{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newDeleteRoute         = func() *DeleteRoute { return &DeleteRoute{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newDeleteRoutetable    = func() *DeleteRoutetable { return &DeleteRoutetable{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newDeleteSecuritygroup = func() *DeleteSecuritygroup {
		return &DeleteSecuritygroup{api: &mockEc2{}, logger: logger.DiscardLogger}
	}
	newDeleteSubnet          = func() *DeleteSubnet { return &DeleteSubnet{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newDeleteTag             = func() *DeleteTag { return &DeleteTag{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newDeleteVpc             = func() *DeleteVpc { return &DeleteVpc{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newDetachInternetgateway = func() *DetachInternetgateway {
		return &DetachInternetgateway{api: &mockEc2{}, logger: logger.DiscardLogger}
	}
	newDetachRoutetable    = func() *DetachRoutetable { return &DetachRoutetable{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newDetachSecuritygroup = func() *DetachSecuritygroup {
		return &DetachSecuritygroup{api: &mockEc2{}, logger: logger.DiscardLogger}
	}
	newUpdateSecuritygroup = func() *UpdateSecuritygroup {
		return &UpdateSecuritygroup{api: &mockEc2{}, logger: logger.DiscardLogger}
	}
	newUpdateSubnet = func() *UpdateSubnet { return &UpdateSubnet{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newAttachPolicy = func() *AttachPolicy { return &AttachPolicy{api: &mockIam{}, logger: logger.DiscardLogger} }
	newCreateGroup  = func() *CreateGroup { return &CreateGroup{api: &mockIam{}, logger: logger.DiscardLogger} }
	newCreatePolicy = func() *CreatePolicy { return &CreatePolicy{api: &mockIam{}, logger: logger.DiscardLogger} }
	newDeleteGroup  = func() *DeleteGroup { return &DeleteGroup{api: &mockIam{}, logger: logger.DiscardLogger} }
	newDeletePolicy = func() *DeletePolicy { return &DeletePolicy{api: &mockIam{}, logger: logger.DiscardLogger} }
	newDetachPolicy = func() *DetachPolicy { return &DetachPolicy{api: &mockIam{}, logger: logger.DiscardLogger} }
	newUpdatePolicy = func() *UpdatePolicy { return &UpdatePolicy{api: &mockIam{}, logger: logger.DiscardLogger} }
)

func init() {

	NewCommandFuncs["attachinternetgateway"] = func() interface{} { return newAttachInternetgateway() }
	NewCommandFuncs["attachroutetable"] = func() interface{} { return newAttachRoutetable() }
	NewCommandFuncs["attachsecuritygroup"] = func() interface{} { return newAttachSecuritygroup() }
	NewCommandFuncs["checkinstance"] = func() interface{} { return newCheckInstance() }
	NewCommandFuncs["checksecuritygroup"] = func() interface{} { return newCheckSecuritygroup() }
	NewCommandFuncs["createinstance"] = func() interface{} { return newCreateInstance() }
	NewCommandFuncs["createinternetgateway"] = func() interface{} { return newCreateInternetgateway() }
	NewCommandFuncs["createkeypair"] = func() interface{} { return newCreateKeypair() }
	NewCommandFuncs["createroute"] = func() interface{} { return newCreateRoute() }
	NewCommandFuncs["createroutetable"] = func() interface{} { return newCreateRoutetable() }
	NewCommandFuncs["createsecuritygroup"] = func() interface{} { return newCreateSecuritygroup() }
	NewCommandFuncs["createsubnet"] = func() interface{} { return newCreateSubnet() }
	NewCommandFuncs["createtag"] = func() interface{} { return newCreateTag() }
	NewCommandFuncs["createvpc"] = func() interface{} { return newCreateVpc() }
	NewCommandFuncs["deleteinstance"] = func() interface{} { return newDeleteInstance() }
	NewCommandFuncs["deleteinternetgateway"] = func() interface{} { return newDeleteInternetgateway() }
	NewCommandFuncs["deletekeypair"] = func() interface{} { return newDeleteKeypair() }
	NewCommandFuncs["deleteroute"] = func() interface{} { return newDeleteRoute() }
	NewCommandFuncs["deleteroutetable"] = func() interface{} { return newDeleteRoutetable() }
	NewCommandFuncs["deletesecuritygroup"] = func() interface{} { return newDeleteSecuritygroup() }
	NewCommandFuncs["deletesubnet"] = func() interface{} { return newDeleteSubnet() }
	NewCommandFuncs["deletetag"] = func() interface{} { return newDeleteTag() }
	NewCommandFuncs["deletevpc"] = func() interface{} { return newDeleteVpc() }
	NewCommandFuncs["detachinternetgateway"] = func() interface{} { return newDetachInternetgateway() }
	NewCommandFuncs["detachroutetable"] = func() interface{} { return newDetachRoutetable() }
	NewCommandFuncs["detachsecuritygroup"] = func() interface{} { return newDetachSecuritygroup() }
	NewCommandFuncs["updatesecuritygroup"] = func() interface{} { return newUpdateSecuritygroup() }
	NewCommandFuncs["updatesubnet"] = func() interface{} { return newUpdateSubnet() }
	NewCommandFuncs["attachpolicy"] = func() interface{} { return newAttachPolicy() }
	NewCommandFuncs["creategroup"] = func() interface{} { return newCreateGroup() }
	NewCommandFuncs["createpolicy"] = func() interface{} { return newCreatePolicy() }
	NewCommandFuncs["deletegroup"] = func() interface{} { return newDeleteGroup() }
	NewCommandFuncs["deletepolicy"] = func() interface{} { return newDeletePolicy() }
	NewCommandFuncs["detachpolicy"] = func() interface{} { return newDetachPolicy() }
	NewCommandFuncs["updatepolicy"] = func() interface{} { return newUpdatePolicy() }
}

type mockEc2 struct {
	ec2iface.EC2API
}

func TestGenAttachInternetgateway(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["attachinternetgateway"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["attachinternetgateway"]
	missings, err := newAttachInternetgateway().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newAttachInternetgateway(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newAttachInternetgateway().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newAttachInternetgateway().Run(genTestsContext["attachinternetgateway"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["attachinternetgateway"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenAttachRoutetable(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["attachroutetable"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["attachroutetable"]
	missings, err := newAttachRoutetable().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newAttachRoutetable(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newAttachRoutetable().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newAttachRoutetable().Run(genTestsContext["attachroutetable"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["attachroutetable"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenAttachSecuritygroup(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["attachsecuritygroup"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["attachsecuritygroup"]
	missings, err := newAttachSecuritygroup().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newAttachSecuritygroup(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newAttachSecuritygroup().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newAttachSecuritygroup().Run(genTestsContext["attachsecuritygroup"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["attachsecuritygroup"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenCheckInstance(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["checkinstance"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["checkinstance"]
	missings, err := newCheckInstance().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newCheckInstance(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newCheckInstance().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newCheckInstance().Run(genTestsContext["checkinstance"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["checkinstance"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenCheckSecuritygroup(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["checksecuritygroup"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["checksecuritygroup"]
	missings, err := newCheckSecuritygroup().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newCheckSecuritygroup(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newCheckSecuritygroup().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newCheckSecuritygroup().Run(genTestsContext["checksecuritygroup"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["checksecuritygroup"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
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
	if errs := newCreateInstance().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newCreateInstance().Run(genTestsContext["createinstance"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["createinstance"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenCreateInternetgateway(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["createinternetgateway"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["createinternetgateway"]
	missings, err := newCreateInternetgateway().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newCreateInternetgateway(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newCreateInternetgateway().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newCreateInternetgateway().Run(genTestsContext["createinternetgateway"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["createinternetgateway"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenCreateKeypair(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["createkeypair"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["createkeypair"]
	missings, err := newCreateKeypair().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newCreateKeypair(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newCreateKeypair().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newCreateKeypair().Run(genTestsContext["createkeypair"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["createkeypair"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenCreateRoute(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["createroute"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["createroute"]
	missings, err := newCreateRoute().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newCreateRoute(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newCreateRoute().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newCreateRoute().Run(genTestsContext["createroute"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["createroute"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenCreateRoutetable(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["createroutetable"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["createroutetable"]
	missings, err := newCreateRoutetable().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newCreateRoutetable(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newCreateRoutetable().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newCreateRoutetable().Run(genTestsContext["createroutetable"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["createroutetable"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenCreateSecuritygroup(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["createsecuritygroup"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["createsecuritygroup"]
	missings, err := newCreateSecuritygroup().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newCreateSecuritygroup(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newCreateSecuritygroup().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newCreateSecuritygroup().Run(genTestsContext["createsecuritygroup"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["createsecuritygroup"]; !reflect.DeepEqual(got, want) {
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
	if errs := newCreateSubnet().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newCreateSubnet().Run(genTestsContext["createsubnet"], params)
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
	if errs := newCreateTag().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newCreateTag().Run(genTestsContext["createtag"], params)
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
	if errs := newCreateVpc().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newCreateVpc().Run(genTestsContext["createvpc"], params)
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
	if errs := newDeleteInstance().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newDeleteInstance().Run(genTestsContext["deleteinstance"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["deleteinstance"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenDeleteInternetgateway(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["deleteinternetgateway"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["deleteinternetgateway"]
	missings, err := newDeleteInternetgateway().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newDeleteInternetgateway(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newDeleteInternetgateway().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newDeleteInternetgateway().Run(genTestsContext["deleteinternetgateway"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["deleteinternetgateway"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenDeleteKeypair(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["deletekeypair"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["deletekeypair"]
	missings, err := newDeleteKeypair().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newDeleteKeypair(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newDeleteKeypair().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newDeleteKeypair().Run(genTestsContext["deletekeypair"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["deletekeypair"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenDeleteRoute(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["deleteroute"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["deleteroute"]
	missings, err := newDeleteRoute().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newDeleteRoute(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newDeleteRoute().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newDeleteRoute().Run(genTestsContext["deleteroute"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["deleteroute"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenDeleteRoutetable(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["deleteroutetable"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["deleteroutetable"]
	missings, err := newDeleteRoutetable().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newDeleteRoutetable(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newDeleteRoutetable().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newDeleteRoutetable().Run(genTestsContext["deleteroutetable"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["deleteroutetable"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenDeleteSecuritygroup(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["deletesecuritygroup"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["deletesecuritygroup"]
	missings, err := newDeleteSecuritygroup().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newDeleteSecuritygroup(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newDeleteSecuritygroup().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newDeleteSecuritygroup().Run(genTestsContext["deletesecuritygroup"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["deletesecuritygroup"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenDeleteSubnet(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["deletesubnet"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["deletesubnet"]
	missings, err := newDeleteSubnet().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newDeleteSubnet(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newDeleteSubnet().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newDeleteSubnet().Run(genTestsContext["deletesubnet"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["deletesubnet"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenDeleteTag(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["deletetag"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["deletetag"]
	missings, err := newDeleteTag().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newDeleteTag(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newDeleteTag().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newDeleteTag().Run(genTestsContext["deletetag"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["deletetag"]; !reflect.DeepEqual(got, want) {
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
	if errs := newDeleteVpc().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newDeleteVpc().Run(genTestsContext["deletevpc"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["deletevpc"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenDetachInternetgateway(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["detachinternetgateway"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["detachinternetgateway"]
	missings, err := newDetachInternetgateway().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newDetachInternetgateway(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newDetachInternetgateway().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newDetachInternetgateway().Run(genTestsContext["detachinternetgateway"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["detachinternetgateway"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenDetachRoutetable(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["detachroutetable"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["detachroutetable"]
	missings, err := newDetachRoutetable().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newDetachRoutetable(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newDetachRoutetable().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newDetachRoutetable().Run(genTestsContext["detachroutetable"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["detachroutetable"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenDetachSecuritygroup(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["detachsecuritygroup"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["detachsecuritygroup"]
	missings, err := newDetachSecuritygroup().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newDetachSecuritygroup(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newDetachSecuritygroup().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newDetachSecuritygroup().Run(genTestsContext["detachsecuritygroup"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["detachsecuritygroup"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenUpdateSecuritygroup(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["updatesecuritygroup"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["updatesecuritygroup"]
	missings, err := newUpdateSecuritygroup().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newUpdateSecuritygroup(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newUpdateSecuritygroup().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newUpdateSecuritygroup().Run(genTestsContext["updatesecuritygroup"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["updatesecuritygroup"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenUpdateSubnet(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["updatesubnet"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["updatesubnet"]
	missings, err := newUpdateSubnet().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newUpdateSubnet(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newUpdateSubnet().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newUpdateSubnet().Run(genTestsContext["updatesubnet"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["updatesubnet"]; !reflect.DeepEqual(got, want) {
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
	if errs := newAttachPolicy().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newAttachPolicy().Run(genTestsContext["attachpolicy"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["attachpolicy"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenCreateGroup(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["creategroup"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["creategroup"]
	missings, err := newCreateGroup().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newCreateGroup(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newCreateGroup().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newCreateGroup().Run(genTestsContext["creategroup"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["creategroup"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenCreatePolicy(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["createpolicy"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["createpolicy"]
	missings, err := newCreatePolicy().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newCreatePolicy(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newCreatePolicy().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newCreatePolicy().Run(genTestsContext["createpolicy"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["createpolicy"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenDeleteGroup(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["deletegroup"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["deletegroup"]
	missings, err := newDeleteGroup().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newDeleteGroup(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newDeleteGroup().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newDeleteGroup().Run(genTestsContext["deletegroup"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["deletegroup"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenDeletePolicy(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["deletepolicy"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["deletepolicy"]
	missings, err := newDeletePolicy().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newDeletePolicy(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newDeletePolicy().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newDeletePolicy().Run(genTestsContext["deletepolicy"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["deletepolicy"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenDetachPolicy(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["detachpolicy"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["detachpolicy"]
	missings, err := newDetachPolicy().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newDetachPolicy(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newDetachPolicy().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newDetachPolicy().Run(genTestsContext["detachpolicy"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["detachpolicy"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
func TestGenUpdatePolicy(t *testing.T) {
	if cleanFn, ok := genTestsCleanupFunc["updatepolicy"]; ok {
		defer cleanFn()
	}
	params := genTestsParams["updatepolicy"]
	missings, err := newUpdatePolicy().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if params, err = convertParamsIfAvailable(newUpdatePolicy(), params); err != nil {
		t.Fatal(err)
	}
	if errs := newUpdatePolicy().ValidateCommand(params, nil); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newUpdatePolicy().Run(genTestsContext["updatepolicy"], params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["updatepolicy"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}

func (m *mockEc2) AttachInternetGateway(input *ec2.AttachInternetGatewayInput) (*ec2.AttachInternetGatewayOutput, error) {
	if got, want := input, genTestsExpected["attachinternetgateway"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["attachinternetgateway"]; ok {
		return outFunc().(*ec2.AttachInternetGatewayOutput), nil
	}
	return nil, nil
}

func (m *mockEc2) AssociateRouteTable(input *ec2.AssociateRouteTableInput) (*ec2.AssociateRouteTableOutput, error) {
	if got, want := input, genTestsExpected["attachroutetable"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["attachroutetable"]; ok {
		return outFunc().(*ec2.AssociateRouteTableOutput), nil
	}
	return nil, nil
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

func (m *mockEc2) CreateInternetGateway(input *ec2.CreateInternetGatewayInput) (*ec2.CreateInternetGatewayOutput, error) {
	if got, want := input, genTestsExpected["createinternetgateway"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["createinternetgateway"]; ok {
		return outFunc().(*ec2.CreateInternetGatewayOutput), nil
	}
	return nil, nil
}

func (m *mockEc2) ImportKeyPair(input *ec2.ImportKeyPairInput) (*ec2.ImportKeyPairOutput, error) {
	if got, want := input, genTestsExpected["createkeypair"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["createkeypair"]; ok {
		return outFunc().(*ec2.ImportKeyPairOutput), nil
	}
	return nil, nil
}

func (m *mockEc2) CreateRoute(input *ec2.CreateRouteInput) (*ec2.CreateRouteOutput, error) {
	if got, want := input, genTestsExpected["createroute"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["createroute"]; ok {
		return outFunc().(*ec2.CreateRouteOutput), nil
	}
	return nil, nil
}

func (m *mockEc2) CreateRouteTable(input *ec2.CreateRouteTableInput) (*ec2.CreateRouteTableOutput, error) {
	if got, want := input, genTestsExpected["createroutetable"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["createroutetable"]; ok {
		return outFunc().(*ec2.CreateRouteTableOutput), nil
	}
	return nil, nil
}

func (m *mockEc2) CreateSecurityGroup(input *ec2.CreateSecurityGroupInput) (*ec2.CreateSecurityGroupOutput, error) {
	if got, want := input, genTestsExpected["createsecuritygroup"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["createsecuritygroup"]; ok {
		return outFunc().(*ec2.CreateSecurityGroupOutput), nil
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

func (m *mockEc2) DeleteInternetGateway(input *ec2.DeleteInternetGatewayInput) (*ec2.DeleteInternetGatewayOutput, error) {
	if got, want := input, genTestsExpected["deleteinternetgateway"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["deleteinternetgateway"]; ok {
		return outFunc().(*ec2.DeleteInternetGatewayOutput), nil
	}
	return nil, nil
}

func (m *mockEc2) DeleteKeyPair(input *ec2.DeleteKeyPairInput) (*ec2.DeleteKeyPairOutput, error) {
	if got, want := input, genTestsExpected["deletekeypair"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["deletekeypair"]; ok {
		return outFunc().(*ec2.DeleteKeyPairOutput), nil
	}
	return nil, nil
}

func (m *mockEc2) DeleteRoute(input *ec2.DeleteRouteInput) (*ec2.DeleteRouteOutput, error) {
	if got, want := input, genTestsExpected["deleteroute"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["deleteroute"]; ok {
		return outFunc().(*ec2.DeleteRouteOutput), nil
	}
	return nil, nil
}

func (m *mockEc2) DeleteRouteTable(input *ec2.DeleteRouteTableInput) (*ec2.DeleteRouteTableOutput, error) {
	if got, want := input, genTestsExpected["deleteroutetable"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["deleteroutetable"]; ok {
		return outFunc().(*ec2.DeleteRouteTableOutput), nil
	}
	return nil, nil
}

func (m *mockEc2) DeleteSecurityGroup(input *ec2.DeleteSecurityGroupInput) (*ec2.DeleteSecurityGroupOutput, error) {
	if got, want := input, genTestsExpected["deletesecuritygroup"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["deletesecuritygroup"]; ok {
		return outFunc().(*ec2.DeleteSecurityGroupOutput), nil
	}
	return nil, nil
}

func (m *mockEc2) DeleteSubnet(input *ec2.DeleteSubnetInput) (*ec2.DeleteSubnetOutput, error) {
	if got, want := input, genTestsExpected["deletesubnet"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["deletesubnet"]; ok {
		return outFunc().(*ec2.DeleteSubnetOutput), nil
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

func (m *mockEc2) DetachInternetGateway(input *ec2.DetachInternetGatewayInput) (*ec2.DetachInternetGatewayOutput, error) {
	if got, want := input, genTestsExpected["detachinternetgateway"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["detachinternetgateway"]; ok {
		return outFunc().(*ec2.DetachInternetGatewayOutput), nil
	}
	return nil, nil
}

func (m *mockEc2) DisassociateRouteTable(input *ec2.DisassociateRouteTableInput) (*ec2.DisassociateRouteTableOutput, error) {
	if got, want := input, genTestsExpected["detachroutetable"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["detachroutetable"]; ok {
		return outFunc().(*ec2.DisassociateRouteTableOutput), nil
	}
	return nil, nil
}

func (m *mockEc2) ModifySubnetAttribute(input *ec2.ModifySubnetAttributeInput) (*ec2.ModifySubnetAttributeOutput, error) {
	if got, want := input, genTestsExpected["updatesubnet"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["updatesubnet"]; ok {
		return outFunc().(*ec2.ModifySubnetAttributeOutput), nil
	}
	return nil, nil
}

func (m *mockIam) CreateGroup(input *iam.CreateGroupInput) (*iam.CreateGroupOutput, error) {
	if got, want := input, genTestsExpected["creategroup"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["creategroup"]; ok {
		return outFunc().(*iam.CreateGroupOutput), nil
	}
	return nil, nil
}

func (m *mockIam) CreatePolicy(input *iam.CreatePolicyInput) (*iam.CreatePolicyOutput, error) {
	if got, want := input, genTestsExpected["createpolicy"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["createpolicy"]; ok {
		return outFunc().(*iam.CreatePolicyOutput), nil
	}
	return nil, nil
}

func (m *mockIam) DeleteGroup(input *iam.DeleteGroupInput) (*iam.DeleteGroupOutput, error) {
	if got, want := input, genTestsExpected["deletegroup"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["deletegroup"]; ok {
		return outFunc().(*iam.DeleteGroupOutput), nil
	}
	return nil, nil
}
