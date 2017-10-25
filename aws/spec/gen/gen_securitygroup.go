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

type CreateSecuritygroup struct {
	_           string `action: "create" entity: "securitygroup" awsAPI: "ec2" awsCall: "CreateSecurityGroup" awsInput: "CreateSecurityGroupInput" awsOutput: "CreateSecurityGroupOutput" awsDryRun: ""`
	logger      *logger.Logger
	api         ec2iface.EC2API
	Name        *string `awsName: "GroupName" awsType: "awsstr" templateName: "name" required: ""`
	Vpc         *string `awsName: "VpcId" awsType: "awsstr" templateName: "vpc" required: ""`
	Description *string `awsName: "Description" awsType: "awsstr" templateName: "description" required: ""`
}
type UpdateSecuritygroup struct {
	_             string `action: "update" entity: "securitygroup" awsAPI: "ec2"`
	logger        *logger.Logger
	api           ec2iface.EC2API
	Id            *struct{} `templateName: "id" required: ""`
	Protocol      *struct{} `templateName: "protocol" required: ""`
	Cidr          *struct{} `templateName: "cidr"`
	Securitygroup *struct{} `templateName: "securitygroup"`
	Inbound       *struct{} `templateName: "inbound"`
	Outbound      *struct{} `templateName: "outbound"`
	Portrange     *struct{} `templateName: "portrange"`
}
type DeleteSecuritygroup struct {
	_      string `action: "delete" entity: "securitygroup" awsAPI: "ec2" awsCall: "DeleteSecurityGroup" awsInput: "DeleteSecurityGroupInput" awsOutput: "DeleteSecurityGroupOutput" awsDryRun: ""`
	logger *logger.Logger
	api    ec2iface.EC2API
	Id     *string `awsName: "GroupId" awsType: "awsstr" templateName: "id" required: ""`
}
type CheckSecuritygroup struct {
	_       string `action: "check" entity: "securitygroup" awsAPI: "ec2"`
	logger  *logger.Logger
	api     ec2iface.EC2API
	Id      *struct{} `templateName: "id" required: ""`
	State   *struct{} `templateName: "state" required: ""`
	Timeout *struct{} `templateName: "timeout" required: ""`
}
type AttachSecuritygroup struct {
	_        string `action: "attach" entity: "securitygroup" awsAPI: "ec2"`
	logger   *logger.Logger
	api      ec2iface.EC2API
	Id       *struct{} `templateName: "id" required: ""`
	Instance *struct{} `templateName: "instance"`
}
type DetachSecuritygroup struct {
	_        string `action: "detach" entity: "securitygroup" awsAPI: "ec2"`
	logger   *logger.Logger
	api      ec2iface.EC2API
	Id       *struct{} `templateName: "id" required: ""`
	Instance *struct{} `templateName: "instance"`
}
