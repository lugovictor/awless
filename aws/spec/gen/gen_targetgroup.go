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
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/elbv2/elbv2iface"
	"github.com/wallix/awless/logger"
)

type CreateTargetgroup struct {
	_                   string `action:"create" entity:"targetgroup" awsAPI:"elbv2" awsCall:"CreateTargetGroup" awsInput:"elbv2.CreateTargetGroupInput" awsOutput:"elbv2.CreateTargetGroupOutput"`
	logger              *logger.Logger
	api                 elbv2iface.ELBV2API
	Name                *string `awsName:"Name" awsType:"awsstr" templateName:"name" required:""`
	Port                *int64  `awsName:"Port" awsType:"awsint64" templateName:"port" required:""`
	Protocol            *string `awsName:"Protocol" awsType:"awsstr" templateName:"protocol" required:""`
	Vpc                 *string `awsName:"VpcId" awsType:"awsstr" templateName:"vpc" required:""`
	Healthcheckinterval *int64  `awsName:"HealthCheckIntervalSeconds" awsType:"awsint64" templateName:"healthcheckinterval"`
	Healthcheckpath     *string `awsName:"HealthCheckPath" awsType:"awsstr" templateName:"healthcheckpath"`
	Healthcheckport     *string `awsName:"HealthCheckPort" awsType:"awsstr" templateName:"healthcheckport"`
	Healthcheckprotocol *string `awsName:"HealthCheckProtocol" awsType:"awsstr" templateName:"healthcheckprotocol"`
	Healthchecktimeout  *int64  `awsName:"HealthCheckTimeoutSeconds" awsType:"awsint64" templateName:"healthchecktimeout"`
	Healthythreshold    *int64  `awsName:"HealthyThresholdCount" awsType:"awsint64" templateName:"healthythreshold"`
	Unhealthythreshold  *int64  `awsName:"UnhealthyThresholdCount" awsType:"awsint64" templateName:"unhealthythreshold"`
	Matcher             *string `awsName:"Matcher.HttpCode" awsType:"awsstr" templateName:"matcher"`
}

func (cmd *CreateTargetgroup) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}

func (cmd *CreateTargetgroup) ExtractResult(i interface{}) string {
	return awssdk.StringValue(i.(*elbv2.CreateTargetGroupOutput).TargetGroups[0].TargetGroupArn)
}

type UpdateTargetgroup struct {
	_                   string `action:"update" entity:"targetgroup" awsAPI:"elbv2"`
	logger              *logger.Logger
	api                 elbv2iface.ELBV2API
	Id                  *string `awsName:"TargetGroupArn" awsType:"awsstr" templateName:"id" required:""`
	Deregistrationdelay *string `awsType:"awsstr" templateName:"deregistrationdelay"`
	Stickiness          *string `awsType:"awsstr" templateName:"stickiness"`
	Stickinessduration  *string `awsType:"awsstr" templateName:"stickinessduration"`
	Healthcheckinterval *int64  `awsName:"HealthCheckIntervalSeconds" awsType:"awsint64" templateName:"healthcheckinterval"`
	Healthcheckpath     *string `awsName:"HealthCheckPath" awsType:"awsstr" templateName:"healthcheckpath"`
	Healthcheckport     *string `awsName:"HealthCheckPort" awsType:"awsstr" templateName:"healthcheckport"`
	Healthcheckprotocol *string `awsName:"HealthCheckProtocol" awsType:"awsstr" templateName:"healthcheckprotocol"`
	Healthchecktimeout  *int64  `awsName:"HealthCheckTimeoutSeconds" awsType:"awsint64" templateName:"healthchecktimeout"`
	Healthythreshold    *int64  `awsName:"HealthyThresholdCount" awsType:"awsint64" templateName:"healthythreshold"`
	Unhealthythreshold  *int64  `awsName:"UnhealthyThresholdCount" awsType:"awsint64" templateName:"unhealthythreshold"`
	Matcher             *string `awsName:"Matcher.HttpCode" awsType:"awsstr" templateName:"matcher"`
}

type DeleteTargetgroup struct {
	_      string `action:"delete" entity:"targetgroup" awsAPI:"elbv2" awsCall:"DeleteTargetGroup" awsInput:"elbv2.DeleteTargetGroupInput" awsOutput:"elbv2.DeleteTargetGroupOutput"`
	logger *logger.Logger
	api    elbv2iface.ELBV2API
	Id     *string `awsName:"TargetGroupArn" awsType:"awsstr" templateName:"id" required:""`
}

func (cmd *DeleteTargetgroup) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}