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
	case "attachalarm":
		return func() interface{} { return NewAttachAlarm(f.Sess, f.Log) }
	case "attachinternetgateway":
		return func() interface{} { return NewAttachInternetgateway(f.Sess, f.Log) }
	case "attachpolicy":
		return func() interface{} { return NewAttachPolicy(f.Sess, f.Log) }
	case "attachroutetable":
		return func() interface{} { return NewAttachRoutetable(f.Sess, f.Log) }
	case "attachsecuritygroup":
		return func() interface{} { return NewAttachSecuritygroup(f.Sess, f.Log) }
	case "attachuser":
		return func() interface{} { return NewAttachUser(f.Sess, f.Log) }
	case "attachvolume":
		return func() interface{} { return NewAttachVolume(f.Sess, f.Log) }
	case "checkinstance":
		return func() interface{} { return NewCheckInstance(f.Sess, f.Log) }
	case "checksecuritygroup":
		return func() interface{} { return NewCheckSecuritygroup(f.Sess, f.Log) }
	case "checkvolume":
		return func() interface{} { return NewCheckVolume(f.Sess, f.Log) }
	case "createaccesskey":
		return func() interface{} { return NewCreateAccesskey(f.Sess, f.Log) }
	case "createalarm":
		return func() interface{} { return NewCreateAlarm(f.Sess, f.Log) }
	case "createappscalingpolicy":
		return func() interface{} { return NewCreateAppscalingpolicy(f.Sess, f.Log) }
	case "createappscalingtarget":
		return func() interface{} { return NewCreateAppscalingtarget(f.Sess, f.Log) }
	case "createbucket":
		return func() interface{} { return NewCreateBucket(f.Sess, f.Log) }
	case "creategroup":
		return func() interface{} { return NewCreateGroup(f.Sess, f.Log) }
	case "createinstance":
		return func() interface{} { return NewCreateInstance(f.Sess, f.Log) }
	case "createinternetgateway":
		return func() interface{} { return NewCreateInternetgateway(f.Sess, f.Log) }
	case "createkeypair":
		return func() interface{} { return NewCreateKeypair(f.Sess, f.Log) }
	case "createpolicy":
		return func() interface{} { return NewCreatePolicy(f.Sess, f.Log) }
	case "createroute":
		return func() interface{} { return NewCreateRoute(f.Sess, f.Log) }
	case "createroutetable":
		return func() interface{} { return NewCreateRoutetable(f.Sess, f.Log) }
	case "createsecuritygroup":
		return func() interface{} { return NewCreateSecuritygroup(f.Sess, f.Log) }
	case "createsubnet":
		return func() interface{} { return NewCreateSubnet(f.Sess, f.Log) }
	case "createtag":
		return func() interface{} { return NewCreateTag(f.Sess, f.Log) }
	case "createuser":
		return func() interface{} { return NewCreateUser(f.Sess, f.Log) }
	case "createvolume":
		return func() interface{} { return NewCreateVolume(f.Sess, f.Log) }
	case "createvpc":
		return func() interface{} { return NewCreateVpc(f.Sess, f.Log) }
	case "createzone":
		return func() interface{} { return NewCreateZone(f.Sess, f.Log) }
	case "deleteaccesskey":
		return func() interface{} { return NewDeleteAccesskey(f.Sess, f.Log) }
	case "deletealarm":
		return func() interface{} { return NewDeleteAlarm(f.Sess, f.Log) }
	case "deleteappscalingpolicy":
		return func() interface{} { return NewDeleteAppscalingpolicy(f.Sess, f.Log) }
	case "deleteappscalingtarget":
		return func() interface{} { return NewDeleteAppscalingtarget(f.Sess, f.Log) }
	case "deletebucket":
		return func() interface{} { return NewDeleteBucket(f.Sess, f.Log) }
	case "deletegroup":
		return func() interface{} { return NewDeleteGroup(f.Sess, f.Log) }
	case "deleteinstance":
		return func() interface{} { return NewDeleteInstance(f.Sess, f.Log) }
	case "deleteinternetgateway":
		return func() interface{} { return NewDeleteInternetgateway(f.Sess, f.Log) }
	case "deletekeypair":
		return func() interface{} { return NewDeleteKeypair(f.Sess, f.Log) }
	case "deletepolicy":
		return func() interface{} { return NewDeletePolicy(f.Sess, f.Log) }
	case "deleteroute":
		return func() interface{} { return NewDeleteRoute(f.Sess, f.Log) }
	case "deleteroutetable":
		return func() interface{} { return NewDeleteRoutetable(f.Sess, f.Log) }
	case "deletesecuritygroup":
		return func() interface{} { return NewDeleteSecuritygroup(f.Sess, f.Log) }
	case "deletesubnet":
		return func() interface{} { return NewDeleteSubnet(f.Sess, f.Log) }
	case "deletetag":
		return func() interface{} { return NewDeleteTag(f.Sess, f.Log) }
	case "deleteuser":
		return func() interface{} { return NewDeleteUser(f.Sess, f.Log) }
	case "deletevolume":
		return func() interface{} { return NewDeleteVolume(f.Sess, f.Log) }
	case "deletevpc":
		return func() interface{} { return NewDeleteVpc(f.Sess, f.Log) }
	case "deletezone":
		return func() interface{} { return NewDeleteZone(f.Sess, f.Log) }
	case "detachalarm":
		return func() interface{} { return NewDetachAlarm(f.Sess, f.Log) }
	case "detachinternetgateway":
		return func() interface{} { return NewDetachInternetgateway(f.Sess, f.Log) }
	case "detachpolicy":
		return func() interface{} { return NewDetachPolicy(f.Sess, f.Log) }
	case "detachroutetable":
		return func() interface{} { return NewDetachRoutetable(f.Sess, f.Log) }
	case "detachsecuritygroup":
		return func() interface{} { return NewDetachSecuritygroup(f.Sess, f.Log) }
	case "detachuser":
		return func() interface{} { return NewDetachUser(f.Sess, f.Log) }
	case "detachvolume":
		return func() interface{} { return NewDetachVolume(f.Sess, f.Log) }
	case "startalarm":
		return func() interface{} { return NewStartAlarm(f.Sess, f.Log) }
	case "stopalarm":
		return func() interface{} { return NewStopAlarm(f.Sess, f.Log) }
	case "updatebucket":
		return func() interface{} { return NewUpdateBucket(f.Sess, f.Log) }
	case "updatepolicy":
		return func() interface{} { return NewUpdatePolicy(f.Sess, f.Log) }
	case "updatesecuritygroup":
		return func() interface{} { return NewUpdateSecuritygroup(f.Sess, f.Log) }
	case "updatesubnet":
		return func() interface{} { return NewUpdateSubnet(f.Sess, f.Log) }
	}
	return nil
}

