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
package awsat

import (
	"github.com/aws/aws-sdk-go/service/applicationautoscaling/applicationautoscalingiface"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
	"github.com/aws/aws-sdk-go/service/route53/route53iface"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/wallix/awless/aws/spec"
)

type AcceptanceFactory struct {
	Mock interface{}
}

func NewAcceptanceFactory(mock interface{}) *AcceptanceFactory {
	return &AcceptanceFactory{Mock: mock}
}

func (f *AcceptanceFactory) Build(key string) func() interface{} {
	switch key {
	case "attachalarm":
		return func() interface{} {
			cmd := awsspec.NewAttachAlarm(nil)
			cmd.SetApi(f.Mock.(cloudwatchiface.CloudWatchAPI))
			return cmd
		}
	case "attachinternetgateway":
		return func() interface{} {
			cmd := awsspec.NewAttachInternetgateway(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "attachpolicy":
		return func() interface{} {
			cmd := awsspec.NewAttachPolicy(nil)
			cmd.SetApi(f.Mock.(iamiface.IAMAPI))
			return cmd
		}
	case "attachroutetable":
		return func() interface{} {
			cmd := awsspec.NewAttachRoutetable(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "attachsecuritygroup":
		return func() interface{} {
			cmd := awsspec.NewAttachSecuritygroup(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "attachvolume":
		return func() interface{} {
			cmd := awsspec.NewAttachVolume(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "checkinstance":
		return func() interface{} {
			cmd := awsspec.NewCheckInstance(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "checksecuritygroup":
		return func() interface{} {
			cmd := awsspec.NewCheckSecuritygroup(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "checkvolume":
		return func() interface{} {
			cmd := awsspec.NewCheckVolume(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "createaccesskey":
		return func() interface{} {
			cmd := awsspec.NewCreateAccesskey(nil)
			cmd.SetApi(f.Mock.(iamiface.IAMAPI))
			return cmd
		}
	case "createalarm":
		return func() interface{} {
			cmd := awsspec.NewCreateAlarm(nil)
			cmd.SetApi(f.Mock.(cloudwatchiface.CloudWatchAPI))
			return cmd
		}
	case "createappscalingpolicy":
		return func() interface{} {
			cmd := awsspec.NewCreateAppscalingpolicy(nil)
			cmd.SetApi(f.Mock.(applicationautoscalingiface.ApplicationAutoScalingAPI))
			return cmd
		}
	case "createappscalingtarget":
		return func() interface{} {
			cmd := awsspec.NewCreateAppscalingtarget(nil)
			cmd.SetApi(f.Mock.(applicationautoscalingiface.ApplicationAutoScalingAPI))
			return cmd
		}
	case "createbucket":
		return func() interface{} {
			cmd := awsspec.NewCreateBucket(nil)
			cmd.SetApi(f.Mock.(s3iface.S3API))
			return cmd
		}
	case "creategroup":
		return func() interface{} {
			cmd := awsspec.NewCreateGroup(nil)
			cmd.SetApi(f.Mock.(iamiface.IAMAPI))
			return cmd
		}
	case "createinstance":
		return func() interface{} {
			cmd := awsspec.NewCreateInstance(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "createinternetgateway":
		return func() interface{} {
			cmd := awsspec.NewCreateInternetgateway(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "createkeypair":
		return func() interface{} {
			cmd := awsspec.NewCreateKeypair(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "createpolicy":
		return func() interface{} {
			cmd := awsspec.NewCreatePolicy(nil)
			cmd.SetApi(f.Mock.(iamiface.IAMAPI))
			return cmd
		}
	case "createroute":
		return func() interface{} {
			cmd := awsspec.NewCreateRoute(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "createroutetable":
		return func() interface{} {
			cmd := awsspec.NewCreateRoutetable(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "createsecuritygroup":
		return func() interface{} {
			cmd := awsspec.NewCreateSecuritygroup(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "createsubnet":
		return func() interface{} {
			cmd := awsspec.NewCreateSubnet(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "createtag":
		return func() interface{} {
			cmd := awsspec.NewCreateTag(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "createvolume":
		return func() interface{} {
			cmd := awsspec.NewCreateVolume(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "createvpc":
		return func() interface{} {
			cmd := awsspec.NewCreateVpc(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "createzone":
		return func() interface{} {
			cmd := awsspec.NewCreateZone(nil)
			cmd.SetApi(f.Mock.(route53iface.Route53API))
			return cmd
		}
	case "deleteaccesskey":
		return func() interface{} {
			cmd := awsspec.NewDeleteAccesskey(nil)
			cmd.SetApi(f.Mock.(iamiface.IAMAPI))
			return cmd
		}
	case "deletealarm":
		return func() interface{} {
			cmd := awsspec.NewDeleteAlarm(nil)
			cmd.SetApi(f.Mock.(cloudwatchiface.CloudWatchAPI))
			return cmd
		}
	case "deleteappscalingpolicy":
		return func() interface{} {
			cmd := awsspec.NewDeleteAppscalingpolicy(nil)
			cmd.SetApi(f.Mock.(applicationautoscalingiface.ApplicationAutoScalingAPI))
			return cmd
		}
	case "deleteappscalingtarget":
		return func() interface{} {
			cmd := awsspec.NewDeleteAppscalingtarget(nil)
			cmd.SetApi(f.Mock.(applicationautoscalingiface.ApplicationAutoScalingAPI))
			return cmd
		}
	case "deletebucket":
		return func() interface{} {
			cmd := awsspec.NewDeleteBucket(nil)
			cmd.SetApi(f.Mock.(s3iface.S3API))
			return cmd
		}
	case "deletegroup":
		return func() interface{} {
			cmd := awsspec.NewDeleteGroup(nil)
			cmd.SetApi(f.Mock.(iamiface.IAMAPI))
			return cmd
		}
	case "deleteinstance":
		return func() interface{} {
			cmd := awsspec.NewDeleteInstance(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "deleteinternetgateway":
		return func() interface{} {
			cmd := awsspec.NewDeleteInternetgateway(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "deletekeypair":
		return func() interface{} {
			cmd := awsspec.NewDeleteKeypair(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "deletepolicy":
		return func() interface{} {
			cmd := awsspec.NewDeletePolicy(nil)
			cmd.SetApi(f.Mock.(iamiface.IAMAPI))
			return cmd
		}
	case "deleteroute":
		return func() interface{} {
			cmd := awsspec.NewDeleteRoute(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "deleteroutetable":
		return func() interface{} {
			cmd := awsspec.NewDeleteRoutetable(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "deletesecuritygroup":
		return func() interface{} {
			cmd := awsspec.NewDeleteSecuritygroup(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "deletesubnet":
		return func() interface{} {
			cmd := awsspec.NewDeleteSubnet(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "deletetag":
		return func() interface{} {
			cmd := awsspec.NewDeleteTag(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "deletevolume":
		return func() interface{} {
			cmd := awsspec.NewDeleteVolume(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "deletevpc":
		return func() interface{} {
			cmd := awsspec.NewDeleteVpc(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "deletezone":
		return func() interface{} {
			cmd := awsspec.NewDeleteZone(nil)
			cmd.SetApi(f.Mock.(route53iface.Route53API))
			return cmd
		}
	case "detachalarm":
		return func() interface{} {
			cmd := awsspec.NewDetachAlarm(nil)
			cmd.SetApi(f.Mock.(cloudwatchiface.CloudWatchAPI))
			return cmd
		}
	case "detachinternetgateway":
		return func() interface{} {
			cmd := awsspec.NewDetachInternetgateway(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "detachpolicy":
		return func() interface{} {
			cmd := awsspec.NewDetachPolicy(nil)
			cmd.SetApi(f.Mock.(iamiface.IAMAPI))
			return cmd
		}
	case "detachroutetable":
		return func() interface{} {
			cmd := awsspec.NewDetachRoutetable(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "detachsecuritygroup":
		return func() interface{} {
			cmd := awsspec.NewDetachSecuritygroup(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "detachvolume":
		return func() interface{} {
			cmd := awsspec.NewDetachVolume(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "startalarm":
		return func() interface{} {
			cmd := awsspec.NewStartAlarm(nil)
			cmd.SetApi(f.Mock.(cloudwatchiface.CloudWatchAPI))
			return cmd
		}
	case "stopalarm":
		return func() interface{} {
			cmd := awsspec.NewStopAlarm(nil)
			cmd.SetApi(f.Mock.(cloudwatchiface.CloudWatchAPI))
			return cmd
		}
	case "updatebucket":
		return func() interface{} {
			cmd := awsspec.NewUpdateBucket(nil)
			cmd.SetApi(f.Mock.(s3iface.S3API))
			return cmd
		}
	case "updatepolicy":
		return func() interface{} {
			cmd := awsspec.NewUpdatePolicy(nil)
			cmd.SetApi(f.Mock.(iamiface.IAMAPI))
			return cmd
		}
	case "updatesecuritygroup":
		return func() interface{} {
			cmd := awsspec.NewUpdateSecuritygroup(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "updatesubnet":
		return func() interface{} {
			cmd := awsspec.NewUpdateSubnet(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	}
	return nil
}
