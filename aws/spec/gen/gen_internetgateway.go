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

type CreateInternetgateway struct {
	_      string `action: "create" entity: "internetgateway" awsAPI: "ec2" awsCall: "CreateInternetGateway" awsInput: "CreateInternetGatewayInput" awsOutput: "CreateInternetGatewayOutput" awsDryRun: ""`
	logger *logger.Logger
	api    ec2iface.EC2API
}
type DeleteInternetgateway struct {
	_      string `action: "delete" entity: "internetgateway" awsAPI: "ec2" awsCall: "DeleteInternetGateway" awsInput: "DeleteInternetGatewayInput" awsOutput: "DeleteInternetGatewayOutput" awsDryRun: ""`
	logger *logger.Logger
	api    ec2iface.EC2API
	Id     *string `awsName: "InternetGatewayId" awsType: "awsstr" templateName: "id" required: ""`
}
type AttachInternetgateway struct {
	_      string `action: "attach" entity: "internetgateway" awsAPI: "ec2" awsCall: "AttachInternetGateway" awsInput: "AttachInternetGatewayInput" awsOutput: "AttachInternetGatewayOutput" awsDryRun: ""`
	logger *logger.Logger
	api    ec2iface.EC2API
	Id     *string `awsName: "InternetGatewayId" awsType: "awsstr" templateName: "id" required: ""`
	Vpc    *string `awsName: "VpcId" awsType: "awsstr" templateName: "vpc" required: ""`
}
type DetachInternetgateway struct {
	_      string `action: "detach" entity: "internetgateway" awsAPI: "ec2" awsCall: "DetachInternetGateway" awsInput: "DetachInternetGatewayInput" awsOutput: "DetachInternetGatewayOutput" awsDryRun: ""`
	logger *logger.Logger
	api    ec2iface.EC2API
	Id     *string `awsName: "InternetGatewayId" awsType: "awsstr" templateName: "id" required: ""`
	Vpc    *string `awsName: "VpcId" awsType: "awsstr" templateName: "vpc" required: ""`
}
