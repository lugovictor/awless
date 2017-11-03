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

	"github.com/aws/aws-sdk-go/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go/service/applicationautoscaling/applicationautoscalingiface"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/aws/aws-sdk-go/service/route53/route53iface"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/wallix/awless/logger"
)

var (
	genTestsParams            = make(map[string]map[string]interface{})
	genTestsContext           = make(map[string]map[string]interface{})
	genTestsExpected          = make(map[string]interface{})
	genTestsOutput            = make(map[string]interface{})
	genTestsCleanupFunc       = make(map[string]func())
	genTestsOutputExtractFunc = make(map[string]func() interface{})

	newCreateAppscalingpolicy = func() *CreateAppscalingpolicy {
		return &CreateAppscalingpolicy{api: &mockApplicationautoscaling{}, logger: logger.DiscardLogger}
	}
	newCreateAppscalingtarget = func() *CreateAppscalingtarget {
		return &CreateAppscalingtarget{api: &mockApplicationautoscaling{}, logger: logger.DiscardLogger}
	}
	newDeleteAppscalingpolicy = func() *DeleteAppscalingpolicy {
		return &DeleteAppscalingpolicy{api: &mockApplicationautoscaling{}, logger: logger.DiscardLogger}
	}
	newDeleteAppscalingtarget = func() *DeleteAppscalingtarget {
		return &DeleteAppscalingtarget{api: &mockApplicationautoscaling{}, logger: logger.DiscardLogger}
	}
	newAttachAlarm           = func() *AttachAlarm { return &AttachAlarm{api: &mockCloudwatch{}, logger: logger.DiscardLogger} }
	newCreateAlarm           = func() *CreateAlarm { return &CreateAlarm{api: &mockCloudwatch{}, logger: logger.DiscardLogger} }
	newDeleteAlarm           = func() *DeleteAlarm { return &DeleteAlarm{api: &mockCloudwatch{}, logger: logger.DiscardLogger} }
	newDetachAlarm           = func() *DetachAlarm { return &DetachAlarm{api: &mockCloudwatch{}, logger: logger.DiscardLogger} }
	newStartAlarm            = func() *StartAlarm { return &StartAlarm{api: &mockCloudwatch{}, logger: logger.DiscardLogger} }
	newStopAlarm             = func() *StopAlarm { return &StopAlarm{api: &mockCloudwatch{}, logger: logger.DiscardLogger} }
	newAttachInternetgateway = func() *AttachInternetgateway {
		return &AttachInternetgateway{api: &mockEc2{}, logger: logger.DiscardLogger}
	}
	newAttachRoutetable    = func() *AttachRoutetable { return &AttachRoutetable{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newAttachSecuritygroup = func() *AttachSecuritygroup {
		return &AttachSecuritygroup{api: &mockEc2{}, logger: logger.DiscardLogger}
	}
	newAttachVolume          = func() *AttachVolume { return &AttachVolume{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newCheckInstance         = func() *CheckInstance { return &CheckInstance{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newCheckSecuritygroup    = func() *CheckSecuritygroup { return &CheckSecuritygroup{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newCheckVolume           = func() *CheckVolume { return &CheckVolume{api: &mockEc2{}, logger: logger.DiscardLogger} }
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
	newCreateVolume          = func() *CreateVolume { return &CreateVolume{api: &mockEc2{}, logger: logger.DiscardLogger} }
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
	newDeleteVolume          = func() *DeleteVolume { return &DeleteVolume{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newDeleteVpc             = func() *DeleteVpc { return &DeleteVpc{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newDetachInternetgateway = func() *DetachInternetgateway {
		return &DetachInternetgateway{api: &mockEc2{}, logger: logger.DiscardLogger}
	}
	newDetachRoutetable    = func() *DetachRoutetable { return &DetachRoutetable{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newDetachSecuritygroup = func() *DetachSecuritygroup {
		return &DetachSecuritygroup{api: &mockEc2{}, logger: logger.DiscardLogger}
	}
	newDetachVolume        = func() *DetachVolume { return &DetachVolume{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newUpdateSecuritygroup = func() *UpdateSecuritygroup {
		return &UpdateSecuritygroup{api: &mockEc2{}, logger: logger.DiscardLogger}
	}
	newUpdateSubnet    = func() *UpdateSubnet { return &UpdateSubnet{api: &mockEc2{}, logger: logger.DiscardLogger} }
	newAttachPolicy    = func() *AttachPolicy { return &AttachPolicy{api: &mockIam{}, logger: logger.DiscardLogger} }
	newCreateAccesskey = func() *CreateAccesskey { return &CreateAccesskey{api: &mockIam{}, logger: logger.DiscardLogger} }
	newCreateGroup     = func() *CreateGroup { return &CreateGroup{api: &mockIam{}, logger: logger.DiscardLogger} }
	newCreatePolicy    = func() *CreatePolicy { return &CreatePolicy{api: &mockIam{}, logger: logger.DiscardLogger} }
	newDeleteAccesskey = func() *DeleteAccesskey { return &DeleteAccesskey{api: &mockIam{}, logger: logger.DiscardLogger} }
	newDeleteGroup     = func() *DeleteGroup { return &DeleteGroup{api: &mockIam{}, logger: logger.DiscardLogger} }
	newDeletePolicy    = func() *DeletePolicy { return &DeletePolicy{api: &mockIam{}, logger: logger.DiscardLogger} }
	newDetachPolicy    = func() *DetachPolicy { return &DetachPolicy{api: &mockIam{}, logger: logger.DiscardLogger} }
	newUpdatePolicy    = func() *UpdatePolicy { return &UpdatePolicy{api: &mockIam{}, logger: logger.DiscardLogger} }
	newCreateZone      = func() *CreateZone { return &CreateZone{api: &mockRoute53{}, logger: logger.DiscardLogger} }
	newDeleteZone      = func() *DeleteZone { return &DeleteZone{api: &mockRoute53{}, logger: logger.DiscardLogger} }
	newCreateBucket    = func() *CreateBucket { return &CreateBucket{api: &mockS3{}, logger: logger.DiscardLogger} }
	newDeleteBucket    = func() *DeleteBucket { return &DeleteBucket{api: &mockS3{}, logger: logger.DiscardLogger} }
	newUpdateBucket    = func() *UpdateBucket { return &UpdateBucket{api: &mockS3{}, logger: logger.DiscardLogger} }
)

type mockApplicationautoscaling struct {
	applicationautoscalingiface.ApplicationAutoScalingAPI
}
type mockCloudwatch struct {
	cloudwatchiface.CloudWatchAPI
}
type mockEc2 struct {
	ec2iface.EC2API
}
type mockIam struct {
	iamiface.IAMAPI
}
type mockRoute53 struct {
	route53iface.Route53API
}
type mockS3 struct {
	s3iface.S3API
}

func (m *mockApplicationautoscaling) PutScalingPolicy(input *applicationautoscaling.PutScalingPolicyInput) (*applicationautoscaling.PutScalingPolicyOutput, error) {
	if got, want := input, genTestsExpected["createappscalingpolicy"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["createappscalingpolicy"]; ok {
		return outFunc().(*applicationautoscaling.PutScalingPolicyOutput), nil
	}
	return nil, nil
}

func (m *mockApplicationautoscaling) RegisterScalableTarget(input *applicationautoscaling.RegisterScalableTargetInput) (*applicationautoscaling.RegisterScalableTargetOutput, error) {
	if got, want := input, genTestsExpected["createappscalingtarget"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["createappscalingtarget"]; ok {
		return outFunc().(*applicationautoscaling.RegisterScalableTargetOutput), nil
	}
	return nil, nil
}

func (m *mockApplicationautoscaling) DeleteScalingPolicy(input *applicationautoscaling.DeleteScalingPolicyInput) (*applicationautoscaling.DeleteScalingPolicyOutput, error) {
	if got, want := input, genTestsExpected["deleteappscalingpolicy"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["deleteappscalingpolicy"]; ok {
		return outFunc().(*applicationautoscaling.DeleteScalingPolicyOutput), nil
	}
	return nil, nil
}

func (m *mockApplicationautoscaling) DeregisterScalableTarget(input *applicationautoscaling.DeregisterScalableTargetInput) (*applicationautoscaling.DeregisterScalableTargetOutput, error) {
	if got, want := input, genTestsExpected["deleteappscalingtarget"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["deleteappscalingtarget"]; ok {
		return outFunc().(*applicationautoscaling.DeregisterScalableTargetOutput), nil
	}
	return nil, nil
}

func (m *mockCloudwatch) PutMetricAlarm(input *cloudwatch.PutMetricAlarmInput) (*cloudwatch.PutMetricAlarmOutput, error) {
	if got, want := input, genTestsExpected["createalarm"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["createalarm"]; ok {
		return outFunc().(*cloudwatch.PutMetricAlarmOutput), nil
	}
	return nil, nil
}

func (m *mockCloudwatch) DeleteAlarms(input *cloudwatch.DeleteAlarmsInput) (*cloudwatch.DeleteAlarmsOutput, error) {
	if got, want := input, genTestsExpected["deletealarm"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["deletealarm"]; ok {
		return outFunc().(*cloudwatch.DeleteAlarmsOutput), nil
	}
	return nil, nil
}

func (m *mockCloudwatch) EnableAlarmActions(input *cloudwatch.EnableAlarmActionsInput) (*cloudwatch.EnableAlarmActionsOutput, error) {
	if got, want := input, genTestsExpected["startalarm"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["startalarm"]; ok {
		return outFunc().(*cloudwatch.EnableAlarmActionsOutput), nil
	}
	return nil, nil
}

func (m *mockCloudwatch) DisableAlarmActions(input *cloudwatch.DisableAlarmActionsInput) (*cloudwatch.DisableAlarmActionsOutput, error) {
	if got, want := input, genTestsExpected["stopalarm"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["stopalarm"]; ok {
		return outFunc().(*cloudwatch.DisableAlarmActionsOutput), nil
	}
	return nil, nil
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

func (m *mockEc2) AttachVolume(input *ec2.AttachVolumeInput) (*ec2.VolumeAttachment, error) {
	if got, want := input, genTestsExpected["attachvolume"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["attachvolume"]; ok {
		return outFunc().(*ec2.VolumeAttachment), nil
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

func (m *mockEc2) CreateVolume(input *ec2.CreateVolumeInput) (*ec2.Volume, error) {
	if got, want := input, genTestsExpected["createvolume"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["createvolume"]; ok {
		return outFunc().(*ec2.Volume), nil
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

func (m *mockEc2) DeleteVolume(input *ec2.DeleteVolumeInput) (*ec2.DeleteVolumeOutput, error) {
	if got, want := input, genTestsExpected["deletevolume"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["deletevolume"]; ok {
		return outFunc().(*ec2.DeleteVolumeOutput), nil
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

func (m *mockEc2) DetachVolume(input *ec2.DetachVolumeInput) (*ec2.VolumeAttachment, error) {
	if got, want := input, genTestsExpected["detachvolume"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["detachvolume"]; ok {
		return outFunc().(*ec2.VolumeAttachment), nil
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

func (m *mockIam) CreateAccessKey(input *iam.CreateAccessKeyInput) (*iam.CreateAccessKeyOutput, error) {
	if got, want := input, genTestsExpected["createaccesskey"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["createaccesskey"]; ok {
		return outFunc().(*iam.CreateAccessKeyOutput), nil
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

func (m *mockIam) DeleteAccessKey(input *iam.DeleteAccessKeyInput) (*iam.DeleteAccessKeyOutput, error) {
	if got, want := input, genTestsExpected["deleteaccesskey"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["deleteaccesskey"]; ok {
		return outFunc().(*iam.DeleteAccessKeyOutput), nil
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

func (m *mockIam) DeletePolicy(input *iam.DeletePolicyInput) (*iam.DeletePolicyOutput, error) {
	if got, want := input, genTestsExpected["deletepolicy"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["deletepolicy"]; ok {
		return outFunc().(*iam.DeletePolicyOutput), nil
	}
	return nil, nil
}

func (m *mockIam) CreatePolicyVersion(input *iam.CreatePolicyVersionInput) (*iam.CreatePolicyVersionOutput, error) {
	if got, want := input, genTestsExpected["updatepolicy"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["updatepolicy"]; ok {
		return outFunc().(*iam.CreatePolicyVersionOutput), nil
	}
	return nil, nil
}

func (m *mockRoute53) CreateHostedZone(input *route53.CreateHostedZoneInput) (*route53.CreateHostedZoneOutput, error) {
	if got, want := input, genTestsExpected["createzone"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["createzone"]; ok {
		return outFunc().(*route53.CreateHostedZoneOutput), nil
	}
	return nil, nil
}

func (m *mockRoute53) DeleteHostedZone(input *route53.DeleteHostedZoneInput) (*route53.DeleteHostedZoneOutput, error) {
	if got, want := input, genTestsExpected["deletezone"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["deletezone"]; ok {
		return outFunc().(*route53.DeleteHostedZoneOutput), nil
	}
	return nil, nil
}

func (m *mockS3) CreateBucket(input *s3.CreateBucketInput) (*s3.CreateBucketOutput, error) {
	if got, want := input, genTestsExpected["createbucket"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["createbucket"]; ok {
		return outFunc().(*s3.CreateBucketOutput), nil
	}
	return nil, nil
}

func (m *mockS3) DeleteBucket(input *s3.DeleteBucketInput) (*s3.DeleteBucketOutput, error) {
	if got, want := input, genTestsExpected["deletebucket"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["deletebucket"]; ok {
		return outFunc().(*s3.DeleteBucketOutput), nil
	}
	return nil, nil
}
