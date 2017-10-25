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

type CreateNatgateway struct {
	_           string `action: "create" entity: "natgateway" awsAPI: "ec2" awsCall: "CreateNatGateway" awsInput: "CreateNatGatewayInput" awsOutput: "CreateNatGatewayOutput"`
	logger      *logger.Logger
	api         ec2iface.EC2API
	ElasticipId *string `awsName: "AllocationId" awsType: "awsstr" templateName: "elasticip-id" required: ""`
	Subnet      *string `awsName: "SubnetId" awsType: "awsstr" templateName: "subnet" required: ""`
}
type DeleteNatgateway struct {
	_      string `action: "delete" entity: "natgateway" awsAPI: "ec2" awsCall: "DeleteNatGateway" awsInput: "DeleteNatGatewayInput" awsOutput: "DeleteNatGatewayOutput"`
	logger *logger.Logger
	api    ec2iface.EC2API
	Id     *string `awsName: "NatGatewayId" awsType: "awsstr" templateName: "id" required: ""`
}
type CheckNatgateway struct {
	_       string `action: "check" entity: "natgateway" awsAPI: "ec2"`
	logger  *logger.Logger
	api     ec2iface.EC2API
	Id      *struct{} `templateName: "id" required: ""`
	State   *struct{} `templateName: "state" required: ""`
	Timeout *struct{} `templateName: "timeout" required: ""`
}
