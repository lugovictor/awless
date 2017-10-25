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

type CreateScalingpolicy struct {
	_                   string `action: "create" entity: "scalingpolicy" awsAPI: "autoscaling" awsCall: "PutScalingPolicy" awsInput: "PutScalingPolicyInput" awsOutput: "PutScalingPolicyOutput"`
	logger              *logger.Logger
	api                 autoscalingiface.AutoScalingAPI
	AdjustmentType      *string `awsName: "AdjustmentType" awsType: "awsstr" templateName: "adjustment-type" required: ""`
	Scalinggroup        *string `awsName: "AutoScalingGroupName" awsType: "awsstr" templateName: "scalinggroup" required: ""`
	Name                *string `awsName: "PolicyName" awsType: "awsstr" templateName: "name" required: ""`
	AdjustmentScaling   *int64  `awsName: "ScalingAdjustment" awsType: "awsint64" templateName: "adjustment-scaling" required: ""`
	Cooldown            *int64  `awsName: "Cooldown" awsType: "awsint64" templateName: "cooldown"`
	AdjustmentMagnitude *int64  `awsName: "MinAdjustmentMagnitude" awsType: "awsint64" templateName: "adjustment-magnitude"`
}
type DeleteScalingpolicy struct {
	_      string `action: "delete" entity: "scalingpolicy" awsAPI: "autoscaling" awsCall: "DeletePolicy" awsInput: "DeletePolicyInput" awsOutput: "DeletePolicyOutput"`
	logger *logger.Logger
	api    autoscalingiface.AutoScalingAPI
	Id     *string `awsName: "PolicyName" awsType: "awsstr" templateName: "id" required: ""`
}