var (
	_ command = &AttachAlarm{}
	_ command = &AttachInternetgateway{}
	_ command = &AttachPolicy{}
	_ command = &AttachRoutetable{}
	_ command = &AttachSecuritygroup{}
	_ command = &AttachUser{}
	_ command = &AttachVolume{}
	_ command = &CheckInstance{}
	_ command = &CheckSecuritygroup{}
	_ command = &CheckVolume{}
	_ command = &CreateAccesskey{}
	_ command = &CreateAlarm{}
	_ command = &CreateAppscalingpolicy{}
	_ command = &CreateAppscalingtarget{}
	_ command = &CreateBucket{}
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
	_ command = &CreateUser{}
	_ command = &CreateVolume{}
	_ command = &CreateVpc{}
	_ command = &CreateZone{}
	_ command = &DeleteAccesskey{}
	_ command = &DeleteAlarm{}
	_ command = &DeleteAppscalingpolicy{}
	_ command = &DeleteAppscalingtarget{}
	_ command = &DeleteBucket{}
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
	_ command = &DeleteUser{}
	_ command = &DeleteVolume{}
	_ command = &DeleteVpc{}
	_ command = &DeleteZone{}
	_ command = &DetachAlarm{}
	_ command = &DetachInternetgateway{}
	_ command = &DetachPolicy{}
	_ command = &DetachRoutetable{}
	_ command = &DetachSecuritygroup{}
	_ command = &DetachUser{}
	_ command = &DetachVolume{}
	_ command = &StartAlarm{}
	_ command = &StopAlarm{}
	_ command = &UpdateBucket{}
	_ command = &UpdatePolicy{}
	_ command = &UpdateSecuritygroup{}
	_ command = &UpdateSubnet{}
)
