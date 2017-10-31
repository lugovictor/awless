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
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/wallix/awless/logger"
)

type Factory interface {
	Build(key string) func() interface{}
}

var CommandFactory Factory

type AWSFactory struct {
	Log  *logger.Logger
	Sess *session.Session
}

func (f *AWSFactory) Build(key string) func() interface{} {
	switch key {
	case "attachinternetgateway":
		return func() interface{} { return NewAttachInternetgateway(f.Log, f.Sess) }
	case "attachpolicy":
		return func() interface{} { return NewAttachPolicy(f.Log, f.Sess) }
	case "attachroutetable":
		return func() interface{} { return NewAttachRoutetable(f.Log, f.Sess) }
	case "attachsecuritygroup":
		return func() interface{} { return NewAttachSecuritygroup(f.Log, f.Sess) }
	case "checkinstance":
		return func() interface{} { return NewCheckInstance(f.Log, f.Sess) }
	case "checksecuritygroup":
		return func() interface{} { return NewCheckSecuritygroup(f.Log, f.Sess) }
	case "creategroup":
		return func() interface{} { return NewCreateGroup(f.Log, f.Sess) }
	case "createinstance":
		return func() interface{} { return NewCreateInstance(f.Log, f.Sess) }
	case "createinternetgateway":
		return func() interface{} { return NewCreateInternetgateway(f.Log, f.Sess) }
	case "createkeypair":
		return func() interface{} { return NewCreateKeypair(f.Log, f.Sess) }
	case "createpolicy":
		return func() interface{} { return NewCreatePolicy(f.Log, f.Sess) }
	case "createroute":
		return func() interface{} { return NewCreateRoute(f.Log, f.Sess) }
	case "createroutetable":
		return func() interface{} { return NewCreateRoutetable(f.Log, f.Sess) }
	case "createsecuritygroup":
		return func() interface{} { return NewCreateSecuritygroup(f.Log, f.Sess) }
	case "createsubnet":
		return func() interface{} { return NewCreateSubnet(f.Log, f.Sess) }
	case "createtag":
		return func() interface{} { return NewCreateTag(f.Log, f.Sess) }
	case "createvpc":
		return func() interface{} { return NewCreateVpc(f.Log, f.Sess) }
	case "deletegroup":
		return func() interface{} { return NewDeleteGroup(f.Log, f.Sess) }
	case "deleteinstance":
		return func() interface{} { return NewDeleteInstance(f.Log, f.Sess) }
	case "deleteinternetgateway":
		return func() interface{} { return NewDeleteInternetgateway(f.Log, f.Sess) }
	case "deletekeypair":
		return func() interface{} { return NewDeleteKeypair(f.Log, f.Sess) }
	case "deletepolicy":
		return func() interface{} { return NewDeletePolicy(f.Log, f.Sess) }
	case "deleteroute":
		return func() interface{} { return NewDeleteRoute(f.Log, f.Sess) }
	case "deleteroutetable":
		return func() interface{} { return NewDeleteRoutetable(f.Log, f.Sess) }
	case "deletesecuritygroup":
		return func() interface{} { return NewDeleteSecuritygroup(f.Log, f.Sess) }
	case "deletesubnet":
		return func() interface{} { return NewDeleteSubnet(f.Log, f.Sess) }
	case "deletetag":
		return func() interface{} { return NewDeleteTag(f.Log, f.Sess) }
	case "deletevpc":
		return func() interface{} { return NewDeleteVpc(f.Log, f.Sess) }
	case "detachinternetgateway":
		return func() interface{} { return NewDetachInternetgateway(f.Log, f.Sess) }
	case "detachpolicy":
		return func() interface{} { return NewDetachPolicy(f.Log, f.Sess) }
	case "detachroutetable":
		return func() interface{} { return NewDetachRoutetable(f.Log, f.Sess) }
	case "detachsecuritygroup":
		return func() interface{} { return NewDetachSecuritygroup(f.Log, f.Sess) }
	case "updatepolicy":
		return func() interface{} { return NewUpdatePolicy(f.Log, f.Sess) }
	case "updatesecuritygroup":
		return func() interface{} { return NewUpdateSecuritygroup(f.Log, f.Sess) }
	case "updatesubnet":
		return func() interface{} { return NewUpdateSubnet(f.Log, f.Sess) }
	}
	return nil
}

var (
	_ command = &AttachInternetgateway{}
	_ command = &AttachPolicy{}
	_ command = &AttachRoutetable{}
	_ command = &AttachSecuritygroup{}
	_ command = &CheckInstance{}
	_ command = &CheckSecuritygroup{}
	_ command = &CreateGroup{}
	_ command = &CreateInstance{}
	_ command = &CreateInternetgateway{}
	_ command = &CreateKeypair{}
	_ command = &CreatePolicy{}
	_ command = &CreateRoute{}
	_ command = &CreateRoutetable{}
	_ command = &CreateSecuritygroup{}
	_ command = &CreateSubnet{}
	_ command = &CreateTag{}
	_ command = &CreateVpc{}
	_ command = &DeleteGroup{}
	_ command = &DeleteInstance{}
	_ command = &DeleteInternetgateway{}
	_ command = &DeleteKeypair{}
	_ command = &DeletePolicy{}
	_ command = &DeleteRoute{}
	_ command = &DeleteRoutetable{}
	_ command = &DeleteSecuritygroup{}
	_ command = &DeleteSubnet{}
	_ command = &DeleteTag{}
	_ command = &DeleteVpc{}
	_ command = &DetachInternetgateway{}
	_ command = &DetachPolicy{}
	_ command = &DetachRoutetable{}
	_ command = &DetachSecuritygroup{}
	_ command = &UpdatePolicy{}
	_ command = &UpdateSecuritygroup{}
	_ command = &UpdateSubnet{}
)
