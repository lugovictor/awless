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
	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/wallix/awless/logger"
)

type CreateNatgateway struct {
	_           string `action:"create" entity:"natgateway" awsAPI:"ec2" awsCall:"CreateNatGateway" awsInput:"ec2.CreateNatGatewayInput" awsOutput:"ec2.CreateNatGatewayOutput"`
	logger      *logger.Logger
	api         ec2iface.EC2API
	ElasticipId *string `awsName:"AllocationId" awsType:"awsstr" templateName:"elasticip-id" required:""`
	Subnet      *string `awsName:"SubnetId" awsType:"awsstr" templateName:"subnet" required:""`
}

func (cmd *CreateNatgateway) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}

func (cmd *CreateNatgateway) ExtractResult(i interface{}) string {
	return awssdk.StringValue(i.(*ec2.CreateNatGatewayOutput).NatGateway.NatGatewayId)
}

type DeleteNatgateway struct {
	_      string `action:"delete" entity:"natgateway" awsAPI:"ec2" awsCall:"DeleteNatGateway" awsInput:"ec2.DeleteNatGatewayInput" awsOutput:"ec2.DeleteNatGatewayOutput"`
	logger *logger.Logger
	api    ec2iface.EC2API
	Id     *string `awsName:"NatGatewayId" awsType:"awsstr" templateName:"id" required:""`
}

func (cmd *DeleteNatgateway) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}

type CheckNatgateway struct {
	_       string `action:"check" entity:"natgateway" awsAPI:"ec2"`
	logger  *logger.Logger
	api     ec2iface.EC2API
	Id      *struct{} `templateName:"id" required:""`
	State   *struct{} `templateName:"state" required:""`
	Timeout *struct{} `templateName:"timeout" required:""`
}
