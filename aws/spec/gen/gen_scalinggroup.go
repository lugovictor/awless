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
	"github.com/aws/aws-sdk-go/service/autoscaling/autoscalingiface"
	"github.com/wallix/awless/logger"
)

type CreateScalinggroup struct {
	_                      string `action:"create" entity:"scalinggroup" awsAPI:"autoscaling" awsCall:"CreateAutoScalingGroup" awsInput:"autoscaling.CreateAutoScalingGroupInput" awsOutput:"autoscaling.CreateAutoScalingGroupOutput"`
	logger                 *logger.Logger
	api                    autoscalingiface.AutoScalingAPI
	Name                   *string   `awsName:"AutoScalingGroupName" awsType:"awsstr" templateName:"name" required:""`
	Launchconfiguration    *string   `awsName:"LaunchConfigurationName" awsType:"awsstr" templateName:"launchconfiguration" required:""`
	MaxSize                *int64    `awsName:"MaxSize" awsType:"awsint64" templateName:"max-size" required:""`
	MinSize                *int64    `awsName:"MinSize" awsType:"awsint64" templateName:"min-size" required:""`
	Subnets                *[]string `awsName:"VPCZoneIdentifier" awsType:"awscsvstr" templateName:"subnets" required:""`
	Cooldown               *int64    `awsName:"DefaultCooldown" awsType:"awsint64" templateName:"cooldown"`
	DesiredCapacity        *int64    `awsName:"DesiredCapacity" awsType:"awsint64" templateName:"desired-capacity"`
	HealthcheckGracePeriod *int64    `awsName:"HealthCheckGracePeriod" awsType:"awsint64" templateName:"healthcheck-grace-period"`
	HealthcheckType        *string   `awsName:"HealthCheckType" awsType:"awsstr" templateName:"healthcheck-type"`
	NewInstancesProtected  *bool     `awsName:"NewInstancesProtectedFromScaleIn" awsType:"awsbool" templateName:"new-instances-protected"`
	Targetgroups           *[]string `awsName:"TargetGroupARNs" awsType:"awsstringslice" templateName:"targetgroups"`
}

func (cmd *CreateScalinggroup) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}

func (cmd *CreateScalinggroup) ExtractResult(i interface{}) string {
	return params["name"]
}

type UpdateScalinggroup struct {
	_                      string `action:"update" entity:"scalinggroup" awsAPI:"autoscaling" awsCall:"UpdateAutoScalingGroup" awsInput:"autoscaling.UpdateAutoScalingGroupInput" awsOutput:"autoscaling.UpdateAutoScalingGroupOutput"`
	logger                 *logger.Logger
	api                    autoscalingiface.AutoScalingAPI
	Name                   *string   `awsName:"AutoScalingGroupName" awsType:"awsstr" templateName:"name" required:""`
	Cooldown               *int64    `awsName:"DefaultCooldown" awsType:"awsint64" templateName:"cooldown"`
	DesiredCapacity        *int64    `awsName:"DesiredCapacity" awsType:"awsint64" templateName:"desired-capacity"`
	HealthcheckGracePeriod *int64    `awsName:"HealthCheckGracePeriod" awsType:"awsint64" templateName:"healthcheck-grace-period"`
	HealthcheckType        *string   `awsName:"HealthCheckType" awsType:"awsstr" templateName:"healthcheck-type"`
	Launchconfiguration    *string   `awsName:"LaunchConfigurationName" awsType:"awsstr" templateName:"launchconfiguration"`
	MaxSize                *int64    `awsName:"MaxSize" awsType:"awsint64" templateName:"max-size"`
	MinSize                *int64    `awsName:"MinSize" awsType:"awsint64" templateName:"min-size"`
	NewInstancesProtected  *bool     `awsName:"NewInstancesProtectedFromScaleIn" awsType:"awsbool" templateName:"new-instances-protected"`
	Subnets                *[]string `awsName:"VPCZoneIdentifier" awsType:"awscsvstr" templateName:"subnets"`
}

func (cmd *UpdateScalinggroup) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}

type DeleteScalinggroup struct {
	_      string `action:"delete" entity:"scalinggroup" awsAPI:"autoscaling" awsCall:"DeleteAutoScalingGroup" awsInput:"autoscaling.DeleteAutoScalingGroupInput" awsOutput:"autoscaling.DeleteAutoScalingGroupOutput"`
	logger *logger.Logger
	api    autoscalingiface.AutoScalingAPI
	Name   *string `awsName:"AutoScalingGroupName" awsType:"awsstr" templateName:"name" required:""`
	Force  *bool   `awsName:"ForceDelete" awsType:"awsbool" templateName:"force"`
}

func (cmd *DeleteScalinggroup) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}

type CheckScalinggroup struct {
	_       string `action:"check" entity:"scalinggroup" awsAPI:"autoscaling"`
	logger  *logger.Logger
	api     autoscalingiface.AutoScalingAPI
	Name    *struct{} `templateName:"name" required:""`
	Count   *struct{} `templateName:"count" required:""`
	Timeout *struct{} `templateName:"timeout" required:""`
}
