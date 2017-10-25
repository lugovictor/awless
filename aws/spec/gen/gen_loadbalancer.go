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
	"github.com/aws/aws-sdk-go/service/elbv2/elbv2iface"
	"github.com/wallix/awless/logger"
)

type CreateLoadbalancer struct {
	_              string `action: "create" entity: "loadbalancer" awsAPI: "elbv2" awsCall: "CreateLoadBalancer" awsInput: "CreateLoadBalancerInput" awsOutput: "CreateLoadBalancerOutput"`
	logger         *logger.Logger
	api            elbv2iface.ELBV2API
	Name           *string   `awsName: "Name" awsType: "awsstr" templateName: "name" required: ""`
	Subnets        *[]string `awsName: "Subnets" awsType: "awsstringslice" templateName: "subnets" required: ""`
	SubnetMappings *struct{} `awsName: "SubnetMappings" awsType: "awssubnetmappings" templateName: "subnet-mappings"`
	Iptype         *string   `awsName: "IpAddressType" awsType: "awsstr" templateName: "iptype"`
	Scheme         *string   `awsName: "Scheme" awsType: "awsstr" templateName: "scheme"`
	Securitygroups *[]string `awsName: "SecurityGroups" awsType: "awsstringslice" templateName: "securitygroups"`
	Type           *string   `awsName: "Type" awsType: "awsstr" templateName: "type"`
}
type DeleteLoadbalancer struct {
	_      string `action: "delete" entity: "loadbalancer" awsAPI: "elbv2" awsCall: "DeleteLoadBalancer" awsInput: "DeleteLoadBalancerInput" awsOutput: "DeleteLoadBalancerOutput"`
	logger *logger.Logger
	api    elbv2iface.ELBV2API
	Id     *string `awsName: "LoadBalancerArn" awsType: "awsstr" templateName: "id" required: ""`
}
type CheckLoadbalancer struct {
	_       string `action: "check" entity: "loadbalancer" awsAPI: "elbv2"`
	logger  *logger.Logger
	api     elbv2iface.ELBV2API
	Id      *struct{} `templateName: "id" required: ""`
	State   *struct{} `templateName: "state" required: ""`
	Timeout *struct{} `templateName: "timeout" required: ""`
}
