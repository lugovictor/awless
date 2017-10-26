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

	newAttachInternetgateway = func() *AttachInternetgateway {
		return &AttachInternetgateway{api: &mockEc2{}, logger: logger.DiscardLogger}
	}
	newAttachRoutetable    = func() *AttachRoutetable { return &AttachRoutetable{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newAttachSecuritygroup = func() *AttachSecuritygroup {
		return &AttachSecuritygroup{api: &mockEc2{}, logger: logger.DiscardLogger}
	}
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
)

func init() {

	NewCommandFuncs["attachinternetgateway"] = func() interface{} { return newAttachInternetgateway() }
	NewCommandFuncs["attachroutetable"] = func() interface{} { return newAttachRoutetable() }
	NewCommandFuncs["attachsecuritygroup"] = func() interface{} { return newAttachSecuritygroup() }
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
	NewCommandFuncs["deletevpc"] = func() interface{} { return newDeleteVpc() }
	NewCommandFuncs["detachinternetgateway"] = func() interface{} { return newDetachInternetgateway() }
	NewCommandFuncs["detachroutetable"] = func() interface{} { return newDetachRoutetable() }
	NewCommandFuncs["detachsecuritygroup"] = func() interface{} { return newDetachSecuritygroup() }
	NewCommandFuncs["updatesecuritygroup"] = func() interface{} { return newUpdateSecuritygroup() }
	NewCommandFuncs["updatesubnet"] = func() interface{} { return newUpdateSubnet() }
	NewCommandFuncs["attachpolicy"] = func() interface{} { return newAttachPolicy() }
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
	if errs := newAttachInternetgateway().ValidateCommand(genTestsParams["attachinternetgateway"]); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newAttachInternetgateway().Run(genTestsContext["attachinternetgateway"], genTestsParams["attachinternetgateway"])
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
	if errs := newAttachRoutetable().ValidateCommand(genTestsParams["attachroutetable"]); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newAttachRoutetable().Run(genTestsContext["attachroutetable"], genTestsParams["attachroutetable"])
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
	if errs := newAttachSecuritygroup().ValidateCommand(genTestsParams["attachsecuritygroup"]); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newAttachSecuritygroup().Run(genTestsContext["attachsecuritygroup"], genTestsParams["attachsecuritygroup"])
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["attachsecuritygroup"]; !reflect.DeepEqual(got, want) {
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
	if errs := newCheckSecuritygroup().ValidateCommand(genTestsParams["checksecuritygroup"]); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newCheckSecuritygroup().Run(genTestsContext["checksecuritygroup"], genTestsParams["checksecuritygroup"])
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
	if errs := newCreateInternetgateway().ValidateCommand(genTestsParams["createinternetgateway"]); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newCreateInternetgateway().Run(genTestsContext["createinternetgateway"], genTestsParams["createinternetgateway"])
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
	if errs := newCreateKeypair().ValidateCommand(genTestsParams["createkeypair"]); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newCreateKeypair().Run(genTestsContext["createkeypair"], genTestsParams["createkeypair"])
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
	if errs := newCreateRoute().ValidateCommand(genTestsParams["createroute"]); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newCreateRoute().Run(genTestsContext["createroute"], genTestsParams["createroute"])
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
	if errs := newCreateRoutetable().ValidateCommand(genTestsParams["createroutetable"]); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newCreateRoutetable().Run(genTestsContext["createroutetable"], genTestsParams["createroutetable"])
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
	if errs := newCreateSecuritygroup().ValidateCommand(genTestsParams["createsecuritygroup"]); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newCreateSecuritygroup().Run(genTestsContext["createsecuritygroup"], genTestsParams["createsecuritygroup"])
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
	if errs := newDeleteInternetgateway().ValidateCommand(genTestsParams["deleteinternetgateway"]); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newDeleteInternetgateway().Run(genTestsContext["deleteinternetgateway"], genTestsParams["deleteinternetgateway"])
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
	if errs := newDeleteKeypair().ValidateCommand(genTestsParams["deletekeypair"]); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newDeleteKeypair().Run(genTestsContext["deletekeypair"], genTestsParams["deletekeypair"])
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
	if errs := newDeleteRoute().ValidateCommand(genTestsParams["deleteroute"]); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newDeleteRoute().Run(genTestsContext["deleteroute"], genTestsParams["deleteroute"])
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
	if errs := newDeleteRoutetable().ValidateCommand(genTestsParams["deleteroutetable"]); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newDeleteRoutetable().Run(genTestsContext["deleteroutetable"], genTestsParams["deleteroutetable"])
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
	if errs := newDeleteSecuritygroup().ValidateCommand(genTestsParams["deletesecuritygroup"]); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newDeleteSecuritygroup().Run(genTestsContext["deletesecuritygroup"], genTestsParams["deletesecuritygroup"])
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
	if errs := newDeleteSubnet().ValidateCommand(genTestsParams["deletesubnet"]); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newDeleteSubnet().Run(genTestsContext["deletesubnet"], genTestsParams["deletesubnet"])
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, genTestsOutput["deletesubnet"]; !reflect.DeepEqual(got, want) {
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
	if errs := newDetachInternetgateway().ValidateCommand(genTestsParams["detachinternetgateway"]); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newDetachInternetgateway().Run(genTestsContext["detachinternetgateway"], genTestsParams["detachinternetgateway"])
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
	if errs := newDetachRoutetable().ValidateCommand(genTestsParams["detachroutetable"]); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newDetachRoutetable().Run(genTestsContext["detachroutetable"], genTestsParams["detachroutetable"])
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
	if errs := newDetachSecuritygroup().ValidateCommand(genTestsParams["detachsecuritygroup"]); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newDetachSecuritygroup().Run(genTestsContext["detachsecuritygroup"], genTestsParams["detachsecuritygroup"])
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
	if errs := newUpdateSecuritygroup().ValidateCommand(genTestsParams["updatesecuritygroup"]); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newUpdateSecuritygroup().Run(genTestsContext["updatesecuritygroup"], genTestsParams["updatesecuritygroup"])
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
	if errs := newUpdateSubnet().ValidateCommand(genTestsParams["updatesubnet"]); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newUpdateSubnet().Run(genTestsContext["updatesubnet"], genTestsParams["updatesubnet"])
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
