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
	"github.com/wallix/awless/logger"
)

type CreateNetworkinterface struct {
	_              string `action: "create" entity: "networkinterface" awsAPI: "ec2" awsCall: "CreateNetworkInterface" awsInput: "CreateNetworkInterfaceInput" awsOutput: "CreateNetworkInterfaceOutput" awsDryRun: ""`
	logger         *logger.Logger
	api            ec2iface.EC2API
	Subnet         *string   `awsName: "SubnetId" awsType: "awsstr" templateName: "subnet" required: ""`
	Description    *string   `awsName: "Description" awsType: "awsstr" templateName: "description"`
	Securitygroups *[]string `awsName: "Groups" awsType: "awsstringslice" templateName: "securitygroups"`
	Privateip      *string   `awsName: "PrivateIpAddress" awsType: "awsstr" templateName: "privateip"`
}
type DeleteNetworkinterface struct {
	_      string `action: "delete" entity: "networkinterface" awsAPI: "ec2" awsCall: "DeleteNetworkInterface" awsInput: "DeleteNetworkInterfaceInput" awsOutput: "DeleteNetworkInterfaceOutput" awsDryRun: ""`
	logger *logger.Logger
	api    ec2iface.EC2API
	Id     *string `awsName: "NetworkInterfaceId" awsType: "awsstr" templateName: "id" required: ""`
}
type AttachNetworkinterface struct {
	_           string `action: "attach" entity: "networkinterface" awsAPI: "ec2" awsCall: "AttachNetworkInterface" awsInput: "AttachNetworkInterfaceInput" awsOutput: "AttachNetworkInterfaceOutput" awsDryRun: ""`
	logger      *logger.Logger
	api         ec2iface.EC2API
	Id          *string `awsName: "NetworkInterfaceId" awsType: "awsstr" templateName: "id" required: ""`
	Instance    *string `awsName: "InstanceId" awsType: "awsstr" templateName: "instance" required: ""`
	DeviceIndex *int64  `awsName: "DeviceIndex" awsType: "awsint64" templateName: "device-index" required: ""`
}
type DetachNetworkinterface struct {
	_          string `action: "detach" entity: "networkinterface" awsAPI: "ec2"`
	logger     *logger.Logger
	api        ec2iface.EC2API
	Attachment *string `awsName: "AttachmentId" awsType: "awsstr" templateName: "attachment"`
	Instance   *string `awsName: "InstanceId" awsType: "awsstr" templateName: "instance"`
	Id         *string `awsName: "NetworkInterfaceId" awsType: "awsstr" templateName: "id"`
	Force      *bool   `awsName: "Force" awsType: "awsbool" templateName: "force"`
}
type CheckNetworkinterface struct {
	_       string `action: "check" entity: "networkinterface" awsAPI: "ec2"`
	logger  *logger.Logger
	api     ec2iface.EC2API
	Id      *struct{} `templateName: "id" required: ""`
	State   *struct{} `templateName: "state" required: ""`
	Timeout *struct{} `templateName: "timeout" required: ""`
}
