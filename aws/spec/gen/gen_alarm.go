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
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
	"github.com/wallix/awless/logger"
)

type CreateAlarm struct {
	_                       string `action: "create" entity: "alarm" awsAPI: "cloudwatch" awsCall: "PutMetricAlarm" awsInput: "PutMetricAlarmInput" awsOutput: "PutMetricAlarmOutput"`
	logger                  *logger.Logger
	api                     cloudwatchiface.CloudWatchAPI
	Name                    *string   `awsName: "AlarmName" awsType: "awsstr" templateName: "name" required: ""`
	Operator                *string   `awsName: "ComparisonOperator" awsType: "awsstr" templateName: "operator" required: ""`
	Metric                  *string   `awsName: "MetricName" awsType: "awsstr" templateName: "metric" required: ""`
	Namespace               *string   `awsName: "Namespace" awsType: "awsstr" templateName: "namespace" required: ""`
	EvaluationPeriods       *int64    `awsName: "EvaluationPeriods" awsType: "awsint64" templateName: "evaluation-periods" required: ""`
	Period                  *int64    `awsName: "Period" awsType: "awsint64" templateName: "period" required: ""`
	StatisticFunction       *string   `awsName: "Statistic" awsType: "awsstr" templateName: "statistic-function" required: ""`
	Threshold               *float64  `awsName: "Threshold" awsType: "awsfloat" templateName: "threshold" required: ""`
	Enabled                 *bool     `awsName: "ActionsEnabled" awsType: "awsbool" templateName: "enabled"`
	AlarmActions            *[]string `awsName: "AlarmActions" awsType: "awsstringslice" templateName: "alarm-actions"`
	InsufficientdataActions *[]string `awsName: "InsufficientDataActions" awsType: "awsstringslice" templateName: "insufficientdata-actions"`
	OkActions               *[]string `awsName: "OKActions" awsType: "awsstringslice" templateName: "ok-actions"`
	Description             *string   `awsName: "AlarmDescription" awsType: "awsstr" templateName: "description"`
	Dimensions              *struct{} `awsName: "Dimensions" awsType: "awsdimensionslice" templateName: "dimensions"`
	Unit                    *string   `awsName: "Unit" awsType: "awsstr" templateName: "unit"`
}
type DeleteAlarm struct {
	_      string `action: "delete" entity: "alarm" awsAPI: "cloudwatch" awsCall: "DeleteAlarms" awsInput: "DeleteAlarmsInput" awsOutput: "DeleteAlarmsOutput"`
	logger *logger.Logger
	api    cloudwatchiface.CloudWatchAPI
	Name   *[]string `awsName: "AlarmNames" awsType: "awsstringslice" templateName: "name" required: ""`
}
type StartAlarm struct {
	_      string `action: "start" entity: "alarm" awsAPI: "cloudwatch" awsCall: "EnableAlarmActions" awsInput: "EnableAlarmActionsInput" awsOutput: "EnableAlarmActionsOutput"`
	logger *logger.Logger
	api    cloudwatchiface.CloudWatchAPI
	Names  *[]string `awsName: "AlarmNames" awsType: "awsstringslice" templateName: "names" required: ""`
}
type StopAlarm struct {
	_      string `action: "stop" entity: "alarm" awsAPI: "cloudwatch" awsCall: "DisableAlarmActions" awsInput: "DisableAlarmActionsInput" awsOutput: "DisableAlarmActionsOutput"`
	logger *logger.Logger
	api    cloudwatchiface.CloudWatchAPI
	Names  *[]string `awsName: "AlarmNames" awsType: "awsstringslice" templateName: "names" required: ""`
}
type AttachAlarm struct {
	_         string `action: "attach" entity: "alarm" awsAPI: "cloudwatch"`
	logger    *logger.Logger
	api       cloudwatchiface.CloudWatchAPI
	Name      *struct{} `templateName: "name" required: ""`
	ActionArn *struct{} `templateName: "action-arn" required: ""`
}
type DetachAlarm struct {
	_         string `action: "detach" entity: "alarm" awsAPI: "cloudwatch"`
	logger    *logger.Logger
	api       cloudwatchiface.CloudWatchAPI
	Name      *struct{} `templateName: "name" required: ""`
	ActionArn *struct{} `templateName: "action-arn" required: ""`
}
