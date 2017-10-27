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

var NewCommandFuncs = map[string]func() interface{}{}

func InitCommands(l *logger.Logger, sess *session.Session) {
	NewCommandFuncs["attachinternetgateway"] = func() interface{} { return NewAttachInternetgateway(l, sess) }
	NewCommandFuncs["attachpolicy"] = func() interface{} { return NewAttachPolicy(l, sess) }
	NewCommandFuncs["attachroutetable"] = func() interface{} { return NewAttachRoutetable(l, sess) }
	NewCommandFuncs["attachsecuritygroup"] = func() interface{} { return NewAttachSecuritygroup(l, sess) }
	NewCommandFuncs["checksecuritygroup"] = func() interface{} { return NewCheckSecuritygroup(l, sess) }
	NewCommandFuncs["creategroup"] = func() interface{} { return NewCreateGroup(l, sess) }
	NewCommandFuncs["createinstance"] = func() interface{} { return NewCreateInstance(l, sess) }
	NewCommandFuncs["createinternetgateway"] = func() interface{} { return NewCreateInternetgateway(l, sess) }
	NewCommandFuncs["createkeypair"] = func() interface{} { return NewCreateKeypair(l, sess) }
	NewCommandFuncs["createpolicy"] = func() interface{} { return NewCreatePolicy(l, sess) }
	NewCommandFuncs["createroute"] = func() interface{} { return NewCreateRoute(l, sess) }
	NewCommandFuncs["createroutetable"] = func() interface{} { return NewCreateRoutetable(l, sess) }
	NewCommandFuncs["createsecuritygroup"] = func() interface{} { return NewCreateSecuritygroup(l, sess) }
	NewCommandFuncs["createsubnet"] = func() interface{} { return NewCreateSubnet(l, sess) }
	NewCommandFuncs["createtag"] = func() interface{} { return NewCreateTag(l, sess) }
	NewCommandFuncs["createvpc"] = func() interface{} { return NewCreateVpc(l, sess) }
	NewCommandFuncs["deletegroup"] = func() interface{} { return NewDeleteGroup(l, sess) }
	NewCommandFuncs["deleteinstance"] = func() interface{} { return NewDeleteInstance(l, sess) }
	NewCommandFuncs["deleteinternetgateway"] = func() interface{} { return NewDeleteInternetgateway(l, sess) }
	NewCommandFuncs["deletekeypair"] = func() interface{} { return NewDeleteKeypair(l, sess) }
	NewCommandFuncs["deletepolicy"] = func() interface{} { return NewDeletePolicy(l, sess) }
	NewCommandFuncs["deleteroute"] = func() interface{} { return NewDeleteRoute(l, sess) }
	NewCommandFuncs["deleteroutetable"] = func() interface{} { return NewDeleteRoutetable(l, sess) }
	NewCommandFuncs["deletesecuritygroup"] = func() interface{} { return NewDeleteSecuritygroup(l, sess) }
	NewCommandFuncs["deletesubnet"] = func() interface{} { return NewDeleteSubnet(l, sess) }
	NewCommandFuncs["deletetag"] = func() interface{} { return NewDeleteTag(l, sess) }
	NewCommandFuncs["deletevpc"] = func() interface{} { return NewDeleteVpc(l, sess) }
	NewCommandFuncs["detachinternetgateway"] = func() interface{} { return NewDetachInternetgateway(l, sess) }
	NewCommandFuncs["detachpolicy"] = func() interface{} { return NewDetachPolicy(l, sess) }
	NewCommandFuncs["detachroutetable"] = func() interface{} { return NewDetachRoutetable(l, sess) }
	NewCommandFuncs["detachsecuritygroup"] = func() interface{} { return NewDetachSecuritygroup(l, sess) }
	NewCommandFuncs["updatepolicy"] = func() interface{} { return NewUpdatePolicy(l, sess) }
	NewCommandFuncs["updatesecuritygroup"] = func() interface{} { return NewUpdateSecuritygroup(l, sess) }
	NewCommandFuncs["updatesubnet"] = func() interface{} { return NewUpdateSubnet(l, sess) }
}

var (
	_ command = &AttachInternetgateway{}
	_ command = &AttachPolicy{}
	_ command = &AttachRoutetable{}
	_ command = &AttachSecuritygroup{}
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
