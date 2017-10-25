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
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/aws/aws-sdk-go/service/elbv2/elbv2iface"
	"github.com/wallix/awless/logger"
)

type CreateInstance struct {
	_             string `action: "create" entity: "instance" awsAPI: "ec2" awsCall: "RunInstances" awsInput: "RunInstancesInput" awsOutput: "Reservation" awsDryRun: ""`
	logger        *logger.Logger
	api           ec2iface.EC2API
	Image         *string   `awsName: "ImageId" awsType: "awsstr" templateName: "image" required: ""`
	Count         *int64    `awsName: "MaxCount" awsType: "awsint64" templateName: "count" required: ""`
	Count         *int64    `awsName: "MinCount" awsType: "awsint64" templateName: "count" required: ""`
	Type          *string   `awsName: "InstanceType" awsType: "awsstr" templateName: "type" required: ""`
	Subnet        *string   `awsName: "SubnetId" awsType: "awsstr" templateName: "subnet" required: ""`
	Name          *struct{} `awsName: "Name" templateName: "name" required: ""`
	Keypair       *string   `awsName: "KeyName" awsType: "awsstr" templateName: "keypair"`
	Ip            *string   `awsName: "PrivateIpAddress" awsType: "awsstr" templateName: "ip"`
	Userdata      *string   `awsName: "UserData" awsType: "awsfiletobase64" templateName: "userdata"`
	Securitygroup *[]string `awsName: "SecurityGroupIds" awsType: "awsstringslice" templateName: "securitygroup"`
	Lock          *bool     `awsName: "DisableApiTermination" awsType: "awsbool" templateName: "lock"`
	Role          *string   `awsName: "IamInstanceProfile.Name" awsType: "awsstr" templateName: "role"`
}
type UpdateInstance struct {
	_      string `action: "update" entity: "instance" awsAPI: "ec2" awsCall: "ModifyInstanceAttribute" awsInput: "ModifyInstanceAttributeInput" awsOutput: "ModifyInstanceAttributeOutput" awsDryRun: ""`
	logger *logger.Logger
	api    ec2iface.EC2API
	Id     *string `awsName: "InstanceId" awsType: "awsstr" templateName: "id" required: ""`
	Type   *string `awsName: "InstanceType.Value" awsType: "awsstr" templateName: "type"`
	Lock   *bool   `awsName: "DisableApiTermination" awsType: "awsboolattribute" templateName: "lock"`
}
type DeleteInstance struct {
	_      string `action: "delete" entity: "instance" awsAPI: "ec2" awsCall: "TerminateInstances" awsInput: "TerminateInstancesInput" awsOutput: "TerminateInstancesOutput" awsDryRun: ""`
	logger *logger.Logger
	api    ec2iface.EC2API
	Id     *[]string `awsName: "InstanceIds" awsType: "awsstringslice" templateName: "id" required: ""`
}
type StartInstance struct {
	_      string `action: "start" entity: "instance" awsAPI: "ec2" awsCall: "StartInstances" awsInput: "StartInstancesInput" awsOutput: "StartInstancesOutput" awsDryRun: ""`
	logger *logger.Logger
	api    ec2iface.EC2API
	Id     *[]string `awsName: "InstanceIds" awsType: "awsstringslice" templateName: "id" required: ""`
}
type StopInstance struct {
	_      string `action: "stop" entity: "instance" awsAPI: "ec2" awsCall: "StopInstances" awsInput: "StopInstancesInput" awsOutput: "StopInstancesOutput" awsDryRun: ""`
	logger *logger.Logger
	api    ec2iface.EC2API
	Id     *[]string `awsName: "InstanceIds" awsType: "awsstringslice" templateName: "id" required: ""`
}
type CheckInstance struct {
	_       string `action: "check" entity: "instance" awsAPI: "ec2"`
	logger  *logger.Logger
	api     ec2iface.EC2API
	Id      *struct{} `templateName: "id" required: ""`
	State   *struct{} `templateName: "state" required: ""`
	Timeout *struct{} `templateName: "timeout" required: ""`
}
type AttachInstance struct {
	_           string `action: "attach" entity: "instance" awsAPI: "elbv2" awsCall: "RegisterTargets" awsInput: "RegisterTargetsInput" awsOutput: "RegisterTargetsOutput"`
	logger      *logger.Logger
	api         elbv2iface.ELBV2API
	Targetgroup *string   `awsName: "TargetGroupArn" awsType: "awsstr" templateName: "targetgroup" required: ""`
	Id          *struct{} `awsName: "Targets[0]Id" awsType: "awsslicestruct" templateName: "id" required: ""`
	Port        *int64    `awsName: "Targets[0]Port" awsType: "awsslicestructint64" templateName: "port"`
}
type DetachInstance struct {
	_           string `action: "detach" entity: "instance" awsAPI: "elbv2" awsCall: "DeregisterTargets" awsInput: "DeregisterTargetsInput" awsOutput: "DeregisterTargetsOutput"`
	logger      *logger.Logger
	api         elbv2iface.ELBV2API
	Targetgroup *string   `awsName: "TargetGroupArn" awsType: "awsstr" templateName: "targetgroup" required: ""`
	Id          *struct{} `awsName: "Targets[0]Id" awsType: "awsslicestruct" templateName: "id" required: ""`
}
