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
	"github.com/aws/aws-sdk-go/service/acm/acmiface"
	"github.com/aws/aws-sdk-go/service/applicationautoscaling/applicationautoscalingiface"
	"github.com/aws/aws-sdk-go/service/autoscaling/autoscalingiface"
	"github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
	"github.com/aws/aws-sdk-go/service/cloudfront/cloudfrontiface"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/aws/aws-sdk-go/service/ecs/ecsiface"
	"github.com/aws/aws-sdk-go/service/elbv2/elbv2iface"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
	"github.com/aws/aws-sdk-go/service/lambda/lambdaiface"
	"github.com/aws/aws-sdk-go/service/rds/rdsiface"
	"github.com/aws/aws-sdk-go/service/route53/route53iface"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
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
	case "attachcontainertask":
		return func() interface{} {
			cmd := awsspec.NewAttachContainertask(nil)
			cmd.SetApi(f.Mock.(ecsiface.ECSAPI))
			return cmd
		}
	case "attachelasticip":
		return func() interface{} {
			cmd := awsspec.NewAttachElasticip(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "attachinstance":
		return func() interface{} {
			cmd := awsspec.NewAttachInstance(nil)
			cmd.SetApi(f.Mock.(elbv2iface.ELBV2API))
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
	case "attachuser":
		return func() interface{} {
			cmd := awsspec.NewAttachUser(nil)
			cmd.SetApi(f.Mock.(iamiface.IAMAPI))
			return cmd
		}
	case "attachvolume":
		return func() interface{} {
			cmd := awsspec.NewAttachVolume(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "checkcertificate":
		return func() interface{} {
			cmd := awsspec.NewCheckCertificate(nil)
			cmd.SetApi(f.Mock.(acmiface.ACMAPI))
			return cmd
		}
	case "checkdatabase":
		return func() interface{} {
			cmd := awsspec.NewCheckDatabase(nil)
			cmd.SetApi(f.Mock.(rdsiface.RDSAPI))
			return cmd
		}
	case "checkdistribution":
		return func() interface{} {
			cmd := awsspec.NewCheckDistribution(nil)
			cmd.SetApi(f.Mock.(cloudfrontiface.CloudFrontAPI))
			return cmd
		}
	case "checkinstance":
		return func() interface{} {
			cmd := awsspec.NewCheckInstance(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "checkscalinggroup":
		return func() interface{} {
			cmd := awsspec.NewCheckScalinggroup(nil)
			cmd.SetApi(f.Mock.(autoscalingiface.AutoScalingAPI))
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
	case "copyimage":
		return func() interface{} {
			cmd := awsspec.NewCopyImage(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "copysnapshot":
		return func() interface{} {
			cmd := awsspec.NewCopySnapshot(nil)
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
	case "createcertificate":
		return func() interface{} {
			cmd := awsspec.NewCreateCertificate(nil)
			cmd.SetApi(f.Mock.(acmiface.ACMAPI))
			return cmd
		}
	case "createcontainercluster":
		return func() interface{} {
			cmd := awsspec.NewCreateContainercluster(nil)
			cmd.SetApi(f.Mock.(ecsiface.ECSAPI))
			return cmd
		}
	case "createdatabase":
		return func() interface{} {
			cmd := awsspec.NewCreateDatabase(nil)
			cmd.SetApi(f.Mock.(rdsiface.RDSAPI))
			return cmd
		}
	case "createdbsubnetgroup":
		return func() interface{} {
			cmd := awsspec.NewCreateDbsubnetgroup(nil)
			cmd.SetApi(f.Mock.(rdsiface.RDSAPI))
			return cmd
		}
	case "createdistribution":
		return func() interface{} {
			cmd := awsspec.NewCreateDistribution(nil)
			cmd.SetApi(f.Mock.(cloudfrontiface.CloudFrontAPI))
			return cmd
		}
	case "createelasticip":
		return func() interface{} {
			cmd := awsspec.NewCreateElasticip(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "createfunction":
		return func() interface{} {
			cmd := awsspec.NewCreateFunction(nil)
			cmd.SetApi(f.Mock.(lambdaiface.LambdaAPI))
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
	case "createscalinggroup":
		return func() interface{} {
			cmd := awsspec.NewCreateScalinggroup(nil)
			cmd.SetApi(f.Mock.(autoscalingiface.AutoScalingAPI))
			return cmd
		}
	case "createscalingpolicy":
		return func() interface{} {
			cmd := awsspec.NewCreateScalingpolicy(nil)
			cmd.SetApi(f.Mock.(autoscalingiface.AutoScalingAPI))
			return cmd
		}
	case "createsecuritygroup":
		return func() interface{} {
			cmd := awsspec.NewCreateSecuritygroup(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "createsnapshot":
		return func() interface{} {
			cmd := awsspec.NewCreateSnapshot(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "createstack":
		return func() interface{} {
			cmd := awsspec.NewCreateStack(nil)
			cmd.SetApi(f.Mock.(cloudformationiface.CloudFormationAPI))
			return cmd
		}
	case "createsubnet":
		return func() interface{} {
			cmd := awsspec.NewCreateSubnet(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "createsubscription":
		return func() interface{} {
			cmd := awsspec.NewCreateSubscription(nil)
			cmd.SetApi(f.Mock.(snsiface.SNSAPI))
			return cmd
		}
	case "createtag":
		return func() interface{} {
			cmd := awsspec.NewCreateTag(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "createtargetgroup":
		return func() interface{} {
			cmd := awsspec.NewCreateTargetgroup(nil)
			cmd.SetApi(f.Mock.(elbv2iface.ELBV2API))
			return cmd
		}
	case "createtopic":
		return func() interface{} {
			cmd := awsspec.NewCreateTopic(nil)
			cmd.SetApi(f.Mock.(snsiface.SNSAPI))
			return cmd
		}
	case "createuser":
		return func() interface{} {
			cmd := awsspec.NewCreateUser(nil)
			cmd.SetApi(f.Mock.(iamiface.IAMAPI))
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
	case "deletecertificate":
		return func() interface{} {
			cmd := awsspec.NewDeleteCertificate(nil)
			cmd.SetApi(f.Mock.(acmiface.ACMAPI))
			return cmd
		}
	case "deletecontainercluster":
		return func() interface{} {
			cmd := awsspec.NewDeleteContainercluster(nil)
			cmd.SetApi(f.Mock.(ecsiface.ECSAPI))
			return cmd
		}
	case "deletecontainertask":
		return func() interface{} {
			cmd := awsspec.NewDeleteContainertask(nil)
			cmd.SetApi(f.Mock.(ecsiface.ECSAPI))
			return cmd
		}
	case "deletedatabase":
		return func() interface{} {
			cmd := awsspec.NewDeleteDatabase(nil)
			cmd.SetApi(f.Mock.(rdsiface.RDSAPI))
			return cmd
		}
	case "deletedbsubnetgroup":
		return func() interface{} {
			cmd := awsspec.NewDeleteDbsubnetgroup(nil)
			cmd.SetApi(f.Mock.(rdsiface.RDSAPI))
			return cmd
		}
	case "deletedistribution":
		return func() interface{} {
			cmd := awsspec.NewDeleteDistribution(nil)
			cmd.SetApi(f.Mock.(cloudfrontiface.CloudFrontAPI))
			return cmd
		}
	case "deleteelasticip":
		return func() interface{} {
			cmd := awsspec.NewDeleteElasticip(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "deletefunction":
		return func() interface{} {
			cmd := awsspec.NewDeleteFunction(nil)
			cmd.SetApi(f.Mock.(lambdaiface.LambdaAPI))
			return cmd
		}
	case "deletegroup":
		return func() interface{} {
			cmd := awsspec.NewDeleteGroup(nil)
			cmd.SetApi(f.Mock.(iamiface.IAMAPI))
			return cmd
		}
	case "deleteimage":
		return func() interface{} {
			cmd := awsspec.NewDeleteImage(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
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
	case "deletescalinggroup":
		return func() interface{} {
			cmd := awsspec.NewDeleteScalinggroup(nil)
			cmd.SetApi(f.Mock.(autoscalingiface.AutoScalingAPI))
			return cmd
		}
	case "deletescalingpolicy":
		return func() interface{} {
			cmd := awsspec.NewDeleteScalingpolicy(nil)
			cmd.SetApi(f.Mock.(autoscalingiface.AutoScalingAPI))
			return cmd
		}
	case "deletesecuritygroup":
		return func() interface{} {
			cmd := awsspec.NewDeleteSecuritygroup(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "deletesnapshot":
		return func() interface{} {
			cmd := awsspec.NewDeleteSnapshot(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "deletestack":
		return func() interface{} {
			cmd := awsspec.NewDeleteStack(nil)
			cmd.SetApi(f.Mock.(cloudformationiface.CloudFormationAPI))
			return cmd
		}
	case "deletesubnet":
		return func() interface{} {
			cmd := awsspec.NewDeleteSubnet(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "deletesubscription":
		return func() interface{} {
			cmd := awsspec.NewDeleteSubscription(nil)
			cmd.SetApi(f.Mock.(snsiface.SNSAPI))
			return cmd
		}
	case "deletetag":
		return func() interface{} {
			cmd := awsspec.NewDeleteTag(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "deletetargetgroup":
		return func() interface{} {
			cmd := awsspec.NewDeleteTargetgroup(nil)
			cmd.SetApi(f.Mock.(elbv2iface.ELBV2API))
			return cmd
		}
	case "deletetopic":
		return func() interface{} {
			cmd := awsspec.NewDeleteTopic(nil)
			cmd.SetApi(f.Mock.(snsiface.SNSAPI))
			return cmd
		}
	case "deleteuser":
		return func() interface{} {
			cmd := awsspec.NewDeleteUser(nil)
			cmd.SetApi(f.Mock.(iamiface.IAMAPI))
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
	case "detachcontainertask":
		return func() interface{} {
			cmd := awsspec.NewDetachContainertask(nil)
			cmd.SetApi(f.Mock.(ecsiface.ECSAPI))
			return cmd
		}
	case "detachelasticip":
		return func() interface{} {
			cmd := awsspec.NewDetachElasticip(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "detachinstance":
		return func() interface{} {
			cmd := awsspec.NewDetachInstance(nil)
			cmd.SetApi(f.Mock.(elbv2iface.ELBV2API))
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
	case "detachuser":
		return func() interface{} {
			cmd := awsspec.NewDetachUser(nil)
			cmd.SetApi(f.Mock.(iamiface.IAMAPI))
			return cmd
		}
	case "detachvolume":
		return func() interface{} {
			cmd := awsspec.NewDetachVolume(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "importimage":
		return func() interface{} {
			cmd := awsspec.NewImportImage(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "startalarm":
		return func() interface{} {
			cmd := awsspec.NewStartAlarm(nil)
			cmd.SetApi(f.Mock.(cloudwatchiface.CloudWatchAPI))
			return cmd
		}
	case "startcontainertask":
		return func() interface{} {
			cmd := awsspec.NewStartContainertask(nil)
			cmd.SetApi(f.Mock.(ecsiface.ECSAPI))
			return cmd
		}
	case "startinstance":
		return func() interface{} {
			cmd := awsspec.NewStartInstance(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "stopalarm":
		return func() interface{} {
			cmd := awsspec.NewStopAlarm(nil)
			cmd.SetApi(f.Mock.(cloudwatchiface.CloudWatchAPI))
			return cmd
		}
	case "stopcontainertask":
		return func() interface{} {
			cmd := awsspec.NewStopContainertask(nil)
			cmd.SetApi(f.Mock.(ecsiface.ECSAPI))
			return cmd
		}
	case "stopinstance":
		return func() interface{} {
			cmd := awsspec.NewStopInstance(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "updatebucket":
		return func() interface{} {
			cmd := awsspec.NewUpdateBucket(nil)
			cmd.SetApi(f.Mock.(s3iface.S3API))
			return cmd
		}
	case "updatecontainertask":
		return func() interface{} {
			cmd := awsspec.NewUpdateContainertask(nil)
			cmd.SetApi(f.Mock.(ecsiface.ECSAPI))
			return cmd
		}
	case "updatedistribution":
		return func() interface{} {
			cmd := awsspec.NewUpdateDistribution(nil)
			cmd.SetApi(f.Mock.(cloudfrontiface.CloudFrontAPI))
			return cmd
		}
	case "updateinstance":
		return func() interface{} {
			cmd := awsspec.NewUpdateInstance(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "updatepolicy":
		return func() interface{} {
			cmd := awsspec.NewUpdatePolicy(nil)
			cmd.SetApi(f.Mock.(iamiface.IAMAPI))
			return cmd
		}
	case "updatescalinggroup":
		return func() interface{} {
			cmd := awsspec.NewUpdateScalinggroup(nil)
			cmd.SetApi(f.Mock.(autoscalingiface.AutoScalingAPI))
			return cmd
		}
	case "updatesecuritygroup":
		return func() interface{} {
			cmd := awsspec.NewUpdateSecuritygroup(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "updatestack":
		return func() interface{} {
			cmd := awsspec.NewUpdateStack(nil)
			cmd.SetApi(f.Mock.(cloudformationiface.CloudFormationAPI))
			return cmd
		}
	case "updatesubnet":
		return func() interface{} {
			cmd := awsspec.NewUpdateSubnet(nil)
			cmd.SetApi(f.Mock.(ec2iface.EC2API))
			return cmd
		}
	case "updatetargetgroup":
		return func() interface{} {
			cmd := awsspec.NewUpdateTargetgroup(nil)
			cmd.SetApi(f.Mock.(elbv2iface.ELBV2API))
			return cmd
		}
	}
	return nil
}
