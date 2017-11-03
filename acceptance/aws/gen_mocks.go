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
package awsat

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/aws/aws-sdk-go/service/route53/route53iface"
)

type cloudwatchMock struct {
	basicMock
	cloudwatchiface.CloudWatchAPI
	DeleteAlarmsFunc                       func(param0 *cloudwatch.DeleteAlarmsInput) (*cloudwatch.DeleteAlarmsOutput, error)
	DeleteAlarmsRequestFunc                func(param0 *cloudwatch.DeleteAlarmsInput) (*request.Request, *cloudwatch.DeleteAlarmsOutput)
	DeleteAlarmsWithContextFunc            func(param0 aws.Context, param1 *cloudwatch.DeleteAlarmsInput, param2 ...request.Option) (*cloudwatch.DeleteAlarmsOutput, error)
	DeleteDashboardsFunc                   func(param0 *cloudwatch.DeleteDashboardsInput) (*cloudwatch.DeleteDashboardsOutput, error)
	DeleteDashboardsRequestFunc            func(param0 *cloudwatch.DeleteDashboardsInput) (*request.Request, *cloudwatch.DeleteDashboardsOutput)
	DeleteDashboardsWithContextFunc        func(param0 aws.Context, param1 *cloudwatch.DeleteDashboardsInput, param2 ...request.Option) (*cloudwatch.DeleteDashboardsOutput, error)
	DescribeAlarmHistoryFunc               func(param0 *cloudwatch.DescribeAlarmHistoryInput) (*cloudwatch.DescribeAlarmHistoryOutput, error)
	DescribeAlarmHistoryRequestFunc        func(param0 *cloudwatch.DescribeAlarmHistoryInput) (*request.Request, *cloudwatch.DescribeAlarmHistoryOutput)
	DescribeAlarmHistoryWithContextFunc    func(param0 aws.Context, param1 *cloudwatch.DescribeAlarmHistoryInput, param2 ...request.Option) (*cloudwatch.DescribeAlarmHistoryOutput, error)
	DescribeAlarmsFunc                     func(param0 *cloudwatch.DescribeAlarmsInput) (*cloudwatch.DescribeAlarmsOutput, error)
	DescribeAlarmsForMetricFunc            func(param0 *cloudwatch.DescribeAlarmsForMetricInput) (*cloudwatch.DescribeAlarmsForMetricOutput, error)
	DescribeAlarmsForMetricRequestFunc     func(param0 *cloudwatch.DescribeAlarmsForMetricInput) (*request.Request, *cloudwatch.DescribeAlarmsForMetricOutput)
	DescribeAlarmsForMetricWithContextFunc func(param0 aws.Context, param1 *cloudwatch.DescribeAlarmsForMetricInput, param2 ...request.Option) (*cloudwatch.DescribeAlarmsForMetricOutput, error)
	DescribeAlarmsRequestFunc              func(param0 *cloudwatch.DescribeAlarmsInput) (*request.Request, *cloudwatch.DescribeAlarmsOutput)
	DescribeAlarmsWithContextFunc          func(param0 aws.Context, param1 *cloudwatch.DescribeAlarmsInput, param2 ...request.Option) (*cloudwatch.DescribeAlarmsOutput, error)
	DisableAlarmActionsFunc                func(param0 *cloudwatch.DisableAlarmActionsInput) (*cloudwatch.DisableAlarmActionsOutput, error)
	DisableAlarmActionsRequestFunc         func(param0 *cloudwatch.DisableAlarmActionsInput) (*request.Request, *cloudwatch.DisableAlarmActionsOutput)
	DisableAlarmActionsWithContextFunc     func(param0 aws.Context, param1 *cloudwatch.DisableAlarmActionsInput, param2 ...request.Option) (*cloudwatch.DisableAlarmActionsOutput, error)
	EnableAlarmActionsFunc                 func(param0 *cloudwatch.EnableAlarmActionsInput) (*cloudwatch.EnableAlarmActionsOutput, error)
	EnableAlarmActionsRequestFunc          func(param0 *cloudwatch.EnableAlarmActionsInput) (*request.Request, *cloudwatch.EnableAlarmActionsOutput)
	EnableAlarmActionsWithContextFunc      func(param0 aws.Context, param1 *cloudwatch.EnableAlarmActionsInput, param2 ...request.Option) (*cloudwatch.EnableAlarmActionsOutput, error)
	GetDashboardFunc                       func(param0 *cloudwatch.GetDashboardInput) (*cloudwatch.GetDashboardOutput, error)
	GetDashboardRequestFunc                func(param0 *cloudwatch.GetDashboardInput) (*request.Request, *cloudwatch.GetDashboardOutput)
	GetDashboardWithContextFunc            func(param0 aws.Context, param1 *cloudwatch.GetDashboardInput, param2 ...request.Option) (*cloudwatch.GetDashboardOutput, error)
	GetMetricStatisticsFunc                func(param0 *cloudwatch.GetMetricStatisticsInput) (*cloudwatch.GetMetricStatisticsOutput, error)
	GetMetricStatisticsRequestFunc         func(param0 *cloudwatch.GetMetricStatisticsInput) (*request.Request, *cloudwatch.GetMetricStatisticsOutput)
	GetMetricStatisticsWithContextFunc     func(param0 aws.Context, param1 *cloudwatch.GetMetricStatisticsInput, param2 ...request.Option) (*cloudwatch.GetMetricStatisticsOutput, error)
	ListDashboardsFunc                     func(param0 *cloudwatch.ListDashboardsInput) (*cloudwatch.ListDashboardsOutput, error)
	ListDashboardsRequestFunc              func(param0 *cloudwatch.ListDashboardsInput) (*request.Request, *cloudwatch.ListDashboardsOutput)
	ListDashboardsWithContextFunc          func(param0 aws.Context, param1 *cloudwatch.ListDashboardsInput, param2 ...request.Option) (*cloudwatch.ListDashboardsOutput, error)
	ListMetricsFunc                        func(param0 *cloudwatch.ListMetricsInput) (*cloudwatch.ListMetricsOutput, error)
	ListMetricsRequestFunc                 func(param0 *cloudwatch.ListMetricsInput) (*request.Request, *cloudwatch.ListMetricsOutput)
	ListMetricsWithContextFunc             func(param0 aws.Context, param1 *cloudwatch.ListMetricsInput, param2 ...request.Option) (*cloudwatch.ListMetricsOutput, error)
	PutDashboardFunc                       func(param0 *cloudwatch.PutDashboardInput) (*cloudwatch.PutDashboardOutput, error)
	PutDashboardRequestFunc                func(param0 *cloudwatch.PutDashboardInput) (*request.Request, *cloudwatch.PutDashboardOutput)
	PutDashboardWithContextFunc            func(param0 aws.Context, param1 *cloudwatch.PutDashboardInput, param2 ...request.Option) (*cloudwatch.PutDashboardOutput, error)
	PutMetricAlarmFunc                     func(param0 *cloudwatch.PutMetricAlarmInput) (*cloudwatch.PutMetricAlarmOutput, error)
	PutMetricAlarmRequestFunc              func(param0 *cloudwatch.PutMetricAlarmInput) (*request.Request, *cloudwatch.PutMetricAlarmOutput)
	PutMetricAlarmWithContextFunc          func(param0 aws.Context, param1 *cloudwatch.PutMetricAlarmInput, param2 ...request.Option) (*cloudwatch.PutMetricAlarmOutput, error)
	PutMetricDataFunc                      func(param0 *cloudwatch.PutMetricDataInput) (*cloudwatch.PutMetricDataOutput, error)
	PutMetricDataRequestFunc               func(param0 *cloudwatch.PutMetricDataInput) (*request.Request, *cloudwatch.PutMetricDataOutput)
	PutMetricDataWithContextFunc           func(param0 aws.Context, param1 *cloudwatch.PutMetricDataInput, param2 ...request.Option) (*cloudwatch.PutMetricDataOutput, error)
	SetAlarmStateFunc                      func(param0 *cloudwatch.SetAlarmStateInput) (*cloudwatch.SetAlarmStateOutput, error)
	SetAlarmStateRequestFunc               func(param0 *cloudwatch.SetAlarmStateInput) (*request.Request, *cloudwatch.SetAlarmStateOutput)
	SetAlarmStateWithContextFunc           func(param0 aws.Context, param1 *cloudwatch.SetAlarmStateInput, param2 ...request.Option) (*cloudwatch.SetAlarmStateOutput, error)
	WaitUntilAlarmExistsFunc               func(param0 *cloudwatch.DescribeAlarmsInput) error
	WaitUntilAlarmExistsWithContextFunc    func(param0 aws.Context, param1 *cloudwatch.DescribeAlarmsInput, param2 ...request.WaiterOption) error
}

func (m *cloudwatchMock) DeleteAlarms(param0 *cloudwatch.DeleteAlarmsInput) (*cloudwatch.DeleteAlarmsOutput, error) {
	m.addCall("DeleteAlarms")
	m.verifyInput("DeleteAlarms", param0)
	return m.DeleteAlarmsFunc(param0)
}

func (m *cloudwatchMock) DeleteAlarmsRequest(param0 *cloudwatch.DeleteAlarmsInput) (*request.Request, *cloudwatch.DeleteAlarmsOutput) {
	m.addCall("DeleteAlarmsRequest")
	m.verifyInput("DeleteAlarmsRequest", param0)
	return m.DeleteAlarmsRequestFunc(param0)
}

func (m *cloudwatchMock) DeleteAlarmsWithContext(param0 aws.Context, param1 *cloudwatch.DeleteAlarmsInput, param2 ...request.Option) (*cloudwatch.DeleteAlarmsOutput, error) {
	m.addCall("DeleteAlarmsWithContext")
	m.verifyInput("DeleteAlarmsWithContext", param0)
	return m.DeleteAlarmsWithContextFunc(param0, param1, param2...)
}

func (m *cloudwatchMock) DeleteDashboards(param0 *cloudwatch.DeleteDashboardsInput) (*cloudwatch.DeleteDashboardsOutput, error) {
	m.addCall("DeleteDashboards")
	m.verifyInput("DeleteDashboards", param0)
	return m.DeleteDashboardsFunc(param0)
}

func (m *cloudwatchMock) DeleteDashboardsRequest(param0 *cloudwatch.DeleteDashboardsInput) (*request.Request, *cloudwatch.DeleteDashboardsOutput) {
	m.addCall("DeleteDashboardsRequest")
	m.verifyInput("DeleteDashboardsRequest", param0)
	return m.DeleteDashboardsRequestFunc(param0)
}

func (m *cloudwatchMock) DeleteDashboardsWithContext(param0 aws.Context, param1 *cloudwatch.DeleteDashboardsInput, param2 ...request.Option) (*cloudwatch.DeleteDashboardsOutput, error) {
	m.addCall("DeleteDashboardsWithContext")
	m.verifyInput("DeleteDashboardsWithContext", param0)
	return m.DeleteDashboardsWithContextFunc(param0, param1, param2...)
}

func (m *cloudwatchMock) DescribeAlarmHistory(param0 *cloudwatch.DescribeAlarmHistoryInput) (*cloudwatch.DescribeAlarmHistoryOutput, error) {
	m.addCall("DescribeAlarmHistory")
	m.verifyInput("DescribeAlarmHistory", param0)
	return m.DescribeAlarmHistoryFunc(param0)
}

func (m *cloudwatchMock) DescribeAlarmHistoryRequest(param0 *cloudwatch.DescribeAlarmHistoryInput) (*request.Request, *cloudwatch.DescribeAlarmHistoryOutput) {
	m.addCall("DescribeAlarmHistoryRequest")
	m.verifyInput("DescribeAlarmHistoryRequest", param0)
	return m.DescribeAlarmHistoryRequestFunc(param0)
}

func (m *cloudwatchMock) DescribeAlarmHistoryWithContext(param0 aws.Context, param1 *cloudwatch.DescribeAlarmHistoryInput, param2 ...request.Option) (*cloudwatch.DescribeAlarmHistoryOutput, error) {
	m.addCall("DescribeAlarmHistoryWithContext")
	m.verifyInput("DescribeAlarmHistoryWithContext", param0)
	return m.DescribeAlarmHistoryWithContextFunc(param0, param1, param2...)
}

func (m *cloudwatchMock) DescribeAlarms(param0 *cloudwatch.DescribeAlarmsInput) (*cloudwatch.DescribeAlarmsOutput, error) {
	m.addCall("DescribeAlarms")
	m.verifyInput("DescribeAlarms", param0)
	return m.DescribeAlarmsFunc(param0)
}

func (m *cloudwatchMock) DescribeAlarmsForMetric(param0 *cloudwatch.DescribeAlarmsForMetricInput) (*cloudwatch.DescribeAlarmsForMetricOutput, error) {
	m.addCall("DescribeAlarmsForMetric")
	m.verifyInput("DescribeAlarmsForMetric", param0)
	return m.DescribeAlarmsForMetricFunc(param0)
}

func (m *cloudwatchMock) DescribeAlarmsForMetricRequest(param0 *cloudwatch.DescribeAlarmsForMetricInput) (*request.Request, *cloudwatch.DescribeAlarmsForMetricOutput) {
	m.addCall("DescribeAlarmsForMetricRequest")
	m.verifyInput("DescribeAlarmsForMetricRequest", param0)
	return m.DescribeAlarmsForMetricRequestFunc(param0)
}

func (m *cloudwatchMock) DescribeAlarmsForMetricWithContext(param0 aws.Context, param1 *cloudwatch.DescribeAlarmsForMetricInput, param2 ...request.Option) (*cloudwatch.DescribeAlarmsForMetricOutput, error) {
	m.addCall("DescribeAlarmsForMetricWithContext")
	m.verifyInput("DescribeAlarmsForMetricWithContext", param0)
	return m.DescribeAlarmsForMetricWithContextFunc(param0, param1, param2...)
}

func (m *cloudwatchMock) DescribeAlarmsRequest(param0 *cloudwatch.DescribeAlarmsInput) (*request.Request, *cloudwatch.DescribeAlarmsOutput) {
	m.addCall("DescribeAlarmsRequest")
	m.verifyInput("DescribeAlarmsRequest", param0)
	return m.DescribeAlarmsRequestFunc(param0)
}

func (m *cloudwatchMock) DescribeAlarmsWithContext(param0 aws.Context, param1 *cloudwatch.DescribeAlarmsInput, param2 ...request.Option) (*cloudwatch.DescribeAlarmsOutput, error) {
	m.addCall("DescribeAlarmsWithContext")
	m.verifyInput("DescribeAlarmsWithContext", param0)
	return m.DescribeAlarmsWithContextFunc(param0, param1, param2...)
}

func (m *cloudwatchMock) DisableAlarmActions(param0 *cloudwatch.DisableAlarmActionsInput) (*cloudwatch.DisableAlarmActionsOutput, error) {
	m.addCall("DisableAlarmActions")
	m.verifyInput("DisableAlarmActions", param0)
	return m.DisableAlarmActionsFunc(param0)
}

func (m *cloudwatchMock) DisableAlarmActionsRequest(param0 *cloudwatch.DisableAlarmActionsInput) (*request.Request, *cloudwatch.DisableAlarmActionsOutput) {
	m.addCall("DisableAlarmActionsRequest")
	m.verifyInput("DisableAlarmActionsRequest", param0)
	return m.DisableAlarmActionsRequestFunc(param0)
}

func (m *cloudwatchMock) DisableAlarmActionsWithContext(param0 aws.Context, param1 *cloudwatch.DisableAlarmActionsInput, param2 ...request.Option) (*cloudwatch.DisableAlarmActionsOutput, error) {
	m.addCall("DisableAlarmActionsWithContext")
	m.verifyInput("DisableAlarmActionsWithContext", param0)
	return m.DisableAlarmActionsWithContextFunc(param0, param1, param2...)
}

func (m *cloudwatchMock) EnableAlarmActions(param0 *cloudwatch.EnableAlarmActionsInput) (*cloudwatch.EnableAlarmActionsOutput, error) {
	m.addCall("EnableAlarmActions")
	m.verifyInput("EnableAlarmActions", param0)
	return m.EnableAlarmActionsFunc(param0)
}

func (m *cloudwatchMock) EnableAlarmActionsRequest(param0 *cloudwatch.EnableAlarmActionsInput) (*request.Request, *cloudwatch.EnableAlarmActionsOutput) {
	m.addCall("EnableAlarmActionsRequest")
	m.verifyInput("EnableAlarmActionsRequest", param0)
	return m.EnableAlarmActionsRequestFunc(param0)
}

func (m *cloudwatchMock) EnableAlarmActionsWithContext(param0 aws.Context, param1 *cloudwatch.EnableAlarmActionsInput, param2 ...request.Option) (*cloudwatch.EnableAlarmActionsOutput, error) {
	m.addCall("EnableAlarmActionsWithContext")
	m.verifyInput("EnableAlarmActionsWithContext", param0)
	return m.EnableAlarmActionsWithContextFunc(param0, param1, param2...)
}

func (m *cloudwatchMock) GetDashboard(param0 *cloudwatch.GetDashboardInput) (*cloudwatch.GetDashboardOutput, error) {
	m.addCall("GetDashboard")
	m.verifyInput("GetDashboard", param0)
	return m.GetDashboardFunc(param0)
}

func (m *cloudwatchMock) GetDashboardRequest(param0 *cloudwatch.GetDashboardInput) (*request.Request, *cloudwatch.GetDashboardOutput) {
	m.addCall("GetDashboardRequest")
	m.verifyInput("GetDashboardRequest", param0)
	return m.GetDashboardRequestFunc(param0)
}

func (m *cloudwatchMock) GetDashboardWithContext(param0 aws.Context, param1 *cloudwatch.GetDashboardInput, param2 ...request.Option) (*cloudwatch.GetDashboardOutput, error) {
	m.addCall("GetDashboardWithContext")
	m.verifyInput("GetDashboardWithContext", param0)
	return m.GetDashboardWithContextFunc(param0, param1, param2...)
}

func (m *cloudwatchMock) GetMetricStatistics(param0 *cloudwatch.GetMetricStatisticsInput) (*cloudwatch.GetMetricStatisticsOutput, error) {
	m.addCall("GetMetricStatistics")
	m.verifyInput("GetMetricStatistics", param0)
	return m.GetMetricStatisticsFunc(param0)
}

func (m *cloudwatchMock) GetMetricStatisticsRequest(param0 *cloudwatch.GetMetricStatisticsInput) (*request.Request, *cloudwatch.GetMetricStatisticsOutput) {
	m.addCall("GetMetricStatisticsRequest")
	m.verifyInput("GetMetricStatisticsRequest", param0)
	return m.GetMetricStatisticsRequestFunc(param0)
}

func (m *cloudwatchMock) GetMetricStatisticsWithContext(param0 aws.Context, param1 *cloudwatch.GetMetricStatisticsInput, param2 ...request.Option) (*cloudwatch.GetMetricStatisticsOutput, error) {
	m.addCall("GetMetricStatisticsWithContext")
	m.verifyInput("GetMetricStatisticsWithContext", param0)
	return m.GetMetricStatisticsWithContextFunc(param0, param1, param2...)
}

func (m *cloudwatchMock) ListDashboards(param0 *cloudwatch.ListDashboardsInput) (*cloudwatch.ListDashboardsOutput, error) {
	m.addCall("ListDashboards")
	m.verifyInput("ListDashboards", param0)
	return m.ListDashboardsFunc(param0)
}

func (m *cloudwatchMock) ListDashboardsRequest(param0 *cloudwatch.ListDashboardsInput) (*request.Request, *cloudwatch.ListDashboardsOutput) {
	m.addCall("ListDashboardsRequest")
	m.verifyInput("ListDashboardsRequest", param0)
	return m.ListDashboardsRequestFunc(param0)
}

func (m *cloudwatchMock) ListDashboardsWithContext(param0 aws.Context, param1 *cloudwatch.ListDashboardsInput, param2 ...request.Option) (*cloudwatch.ListDashboardsOutput, error) {
	m.addCall("ListDashboardsWithContext")
	m.verifyInput("ListDashboardsWithContext", param0)
	return m.ListDashboardsWithContextFunc(param0, param1, param2...)
}

func (m *cloudwatchMock) ListMetrics(param0 *cloudwatch.ListMetricsInput) (*cloudwatch.ListMetricsOutput, error) {
	m.addCall("ListMetrics")
	m.verifyInput("ListMetrics", param0)
	return m.ListMetricsFunc(param0)
}

func (m *cloudwatchMock) ListMetricsRequest(param0 *cloudwatch.ListMetricsInput) (*request.Request, *cloudwatch.ListMetricsOutput) {
	m.addCall("ListMetricsRequest")
	m.verifyInput("ListMetricsRequest", param0)
	return m.ListMetricsRequestFunc(param0)
}

func (m *cloudwatchMock) ListMetricsWithContext(param0 aws.Context, param1 *cloudwatch.ListMetricsInput, param2 ...request.Option) (*cloudwatch.ListMetricsOutput, error) {
	m.addCall("ListMetricsWithContext")
	m.verifyInput("ListMetricsWithContext", param0)
	return m.ListMetricsWithContextFunc(param0, param1, param2...)
}

func (m *cloudwatchMock) PutDashboard(param0 *cloudwatch.PutDashboardInput) (*cloudwatch.PutDashboardOutput, error) {
	m.addCall("PutDashboard")
	m.verifyInput("PutDashboard", param0)
	return m.PutDashboardFunc(param0)
}

func (m *cloudwatchMock) PutDashboardRequest(param0 *cloudwatch.PutDashboardInput) (*request.Request, *cloudwatch.PutDashboardOutput) {
	m.addCall("PutDashboardRequest")
	m.verifyInput("PutDashboardRequest", param0)
	return m.PutDashboardRequestFunc(param0)
}

func (m *cloudwatchMock) PutDashboardWithContext(param0 aws.Context, param1 *cloudwatch.PutDashboardInput, param2 ...request.Option) (*cloudwatch.PutDashboardOutput, error) {
	m.addCall("PutDashboardWithContext")
	m.verifyInput("PutDashboardWithContext", param0)
	return m.PutDashboardWithContextFunc(param0, param1, param2...)
}

func (m *cloudwatchMock) PutMetricAlarm(param0 *cloudwatch.PutMetricAlarmInput) (*cloudwatch.PutMetricAlarmOutput, error) {
	m.addCall("PutMetricAlarm")
	m.verifyInput("PutMetricAlarm", param0)
	return m.PutMetricAlarmFunc(param0)
}

func (m *cloudwatchMock) PutMetricAlarmRequest(param0 *cloudwatch.PutMetricAlarmInput) (*request.Request, *cloudwatch.PutMetricAlarmOutput) {
	m.addCall("PutMetricAlarmRequest")
	m.verifyInput("PutMetricAlarmRequest", param0)
	return m.PutMetricAlarmRequestFunc(param0)
}

func (m *cloudwatchMock) PutMetricAlarmWithContext(param0 aws.Context, param1 *cloudwatch.PutMetricAlarmInput, param2 ...request.Option) (*cloudwatch.PutMetricAlarmOutput, error) {
	m.addCall("PutMetricAlarmWithContext")
	m.verifyInput("PutMetricAlarmWithContext", param0)
	return m.PutMetricAlarmWithContextFunc(param0, param1, param2...)
}

func (m *cloudwatchMock) PutMetricData(param0 *cloudwatch.PutMetricDataInput) (*cloudwatch.PutMetricDataOutput, error) {
	m.addCall("PutMetricData")
	m.verifyInput("PutMetricData", param0)
	return m.PutMetricDataFunc(param0)
}

func (m *cloudwatchMock) PutMetricDataRequest(param0 *cloudwatch.PutMetricDataInput) (*request.Request, *cloudwatch.PutMetricDataOutput) {
	m.addCall("PutMetricDataRequest")
	m.verifyInput("PutMetricDataRequest", param0)
	return m.PutMetricDataRequestFunc(param0)
}

func (m *cloudwatchMock) PutMetricDataWithContext(param0 aws.Context, param1 *cloudwatch.PutMetricDataInput, param2 ...request.Option) (*cloudwatch.PutMetricDataOutput, error) {
	m.addCall("PutMetricDataWithContext")
	m.verifyInput("PutMetricDataWithContext", param0)
	return m.PutMetricDataWithContextFunc(param0, param1, param2...)
}

func (m *cloudwatchMock) SetAlarmState(param0 *cloudwatch.SetAlarmStateInput) (*cloudwatch.SetAlarmStateOutput, error) {
	m.addCall("SetAlarmState")
	m.verifyInput("SetAlarmState", param0)
	return m.SetAlarmStateFunc(param0)
}

func (m *cloudwatchMock) SetAlarmStateRequest(param0 *cloudwatch.SetAlarmStateInput) (*request.Request, *cloudwatch.SetAlarmStateOutput) {
	m.addCall("SetAlarmStateRequest")
	m.verifyInput("SetAlarmStateRequest", param0)
	return m.SetAlarmStateRequestFunc(param0)
}

func (m *cloudwatchMock) SetAlarmStateWithContext(param0 aws.Context, param1 *cloudwatch.SetAlarmStateInput, param2 ...request.Option) (*cloudwatch.SetAlarmStateOutput, error) {
	m.addCall("SetAlarmStateWithContext")
	m.verifyInput("SetAlarmStateWithContext", param0)
	return m.SetAlarmStateWithContextFunc(param0, param1, param2...)
}

func (m *cloudwatchMock) WaitUntilAlarmExists(param0 *cloudwatch.DescribeAlarmsInput) error {
	m.addCall("WaitUntilAlarmExists")
	m.verifyInput("WaitUntilAlarmExists", param0)
	return m.WaitUntilAlarmExistsFunc(param0)
}

func (m *cloudwatchMock) WaitUntilAlarmExistsWithContext(param0 aws.Context, param1 *cloudwatch.DescribeAlarmsInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilAlarmExistsWithContext")
	m.verifyInput("WaitUntilAlarmExistsWithContext", param0)
	return m.WaitUntilAlarmExistsWithContextFunc(param0, param1, param2...)
}

type ec2Mock struct {
	basicMock
	ec2iface.EC2API
	AcceptReservedInstancesExchangeQuoteFunc                  func(param0 *ec2.AcceptReservedInstancesExchangeQuoteInput) (*ec2.AcceptReservedInstancesExchangeQuoteOutput, error)
	AcceptReservedInstancesExchangeQuoteRequestFunc           func(param0 *ec2.AcceptReservedInstancesExchangeQuoteInput) (*request.Request, *ec2.AcceptReservedInstancesExchangeQuoteOutput)
	AcceptReservedInstancesExchangeQuoteWithContextFunc       func(param0 aws.Context, param1 *ec2.AcceptReservedInstancesExchangeQuoteInput, param2 ...request.Option) (*ec2.AcceptReservedInstancesExchangeQuoteOutput, error)
	AcceptVpcPeeringConnectionFunc                            func(param0 *ec2.AcceptVpcPeeringConnectionInput) (*ec2.AcceptVpcPeeringConnectionOutput, error)
	AcceptVpcPeeringConnectionRequestFunc                     func(param0 *ec2.AcceptVpcPeeringConnectionInput) (*request.Request, *ec2.AcceptVpcPeeringConnectionOutput)
	AcceptVpcPeeringConnectionWithContextFunc                 func(param0 aws.Context, param1 *ec2.AcceptVpcPeeringConnectionInput, param2 ...request.Option) (*ec2.AcceptVpcPeeringConnectionOutput, error)
	AllocateAddressFunc                                       func(param0 *ec2.AllocateAddressInput) (*ec2.AllocateAddressOutput, error)
	AllocateAddressRequestFunc                                func(param0 *ec2.AllocateAddressInput) (*request.Request, *ec2.AllocateAddressOutput)
	AllocateAddressWithContextFunc                            func(param0 aws.Context, param1 *ec2.AllocateAddressInput, param2 ...request.Option) (*ec2.AllocateAddressOutput, error)
	AllocateHostsFunc                                         func(param0 *ec2.AllocateHostsInput) (*ec2.AllocateHostsOutput, error)
	AllocateHostsRequestFunc                                  func(param0 *ec2.AllocateHostsInput) (*request.Request, *ec2.AllocateHostsOutput)
	AllocateHostsWithContextFunc                              func(param0 aws.Context, param1 *ec2.AllocateHostsInput, param2 ...request.Option) (*ec2.AllocateHostsOutput, error)
	AssignIpv6AddressesFunc                                   func(param0 *ec2.AssignIpv6AddressesInput) (*ec2.AssignIpv6AddressesOutput, error)
	AssignIpv6AddressesRequestFunc                            func(param0 *ec2.AssignIpv6AddressesInput) (*request.Request, *ec2.AssignIpv6AddressesOutput)
	AssignIpv6AddressesWithContextFunc                        func(param0 aws.Context, param1 *ec2.AssignIpv6AddressesInput, param2 ...request.Option) (*ec2.AssignIpv6AddressesOutput, error)
	AssignPrivateIpAddressesFunc                              func(param0 *ec2.AssignPrivateIpAddressesInput) (*ec2.AssignPrivateIpAddressesOutput, error)
	AssignPrivateIpAddressesRequestFunc                       func(param0 *ec2.AssignPrivateIpAddressesInput) (*request.Request, *ec2.AssignPrivateIpAddressesOutput)
	AssignPrivateIpAddressesWithContextFunc                   func(param0 aws.Context, param1 *ec2.AssignPrivateIpAddressesInput, param2 ...request.Option) (*ec2.AssignPrivateIpAddressesOutput, error)
	AssociateAddressFunc                                      func(param0 *ec2.AssociateAddressInput) (*ec2.AssociateAddressOutput, error)
	AssociateAddressRequestFunc                               func(param0 *ec2.AssociateAddressInput) (*request.Request, *ec2.AssociateAddressOutput)
	AssociateAddressWithContextFunc                           func(param0 aws.Context, param1 *ec2.AssociateAddressInput, param2 ...request.Option) (*ec2.AssociateAddressOutput, error)
	AssociateDhcpOptionsFunc                                  func(param0 *ec2.AssociateDhcpOptionsInput) (*ec2.AssociateDhcpOptionsOutput, error)
	AssociateDhcpOptionsRequestFunc                           func(param0 *ec2.AssociateDhcpOptionsInput) (*request.Request, *ec2.AssociateDhcpOptionsOutput)
	AssociateDhcpOptionsWithContextFunc                       func(param0 aws.Context, param1 *ec2.AssociateDhcpOptionsInput, param2 ...request.Option) (*ec2.AssociateDhcpOptionsOutput, error)
	AssociateIamInstanceProfileFunc                           func(param0 *ec2.AssociateIamInstanceProfileInput) (*ec2.AssociateIamInstanceProfileOutput, error)
	AssociateIamInstanceProfileRequestFunc                    func(param0 *ec2.AssociateIamInstanceProfileInput) (*request.Request, *ec2.AssociateIamInstanceProfileOutput)
	AssociateIamInstanceProfileWithContextFunc                func(param0 aws.Context, param1 *ec2.AssociateIamInstanceProfileInput, param2 ...request.Option) (*ec2.AssociateIamInstanceProfileOutput, error)
	AssociateRouteTableFunc                                   func(param0 *ec2.AssociateRouteTableInput) (*ec2.AssociateRouteTableOutput, error)
	AssociateRouteTableRequestFunc                            func(param0 *ec2.AssociateRouteTableInput) (*request.Request, *ec2.AssociateRouteTableOutput)
	AssociateRouteTableWithContextFunc                        func(param0 aws.Context, param1 *ec2.AssociateRouteTableInput, param2 ...request.Option) (*ec2.AssociateRouteTableOutput, error)
	AssociateSubnetCidrBlockFunc                              func(param0 *ec2.AssociateSubnetCidrBlockInput) (*ec2.AssociateSubnetCidrBlockOutput, error)
	AssociateSubnetCidrBlockRequestFunc                       func(param0 *ec2.AssociateSubnetCidrBlockInput) (*request.Request, *ec2.AssociateSubnetCidrBlockOutput)
	AssociateSubnetCidrBlockWithContextFunc                   func(param0 aws.Context, param1 *ec2.AssociateSubnetCidrBlockInput, param2 ...request.Option) (*ec2.AssociateSubnetCidrBlockOutput, error)
	AssociateVpcCidrBlockFunc                                 func(param0 *ec2.AssociateVpcCidrBlockInput) (*ec2.AssociateVpcCidrBlockOutput, error)
	AssociateVpcCidrBlockRequestFunc                          func(param0 *ec2.AssociateVpcCidrBlockInput) (*request.Request, *ec2.AssociateVpcCidrBlockOutput)
	AssociateVpcCidrBlockWithContextFunc                      func(param0 aws.Context, param1 *ec2.AssociateVpcCidrBlockInput, param2 ...request.Option) (*ec2.AssociateVpcCidrBlockOutput, error)
	AttachClassicLinkVpcFunc                                  func(param0 *ec2.AttachClassicLinkVpcInput) (*ec2.AttachClassicLinkVpcOutput, error)
	AttachClassicLinkVpcRequestFunc                           func(param0 *ec2.AttachClassicLinkVpcInput) (*request.Request, *ec2.AttachClassicLinkVpcOutput)
	AttachClassicLinkVpcWithContextFunc                       func(param0 aws.Context, param1 *ec2.AttachClassicLinkVpcInput, param2 ...request.Option) (*ec2.AttachClassicLinkVpcOutput, error)
	AttachInternetGatewayFunc                                 func(param0 *ec2.AttachInternetGatewayInput) (*ec2.AttachInternetGatewayOutput, error)
	AttachInternetGatewayRequestFunc                          func(param0 *ec2.AttachInternetGatewayInput) (*request.Request, *ec2.AttachInternetGatewayOutput)
	AttachInternetGatewayWithContextFunc                      func(param0 aws.Context, param1 *ec2.AttachInternetGatewayInput, param2 ...request.Option) (*ec2.AttachInternetGatewayOutput, error)
	AttachNetworkInterfaceFunc                                func(param0 *ec2.AttachNetworkInterfaceInput) (*ec2.AttachNetworkInterfaceOutput, error)
	AttachNetworkInterfaceRequestFunc                         func(param0 *ec2.AttachNetworkInterfaceInput) (*request.Request, *ec2.AttachNetworkInterfaceOutput)
	AttachNetworkInterfaceWithContextFunc                     func(param0 aws.Context, param1 *ec2.AttachNetworkInterfaceInput, param2 ...request.Option) (*ec2.AttachNetworkInterfaceOutput, error)
	AttachVolumeFunc                                          func(param0 *ec2.AttachVolumeInput) (*ec2.VolumeAttachment, error)
	AttachVolumeRequestFunc                                   func(param0 *ec2.AttachVolumeInput) (*request.Request, *ec2.VolumeAttachment)
	AttachVolumeWithContextFunc                               func(param0 aws.Context, param1 *ec2.AttachVolumeInput, param2 ...request.Option) (*ec2.VolumeAttachment, error)
	AttachVpnGatewayFunc                                      func(param0 *ec2.AttachVpnGatewayInput) (*ec2.AttachVpnGatewayOutput, error)
	AttachVpnGatewayRequestFunc                               func(param0 *ec2.AttachVpnGatewayInput) (*request.Request, *ec2.AttachVpnGatewayOutput)
	AttachVpnGatewayWithContextFunc                           func(param0 aws.Context, param1 *ec2.AttachVpnGatewayInput, param2 ...request.Option) (*ec2.AttachVpnGatewayOutput, error)
	AuthorizeSecurityGroupEgressFunc                          func(param0 *ec2.AuthorizeSecurityGroupEgressInput) (*ec2.AuthorizeSecurityGroupEgressOutput, error)
	AuthorizeSecurityGroupEgressRequestFunc                   func(param0 *ec2.AuthorizeSecurityGroupEgressInput) (*request.Request, *ec2.AuthorizeSecurityGroupEgressOutput)
	AuthorizeSecurityGroupEgressWithContextFunc               func(param0 aws.Context, param1 *ec2.AuthorizeSecurityGroupEgressInput, param2 ...request.Option) (*ec2.AuthorizeSecurityGroupEgressOutput, error)
	AuthorizeSecurityGroupIngressFunc                         func(param0 *ec2.AuthorizeSecurityGroupIngressInput) (*ec2.AuthorizeSecurityGroupIngressOutput, error)
	AuthorizeSecurityGroupIngressRequestFunc                  func(param0 *ec2.AuthorizeSecurityGroupIngressInput) (*request.Request, *ec2.AuthorizeSecurityGroupIngressOutput)
	AuthorizeSecurityGroupIngressWithContextFunc              func(param0 aws.Context, param1 *ec2.AuthorizeSecurityGroupIngressInput, param2 ...request.Option) (*ec2.AuthorizeSecurityGroupIngressOutput, error)
	BundleInstanceFunc                                        func(param0 *ec2.BundleInstanceInput) (*ec2.BundleInstanceOutput, error)
	BundleInstanceRequestFunc                                 func(param0 *ec2.BundleInstanceInput) (*request.Request, *ec2.BundleInstanceOutput)
	BundleInstanceWithContextFunc                             func(param0 aws.Context, param1 *ec2.BundleInstanceInput, param2 ...request.Option) (*ec2.BundleInstanceOutput, error)
	CancelBundleTaskFunc                                      func(param0 *ec2.CancelBundleTaskInput) (*ec2.CancelBundleTaskOutput, error)
	CancelBundleTaskRequestFunc                               func(param0 *ec2.CancelBundleTaskInput) (*request.Request, *ec2.CancelBundleTaskOutput)
	CancelBundleTaskWithContextFunc                           func(param0 aws.Context, param1 *ec2.CancelBundleTaskInput, param2 ...request.Option) (*ec2.CancelBundleTaskOutput, error)
	CancelConversionTaskFunc                                  func(param0 *ec2.CancelConversionTaskInput) (*ec2.CancelConversionTaskOutput, error)
	CancelConversionTaskRequestFunc                           func(param0 *ec2.CancelConversionTaskInput) (*request.Request, *ec2.CancelConversionTaskOutput)
	CancelConversionTaskWithContextFunc                       func(param0 aws.Context, param1 *ec2.CancelConversionTaskInput, param2 ...request.Option) (*ec2.CancelConversionTaskOutput, error)
	CancelExportTaskFunc                                      func(param0 *ec2.CancelExportTaskInput) (*ec2.CancelExportTaskOutput, error)
	CancelExportTaskRequestFunc                               func(param0 *ec2.CancelExportTaskInput) (*request.Request, *ec2.CancelExportTaskOutput)
	CancelExportTaskWithContextFunc                           func(param0 aws.Context, param1 *ec2.CancelExportTaskInput, param2 ...request.Option) (*ec2.CancelExportTaskOutput, error)
	CancelImportTaskFunc                                      func(param0 *ec2.CancelImportTaskInput) (*ec2.CancelImportTaskOutput, error)
	CancelImportTaskRequestFunc                               func(param0 *ec2.CancelImportTaskInput) (*request.Request, *ec2.CancelImportTaskOutput)
	CancelImportTaskWithContextFunc                           func(param0 aws.Context, param1 *ec2.CancelImportTaskInput, param2 ...request.Option) (*ec2.CancelImportTaskOutput, error)
	CancelReservedInstancesListingFunc                        func(param0 *ec2.CancelReservedInstancesListingInput) (*ec2.CancelReservedInstancesListingOutput, error)
	CancelReservedInstancesListingRequestFunc                 func(param0 *ec2.CancelReservedInstancesListingInput) (*request.Request, *ec2.CancelReservedInstancesListingOutput)
	CancelReservedInstancesListingWithContextFunc             func(param0 aws.Context, param1 *ec2.CancelReservedInstancesListingInput, param2 ...request.Option) (*ec2.CancelReservedInstancesListingOutput, error)
	CancelSpotFleetRequestsFunc                               func(param0 *ec2.CancelSpotFleetRequestsInput) (*ec2.CancelSpotFleetRequestsOutput, error)
	CancelSpotFleetRequestsRequestFunc                        func(param0 *ec2.CancelSpotFleetRequestsInput) (*request.Request, *ec2.CancelSpotFleetRequestsOutput)
	CancelSpotFleetRequestsWithContextFunc                    func(param0 aws.Context, param1 *ec2.CancelSpotFleetRequestsInput, param2 ...request.Option) (*ec2.CancelSpotFleetRequestsOutput, error)
	CancelSpotInstanceRequestsFunc                            func(param0 *ec2.CancelSpotInstanceRequestsInput) (*ec2.CancelSpotInstanceRequestsOutput, error)
	CancelSpotInstanceRequestsRequestFunc                     func(param0 *ec2.CancelSpotInstanceRequestsInput) (*request.Request, *ec2.CancelSpotInstanceRequestsOutput)
	CancelSpotInstanceRequestsWithContextFunc                 func(param0 aws.Context, param1 *ec2.CancelSpotInstanceRequestsInput, param2 ...request.Option) (*ec2.CancelSpotInstanceRequestsOutput, error)
	ConfirmProductInstanceFunc                                func(param0 *ec2.ConfirmProductInstanceInput) (*ec2.ConfirmProductInstanceOutput, error)
	ConfirmProductInstanceRequestFunc                         func(param0 *ec2.ConfirmProductInstanceInput) (*request.Request, *ec2.ConfirmProductInstanceOutput)
	ConfirmProductInstanceWithContextFunc                     func(param0 aws.Context, param1 *ec2.ConfirmProductInstanceInput, param2 ...request.Option) (*ec2.ConfirmProductInstanceOutput, error)
	CopyFpgaImageFunc                                         func(param0 *ec2.CopyFpgaImageInput) (*ec2.CopyFpgaImageOutput, error)
	CopyFpgaImageRequestFunc                                  func(param0 *ec2.CopyFpgaImageInput) (*request.Request, *ec2.CopyFpgaImageOutput)
	CopyFpgaImageWithContextFunc                              func(param0 aws.Context, param1 *ec2.CopyFpgaImageInput, param2 ...request.Option) (*ec2.CopyFpgaImageOutput, error)
	CopyImageFunc                                             func(param0 *ec2.CopyImageInput) (*ec2.CopyImageOutput, error)
	CopyImageRequestFunc                                      func(param0 *ec2.CopyImageInput) (*request.Request, *ec2.CopyImageOutput)
	CopyImageWithContextFunc                                  func(param0 aws.Context, param1 *ec2.CopyImageInput, param2 ...request.Option) (*ec2.CopyImageOutput, error)
	CopySnapshotFunc                                          func(param0 *ec2.CopySnapshotInput) (*ec2.CopySnapshotOutput, error)
	CopySnapshotRequestFunc                                   func(param0 *ec2.CopySnapshotInput) (*request.Request, *ec2.CopySnapshotOutput)
	CopySnapshotWithContextFunc                               func(param0 aws.Context, param1 *ec2.CopySnapshotInput, param2 ...request.Option) (*ec2.CopySnapshotOutput, error)
	CreateCustomerGatewayFunc                                 func(param0 *ec2.CreateCustomerGatewayInput) (*ec2.CreateCustomerGatewayOutput, error)
	CreateCustomerGatewayRequestFunc                          func(param0 *ec2.CreateCustomerGatewayInput) (*request.Request, *ec2.CreateCustomerGatewayOutput)
	CreateCustomerGatewayWithContextFunc                      func(param0 aws.Context, param1 *ec2.CreateCustomerGatewayInput, param2 ...request.Option) (*ec2.CreateCustomerGatewayOutput, error)
	CreateDefaultVpcFunc                                      func(param0 *ec2.CreateDefaultVpcInput) (*ec2.CreateDefaultVpcOutput, error)
	CreateDefaultVpcRequestFunc                               func(param0 *ec2.CreateDefaultVpcInput) (*request.Request, *ec2.CreateDefaultVpcOutput)
	CreateDefaultVpcWithContextFunc                           func(param0 aws.Context, param1 *ec2.CreateDefaultVpcInput, param2 ...request.Option) (*ec2.CreateDefaultVpcOutput, error)
	CreateDhcpOptionsFunc                                     func(param0 *ec2.CreateDhcpOptionsInput) (*ec2.CreateDhcpOptionsOutput, error)
	CreateDhcpOptionsRequestFunc                              func(param0 *ec2.CreateDhcpOptionsInput) (*request.Request, *ec2.CreateDhcpOptionsOutput)
	CreateDhcpOptionsWithContextFunc                          func(param0 aws.Context, param1 *ec2.CreateDhcpOptionsInput, param2 ...request.Option) (*ec2.CreateDhcpOptionsOutput, error)
	CreateEgressOnlyInternetGatewayFunc                       func(param0 *ec2.CreateEgressOnlyInternetGatewayInput) (*ec2.CreateEgressOnlyInternetGatewayOutput, error)
	CreateEgressOnlyInternetGatewayRequestFunc                func(param0 *ec2.CreateEgressOnlyInternetGatewayInput) (*request.Request, *ec2.CreateEgressOnlyInternetGatewayOutput)
	CreateEgressOnlyInternetGatewayWithContextFunc            func(param0 aws.Context, param1 *ec2.CreateEgressOnlyInternetGatewayInput, param2 ...request.Option) (*ec2.CreateEgressOnlyInternetGatewayOutput, error)
	CreateFlowLogsFunc                                        func(param0 *ec2.CreateFlowLogsInput) (*ec2.CreateFlowLogsOutput, error)
	CreateFlowLogsRequestFunc                                 func(param0 *ec2.CreateFlowLogsInput) (*request.Request, *ec2.CreateFlowLogsOutput)
	CreateFlowLogsWithContextFunc                             func(param0 aws.Context, param1 *ec2.CreateFlowLogsInput, param2 ...request.Option) (*ec2.CreateFlowLogsOutput, error)
	CreateFpgaImageFunc                                       func(param0 *ec2.CreateFpgaImageInput) (*ec2.CreateFpgaImageOutput, error)
	CreateFpgaImageRequestFunc                                func(param0 *ec2.CreateFpgaImageInput) (*request.Request, *ec2.CreateFpgaImageOutput)
	CreateFpgaImageWithContextFunc                            func(param0 aws.Context, param1 *ec2.CreateFpgaImageInput, param2 ...request.Option) (*ec2.CreateFpgaImageOutput, error)
	CreateImageFunc                                           func(param0 *ec2.CreateImageInput) (*ec2.CreateImageOutput, error)
	CreateImageRequestFunc                                    func(param0 *ec2.CreateImageInput) (*request.Request, *ec2.CreateImageOutput)
	CreateImageWithContextFunc                                func(param0 aws.Context, param1 *ec2.CreateImageInput, param2 ...request.Option) (*ec2.CreateImageOutput, error)
	CreateInstanceExportTaskFunc                              func(param0 *ec2.CreateInstanceExportTaskInput) (*ec2.CreateInstanceExportTaskOutput, error)
	CreateInstanceExportTaskRequestFunc                       func(param0 *ec2.CreateInstanceExportTaskInput) (*request.Request, *ec2.CreateInstanceExportTaskOutput)
	CreateInstanceExportTaskWithContextFunc                   func(param0 aws.Context, param1 *ec2.CreateInstanceExportTaskInput, param2 ...request.Option) (*ec2.CreateInstanceExportTaskOutput, error)
	CreateInternetGatewayFunc                                 func(param0 *ec2.CreateInternetGatewayInput) (*ec2.CreateInternetGatewayOutput, error)
	CreateInternetGatewayRequestFunc                          func(param0 *ec2.CreateInternetGatewayInput) (*request.Request, *ec2.CreateInternetGatewayOutput)
	CreateInternetGatewayWithContextFunc                      func(param0 aws.Context, param1 *ec2.CreateInternetGatewayInput, param2 ...request.Option) (*ec2.CreateInternetGatewayOutput, error)
	CreateKeyPairFunc                                         func(param0 *ec2.CreateKeyPairInput) (*ec2.CreateKeyPairOutput, error)
	CreateKeyPairRequestFunc                                  func(param0 *ec2.CreateKeyPairInput) (*request.Request, *ec2.CreateKeyPairOutput)
	CreateKeyPairWithContextFunc                              func(param0 aws.Context, param1 *ec2.CreateKeyPairInput, param2 ...request.Option) (*ec2.CreateKeyPairOutput, error)
	CreateNatGatewayFunc                                      func(param0 *ec2.CreateNatGatewayInput) (*ec2.CreateNatGatewayOutput, error)
	CreateNatGatewayRequestFunc                               func(param0 *ec2.CreateNatGatewayInput) (*request.Request, *ec2.CreateNatGatewayOutput)
	CreateNatGatewayWithContextFunc                           func(param0 aws.Context, param1 *ec2.CreateNatGatewayInput, param2 ...request.Option) (*ec2.CreateNatGatewayOutput, error)
	CreateNetworkAclFunc                                      func(param0 *ec2.CreateNetworkAclInput) (*ec2.CreateNetworkAclOutput, error)
	CreateNetworkAclEntryFunc                                 func(param0 *ec2.CreateNetworkAclEntryInput) (*ec2.CreateNetworkAclEntryOutput, error)
	CreateNetworkAclEntryRequestFunc                          func(param0 *ec2.CreateNetworkAclEntryInput) (*request.Request, *ec2.CreateNetworkAclEntryOutput)
	CreateNetworkAclEntryWithContextFunc                      func(param0 aws.Context, param1 *ec2.CreateNetworkAclEntryInput, param2 ...request.Option) (*ec2.CreateNetworkAclEntryOutput, error)
	CreateNetworkAclRequestFunc                               func(param0 *ec2.CreateNetworkAclInput) (*request.Request, *ec2.CreateNetworkAclOutput)
	CreateNetworkAclWithContextFunc                           func(param0 aws.Context, param1 *ec2.CreateNetworkAclInput, param2 ...request.Option) (*ec2.CreateNetworkAclOutput, error)
	CreateNetworkInterfaceFunc                                func(param0 *ec2.CreateNetworkInterfaceInput) (*ec2.CreateNetworkInterfaceOutput, error)
	CreateNetworkInterfacePermissionFunc                      func(param0 *ec2.CreateNetworkInterfacePermissionInput) (*ec2.CreateNetworkInterfacePermissionOutput, error)
	CreateNetworkInterfacePermissionRequestFunc               func(param0 *ec2.CreateNetworkInterfacePermissionInput) (*request.Request, *ec2.CreateNetworkInterfacePermissionOutput)
	CreateNetworkInterfacePermissionWithContextFunc           func(param0 aws.Context, param1 *ec2.CreateNetworkInterfacePermissionInput, param2 ...request.Option) (*ec2.CreateNetworkInterfacePermissionOutput, error)
	CreateNetworkInterfaceRequestFunc                         func(param0 *ec2.CreateNetworkInterfaceInput) (*request.Request, *ec2.CreateNetworkInterfaceOutput)
	CreateNetworkInterfaceWithContextFunc                     func(param0 aws.Context, param1 *ec2.CreateNetworkInterfaceInput, param2 ...request.Option) (*ec2.CreateNetworkInterfaceOutput, error)
	CreatePlacementGroupFunc                                  func(param0 *ec2.CreatePlacementGroupInput) (*ec2.CreatePlacementGroupOutput, error)
	CreatePlacementGroupRequestFunc                           func(param0 *ec2.CreatePlacementGroupInput) (*request.Request, *ec2.CreatePlacementGroupOutput)
	CreatePlacementGroupWithContextFunc                       func(param0 aws.Context, param1 *ec2.CreatePlacementGroupInput, param2 ...request.Option) (*ec2.CreatePlacementGroupOutput, error)
	CreateReservedInstancesListingFunc                        func(param0 *ec2.CreateReservedInstancesListingInput) (*ec2.CreateReservedInstancesListingOutput, error)
	CreateReservedInstancesListingRequestFunc                 func(param0 *ec2.CreateReservedInstancesListingInput) (*request.Request, *ec2.CreateReservedInstancesListingOutput)
	CreateReservedInstancesListingWithContextFunc             func(param0 aws.Context, param1 *ec2.CreateReservedInstancesListingInput, param2 ...request.Option) (*ec2.CreateReservedInstancesListingOutput, error)
	CreateRouteFunc                                           func(param0 *ec2.CreateRouteInput) (*ec2.CreateRouteOutput, error)
	CreateRouteRequestFunc                                    func(param0 *ec2.CreateRouteInput) (*request.Request, *ec2.CreateRouteOutput)
	CreateRouteTableFunc                                      func(param0 *ec2.CreateRouteTableInput) (*ec2.CreateRouteTableOutput, error)
	CreateRouteTableRequestFunc                               func(param0 *ec2.CreateRouteTableInput) (*request.Request, *ec2.CreateRouteTableOutput)
	CreateRouteTableWithContextFunc                           func(param0 aws.Context, param1 *ec2.CreateRouteTableInput, param2 ...request.Option) (*ec2.CreateRouteTableOutput, error)
	CreateRouteWithContextFunc                                func(param0 aws.Context, param1 *ec2.CreateRouteInput, param2 ...request.Option) (*ec2.CreateRouteOutput, error)
	CreateSecurityGroupFunc                                   func(param0 *ec2.CreateSecurityGroupInput) (*ec2.CreateSecurityGroupOutput, error)
	CreateSecurityGroupRequestFunc                            func(param0 *ec2.CreateSecurityGroupInput) (*request.Request, *ec2.CreateSecurityGroupOutput)
	CreateSecurityGroupWithContextFunc                        func(param0 aws.Context, param1 *ec2.CreateSecurityGroupInput, param2 ...request.Option) (*ec2.CreateSecurityGroupOutput, error)
	CreateSnapshotFunc                                        func(param0 *ec2.CreateSnapshotInput) (*ec2.Snapshot, error)
	CreateSnapshotRequestFunc                                 func(param0 *ec2.CreateSnapshotInput) (*request.Request, *ec2.Snapshot)
	CreateSnapshotWithContextFunc                             func(param0 aws.Context, param1 *ec2.CreateSnapshotInput, param2 ...request.Option) (*ec2.Snapshot, error)
	CreateSpotDatafeedSubscriptionFunc                        func(param0 *ec2.CreateSpotDatafeedSubscriptionInput) (*ec2.CreateSpotDatafeedSubscriptionOutput, error)
	CreateSpotDatafeedSubscriptionRequestFunc                 func(param0 *ec2.CreateSpotDatafeedSubscriptionInput) (*request.Request, *ec2.CreateSpotDatafeedSubscriptionOutput)
	CreateSpotDatafeedSubscriptionWithContextFunc             func(param0 aws.Context, param1 *ec2.CreateSpotDatafeedSubscriptionInput, param2 ...request.Option) (*ec2.CreateSpotDatafeedSubscriptionOutput, error)
	CreateSubnetFunc                                          func(param0 *ec2.CreateSubnetInput) (*ec2.CreateSubnetOutput, error)
	CreateSubnetRequestFunc                                   func(param0 *ec2.CreateSubnetInput) (*request.Request, *ec2.CreateSubnetOutput)
	CreateSubnetWithContextFunc                               func(param0 aws.Context, param1 *ec2.CreateSubnetInput, param2 ...request.Option) (*ec2.CreateSubnetOutput, error)
	CreateTagsFunc                                            func(param0 *ec2.CreateTagsInput) (*ec2.CreateTagsOutput, error)
	CreateTagsRequestFunc                                     func(param0 *ec2.CreateTagsInput) (*request.Request, *ec2.CreateTagsOutput)
	CreateTagsWithContextFunc                                 func(param0 aws.Context, param1 *ec2.CreateTagsInput, param2 ...request.Option) (*ec2.CreateTagsOutput, error)
	CreateVolumeFunc                                          func(param0 *ec2.CreateVolumeInput) (*ec2.Volume, error)
	CreateVolumeRequestFunc                                   func(param0 *ec2.CreateVolumeInput) (*request.Request, *ec2.Volume)
	CreateVolumeWithContextFunc                               func(param0 aws.Context, param1 *ec2.CreateVolumeInput, param2 ...request.Option) (*ec2.Volume, error)
	CreateVpcFunc                                             func(param0 *ec2.CreateVpcInput) (*ec2.CreateVpcOutput, error)
	CreateVpcEndpointFunc                                     func(param0 *ec2.CreateVpcEndpointInput) (*ec2.CreateVpcEndpointOutput, error)
	CreateVpcEndpointRequestFunc                              func(param0 *ec2.CreateVpcEndpointInput) (*request.Request, *ec2.CreateVpcEndpointOutput)
	CreateVpcEndpointWithContextFunc                          func(param0 aws.Context, param1 *ec2.CreateVpcEndpointInput, param2 ...request.Option) (*ec2.CreateVpcEndpointOutput, error)
	CreateVpcPeeringConnectionFunc                            func(param0 *ec2.CreateVpcPeeringConnectionInput) (*ec2.CreateVpcPeeringConnectionOutput, error)
	CreateVpcPeeringConnectionRequestFunc                     func(param0 *ec2.CreateVpcPeeringConnectionInput) (*request.Request, *ec2.CreateVpcPeeringConnectionOutput)
	CreateVpcPeeringConnectionWithContextFunc                 func(param0 aws.Context, param1 *ec2.CreateVpcPeeringConnectionInput, param2 ...request.Option) (*ec2.CreateVpcPeeringConnectionOutput, error)
	CreateVpcRequestFunc                                      func(param0 *ec2.CreateVpcInput) (*request.Request, *ec2.CreateVpcOutput)
	CreateVpcWithContextFunc                                  func(param0 aws.Context, param1 *ec2.CreateVpcInput, param2 ...request.Option) (*ec2.CreateVpcOutput, error)
	CreateVpnConnectionFunc                                   func(param0 *ec2.CreateVpnConnectionInput) (*ec2.CreateVpnConnectionOutput, error)
	CreateVpnConnectionRequestFunc                            func(param0 *ec2.CreateVpnConnectionInput) (*request.Request, *ec2.CreateVpnConnectionOutput)
	CreateVpnConnectionRouteFunc                              func(param0 *ec2.CreateVpnConnectionRouteInput) (*ec2.CreateVpnConnectionRouteOutput, error)
	CreateVpnConnectionRouteRequestFunc                       func(param0 *ec2.CreateVpnConnectionRouteInput) (*request.Request, *ec2.CreateVpnConnectionRouteOutput)
	CreateVpnConnectionRouteWithContextFunc                   func(param0 aws.Context, param1 *ec2.CreateVpnConnectionRouteInput, param2 ...request.Option) (*ec2.CreateVpnConnectionRouteOutput, error)
	CreateVpnConnectionWithContextFunc                        func(param0 aws.Context, param1 *ec2.CreateVpnConnectionInput, param2 ...request.Option) (*ec2.CreateVpnConnectionOutput, error)
	CreateVpnGatewayFunc                                      func(param0 *ec2.CreateVpnGatewayInput) (*ec2.CreateVpnGatewayOutput, error)
	CreateVpnGatewayRequestFunc                               func(param0 *ec2.CreateVpnGatewayInput) (*request.Request, *ec2.CreateVpnGatewayOutput)
	CreateVpnGatewayWithContextFunc                           func(param0 aws.Context, param1 *ec2.CreateVpnGatewayInput, param2 ...request.Option) (*ec2.CreateVpnGatewayOutput, error)
	DeleteCustomerGatewayFunc                                 func(param0 *ec2.DeleteCustomerGatewayInput) (*ec2.DeleteCustomerGatewayOutput, error)
	DeleteCustomerGatewayRequestFunc                          func(param0 *ec2.DeleteCustomerGatewayInput) (*request.Request, *ec2.DeleteCustomerGatewayOutput)
	DeleteCustomerGatewayWithContextFunc                      func(param0 aws.Context, param1 *ec2.DeleteCustomerGatewayInput, param2 ...request.Option) (*ec2.DeleteCustomerGatewayOutput, error)
	DeleteDhcpOptionsFunc                                     func(param0 *ec2.DeleteDhcpOptionsInput) (*ec2.DeleteDhcpOptionsOutput, error)
	DeleteDhcpOptionsRequestFunc                              func(param0 *ec2.DeleteDhcpOptionsInput) (*request.Request, *ec2.DeleteDhcpOptionsOutput)
	DeleteDhcpOptionsWithContextFunc                          func(param0 aws.Context, param1 *ec2.DeleteDhcpOptionsInput, param2 ...request.Option) (*ec2.DeleteDhcpOptionsOutput, error)
	DeleteEgressOnlyInternetGatewayFunc                       func(param0 *ec2.DeleteEgressOnlyInternetGatewayInput) (*ec2.DeleteEgressOnlyInternetGatewayOutput, error)
	DeleteEgressOnlyInternetGatewayRequestFunc                func(param0 *ec2.DeleteEgressOnlyInternetGatewayInput) (*request.Request, *ec2.DeleteEgressOnlyInternetGatewayOutput)
	DeleteEgressOnlyInternetGatewayWithContextFunc            func(param0 aws.Context, param1 *ec2.DeleteEgressOnlyInternetGatewayInput, param2 ...request.Option) (*ec2.DeleteEgressOnlyInternetGatewayOutput, error)
	DeleteFlowLogsFunc                                        func(param0 *ec2.DeleteFlowLogsInput) (*ec2.DeleteFlowLogsOutput, error)
	DeleteFlowLogsRequestFunc                                 func(param0 *ec2.DeleteFlowLogsInput) (*request.Request, *ec2.DeleteFlowLogsOutput)
	DeleteFlowLogsWithContextFunc                             func(param0 aws.Context, param1 *ec2.DeleteFlowLogsInput, param2 ...request.Option) (*ec2.DeleteFlowLogsOutput, error)
	DeleteFpgaImageFunc                                       func(param0 *ec2.DeleteFpgaImageInput) (*ec2.DeleteFpgaImageOutput, error)
	DeleteFpgaImageRequestFunc                                func(param0 *ec2.DeleteFpgaImageInput) (*request.Request, *ec2.DeleteFpgaImageOutput)
	DeleteFpgaImageWithContextFunc                            func(param0 aws.Context, param1 *ec2.DeleteFpgaImageInput, param2 ...request.Option) (*ec2.DeleteFpgaImageOutput, error)
	DeleteInternetGatewayFunc                                 func(param0 *ec2.DeleteInternetGatewayInput) (*ec2.DeleteInternetGatewayOutput, error)
	DeleteInternetGatewayRequestFunc                          func(param0 *ec2.DeleteInternetGatewayInput) (*request.Request, *ec2.DeleteInternetGatewayOutput)
	DeleteInternetGatewayWithContextFunc                      func(param0 aws.Context, param1 *ec2.DeleteInternetGatewayInput, param2 ...request.Option) (*ec2.DeleteInternetGatewayOutput, error)
	DeleteKeyPairFunc                                         func(param0 *ec2.DeleteKeyPairInput) (*ec2.DeleteKeyPairOutput, error)
	DeleteKeyPairRequestFunc                                  func(param0 *ec2.DeleteKeyPairInput) (*request.Request, *ec2.DeleteKeyPairOutput)
	DeleteKeyPairWithContextFunc                              func(param0 aws.Context, param1 *ec2.DeleteKeyPairInput, param2 ...request.Option) (*ec2.DeleteKeyPairOutput, error)
	DeleteNatGatewayFunc                                      func(param0 *ec2.DeleteNatGatewayInput) (*ec2.DeleteNatGatewayOutput, error)
	DeleteNatGatewayRequestFunc                               func(param0 *ec2.DeleteNatGatewayInput) (*request.Request, *ec2.DeleteNatGatewayOutput)
	DeleteNatGatewayWithContextFunc                           func(param0 aws.Context, param1 *ec2.DeleteNatGatewayInput, param2 ...request.Option) (*ec2.DeleteNatGatewayOutput, error)
	DeleteNetworkAclFunc                                      func(param0 *ec2.DeleteNetworkAclInput) (*ec2.DeleteNetworkAclOutput, error)
	DeleteNetworkAclEntryFunc                                 func(param0 *ec2.DeleteNetworkAclEntryInput) (*ec2.DeleteNetworkAclEntryOutput, error)
	DeleteNetworkAclEntryRequestFunc                          func(param0 *ec2.DeleteNetworkAclEntryInput) (*request.Request, *ec2.DeleteNetworkAclEntryOutput)
	DeleteNetworkAclEntryWithContextFunc                      func(param0 aws.Context, param1 *ec2.DeleteNetworkAclEntryInput, param2 ...request.Option) (*ec2.DeleteNetworkAclEntryOutput, error)
	DeleteNetworkAclRequestFunc                               func(param0 *ec2.DeleteNetworkAclInput) (*request.Request, *ec2.DeleteNetworkAclOutput)
	DeleteNetworkAclWithContextFunc                           func(param0 aws.Context, param1 *ec2.DeleteNetworkAclInput, param2 ...request.Option) (*ec2.DeleteNetworkAclOutput, error)
	DeleteNetworkInterfaceFunc                                func(param0 *ec2.DeleteNetworkInterfaceInput) (*ec2.DeleteNetworkInterfaceOutput, error)
	DeleteNetworkInterfacePermissionFunc                      func(param0 *ec2.DeleteNetworkInterfacePermissionInput) (*ec2.DeleteNetworkInterfacePermissionOutput, error)
	DeleteNetworkInterfacePermissionRequestFunc               func(param0 *ec2.DeleteNetworkInterfacePermissionInput) (*request.Request, *ec2.DeleteNetworkInterfacePermissionOutput)
	DeleteNetworkInterfacePermissionWithContextFunc           func(param0 aws.Context, param1 *ec2.DeleteNetworkInterfacePermissionInput, param2 ...request.Option) (*ec2.DeleteNetworkInterfacePermissionOutput, error)
	DeleteNetworkInterfaceRequestFunc                         func(param0 *ec2.DeleteNetworkInterfaceInput) (*request.Request, *ec2.DeleteNetworkInterfaceOutput)
	DeleteNetworkInterfaceWithContextFunc                     func(param0 aws.Context, param1 *ec2.DeleteNetworkInterfaceInput, param2 ...request.Option) (*ec2.DeleteNetworkInterfaceOutput, error)
	DeletePlacementGroupFunc                                  func(param0 *ec2.DeletePlacementGroupInput) (*ec2.DeletePlacementGroupOutput, error)
	DeletePlacementGroupRequestFunc                           func(param0 *ec2.DeletePlacementGroupInput) (*request.Request, *ec2.DeletePlacementGroupOutput)
	DeletePlacementGroupWithContextFunc                       func(param0 aws.Context, param1 *ec2.DeletePlacementGroupInput, param2 ...request.Option) (*ec2.DeletePlacementGroupOutput, error)
	DeleteRouteFunc                                           func(param0 *ec2.DeleteRouteInput) (*ec2.DeleteRouteOutput, error)
	DeleteRouteRequestFunc                                    func(param0 *ec2.DeleteRouteInput) (*request.Request, *ec2.DeleteRouteOutput)
	DeleteRouteTableFunc                                      func(param0 *ec2.DeleteRouteTableInput) (*ec2.DeleteRouteTableOutput, error)
	DeleteRouteTableRequestFunc                               func(param0 *ec2.DeleteRouteTableInput) (*request.Request, *ec2.DeleteRouteTableOutput)
	DeleteRouteTableWithContextFunc                           func(param0 aws.Context, param1 *ec2.DeleteRouteTableInput, param2 ...request.Option) (*ec2.DeleteRouteTableOutput, error)
	DeleteRouteWithContextFunc                                func(param0 aws.Context, param1 *ec2.DeleteRouteInput, param2 ...request.Option) (*ec2.DeleteRouteOutput, error)
	DeleteSecurityGroupFunc                                   func(param0 *ec2.DeleteSecurityGroupInput) (*ec2.DeleteSecurityGroupOutput, error)
	DeleteSecurityGroupRequestFunc                            func(param0 *ec2.DeleteSecurityGroupInput) (*request.Request, *ec2.DeleteSecurityGroupOutput)
	DeleteSecurityGroupWithContextFunc                        func(param0 aws.Context, param1 *ec2.DeleteSecurityGroupInput, param2 ...request.Option) (*ec2.DeleteSecurityGroupOutput, error)
	DeleteSnapshotFunc                                        func(param0 *ec2.DeleteSnapshotInput) (*ec2.DeleteSnapshotOutput, error)
	DeleteSnapshotRequestFunc                                 func(param0 *ec2.DeleteSnapshotInput) (*request.Request, *ec2.DeleteSnapshotOutput)
	DeleteSnapshotWithContextFunc                             func(param0 aws.Context, param1 *ec2.DeleteSnapshotInput, param2 ...request.Option) (*ec2.DeleteSnapshotOutput, error)
	DeleteSpotDatafeedSubscriptionFunc                        func(param0 *ec2.DeleteSpotDatafeedSubscriptionInput) (*ec2.DeleteSpotDatafeedSubscriptionOutput, error)
	DeleteSpotDatafeedSubscriptionRequestFunc                 func(param0 *ec2.DeleteSpotDatafeedSubscriptionInput) (*request.Request, *ec2.DeleteSpotDatafeedSubscriptionOutput)
	DeleteSpotDatafeedSubscriptionWithContextFunc             func(param0 aws.Context, param1 *ec2.DeleteSpotDatafeedSubscriptionInput, param2 ...request.Option) (*ec2.DeleteSpotDatafeedSubscriptionOutput, error)
	DeleteSubnetFunc                                          func(param0 *ec2.DeleteSubnetInput) (*ec2.DeleteSubnetOutput, error)
	DeleteSubnetRequestFunc                                   func(param0 *ec2.DeleteSubnetInput) (*request.Request, *ec2.DeleteSubnetOutput)
	DeleteSubnetWithContextFunc                               func(param0 aws.Context, param1 *ec2.DeleteSubnetInput, param2 ...request.Option) (*ec2.DeleteSubnetOutput, error)
	DeleteTagsFunc                                            func(param0 *ec2.DeleteTagsInput) (*ec2.DeleteTagsOutput, error)
	DeleteTagsRequestFunc                                     func(param0 *ec2.DeleteTagsInput) (*request.Request, *ec2.DeleteTagsOutput)
	DeleteTagsWithContextFunc                                 func(param0 aws.Context, param1 *ec2.DeleteTagsInput, param2 ...request.Option) (*ec2.DeleteTagsOutput, error)
	DeleteVolumeFunc                                          func(param0 *ec2.DeleteVolumeInput) (*ec2.DeleteVolumeOutput, error)
	DeleteVolumeRequestFunc                                   func(param0 *ec2.DeleteVolumeInput) (*request.Request, *ec2.DeleteVolumeOutput)
	DeleteVolumeWithContextFunc                               func(param0 aws.Context, param1 *ec2.DeleteVolumeInput, param2 ...request.Option) (*ec2.DeleteVolumeOutput, error)
	DeleteVpcFunc                                             func(param0 *ec2.DeleteVpcInput) (*ec2.DeleteVpcOutput, error)
	DeleteVpcEndpointsFunc                                    func(param0 *ec2.DeleteVpcEndpointsInput) (*ec2.DeleteVpcEndpointsOutput, error)
	DeleteVpcEndpointsRequestFunc                             func(param0 *ec2.DeleteVpcEndpointsInput) (*request.Request, *ec2.DeleteVpcEndpointsOutput)
	DeleteVpcEndpointsWithContextFunc                         func(param0 aws.Context, param1 *ec2.DeleteVpcEndpointsInput, param2 ...request.Option) (*ec2.DeleteVpcEndpointsOutput, error)
	DeleteVpcPeeringConnectionFunc                            func(param0 *ec2.DeleteVpcPeeringConnectionInput) (*ec2.DeleteVpcPeeringConnectionOutput, error)
	DeleteVpcPeeringConnectionRequestFunc                     func(param0 *ec2.DeleteVpcPeeringConnectionInput) (*request.Request, *ec2.DeleteVpcPeeringConnectionOutput)
	DeleteVpcPeeringConnectionWithContextFunc                 func(param0 aws.Context, param1 *ec2.DeleteVpcPeeringConnectionInput, param2 ...request.Option) (*ec2.DeleteVpcPeeringConnectionOutput, error)
	DeleteVpcRequestFunc                                      func(param0 *ec2.DeleteVpcInput) (*request.Request, *ec2.DeleteVpcOutput)
	DeleteVpcWithContextFunc                                  func(param0 aws.Context, param1 *ec2.DeleteVpcInput, param2 ...request.Option) (*ec2.DeleteVpcOutput, error)
	DeleteVpnConnectionFunc                                   func(param0 *ec2.DeleteVpnConnectionInput) (*ec2.DeleteVpnConnectionOutput, error)
	DeleteVpnConnectionRequestFunc                            func(param0 *ec2.DeleteVpnConnectionInput) (*request.Request, *ec2.DeleteVpnConnectionOutput)
	DeleteVpnConnectionRouteFunc                              func(param0 *ec2.DeleteVpnConnectionRouteInput) (*ec2.DeleteVpnConnectionRouteOutput, error)
	DeleteVpnConnectionRouteRequestFunc                       func(param0 *ec2.DeleteVpnConnectionRouteInput) (*request.Request, *ec2.DeleteVpnConnectionRouteOutput)
	DeleteVpnConnectionRouteWithContextFunc                   func(param0 aws.Context, param1 *ec2.DeleteVpnConnectionRouteInput, param2 ...request.Option) (*ec2.DeleteVpnConnectionRouteOutput, error)
	DeleteVpnConnectionWithContextFunc                        func(param0 aws.Context, param1 *ec2.DeleteVpnConnectionInput, param2 ...request.Option) (*ec2.DeleteVpnConnectionOutput, error)
	DeleteVpnGatewayFunc                                      func(param0 *ec2.DeleteVpnGatewayInput) (*ec2.DeleteVpnGatewayOutput, error)
	DeleteVpnGatewayRequestFunc                               func(param0 *ec2.DeleteVpnGatewayInput) (*request.Request, *ec2.DeleteVpnGatewayOutput)
	DeleteVpnGatewayWithContextFunc                           func(param0 aws.Context, param1 *ec2.DeleteVpnGatewayInput, param2 ...request.Option) (*ec2.DeleteVpnGatewayOutput, error)
	DeregisterImageFunc                                       func(param0 *ec2.DeregisterImageInput) (*ec2.DeregisterImageOutput, error)
	DeregisterImageRequestFunc                                func(param0 *ec2.DeregisterImageInput) (*request.Request, *ec2.DeregisterImageOutput)
	DeregisterImageWithContextFunc                            func(param0 aws.Context, param1 *ec2.DeregisterImageInput, param2 ...request.Option) (*ec2.DeregisterImageOutput, error)
	DescribeAccountAttributesFunc                             func(param0 *ec2.DescribeAccountAttributesInput) (*ec2.DescribeAccountAttributesOutput, error)
	DescribeAccountAttributesRequestFunc                      func(param0 *ec2.DescribeAccountAttributesInput) (*request.Request, *ec2.DescribeAccountAttributesOutput)
	DescribeAccountAttributesWithContextFunc                  func(param0 aws.Context, param1 *ec2.DescribeAccountAttributesInput, param2 ...request.Option) (*ec2.DescribeAccountAttributesOutput, error)
	DescribeAddressesFunc                                     func(param0 *ec2.DescribeAddressesInput) (*ec2.DescribeAddressesOutput, error)
	DescribeAddressesRequestFunc                              func(param0 *ec2.DescribeAddressesInput) (*request.Request, *ec2.DescribeAddressesOutput)
	DescribeAddressesWithContextFunc                          func(param0 aws.Context, param1 *ec2.DescribeAddressesInput, param2 ...request.Option) (*ec2.DescribeAddressesOutput, error)
	DescribeAvailabilityZonesFunc                             func(param0 *ec2.DescribeAvailabilityZonesInput) (*ec2.DescribeAvailabilityZonesOutput, error)
	DescribeAvailabilityZonesRequestFunc                      func(param0 *ec2.DescribeAvailabilityZonesInput) (*request.Request, *ec2.DescribeAvailabilityZonesOutput)
	DescribeAvailabilityZonesWithContextFunc                  func(param0 aws.Context, param1 *ec2.DescribeAvailabilityZonesInput, param2 ...request.Option) (*ec2.DescribeAvailabilityZonesOutput, error)
	DescribeBundleTasksFunc                                   func(param0 *ec2.DescribeBundleTasksInput) (*ec2.DescribeBundleTasksOutput, error)
	DescribeBundleTasksRequestFunc                            func(param0 *ec2.DescribeBundleTasksInput) (*request.Request, *ec2.DescribeBundleTasksOutput)
	DescribeBundleTasksWithContextFunc                        func(param0 aws.Context, param1 *ec2.DescribeBundleTasksInput, param2 ...request.Option) (*ec2.DescribeBundleTasksOutput, error)
	DescribeClassicLinkInstancesFunc                          func(param0 *ec2.DescribeClassicLinkInstancesInput) (*ec2.DescribeClassicLinkInstancesOutput, error)
	DescribeClassicLinkInstancesRequestFunc                   func(param0 *ec2.DescribeClassicLinkInstancesInput) (*request.Request, *ec2.DescribeClassicLinkInstancesOutput)
	DescribeClassicLinkInstancesWithContextFunc               func(param0 aws.Context, param1 *ec2.DescribeClassicLinkInstancesInput, param2 ...request.Option) (*ec2.DescribeClassicLinkInstancesOutput, error)
	DescribeConversionTasksFunc                               func(param0 *ec2.DescribeConversionTasksInput) (*ec2.DescribeConversionTasksOutput, error)
	DescribeConversionTasksRequestFunc                        func(param0 *ec2.DescribeConversionTasksInput) (*request.Request, *ec2.DescribeConversionTasksOutput)
	DescribeConversionTasksWithContextFunc                    func(param0 aws.Context, param1 *ec2.DescribeConversionTasksInput, param2 ...request.Option) (*ec2.DescribeConversionTasksOutput, error)
	DescribeCustomerGatewaysFunc                              func(param0 *ec2.DescribeCustomerGatewaysInput) (*ec2.DescribeCustomerGatewaysOutput, error)
	DescribeCustomerGatewaysRequestFunc                       func(param0 *ec2.DescribeCustomerGatewaysInput) (*request.Request, *ec2.DescribeCustomerGatewaysOutput)
	DescribeCustomerGatewaysWithContextFunc                   func(param0 aws.Context, param1 *ec2.DescribeCustomerGatewaysInput, param2 ...request.Option) (*ec2.DescribeCustomerGatewaysOutput, error)
	DescribeDhcpOptionsFunc                                   func(param0 *ec2.DescribeDhcpOptionsInput) (*ec2.DescribeDhcpOptionsOutput, error)
	DescribeDhcpOptionsRequestFunc                            func(param0 *ec2.DescribeDhcpOptionsInput) (*request.Request, *ec2.DescribeDhcpOptionsOutput)
	DescribeDhcpOptionsWithContextFunc                        func(param0 aws.Context, param1 *ec2.DescribeDhcpOptionsInput, param2 ...request.Option) (*ec2.DescribeDhcpOptionsOutput, error)
	DescribeEgressOnlyInternetGatewaysFunc                    func(param0 *ec2.DescribeEgressOnlyInternetGatewaysInput) (*ec2.DescribeEgressOnlyInternetGatewaysOutput, error)
	DescribeEgressOnlyInternetGatewaysRequestFunc             func(param0 *ec2.DescribeEgressOnlyInternetGatewaysInput) (*request.Request, *ec2.DescribeEgressOnlyInternetGatewaysOutput)
	DescribeEgressOnlyInternetGatewaysWithContextFunc         func(param0 aws.Context, param1 *ec2.DescribeEgressOnlyInternetGatewaysInput, param2 ...request.Option) (*ec2.DescribeEgressOnlyInternetGatewaysOutput, error)
	DescribeElasticGpusFunc                                   func(param0 *ec2.DescribeElasticGpusInput) (*ec2.DescribeElasticGpusOutput, error)
	DescribeElasticGpusRequestFunc                            func(param0 *ec2.DescribeElasticGpusInput) (*request.Request, *ec2.DescribeElasticGpusOutput)
	DescribeElasticGpusWithContextFunc                        func(param0 aws.Context, param1 *ec2.DescribeElasticGpusInput, param2 ...request.Option) (*ec2.DescribeElasticGpusOutput, error)
	DescribeExportTasksFunc                                   func(param0 *ec2.DescribeExportTasksInput) (*ec2.DescribeExportTasksOutput, error)
	DescribeExportTasksRequestFunc                            func(param0 *ec2.DescribeExportTasksInput) (*request.Request, *ec2.DescribeExportTasksOutput)
	DescribeExportTasksWithContextFunc                        func(param0 aws.Context, param1 *ec2.DescribeExportTasksInput, param2 ...request.Option) (*ec2.DescribeExportTasksOutput, error)
	DescribeFlowLogsFunc                                      func(param0 *ec2.DescribeFlowLogsInput) (*ec2.DescribeFlowLogsOutput, error)
	DescribeFlowLogsRequestFunc                               func(param0 *ec2.DescribeFlowLogsInput) (*request.Request, *ec2.DescribeFlowLogsOutput)
	DescribeFlowLogsWithContextFunc                           func(param0 aws.Context, param1 *ec2.DescribeFlowLogsInput, param2 ...request.Option) (*ec2.DescribeFlowLogsOutput, error)
	DescribeFpgaImageAttributeFunc                            func(param0 *ec2.DescribeFpgaImageAttributeInput) (*ec2.DescribeFpgaImageAttributeOutput, error)
	DescribeFpgaImageAttributeRequestFunc                     func(param0 *ec2.DescribeFpgaImageAttributeInput) (*request.Request, *ec2.DescribeFpgaImageAttributeOutput)
	DescribeFpgaImageAttributeWithContextFunc                 func(param0 aws.Context, param1 *ec2.DescribeFpgaImageAttributeInput, param2 ...request.Option) (*ec2.DescribeFpgaImageAttributeOutput, error)
	DescribeFpgaImagesFunc                                    func(param0 *ec2.DescribeFpgaImagesInput) (*ec2.DescribeFpgaImagesOutput, error)
	DescribeFpgaImagesRequestFunc                             func(param0 *ec2.DescribeFpgaImagesInput) (*request.Request, *ec2.DescribeFpgaImagesOutput)
	DescribeFpgaImagesWithContextFunc                         func(param0 aws.Context, param1 *ec2.DescribeFpgaImagesInput, param2 ...request.Option) (*ec2.DescribeFpgaImagesOutput, error)
	DescribeHostReservationOfferingsFunc                      func(param0 *ec2.DescribeHostReservationOfferingsInput) (*ec2.DescribeHostReservationOfferingsOutput, error)
	DescribeHostReservationOfferingsRequestFunc               func(param0 *ec2.DescribeHostReservationOfferingsInput) (*request.Request, *ec2.DescribeHostReservationOfferingsOutput)
	DescribeHostReservationOfferingsWithContextFunc           func(param0 aws.Context, param1 *ec2.DescribeHostReservationOfferingsInput, param2 ...request.Option) (*ec2.DescribeHostReservationOfferingsOutput, error)
	DescribeHostReservationsFunc                              func(param0 *ec2.DescribeHostReservationsInput) (*ec2.DescribeHostReservationsOutput, error)
	DescribeHostReservationsRequestFunc                       func(param0 *ec2.DescribeHostReservationsInput) (*request.Request, *ec2.DescribeHostReservationsOutput)
	DescribeHostReservationsWithContextFunc                   func(param0 aws.Context, param1 *ec2.DescribeHostReservationsInput, param2 ...request.Option) (*ec2.DescribeHostReservationsOutput, error)
	DescribeHostsFunc                                         func(param0 *ec2.DescribeHostsInput) (*ec2.DescribeHostsOutput, error)
	DescribeHostsRequestFunc                                  func(param0 *ec2.DescribeHostsInput) (*request.Request, *ec2.DescribeHostsOutput)
	DescribeHostsWithContextFunc                              func(param0 aws.Context, param1 *ec2.DescribeHostsInput, param2 ...request.Option) (*ec2.DescribeHostsOutput, error)
	DescribeIamInstanceProfileAssociationsFunc                func(param0 *ec2.DescribeIamInstanceProfileAssociationsInput) (*ec2.DescribeIamInstanceProfileAssociationsOutput, error)
	DescribeIamInstanceProfileAssociationsRequestFunc         func(param0 *ec2.DescribeIamInstanceProfileAssociationsInput) (*request.Request, *ec2.DescribeIamInstanceProfileAssociationsOutput)
	DescribeIamInstanceProfileAssociationsWithContextFunc     func(param0 aws.Context, param1 *ec2.DescribeIamInstanceProfileAssociationsInput, param2 ...request.Option) (*ec2.DescribeIamInstanceProfileAssociationsOutput, error)
	DescribeIdFormatFunc                                      func(param0 *ec2.DescribeIdFormatInput) (*ec2.DescribeIdFormatOutput, error)
	DescribeIdFormatRequestFunc                               func(param0 *ec2.DescribeIdFormatInput) (*request.Request, *ec2.DescribeIdFormatOutput)
	DescribeIdFormatWithContextFunc                           func(param0 aws.Context, param1 *ec2.DescribeIdFormatInput, param2 ...request.Option) (*ec2.DescribeIdFormatOutput, error)
	DescribeIdentityIdFormatFunc                              func(param0 *ec2.DescribeIdentityIdFormatInput) (*ec2.DescribeIdentityIdFormatOutput, error)
	DescribeIdentityIdFormatRequestFunc                       func(param0 *ec2.DescribeIdentityIdFormatInput) (*request.Request, *ec2.DescribeIdentityIdFormatOutput)
	DescribeIdentityIdFormatWithContextFunc                   func(param0 aws.Context, param1 *ec2.DescribeIdentityIdFormatInput, param2 ...request.Option) (*ec2.DescribeIdentityIdFormatOutput, error)
	DescribeImageAttributeFunc                                func(param0 *ec2.DescribeImageAttributeInput) (*ec2.DescribeImageAttributeOutput, error)
	DescribeImageAttributeRequestFunc                         func(param0 *ec2.DescribeImageAttributeInput) (*request.Request, *ec2.DescribeImageAttributeOutput)
	DescribeImageAttributeWithContextFunc                     func(param0 aws.Context, param1 *ec2.DescribeImageAttributeInput, param2 ...request.Option) (*ec2.DescribeImageAttributeOutput, error)
	DescribeImagesFunc                                        func(param0 *ec2.DescribeImagesInput) (*ec2.DescribeImagesOutput, error)
	DescribeImagesRequestFunc                                 func(param0 *ec2.DescribeImagesInput) (*request.Request, *ec2.DescribeImagesOutput)
	DescribeImagesWithContextFunc                             func(param0 aws.Context, param1 *ec2.DescribeImagesInput, param2 ...request.Option) (*ec2.DescribeImagesOutput, error)
	DescribeImportImageTasksFunc                              func(param0 *ec2.DescribeImportImageTasksInput) (*ec2.DescribeImportImageTasksOutput, error)
	DescribeImportImageTasksRequestFunc                       func(param0 *ec2.DescribeImportImageTasksInput) (*request.Request, *ec2.DescribeImportImageTasksOutput)
	DescribeImportImageTasksWithContextFunc                   func(param0 aws.Context, param1 *ec2.DescribeImportImageTasksInput, param2 ...request.Option) (*ec2.DescribeImportImageTasksOutput, error)
	DescribeImportSnapshotTasksFunc                           func(param0 *ec2.DescribeImportSnapshotTasksInput) (*ec2.DescribeImportSnapshotTasksOutput, error)
	DescribeImportSnapshotTasksRequestFunc                    func(param0 *ec2.DescribeImportSnapshotTasksInput) (*request.Request, *ec2.DescribeImportSnapshotTasksOutput)
	DescribeImportSnapshotTasksWithContextFunc                func(param0 aws.Context, param1 *ec2.DescribeImportSnapshotTasksInput, param2 ...request.Option) (*ec2.DescribeImportSnapshotTasksOutput, error)
	DescribeInstanceAttributeFunc                             func(param0 *ec2.DescribeInstanceAttributeInput) (*ec2.DescribeInstanceAttributeOutput, error)
	DescribeInstanceAttributeRequestFunc                      func(param0 *ec2.DescribeInstanceAttributeInput) (*request.Request, *ec2.DescribeInstanceAttributeOutput)
	DescribeInstanceAttributeWithContextFunc                  func(param0 aws.Context, param1 *ec2.DescribeInstanceAttributeInput, param2 ...request.Option) (*ec2.DescribeInstanceAttributeOutput, error)
	DescribeInstanceStatusFunc                                func(param0 *ec2.DescribeInstanceStatusInput) (*ec2.DescribeInstanceStatusOutput, error)
	DescribeInstanceStatusRequestFunc                         func(param0 *ec2.DescribeInstanceStatusInput) (*request.Request, *ec2.DescribeInstanceStatusOutput)
	DescribeInstanceStatusWithContextFunc                     func(param0 aws.Context, param1 *ec2.DescribeInstanceStatusInput, param2 ...request.Option) (*ec2.DescribeInstanceStatusOutput, error)
	DescribeInstancesFunc                                     func(param0 *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error)
	DescribeInstancesRequestFunc                              func(param0 *ec2.DescribeInstancesInput) (*request.Request, *ec2.DescribeInstancesOutput)
	DescribeInstancesWithContextFunc                          func(param0 aws.Context, param1 *ec2.DescribeInstancesInput, param2 ...request.Option) (*ec2.DescribeInstancesOutput, error)
	DescribeInternetGatewaysFunc                              func(param0 *ec2.DescribeInternetGatewaysInput) (*ec2.DescribeInternetGatewaysOutput, error)
	DescribeInternetGatewaysRequestFunc                       func(param0 *ec2.DescribeInternetGatewaysInput) (*request.Request, *ec2.DescribeInternetGatewaysOutput)
	DescribeInternetGatewaysWithContextFunc                   func(param0 aws.Context, param1 *ec2.DescribeInternetGatewaysInput, param2 ...request.Option) (*ec2.DescribeInternetGatewaysOutput, error)
	DescribeKeyPairsFunc                                      func(param0 *ec2.DescribeKeyPairsInput) (*ec2.DescribeKeyPairsOutput, error)
	DescribeKeyPairsRequestFunc                               func(param0 *ec2.DescribeKeyPairsInput) (*request.Request, *ec2.DescribeKeyPairsOutput)
	DescribeKeyPairsWithContextFunc                           func(param0 aws.Context, param1 *ec2.DescribeKeyPairsInput, param2 ...request.Option) (*ec2.DescribeKeyPairsOutput, error)
	DescribeMovingAddressesFunc                               func(param0 *ec2.DescribeMovingAddressesInput) (*ec2.DescribeMovingAddressesOutput, error)
	DescribeMovingAddressesRequestFunc                        func(param0 *ec2.DescribeMovingAddressesInput) (*request.Request, *ec2.DescribeMovingAddressesOutput)
	DescribeMovingAddressesWithContextFunc                    func(param0 aws.Context, param1 *ec2.DescribeMovingAddressesInput, param2 ...request.Option) (*ec2.DescribeMovingAddressesOutput, error)
	DescribeNatGatewaysFunc                                   func(param0 *ec2.DescribeNatGatewaysInput) (*ec2.DescribeNatGatewaysOutput, error)
	DescribeNatGatewaysRequestFunc                            func(param0 *ec2.DescribeNatGatewaysInput) (*request.Request, *ec2.DescribeNatGatewaysOutput)
	DescribeNatGatewaysWithContextFunc                        func(param0 aws.Context, param1 *ec2.DescribeNatGatewaysInput, param2 ...request.Option) (*ec2.DescribeNatGatewaysOutput, error)
	DescribeNetworkAclsFunc                                   func(param0 *ec2.DescribeNetworkAclsInput) (*ec2.DescribeNetworkAclsOutput, error)
	DescribeNetworkAclsRequestFunc                            func(param0 *ec2.DescribeNetworkAclsInput) (*request.Request, *ec2.DescribeNetworkAclsOutput)
	DescribeNetworkAclsWithContextFunc                        func(param0 aws.Context, param1 *ec2.DescribeNetworkAclsInput, param2 ...request.Option) (*ec2.DescribeNetworkAclsOutput, error)
	DescribeNetworkInterfaceAttributeFunc                     func(param0 *ec2.DescribeNetworkInterfaceAttributeInput) (*ec2.DescribeNetworkInterfaceAttributeOutput, error)
	DescribeNetworkInterfaceAttributeRequestFunc              func(param0 *ec2.DescribeNetworkInterfaceAttributeInput) (*request.Request, *ec2.DescribeNetworkInterfaceAttributeOutput)
	DescribeNetworkInterfaceAttributeWithContextFunc          func(param0 aws.Context, param1 *ec2.DescribeNetworkInterfaceAttributeInput, param2 ...request.Option) (*ec2.DescribeNetworkInterfaceAttributeOutput, error)
	DescribeNetworkInterfacePermissionsFunc                   func(param0 *ec2.DescribeNetworkInterfacePermissionsInput) (*ec2.DescribeNetworkInterfacePermissionsOutput, error)
	DescribeNetworkInterfacePermissionsRequestFunc            func(param0 *ec2.DescribeNetworkInterfacePermissionsInput) (*request.Request, *ec2.DescribeNetworkInterfacePermissionsOutput)
	DescribeNetworkInterfacePermissionsWithContextFunc        func(param0 aws.Context, param1 *ec2.DescribeNetworkInterfacePermissionsInput, param2 ...request.Option) (*ec2.DescribeNetworkInterfacePermissionsOutput, error)
	DescribeNetworkInterfacesFunc                             func(param0 *ec2.DescribeNetworkInterfacesInput) (*ec2.DescribeNetworkInterfacesOutput, error)
	DescribeNetworkInterfacesRequestFunc                      func(param0 *ec2.DescribeNetworkInterfacesInput) (*request.Request, *ec2.DescribeNetworkInterfacesOutput)
	DescribeNetworkInterfacesWithContextFunc                  func(param0 aws.Context, param1 *ec2.DescribeNetworkInterfacesInput, param2 ...request.Option) (*ec2.DescribeNetworkInterfacesOutput, error)
	DescribePlacementGroupsFunc                               func(param0 *ec2.DescribePlacementGroupsInput) (*ec2.DescribePlacementGroupsOutput, error)
	DescribePlacementGroupsRequestFunc                        func(param0 *ec2.DescribePlacementGroupsInput) (*request.Request, *ec2.DescribePlacementGroupsOutput)
	DescribePlacementGroupsWithContextFunc                    func(param0 aws.Context, param1 *ec2.DescribePlacementGroupsInput, param2 ...request.Option) (*ec2.DescribePlacementGroupsOutput, error)
	DescribePrefixListsFunc                                   func(param0 *ec2.DescribePrefixListsInput) (*ec2.DescribePrefixListsOutput, error)
	DescribePrefixListsRequestFunc                            func(param0 *ec2.DescribePrefixListsInput) (*request.Request, *ec2.DescribePrefixListsOutput)
	DescribePrefixListsWithContextFunc                        func(param0 aws.Context, param1 *ec2.DescribePrefixListsInput, param2 ...request.Option) (*ec2.DescribePrefixListsOutput, error)
	DescribeRegionsFunc                                       func(param0 *ec2.DescribeRegionsInput) (*ec2.DescribeRegionsOutput, error)
	DescribeRegionsRequestFunc                                func(param0 *ec2.DescribeRegionsInput) (*request.Request, *ec2.DescribeRegionsOutput)
	DescribeRegionsWithContextFunc                            func(param0 aws.Context, param1 *ec2.DescribeRegionsInput, param2 ...request.Option) (*ec2.DescribeRegionsOutput, error)
	DescribeReservedInstancesFunc                             func(param0 *ec2.DescribeReservedInstancesInput) (*ec2.DescribeReservedInstancesOutput, error)
	DescribeReservedInstancesListingsFunc                     func(param0 *ec2.DescribeReservedInstancesListingsInput) (*ec2.DescribeReservedInstancesListingsOutput, error)
	DescribeReservedInstancesListingsRequestFunc              func(param0 *ec2.DescribeReservedInstancesListingsInput) (*request.Request, *ec2.DescribeReservedInstancesListingsOutput)
	DescribeReservedInstancesListingsWithContextFunc          func(param0 aws.Context, param1 *ec2.DescribeReservedInstancesListingsInput, param2 ...request.Option) (*ec2.DescribeReservedInstancesListingsOutput, error)
	DescribeReservedInstancesModificationsFunc                func(param0 *ec2.DescribeReservedInstancesModificationsInput) (*ec2.DescribeReservedInstancesModificationsOutput, error)
	DescribeReservedInstancesModificationsRequestFunc         func(param0 *ec2.DescribeReservedInstancesModificationsInput) (*request.Request, *ec2.DescribeReservedInstancesModificationsOutput)
	DescribeReservedInstancesModificationsWithContextFunc     func(param0 aws.Context, param1 *ec2.DescribeReservedInstancesModificationsInput, param2 ...request.Option) (*ec2.DescribeReservedInstancesModificationsOutput, error)
	DescribeReservedInstancesOfferingsFunc                    func(param0 *ec2.DescribeReservedInstancesOfferingsInput) (*ec2.DescribeReservedInstancesOfferingsOutput, error)
	DescribeReservedInstancesOfferingsRequestFunc             func(param0 *ec2.DescribeReservedInstancesOfferingsInput) (*request.Request, *ec2.DescribeReservedInstancesOfferingsOutput)
	DescribeReservedInstancesOfferingsWithContextFunc         func(param0 aws.Context, param1 *ec2.DescribeReservedInstancesOfferingsInput, param2 ...request.Option) (*ec2.DescribeReservedInstancesOfferingsOutput, error)
	DescribeReservedInstancesRequestFunc                      func(param0 *ec2.DescribeReservedInstancesInput) (*request.Request, *ec2.DescribeReservedInstancesOutput)
	DescribeReservedInstancesWithContextFunc                  func(param0 aws.Context, param1 *ec2.DescribeReservedInstancesInput, param2 ...request.Option) (*ec2.DescribeReservedInstancesOutput, error)
	DescribeRouteTablesFunc                                   func(param0 *ec2.DescribeRouteTablesInput) (*ec2.DescribeRouteTablesOutput, error)
	DescribeRouteTablesRequestFunc                            func(param0 *ec2.DescribeRouteTablesInput) (*request.Request, *ec2.DescribeRouteTablesOutput)
	DescribeRouteTablesWithContextFunc                        func(param0 aws.Context, param1 *ec2.DescribeRouteTablesInput, param2 ...request.Option) (*ec2.DescribeRouteTablesOutput, error)
	DescribeScheduledInstanceAvailabilityFunc                 func(param0 *ec2.DescribeScheduledInstanceAvailabilityInput) (*ec2.DescribeScheduledInstanceAvailabilityOutput, error)
	DescribeScheduledInstanceAvailabilityRequestFunc          func(param0 *ec2.DescribeScheduledInstanceAvailabilityInput) (*request.Request, *ec2.DescribeScheduledInstanceAvailabilityOutput)
	DescribeScheduledInstanceAvailabilityWithContextFunc      func(param0 aws.Context, param1 *ec2.DescribeScheduledInstanceAvailabilityInput, param2 ...request.Option) (*ec2.DescribeScheduledInstanceAvailabilityOutput, error)
	DescribeScheduledInstancesFunc                            func(param0 *ec2.DescribeScheduledInstancesInput) (*ec2.DescribeScheduledInstancesOutput, error)
	DescribeScheduledInstancesRequestFunc                     func(param0 *ec2.DescribeScheduledInstancesInput) (*request.Request, *ec2.DescribeScheduledInstancesOutput)
	DescribeScheduledInstancesWithContextFunc                 func(param0 aws.Context, param1 *ec2.DescribeScheduledInstancesInput, param2 ...request.Option) (*ec2.DescribeScheduledInstancesOutput, error)
	DescribeSecurityGroupReferencesFunc                       func(param0 *ec2.DescribeSecurityGroupReferencesInput) (*ec2.DescribeSecurityGroupReferencesOutput, error)
	DescribeSecurityGroupReferencesRequestFunc                func(param0 *ec2.DescribeSecurityGroupReferencesInput) (*request.Request, *ec2.DescribeSecurityGroupReferencesOutput)
	DescribeSecurityGroupReferencesWithContextFunc            func(param0 aws.Context, param1 *ec2.DescribeSecurityGroupReferencesInput, param2 ...request.Option) (*ec2.DescribeSecurityGroupReferencesOutput, error)
	DescribeSecurityGroupsFunc                                func(param0 *ec2.DescribeSecurityGroupsInput) (*ec2.DescribeSecurityGroupsOutput, error)
	DescribeSecurityGroupsRequestFunc                         func(param0 *ec2.DescribeSecurityGroupsInput) (*request.Request, *ec2.DescribeSecurityGroupsOutput)
	DescribeSecurityGroupsWithContextFunc                     func(param0 aws.Context, param1 *ec2.DescribeSecurityGroupsInput, param2 ...request.Option) (*ec2.DescribeSecurityGroupsOutput, error)
	DescribeSnapshotAttributeFunc                             func(param0 *ec2.DescribeSnapshotAttributeInput) (*ec2.DescribeSnapshotAttributeOutput, error)
	DescribeSnapshotAttributeRequestFunc                      func(param0 *ec2.DescribeSnapshotAttributeInput) (*request.Request, *ec2.DescribeSnapshotAttributeOutput)
	DescribeSnapshotAttributeWithContextFunc                  func(param0 aws.Context, param1 *ec2.DescribeSnapshotAttributeInput, param2 ...request.Option) (*ec2.DescribeSnapshotAttributeOutput, error)
	DescribeSnapshotsFunc                                     func(param0 *ec2.DescribeSnapshotsInput) (*ec2.DescribeSnapshotsOutput, error)
	DescribeSnapshotsRequestFunc                              func(param0 *ec2.DescribeSnapshotsInput) (*request.Request, *ec2.DescribeSnapshotsOutput)
	DescribeSnapshotsWithContextFunc                          func(param0 aws.Context, param1 *ec2.DescribeSnapshotsInput, param2 ...request.Option) (*ec2.DescribeSnapshotsOutput, error)
	DescribeSpotDatafeedSubscriptionFunc                      func(param0 *ec2.DescribeSpotDatafeedSubscriptionInput) (*ec2.DescribeSpotDatafeedSubscriptionOutput, error)
	DescribeSpotDatafeedSubscriptionRequestFunc               func(param0 *ec2.DescribeSpotDatafeedSubscriptionInput) (*request.Request, *ec2.DescribeSpotDatafeedSubscriptionOutput)
	DescribeSpotDatafeedSubscriptionWithContextFunc           func(param0 aws.Context, param1 *ec2.DescribeSpotDatafeedSubscriptionInput, param2 ...request.Option) (*ec2.DescribeSpotDatafeedSubscriptionOutput, error)
	DescribeSpotFleetInstancesFunc                            func(param0 *ec2.DescribeSpotFleetInstancesInput) (*ec2.DescribeSpotFleetInstancesOutput, error)
	DescribeSpotFleetInstancesRequestFunc                     func(param0 *ec2.DescribeSpotFleetInstancesInput) (*request.Request, *ec2.DescribeSpotFleetInstancesOutput)
	DescribeSpotFleetInstancesWithContextFunc                 func(param0 aws.Context, param1 *ec2.DescribeSpotFleetInstancesInput, param2 ...request.Option) (*ec2.DescribeSpotFleetInstancesOutput, error)
	DescribeSpotFleetRequestHistoryFunc                       func(param0 *ec2.DescribeSpotFleetRequestHistoryInput) (*ec2.DescribeSpotFleetRequestHistoryOutput, error)
	DescribeSpotFleetRequestHistoryRequestFunc                func(param0 *ec2.DescribeSpotFleetRequestHistoryInput) (*request.Request, *ec2.DescribeSpotFleetRequestHistoryOutput)
	DescribeSpotFleetRequestHistoryWithContextFunc            func(param0 aws.Context, param1 *ec2.DescribeSpotFleetRequestHistoryInput, param2 ...request.Option) (*ec2.DescribeSpotFleetRequestHistoryOutput, error)
	DescribeSpotFleetRequestsFunc                             func(param0 *ec2.DescribeSpotFleetRequestsInput) (*ec2.DescribeSpotFleetRequestsOutput, error)
	DescribeSpotFleetRequestsRequestFunc                      func(param0 *ec2.DescribeSpotFleetRequestsInput) (*request.Request, *ec2.DescribeSpotFleetRequestsOutput)
	DescribeSpotFleetRequestsWithContextFunc                  func(param0 aws.Context, param1 *ec2.DescribeSpotFleetRequestsInput, param2 ...request.Option) (*ec2.DescribeSpotFleetRequestsOutput, error)
	DescribeSpotInstanceRequestsFunc                          func(param0 *ec2.DescribeSpotInstanceRequestsInput) (*ec2.DescribeSpotInstanceRequestsOutput, error)
	DescribeSpotInstanceRequestsRequestFunc                   func(param0 *ec2.DescribeSpotInstanceRequestsInput) (*request.Request, *ec2.DescribeSpotInstanceRequestsOutput)
	DescribeSpotInstanceRequestsWithContextFunc               func(param0 aws.Context, param1 *ec2.DescribeSpotInstanceRequestsInput, param2 ...request.Option) (*ec2.DescribeSpotInstanceRequestsOutput, error)
	DescribeSpotPriceHistoryFunc                              func(param0 *ec2.DescribeSpotPriceHistoryInput) (*ec2.DescribeSpotPriceHistoryOutput, error)
	DescribeSpotPriceHistoryRequestFunc                       func(param0 *ec2.DescribeSpotPriceHistoryInput) (*request.Request, *ec2.DescribeSpotPriceHistoryOutput)
	DescribeSpotPriceHistoryWithContextFunc                   func(param0 aws.Context, param1 *ec2.DescribeSpotPriceHistoryInput, param2 ...request.Option) (*ec2.DescribeSpotPriceHistoryOutput, error)
	DescribeStaleSecurityGroupsFunc                           func(param0 *ec2.DescribeStaleSecurityGroupsInput) (*ec2.DescribeStaleSecurityGroupsOutput, error)
	DescribeStaleSecurityGroupsRequestFunc                    func(param0 *ec2.DescribeStaleSecurityGroupsInput) (*request.Request, *ec2.DescribeStaleSecurityGroupsOutput)
	DescribeStaleSecurityGroupsWithContextFunc                func(param0 aws.Context, param1 *ec2.DescribeStaleSecurityGroupsInput, param2 ...request.Option) (*ec2.DescribeStaleSecurityGroupsOutput, error)
	DescribeSubnetsFunc                                       func(param0 *ec2.DescribeSubnetsInput) (*ec2.DescribeSubnetsOutput, error)
	DescribeSubnetsRequestFunc                                func(param0 *ec2.DescribeSubnetsInput) (*request.Request, *ec2.DescribeSubnetsOutput)
	DescribeSubnetsWithContextFunc                            func(param0 aws.Context, param1 *ec2.DescribeSubnetsInput, param2 ...request.Option) (*ec2.DescribeSubnetsOutput, error)
	DescribeTagsFunc                                          func(param0 *ec2.DescribeTagsInput) (*ec2.DescribeTagsOutput, error)
	DescribeTagsRequestFunc                                   func(param0 *ec2.DescribeTagsInput) (*request.Request, *ec2.DescribeTagsOutput)
	DescribeTagsWithContextFunc                               func(param0 aws.Context, param1 *ec2.DescribeTagsInput, param2 ...request.Option) (*ec2.DescribeTagsOutput, error)
	DescribeVolumeAttributeFunc                               func(param0 *ec2.DescribeVolumeAttributeInput) (*ec2.DescribeVolumeAttributeOutput, error)
	DescribeVolumeAttributeRequestFunc                        func(param0 *ec2.DescribeVolumeAttributeInput) (*request.Request, *ec2.DescribeVolumeAttributeOutput)
	DescribeVolumeAttributeWithContextFunc                    func(param0 aws.Context, param1 *ec2.DescribeVolumeAttributeInput, param2 ...request.Option) (*ec2.DescribeVolumeAttributeOutput, error)
	DescribeVolumeStatusFunc                                  func(param0 *ec2.DescribeVolumeStatusInput) (*ec2.DescribeVolumeStatusOutput, error)
	DescribeVolumeStatusRequestFunc                           func(param0 *ec2.DescribeVolumeStatusInput) (*request.Request, *ec2.DescribeVolumeStatusOutput)
	DescribeVolumeStatusWithContextFunc                       func(param0 aws.Context, param1 *ec2.DescribeVolumeStatusInput, param2 ...request.Option) (*ec2.DescribeVolumeStatusOutput, error)
	DescribeVolumesFunc                                       func(param0 *ec2.DescribeVolumesInput) (*ec2.DescribeVolumesOutput, error)
	DescribeVolumesModificationsFunc                          func(param0 *ec2.DescribeVolumesModificationsInput) (*ec2.DescribeVolumesModificationsOutput, error)
	DescribeVolumesModificationsRequestFunc                   func(param0 *ec2.DescribeVolumesModificationsInput) (*request.Request, *ec2.DescribeVolumesModificationsOutput)
	DescribeVolumesModificationsWithContextFunc               func(param0 aws.Context, param1 *ec2.DescribeVolumesModificationsInput, param2 ...request.Option) (*ec2.DescribeVolumesModificationsOutput, error)
	DescribeVolumesRequestFunc                                func(param0 *ec2.DescribeVolumesInput) (*request.Request, *ec2.DescribeVolumesOutput)
	DescribeVolumesWithContextFunc                            func(param0 aws.Context, param1 *ec2.DescribeVolumesInput, param2 ...request.Option) (*ec2.DescribeVolumesOutput, error)
	DescribeVpcAttributeFunc                                  func(param0 *ec2.DescribeVpcAttributeInput) (*ec2.DescribeVpcAttributeOutput, error)
	DescribeVpcAttributeRequestFunc                           func(param0 *ec2.DescribeVpcAttributeInput) (*request.Request, *ec2.DescribeVpcAttributeOutput)
	DescribeVpcAttributeWithContextFunc                       func(param0 aws.Context, param1 *ec2.DescribeVpcAttributeInput, param2 ...request.Option) (*ec2.DescribeVpcAttributeOutput, error)
	DescribeVpcClassicLinkFunc                                func(param0 *ec2.DescribeVpcClassicLinkInput) (*ec2.DescribeVpcClassicLinkOutput, error)
	DescribeVpcClassicLinkDnsSupportFunc                      func(param0 *ec2.DescribeVpcClassicLinkDnsSupportInput) (*ec2.DescribeVpcClassicLinkDnsSupportOutput, error)
	DescribeVpcClassicLinkDnsSupportRequestFunc               func(param0 *ec2.DescribeVpcClassicLinkDnsSupportInput) (*request.Request, *ec2.DescribeVpcClassicLinkDnsSupportOutput)
	DescribeVpcClassicLinkDnsSupportWithContextFunc           func(param0 aws.Context, param1 *ec2.DescribeVpcClassicLinkDnsSupportInput, param2 ...request.Option) (*ec2.DescribeVpcClassicLinkDnsSupportOutput, error)
	DescribeVpcClassicLinkRequestFunc                         func(param0 *ec2.DescribeVpcClassicLinkInput) (*request.Request, *ec2.DescribeVpcClassicLinkOutput)
	DescribeVpcClassicLinkWithContextFunc                     func(param0 aws.Context, param1 *ec2.DescribeVpcClassicLinkInput, param2 ...request.Option) (*ec2.DescribeVpcClassicLinkOutput, error)
	DescribeVpcEndpointServicesFunc                           func(param0 *ec2.DescribeVpcEndpointServicesInput) (*ec2.DescribeVpcEndpointServicesOutput, error)
	DescribeVpcEndpointServicesRequestFunc                    func(param0 *ec2.DescribeVpcEndpointServicesInput) (*request.Request, *ec2.DescribeVpcEndpointServicesOutput)
	DescribeVpcEndpointServicesWithContextFunc                func(param0 aws.Context, param1 *ec2.DescribeVpcEndpointServicesInput, param2 ...request.Option) (*ec2.DescribeVpcEndpointServicesOutput, error)
	DescribeVpcEndpointsFunc                                  func(param0 *ec2.DescribeVpcEndpointsInput) (*ec2.DescribeVpcEndpointsOutput, error)
	DescribeVpcEndpointsRequestFunc                           func(param0 *ec2.DescribeVpcEndpointsInput) (*request.Request, *ec2.DescribeVpcEndpointsOutput)
	DescribeVpcEndpointsWithContextFunc                       func(param0 aws.Context, param1 *ec2.DescribeVpcEndpointsInput, param2 ...request.Option) (*ec2.DescribeVpcEndpointsOutput, error)
	DescribeVpcPeeringConnectionsFunc                         func(param0 *ec2.DescribeVpcPeeringConnectionsInput) (*ec2.DescribeVpcPeeringConnectionsOutput, error)
	DescribeVpcPeeringConnectionsRequestFunc                  func(param0 *ec2.DescribeVpcPeeringConnectionsInput) (*request.Request, *ec2.DescribeVpcPeeringConnectionsOutput)
	DescribeVpcPeeringConnectionsWithContextFunc              func(param0 aws.Context, param1 *ec2.DescribeVpcPeeringConnectionsInput, param2 ...request.Option) (*ec2.DescribeVpcPeeringConnectionsOutput, error)
	DescribeVpcsFunc                                          func(param0 *ec2.DescribeVpcsInput) (*ec2.DescribeVpcsOutput, error)
	DescribeVpcsRequestFunc                                   func(param0 *ec2.DescribeVpcsInput) (*request.Request, *ec2.DescribeVpcsOutput)
	DescribeVpcsWithContextFunc                               func(param0 aws.Context, param1 *ec2.DescribeVpcsInput, param2 ...request.Option) (*ec2.DescribeVpcsOutput, error)
	DescribeVpnConnectionsFunc                                func(param0 *ec2.DescribeVpnConnectionsInput) (*ec2.DescribeVpnConnectionsOutput, error)
	DescribeVpnConnectionsRequestFunc                         func(param0 *ec2.DescribeVpnConnectionsInput) (*request.Request, *ec2.DescribeVpnConnectionsOutput)
	DescribeVpnConnectionsWithContextFunc                     func(param0 aws.Context, param1 *ec2.DescribeVpnConnectionsInput, param2 ...request.Option) (*ec2.DescribeVpnConnectionsOutput, error)
	DescribeVpnGatewaysFunc                                   func(param0 *ec2.DescribeVpnGatewaysInput) (*ec2.DescribeVpnGatewaysOutput, error)
	DescribeVpnGatewaysRequestFunc                            func(param0 *ec2.DescribeVpnGatewaysInput) (*request.Request, *ec2.DescribeVpnGatewaysOutput)
	DescribeVpnGatewaysWithContextFunc                        func(param0 aws.Context, param1 *ec2.DescribeVpnGatewaysInput, param2 ...request.Option) (*ec2.DescribeVpnGatewaysOutput, error)
	DetachClassicLinkVpcFunc                                  func(param0 *ec2.DetachClassicLinkVpcInput) (*ec2.DetachClassicLinkVpcOutput, error)
	DetachClassicLinkVpcRequestFunc                           func(param0 *ec2.DetachClassicLinkVpcInput) (*request.Request, *ec2.DetachClassicLinkVpcOutput)
	DetachClassicLinkVpcWithContextFunc                       func(param0 aws.Context, param1 *ec2.DetachClassicLinkVpcInput, param2 ...request.Option) (*ec2.DetachClassicLinkVpcOutput, error)
	DetachInternetGatewayFunc                                 func(param0 *ec2.DetachInternetGatewayInput) (*ec2.DetachInternetGatewayOutput, error)
	DetachInternetGatewayRequestFunc                          func(param0 *ec2.DetachInternetGatewayInput) (*request.Request, *ec2.DetachInternetGatewayOutput)
	DetachInternetGatewayWithContextFunc                      func(param0 aws.Context, param1 *ec2.DetachInternetGatewayInput, param2 ...request.Option) (*ec2.DetachInternetGatewayOutput, error)
	DetachNetworkInterfaceFunc                                func(param0 *ec2.DetachNetworkInterfaceInput) (*ec2.DetachNetworkInterfaceOutput, error)
	DetachNetworkInterfaceRequestFunc                         func(param0 *ec2.DetachNetworkInterfaceInput) (*request.Request, *ec2.DetachNetworkInterfaceOutput)
	DetachNetworkInterfaceWithContextFunc                     func(param0 aws.Context, param1 *ec2.DetachNetworkInterfaceInput, param2 ...request.Option) (*ec2.DetachNetworkInterfaceOutput, error)
	DetachVolumeFunc                                          func(param0 *ec2.DetachVolumeInput) (*ec2.VolumeAttachment, error)
	DetachVolumeRequestFunc                                   func(param0 *ec2.DetachVolumeInput) (*request.Request, *ec2.VolumeAttachment)
	DetachVolumeWithContextFunc                               func(param0 aws.Context, param1 *ec2.DetachVolumeInput, param2 ...request.Option) (*ec2.VolumeAttachment, error)
	DetachVpnGatewayFunc                                      func(param0 *ec2.DetachVpnGatewayInput) (*ec2.DetachVpnGatewayOutput, error)
	DetachVpnGatewayRequestFunc                               func(param0 *ec2.DetachVpnGatewayInput) (*request.Request, *ec2.DetachVpnGatewayOutput)
	DetachVpnGatewayWithContextFunc                           func(param0 aws.Context, param1 *ec2.DetachVpnGatewayInput, param2 ...request.Option) (*ec2.DetachVpnGatewayOutput, error)
	DisableVgwRoutePropagationFunc                            func(param0 *ec2.DisableVgwRoutePropagationInput) (*ec2.DisableVgwRoutePropagationOutput, error)
	DisableVgwRoutePropagationRequestFunc                     func(param0 *ec2.DisableVgwRoutePropagationInput) (*request.Request, *ec2.DisableVgwRoutePropagationOutput)
	DisableVgwRoutePropagationWithContextFunc                 func(param0 aws.Context, param1 *ec2.DisableVgwRoutePropagationInput, param2 ...request.Option) (*ec2.DisableVgwRoutePropagationOutput, error)
	DisableVpcClassicLinkFunc                                 func(param0 *ec2.DisableVpcClassicLinkInput) (*ec2.DisableVpcClassicLinkOutput, error)
	DisableVpcClassicLinkDnsSupportFunc                       func(param0 *ec2.DisableVpcClassicLinkDnsSupportInput) (*ec2.DisableVpcClassicLinkDnsSupportOutput, error)
	DisableVpcClassicLinkDnsSupportRequestFunc                func(param0 *ec2.DisableVpcClassicLinkDnsSupportInput) (*request.Request, *ec2.DisableVpcClassicLinkDnsSupportOutput)
	DisableVpcClassicLinkDnsSupportWithContextFunc            func(param0 aws.Context, param1 *ec2.DisableVpcClassicLinkDnsSupportInput, param2 ...request.Option) (*ec2.DisableVpcClassicLinkDnsSupportOutput, error)
	DisableVpcClassicLinkRequestFunc                          func(param0 *ec2.DisableVpcClassicLinkInput) (*request.Request, *ec2.DisableVpcClassicLinkOutput)
	DisableVpcClassicLinkWithContextFunc                      func(param0 aws.Context, param1 *ec2.DisableVpcClassicLinkInput, param2 ...request.Option) (*ec2.DisableVpcClassicLinkOutput, error)
	DisassociateAddressFunc                                   func(param0 *ec2.DisassociateAddressInput) (*ec2.DisassociateAddressOutput, error)
	DisassociateAddressRequestFunc                            func(param0 *ec2.DisassociateAddressInput) (*request.Request, *ec2.DisassociateAddressOutput)
	DisassociateAddressWithContextFunc                        func(param0 aws.Context, param1 *ec2.DisassociateAddressInput, param2 ...request.Option) (*ec2.DisassociateAddressOutput, error)
	DisassociateIamInstanceProfileFunc                        func(param0 *ec2.DisassociateIamInstanceProfileInput) (*ec2.DisassociateIamInstanceProfileOutput, error)
	DisassociateIamInstanceProfileRequestFunc                 func(param0 *ec2.DisassociateIamInstanceProfileInput) (*request.Request, *ec2.DisassociateIamInstanceProfileOutput)
	DisassociateIamInstanceProfileWithContextFunc             func(param0 aws.Context, param1 *ec2.DisassociateIamInstanceProfileInput, param2 ...request.Option) (*ec2.DisassociateIamInstanceProfileOutput, error)
	DisassociateRouteTableFunc                                func(param0 *ec2.DisassociateRouteTableInput) (*ec2.DisassociateRouteTableOutput, error)
	DisassociateRouteTableRequestFunc                         func(param0 *ec2.DisassociateRouteTableInput) (*request.Request, *ec2.DisassociateRouteTableOutput)
	DisassociateRouteTableWithContextFunc                     func(param0 aws.Context, param1 *ec2.DisassociateRouteTableInput, param2 ...request.Option) (*ec2.DisassociateRouteTableOutput, error)
	DisassociateSubnetCidrBlockFunc                           func(param0 *ec2.DisassociateSubnetCidrBlockInput) (*ec2.DisassociateSubnetCidrBlockOutput, error)
	DisassociateSubnetCidrBlockRequestFunc                    func(param0 *ec2.DisassociateSubnetCidrBlockInput) (*request.Request, *ec2.DisassociateSubnetCidrBlockOutput)
	DisassociateSubnetCidrBlockWithContextFunc                func(param0 aws.Context, param1 *ec2.DisassociateSubnetCidrBlockInput, param2 ...request.Option) (*ec2.DisassociateSubnetCidrBlockOutput, error)
	DisassociateVpcCidrBlockFunc                              func(param0 *ec2.DisassociateVpcCidrBlockInput) (*ec2.DisassociateVpcCidrBlockOutput, error)
	DisassociateVpcCidrBlockRequestFunc                       func(param0 *ec2.DisassociateVpcCidrBlockInput) (*request.Request, *ec2.DisassociateVpcCidrBlockOutput)
	DisassociateVpcCidrBlockWithContextFunc                   func(param0 aws.Context, param1 *ec2.DisassociateVpcCidrBlockInput, param2 ...request.Option) (*ec2.DisassociateVpcCidrBlockOutput, error)
	EnableVgwRoutePropagationFunc                             func(param0 *ec2.EnableVgwRoutePropagationInput) (*ec2.EnableVgwRoutePropagationOutput, error)
	EnableVgwRoutePropagationRequestFunc                      func(param0 *ec2.EnableVgwRoutePropagationInput) (*request.Request, *ec2.EnableVgwRoutePropagationOutput)
	EnableVgwRoutePropagationWithContextFunc                  func(param0 aws.Context, param1 *ec2.EnableVgwRoutePropagationInput, param2 ...request.Option) (*ec2.EnableVgwRoutePropagationOutput, error)
	EnableVolumeIOFunc                                        func(param0 *ec2.EnableVolumeIOInput) (*ec2.EnableVolumeIOOutput, error)
	EnableVolumeIORequestFunc                                 func(param0 *ec2.EnableVolumeIOInput) (*request.Request, *ec2.EnableVolumeIOOutput)
	EnableVolumeIOWithContextFunc                             func(param0 aws.Context, param1 *ec2.EnableVolumeIOInput, param2 ...request.Option) (*ec2.EnableVolumeIOOutput, error)
	EnableVpcClassicLinkFunc                                  func(param0 *ec2.EnableVpcClassicLinkInput) (*ec2.EnableVpcClassicLinkOutput, error)
	EnableVpcClassicLinkDnsSupportFunc                        func(param0 *ec2.EnableVpcClassicLinkDnsSupportInput) (*ec2.EnableVpcClassicLinkDnsSupportOutput, error)
	EnableVpcClassicLinkDnsSupportRequestFunc                 func(param0 *ec2.EnableVpcClassicLinkDnsSupportInput) (*request.Request, *ec2.EnableVpcClassicLinkDnsSupportOutput)
	EnableVpcClassicLinkDnsSupportWithContextFunc             func(param0 aws.Context, param1 *ec2.EnableVpcClassicLinkDnsSupportInput, param2 ...request.Option) (*ec2.EnableVpcClassicLinkDnsSupportOutput, error)
	EnableVpcClassicLinkRequestFunc                           func(param0 *ec2.EnableVpcClassicLinkInput) (*request.Request, *ec2.EnableVpcClassicLinkOutput)
	EnableVpcClassicLinkWithContextFunc                       func(param0 aws.Context, param1 *ec2.EnableVpcClassicLinkInput, param2 ...request.Option) (*ec2.EnableVpcClassicLinkOutput, error)
	GetConsoleOutputFunc                                      func(param0 *ec2.GetConsoleOutputInput) (*ec2.GetConsoleOutputOutput, error)
	GetConsoleOutputRequestFunc                               func(param0 *ec2.GetConsoleOutputInput) (*request.Request, *ec2.GetConsoleOutputOutput)
	GetConsoleOutputWithContextFunc                           func(param0 aws.Context, param1 *ec2.GetConsoleOutputInput, param2 ...request.Option) (*ec2.GetConsoleOutputOutput, error)
	GetConsoleScreenshotFunc                                  func(param0 *ec2.GetConsoleScreenshotInput) (*ec2.GetConsoleScreenshotOutput, error)
	GetConsoleScreenshotRequestFunc                           func(param0 *ec2.GetConsoleScreenshotInput) (*request.Request, *ec2.GetConsoleScreenshotOutput)
	GetConsoleScreenshotWithContextFunc                       func(param0 aws.Context, param1 *ec2.GetConsoleScreenshotInput, param2 ...request.Option) (*ec2.GetConsoleScreenshotOutput, error)
	GetHostReservationPurchasePreviewFunc                     func(param0 *ec2.GetHostReservationPurchasePreviewInput) (*ec2.GetHostReservationPurchasePreviewOutput, error)
	GetHostReservationPurchasePreviewRequestFunc              func(param0 *ec2.GetHostReservationPurchasePreviewInput) (*request.Request, *ec2.GetHostReservationPurchasePreviewOutput)
	GetHostReservationPurchasePreviewWithContextFunc          func(param0 aws.Context, param1 *ec2.GetHostReservationPurchasePreviewInput, param2 ...request.Option) (*ec2.GetHostReservationPurchasePreviewOutput, error)
	GetPasswordDataFunc                                       func(param0 *ec2.GetPasswordDataInput) (*ec2.GetPasswordDataOutput, error)
	GetPasswordDataRequestFunc                                func(param0 *ec2.GetPasswordDataInput) (*request.Request, *ec2.GetPasswordDataOutput)
	GetPasswordDataWithContextFunc                            func(param0 aws.Context, param1 *ec2.GetPasswordDataInput, param2 ...request.Option) (*ec2.GetPasswordDataOutput, error)
	GetReservedInstancesExchangeQuoteFunc                     func(param0 *ec2.GetReservedInstancesExchangeQuoteInput) (*ec2.GetReservedInstancesExchangeQuoteOutput, error)
	GetReservedInstancesExchangeQuoteRequestFunc              func(param0 *ec2.GetReservedInstancesExchangeQuoteInput) (*request.Request, *ec2.GetReservedInstancesExchangeQuoteOutput)
	GetReservedInstancesExchangeQuoteWithContextFunc          func(param0 aws.Context, param1 *ec2.GetReservedInstancesExchangeQuoteInput, param2 ...request.Option) (*ec2.GetReservedInstancesExchangeQuoteOutput, error)
	ImportImageFunc                                           func(param0 *ec2.ImportImageInput) (*ec2.ImportImageOutput, error)
	ImportImageRequestFunc                                    func(param0 *ec2.ImportImageInput) (*request.Request, *ec2.ImportImageOutput)
	ImportImageWithContextFunc                                func(param0 aws.Context, param1 *ec2.ImportImageInput, param2 ...request.Option) (*ec2.ImportImageOutput, error)
	ImportInstanceFunc                                        func(param0 *ec2.ImportInstanceInput) (*ec2.ImportInstanceOutput, error)
	ImportInstanceRequestFunc                                 func(param0 *ec2.ImportInstanceInput) (*request.Request, *ec2.ImportInstanceOutput)
	ImportInstanceWithContextFunc                             func(param0 aws.Context, param1 *ec2.ImportInstanceInput, param2 ...request.Option) (*ec2.ImportInstanceOutput, error)
	ImportKeyPairFunc                                         func(param0 *ec2.ImportKeyPairInput) (*ec2.ImportKeyPairOutput, error)
	ImportKeyPairRequestFunc                                  func(param0 *ec2.ImportKeyPairInput) (*request.Request, *ec2.ImportKeyPairOutput)
	ImportKeyPairWithContextFunc                              func(param0 aws.Context, param1 *ec2.ImportKeyPairInput, param2 ...request.Option) (*ec2.ImportKeyPairOutput, error)
	ImportSnapshotFunc                                        func(param0 *ec2.ImportSnapshotInput) (*ec2.ImportSnapshotOutput, error)
	ImportSnapshotRequestFunc                                 func(param0 *ec2.ImportSnapshotInput) (*request.Request, *ec2.ImportSnapshotOutput)
	ImportSnapshotWithContextFunc                             func(param0 aws.Context, param1 *ec2.ImportSnapshotInput, param2 ...request.Option) (*ec2.ImportSnapshotOutput, error)
	ImportVolumeFunc                                          func(param0 *ec2.ImportVolumeInput) (*ec2.ImportVolumeOutput, error)
	ImportVolumeRequestFunc                                   func(param0 *ec2.ImportVolumeInput) (*request.Request, *ec2.ImportVolumeOutput)
	ImportVolumeWithContextFunc                               func(param0 aws.Context, param1 *ec2.ImportVolumeInput, param2 ...request.Option) (*ec2.ImportVolumeOutput, error)
	ModifyFpgaImageAttributeFunc                              func(param0 *ec2.ModifyFpgaImageAttributeInput) (*ec2.ModifyFpgaImageAttributeOutput, error)
	ModifyFpgaImageAttributeRequestFunc                       func(param0 *ec2.ModifyFpgaImageAttributeInput) (*request.Request, *ec2.ModifyFpgaImageAttributeOutput)
	ModifyFpgaImageAttributeWithContextFunc                   func(param0 aws.Context, param1 *ec2.ModifyFpgaImageAttributeInput, param2 ...request.Option) (*ec2.ModifyFpgaImageAttributeOutput, error)
	ModifyHostsFunc                                           func(param0 *ec2.ModifyHostsInput) (*ec2.ModifyHostsOutput, error)
	ModifyHostsRequestFunc                                    func(param0 *ec2.ModifyHostsInput) (*request.Request, *ec2.ModifyHostsOutput)
	ModifyHostsWithContextFunc                                func(param0 aws.Context, param1 *ec2.ModifyHostsInput, param2 ...request.Option) (*ec2.ModifyHostsOutput, error)
	ModifyIdFormatFunc                                        func(param0 *ec2.ModifyIdFormatInput) (*ec2.ModifyIdFormatOutput, error)
	ModifyIdFormatRequestFunc                                 func(param0 *ec2.ModifyIdFormatInput) (*request.Request, *ec2.ModifyIdFormatOutput)
	ModifyIdFormatWithContextFunc                             func(param0 aws.Context, param1 *ec2.ModifyIdFormatInput, param2 ...request.Option) (*ec2.ModifyIdFormatOutput, error)
	ModifyIdentityIdFormatFunc                                func(param0 *ec2.ModifyIdentityIdFormatInput) (*ec2.ModifyIdentityIdFormatOutput, error)
	ModifyIdentityIdFormatRequestFunc                         func(param0 *ec2.ModifyIdentityIdFormatInput) (*request.Request, *ec2.ModifyIdentityIdFormatOutput)
	ModifyIdentityIdFormatWithContextFunc                     func(param0 aws.Context, param1 *ec2.ModifyIdentityIdFormatInput, param2 ...request.Option) (*ec2.ModifyIdentityIdFormatOutput, error)
	ModifyImageAttributeFunc                                  func(param0 *ec2.ModifyImageAttributeInput) (*ec2.ModifyImageAttributeOutput, error)
	ModifyImageAttributeRequestFunc                           func(param0 *ec2.ModifyImageAttributeInput) (*request.Request, *ec2.ModifyImageAttributeOutput)
	ModifyImageAttributeWithContextFunc                       func(param0 aws.Context, param1 *ec2.ModifyImageAttributeInput, param2 ...request.Option) (*ec2.ModifyImageAttributeOutput, error)
	ModifyInstanceAttributeFunc                               func(param0 *ec2.ModifyInstanceAttributeInput) (*ec2.ModifyInstanceAttributeOutput, error)
	ModifyInstanceAttributeRequestFunc                        func(param0 *ec2.ModifyInstanceAttributeInput) (*request.Request, *ec2.ModifyInstanceAttributeOutput)
	ModifyInstanceAttributeWithContextFunc                    func(param0 aws.Context, param1 *ec2.ModifyInstanceAttributeInput, param2 ...request.Option) (*ec2.ModifyInstanceAttributeOutput, error)
	ModifyInstancePlacementFunc                               func(param0 *ec2.ModifyInstancePlacementInput) (*ec2.ModifyInstancePlacementOutput, error)
	ModifyInstancePlacementRequestFunc                        func(param0 *ec2.ModifyInstancePlacementInput) (*request.Request, *ec2.ModifyInstancePlacementOutput)
	ModifyInstancePlacementWithContextFunc                    func(param0 aws.Context, param1 *ec2.ModifyInstancePlacementInput, param2 ...request.Option) (*ec2.ModifyInstancePlacementOutput, error)
	ModifyNetworkInterfaceAttributeFunc                       func(param0 *ec2.ModifyNetworkInterfaceAttributeInput) (*ec2.ModifyNetworkInterfaceAttributeOutput, error)
	ModifyNetworkInterfaceAttributeRequestFunc                func(param0 *ec2.ModifyNetworkInterfaceAttributeInput) (*request.Request, *ec2.ModifyNetworkInterfaceAttributeOutput)
	ModifyNetworkInterfaceAttributeWithContextFunc            func(param0 aws.Context, param1 *ec2.ModifyNetworkInterfaceAttributeInput, param2 ...request.Option) (*ec2.ModifyNetworkInterfaceAttributeOutput, error)
	ModifyReservedInstancesFunc                               func(param0 *ec2.ModifyReservedInstancesInput) (*ec2.ModifyReservedInstancesOutput, error)
	ModifyReservedInstancesRequestFunc                        func(param0 *ec2.ModifyReservedInstancesInput) (*request.Request, *ec2.ModifyReservedInstancesOutput)
	ModifyReservedInstancesWithContextFunc                    func(param0 aws.Context, param1 *ec2.ModifyReservedInstancesInput, param2 ...request.Option) (*ec2.ModifyReservedInstancesOutput, error)
	ModifySnapshotAttributeFunc                               func(param0 *ec2.ModifySnapshotAttributeInput) (*ec2.ModifySnapshotAttributeOutput, error)
	ModifySnapshotAttributeRequestFunc                        func(param0 *ec2.ModifySnapshotAttributeInput) (*request.Request, *ec2.ModifySnapshotAttributeOutput)
	ModifySnapshotAttributeWithContextFunc                    func(param0 aws.Context, param1 *ec2.ModifySnapshotAttributeInput, param2 ...request.Option) (*ec2.ModifySnapshotAttributeOutput, error)
	ModifySpotFleetRequestFunc                                func(param0 *ec2.ModifySpotFleetRequestInput) (*ec2.ModifySpotFleetRequestOutput, error)
	ModifySpotFleetRequestRequestFunc                         func(param0 *ec2.ModifySpotFleetRequestInput) (*request.Request, *ec2.ModifySpotFleetRequestOutput)
	ModifySpotFleetRequestWithContextFunc                     func(param0 aws.Context, param1 *ec2.ModifySpotFleetRequestInput, param2 ...request.Option) (*ec2.ModifySpotFleetRequestOutput, error)
	ModifySubnetAttributeFunc                                 func(param0 *ec2.ModifySubnetAttributeInput) (*ec2.ModifySubnetAttributeOutput, error)
	ModifySubnetAttributeRequestFunc                          func(param0 *ec2.ModifySubnetAttributeInput) (*request.Request, *ec2.ModifySubnetAttributeOutput)
	ModifySubnetAttributeWithContextFunc                      func(param0 aws.Context, param1 *ec2.ModifySubnetAttributeInput, param2 ...request.Option) (*ec2.ModifySubnetAttributeOutput, error)
	ModifyVolumeFunc                                          func(param0 *ec2.ModifyVolumeInput) (*ec2.ModifyVolumeOutput, error)
	ModifyVolumeAttributeFunc                                 func(param0 *ec2.ModifyVolumeAttributeInput) (*ec2.ModifyVolumeAttributeOutput, error)
	ModifyVolumeAttributeRequestFunc                          func(param0 *ec2.ModifyVolumeAttributeInput) (*request.Request, *ec2.ModifyVolumeAttributeOutput)
	ModifyVolumeAttributeWithContextFunc                      func(param0 aws.Context, param1 *ec2.ModifyVolumeAttributeInput, param2 ...request.Option) (*ec2.ModifyVolumeAttributeOutput, error)
	ModifyVolumeRequestFunc                                   func(param0 *ec2.ModifyVolumeInput) (*request.Request, *ec2.ModifyVolumeOutput)
	ModifyVolumeWithContextFunc                               func(param0 aws.Context, param1 *ec2.ModifyVolumeInput, param2 ...request.Option) (*ec2.ModifyVolumeOutput, error)
	ModifyVpcAttributeFunc                                    func(param0 *ec2.ModifyVpcAttributeInput) (*ec2.ModifyVpcAttributeOutput, error)
	ModifyVpcAttributeRequestFunc                             func(param0 *ec2.ModifyVpcAttributeInput) (*request.Request, *ec2.ModifyVpcAttributeOutput)
	ModifyVpcAttributeWithContextFunc                         func(param0 aws.Context, param1 *ec2.ModifyVpcAttributeInput, param2 ...request.Option) (*ec2.ModifyVpcAttributeOutput, error)
	ModifyVpcEndpointFunc                                     func(param0 *ec2.ModifyVpcEndpointInput) (*ec2.ModifyVpcEndpointOutput, error)
	ModifyVpcEndpointRequestFunc                              func(param0 *ec2.ModifyVpcEndpointInput) (*request.Request, *ec2.ModifyVpcEndpointOutput)
	ModifyVpcEndpointWithContextFunc                          func(param0 aws.Context, param1 *ec2.ModifyVpcEndpointInput, param2 ...request.Option) (*ec2.ModifyVpcEndpointOutput, error)
	ModifyVpcPeeringConnectionOptionsFunc                     func(param0 *ec2.ModifyVpcPeeringConnectionOptionsInput) (*ec2.ModifyVpcPeeringConnectionOptionsOutput, error)
	ModifyVpcPeeringConnectionOptionsRequestFunc              func(param0 *ec2.ModifyVpcPeeringConnectionOptionsInput) (*request.Request, *ec2.ModifyVpcPeeringConnectionOptionsOutput)
	ModifyVpcPeeringConnectionOptionsWithContextFunc          func(param0 aws.Context, param1 *ec2.ModifyVpcPeeringConnectionOptionsInput, param2 ...request.Option) (*ec2.ModifyVpcPeeringConnectionOptionsOutput, error)
	MonitorInstancesFunc                                      func(param0 *ec2.MonitorInstancesInput) (*ec2.MonitorInstancesOutput, error)
	MonitorInstancesRequestFunc                               func(param0 *ec2.MonitorInstancesInput) (*request.Request, *ec2.MonitorInstancesOutput)
	MonitorInstancesWithContextFunc                           func(param0 aws.Context, param1 *ec2.MonitorInstancesInput, param2 ...request.Option) (*ec2.MonitorInstancesOutput, error)
	MoveAddressToVpcFunc                                      func(param0 *ec2.MoveAddressToVpcInput) (*ec2.MoveAddressToVpcOutput, error)
	MoveAddressToVpcRequestFunc                               func(param0 *ec2.MoveAddressToVpcInput) (*request.Request, *ec2.MoveAddressToVpcOutput)
	MoveAddressToVpcWithContextFunc                           func(param0 aws.Context, param1 *ec2.MoveAddressToVpcInput, param2 ...request.Option) (*ec2.MoveAddressToVpcOutput, error)
	PurchaseHostReservationFunc                               func(param0 *ec2.PurchaseHostReservationInput) (*ec2.PurchaseHostReservationOutput, error)
	PurchaseHostReservationRequestFunc                        func(param0 *ec2.PurchaseHostReservationInput) (*request.Request, *ec2.PurchaseHostReservationOutput)
	PurchaseHostReservationWithContextFunc                    func(param0 aws.Context, param1 *ec2.PurchaseHostReservationInput, param2 ...request.Option) (*ec2.PurchaseHostReservationOutput, error)
	PurchaseReservedInstancesOfferingFunc                     func(param0 *ec2.PurchaseReservedInstancesOfferingInput) (*ec2.PurchaseReservedInstancesOfferingOutput, error)
	PurchaseReservedInstancesOfferingRequestFunc              func(param0 *ec2.PurchaseReservedInstancesOfferingInput) (*request.Request, *ec2.PurchaseReservedInstancesOfferingOutput)
	PurchaseReservedInstancesOfferingWithContextFunc          func(param0 aws.Context, param1 *ec2.PurchaseReservedInstancesOfferingInput, param2 ...request.Option) (*ec2.PurchaseReservedInstancesOfferingOutput, error)
	PurchaseScheduledInstancesFunc                            func(param0 *ec2.PurchaseScheduledInstancesInput) (*ec2.PurchaseScheduledInstancesOutput, error)
	PurchaseScheduledInstancesRequestFunc                     func(param0 *ec2.PurchaseScheduledInstancesInput) (*request.Request, *ec2.PurchaseScheduledInstancesOutput)
	PurchaseScheduledInstancesWithContextFunc                 func(param0 aws.Context, param1 *ec2.PurchaseScheduledInstancesInput, param2 ...request.Option) (*ec2.PurchaseScheduledInstancesOutput, error)
	RebootInstancesFunc                                       func(param0 *ec2.RebootInstancesInput) (*ec2.RebootInstancesOutput, error)
	RebootInstancesRequestFunc                                func(param0 *ec2.RebootInstancesInput) (*request.Request, *ec2.RebootInstancesOutput)
	RebootInstancesWithContextFunc                            func(param0 aws.Context, param1 *ec2.RebootInstancesInput, param2 ...request.Option) (*ec2.RebootInstancesOutput, error)
	RegisterImageFunc                                         func(param0 *ec2.RegisterImageInput) (*ec2.RegisterImageOutput, error)
	RegisterImageRequestFunc                                  func(param0 *ec2.RegisterImageInput) (*request.Request, *ec2.RegisterImageOutput)
	RegisterImageWithContextFunc                              func(param0 aws.Context, param1 *ec2.RegisterImageInput, param2 ...request.Option) (*ec2.RegisterImageOutput, error)
	RejectVpcPeeringConnectionFunc                            func(param0 *ec2.RejectVpcPeeringConnectionInput) (*ec2.RejectVpcPeeringConnectionOutput, error)
	RejectVpcPeeringConnectionRequestFunc                     func(param0 *ec2.RejectVpcPeeringConnectionInput) (*request.Request, *ec2.RejectVpcPeeringConnectionOutput)
	RejectVpcPeeringConnectionWithContextFunc                 func(param0 aws.Context, param1 *ec2.RejectVpcPeeringConnectionInput, param2 ...request.Option) (*ec2.RejectVpcPeeringConnectionOutput, error)
	ReleaseAddressFunc                                        func(param0 *ec2.ReleaseAddressInput) (*ec2.ReleaseAddressOutput, error)
	ReleaseAddressRequestFunc                                 func(param0 *ec2.ReleaseAddressInput) (*request.Request, *ec2.ReleaseAddressOutput)
	ReleaseAddressWithContextFunc                             func(param0 aws.Context, param1 *ec2.ReleaseAddressInput, param2 ...request.Option) (*ec2.ReleaseAddressOutput, error)
	ReleaseHostsFunc                                          func(param0 *ec2.ReleaseHostsInput) (*ec2.ReleaseHostsOutput, error)
	ReleaseHostsRequestFunc                                   func(param0 *ec2.ReleaseHostsInput) (*request.Request, *ec2.ReleaseHostsOutput)
	ReleaseHostsWithContextFunc                               func(param0 aws.Context, param1 *ec2.ReleaseHostsInput, param2 ...request.Option) (*ec2.ReleaseHostsOutput, error)
	ReplaceIamInstanceProfileAssociationFunc                  func(param0 *ec2.ReplaceIamInstanceProfileAssociationInput) (*ec2.ReplaceIamInstanceProfileAssociationOutput, error)
	ReplaceIamInstanceProfileAssociationRequestFunc           func(param0 *ec2.ReplaceIamInstanceProfileAssociationInput) (*request.Request, *ec2.ReplaceIamInstanceProfileAssociationOutput)
	ReplaceIamInstanceProfileAssociationWithContextFunc       func(param0 aws.Context, param1 *ec2.ReplaceIamInstanceProfileAssociationInput, param2 ...request.Option) (*ec2.ReplaceIamInstanceProfileAssociationOutput, error)
	ReplaceNetworkAclAssociationFunc                          func(param0 *ec2.ReplaceNetworkAclAssociationInput) (*ec2.ReplaceNetworkAclAssociationOutput, error)
	ReplaceNetworkAclAssociationRequestFunc                   func(param0 *ec2.ReplaceNetworkAclAssociationInput) (*request.Request, *ec2.ReplaceNetworkAclAssociationOutput)
	ReplaceNetworkAclAssociationWithContextFunc               func(param0 aws.Context, param1 *ec2.ReplaceNetworkAclAssociationInput, param2 ...request.Option) (*ec2.ReplaceNetworkAclAssociationOutput, error)
	ReplaceNetworkAclEntryFunc                                func(param0 *ec2.ReplaceNetworkAclEntryInput) (*ec2.ReplaceNetworkAclEntryOutput, error)
	ReplaceNetworkAclEntryRequestFunc                         func(param0 *ec2.ReplaceNetworkAclEntryInput) (*request.Request, *ec2.ReplaceNetworkAclEntryOutput)
	ReplaceNetworkAclEntryWithContextFunc                     func(param0 aws.Context, param1 *ec2.ReplaceNetworkAclEntryInput, param2 ...request.Option) (*ec2.ReplaceNetworkAclEntryOutput, error)
	ReplaceRouteFunc                                          func(param0 *ec2.ReplaceRouteInput) (*ec2.ReplaceRouteOutput, error)
	ReplaceRouteRequestFunc                                   func(param0 *ec2.ReplaceRouteInput) (*request.Request, *ec2.ReplaceRouteOutput)
	ReplaceRouteTableAssociationFunc                          func(param0 *ec2.ReplaceRouteTableAssociationInput) (*ec2.ReplaceRouteTableAssociationOutput, error)
	ReplaceRouteTableAssociationRequestFunc                   func(param0 *ec2.ReplaceRouteTableAssociationInput) (*request.Request, *ec2.ReplaceRouteTableAssociationOutput)
	ReplaceRouteTableAssociationWithContextFunc               func(param0 aws.Context, param1 *ec2.ReplaceRouteTableAssociationInput, param2 ...request.Option) (*ec2.ReplaceRouteTableAssociationOutput, error)
	ReplaceRouteWithContextFunc                               func(param0 aws.Context, param1 *ec2.ReplaceRouteInput, param2 ...request.Option) (*ec2.ReplaceRouteOutput, error)
	ReportInstanceStatusFunc                                  func(param0 *ec2.ReportInstanceStatusInput) (*ec2.ReportInstanceStatusOutput, error)
	ReportInstanceStatusRequestFunc                           func(param0 *ec2.ReportInstanceStatusInput) (*request.Request, *ec2.ReportInstanceStatusOutput)
	ReportInstanceStatusWithContextFunc                       func(param0 aws.Context, param1 *ec2.ReportInstanceStatusInput, param2 ...request.Option) (*ec2.ReportInstanceStatusOutput, error)
	RequestSpotFleetFunc                                      func(param0 *ec2.RequestSpotFleetInput) (*ec2.RequestSpotFleetOutput, error)
	RequestSpotFleetRequestFunc                               func(param0 *ec2.RequestSpotFleetInput) (*request.Request, *ec2.RequestSpotFleetOutput)
	RequestSpotFleetWithContextFunc                           func(param0 aws.Context, param1 *ec2.RequestSpotFleetInput, param2 ...request.Option) (*ec2.RequestSpotFleetOutput, error)
	RequestSpotInstancesFunc                                  func(param0 *ec2.RequestSpotInstancesInput) (*ec2.RequestSpotInstancesOutput, error)
	RequestSpotInstancesRequestFunc                           func(param0 *ec2.RequestSpotInstancesInput) (*request.Request, *ec2.RequestSpotInstancesOutput)
	RequestSpotInstancesWithContextFunc                       func(param0 aws.Context, param1 *ec2.RequestSpotInstancesInput, param2 ...request.Option) (*ec2.RequestSpotInstancesOutput, error)
	ResetFpgaImageAttributeFunc                               func(param0 *ec2.ResetFpgaImageAttributeInput) (*ec2.ResetFpgaImageAttributeOutput, error)
	ResetFpgaImageAttributeRequestFunc                        func(param0 *ec2.ResetFpgaImageAttributeInput) (*request.Request, *ec2.ResetFpgaImageAttributeOutput)
	ResetFpgaImageAttributeWithContextFunc                    func(param0 aws.Context, param1 *ec2.ResetFpgaImageAttributeInput, param2 ...request.Option) (*ec2.ResetFpgaImageAttributeOutput, error)
	ResetImageAttributeFunc                                   func(param0 *ec2.ResetImageAttributeInput) (*ec2.ResetImageAttributeOutput, error)
	ResetImageAttributeRequestFunc                            func(param0 *ec2.ResetImageAttributeInput) (*request.Request, *ec2.ResetImageAttributeOutput)
	ResetImageAttributeWithContextFunc                        func(param0 aws.Context, param1 *ec2.ResetImageAttributeInput, param2 ...request.Option) (*ec2.ResetImageAttributeOutput, error)
	ResetInstanceAttributeFunc                                func(param0 *ec2.ResetInstanceAttributeInput) (*ec2.ResetInstanceAttributeOutput, error)
	ResetInstanceAttributeRequestFunc                         func(param0 *ec2.ResetInstanceAttributeInput) (*request.Request, *ec2.ResetInstanceAttributeOutput)
	ResetInstanceAttributeWithContextFunc                     func(param0 aws.Context, param1 *ec2.ResetInstanceAttributeInput, param2 ...request.Option) (*ec2.ResetInstanceAttributeOutput, error)
	ResetNetworkInterfaceAttributeFunc                        func(param0 *ec2.ResetNetworkInterfaceAttributeInput) (*ec2.ResetNetworkInterfaceAttributeOutput, error)
	ResetNetworkInterfaceAttributeRequestFunc                 func(param0 *ec2.ResetNetworkInterfaceAttributeInput) (*request.Request, *ec2.ResetNetworkInterfaceAttributeOutput)
	ResetNetworkInterfaceAttributeWithContextFunc             func(param0 aws.Context, param1 *ec2.ResetNetworkInterfaceAttributeInput, param2 ...request.Option) (*ec2.ResetNetworkInterfaceAttributeOutput, error)
	ResetSnapshotAttributeFunc                                func(param0 *ec2.ResetSnapshotAttributeInput) (*ec2.ResetSnapshotAttributeOutput, error)
	ResetSnapshotAttributeRequestFunc                         func(param0 *ec2.ResetSnapshotAttributeInput) (*request.Request, *ec2.ResetSnapshotAttributeOutput)
	ResetSnapshotAttributeWithContextFunc                     func(param0 aws.Context, param1 *ec2.ResetSnapshotAttributeInput, param2 ...request.Option) (*ec2.ResetSnapshotAttributeOutput, error)
	RestoreAddressToClassicFunc                               func(param0 *ec2.RestoreAddressToClassicInput) (*ec2.RestoreAddressToClassicOutput, error)
	RestoreAddressToClassicRequestFunc                        func(param0 *ec2.RestoreAddressToClassicInput) (*request.Request, *ec2.RestoreAddressToClassicOutput)
	RestoreAddressToClassicWithContextFunc                    func(param0 aws.Context, param1 *ec2.RestoreAddressToClassicInput, param2 ...request.Option) (*ec2.RestoreAddressToClassicOutput, error)
	RevokeSecurityGroupEgressFunc                             func(param0 *ec2.RevokeSecurityGroupEgressInput) (*ec2.RevokeSecurityGroupEgressOutput, error)
	RevokeSecurityGroupEgressRequestFunc                      func(param0 *ec2.RevokeSecurityGroupEgressInput) (*request.Request, *ec2.RevokeSecurityGroupEgressOutput)
	RevokeSecurityGroupEgressWithContextFunc                  func(param0 aws.Context, param1 *ec2.RevokeSecurityGroupEgressInput, param2 ...request.Option) (*ec2.RevokeSecurityGroupEgressOutput, error)
	RevokeSecurityGroupIngressFunc                            func(param0 *ec2.RevokeSecurityGroupIngressInput) (*ec2.RevokeSecurityGroupIngressOutput, error)
	RevokeSecurityGroupIngressRequestFunc                     func(param0 *ec2.RevokeSecurityGroupIngressInput) (*request.Request, *ec2.RevokeSecurityGroupIngressOutput)
	RevokeSecurityGroupIngressWithContextFunc                 func(param0 aws.Context, param1 *ec2.RevokeSecurityGroupIngressInput, param2 ...request.Option) (*ec2.RevokeSecurityGroupIngressOutput, error)
	RunInstancesFunc                                          func(param0 *ec2.RunInstancesInput) (*ec2.Reservation, error)
	RunInstancesRequestFunc                                   func(param0 *ec2.RunInstancesInput) (*request.Request, *ec2.Reservation)
	RunInstancesWithContextFunc                               func(param0 aws.Context, param1 *ec2.RunInstancesInput, param2 ...request.Option) (*ec2.Reservation, error)
	RunScheduledInstancesFunc                                 func(param0 *ec2.RunScheduledInstancesInput) (*ec2.RunScheduledInstancesOutput, error)
	RunScheduledInstancesRequestFunc                          func(param0 *ec2.RunScheduledInstancesInput) (*request.Request, *ec2.RunScheduledInstancesOutput)
	RunScheduledInstancesWithContextFunc                      func(param0 aws.Context, param1 *ec2.RunScheduledInstancesInput, param2 ...request.Option) (*ec2.RunScheduledInstancesOutput, error)
	StartInstancesFunc                                        func(param0 *ec2.StartInstancesInput) (*ec2.StartInstancesOutput, error)
	StartInstancesRequestFunc                                 func(param0 *ec2.StartInstancesInput) (*request.Request, *ec2.StartInstancesOutput)
	StartInstancesWithContextFunc                             func(param0 aws.Context, param1 *ec2.StartInstancesInput, param2 ...request.Option) (*ec2.StartInstancesOutput, error)
	StopInstancesFunc                                         func(param0 *ec2.StopInstancesInput) (*ec2.StopInstancesOutput, error)
	StopInstancesRequestFunc                                  func(param0 *ec2.StopInstancesInput) (*request.Request, *ec2.StopInstancesOutput)
	StopInstancesWithContextFunc                              func(param0 aws.Context, param1 *ec2.StopInstancesInput, param2 ...request.Option) (*ec2.StopInstancesOutput, error)
	TerminateInstancesFunc                                    func(param0 *ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error)
	TerminateInstancesRequestFunc                             func(param0 *ec2.TerminateInstancesInput) (*request.Request, *ec2.TerminateInstancesOutput)
	TerminateInstancesWithContextFunc                         func(param0 aws.Context, param1 *ec2.TerminateInstancesInput, param2 ...request.Option) (*ec2.TerminateInstancesOutput, error)
	UnassignIpv6AddressesFunc                                 func(param0 *ec2.UnassignIpv6AddressesInput) (*ec2.UnassignIpv6AddressesOutput, error)
	UnassignIpv6AddressesRequestFunc                          func(param0 *ec2.UnassignIpv6AddressesInput) (*request.Request, *ec2.UnassignIpv6AddressesOutput)
	UnassignIpv6AddressesWithContextFunc                      func(param0 aws.Context, param1 *ec2.UnassignIpv6AddressesInput, param2 ...request.Option) (*ec2.UnassignIpv6AddressesOutput, error)
	UnassignPrivateIpAddressesFunc                            func(param0 *ec2.UnassignPrivateIpAddressesInput) (*ec2.UnassignPrivateIpAddressesOutput, error)
	UnassignPrivateIpAddressesRequestFunc                     func(param0 *ec2.UnassignPrivateIpAddressesInput) (*request.Request, *ec2.UnassignPrivateIpAddressesOutput)
	UnassignPrivateIpAddressesWithContextFunc                 func(param0 aws.Context, param1 *ec2.UnassignPrivateIpAddressesInput, param2 ...request.Option) (*ec2.UnassignPrivateIpAddressesOutput, error)
	UnmonitorInstancesFunc                                    func(param0 *ec2.UnmonitorInstancesInput) (*ec2.UnmonitorInstancesOutput, error)
	UnmonitorInstancesRequestFunc                             func(param0 *ec2.UnmonitorInstancesInput) (*request.Request, *ec2.UnmonitorInstancesOutput)
	UnmonitorInstancesWithContextFunc                         func(param0 aws.Context, param1 *ec2.UnmonitorInstancesInput, param2 ...request.Option) (*ec2.UnmonitorInstancesOutput, error)
	UpdateSecurityGroupRuleDescriptionsEgressFunc             func(param0 *ec2.UpdateSecurityGroupRuleDescriptionsEgressInput) (*ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput, error)
	UpdateSecurityGroupRuleDescriptionsEgressRequestFunc      func(param0 *ec2.UpdateSecurityGroupRuleDescriptionsEgressInput) (*request.Request, *ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput)
	UpdateSecurityGroupRuleDescriptionsEgressWithContextFunc  func(param0 aws.Context, param1 *ec2.UpdateSecurityGroupRuleDescriptionsEgressInput, param2 ...request.Option) (*ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput, error)
	UpdateSecurityGroupRuleDescriptionsIngressFunc            func(param0 *ec2.UpdateSecurityGroupRuleDescriptionsIngressInput) (*ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput, error)
	UpdateSecurityGroupRuleDescriptionsIngressRequestFunc     func(param0 *ec2.UpdateSecurityGroupRuleDescriptionsIngressInput) (*request.Request, *ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput)
	UpdateSecurityGroupRuleDescriptionsIngressWithContextFunc func(param0 aws.Context, param1 *ec2.UpdateSecurityGroupRuleDescriptionsIngressInput, param2 ...request.Option) (*ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput, error)
	WaitUntilBundleTaskCompleteFunc                           func(param0 *ec2.DescribeBundleTasksInput) error
	WaitUntilBundleTaskCompleteWithContextFunc                func(param0 aws.Context, param1 *ec2.DescribeBundleTasksInput, param2 ...request.WaiterOption) error
	WaitUntilConversionTaskCancelledFunc                      func(param0 *ec2.DescribeConversionTasksInput) error
	WaitUntilConversionTaskCancelledWithContextFunc           func(param0 aws.Context, param1 *ec2.DescribeConversionTasksInput, param2 ...request.WaiterOption) error
	WaitUntilConversionTaskCompletedFunc                      func(param0 *ec2.DescribeConversionTasksInput) error
	WaitUntilConversionTaskCompletedWithContextFunc           func(param0 aws.Context, param1 *ec2.DescribeConversionTasksInput, param2 ...request.WaiterOption) error
	WaitUntilConversionTaskDeletedFunc                        func(param0 *ec2.DescribeConversionTasksInput) error
	WaitUntilConversionTaskDeletedWithContextFunc             func(param0 aws.Context, param1 *ec2.DescribeConversionTasksInput, param2 ...request.WaiterOption) error
	WaitUntilCustomerGatewayAvailableFunc                     func(param0 *ec2.DescribeCustomerGatewaysInput) error
	WaitUntilCustomerGatewayAvailableWithContextFunc          func(param0 aws.Context, param1 *ec2.DescribeCustomerGatewaysInput, param2 ...request.WaiterOption) error
	WaitUntilExportTaskCancelledFunc                          func(param0 *ec2.DescribeExportTasksInput) error
	WaitUntilExportTaskCancelledWithContextFunc               func(param0 aws.Context, param1 *ec2.DescribeExportTasksInput, param2 ...request.WaiterOption) error
	WaitUntilExportTaskCompletedFunc                          func(param0 *ec2.DescribeExportTasksInput) error
	WaitUntilExportTaskCompletedWithContextFunc               func(param0 aws.Context, param1 *ec2.DescribeExportTasksInput, param2 ...request.WaiterOption) error
	WaitUntilImageAvailableFunc                               func(param0 *ec2.DescribeImagesInput) error
	WaitUntilImageAvailableWithContextFunc                    func(param0 aws.Context, param1 *ec2.DescribeImagesInput, param2 ...request.WaiterOption) error
	WaitUntilImageExistsFunc                                  func(param0 *ec2.DescribeImagesInput) error
	WaitUntilImageExistsWithContextFunc                       func(param0 aws.Context, param1 *ec2.DescribeImagesInput, param2 ...request.WaiterOption) error
	WaitUntilInstanceExistsFunc                               func(param0 *ec2.DescribeInstancesInput) error
	WaitUntilInstanceExistsWithContextFunc                    func(param0 aws.Context, param1 *ec2.DescribeInstancesInput, param2 ...request.WaiterOption) error
	WaitUntilInstanceRunningFunc                              func(param0 *ec2.DescribeInstancesInput) error
	WaitUntilInstanceRunningWithContextFunc                   func(param0 aws.Context, param1 *ec2.DescribeInstancesInput, param2 ...request.WaiterOption) error
	WaitUntilInstanceStatusOkFunc                             func(param0 *ec2.DescribeInstanceStatusInput) error
	WaitUntilInstanceStatusOkWithContextFunc                  func(param0 aws.Context, param1 *ec2.DescribeInstanceStatusInput, param2 ...request.WaiterOption) error
	WaitUntilInstanceStoppedFunc                              func(param0 *ec2.DescribeInstancesInput) error
	WaitUntilInstanceStoppedWithContextFunc                   func(param0 aws.Context, param1 *ec2.DescribeInstancesInput, param2 ...request.WaiterOption) error
	WaitUntilInstanceTerminatedFunc                           func(param0 *ec2.DescribeInstancesInput) error
	WaitUntilInstanceTerminatedWithContextFunc                func(param0 aws.Context, param1 *ec2.DescribeInstancesInput, param2 ...request.WaiterOption) error
	WaitUntilKeyPairExistsFunc                                func(param0 *ec2.DescribeKeyPairsInput) error
	WaitUntilKeyPairExistsWithContextFunc                     func(param0 aws.Context, param1 *ec2.DescribeKeyPairsInput, param2 ...request.WaiterOption) error
	WaitUntilNatGatewayAvailableFunc                          func(param0 *ec2.DescribeNatGatewaysInput) error
	WaitUntilNatGatewayAvailableWithContextFunc               func(param0 aws.Context, param1 *ec2.DescribeNatGatewaysInput, param2 ...request.WaiterOption) error
	WaitUntilNetworkInterfaceAvailableFunc                    func(param0 *ec2.DescribeNetworkInterfacesInput) error
	WaitUntilNetworkInterfaceAvailableWithContextFunc         func(param0 aws.Context, param1 *ec2.DescribeNetworkInterfacesInput, param2 ...request.WaiterOption) error
	WaitUntilPasswordDataAvailableFunc                        func(param0 *ec2.GetPasswordDataInput) error
	WaitUntilPasswordDataAvailableWithContextFunc             func(param0 aws.Context, param1 *ec2.GetPasswordDataInput, param2 ...request.WaiterOption) error
	WaitUntilSnapshotCompletedFunc                            func(param0 *ec2.DescribeSnapshotsInput) error
	WaitUntilSnapshotCompletedWithContextFunc                 func(param0 aws.Context, param1 *ec2.DescribeSnapshotsInput, param2 ...request.WaiterOption) error
	WaitUntilSpotInstanceRequestFulfilledFunc                 func(param0 *ec2.DescribeSpotInstanceRequestsInput) error
	WaitUntilSpotInstanceRequestFulfilledWithContextFunc      func(param0 aws.Context, param1 *ec2.DescribeSpotInstanceRequestsInput, param2 ...request.WaiterOption) error
	WaitUntilSubnetAvailableFunc                              func(param0 *ec2.DescribeSubnetsInput) error
	WaitUntilSubnetAvailableWithContextFunc                   func(param0 aws.Context, param1 *ec2.DescribeSubnetsInput, param2 ...request.WaiterOption) error
	WaitUntilSystemStatusOkFunc                               func(param0 *ec2.DescribeInstanceStatusInput) error
	WaitUntilSystemStatusOkWithContextFunc                    func(param0 aws.Context, param1 *ec2.DescribeInstanceStatusInput, param2 ...request.WaiterOption) error
	WaitUntilVolumeAvailableFunc                              func(param0 *ec2.DescribeVolumesInput) error
	WaitUntilVolumeAvailableWithContextFunc                   func(param0 aws.Context, param1 *ec2.DescribeVolumesInput, param2 ...request.WaiterOption) error
	WaitUntilVolumeDeletedFunc                                func(param0 *ec2.DescribeVolumesInput) error
	WaitUntilVolumeDeletedWithContextFunc                     func(param0 aws.Context, param1 *ec2.DescribeVolumesInput, param2 ...request.WaiterOption) error
	WaitUntilVolumeInUseFunc                                  func(param0 *ec2.DescribeVolumesInput) error
	WaitUntilVolumeInUseWithContextFunc                       func(param0 aws.Context, param1 *ec2.DescribeVolumesInput, param2 ...request.WaiterOption) error
	WaitUntilVpcAvailableFunc                                 func(param0 *ec2.DescribeVpcsInput) error
	WaitUntilVpcAvailableWithContextFunc                      func(param0 aws.Context, param1 *ec2.DescribeVpcsInput, param2 ...request.WaiterOption) error
	WaitUntilVpcExistsFunc                                    func(param0 *ec2.DescribeVpcsInput) error
	WaitUntilVpcExistsWithContextFunc                         func(param0 aws.Context, param1 *ec2.DescribeVpcsInput, param2 ...request.WaiterOption) error
	WaitUntilVpcPeeringConnectionDeletedFunc                  func(param0 *ec2.DescribeVpcPeeringConnectionsInput) error
	WaitUntilVpcPeeringConnectionDeletedWithContextFunc       func(param0 aws.Context, param1 *ec2.DescribeVpcPeeringConnectionsInput, param2 ...request.WaiterOption) error
	WaitUntilVpcPeeringConnectionExistsFunc                   func(param0 *ec2.DescribeVpcPeeringConnectionsInput) error
	WaitUntilVpcPeeringConnectionExistsWithContextFunc        func(param0 aws.Context, param1 *ec2.DescribeVpcPeeringConnectionsInput, param2 ...request.WaiterOption) error
	WaitUntilVpnConnectionAvailableFunc                       func(param0 *ec2.DescribeVpnConnectionsInput) error
	WaitUntilVpnConnectionAvailableWithContextFunc            func(param0 aws.Context, param1 *ec2.DescribeVpnConnectionsInput, param2 ...request.WaiterOption) error
	WaitUntilVpnConnectionDeletedFunc                         func(param0 *ec2.DescribeVpnConnectionsInput) error
	WaitUntilVpnConnectionDeletedWithContextFunc              func(param0 aws.Context, param1 *ec2.DescribeVpnConnectionsInput, param2 ...request.WaiterOption) error
}

func (m *ec2Mock) AcceptReservedInstancesExchangeQuote(param0 *ec2.AcceptReservedInstancesExchangeQuoteInput) (*ec2.AcceptReservedInstancesExchangeQuoteOutput, error) {
	m.addCall("AcceptReservedInstancesExchangeQuote")
	m.verifyInput("AcceptReservedInstancesExchangeQuote", param0)
	return m.AcceptReservedInstancesExchangeQuoteFunc(param0)
}

func (m *ec2Mock) AcceptReservedInstancesExchangeQuoteRequest(param0 *ec2.AcceptReservedInstancesExchangeQuoteInput) (*request.Request, *ec2.AcceptReservedInstancesExchangeQuoteOutput) {
	m.addCall("AcceptReservedInstancesExchangeQuoteRequest")
	m.verifyInput("AcceptReservedInstancesExchangeQuoteRequest", param0)
	return m.AcceptReservedInstancesExchangeQuoteRequestFunc(param0)
}

func (m *ec2Mock) AcceptReservedInstancesExchangeQuoteWithContext(param0 aws.Context, param1 *ec2.AcceptReservedInstancesExchangeQuoteInput, param2 ...request.Option) (*ec2.AcceptReservedInstancesExchangeQuoteOutput, error) {
	m.addCall("AcceptReservedInstancesExchangeQuoteWithContext")
	m.verifyInput("AcceptReservedInstancesExchangeQuoteWithContext", param0)
	return m.AcceptReservedInstancesExchangeQuoteWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) AcceptVpcPeeringConnection(param0 *ec2.AcceptVpcPeeringConnectionInput) (*ec2.AcceptVpcPeeringConnectionOutput, error) {
	m.addCall("AcceptVpcPeeringConnection")
	m.verifyInput("AcceptVpcPeeringConnection", param0)
	return m.AcceptVpcPeeringConnectionFunc(param0)
}

func (m *ec2Mock) AcceptVpcPeeringConnectionRequest(param0 *ec2.AcceptVpcPeeringConnectionInput) (*request.Request, *ec2.AcceptVpcPeeringConnectionOutput) {
	m.addCall("AcceptVpcPeeringConnectionRequest")
	m.verifyInput("AcceptVpcPeeringConnectionRequest", param0)
	return m.AcceptVpcPeeringConnectionRequestFunc(param0)
}

func (m *ec2Mock) AcceptVpcPeeringConnectionWithContext(param0 aws.Context, param1 *ec2.AcceptVpcPeeringConnectionInput, param2 ...request.Option) (*ec2.AcceptVpcPeeringConnectionOutput, error) {
	m.addCall("AcceptVpcPeeringConnectionWithContext")
	m.verifyInput("AcceptVpcPeeringConnectionWithContext", param0)
	return m.AcceptVpcPeeringConnectionWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) AllocateAddress(param0 *ec2.AllocateAddressInput) (*ec2.AllocateAddressOutput, error) {
	m.addCall("AllocateAddress")
	m.verifyInput("AllocateAddress", param0)
	return m.AllocateAddressFunc(param0)
}

func (m *ec2Mock) AllocateAddressRequest(param0 *ec2.AllocateAddressInput) (*request.Request, *ec2.AllocateAddressOutput) {
	m.addCall("AllocateAddressRequest")
	m.verifyInput("AllocateAddressRequest", param0)
	return m.AllocateAddressRequestFunc(param0)
}

func (m *ec2Mock) AllocateAddressWithContext(param0 aws.Context, param1 *ec2.AllocateAddressInput, param2 ...request.Option) (*ec2.AllocateAddressOutput, error) {
	m.addCall("AllocateAddressWithContext")
	m.verifyInput("AllocateAddressWithContext", param0)
	return m.AllocateAddressWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) AllocateHosts(param0 *ec2.AllocateHostsInput) (*ec2.AllocateHostsOutput, error) {
	m.addCall("AllocateHosts")
	m.verifyInput("AllocateHosts", param0)
	return m.AllocateHostsFunc(param0)
}

func (m *ec2Mock) AllocateHostsRequest(param0 *ec2.AllocateHostsInput) (*request.Request, *ec2.AllocateHostsOutput) {
	m.addCall("AllocateHostsRequest")
	m.verifyInput("AllocateHostsRequest", param0)
	return m.AllocateHostsRequestFunc(param0)
}

func (m *ec2Mock) AllocateHostsWithContext(param0 aws.Context, param1 *ec2.AllocateHostsInput, param2 ...request.Option) (*ec2.AllocateHostsOutput, error) {
	m.addCall("AllocateHostsWithContext")
	m.verifyInput("AllocateHostsWithContext", param0)
	return m.AllocateHostsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) AssignIpv6Addresses(param0 *ec2.AssignIpv6AddressesInput) (*ec2.AssignIpv6AddressesOutput, error) {
	m.addCall("AssignIpv6Addresses")
	m.verifyInput("AssignIpv6Addresses", param0)
	return m.AssignIpv6AddressesFunc(param0)
}

func (m *ec2Mock) AssignIpv6AddressesRequest(param0 *ec2.AssignIpv6AddressesInput) (*request.Request, *ec2.AssignIpv6AddressesOutput) {
	m.addCall("AssignIpv6AddressesRequest")
	m.verifyInput("AssignIpv6AddressesRequest", param0)
	return m.AssignIpv6AddressesRequestFunc(param0)
}

func (m *ec2Mock) AssignIpv6AddressesWithContext(param0 aws.Context, param1 *ec2.AssignIpv6AddressesInput, param2 ...request.Option) (*ec2.AssignIpv6AddressesOutput, error) {
	m.addCall("AssignIpv6AddressesWithContext")
	m.verifyInput("AssignIpv6AddressesWithContext", param0)
	return m.AssignIpv6AddressesWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) AssignPrivateIpAddresses(param0 *ec2.AssignPrivateIpAddressesInput) (*ec2.AssignPrivateIpAddressesOutput, error) {
	m.addCall("AssignPrivateIpAddresses")
	m.verifyInput("AssignPrivateIpAddresses", param0)
	return m.AssignPrivateIpAddressesFunc(param0)
}

func (m *ec2Mock) AssignPrivateIpAddressesRequest(param0 *ec2.AssignPrivateIpAddressesInput) (*request.Request, *ec2.AssignPrivateIpAddressesOutput) {
	m.addCall("AssignPrivateIpAddressesRequest")
	m.verifyInput("AssignPrivateIpAddressesRequest", param0)
	return m.AssignPrivateIpAddressesRequestFunc(param0)
}

func (m *ec2Mock) AssignPrivateIpAddressesWithContext(param0 aws.Context, param1 *ec2.AssignPrivateIpAddressesInput, param2 ...request.Option) (*ec2.AssignPrivateIpAddressesOutput, error) {
	m.addCall("AssignPrivateIpAddressesWithContext")
	m.verifyInput("AssignPrivateIpAddressesWithContext", param0)
	return m.AssignPrivateIpAddressesWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) AssociateAddress(param0 *ec2.AssociateAddressInput) (*ec2.AssociateAddressOutput, error) {
	m.addCall("AssociateAddress")
	m.verifyInput("AssociateAddress", param0)
	return m.AssociateAddressFunc(param0)
}

func (m *ec2Mock) AssociateAddressRequest(param0 *ec2.AssociateAddressInput) (*request.Request, *ec2.AssociateAddressOutput) {
	m.addCall("AssociateAddressRequest")
	m.verifyInput("AssociateAddressRequest", param0)
	return m.AssociateAddressRequestFunc(param0)
}

func (m *ec2Mock) AssociateAddressWithContext(param0 aws.Context, param1 *ec2.AssociateAddressInput, param2 ...request.Option) (*ec2.AssociateAddressOutput, error) {
	m.addCall("AssociateAddressWithContext")
	m.verifyInput("AssociateAddressWithContext", param0)
	return m.AssociateAddressWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) AssociateDhcpOptions(param0 *ec2.AssociateDhcpOptionsInput) (*ec2.AssociateDhcpOptionsOutput, error) {
	m.addCall("AssociateDhcpOptions")
	m.verifyInput("AssociateDhcpOptions", param0)
	return m.AssociateDhcpOptionsFunc(param0)
}

func (m *ec2Mock) AssociateDhcpOptionsRequest(param0 *ec2.AssociateDhcpOptionsInput) (*request.Request, *ec2.AssociateDhcpOptionsOutput) {
	m.addCall("AssociateDhcpOptionsRequest")
	m.verifyInput("AssociateDhcpOptionsRequest", param0)
	return m.AssociateDhcpOptionsRequestFunc(param0)
}

func (m *ec2Mock) AssociateDhcpOptionsWithContext(param0 aws.Context, param1 *ec2.AssociateDhcpOptionsInput, param2 ...request.Option) (*ec2.AssociateDhcpOptionsOutput, error) {
	m.addCall("AssociateDhcpOptionsWithContext")
	m.verifyInput("AssociateDhcpOptionsWithContext", param0)
	return m.AssociateDhcpOptionsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) AssociateIamInstanceProfile(param0 *ec2.AssociateIamInstanceProfileInput) (*ec2.AssociateIamInstanceProfileOutput, error) {
	m.addCall("AssociateIamInstanceProfile")
	m.verifyInput("AssociateIamInstanceProfile", param0)
	return m.AssociateIamInstanceProfileFunc(param0)
}

func (m *ec2Mock) AssociateIamInstanceProfileRequest(param0 *ec2.AssociateIamInstanceProfileInput) (*request.Request, *ec2.AssociateIamInstanceProfileOutput) {
	m.addCall("AssociateIamInstanceProfileRequest")
	m.verifyInput("AssociateIamInstanceProfileRequest", param0)
	return m.AssociateIamInstanceProfileRequestFunc(param0)
}

func (m *ec2Mock) AssociateIamInstanceProfileWithContext(param0 aws.Context, param1 *ec2.AssociateIamInstanceProfileInput, param2 ...request.Option) (*ec2.AssociateIamInstanceProfileOutput, error) {
	m.addCall("AssociateIamInstanceProfileWithContext")
	m.verifyInput("AssociateIamInstanceProfileWithContext", param0)
	return m.AssociateIamInstanceProfileWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) AssociateRouteTable(param0 *ec2.AssociateRouteTableInput) (*ec2.AssociateRouteTableOutput, error) {
	m.addCall("AssociateRouteTable")
	m.verifyInput("AssociateRouteTable", param0)
	return m.AssociateRouteTableFunc(param0)
}

func (m *ec2Mock) AssociateRouteTableRequest(param0 *ec2.AssociateRouteTableInput) (*request.Request, *ec2.AssociateRouteTableOutput) {
	m.addCall("AssociateRouteTableRequest")
	m.verifyInput("AssociateRouteTableRequest", param0)
	return m.AssociateRouteTableRequestFunc(param0)
}

func (m *ec2Mock) AssociateRouteTableWithContext(param0 aws.Context, param1 *ec2.AssociateRouteTableInput, param2 ...request.Option) (*ec2.AssociateRouteTableOutput, error) {
	m.addCall("AssociateRouteTableWithContext")
	m.verifyInput("AssociateRouteTableWithContext", param0)
	return m.AssociateRouteTableWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) AssociateSubnetCidrBlock(param0 *ec2.AssociateSubnetCidrBlockInput) (*ec2.AssociateSubnetCidrBlockOutput, error) {
	m.addCall("AssociateSubnetCidrBlock")
	m.verifyInput("AssociateSubnetCidrBlock", param0)
	return m.AssociateSubnetCidrBlockFunc(param0)
}

func (m *ec2Mock) AssociateSubnetCidrBlockRequest(param0 *ec2.AssociateSubnetCidrBlockInput) (*request.Request, *ec2.AssociateSubnetCidrBlockOutput) {
	m.addCall("AssociateSubnetCidrBlockRequest")
	m.verifyInput("AssociateSubnetCidrBlockRequest", param0)
	return m.AssociateSubnetCidrBlockRequestFunc(param0)
}

func (m *ec2Mock) AssociateSubnetCidrBlockWithContext(param0 aws.Context, param1 *ec2.AssociateSubnetCidrBlockInput, param2 ...request.Option) (*ec2.AssociateSubnetCidrBlockOutput, error) {
	m.addCall("AssociateSubnetCidrBlockWithContext")
	m.verifyInput("AssociateSubnetCidrBlockWithContext", param0)
	return m.AssociateSubnetCidrBlockWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) AssociateVpcCidrBlock(param0 *ec2.AssociateVpcCidrBlockInput) (*ec2.AssociateVpcCidrBlockOutput, error) {
	m.addCall("AssociateVpcCidrBlock")
	m.verifyInput("AssociateVpcCidrBlock", param0)
	return m.AssociateVpcCidrBlockFunc(param0)
}

func (m *ec2Mock) AssociateVpcCidrBlockRequest(param0 *ec2.AssociateVpcCidrBlockInput) (*request.Request, *ec2.AssociateVpcCidrBlockOutput) {
	m.addCall("AssociateVpcCidrBlockRequest")
	m.verifyInput("AssociateVpcCidrBlockRequest", param0)
	return m.AssociateVpcCidrBlockRequestFunc(param0)
}

func (m *ec2Mock) AssociateVpcCidrBlockWithContext(param0 aws.Context, param1 *ec2.AssociateVpcCidrBlockInput, param2 ...request.Option) (*ec2.AssociateVpcCidrBlockOutput, error) {
	m.addCall("AssociateVpcCidrBlockWithContext")
	m.verifyInput("AssociateVpcCidrBlockWithContext", param0)
	return m.AssociateVpcCidrBlockWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) AttachClassicLinkVpc(param0 *ec2.AttachClassicLinkVpcInput) (*ec2.AttachClassicLinkVpcOutput, error) {
	m.addCall("AttachClassicLinkVpc")
	m.verifyInput("AttachClassicLinkVpc", param0)
	return m.AttachClassicLinkVpcFunc(param0)
}

func (m *ec2Mock) AttachClassicLinkVpcRequest(param0 *ec2.AttachClassicLinkVpcInput) (*request.Request, *ec2.AttachClassicLinkVpcOutput) {
	m.addCall("AttachClassicLinkVpcRequest")
	m.verifyInput("AttachClassicLinkVpcRequest", param0)
	return m.AttachClassicLinkVpcRequestFunc(param0)
}

func (m *ec2Mock) AttachClassicLinkVpcWithContext(param0 aws.Context, param1 *ec2.AttachClassicLinkVpcInput, param2 ...request.Option) (*ec2.AttachClassicLinkVpcOutput, error) {
	m.addCall("AttachClassicLinkVpcWithContext")
	m.verifyInput("AttachClassicLinkVpcWithContext", param0)
	return m.AttachClassicLinkVpcWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) AttachInternetGateway(param0 *ec2.AttachInternetGatewayInput) (*ec2.AttachInternetGatewayOutput, error) {
	m.addCall("AttachInternetGateway")
	m.verifyInput("AttachInternetGateway", param0)
	return m.AttachInternetGatewayFunc(param0)
}

func (m *ec2Mock) AttachInternetGatewayRequest(param0 *ec2.AttachInternetGatewayInput) (*request.Request, *ec2.AttachInternetGatewayOutput) {
	m.addCall("AttachInternetGatewayRequest")
	m.verifyInput("AttachInternetGatewayRequest", param0)
	return m.AttachInternetGatewayRequestFunc(param0)
}

func (m *ec2Mock) AttachInternetGatewayWithContext(param0 aws.Context, param1 *ec2.AttachInternetGatewayInput, param2 ...request.Option) (*ec2.AttachInternetGatewayOutput, error) {
	m.addCall("AttachInternetGatewayWithContext")
	m.verifyInput("AttachInternetGatewayWithContext", param0)
	return m.AttachInternetGatewayWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) AttachNetworkInterface(param0 *ec2.AttachNetworkInterfaceInput) (*ec2.AttachNetworkInterfaceOutput, error) {
	m.addCall("AttachNetworkInterface")
	m.verifyInput("AttachNetworkInterface", param0)
	return m.AttachNetworkInterfaceFunc(param0)
}

func (m *ec2Mock) AttachNetworkInterfaceRequest(param0 *ec2.AttachNetworkInterfaceInput) (*request.Request, *ec2.AttachNetworkInterfaceOutput) {
	m.addCall("AttachNetworkInterfaceRequest")
	m.verifyInput("AttachNetworkInterfaceRequest", param0)
	return m.AttachNetworkInterfaceRequestFunc(param0)
}

func (m *ec2Mock) AttachNetworkInterfaceWithContext(param0 aws.Context, param1 *ec2.AttachNetworkInterfaceInput, param2 ...request.Option) (*ec2.AttachNetworkInterfaceOutput, error) {
	m.addCall("AttachNetworkInterfaceWithContext")
	m.verifyInput("AttachNetworkInterfaceWithContext", param0)
	return m.AttachNetworkInterfaceWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) AttachVolume(param0 *ec2.AttachVolumeInput) (*ec2.VolumeAttachment, error) {
	m.addCall("AttachVolume")
	m.verifyInput("AttachVolume", param0)
	return m.AttachVolumeFunc(param0)
}

func (m *ec2Mock) AttachVolumeRequest(param0 *ec2.AttachVolumeInput) (*request.Request, *ec2.VolumeAttachment) {
	m.addCall("AttachVolumeRequest")
	m.verifyInput("AttachVolumeRequest", param0)
	return m.AttachVolumeRequestFunc(param0)
}

func (m *ec2Mock) AttachVolumeWithContext(param0 aws.Context, param1 *ec2.AttachVolumeInput, param2 ...request.Option) (*ec2.VolumeAttachment, error) {
	m.addCall("AttachVolumeWithContext")
	m.verifyInput("AttachVolumeWithContext", param0)
	return m.AttachVolumeWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) AttachVpnGateway(param0 *ec2.AttachVpnGatewayInput) (*ec2.AttachVpnGatewayOutput, error) {
	m.addCall("AttachVpnGateway")
	m.verifyInput("AttachVpnGateway", param0)
	return m.AttachVpnGatewayFunc(param0)
}

func (m *ec2Mock) AttachVpnGatewayRequest(param0 *ec2.AttachVpnGatewayInput) (*request.Request, *ec2.AttachVpnGatewayOutput) {
	m.addCall("AttachVpnGatewayRequest")
	m.verifyInput("AttachVpnGatewayRequest", param0)
	return m.AttachVpnGatewayRequestFunc(param0)
}

func (m *ec2Mock) AttachVpnGatewayWithContext(param0 aws.Context, param1 *ec2.AttachVpnGatewayInput, param2 ...request.Option) (*ec2.AttachVpnGatewayOutput, error) {
	m.addCall("AttachVpnGatewayWithContext")
	m.verifyInput("AttachVpnGatewayWithContext", param0)
	return m.AttachVpnGatewayWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) AuthorizeSecurityGroupEgress(param0 *ec2.AuthorizeSecurityGroupEgressInput) (*ec2.AuthorizeSecurityGroupEgressOutput, error) {
	m.addCall("AuthorizeSecurityGroupEgress")
	m.verifyInput("AuthorizeSecurityGroupEgress", param0)
	return m.AuthorizeSecurityGroupEgressFunc(param0)
}

func (m *ec2Mock) AuthorizeSecurityGroupEgressRequest(param0 *ec2.AuthorizeSecurityGroupEgressInput) (*request.Request, *ec2.AuthorizeSecurityGroupEgressOutput) {
	m.addCall("AuthorizeSecurityGroupEgressRequest")
	m.verifyInput("AuthorizeSecurityGroupEgressRequest", param0)
	return m.AuthorizeSecurityGroupEgressRequestFunc(param0)
}

func (m *ec2Mock) AuthorizeSecurityGroupEgressWithContext(param0 aws.Context, param1 *ec2.AuthorizeSecurityGroupEgressInput, param2 ...request.Option) (*ec2.AuthorizeSecurityGroupEgressOutput, error) {
	m.addCall("AuthorizeSecurityGroupEgressWithContext")
	m.verifyInput("AuthorizeSecurityGroupEgressWithContext", param0)
	return m.AuthorizeSecurityGroupEgressWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) AuthorizeSecurityGroupIngress(param0 *ec2.AuthorizeSecurityGroupIngressInput) (*ec2.AuthorizeSecurityGroupIngressOutput, error) {
	m.addCall("AuthorizeSecurityGroupIngress")
	m.verifyInput("AuthorizeSecurityGroupIngress", param0)
	return m.AuthorizeSecurityGroupIngressFunc(param0)
}

func (m *ec2Mock) AuthorizeSecurityGroupIngressRequest(param0 *ec2.AuthorizeSecurityGroupIngressInput) (*request.Request, *ec2.AuthorizeSecurityGroupIngressOutput) {
	m.addCall("AuthorizeSecurityGroupIngressRequest")
	m.verifyInput("AuthorizeSecurityGroupIngressRequest", param0)
	return m.AuthorizeSecurityGroupIngressRequestFunc(param0)
}

func (m *ec2Mock) AuthorizeSecurityGroupIngressWithContext(param0 aws.Context, param1 *ec2.AuthorizeSecurityGroupIngressInput, param2 ...request.Option) (*ec2.AuthorizeSecurityGroupIngressOutput, error) {
	m.addCall("AuthorizeSecurityGroupIngressWithContext")
	m.verifyInput("AuthorizeSecurityGroupIngressWithContext", param0)
	return m.AuthorizeSecurityGroupIngressWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) BundleInstance(param0 *ec2.BundleInstanceInput) (*ec2.BundleInstanceOutput, error) {
	m.addCall("BundleInstance")
	m.verifyInput("BundleInstance", param0)
	return m.BundleInstanceFunc(param0)
}

func (m *ec2Mock) BundleInstanceRequest(param0 *ec2.BundleInstanceInput) (*request.Request, *ec2.BundleInstanceOutput) {
	m.addCall("BundleInstanceRequest")
	m.verifyInput("BundleInstanceRequest", param0)
	return m.BundleInstanceRequestFunc(param0)
}

func (m *ec2Mock) BundleInstanceWithContext(param0 aws.Context, param1 *ec2.BundleInstanceInput, param2 ...request.Option) (*ec2.BundleInstanceOutput, error) {
	m.addCall("BundleInstanceWithContext")
	m.verifyInput("BundleInstanceWithContext", param0)
	return m.BundleInstanceWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CancelBundleTask(param0 *ec2.CancelBundleTaskInput) (*ec2.CancelBundleTaskOutput, error) {
	m.addCall("CancelBundleTask")
	m.verifyInput("CancelBundleTask", param0)
	return m.CancelBundleTaskFunc(param0)
}

func (m *ec2Mock) CancelBundleTaskRequest(param0 *ec2.CancelBundleTaskInput) (*request.Request, *ec2.CancelBundleTaskOutput) {
	m.addCall("CancelBundleTaskRequest")
	m.verifyInput("CancelBundleTaskRequest", param0)
	return m.CancelBundleTaskRequestFunc(param0)
}

func (m *ec2Mock) CancelBundleTaskWithContext(param0 aws.Context, param1 *ec2.CancelBundleTaskInput, param2 ...request.Option) (*ec2.CancelBundleTaskOutput, error) {
	m.addCall("CancelBundleTaskWithContext")
	m.verifyInput("CancelBundleTaskWithContext", param0)
	return m.CancelBundleTaskWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CancelConversionTask(param0 *ec2.CancelConversionTaskInput) (*ec2.CancelConversionTaskOutput, error) {
	m.addCall("CancelConversionTask")
	m.verifyInput("CancelConversionTask", param0)
	return m.CancelConversionTaskFunc(param0)
}

func (m *ec2Mock) CancelConversionTaskRequest(param0 *ec2.CancelConversionTaskInput) (*request.Request, *ec2.CancelConversionTaskOutput) {
	m.addCall("CancelConversionTaskRequest")
	m.verifyInput("CancelConversionTaskRequest", param0)
	return m.CancelConversionTaskRequestFunc(param0)
}

func (m *ec2Mock) CancelConversionTaskWithContext(param0 aws.Context, param1 *ec2.CancelConversionTaskInput, param2 ...request.Option) (*ec2.CancelConversionTaskOutput, error) {
	m.addCall("CancelConversionTaskWithContext")
	m.verifyInput("CancelConversionTaskWithContext", param0)
	return m.CancelConversionTaskWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CancelExportTask(param0 *ec2.CancelExportTaskInput) (*ec2.CancelExportTaskOutput, error) {
	m.addCall("CancelExportTask")
	m.verifyInput("CancelExportTask", param0)
	return m.CancelExportTaskFunc(param0)
}

func (m *ec2Mock) CancelExportTaskRequest(param0 *ec2.CancelExportTaskInput) (*request.Request, *ec2.CancelExportTaskOutput) {
	m.addCall("CancelExportTaskRequest")
	m.verifyInput("CancelExportTaskRequest", param0)
	return m.CancelExportTaskRequestFunc(param0)
}

func (m *ec2Mock) CancelExportTaskWithContext(param0 aws.Context, param1 *ec2.CancelExportTaskInput, param2 ...request.Option) (*ec2.CancelExportTaskOutput, error) {
	m.addCall("CancelExportTaskWithContext")
	m.verifyInput("CancelExportTaskWithContext", param0)
	return m.CancelExportTaskWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CancelImportTask(param0 *ec2.CancelImportTaskInput) (*ec2.CancelImportTaskOutput, error) {
	m.addCall("CancelImportTask")
	m.verifyInput("CancelImportTask", param0)
	return m.CancelImportTaskFunc(param0)
}

func (m *ec2Mock) CancelImportTaskRequest(param0 *ec2.CancelImportTaskInput) (*request.Request, *ec2.CancelImportTaskOutput) {
	m.addCall("CancelImportTaskRequest")
	m.verifyInput("CancelImportTaskRequest", param0)
	return m.CancelImportTaskRequestFunc(param0)
}

func (m *ec2Mock) CancelImportTaskWithContext(param0 aws.Context, param1 *ec2.CancelImportTaskInput, param2 ...request.Option) (*ec2.CancelImportTaskOutput, error) {
	m.addCall("CancelImportTaskWithContext")
	m.verifyInput("CancelImportTaskWithContext", param0)
	return m.CancelImportTaskWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CancelReservedInstancesListing(param0 *ec2.CancelReservedInstancesListingInput) (*ec2.CancelReservedInstancesListingOutput, error) {
	m.addCall("CancelReservedInstancesListing")
	m.verifyInput("CancelReservedInstancesListing", param0)
	return m.CancelReservedInstancesListingFunc(param0)
}

func (m *ec2Mock) CancelReservedInstancesListingRequest(param0 *ec2.CancelReservedInstancesListingInput) (*request.Request, *ec2.CancelReservedInstancesListingOutput) {
	m.addCall("CancelReservedInstancesListingRequest")
	m.verifyInput("CancelReservedInstancesListingRequest", param0)
	return m.CancelReservedInstancesListingRequestFunc(param0)
}

func (m *ec2Mock) CancelReservedInstancesListingWithContext(param0 aws.Context, param1 *ec2.CancelReservedInstancesListingInput, param2 ...request.Option) (*ec2.CancelReservedInstancesListingOutput, error) {
	m.addCall("CancelReservedInstancesListingWithContext")
	m.verifyInput("CancelReservedInstancesListingWithContext", param0)
	return m.CancelReservedInstancesListingWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CancelSpotFleetRequests(param0 *ec2.CancelSpotFleetRequestsInput) (*ec2.CancelSpotFleetRequestsOutput, error) {
	m.addCall("CancelSpotFleetRequests")
	m.verifyInput("CancelSpotFleetRequests", param0)
	return m.CancelSpotFleetRequestsFunc(param0)
}

func (m *ec2Mock) CancelSpotFleetRequestsRequest(param0 *ec2.CancelSpotFleetRequestsInput) (*request.Request, *ec2.CancelSpotFleetRequestsOutput) {
	m.addCall("CancelSpotFleetRequestsRequest")
	m.verifyInput("CancelSpotFleetRequestsRequest", param0)
	return m.CancelSpotFleetRequestsRequestFunc(param0)
}

func (m *ec2Mock) CancelSpotFleetRequestsWithContext(param0 aws.Context, param1 *ec2.CancelSpotFleetRequestsInput, param2 ...request.Option) (*ec2.CancelSpotFleetRequestsOutput, error) {
	m.addCall("CancelSpotFleetRequestsWithContext")
	m.verifyInput("CancelSpotFleetRequestsWithContext", param0)
	return m.CancelSpotFleetRequestsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CancelSpotInstanceRequests(param0 *ec2.CancelSpotInstanceRequestsInput) (*ec2.CancelSpotInstanceRequestsOutput, error) {
	m.addCall("CancelSpotInstanceRequests")
	m.verifyInput("CancelSpotInstanceRequests", param0)
	return m.CancelSpotInstanceRequestsFunc(param0)
}

func (m *ec2Mock) CancelSpotInstanceRequestsRequest(param0 *ec2.CancelSpotInstanceRequestsInput) (*request.Request, *ec2.CancelSpotInstanceRequestsOutput) {
	m.addCall("CancelSpotInstanceRequestsRequest")
	m.verifyInput("CancelSpotInstanceRequestsRequest", param0)
	return m.CancelSpotInstanceRequestsRequestFunc(param0)
}

func (m *ec2Mock) CancelSpotInstanceRequestsWithContext(param0 aws.Context, param1 *ec2.CancelSpotInstanceRequestsInput, param2 ...request.Option) (*ec2.CancelSpotInstanceRequestsOutput, error) {
	m.addCall("CancelSpotInstanceRequestsWithContext")
	m.verifyInput("CancelSpotInstanceRequestsWithContext", param0)
	return m.CancelSpotInstanceRequestsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ConfirmProductInstance(param0 *ec2.ConfirmProductInstanceInput) (*ec2.ConfirmProductInstanceOutput, error) {
	m.addCall("ConfirmProductInstance")
	m.verifyInput("ConfirmProductInstance", param0)
	return m.ConfirmProductInstanceFunc(param0)
}

func (m *ec2Mock) ConfirmProductInstanceRequest(param0 *ec2.ConfirmProductInstanceInput) (*request.Request, *ec2.ConfirmProductInstanceOutput) {
	m.addCall("ConfirmProductInstanceRequest")
	m.verifyInput("ConfirmProductInstanceRequest", param0)
	return m.ConfirmProductInstanceRequestFunc(param0)
}

func (m *ec2Mock) ConfirmProductInstanceWithContext(param0 aws.Context, param1 *ec2.ConfirmProductInstanceInput, param2 ...request.Option) (*ec2.ConfirmProductInstanceOutput, error) {
	m.addCall("ConfirmProductInstanceWithContext")
	m.verifyInput("ConfirmProductInstanceWithContext", param0)
	return m.ConfirmProductInstanceWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CopyFpgaImage(param0 *ec2.CopyFpgaImageInput) (*ec2.CopyFpgaImageOutput, error) {
	m.addCall("CopyFpgaImage")
	m.verifyInput("CopyFpgaImage", param0)
	return m.CopyFpgaImageFunc(param0)
}

func (m *ec2Mock) CopyFpgaImageRequest(param0 *ec2.CopyFpgaImageInput) (*request.Request, *ec2.CopyFpgaImageOutput) {
	m.addCall("CopyFpgaImageRequest")
	m.verifyInput("CopyFpgaImageRequest", param0)
	return m.CopyFpgaImageRequestFunc(param0)
}

func (m *ec2Mock) CopyFpgaImageWithContext(param0 aws.Context, param1 *ec2.CopyFpgaImageInput, param2 ...request.Option) (*ec2.CopyFpgaImageOutput, error) {
	m.addCall("CopyFpgaImageWithContext")
	m.verifyInput("CopyFpgaImageWithContext", param0)
	return m.CopyFpgaImageWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CopyImage(param0 *ec2.CopyImageInput) (*ec2.CopyImageOutput, error) {
	m.addCall("CopyImage")
	m.verifyInput("CopyImage", param0)
	return m.CopyImageFunc(param0)
}

func (m *ec2Mock) CopyImageRequest(param0 *ec2.CopyImageInput) (*request.Request, *ec2.CopyImageOutput) {
	m.addCall("CopyImageRequest")
	m.verifyInput("CopyImageRequest", param0)
	return m.CopyImageRequestFunc(param0)
}

func (m *ec2Mock) CopyImageWithContext(param0 aws.Context, param1 *ec2.CopyImageInput, param2 ...request.Option) (*ec2.CopyImageOutput, error) {
	m.addCall("CopyImageWithContext")
	m.verifyInput("CopyImageWithContext", param0)
	return m.CopyImageWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CopySnapshot(param0 *ec2.CopySnapshotInput) (*ec2.CopySnapshotOutput, error) {
	m.addCall("CopySnapshot")
	m.verifyInput("CopySnapshot", param0)
	return m.CopySnapshotFunc(param0)
}

func (m *ec2Mock) CopySnapshotRequest(param0 *ec2.CopySnapshotInput) (*request.Request, *ec2.CopySnapshotOutput) {
	m.addCall("CopySnapshotRequest")
	m.verifyInput("CopySnapshotRequest", param0)
	return m.CopySnapshotRequestFunc(param0)
}

func (m *ec2Mock) CopySnapshotWithContext(param0 aws.Context, param1 *ec2.CopySnapshotInput, param2 ...request.Option) (*ec2.CopySnapshotOutput, error) {
	m.addCall("CopySnapshotWithContext")
	m.verifyInput("CopySnapshotWithContext", param0)
	return m.CopySnapshotWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CreateCustomerGateway(param0 *ec2.CreateCustomerGatewayInput) (*ec2.CreateCustomerGatewayOutput, error) {
	m.addCall("CreateCustomerGateway")
	m.verifyInput("CreateCustomerGateway", param0)
	return m.CreateCustomerGatewayFunc(param0)
}

func (m *ec2Mock) CreateCustomerGatewayRequest(param0 *ec2.CreateCustomerGatewayInput) (*request.Request, *ec2.CreateCustomerGatewayOutput) {
	m.addCall("CreateCustomerGatewayRequest")
	m.verifyInput("CreateCustomerGatewayRequest", param0)
	return m.CreateCustomerGatewayRequestFunc(param0)
}

func (m *ec2Mock) CreateCustomerGatewayWithContext(param0 aws.Context, param1 *ec2.CreateCustomerGatewayInput, param2 ...request.Option) (*ec2.CreateCustomerGatewayOutput, error) {
	m.addCall("CreateCustomerGatewayWithContext")
	m.verifyInput("CreateCustomerGatewayWithContext", param0)
	return m.CreateCustomerGatewayWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CreateDefaultVpc(param0 *ec2.CreateDefaultVpcInput) (*ec2.CreateDefaultVpcOutput, error) {
	m.addCall("CreateDefaultVpc")
	m.verifyInput("CreateDefaultVpc", param0)
	return m.CreateDefaultVpcFunc(param0)
}

func (m *ec2Mock) CreateDefaultVpcRequest(param0 *ec2.CreateDefaultVpcInput) (*request.Request, *ec2.CreateDefaultVpcOutput) {
	m.addCall("CreateDefaultVpcRequest")
	m.verifyInput("CreateDefaultVpcRequest", param0)
	return m.CreateDefaultVpcRequestFunc(param0)
}

func (m *ec2Mock) CreateDefaultVpcWithContext(param0 aws.Context, param1 *ec2.CreateDefaultVpcInput, param2 ...request.Option) (*ec2.CreateDefaultVpcOutput, error) {
	m.addCall("CreateDefaultVpcWithContext")
	m.verifyInput("CreateDefaultVpcWithContext", param0)
	return m.CreateDefaultVpcWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CreateDhcpOptions(param0 *ec2.CreateDhcpOptionsInput) (*ec2.CreateDhcpOptionsOutput, error) {
	m.addCall("CreateDhcpOptions")
	m.verifyInput("CreateDhcpOptions", param0)
	return m.CreateDhcpOptionsFunc(param0)
}

func (m *ec2Mock) CreateDhcpOptionsRequest(param0 *ec2.CreateDhcpOptionsInput) (*request.Request, *ec2.CreateDhcpOptionsOutput) {
	m.addCall("CreateDhcpOptionsRequest")
	m.verifyInput("CreateDhcpOptionsRequest", param0)
	return m.CreateDhcpOptionsRequestFunc(param0)
}

func (m *ec2Mock) CreateDhcpOptionsWithContext(param0 aws.Context, param1 *ec2.CreateDhcpOptionsInput, param2 ...request.Option) (*ec2.CreateDhcpOptionsOutput, error) {
	m.addCall("CreateDhcpOptionsWithContext")
	m.verifyInput("CreateDhcpOptionsWithContext", param0)
	return m.CreateDhcpOptionsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CreateEgressOnlyInternetGateway(param0 *ec2.CreateEgressOnlyInternetGatewayInput) (*ec2.CreateEgressOnlyInternetGatewayOutput, error) {
	m.addCall("CreateEgressOnlyInternetGateway")
	m.verifyInput("CreateEgressOnlyInternetGateway", param0)
	return m.CreateEgressOnlyInternetGatewayFunc(param0)
}

func (m *ec2Mock) CreateEgressOnlyInternetGatewayRequest(param0 *ec2.CreateEgressOnlyInternetGatewayInput) (*request.Request, *ec2.CreateEgressOnlyInternetGatewayOutput) {
	m.addCall("CreateEgressOnlyInternetGatewayRequest")
	m.verifyInput("CreateEgressOnlyInternetGatewayRequest", param0)
	return m.CreateEgressOnlyInternetGatewayRequestFunc(param0)
}

func (m *ec2Mock) CreateEgressOnlyInternetGatewayWithContext(param0 aws.Context, param1 *ec2.CreateEgressOnlyInternetGatewayInput, param2 ...request.Option) (*ec2.CreateEgressOnlyInternetGatewayOutput, error) {
	m.addCall("CreateEgressOnlyInternetGatewayWithContext")
	m.verifyInput("CreateEgressOnlyInternetGatewayWithContext", param0)
	return m.CreateEgressOnlyInternetGatewayWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CreateFlowLogs(param0 *ec2.CreateFlowLogsInput) (*ec2.CreateFlowLogsOutput, error) {
	m.addCall("CreateFlowLogs")
	m.verifyInput("CreateFlowLogs", param0)
	return m.CreateFlowLogsFunc(param0)
}

func (m *ec2Mock) CreateFlowLogsRequest(param0 *ec2.CreateFlowLogsInput) (*request.Request, *ec2.CreateFlowLogsOutput) {
	m.addCall("CreateFlowLogsRequest")
	m.verifyInput("CreateFlowLogsRequest", param0)
	return m.CreateFlowLogsRequestFunc(param0)
}

func (m *ec2Mock) CreateFlowLogsWithContext(param0 aws.Context, param1 *ec2.CreateFlowLogsInput, param2 ...request.Option) (*ec2.CreateFlowLogsOutput, error) {
	m.addCall("CreateFlowLogsWithContext")
	m.verifyInput("CreateFlowLogsWithContext", param0)
	return m.CreateFlowLogsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CreateFpgaImage(param0 *ec2.CreateFpgaImageInput) (*ec2.CreateFpgaImageOutput, error) {
	m.addCall("CreateFpgaImage")
	m.verifyInput("CreateFpgaImage", param0)
	return m.CreateFpgaImageFunc(param0)
}

func (m *ec2Mock) CreateFpgaImageRequest(param0 *ec2.CreateFpgaImageInput) (*request.Request, *ec2.CreateFpgaImageOutput) {
	m.addCall("CreateFpgaImageRequest")
	m.verifyInput("CreateFpgaImageRequest", param0)
	return m.CreateFpgaImageRequestFunc(param0)
}

func (m *ec2Mock) CreateFpgaImageWithContext(param0 aws.Context, param1 *ec2.CreateFpgaImageInput, param2 ...request.Option) (*ec2.CreateFpgaImageOutput, error) {
	m.addCall("CreateFpgaImageWithContext")
	m.verifyInput("CreateFpgaImageWithContext", param0)
	return m.CreateFpgaImageWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CreateImage(param0 *ec2.CreateImageInput) (*ec2.CreateImageOutput, error) {
	m.addCall("CreateImage")
	m.verifyInput("CreateImage", param0)
	return m.CreateImageFunc(param0)
}

func (m *ec2Mock) CreateImageRequest(param0 *ec2.CreateImageInput) (*request.Request, *ec2.CreateImageOutput) {
	m.addCall("CreateImageRequest")
	m.verifyInput("CreateImageRequest", param0)
	return m.CreateImageRequestFunc(param0)
}

func (m *ec2Mock) CreateImageWithContext(param0 aws.Context, param1 *ec2.CreateImageInput, param2 ...request.Option) (*ec2.CreateImageOutput, error) {
	m.addCall("CreateImageWithContext")
	m.verifyInput("CreateImageWithContext", param0)
	return m.CreateImageWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CreateInstanceExportTask(param0 *ec2.CreateInstanceExportTaskInput) (*ec2.CreateInstanceExportTaskOutput, error) {
	m.addCall("CreateInstanceExportTask")
	m.verifyInput("CreateInstanceExportTask", param0)
	return m.CreateInstanceExportTaskFunc(param0)
}

func (m *ec2Mock) CreateInstanceExportTaskRequest(param0 *ec2.CreateInstanceExportTaskInput) (*request.Request, *ec2.CreateInstanceExportTaskOutput) {
	m.addCall("CreateInstanceExportTaskRequest")
	m.verifyInput("CreateInstanceExportTaskRequest", param0)
	return m.CreateInstanceExportTaskRequestFunc(param0)
}

func (m *ec2Mock) CreateInstanceExportTaskWithContext(param0 aws.Context, param1 *ec2.CreateInstanceExportTaskInput, param2 ...request.Option) (*ec2.CreateInstanceExportTaskOutput, error) {
	m.addCall("CreateInstanceExportTaskWithContext")
	m.verifyInput("CreateInstanceExportTaskWithContext", param0)
	return m.CreateInstanceExportTaskWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CreateInternetGateway(param0 *ec2.CreateInternetGatewayInput) (*ec2.CreateInternetGatewayOutput, error) {
	m.addCall("CreateInternetGateway")
	m.verifyInput("CreateInternetGateway", param0)
	return m.CreateInternetGatewayFunc(param0)
}

func (m *ec2Mock) CreateInternetGatewayRequest(param0 *ec2.CreateInternetGatewayInput) (*request.Request, *ec2.CreateInternetGatewayOutput) {
	m.addCall("CreateInternetGatewayRequest")
	m.verifyInput("CreateInternetGatewayRequest", param0)
	return m.CreateInternetGatewayRequestFunc(param0)
}

func (m *ec2Mock) CreateInternetGatewayWithContext(param0 aws.Context, param1 *ec2.CreateInternetGatewayInput, param2 ...request.Option) (*ec2.CreateInternetGatewayOutput, error) {
	m.addCall("CreateInternetGatewayWithContext")
	m.verifyInput("CreateInternetGatewayWithContext", param0)
	return m.CreateInternetGatewayWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CreateKeyPair(param0 *ec2.CreateKeyPairInput) (*ec2.CreateKeyPairOutput, error) {
	m.addCall("CreateKeyPair")
	m.verifyInput("CreateKeyPair", param0)
	return m.CreateKeyPairFunc(param0)
}

func (m *ec2Mock) CreateKeyPairRequest(param0 *ec2.CreateKeyPairInput) (*request.Request, *ec2.CreateKeyPairOutput) {
	m.addCall("CreateKeyPairRequest")
	m.verifyInput("CreateKeyPairRequest", param0)
	return m.CreateKeyPairRequestFunc(param0)
}

func (m *ec2Mock) CreateKeyPairWithContext(param0 aws.Context, param1 *ec2.CreateKeyPairInput, param2 ...request.Option) (*ec2.CreateKeyPairOutput, error) {
	m.addCall("CreateKeyPairWithContext")
	m.verifyInput("CreateKeyPairWithContext", param0)
	return m.CreateKeyPairWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CreateNatGateway(param0 *ec2.CreateNatGatewayInput) (*ec2.CreateNatGatewayOutput, error) {
	m.addCall("CreateNatGateway")
	m.verifyInput("CreateNatGateway", param0)
	return m.CreateNatGatewayFunc(param0)
}

func (m *ec2Mock) CreateNatGatewayRequest(param0 *ec2.CreateNatGatewayInput) (*request.Request, *ec2.CreateNatGatewayOutput) {
	m.addCall("CreateNatGatewayRequest")
	m.verifyInput("CreateNatGatewayRequest", param0)
	return m.CreateNatGatewayRequestFunc(param0)
}

func (m *ec2Mock) CreateNatGatewayWithContext(param0 aws.Context, param1 *ec2.CreateNatGatewayInput, param2 ...request.Option) (*ec2.CreateNatGatewayOutput, error) {
	m.addCall("CreateNatGatewayWithContext")
	m.verifyInput("CreateNatGatewayWithContext", param0)
	return m.CreateNatGatewayWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CreateNetworkAcl(param0 *ec2.CreateNetworkAclInput) (*ec2.CreateNetworkAclOutput, error) {
	m.addCall("CreateNetworkAcl")
	m.verifyInput("CreateNetworkAcl", param0)
	return m.CreateNetworkAclFunc(param0)
}

func (m *ec2Mock) CreateNetworkAclEntry(param0 *ec2.CreateNetworkAclEntryInput) (*ec2.CreateNetworkAclEntryOutput, error) {
	m.addCall("CreateNetworkAclEntry")
	m.verifyInput("CreateNetworkAclEntry", param0)
	return m.CreateNetworkAclEntryFunc(param0)
}

func (m *ec2Mock) CreateNetworkAclEntryRequest(param0 *ec2.CreateNetworkAclEntryInput) (*request.Request, *ec2.CreateNetworkAclEntryOutput) {
	m.addCall("CreateNetworkAclEntryRequest")
	m.verifyInput("CreateNetworkAclEntryRequest", param0)
	return m.CreateNetworkAclEntryRequestFunc(param0)
}

func (m *ec2Mock) CreateNetworkAclEntryWithContext(param0 aws.Context, param1 *ec2.CreateNetworkAclEntryInput, param2 ...request.Option) (*ec2.CreateNetworkAclEntryOutput, error) {
	m.addCall("CreateNetworkAclEntryWithContext")
	m.verifyInput("CreateNetworkAclEntryWithContext", param0)
	return m.CreateNetworkAclEntryWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CreateNetworkAclRequest(param0 *ec2.CreateNetworkAclInput) (*request.Request, *ec2.CreateNetworkAclOutput) {
	m.addCall("CreateNetworkAclRequest")
	m.verifyInput("CreateNetworkAclRequest", param0)
	return m.CreateNetworkAclRequestFunc(param0)
}

func (m *ec2Mock) CreateNetworkAclWithContext(param0 aws.Context, param1 *ec2.CreateNetworkAclInput, param2 ...request.Option) (*ec2.CreateNetworkAclOutput, error) {
	m.addCall("CreateNetworkAclWithContext")
	m.verifyInput("CreateNetworkAclWithContext", param0)
	return m.CreateNetworkAclWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CreateNetworkInterface(param0 *ec2.CreateNetworkInterfaceInput) (*ec2.CreateNetworkInterfaceOutput, error) {
	m.addCall("CreateNetworkInterface")
	m.verifyInput("CreateNetworkInterface", param0)
	return m.CreateNetworkInterfaceFunc(param0)
}

func (m *ec2Mock) CreateNetworkInterfacePermission(param0 *ec2.CreateNetworkInterfacePermissionInput) (*ec2.CreateNetworkInterfacePermissionOutput, error) {
	m.addCall("CreateNetworkInterfacePermission")
	m.verifyInput("CreateNetworkInterfacePermission", param0)
	return m.CreateNetworkInterfacePermissionFunc(param0)
}

func (m *ec2Mock) CreateNetworkInterfacePermissionRequest(param0 *ec2.CreateNetworkInterfacePermissionInput) (*request.Request, *ec2.CreateNetworkInterfacePermissionOutput) {
	m.addCall("CreateNetworkInterfacePermissionRequest")
	m.verifyInput("CreateNetworkInterfacePermissionRequest", param0)
	return m.CreateNetworkInterfacePermissionRequestFunc(param0)
}

func (m *ec2Mock) CreateNetworkInterfacePermissionWithContext(param0 aws.Context, param1 *ec2.CreateNetworkInterfacePermissionInput, param2 ...request.Option) (*ec2.CreateNetworkInterfacePermissionOutput, error) {
	m.addCall("CreateNetworkInterfacePermissionWithContext")
	m.verifyInput("CreateNetworkInterfacePermissionWithContext", param0)
	return m.CreateNetworkInterfacePermissionWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CreateNetworkInterfaceRequest(param0 *ec2.CreateNetworkInterfaceInput) (*request.Request, *ec2.CreateNetworkInterfaceOutput) {
	m.addCall("CreateNetworkInterfaceRequest")
	m.verifyInput("CreateNetworkInterfaceRequest", param0)
	return m.CreateNetworkInterfaceRequestFunc(param0)
}

func (m *ec2Mock) CreateNetworkInterfaceWithContext(param0 aws.Context, param1 *ec2.CreateNetworkInterfaceInput, param2 ...request.Option) (*ec2.CreateNetworkInterfaceOutput, error) {
	m.addCall("CreateNetworkInterfaceWithContext")
	m.verifyInput("CreateNetworkInterfaceWithContext", param0)
	return m.CreateNetworkInterfaceWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CreatePlacementGroup(param0 *ec2.CreatePlacementGroupInput) (*ec2.CreatePlacementGroupOutput, error) {
	m.addCall("CreatePlacementGroup")
	m.verifyInput("CreatePlacementGroup", param0)
	return m.CreatePlacementGroupFunc(param0)
}

func (m *ec2Mock) CreatePlacementGroupRequest(param0 *ec2.CreatePlacementGroupInput) (*request.Request, *ec2.CreatePlacementGroupOutput) {
	m.addCall("CreatePlacementGroupRequest")
	m.verifyInput("CreatePlacementGroupRequest", param0)
	return m.CreatePlacementGroupRequestFunc(param0)
}

func (m *ec2Mock) CreatePlacementGroupWithContext(param0 aws.Context, param1 *ec2.CreatePlacementGroupInput, param2 ...request.Option) (*ec2.CreatePlacementGroupOutput, error) {
	m.addCall("CreatePlacementGroupWithContext")
	m.verifyInput("CreatePlacementGroupWithContext", param0)
	return m.CreatePlacementGroupWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CreateReservedInstancesListing(param0 *ec2.CreateReservedInstancesListingInput) (*ec2.CreateReservedInstancesListingOutput, error) {
	m.addCall("CreateReservedInstancesListing")
	m.verifyInput("CreateReservedInstancesListing", param0)
	return m.CreateReservedInstancesListingFunc(param0)
}

func (m *ec2Mock) CreateReservedInstancesListingRequest(param0 *ec2.CreateReservedInstancesListingInput) (*request.Request, *ec2.CreateReservedInstancesListingOutput) {
	m.addCall("CreateReservedInstancesListingRequest")
	m.verifyInput("CreateReservedInstancesListingRequest", param0)
	return m.CreateReservedInstancesListingRequestFunc(param0)
}

func (m *ec2Mock) CreateReservedInstancesListingWithContext(param0 aws.Context, param1 *ec2.CreateReservedInstancesListingInput, param2 ...request.Option) (*ec2.CreateReservedInstancesListingOutput, error) {
	m.addCall("CreateReservedInstancesListingWithContext")
	m.verifyInput("CreateReservedInstancesListingWithContext", param0)
	return m.CreateReservedInstancesListingWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CreateRoute(param0 *ec2.CreateRouteInput) (*ec2.CreateRouteOutput, error) {
	m.addCall("CreateRoute")
	m.verifyInput("CreateRoute", param0)
	return m.CreateRouteFunc(param0)
}

func (m *ec2Mock) CreateRouteRequest(param0 *ec2.CreateRouteInput) (*request.Request, *ec2.CreateRouteOutput) {
	m.addCall("CreateRouteRequest")
	m.verifyInput("CreateRouteRequest", param0)
	return m.CreateRouteRequestFunc(param0)
}

func (m *ec2Mock) CreateRouteTable(param0 *ec2.CreateRouteTableInput) (*ec2.CreateRouteTableOutput, error) {
	m.addCall("CreateRouteTable")
	m.verifyInput("CreateRouteTable", param0)
	return m.CreateRouteTableFunc(param0)
}

func (m *ec2Mock) CreateRouteTableRequest(param0 *ec2.CreateRouteTableInput) (*request.Request, *ec2.CreateRouteTableOutput) {
	m.addCall("CreateRouteTableRequest")
	m.verifyInput("CreateRouteTableRequest", param0)
	return m.CreateRouteTableRequestFunc(param0)
}

func (m *ec2Mock) CreateRouteTableWithContext(param0 aws.Context, param1 *ec2.CreateRouteTableInput, param2 ...request.Option) (*ec2.CreateRouteTableOutput, error) {
	m.addCall("CreateRouteTableWithContext")
	m.verifyInput("CreateRouteTableWithContext", param0)
	return m.CreateRouteTableWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CreateRouteWithContext(param0 aws.Context, param1 *ec2.CreateRouteInput, param2 ...request.Option) (*ec2.CreateRouteOutput, error) {
	m.addCall("CreateRouteWithContext")
	m.verifyInput("CreateRouteWithContext", param0)
	return m.CreateRouteWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CreateSecurityGroup(param0 *ec2.CreateSecurityGroupInput) (*ec2.CreateSecurityGroupOutput, error) {
	m.addCall("CreateSecurityGroup")
	m.verifyInput("CreateSecurityGroup", param0)
	return m.CreateSecurityGroupFunc(param0)
}

func (m *ec2Mock) CreateSecurityGroupRequest(param0 *ec2.CreateSecurityGroupInput) (*request.Request, *ec2.CreateSecurityGroupOutput) {
	m.addCall("CreateSecurityGroupRequest")
	m.verifyInput("CreateSecurityGroupRequest", param0)
	return m.CreateSecurityGroupRequestFunc(param0)
}

func (m *ec2Mock) CreateSecurityGroupWithContext(param0 aws.Context, param1 *ec2.CreateSecurityGroupInput, param2 ...request.Option) (*ec2.CreateSecurityGroupOutput, error) {
	m.addCall("CreateSecurityGroupWithContext")
	m.verifyInput("CreateSecurityGroupWithContext", param0)
	return m.CreateSecurityGroupWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CreateSnapshot(param0 *ec2.CreateSnapshotInput) (*ec2.Snapshot, error) {
	m.addCall("CreateSnapshot")
	m.verifyInput("CreateSnapshot", param0)
	return m.CreateSnapshotFunc(param0)
}

func (m *ec2Mock) CreateSnapshotRequest(param0 *ec2.CreateSnapshotInput) (*request.Request, *ec2.Snapshot) {
	m.addCall("CreateSnapshotRequest")
	m.verifyInput("CreateSnapshotRequest", param0)
	return m.CreateSnapshotRequestFunc(param0)
}

func (m *ec2Mock) CreateSnapshotWithContext(param0 aws.Context, param1 *ec2.CreateSnapshotInput, param2 ...request.Option) (*ec2.Snapshot, error) {
	m.addCall("CreateSnapshotWithContext")
	m.verifyInput("CreateSnapshotWithContext", param0)
	return m.CreateSnapshotWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CreateSpotDatafeedSubscription(param0 *ec2.CreateSpotDatafeedSubscriptionInput) (*ec2.CreateSpotDatafeedSubscriptionOutput, error) {
	m.addCall("CreateSpotDatafeedSubscription")
	m.verifyInput("CreateSpotDatafeedSubscription", param0)
	return m.CreateSpotDatafeedSubscriptionFunc(param0)
}

func (m *ec2Mock) CreateSpotDatafeedSubscriptionRequest(param0 *ec2.CreateSpotDatafeedSubscriptionInput) (*request.Request, *ec2.CreateSpotDatafeedSubscriptionOutput) {
	m.addCall("CreateSpotDatafeedSubscriptionRequest")
	m.verifyInput("CreateSpotDatafeedSubscriptionRequest", param0)
	return m.CreateSpotDatafeedSubscriptionRequestFunc(param0)
}

func (m *ec2Mock) CreateSpotDatafeedSubscriptionWithContext(param0 aws.Context, param1 *ec2.CreateSpotDatafeedSubscriptionInput, param2 ...request.Option) (*ec2.CreateSpotDatafeedSubscriptionOutput, error) {
	m.addCall("CreateSpotDatafeedSubscriptionWithContext")
	m.verifyInput("CreateSpotDatafeedSubscriptionWithContext", param0)
	return m.CreateSpotDatafeedSubscriptionWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CreateSubnet(param0 *ec2.CreateSubnetInput) (*ec2.CreateSubnetOutput, error) {
	m.addCall("CreateSubnet")
	m.verifyInput("CreateSubnet", param0)
	return m.CreateSubnetFunc(param0)
}

func (m *ec2Mock) CreateSubnetRequest(param0 *ec2.CreateSubnetInput) (*request.Request, *ec2.CreateSubnetOutput) {
	m.addCall("CreateSubnetRequest")
	m.verifyInput("CreateSubnetRequest", param0)
	return m.CreateSubnetRequestFunc(param0)
}

func (m *ec2Mock) CreateSubnetWithContext(param0 aws.Context, param1 *ec2.CreateSubnetInput, param2 ...request.Option) (*ec2.CreateSubnetOutput, error) {
	m.addCall("CreateSubnetWithContext")
	m.verifyInput("CreateSubnetWithContext", param0)
	return m.CreateSubnetWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CreateTags(param0 *ec2.CreateTagsInput) (*ec2.CreateTagsOutput, error) {
	m.addCall("CreateTags")
	m.verifyInput("CreateTags", param0)
	return m.CreateTagsFunc(param0)
}

func (m *ec2Mock) CreateTagsRequest(param0 *ec2.CreateTagsInput) (*request.Request, *ec2.CreateTagsOutput) {
	m.addCall("CreateTagsRequest")
	m.verifyInput("CreateTagsRequest", param0)
	return m.CreateTagsRequestFunc(param0)
}

func (m *ec2Mock) CreateTagsWithContext(param0 aws.Context, param1 *ec2.CreateTagsInput, param2 ...request.Option) (*ec2.CreateTagsOutput, error) {
	m.addCall("CreateTagsWithContext")
	m.verifyInput("CreateTagsWithContext", param0)
	return m.CreateTagsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CreateVolume(param0 *ec2.CreateVolumeInput) (*ec2.Volume, error) {
	m.addCall("CreateVolume")
	m.verifyInput("CreateVolume", param0)
	return m.CreateVolumeFunc(param0)
}

func (m *ec2Mock) CreateVolumeRequest(param0 *ec2.CreateVolumeInput) (*request.Request, *ec2.Volume) {
	m.addCall("CreateVolumeRequest")
	m.verifyInput("CreateVolumeRequest", param0)
	return m.CreateVolumeRequestFunc(param0)
}

func (m *ec2Mock) CreateVolumeWithContext(param0 aws.Context, param1 *ec2.CreateVolumeInput, param2 ...request.Option) (*ec2.Volume, error) {
	m.addCall("CreateVolumeWithContext")
	m.verifyInput("CreateVolumeWithContext", param0)
	return m.CreateVolumeWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CreateVpc(param0 *ec2.CreateVpcInput) (*ec2.CreateVpcOutput, error) {
	m.addCall("CreateVpc")
	m.verifyInput("CreateVpc", param0)
	return m.CreateVpcFunc(param0)
}

func (m *ec2Mock) CreateVpcEndpoint(param0 *ec2.CreateVpcEndpointInput) (*ec2.CreateVpcEndpointOutput, error) {
	m.addCall("CreateVpcEndpoint")
	m.verifyInput("CreateVpcEndpoint", param0)
	return m.CreateVpcEndpointFunc(param0)
}

func (m *ec2Mock) CreateVpcEndpointRequest(param0 *ec2.CreateVpcEndpointInput) (*request.Request, *ec2.CreateVpcEndpointOutput) {
	m.addCall("CreateVpcEndpointRequest")
	m.verifyInput("CreateVpcEndpointRequest", param0)
	return m.CreateVpcEndpointRequestFunc(param0)
}

func (m *ec2Mock) CreateVpcEndpointWithContext(param0 aws.Context, param1 *ec2.CreateVpcEndpointInput, param2 ...request.Option) (*ec2.CreateVpcEndpointOutput, error) {
	m.addCall("CreateVpcEndpointWithContext")
	m.verifyInput("CreateVpcEndpointWithContext", param0)
	return m.CreateVpcEndpointWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CreateVpcPeeringConnection(param0 *ec2.CreateVpcPeeringConnectionInput) (*ec2.CreateVpcPeeringConnectionOutput, error) {
	m.addCall("CreateVpcPeeringConnection")
	m.verifyInput("CreateVpcPeeringConnection", param0)
	return m.CreateVpcPeeringConnectionFunc(param0)
}

func (m *ec2Mock) CreateVpcPeeringConnectionRequest(param0 *ec2.CreateVpcPeeringConnectionInput) (*request.Request, *ec2.CreateVpcPeeringConnectionOutput) {
	m.addCall("CreateVpcPeeringConnectionRequest")
	m.verifyInput("CreateVpcPeeringConnectionRequest", param0)
	return m.CreateVpcPeeringConnectionRequestFunc(param0)
}

func (m *ec2Mock) CreateVpcPeeringConnectionWithContext(param0 aws.Context, param1 *ec2.CreateVpcPeeringConnectionInput, param2 ...request.Option) (*ec2.CreateVpcPeeringConnectionOutput, error) {
	m.addCall("CreateVpcPeeringConnectionWithContext")
	m.verifyInput("CreateVpcPeeringConnectionWithContext", param0)
	return m.CreateVpcPeeringConnectionWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CreateVpcRequest(param0 *ec2.CreateVpcInput) (*request.Request, *ec2.CreateVpcOutput) {
	m.addCall("CreateVpcRequest")
	m.verifyInput("CreateVpcRequest", param0)
	return m.CreateVpcRequestFunc(param0)
}

func (m *ec2Mock) CreateVpcWithContext(param0 aws.Context, param1 *ec2.CreateVpcInput, param2 ...request.Option) (*ec2.CreateVpcOutput, error) {
	m.addCall("CreateVpcWithContext")
	m.verifyInput("CreateVpcWithContext", param0)
	return m.CreateVpcWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CreateVpnConnection(param0 *ec2.CreateVpnConnectionInput) (*ec2.CreateVpnConnectionOutput, error) {
	m.addCall("CreateVpnConnection")
	m.verifyInput("CreateVpnConnection", param0)
	return m.CreateVpnConnectionFunc(param0)
}

func (m *ec2Mock) CreateVpnConnectionRequest(param0 *ec2.CreateVpnConnectionInput) (*request.Request, *ec2.CreateVpnConnectionOutput) {
	m.addCall("CreateVpnConnectionRequest")
	m.verifyInput("CreateVpnConnectionRequest", param0)
	return m.CreateVpnConnectionRequestFunc(param0)
}

func (m *ec2Mock) CreateVpnConnectionRoute(param0 *ec2.CreateVpnConnectionRouteInput) (*ec2.CreateVpnConnectionRouteOutput, error) {
	m.addCall("CreateVpnConnectionRoute")
	m.verifyInput("CreateVpnConnectionRoute", param0)
	return m.CreateVpnConnectionRouteFunc(param0)
}

func (m *ec2Mock) CreateVpnConnectionRouteRequest(param0 *ec2.CreateVpnConnectionRouteInput) (*request.Request, *ec2.CreateVpnConnectionRouteOutput) {
	m.addCall("CreateVpnConnectionRouteRequest")
	m.verifyInput("CreateVpnConnectionRouteRequest", param0)
	return m.CreateVpnConnectionRouteRequestFunc(param0)
}

func (m *ec2Mock) CreateVpnConnectionRouteWithContext(param0 aws.Context, param1 *ec2.CreateVpnConnectionRouteInput, param2 ...request.Option) (*ec2.CreateVpnConnectionRouteOutput, error) {
	m.addCall("CreateVpnConnectionRouteWithContext")
	m.verifyInput("CreateVpnConnectionRouteWithContext", param0)
	return m.CreateVpnConnectionRouteWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CreateVpnConnectionWithContext(param0 aws.Context, param1 *ec2.CreateVpnConnectionInput, param2 ...request.Option) (*ec2.CreateVpnConnectionOutput, error) {
	m.addCall("CreateVpnConnectionWithContext")
	m.verifyInput("CreateVpnConnectionWithContext", param0)
	return m.CreateVpnConnectionWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) CreateVpnGateway(param0 *ec2.CreateVpnGatewayInput) (*ec2.CreateVpnGatewayOutput, error) {
	m.addCall("CreateVpnGateway")
	m.verifyInput("CreateVpnGateway", param0)
	return m.CreateVpnGatewayFunc(param0)
}

func (m *ec2Mock) CreateVpnGatewayRequest(param0 *ec2.CreateVpnGatewayInput) (*request.Request, *ec2.CreateVpnGatewayOutput) {
	m.addCall("CreateVpnGatewayRequest")
	m.verifyInput("CreateVpnGatewayRequest", param0)
	return m.CreateVpnGatewayRequestFunc(param0)
}

func (m *ec2Mock) CreateVpnGatewayWithContext(param0 aws.Context, param1 *ec2.CreateVpnGatewayInput, param2 ...request.Option) (*ec2.CreateVpnGatewayOutput, error) {
	m.addCall("CreateVpnGatewayWithContext")
	m.verifyInput("CreateVpnGatewayWithContext", param0)
	return m.CreateVpnGatewayWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DeleteCustomerGateway(param0 *ec2.DeleteCustomerGatewayInput) (*ec2.DeleteCustomerGatewayOutput, error) {
	m.addCall("DeleteCustomerGateway")
	m.verifyInput("DeleteCustomerGateway", param0)
	return m.DeleteCustomerGatewayFunc(param0)
}

func (m *ec2Mock) DeleteCustomerGatewayRequest(param0 *ec2.DeleteCustomerGatewayInput) (*request.Request, *ec2.DeleteCustomerGatewayOutput) {
	m.addCall("DeleteCustomerGatewayRequest")
	m.verifyInput("DeleteCustomerGatewayRequest", param0)
	return m.DeleteCustomerGatewayRequestFunc(param0)
}

func (m *ec2Mock) DeleteCustomerGatewayWithContext(param0 aws.Context, param1 *ec2.DeleteCustomerGatewayInput, param2 ...request.Option) (*ec2.DeleteCustomerGatewayOutput, error) {
	m.addCall("DeleteCustomerGatewayWithContext")
	m.verifyInput("DeleteCustomerGatewayWithContext", param0)
	return m.DeleteCustomerGatewayWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DeleteDhcpOptions(param0 *ec2.DeleteDhcpOptionsInput) (*ec2.DeleteDhcpOptionsOutput, error) {
	m.addCall("DeleteDhcpOptions")
	m.verifyInput("DeleteDhcpOptions", param0)
	return m.DeleteDhcpOptionsFunc(param0)
}

func (m *ec2Mock) DeleteDhcpOptionsRequest(param0 *ec2.DeleteDhcpOptionsInput) (*request.Request, *ec2.DeleteDhcpOptionsOutput) {
	m.addCall("DeleteDhcpOptionsRequest")
	m.verifyInput("DeleteDhcpOptionsRequest", param0)
	return m.DeleteDhcpOptionsRequestFunc(param0)
}

func (m *ec2Mock) DeleteDhcpOptionsWithContext(param0 aws.Context, param1 *ec2.DeleteDhcpOptionsInput, param2 ...request.Option) (*ec2.DeleteDhcpOptionsOutput, error) {
	m.addCall("DeleteDhcpOptionsWithContext")
	m.verifyInput("DeleteDhcpOptionsWithContext", param0)
	return m.DeleteDhcpOptionsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DeleteEgressOnlyInternetGateway(param0 *ec2.DeleteEgressOnlyInternetGatewayInput) (*ec2.DeleteEgressOnlyInternetGatewayOutput, error) {
	m.addCall("DeleteEgressOnlyInternetGateway")
	m.verifyInput("DeleteEgressOnlyInternetGateway", param0)
	return m.DeleteEgressOnlyInternetGatewayFunc(param0)
}

func (m *ec2Mock) DeleteEgressOnlyInternetGatewayRequest(param0 *ec2.DeleteEgressOnlyInternetGatewayInput) (*request.Request, *ec2.DeleteEgressOnlyInternetGatewayOutput) {
	m.addCall("DeleteEgressOnlyInternetGatewayRequest")
	m.verifyInput("DeleteEgressOnlyInternetGatewayRequest", param0)
	return m.DeleteEgressOnlyInternetGatewayRequestFunc(param0)
}

func (m *ec2Mock) DeleteEgressOnlyInternetGatewayWithContext(param0 aws.Context, param1 *ec2.DeleteEgressOnlyInternetGatewayInput, param2 ...request.Option) (*ec2.DeleteEgressOnlyInternetGatewayOutput, error) {
	m.addCall("DeleteEgressOnlyInternetGatewayWithContext")
	m.verifyInput("DeleteEgressOnlyInternetGatewayWithContext", param0)
	return m.DeleteEgressOnlyInternetGatewayWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DeleteFlowLogs(param0 *ec2.DeleteFlowLogsInput) (*ec2.DeleteFlowLogsOutput, error) {
	m.addCall("DeleteFlowLogs")
	m.verifyInput("DeleteFlowLogs", param0)
	return m.DeleteFlowLogsFunc(param0)
}

func (m *ec2Mock) DeleteFlowLogsRequest(param0 *ec2.DeleteFlowLogsInput) (*request.Request, *ec2.DeleteFlowLogsOutput) {
	m.addCall("DeleteFlowLogsRequest")
	m.verifyInput("DeleteFlowLogsRequest", param0)
	return m.DeleteFlowLogsRequestFunc(param0)
}

func (m *ec2Mock) DeleteFlowLogsWithContext(param0 aws.Context, param1 *ec2.DeleteFlowLogsInput, param2 ...request.Option) (*ec2.DeleteFlowLogsOutput, error) {
	m.addCall("DeleteFlowLogsWithContext")
	m.verifyInput("DeleteFlowLogsWithContext", param0)
	return m.DeleteFlowLogsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DeleteFpgaImage(param0 *ec2.DeleteFpgaImageInput) (*ec2.DeleteFpgaImageOutput, error) {
	m.addCall("DeleteFpgaImage")
	m.verifyInput("DeleteFpgaImage", param0)
	return m.DeleteFpgaImageFunc(param0)
}

func (m *ec2Mock) DeleteFpgaImageRequest(param0 *ec2.DeleteFpgaImageInput) (*request.Request, *ec2.DeleteFpgaImageOutput) {
	m.addCall("DeleteFpgaImageRequest")
	m.verifyInput("DeleteFpgaImageRequest", param0)
	return m.DeleteFpgaImageRequestFunc(param0)
}

func (m *ec2Mock) DeleteFpgaImageWithContext(param0 aws.Context, param1 *ec2.DeleteFpgaImageInput, param2 ...request.Option) (*ec2.DeleteFpgaImageOutput, error) {
	m.addCall("DeleteFpgaImageWithContext")
	m.verifyInput("DeleteFpgaImageWithContext", param0)
	return m.DeleteFpgaImageWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DeleteInternetGateway(param0 *ec2.DeleteInternetGatewayInput) (*ec2.DeleteInternetGatewayOutput, error) {
	m.addCall("DeleteInternetGateway")
	m.verifyInput("DeleteInternetGateway", param0)
	return m.DeleteInternetGatewayFunc(param0)
}

func (m *ec2Mock) DeleteInternetGatewayRequest(param0 *ec2.DeleteInternetGatewayInput) (*request.Request, *ec2.DeleteInternetGatewayOutput) {
	m.addCall("DeleteInternetGatewayRequest")
	m.verifyInput("DeleteInternetGatewayRequest", param0)
	return m.DeleteInternetGatewayRequestFunc(param0)
}

func (m *ec2Mock) DeleteInternetGatewayWithContext(param0 aws.Context, param1 *ec2.DeleteInternetGatewayInput, param2 ...request.Option) (*ec2.DeleteInternetGatewayOutput, error) {
	m.addCall("DeleteInternetGatewayWithContext")
	m.verifyInput("DeleteInternetGatewayWithContext", param0)
	return m.DeleteInternetGatewayWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DeleteKeyPair(param0 *ec2.DeleteKeyPairInput) (*ec2.DeleteKeyPairOutput, error) {
	m.addCall("DeleteKeyPair")
	m.verifyInput("DeleteKeyPair", param0)
	return m.DeleteKeyPairFunc(param0)
}

func (m *ec2Mock) DeleteKeyPairRequest(param0 *ec2.DeleteKeyPairInput) (*request.Request, *ec2.DeleteKeyPairOutput) {
	m.addCall("DeleteKeyPairRequest")
	m.verifyInput("DeleteKeyPairRequest", param0)
	return m.DeleteKeyPairRequestFunc(param0)
}

func (m *ec2Mock) DeleteKeyPairWithContext(param0 aws.Context, param1 *ec2.DeleteKeyPairInput, param2 ...request.Option) (*ec2.DeleteKeyPairOutput, error) {
	m.addCall("DeleteKeyPairWithContext")
	m.verifyInput("DeleteKeyPairWithContext", param0)
	return m.DeleteKeyPairWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DeleteNatGateway(param0 *ec2.DeleteNatGatewayInput) (*ec2.DeleteNatGatewayOutput, error) {
	m.addCall("DeleteNatGateway")
	m.verifyInput("DeleteNatGateway", param0)
	return m.DeleteNatGatewayFunc(param0)
}

func (m *ec2Mock) DeleteNatGatewayRequest(param0 *ec2.DeleteNatGatewayInput) (*request.Request, *ec2.DeleteNatGatewayOutput) {
	m.addCall("DeleteNatGatewayRequest")
	m.verifyInput("DeleteNatGatewayRequest", param0)
	return m.DeleteNatGatewayRequestFunc(param0)
}

func (m *ec2Mock) DeleteNatGatewayWithContext(param0 aws.Context, param1 *ec2.DeleteNatGatewayInput, param2 ...request.Option) (*ec2.DeleteNatGatewayOutput, error) {
	m.addCall("DeleteNatGatewayWithContext")
	m.verifyInput("DeleteNatGatewayWithContext", param0)
	return m.DeleteNatGatewayWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DeleteNetworkAcl(param0 *ec2.DeleteNetworkAclInput) (*ec2.DeleteNetworkAclOutput, error) {
	m.addCall("DeleteNetworkAcl")
	m.verifyInput("DeleteNetworkAcl", param0)
	return m.DeleteNetworkAclFunc(param0)
}

func (m *ec2Mock) DeleteNetworkAclEntry(param0 *ec2.DeleteNetworkAclEntryInput) (*ec2.DeleteNetworkAclEntryOutput, error) {
	m.addCall("DeleteNetworkAclEntry")
	m.verifyInput("DeleteNetworkAclEntry", param0)
	return m.DeleteNetworkAclEntryFunc(param0)
}

func (m *ec2Mock) DeleteNetworkAclEntryRequest(param0 *ec2.DeleteNetworkAclEntryInput) (*request.Request, *ec2.DeleteNetworkAclEntryOutput) {
	m.addCall("DeleteNetworkAclEntryRequest")
	m.verifyInput("DeleteNetworkAclEntryRequest", param0)
	return m.DeleteNetworkAclEntryRequestFunc(param0)
}

func (m *ec2Mock) DeleteNetworkAclEntryWithContext(param0 aws.Context, param1 *ec2.DeleteNetworkAclEntryInput, param2 ...request.Option) (*ec2.DeleteNetworkAclEntryOutput, error) {
	m.addCall("DeleteNetworkAclEntryWithContext")
	m.verifyInput("DeleteNetworkAclEntryWithContext", param0)
	return m.DeleteNetworkAclEntryWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DeleteNetworkAclRequest(param0 *ec2.DeleteNetworkAclInput) (*request.Request, *ec2.DeleteNetworkAclOutput) {
	m.addCall("DeleteNetworkAclRequest")
	m.verifyInput("DeleteNetworkAclRequest", param0)
	return m.DeleteNetworkAclRequestFunc(param0)
}

func (m *ec2Mock) DeleteNetworkAclWithContext(param0 aws.Context, param1 *ec2.DeleteNetworkAclInput, param2 ...request.Option) (*ec2.DeleteNetworkAclOutput, error) {
	m.addCall("DeleteNetworkAclWithContext")
	m.verifyInput("DeleteNetworkAclWithContext", param0)
	return m.DeleteNetworkAclWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DeleteNetworkInterface(param0 *ec2.DeleteNetworkInterfaceInput) (*ec2.DeleteNetworkInterfaceOutput, error) {
	m.addCall("DeleteNetworkInterface")
	m.verifyInput("DeleteNetworkInterface", param0)
	return m.DeleteNetworkInterfaceFunc(param0)
}

func (m *ec2Mock) DeleteNetworkInterfacePermission(param0 *ec2.DeleteNetworkInterfacePermissionInput) (*ec2.DeleteNetworkInterfacePermissionOutput, error) {
	m.addCall("DeleteNetworkInterfacePermission")
	m.verifyInput("DeleteNetworkInterfacePermission", param0)
	return m.DeleteNetworkInterfacePermissionFunc(param0)
}

func (m *ec2Mock) DeleteNetworkInterfacePermissionRequest(param0 *ec2.DeleteNetworkInterfacePermissionInput) (*request.Request, *ec2.DeleteNetworkInterfacePermissionOutput) {
	m.addCall("DeleteNetworkInterfacePermissionRequest")
	m.verifyInput("DeleteNetworkInterfacePermissionRequest", param0)
	return m.DeleteNetworkInterfacePermissionRequestFunc(param0)
}

func (m *ec2Mock) DeleteNetworkInterfacePermissionWithContext(param0 aws.Context, param1 *ec2.DeleteNetworkInterfacePermissionInput, param2 ...request.Option) (*ec2.DeleteNetworkInterfacePermissionOutput, error) {
	m.addCall("DeleteNetworkInterfacePermissionWithContext")
	m.verifyInput("DeleteNetworkInterfacePermissionWithContext", param0)
	return m.DeleteNetworkInterfacePermissionWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DeleteNetworkInterfaceRequest(param0 *ec2.DeleteNetworkInterfaceInput) (*request.Request, *ec2.DeleteNetworkInterfaceOutput) {
	m.addCall("DeleteNetworkInterfaceRequest")
	m.verifyInput("DeleteNetworkInterfaceRequest", param0)
	return m.DeleteNetworkInterfaceRequestFunc(param0)
}

func (m *ec2Mock) DeleteNetworkInterfaceWithContext(param0 aws.Context, param1 *ec2.DeleteNetworkInterfaceInput, param2 ...request.Option) (*ec2.DeleteNetworkInterfaceOutput, error) {
	m.addCall("DeleteNetworkInterfaceWithContext")
	m.verifyInput("DeleteNetworkInterfaceWithContext", param0)
	return m.DeleteNetworkInterfaceWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DeletePlacementGroup(param0 *ec2.DeletePlacementGroupInput) (*ec2.DeletePlacementGroupOutput, error) {
	m.addCall("DeletePlacementGroup")
	m.verifyInput("DeletePlacementGroup", param0)
	return m.DeletePlacementGroupFunc(param0)
}

func (m *ec2Mock) DeletePlacementGroupRequest(param0 *ec2.DeletePlacementGroupInput) (*request.Request, *ec2.DeletePlacementGroupOutput) {
	m.addCall("DeletePlacementGroupRequest")
	m.verifyInput("DeletePlacementGroupRequest", param0)
	return m.DeletePlacementGroupRequestFunc(param0)
}

func (m *ec2Mock) DeletePlacementGroupWithContext(param0 aws.Context, param1 *ec2.DeletePlacementGroupInput, param2 ...request.Option) (*ec2.DeletePlacementGroupOutput, error) {
	m.addCall("DeletePlacementGroupWithContext")
	m.verifyInput("DeletePlacementGroupWithContext", param0)
	return m.DeletePlacementGroupWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DeleteRoute(param0 *ec2.DeleteRouteInput) (*ec2.DeleteRouteOutput, error) {
	m.addCall("DeleteRoute")
	m.verifyInput("DeleteRoute", param0)
	return m.DeleteRouteFunc(param0)
}

func (m *ec2Mock) DeleteRouteRequest(param0 *ec2.DeleteRouteInput) (*request.Request, *ec2.DeleteRouteOutput) {
	m.addCall("DeleteRouteRequest")
	m.verifyInput("DeleteRouteRequest", param0)
	return m.DeleteRouteRequestFunc(param0)
}

func (m *ec2Mock) DeleteRouteTable(param0 *ec2.DeleteRouteTableInput) (*ec2.DeleteRouteTableOutput, error) {
	m.addCall("DeleteRouteTable")
	m.verifyInput("DeleteRouteTable", param0)
	return m.DeleteRouteTableFunc(param0)
}

func (m *ec2Mock) DeleteRouteTableRequest(param0 *ec2.DeleteRouteTableInput) (*request.Request, *ec2.DeleteRouteTableOutput) {
	m.addCall("DeleteRouteTableRequest")
	m.verifyInput("DeleteRouteTableRequest", param0)
	return m.DeleteRouteTableRequestFunc(param0)
}

func (m *ec2Mock) DeleteRouteTableWithContext(param0 aws.Context, param1 *ec2.DeleteRouteTableInput, param2 ...request.Option) (*ec2.DeleteRouteTableOutput, error) {
	m.addCall("DeleteRouteTableWithContext")
	m.verifyInput("DeleteRouteTableWithContext", param0)
	return m.DeleteRouteTableWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DeleteRouteWithContext(param0 aws.Context, param1 *ec2.DeleteRouteInput, param2 ...request.Option) (*ec2.DeleteRouteOutput, error) {
	m.addCall("DeleteRouteWithContext")
	m.verifyInput("DeleteRouteWithContext", param0)
	return m.DeleteRouteWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DeleteSecurityGroup(param0 *ec2.DeleteSecurityGroupInput) (*ec2.DeleteSecurityGroupOutput, error) {
	m.addCall("DeleteSecurityGroup")
	m.verifyInput("DeleteSecurityGroup", param0)
	return m.DeleteSecurityGroupFunc(param0)
}

func (m *ec2Mock) DeleteSecurityGroupRequest(param0 *ec2.DeleteSecurityGroupInput) (*request.Request, *ec2.DeleteSecurityGroupOutput) {
	m.addCall("DeleteSecurityGroupRequest")
	m.verifyInput("DeleteSecurityGroupRequest", param0)
	return m.DeleteSecurityGroupRequestFunc(param0)
}

func (m *ec2Mock) DeleteSecurityGroupWithContext(param0 aws.Context, param1 *ec2.DeleteSecurityGroupInput, param2 ...request.Option) (*ec2.DeleteSecurityGroupOutput, error) {
	m.addCall("DeleteSecurityGroupWithContext")
	m.verifyInput("DeleteSecurityGroupWithContext", param0)
	return m.DeleteSecurityGroupWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DeleteSnapshot(param0 *ec2.DeleteSnapshotInput) (*ec2.DeleteSnapshotOutput, error) {
	m.addCall("DeleteSnapshot")
	m.verifyInput("DeleteSnapshot", param0)
	return m.DeleteSnapshotFunc(param0)
}

func (m *ec2Mock) DeleteSnapshotRequest(param0 *ec2.DeleteSnapshotInput) (*request.Request, *ec2.DeleteSnapshotOutput) {
	m.addCall("DeleteSnapshotRequest")
	m.verifyInput("DeleteSnapshotRequest", param0)
	return m.DeleteSnapshotRequestFunc(param0)
}

func (m *ec2Mock) DeleteSnapshotWithContext(param0 aws.Context, param1 *ec2.DeleteSnapshotInput, param2 ...request.Option) (*ec2.DeleteSnapshotOutput, error) {
	m.addCall("DeleteSnapshotWithContext")
	m.verifyInput("DeleteSnapshotWithContext", param0)
	return m.DeleteSnapshotWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DeleteSpotDatafeedSubscription(param0 *ec2.DeleteSpotDatafeedSubscriptionInput) (*ec2.DeleteSpotDatafeedSubscriptionOutput, error) {
	m.addCall("DeleteSpotDatafeedSubscription")
	m.verifyInput("DeleteSpotDatafeedSubscription", param0)
	return m.DeleteSpotDatafeedSubscriptionFunc(param0)
}

func (m *ec2Mock) DeleteSpotDatafeedSubscriptionRequest(param0 *ec2.DeleteSpotDatafeedSubscriptionInput) (*request.Request, *ec2.DeleteSpotDatafeedSubscriptionOutput) {
	m.addCall("DeleteSpotDatafeedSubscriptionRequest")
	m.verifyInput("DeleteSpotDatafeedSubscriptionRequest", param0)
	return m.DeleteSpotDatafeedSubscriptionRequestFunc(param0)
}

func (m *ec2Mock) DeleteSpotDatafeedSubscriptionWithContext(param0 aws.Context, param1 *ec2.DeleteSpotDatafeedSubscriptionInput, param2 ...request.Option) (*ec2.DeleteSpotDatafeedSubscriptionOutput, error) {
	m.addCall("DeleteSpotDatafeedSubscriptionWithContext")
	m.verifyInput("DeleteSpotDatafeedSubscriptionWithContext", param0)
	return m.DeleteSpotDatafeedSubscriptionWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DeleteSubnet(param0 *ec2.DeleteSubnetInput) (*ec2.DeleteSubnetOutput, error) {
	m.addCall("DeleteSubnet")
	m.verifyInput("DeleteSubnet", param0)
	return m.DeleteSubnetFunc(param0)
}

func (m *ec2Mock) DeleteSubnetRequest(param0 *ec2.DeleteSubnetInput) (*request.Request, *ec2.DeleteSubnetOutput) {
	m.addCall("DeleteSubnetRequest")
	m.verifyInput("DeleteSubnetRequest", param0)
	return m.DeleteSubnetRequestFunc(param0)
}

func (m *ec2Mock) DeleteSubnetWithContext(param0 aws.Context, param1 *ec2.DeleteSubnetInput, param2 ...request.Option) (*ec2.DeleteSubnetOutput, error) {
	m.addCall("DeleteSubnetWithContext")
	m.verifyInput("DeleteSubnetWithContext", param0)
	return m.DeleteSubnetWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DeleteTags(param0 *ec2.DeleteTagsInput) (*ec2.DeleteTagsOutput, error) {
	m.addCall("DeleteTags")
	m.verifyInput("DeleteTags", param0)
	return m.DeleteTagsFunc(param0)
}

func (m *ec2Mock) DeleteTagsRequest(param0 *ec2.DeleteTagsInput) (*request.Request, *ec2.DeleteTagsOutput) {
	m.addCall("DeleteTagsRequest")
	m.verifyInput("DeleteTagsRequest", param0)
	return m.DeleteTagsRequestFunc(param0)
}

func (m *ec2Mock) DeleteTagsWithContext(param0 aws.Context, param1 *ec2.DeleteTagsInput, param2 ...request.Option) (*ec2.DeleteTagsOutput, error) {
	m.addCall("DeleteTagsWithContext")
	m.verifyInput("DeleteTagsWithContext", param0)
	return m.DeleteTagsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DeleteVolume(param0 *ec2.DeleteVolumeInput) (*ec2.DeleteVolumeOutput, error) {
	m.addCall("DeleteVolume")
	m.verifyInput("DeleteVolume", param0)
	return m.DeleteVolumeFunc(param0)
}

func (m *ec2Mock) DeleteVolumeRequest(param0 *ec2.DeleteVolumeInput) (*request.Request, *ec2.DeleteVolumeOutput) {
	m.addCall("DeleteVolumeRequest")
	m.verifyInput("DeleteVolumeRequest", param0)
	return m.DeleteVolumeRequestFunc(param0)
}

func (m *ec2Mock) DeleteVolumeWithContext(param0 aws.Context, param1 *ec2.DeleteVolumeInput, param2 ...request.Option) (*ec2.DeleteVolumeOutput, error) {
	m.addCall("DeleteVolumeWithContext")
	m.verifyInput("DeleteVolumeWithContext", param0)
	return m.DeleteVolumeWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DeleteVpc(param0 *ec2.DeleteVpcInput) (*ec2.DeleteVpcOutput, error) {
	m.addCall("DeleteVpc")
	m.verifyInput("DeleteVpc", param0)
	return m.DeleteVpcFunc(param0)
}

func (m *ec2Mock) DeleteVpcEndpoints(param0 *ec2.DeleteVpcEndpointsInput) (*ec2.DeleteVpcEndpointsOutput, error) {
	m.addCall("DeleteVpcEndpoints")
	m.verifyInput("DeleteVpcEndpoints", param0)
	return m.DeleteVpcEndpointsFunc(param0)
}

func (m *ec2Mock) DeleteVpcEndpointsRequest(param0 *ec2.DeleteVpcEndpointsInput) (*request.Request, *ec2.DeleteVpcEndpointsOutput) {
	m.addCall("DeleteVpcEndpointsRequest")
	m.verifyInput("DeleteVpcEndpointsRequest", param0)
	return m.DeleteVpcEndpointsRequestFunc(param0)
}

func (m *ec2Mock) DeleteVpcEndpointsWithContext(param0 aws.Context, param1 *ec2.DeleteVpcEndpointsInput, param2 ...request.Option) (*ec2.DeleteVpcEndpointsOutput, error) {
	m.addCall("DeleteVpcEndpointsWithContext")
	m.verifyInput("DeleteVpcEndpointsWithContext", param0)
	return m.DeleteVpcEndpointsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DeleteVpcPeeringConnection(param0 *ec2.DeleteVpcPeeringConnectionInput) (*ec2.DeleteVpcPeeringConnectionOutput, error) {
	m.addCall("DeleteVpcPeeringConnection")
	m.verifyInput("DeleteVpcPeeringConnection", param0)
	return m.DeleteVpcPeeringConnectionFunc(param0)
}

func (m *ec2Mock) DeleteVpcPeeringConnectionRequest(param0 *ec2.DeleteVpcPeeringConnectionInput) (*request.Request, *ec2.DeleteVpcPeeringConnectionOutput) {
	m.addCall("DeleteVpcPeeringConnectionRequest")
	m.verifyInput("DeleteVpcPeeringConnectionRequest", param0)
	return m.DeleteVpcPeeringConnectionRequestFunc(param0)
}

func (m *ec2Mock) DeleteVpcPeeringConnectionWithContext(param0 aws.Context, param1 *ec2.DeleteVpcPeeringConnectionInput, param2 ...request.Option) (*ec2.DeleteVpcPeeringConnectionOutput, error) {
	m.addCall("DeleteVpcPeeringConnectionWithContext")
	m.verifyInput("DeleteVpcPeeringConnectionWithContext", param0)
	return m.DeleteVpcPeeringConnectionWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DeleteVpcRequest(param0 *ec2.DeleteVpcInput) (*request.Request, *ec2.DeleteVpcOutput) {
	m.addCall("DeleteVpcRequest")
	m.verifyInput("DeleteVpcRequest", param0)
	return m.DeleteVpcRequestFunc(param0)
}

func (m *ec2Mock) DeleteVpcWithContext(param0 aws.Context, param1 *ec2.DeleteVpcInput, param2 ...request.Option) (*ec2.DeleteVpcOutput, error) {
	m.addCall("DeleteVpcWithContext")
	m.verifyInput("DeleteVpcWithContext", param0)
	return m.DeleteVpcWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DeleteVpnConnection(param0 *ec2.DeleteVpnConnectionInput) (*ec2.DeleteVpnConnectionOutput, error) {
	m.addCall("DeleteVpnConnection")
	m.verifyInput("DeleteVpnConnection", param0)
	return m.DeleteVpnConnectionFunc(param0)
}

func (m *ec2Mock) DeleteVpnConnectionRequest(param0 *ec2.DeleteVpnConnectionInput) (*request.Request, *ec2.DeleteVpnConnectionOutput) {
	m.addCall("DeleteVpnConnectionRequest")
	m.verifyInput("DeleteVpnConnectionRequest", param0)
	return m.DeleteVpnConnectionRequestFunc(param0)
}

func (m *ec2Mock) DeleteVpnConnectionRoute(param0 *ec2.DeleteVpnConnectionRouteInput) (*ec2.DeleteVpnConnectionRouteOutput, error) {
	m.addCall("DeleteVpnConnectionRoute")
	m.verifyInput("DeleteVpnConnectionRoute", param0)
	return m.DeleteVpnConnectionRouteFunc(param0)
}

func (m *ec2Mock) DeleteVpnConnectionRouteRequest(param0 *ec2.DeleteVpnConnectionRouteInput) (*request.Request, *ec2.DeleteVpnConnectionRouteOutput) {
	m.addCall("DeleteVpnConnectionRouteRequest")
	m.verifyInput("DeleteVpnConnectionRouteRequest", param0)
	return m.DeleteVpnConnectionRouteRequestFunc(param0)
}

func (m *ec2Mock) DeleteVpnConnectionRouteWithContext(param0 aws.Context, param1 *ec2.DeleteVpnConnectionRouteInput, param2 ...request.Option) (*ec2.DeleteVpnConnectionRouteOutput, error) {
	m.addCall("DeleteVpnConnectionRouteWithContext")
	m.verifyInput("DeleteVpnConnectionRouteWithContext", param0)
	return m.DeleteVpnConnectionRouteWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DeleteVpnConnectionWithContext(param0 aws.Context, param1 *ec2.DeleteVpnConnectionInput, param2 ...request.Option) (*ec2.DeleteVpnConnectionOutput, error) {
	m.addCall("DeleteVpnConnectionWithContext")
	m.verifyInput("DeleteVpnConnectionWithContext", param0)
	return m.DeleteVpnConnectionWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DeleteVpnGateway(param0 *ec2.DeleteVpnGatewayInput) (*ec2.DeleteVpnGatewayOutput, error) {
	m.addCall("DeleteVpnGateway")
	m.verifyInput("DeleteVpnGateway", param0)
	return m.DeleteVpnGatewayFunc(param0)
}

func (m *ec2Mock) DeleteVpnGatewayRequest(param0 *ec2.DeleteVpnGatewayInput) (*request.Request, *ec2.DeleteVpnGatewayOutput) {
	m.addCall("DeleteVpnGatewayRequest")
	m.verifyInput("DeleteVpnGatewayRequest", param0)
	return m.DeleteVpnGatewayRequestFunc(param0)
}

func (m *ec2Mock) DeleteVpnGatewayWithContext(param0 aws.Context, param1 *ec2.DeleteVpnGatewayInput, param2 ...request.Option) (*ec2.DeleteVpnGatewayOutput, error) {
	m.addCall("DeleteVpnGatewayWithContext")
	m.verifyInput("DeleteVpnGatewayWithContext", param0)
	return m.DeleteVpnGatewayWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DeregisterImage(param0 *ec2.DeregisterImageInput) (*ec2.DeregisterImageOutput, error) {
	m.addCall("DeregisterImage")
	m.verifyInput("DeregisterImage", param0)
	return m.DeregisterImageFunc(param0)
}

func (m *ec2Mock) DeregisterImageRequest(param0 *ec2.DeregisterImageInput) (*request.Request, *ec2.DeregisterImageOutput) {
	m.addCall("DeregisterImageRequest")
	m.verifyInput("DeregisterImageRequest", param0)
	return m.DeregisterImageRequestFunc(param0)
}

func (m *ec2Mock) DeregisterImageWithContext(param0 aws.Context, param1 *ec2.DeregisterImageInput, param2 ...request.Option) (*ec2.DeregisterImageOutput, error) {
	m.addCall("DeregisterImageWithContext")
	m.verifyInput("DeregisterImageWithContext", param0)
	return m.DeregisterImageWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeAccountAttributes(param0 *ec2.DescribeAccountAttributesInput) (*ec2.DescribeAccountAttributesOutput, error) {
	m.addCall("DescribeAccountAttributes")
	m.verifyInput("DescribeAccountAttributes", param0)
	return m.DescribeAccountAttributesFunc(param0)
}

func (m *ec2Mock) DescribeAccountAttributesRequest(param0 *ec2.DescribeAccountAttributesInput) (*request.Request, *ec2.DescribeAccountAttributesOutput) {
	m.addCall("DescribeAccountAttributesRequest")
	m.verifyInput("DescribeAccountAttributesRequest", param0)
	return m.DescribeAccountAttributesRequestFunc(param0)
}

func (m *ec2Mock) DescribeAccountAttributesWithContext(param0 aws.Context, param1 *ec2.DescribeAccountAttributesInput, param2 ...request.Option) (*ec2.DescribeAccountAttributesOutput, error) {
	m.addCall("DescribeAccountAttributesWithContext")
	m.verifyInput("DescribeAccountAttributesWithContext", param0)
	return m.DescribeAccountAttributesWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeAddresses(param0 *ec2.DescribeAddressesInput) (*ec2.DescribeAddressesOutput, error) {
	m.addCall("DescribeAddresses")
	m.verifyInput("DescribeAddresses", param0)
	return m.DescribeAddressesFunc(param0)
}

func (m *ec2Mock) DescribeAddressesRequest(param0 *ec2.DescribeAddressesInput) (*request.Request, *ec2.DescribeAddressesOutput) {
	m.addCall("DescribeAddressesRequest")
	m.verifyInput("DescribeAddressesRequest", param0)
	return m.DescribeAddressesRequestFunc(param0)
}

func (m *ec2Mock) DescribeAddressesWithContext(param0 aws.Context, param1 *ec2.DescribeAddressesInput, param2 ...request.Option) (*ec2.DescribeAddressesOutput, error) {
	m.addCall("DescribeAddressesWithContext")
	m.verifyInput("DescribeAddressesWithContext", param0)
	return m.DescribeAddressesWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeAvailabilityZones(param0 *ec2.DescribeAvailabilityZonesInput) (*ec2.DescribeAvailabilityZonesOutput, error) {
	m.addCall("DescribeAvailabilityZones")
	m.verifyInput("DescribeAvailabilityZones", param0)
	return m.DescribeAvailabilityZonesFunc(param0)
}

func (m *ec2Mock) DescribeAvailabilityZonesRequest(param0 *ec2.DescribeAvailabilityZonesInput) (*request.Request, *ec2.DescribeAvailabilityZonesOutput) {
	m.addCall("DescribeAvailabilityZonesRequest")
	m.verifyInput("DescribeAvailabilityZonesRequest", param0)
	return m.DescribeAvailabilityZonesRequestFunc(param0)
}

func (m *ec2Mock) DescribeAvailabilityZonesWithContext(param0 aws.Context, param1 *ec2.DescribeAvailabilityZonesInput, param2 ...request.Option) (*ec2.DescribeAvailabilityZonesOutput, error) {
	m.addCall("DescribeAvailabilityZonesWithContext")
	m.verifyInput("DescribeAvailabilityZonesWithContext", param0)
	return m.DescribeAvailabilityZonesWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeBundleTasks(param0 *ec2.DescribeBundleTasksInput) (*ec2.DescribeBundleTasksOutput, error) {
	m.addCall("DescribeBundleTasks")
	m.verifyInput("DescribeBundleTasks", param0)
	return m.DescribeBundleTasksFunc(param0)
}

func (m *ec2Mock) DescribeBundleTasksRequest(param0 *ec2.DescribeBundleTasksInput) (*request.Request, *ec2.DescribeBundleTasksOutput) {
	m.addCall("DescribeBundleTasksRequest")
	m.verifyInput("DescribeBundleTasksRequest", param0)
	return m.DescribeBundleTasksRequestFunc(param0)
}

func (m *ec2Mock) DescribeBundleTasksWithContext(param0 aws.Context, param1 *ec2.DescribeBundleTasksInput, param2 ...request.Option) (*ec2.DescribeBundleTasksOutput, error) {
	m.addCall("DescribeBundleTasksWithContext")
	m.verifyInput("DescribeBundleTasksWithContext", param0)
	return m.DescribeBundleTasksWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeClassicLinkInstances(param0 *ec2.DescribeClassicLinkInstancesInput) (*ec2.DescribeClassicLinkInstancesOutput, error) {
	m.addCall("DescribeClassicLinkInstances")
	m.verifyInput("DescribeClassicLinkInstances", param0)
	return m.DescribeClassicLinkInstancesFunc(param0)
}

func (m *ec2Mock) DescribeClassicLinkInstancesRequest(param0 *ec2.DescribeClassicLinkInstancesInput) (*request.Request, *ec2.DescribeClassicLinkInstancesOutput) {
	m.addCall("DescribeClassicLinkInstancesRequest")
	m.verifyInput("DescribeClassicLinkInstancesRequest", param0)
	return m.DescribeClassicLinkInstancesRequestFunc(param0)
}

func (m *ec2Mock) DescribeClassicLinkInstancesWithContext(param0 aws.Context, param1 *ec2.DescribeClassicLinkInstancesInput, param2 ...request.Option) (*ec2.DescribeClassicLinkInstancesOutput, error) {
	m.addCall("DescribeClassicLinkInstancesWithContext")
	m.verifyInput("DescribeClassicLinkInstancesWithContext", param0)
	return m.DescribeClassicLinkInstancesWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeConversionTasks(param0 *ec2.DescribeConversionTasksInput) (*ec2.DescribeConversionTasksOutput, error) {
	m.addCall("DescribeConversionTasks")
	m.verifyInput("DescribeConversionTasks", param0)
	return m.DescribeConversionTasksFunc(param0)
}

func (m *ec2Mock) DescribeConversionTasksRequest(param0 *ec2.DescribeConversionTasksInput) (*request.Request, *ec2.DescribeConversionTasksOutput) {
	m.addCall("DescribeConversionTasksRequest")
	m.verifyInput("DescribeConversionTasksRequest", param0)
	return m.DescribeConversionTasksRequestFunc(param0)
}

func (m *ec2Mock) DescribeConversionTasksWithContext(param0 aws.Context, param1 *ec2.DescribeConversionTasksInput, param2 ...request.Option) (*ec2.DescribeConversionTasksOutput, error) {
	m.addCall("DescribeConversionTasksWithContext")
	m.verifyInput("DescribeConversionTasksWithContext", param0)
	return m.DescribeConversionTasksWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeCustomerGateways(param0 *ec2.DescribeCustomerGatewaysInput) (*ec2.DescribeCustomerGatewaysOutput, error) {
	m.addCall("DescribeCustomerGateways")
	m.verifyInput("DescribeCustomerGateways", param0)
	return m.DescribeCustomerGatewaysFunc(param0)
}

func (m *ec2Mock) DescribeCustomerGatewaysRequest(param0 *ec2.DescribeCustomerGatewaysInput) (*request.Request, *ec2.DescribeCustomerGatewaysOutput) {
	m.addCall("DescribeCustomerGatewaysRequest")
	m.verifyInput("DescribeCustomerGatewaysRequest", param0)
	return m.DescribeCustomerGatewaysRequestFunc(param0)
}

func (m *ec2Mock) DescribeCustomerGatewaysWithContext(param0 aws.Context, param1 *ec2.DescribeCustomerGatewaysInput, param2 ...request.Option) (*ec2.DescribeCustomerGatewaysOutput, error) {
	m.addCall("DescribeCustomerGatewaysWithContext")
	m.verifyInput("DescribeCustomerGatewaysWithContext", param0)
	return m.DescribeCustomerGatewaysWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeDhcpOptions(param0 *ec2.DescribeDhcpOptionsInput) (*ec2.DescribeDhcpOptionsOutput, error) {
	m.addCall("DescribeDhcpOptions")
	m.verifyInput("DescribeDhcpOptions", param0)
	return m.DescribeDhcpOptionsFunc(param0)
}

func (m *ec2Mock) DescribeDhcpOptionsRequest(param0 *ec2.DescribeDhcpOptionsInput) (*request.Request, *ec2.DescribeDhcpOptionsOutput) {
	m.addCall("DescribeDhcpOptionsRequest")
	m.verifyInput("DescribeDhcpOptionsRequest", param0)
	return m.DescribeDhcpOptionsRequestFunc(param0)
}

func (m *ec2Mock) DescribeDhcpOptionsWithContext(param0 aws.Context, param1 *ec2.DescribeDhcpOptionsInput, param2 ...request.Option) (*ec2.DescribeDhcpOptionsOutput, error) {
	m.addCall("DescribeDhcpOptionsWithContext")
	m.verifyInput("DescribeDhcpOptionsWithContext", param0)
	return m.DescribeDhcpOptionsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeEgressOnlyInternetGateways(param0 *ec2.DescribeEgressOnlyInternetGatewaysInput) (*ec2.DescribeEgressOnlyInternetGatewaysOutput, error) {
	m.addCall("DescribeEgressOnlyInternetGateways")
	m.verifyInput("DescribeEgressOnlyInternetGateways", param0)
	return m.DescribeEgressOnlyInternetGatewaysFunc(param0)
}

func (m *ec2Mock) DescribeEgressOnlyInternetGatewaysRequest(param0 *ec2.DescribeEgressOnlyInternetGatewaysInput) (*request.Request, *ec2.DescribeEgressOnlyInternetGatewaysOutput) {
	m.addCall("DescribeEgressOnlyInternetGatewaysRequest")
	m.verifyInput("DescribeEgressOnlyInternetGatewaysRequest", param0)
	return m.DescribeEgressOnlyInternetGatewaysRequestFunc(param0)
}

func (m *ec2Mock) DescribeEgressOnlyInternetGatewaysWithContext(param0 aws.Context, param1 *ec2.DescribeEgressOnlyInternetGatewaysInput, param2 ...request.Option) (*ec2.DescribeEgressOnlyInternetGatewaysOutput, error) {
	m.addCall("DescribeEgressOnlyInternetGatewaysWithContext")
	m.verifyInput("DescribeEgressOnlyInternetGatewaysWithContext", param0)
	return m.DescribeEgressOnlyInternetGatewaysWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeElasticGpus(param0 *ec2.DescribeElasticGpusInput) (*ec2.DescribeElasticGpusOutput, error) {
	m.addCall("DescribeElasticGpus")
	m.verifyInput("DescribeElasticGpus", param0)
	return m.DescribeElasticGpusFunc(param0)
}

func (m *ec2Mock) DescribeElasticGpusRequest(param0 *ec2.DescribeElasticGpusInput) (*request.Request, *ec2.DescribeElasticGpusOutput) {
	m.addCall("DescribeElasticGpusRequest")
	m.verifyInput("DescribeElasticGpusRequest", param0)
	return m.DescribeElasticGpusRequestFunc(param0)
}

func (m *ec2Mock) DescribeElasticGpusWithContext(param0 aws.Context, param1 *ec2.DescribeElasticGpusInput, param2 ...request.Option) (*ec2.DescribeElasticGpusOutput, error) {
	m.addCall("DescribeElasticGpusWithContext")
	m.verifyInput("DescribeElasticGpusWithContext", param0)
	return m.DescribeElasticGpusWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeExportTasks(param0 *ec2.DescribeExportTasksInput) (*ec2.DescribeExportTasksOutput, error) {
	m.addCall("DescribeExportTasks")
	m.verifyInput("DescribeExportTasks", param0)
	return m.DescribeExportTasksFunc(param0)
}

func (m *ec2Mock) DescribeExportTasksRequest(param0 *ec2.DescribeExportTasksInput) (*request.Request, *ec2.DescribeExportTasksOutput) {
	m.addCall("DescribeExportTasksRequest")
	m.verifyInput("DescribeExportTasksRequest", param0)
	return m.DescribeExportTasksRequestFunc(param0)
}

func (m *ec2Mock) DescribeExportTasksWithContext(param0 aws.Context, param1 *ec2.DescribeExportTasksInput, param2 ...request.Option) (*ec2.DescribeExportTasksOutput, error) {
	m.addCall("DescribeExportTasksWithContext")
	m.verifyInput("DescribeExportTasksWithContext", param0)
	return m.DescribeExportTasksWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeFlowLogs(param0 *ec2.DescribeFlowLogsInput) (*ec2.DescribeFlowLogsOutput, error) {
	m.addCall("DescribeFlowLogs")
	m.verifyInput("DescribeFlowLogs", param0)
	return m.DescribeFlowLogsFunc(param0)
}

func (m *ec2Mock) DescribeFlowLogsRequest(param0 *ec2.DescribeFlowLogsInput) (*request.Request, *ec2.DescribeFlowLogsOutput) {
	m.addCall("DescribeFlowLogsRequest")
	m.verifyInput("DescribeFlowLogsRequest", param0)
	return m.DescribeFlowLogsRequestFunc(param0)
}

func (m *ec2Mock) DescribeFlowLogsWithContext(param0 aws.Context, param1 *ec2.DescribeFlowLogsInput, param2 ...request.Option) (*ec2.DescribeFlowLogsOutput, error) {
	m.addCall("DescribeFlowLogsWithContext")
	m.verifyInput("DescribeFlowLogsWithContext", param0)
	return m.DescribeFlowLogsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeFpgaImageAttribute(param0 *ec2.DescribeFpgaImageAttributeInput) (*ec2.DescribeFpgaImageAttributeOutput, error) {
	m.addCall("DescribeFpgaImageAttribute")
	m.verifyInput("DescribeFpgaImageAttribute", param0)
	return m.DescribeFpgaImageAttributeFunc(param0)
}

func (m *ec2Mock) DescribeFpgaImageAttributeRequest(param0 *ec2.DescribeFpgaImageAttributeInput) (*request.Request, *ec2.DescribeFpgaImageAttributeOutput) {
	m.addCall("DescribeFpgaImageAttributeRequest")
	m.verifyInput("DescribeFpgaImageAttributeRequest", param0)
	return m.DescribeFpgaImageAttributeRequestFunc(param0)
}

func (m *ec2Mock) DescribeFpgaImageAttributeWithContext(param0 aws.Context, param1 *ec2.DescribeFpgaImageAttributeInput, param2 ...request.Option) (*ec2.DescribeFpgaImageAttributeOutput, error) {
	m.addCall("DescribeFpgaImageAttributeWithContext")
	m.verifyInput("DescribeFpgaImageAttributeWithContext", param0)
	return m.DescribeFpgaImageAttributeWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeFpgaImages(param0 *ec2.DescribeFpgaImagesInput) (*ec2.DescribeFpgaImagesOutput, error) {
	m.addCall("DescribeFpgaImages")
	m.verifyInput("DescribeFpgaImages", param0)
	return m.DescribeFpgaImagesFunc(param0)
}

func (m *ec2Mock) DescribeFpgaImagesRequest(param0 *ec2.DescribeFpgaImagesInput) (*request.Request, *ec2.DescribeFpgaImagesOutput) {
	m.addCall("DescribeFpgaImagesRequest")
	m.verifyInput("DescribeFpgaImagesRequest", param0)
	return m.DescribeFpgaImagesRequestFunc(param0)
}

func (m *ec2Mock) DescribeFpgaImagesWithContext(param0 aws.Context, param1 *ec2.DescribeFpgaImagesInput, param2 ...request.Option) (*ec2.DescribeFpgaImagesOutput, error) {
	m.addCall("DescribeFpgaImagesWithContext")
	m.verifyInput("DescribeFpgaImagesWithContext", param0)
	return m.DescribeFpgaImagesWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeHostReservationOfferings(param0 *ec2.DescribeHostReservationOfferingsInput) (*ec2.DescribeHostReservationOfferingsOutput, error) {
	m.addCall("DescribeHostReservationOfferings")
	m.verifyInput("DescribeHostReservationOfferings", param0)
	return m.DescribeHostReservationOfferingsFunc(param0)
}

func (m *ec2Mock) DescribeHostReservationOfferingsRequest(param0 *ec2.DescribeHostReservationOfferingsInput) (*request.Request, *ec2.DescribeHostReservationOfferingsOutput) {
	m.addCall("DescribeHostReservationOfferingsRequest")
	m.verifyInput("DescribeHostReservationOfferingsRequest", param0)
	return m.DescribeHostReservationOfferingsRequestFunc(param0)
}

func (m *ec2Mock) DescribeHostReservationOfferingsWithContext(param0 aws.Context, param1 *ec2.DescribeHostReservationOfferingsInput, param2 ...request.Option) (*ec2.DescribeHostReservationOfferingsOutput, error) {
	m.addCall("DescribeHostReservationOfferingsWithContext")
	m.verifyInput("DescribeHostReservationOfferingsWithContext", param0)
	return m.DescribeHostReservationOfferingsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeHostReservations(param0 *ec2.DescribeHostReservationsInput) (*ec2.DescribeHostReservationsOutput, error) {
	m.addCall("DescribeHostReservations")
	m.verifyInput("DescribeHostReservations", param0)
	return m.DescribeHostReservationsFunc(param0)
}

func (m *ec2Mock) DescribeHostReservationsRequest(param0 *ec2.DescribeHostReservationsInput) (*request.Request, *ec2.DescribeHostReservationsOutput) {
	m.addCall("DescribeHostReservationsRequest")
	m.verifyInput("DescribeHostReservationsRequest", param0)
	return m.DescribeHostReservationsRequestFunc(param0)
}

func (m *ec2Mock) DescribeHostReservationsWithContext(param0 aws.Context, param1 *ec2.DescribeHostReservationsInput, param2 ...request.Option) (*ec2.DescribeHostReservationsOutput, error) {
	m.addCall("DescribeHostReservationsWithContext")
	m.verifyInput("DescribeHostReservationsWithContext", param0)
	return m.DescribeHostReservationsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeHosts(param0 *ec2.DescribeHostsInput) (*ec2.DescribeHostsOutput, error) {
	m.addCall("DescribeHosts")
	m.verifyInput("DescribeHosts", param0)
	return m.DescribeHostsFunc(param0)
}

func (m *ec2Mock) DescribeHostsRequest(param0 *ec2.DescribeHostsInput) (*request.Request, *ec2.DescribeHostsOutput) {
	m.addCall("DescribeHostsRequest")
	m.verifyInput("DescribeHostsRequest", param0)
	return m.DescribeHostsRequestFunc(param0)
}

func (m *ec2Mock) DescribeHostsWithContext(param0 aws.Context, param1 *ec2.DescribeHostsInput, param2 ...request.Option) (*ec2.DescribeHostsOutput, error) {
	m.addCall("DescribeHostsWithContext")
	m.verifyInput("DescribeHostsWithContext", param0)
	return m.DescribeHostsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeIamInstanceProfileAssociations(param0 *ec2.DescribeIamInstanceProfileAssociationsInput) (*ec2.DescribeIamInstanceProfileAssociationsOutput, error) {
	m.addCall("DescribeIamInstanceProfileAssociations")
	m.verifyInput("DescribeIamInstanceProfileAssociations", param0)
	return m.DescribeIamInstanceProfileAssociationsFunc(param0)
}

func (m *ec2Mock) DescribeIamInstanceProfileAssociationsRequest(param0 *ec2.DescribeIamInstanceProfileAssociationsInput) (*request.Request, *ec2.DescribeIamInstanceProfileAssociationsOutput) {
	m.addCall("DescribeIamInstanceProfileAssociationsRequest")
	m.verifyInput("DescribeIamInstanceProfileAssociationsRequest", param0)
	return m.DescribeIamInstanceProfileAssociationsRequestFunc(param0)
}

func (m *ec2Mock) DescribeIamInstanceProfileAssociationsWithContext(param0 aws.Context, param1 *ec2.DescribeIamInstanceProfileAssociationsInput, param2 ...request.Option) (*ec2.DescribeIamInstanceProfileAssociationsOutput, error) {
	m.addCall("DescribeIamInstanceProfileAssociationsWithContext")
	m.verifyInput("DescribeIamInstanceProfileAssociationsWithContext", param0)
	return m.DescribeIamInstanceProfileAssociationsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeIdFormat(param0 *ec2.DescribeIdFormatInput) (*ec2.DescribeIdFormatOutput, error) {
	m.addCall("DescribeIdFormat")
	m.verifyInput("DescribeIdFormat", param0)
	return m.DescribeIdFormatFunc(param0)
}

func (m *ec2Mock) DescribeIdFormatRequest(param0 *ec2.DescribeIdFormatInput) (*request.Request, *ec2.DescribeIdFormatOutput) {
	m.addCall("DescribeIdFormatRequest")
	m.verifyInput("DescribeIdFormatRequest", param0)
	return m.DescribeIdFormatRequestFunc(param0)
}

func (m *ec2Mock) DescribeIdFormatWithContext(param0 aws.Context, param1 *ec2.DescribeIdFormatInput, param2 ...request.Option) (*ec2.DescribeIdFormatOutput, error) {
	m.addCall("DescribeIdFormatWithContext")
	m.verifyInput("DescribeIdFormatWithContext", param0)
	return m.DescribeIdFormatWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeIdentityIdFormat(param0 *ec2.DescribeIdentityIdFormatInput) (*ec2.DescribeIdentityIdFormatOutput, error) {
	m.addCall("DescribeIdentityIdFormat")
	m.verifyInput("DescribeIdentityIdFormat", param0)
	return m.DescribeIdentityIdFormatFunc(param0)
}

func (m *ec2Mock) DescribeIdentityIdFormatRequest(param0 *ec2.DescribeIdentityIdFormatInput) (*request.Request, *ec2.DescribeIdentityIdFormatOutput) {
	m.addCall("DescribeIdentityIdFormatRequest")
	m.verifyInput("DescribeIdentityIdFormatRequest", param0)
	return m.DescribeIdentityIdFormatRequestFunc(param0)
}

func (m *ec2Mock) DescribeIdentityIdFormatWithContext(param0 aws.Context, param1 *ec2.DescribeIdentityIdFormatInput, param2 ...request.Option) (*ec2.DescribeIdentityIdFormatOutput, error) {
	m.addCall("DescribeIdentityIdFormatWithContext")
	m.verifyInput("DescribeIdentityIdFormatWithContext", param0)
	return m.DescribeIdentityIdFormatWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeImageAttribute(param0 *ec2.DescribeImageAttributeInput) (*ec2.DescribeImageAttributeOutput, error) {
	m.addCall("DescribeImageAttribute")
	m.verifyInput("DescribeImageAttribute", param0)
	return m.DescribeImageAttributeFunc(param0)
}

func (m *ec2Mock) DescribeImageAttributeRequest(param0 *ec2.DescribeImageAttributeInput) (*request.Request, *ec2.DescribeImageAttributeOutput) {
	m.addCall("DescribeImageAttributeRequest")
	m.verifyInput("DescribeImageAttributeRequest", param0)
	return m.DescribeImageAttributeRequestFunc(param0)
}

func (m *ec2Mock) DescribeImageAttributeWithContext(param0 aws.Context, param1 *ec2.DescribeImageAttributeInput, param2 ...request.Option) (*ec2.DescribeImageAttributeOutput, error) {
	m.addCall("DescribeImageAttributeWithContext")
	m.verifyInput("DescribeImageAttributeWithContext", param0)
	return m.DescribeImageAttributeWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeImages(param0 *ec2.DescribeImagesInput) (*ec2.DescribeImagesOutput, error) {
	m.addCall("DescribeImages")
	m.verifyInput("DescribeImages", param0)
	return m.DescribeImagesFunc(param0)
}

func (m *ec2Mock) DescribeImagesRequest(param0 *ec2.DescribeImagesInput) (*request.Request, *ec2.DescribeImagesOutput) {
	m.addCall("DescribeImagesRequest")
	m.verifyInput("DescribeImagesRequest", param0)
	return m.DescribeImagesRequestFunc(param0)
}

func (m *ec2Mock) DescribeImagesWithContext(param0 aws.Context, param1 *ec2.DescribeImagesInput, param2 ...request.Option) (*ec2.DescribeImagesOutput, error) {
	m.addCall("DescribeImagesWithContext")
	m.verifyInput("DescribeImagesWithContext", param0)
	return m.DescribeImagesWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeImportImageTasks(param0 *ec2.DescribeImportImageTasksInput) (*ec2.DescribeImportImageTasksOutput, error) {
	m.addCall("DescribeImportImageTasks")
	m.verifyInput("DescribeImportImageTasks", param0)
	return m.DescribeImportImageTasksFunc(param0)
}

func (m *ec2Mock) DescribeImportImageTasksRequest(param0 *ec2.DescribeImportImageTasksInput) (*request.Request, *ec2.DescribeImportImageTasksOutput) {
	m.addCall("DescribeImportImageTasksRequest")
	m.verifyInput("DescribeImportImageTasksRequest", param0)
	return m.DescribeImportImageTasksRequestFunc(param0)
}

func (m *ec2Mock) DescribeImportImageTasksWithContext(param0 aws.Context, param1 *ec2.DescribeImportImageTasksInput, param2 ...request.Option) (*ec2.DescribeImportImageTasksOutput, error) {
	m.addCall("DescribeImportImageTasksWithContext")
	m.verifyInput("DescribeImportImageTasksWithContext", param0)
	return m.DescribeImportImageTasksWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeImportSnapshotTasks(param0 *ec2.DescribeImportSnapshotTasksInput) (*ec2.DescribeImportSnapshotTasksOutput, error) {
	m.addCall("DescribeImportSnapshotTasks")
	m.verifyInput("DescribeImportSnapshotTasks", param0)
	return m.DescribeImportSnapshotTasksFunc(param0)
}

func (m *ec2Mock) DescribeImportSnapshotTasksRequest(param0 *ec2.DescribeImportSnapshotTasksInput) (*request.Request, *ec2.DescribeImportSnapshotTasksOutput) {
	m.addCall("DescribeImportSnapshotTasksRequest")
	m.verifyInput("DescribeImportSnapshotTasksRequest", param0)
	return m.DescribeImportSnapshotTasksRequestFunc(param0)
}

func (m *ec2Mock) DescribeImportSnapshotTasksWithContext(param0 aws.Context, param1 *ec2.DescribeImportSnapshotTasksInput, param2 ...request.Option) (*ec2.DescribeImportSnapshotTasksOutput, error) {
	m.addCall("DescribeImportSnapshotTasksWithContext")
	m.verifyInput("DescribeImportSnapshotTasksWithContext", param0)
	return m.DescribeImportSnapshotTasksWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeInstanceAttribute(param0 *ec2.DescribeInstanceAttributeInput) (*ec2.DescribeInstanceAttributeOutput, error) {
	m.addCall("DescribeInstanceAttribute")
	m.verifyInput("DescribeInstanceAttribute", param0)
	return m.DescribeInstanceAttributeFunc(param0)
}

func (m *ec2Mock) DescribeInstanceAttributeRequest(param0 *ec2.DescribeInstanceAttributeInput) (*request.Request, *ec2.DescribeInstanceAttributeOutput) {
	m.addCall("DescribeInstanceAttributeRequest")
	m.verifyInput("DescribeInstanceAttributeRequest", param0)
	return m.DescribeInstanceAttributeRequestFunc(param0)
}

func (m *ec2Mock) DescribeInstanceAttributeWithContext(param0 aws.Context, param1 *ec2.DescribeInstanceAttributeInput, param2 ...request.Option) (*ec2.DescribeInstanceAttributeOutput, error) {
	m.addCall("DescribeInstanceAttributeWithContext")
	m.verifyInput("DescribeInstanceAttributeWithContext", param0)
	return m.DescribeInstanceAttributeWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeInstanceStatus(param0 *ec2.DescribeInstanceStatusInput) (*ec2.DescribeInstanceStatusOutput, error) {
	m.addCall("DescribeInstanceStatus")
	m.verifyInput("DescribeInstanceStatus", param0)
	return m.DescribeInstanceStatusFunc(param0)
}

func (m *ec2Mock) DescribeInstanceStatusRequest(param0 *ec2.DescribeInstanceStatusInput) (*request.Request, *ec2.DescribeInstanceStatusOutput) {
	m.addCall("DescribeInstanceStatusRequest")
	m.verifyInput("DescribeInstanceStatusRequest", param0)
	return m.DescribeInstanceStatusRequestFunc(param0)
}

func (m *ec2Mock) DescribeInstanceStatusWithContext(param0 aws.Context, param1 *ec2.DescribeInstanceStatusInput, param2 ...request.Option) (*ec2.DescribeInstanceStatusOutput, error) {
	m.addCall("DescribeInstanceStatusWithContext")
	m.verifyInput("DescribeInstanceStatusWithContext", param0)
	return m.DescribeInstanceStatusWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeInstances(param0 *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	m.addCall("DescribeInstances")
	m.verifyInput("DescribeInstances", param0)
	return m.DescribeInstancesFunc(param0)
}

func (m *ec2Mock) DescribeInstancesRequest(param0 *ec2.DescribeInstancesInput) (*request.Request, *ec2.DescribeInstancesOutput) {
	m.addCall("DescribeInstancesRequest")
	m.verifyInput("DescribeInstancesRequest", param0)
	return m.DescribeInstancesRequestFunc(param0)
}

func (m *ec2Mock) DescribeInstancesWithContext(param0 aws.Context, param1 *ec2.DescribeInstancesInput, param2 ...request.Option) (*ec2.DescribeInstancesOutput, error) {
	m.addCall("DescribeInstancesWithContext")
	m.verifyInput("DescribeInstancesWithContext", param0)
	return m.DescribeInstancesWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeInternetGateways(param0 *ec2.DescribeInternetGatewaysInput) (*ec2.DescribeInternetGatewaysOutput, error) {
	m.addCall("DescribeInternetGateways")
	m.verifyInput("DescribeInternetGateways", param0)
	return m.DescribeInternetGatewaysFunc(param0)
}

func (m *ec2Mock) DescribeInternetGatewaysRequest(param0 *ec2.DescribeInternetGatewaysInput) (*request.Request, *ec2.DescribeInternetGatewaysOutput) {
	m.addCall("DescribeInternetGatewaysRequest")
	m.verifyInput("DescribeInternetGatewaysRequest", param0)
	return m.DescribeInternetGatewaysRequestFunc(param0)
}

func (m *ec2Mock) DescribeInternetGatewaysWithContext(param0 aws.Context, param1 *ec2.DescribeInternetGatewaysInput, param2 ...request.Option) (*ec2.DescribeInternetGatewaysOutput, error) {
	m.addCall("DescribeInternetGatewaysWithContext")
	m.verifyInput("DescribeInternetGatewaysWithContext", param0)
	return m.DescribeInternetGatewaysWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeKeyPairs(param0 *ec2.DescribeKeyPairsInput) (*ec2.DescribeKeyPairsOutput, error) {
	m.addCall("DescribeKeyPairs")
	m.verifyInput("DescribeKeyPairs", param0)
	return m.DescribeKeyPairsFunc(param0)
}

func (m *ec2Mock) DescribeKeyPairsRequest(param0 *ec2.DescribeKeyPairsInput) (*request.Request, *ec2.DescribeKeyPairsOutput) {
	m.addCall("DescribeKeyPairsRequest")
	m.verifyInput("DescribeKeyPairsRequest", param0)
	return m.DescribeKeyPairsRequestFunc(param0)
}

func (m *ec2Mock) DescribeKeyPairsWithContext(param0 aws.Context, param1 *ec2.DescribeKeyPairsInput, param2 ...request.Option) (*ec2.DescribeKeyPairsOutput, error) {
	m.addCall("DescribeKeyPairsWithContext")
	m.verifyInput("DescribeKeyPairsWithContext", param0)
	return m.DescribeKeyPairsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeMovingAddresses(param0 *ec2.DescribeMovingAddressesInput) (*ec2.DescribeMovingAddressesOutput, error) {
	m.addCall("DescribeMovingAddresses")
	m.verifyInput("DescribeMovingAddresses", param0)
	return m.DescribeMovingAddressesFunc(param0)
}

func (m *ec2Mock) DescribeMovingAddressesRequest(param0 *ec2.DescribeMovingAddressesInput) (*request.Request, *ec2.DescribeMovingAddressesOutput) {
	m.addCall("DescribeMovingAddressesRequest")
	m.verifyInput("DescribeMovingAddressesRequest", param0)
	return m.DescribeMovingAddressesRequestFunc(param0)
}

func (m *ec2Mock) DescribeMovingAddressesWithContext(param0 aws.Context, param1 *ec2.DescribeMovingAddressesInput, param2 ...request.Option) (*ec2.DescribeMovingAddressesOutput, error) {
	m.addCall("DescribeMovingAddressesWithContext")
	m.verifyInput("DescribeMovingAddressesWithContext", param0)
	return m.DescribeMovingAddressesWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeNatGateways(param0 *ec2.DescribeNatGatewaysInput) (*ec2.DescribeNatGatewaysOutput, error) {
	m.addCall("DescribeNatGateways")
	m.verifyInput("DescribeNatGateways", param0)
	return m.DescribeNatGatewaysFunc(param0)
}

func (m *ec2Mock) DescribeNatGatewaysRequest(param0 *ec2.DescribeNatGatewaysInput) (*request.Request, *ec2.DescribeNatGatewaysOutput) {
	m.addCall("DescribeNatGatewaysRequest")
	m.verifyInput("DescribeNatGatewaysRequest", param0)
	return m.DescribeNatGatewaysRequestFunc(param0)
}

func (m *ec2Mock) DescribeNatGatewaysWithContext(param0 aws.Context, param1 *ec2.DescribeNatGatewaysInput, param2 ...request.Option) (*ec2.DescribeNatGatewaysOutput, error) {
	m.addCall("DescribeNatGatewaysWithContext")
	m.verifyInput("DescribeNatGatewaysWithContext", param0)
	return m.DescribeNatGatewaysWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeNetworkAcls(param0 *ec2.DescribeNetworkAclsInput) (*ec2.DescribeNetworkAclsOutput, error) {
	m.addCall("DescribeNetworkAcls")
	m.verifyInput("DescribeNetworkAcls", param0)
	return m.DescribeNetworkAclsFunc(param0)
}

func (m *ec2Mock) DescribeNetworkAclsRequest(param0 *ec2.DescribeNetworkAclsInput) (*request.Request, *ec2.DescribeNetworkAclsOutput) {
	m.addCall("DescribeNetworkAclsRequest")
	m.verifyInput("DescribeNetworkAclsRequest", param0)
	return m.DescribeNetworkAclsRequestFunc(param0)
}

func (m *ec2Mock) DescribeNetworkAclsWithContext(param0 aws.Context, param1 *ec2.DescribeNetworkAclsInput, param2 ...request.Option) (*ec2.DescribeNetworkAclsOutput, error) {
	m.addCall("DescribeNetworkAclsWithContext")
	m.verifyInput("DescribeNetworkAclsWithContext", param0)
	return m.DescribeNetworkAclsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeNetworkInterfaceAttribute(param0 *ec2.DescribeNetworkInterfaceAttributeInput) (*ec2.DescribeNetworkInterfaceAttributeOutput, error) {
	m.addCall("DescribeNetworkInterfaceAttribute")
	m.verifyInput("DescribeNetworkInterfaceAttribute", param0)
	return m.DescribeNetworkInterfaceAttributeFunc(param0)
}

func (m *ec2Mock) DescribeNetworkInterfaceAttributeRequest(param0 *ec2.DescribeNetworkInterfaceAttributeInput) (*request.Request, *ec2.DescribeNetworkInterfaceAttributeOutput) {
	m.addCall("DescribeNetworkInterfaceAttributeRequest")
	m.verifyInput("DescribeNetworkInterfaceAttributeRequest", param0)
	return m.DescribeNetworkInterfaceAttributeRequestFunc(param0)
}

func (m *ec2Mock) DescribeNetworkInterfaceAttributeWithContext(param0 aws.Context, param1 *ec2.DescribeNetworkInterfaceAttributeInput, param2 ...request.Option) (*ec2.DescribeNetworkInterfaceAttributeOutput, error) {
	m.addCall("DescribeNetworkInterfaceAttributeWithContext")
	m.verifyInput("DescribeNetworkInterfaceAttributeWithContext", param0)
	return m.DescribeNetworkInterfaceAttributeWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeNetworkInterfacePermissions(param0 *ec2.DescribeNetworkInterfacePermissionsInput) (*ec2.DescribeNetworkInterfacePermissionsOutput, error) {
	m.addCall("DescribeNetworkInterfacePermissions")
	m.verifyInput("DescribeNetworkInterfacePermissions", param0)
	return m.DescribeNetworkInterfacePermissionsFunc(param0)
}

func (m *ec2Mock) DescribeNetworkInterfacePermissionsRequest(param0 *ec2.DescribeNetworkInterfacePermissionsInput) (*request.Request, *ec2.DescribeNetworkInterfacePermissionsOutput) {
	m.addCall("DescribeNetworkInterfacePermissionsRequest")
	m.verifyInput("DescribeNetworkInterfacePermissionsRequest", param0)
	return m.DescribeNetworkInterfacePermissionsRequestFunc(param0)
}

func (m *ec2Mock) DescribeNetworkInterfacePermissionsWithContext(param0 aws.Context, param1 *ec2.DescribeNetworkInterfacePermissionsInput, param2 ...request.Option) (*ec2.DescribeNetworkInterfacePermissionsOutput, error) {
	m.addCall("DescribeNetworkInterfacePermissionsWithContext")
	m.verifyInput("DescribeNetworkInterfacePermissionsWithContext", param0)
	return m.DescribeNetworkInterfacePermissionsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeNetworkInterfaces(param0 *ec2.DescribeNetworkInterfacesInput) (*ec2.DescribeNetworkInterfacesOutput, error) {
	m.addCall("DescribeNetworkInterfaces")
	m.verifyInput("DescribeNetworkInterfaces", param0)
	return m.DescribeNetworkInterfacesFunc(param0)
}

func (m *ec2Mock) DescribeNetworkInterfacesRequest(param0 *ec2.DescribeNetworkInterfacesInput) (*request.Request, *ec2.DescribeNetworkInterfacesOutput) {
	m.addCall("DescribeNetworkInterfacesRequest")
	m.verifyInput("DescribeNetworkInterfacesRequest", param0)
	return m.DescribeNetworkInterfacesRequestFunc(param0)
}

func (m *ec2Mock) DescribeNetworkInterfacesWithContext(param0 aws.Context, param1 *ec2.DescribeNetworkInterfacesInput, param2 ...request.Option) (*ec2.DescribeNetworkInterfacesOutput, error) {
	m.addCall("DescribeNetworkInterfacesWithContext")
	m.verifyInput("DescribeNetworkInterfacesWithContext", param0)
	return m.DescribeNetworkInterfacesWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribePlacementGroups(param0 *ec2.DescribePlacementGroupsInput) (*ec2.DescribePlacementGroupsOutput, error) {
	m.addCall("DescribePlacementGroups")
	m.verifyInput("DescribePlacementGroups", param0)
	return m.DescribePlacementGroupsFunc(param0)
}

func (m *ec2Mock) DescribePlacementGroupsRequest(param0 *ec2.DescribePlacementGroupsInput) (*request.Request, *ec2.DescribePlacementGroupsOutput) {
	m.addCall("DescribePlacementGroupsRequest")
	m.verifyInput("DescribePlacementGroupsRequest", param0)
	return m.DescribePlacementGroupsRequestFunc(param0)
}

func (m *ec2Mock) DescribePlacementGroupsWithContext(param0 aws.Context, param1 *ec2.DescribePlacementGroupsInput, param2 ...request.Option) (*ec2.DescribePlacementGroupsOutput, error) {
	m.addCall("DescribePlacementGroupsWithContext")
	m.verifyInput("DescribePlacementGroupsWithContext", param0)
	return m.DescribePlacementGroupsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribePrefixLists(param0 *ec2.DescribePrefixListsInput) (*ec2.DescribePrefixListsOutput, error) {
	m.addCall("DescribePrefixLists")
	m.verifyInput("DescribePrefixLists", param0)
	return m.DescribePrefixListsFunc(param0)
}

func (m *ec2Mock) DescribePrefixListsRequest(param0 *ec2.DescribePrefixListsInput) (*request.Request, *ec2.DescribePrefixListsOutput) {
	m.addCall("DescribePrefixListsRequest")
	m.verifyInput("DescribePrefixListsRequest", param0)
	return m.DescribePrefixListsRequestFunc(param0)
}

func (m *ec2Mock) DescribePrefixListsWithContext(param0 aws.Context, param1 *ec2.DescribePrefixListsInput, param2 ...request.Option) (*ec2.DescribePrefixListsOutput, error) {
	m.addCall("DescribePrefixListsWithContext")
	m.verifyInput("DescribePrefixListsWithContext", param0)
	return m.DescribePrefixListsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeRegions(param0 *ec2.DescribeRegionsInput) (*ec2.DescribeRegionsOutput, error) {
	m.addCall("DescribeRegions")
	m.verifyInput("DescribeRegions", param0)
	return m.DescribeRegionsFunc(param0)
}

func (m *ec2Mock) DescribeRegionsRequest(param0 *ec2.DescribeRegionsInput) (*request.Request, *ec2.DescribeRegionsOutput) {
	m.addCall("DescribeRegionsRequest")
	m.verifyInput("DescribeRegionsRequest", param0)
	return m.DescribeRegionsRequestFunc(param0)
}

func (m *ec2Mock) DescribeRegionsWithContext(param0 aws.Context, param1 *ec2.DescribeRegionsInput, param2 ...request.Option) (*ec2.DescribeRegionsOutput, error) {
	m.addCall("DescribeRegionsWithContext")
	m.verifyInput("DescribeRegionsWithContext", param0)
	return m.DescribeRegionsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeReservedInstances(param0 *ec2.DescribeReservedInstancesInput) (*ec2.DescribeReservedInstancesOutput, error) {
	m.addCall("DescribeReservedInstances")
	m.verifyInput("DescribeReservedInstances", param0)
	return m.DescribeReservedInstancesFunc(param0)
}

func (m *ec2Mock) DescribeReservedInstancesListings(param0 *ec2.DescribeReservedInstancesListingsInput) (*ec2.DescribeReservedInstancesListingsOutput, error) {
	m.addCall("DescribeReservedInstancesListings")
	m.verifyInput("DescribeReservedInstancesListings", param0)
	return m.DescribeReservedInstancesListingsFunc(param0)
}

func (m *ec2Mock) DescribeReservedInstancesListingsRequest(param0 *ec2.DescribeReservedInstancesListingsInput) (*request.Request, *ec2.DescribeReservedInstancesListingsOutput) {
	m.addCall("DescribeReservedInstancesListingsRequest")
	m.verifyInput("DescribeReservedInstancesListingsRequest", param0)
	return m.DescribeReservedInstancesListingsRequestFunc(param0)
}

func (m *ec2Mock) DescribeReservedInstancesListingsWithContext(param0 aws.Context, param1 *ec2.DescribeReservedInstancesListingsInput, param2 ...request.Option) (*ec2.DescribeReservedInstancesListingsOutput, error) {
	m.addCall("DescribeReservedInstancesListingsWithContext")
	m.verifyInput("DescribeReservedInstancesListingsWithContext", param0)
	return m.DescribeReservedInstancesListingsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeReservedInstancesModifications(param0 *ec2.DescribeReservedInstancesModificationsInput) (*ec2.DescribeReservedInstancesModificationsOutput, error) {
	m.addCall("DescribeReservedInstancesModifications")
	m.verifyInput("DescribeReservedInstancesModifications", param0)
	return m.DescribeReservedInstancesModificationsFunc(param0)
}

func (m *ec2Mock) DescribeReservedInstancesModificationsRequest(param0 *ec2.DescribeReservedInstancesModificationsInput) (*request.Request, *ec2.DescribeReservedInstancesModificationsOutput) {
	m.addCall("DescribeReservedInstancesModificationsRequest")
	m.verifyInput("DescribeReservedInstancesModificationsRequest", param0)
	return m.DescribeReservedInstancesModificationsRequestFunc(param0)
}

func (m *ec2Mock) DescribeReservedInstancesModificationsWithContext(param0 aws.Context, param1 *ec2.DescribeReservedInstancesModificationsInput, param2 ...request.Option) (*ec2.DescribeReservedInstancesModificationsOutput, error) {
	m.addCall("DescribeReservedInstancesModificationsWithContext")
	m.verifyInput("DescribeReservedInstancesModificationsWithContext", param0)
	return m.DescribeReservedInstancesModificationsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeReservedInstancesOfferings(param0 *ec2.DescribeReservedInstancesOfferingsInput) (*ec2.DescribeReservedInstancesOfferingsOutput, error) {
	m.addCall("DescribeReservedInstancesOfferings")
	m.verifyInput("DescribeReservedInstancesOfferings", param0)
	return m.DescribeReservedInstancesOfferingsFunc(param0)
}

func (m *ec2Mock) DescribeReservedInstancesOfferingsRequest(param0 *ec2.DescribeReservedInstancesOfferingsInput) (*request.Request, *ec2.DescribeReservedInstancesOfferingsOutput) {
	m.addCall("DescribeReservedInstancesOfferingsRequest")
	m.verifyInput("DescribeReservedInstancesOfferingsRequest", param0)
	return m.DescribeReservedInstancesOfferingsRequestFunc(param0)
}

func (m *ec2Mock) DescribeReservedInstancesOfferingsWithContext(param0 aws.Context, param1 *ec2.DescribeReservedInstancesOfferingsInput, param2 ...request.Option) (*ec2.DescribeReservedInstancesOfferingsOutput, error) {
	m.addCall("DescribeReservedInstancesOfferingsWithContext")
	m.verifyInput("DescribeReservedInstancesOfferingsWithContext", param0)
	return m.DescribeReservedInstancesOfferingsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeReservedInstancesRequest(param0 *ec2.DescribeReservedInstancesInput) (*request.Request, *ec2.DescribeReservedInstancesOutput) {
	m.addCall("DescribeReservedInstancesRequest")
	m.verifyInput("DescribeReservedInstancesRequest", param0)
	return m.DescribeReservedInstancesRequestFunc(param0)
}

func (m *ec2Mock) DescribeReservedInstancesWithContext(param0 aws.Context, param1 *ec2.DescribeReservedInstancesInput, param2 ...request.Option) (*ec2.DescribeReservedInstancesOutput, error) {
	m.addCall("DescribeReservedInstancesWithContext")
	m.verifyInput("DescribeReservedInstancesWithContext", param0)
	return m.DescribeReservedInstancesWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeRouteTables(param0 *ec2.DescribeRouteTablesInput) (*ec2.DescribeRouteTablesOutput, error) {
	m.addCall("DescribeRouteTables")
	m.verifyInput("DescribeRouteTables", param0)
	return m.DescribeRouteTablesFunc(param0)
}

func (m *ec2Mock) DescribeRouteTablesRequest(param0 *ec2.DescribeRouteTablesInput) (*request.Request, *ec2.DescribeRouteTablesOutput) {
	m.addCall("DescribeRouteTablesRequest")
	m.verifyInput("DescribeRouteTablesRequest", param0)
	return m.DescribeRouteTablesRequestFunc(param0)
}

func (m *ec2Mock) DescribeRouteTablesWithContext(param0 aws.Context, param1 *ec2.DescribeRouteTablesInput, param2 ...request.Option) (*ec2.DescribeRouteTablesOutput, error) {
	m.addCall("DescribeRouteTablesWithContext")
	m.verifyInput("DescribeRouteTablesWithContext", param0)
	return m.DescribeRouteTablesWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeScheduledInstanceAvailability(param0 *ec2.DescribeScheduledInstanceAvailabilityInput) (*ec2.DescribeScheduledInstanceAvailabilityOutput, error) {
	m.addCall("DescribeScheduledInstanceAvailability")
	m.verifyInput("DescribeScheduledInstanceAvailability", param0)
	return m.DescribeScheduledInstanceAvailabilityFunc(param0)
}

func (m *ec2Mock) DescribeScheduledInstanceAvailabilityRequest(param0 *ec2.DescribeScheduledInstanceAvailabilityInput) (*request.Request, *ec2.DescribeScheduledInstanceAvailabilityOutput) {
	m.addCall("DescribeScheduledInstanceAvailabilityRequest")
	m.verifyInput("DescribeScheduledInstanceAvailabilityRequest", param0)
	return m.DescribeScheduledInstanceAvailabilityRequestFunc(param0)
}

func (m *ec2Mock) DescribeScheduledInstanceAvailabilityWithContext(param0 aws.Context, param1 *ec2.DescribeScheduledInstanceAvailabilityInput, param2 ...request.Option) (*ec2.DescribeScheduledInstanceAvailabilityOutput, error) {
	m.addCall("DescribeScheduledInstanceAvailabilityWithContext")
	m.verifyInput("DescribeScheduledInstanceAvailabilityWithContext", param0)
	return m.DescribeScheduledInstanceAvailabilityWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeScheduledInstances(param0 *ec2.DescribeScheduledInstancesInput) (*ec2.DescribeScheduledInstancesOutput, error) {
	m.addCall("DescribeScheduledInstances")
	m.verifyInput("DescribeScheduledInstances", param0)
	return m.DescribeScheduledInstancesFunc(param0)
}

func (m *ec2Mock) DescribeScheduledInstancesRequest(param0 *ec2.DescribeScheduledInstancesInput) (*request.Request, *ec2.DescribeScheduledInstancesOutput) {
	m.addCall("DescribeScheduledInstancesRequest")
	m.verifyInput("DescribeScheduledInstancesRequest", param0)
	return m.DescribeScheduledInstancesRequestFunc(param0)
}

func (m *ec2Mock) DescribeScheduledInstancesWithContext(param0 aws.Context, param1 *ec2.DescribeScheduledInstancesInput, param2 ...request.Option) (*ec2.DescribeScheduledInstancesOutput, error) {
	m.addCall("DescribeScheduledInstancesWithContext")
	m.verifyInput("DescribeScheduledInstancesWithContext", param0)
	return m.DescribeScheduledInstancesWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeSecurityGroupReferences(param0 *ec2.DescribeSecurityGroupReferencesInput) (*ec2.DescribeSecurityGroupReferencesOutput, error) {
	m.addCall("DescribeSecurityGroupReferences")
	m.verifyInput("DescribeSecurityGroupReferences", param0)
	return m.DescribeSecurityGroupReferencesFunc(param0)
}

func (m *ec2Mock) DescribeSecurityGroupReferencesRequest(param0 *ec2.DescribeSecurityGroupReferencesInput) (*request.Request, *ec2.DescribeSecurityGroupReferencesOutput) {
	m.addCall("DescribeSecurityGroupReferencesRequest")
	m.verifyInput("DescribeSecurityGroupReferencesRequest", param0)
	return m.DescribeSecurityGroupReferencesRequestFunc(param0)
}

func (m *ec2Mock) DescribeSecurityGroupReferencesWithContext(param0 aws.Context, param1 *ec2.DescribeSecurityGroupReferencesInput, param2 ...request.Option) (*ec2.DescribeSecurityGroupReferencesOutput, error) {
	m.addCall("DescribeSecurityGroupReferencesWithContext")
	m.verifyInput("DescribeSecurityGroupReferencesWithContext", param0)
	return m.DescribeSecurityGroupReferencesWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeSecurityGroups(param0 *ec2.DescribeSecurityGroupsInput) (*ec2.DescribeSecurityGroupsOutput, error) {
	m.addCall("DescribeSecurityGroups")
	m.verifyInput("DescribeSecurityGroups", param0)
	return m.DescribeSecurityGroupsFunc(param0)
}

func (m *ec2Mock) DescribeSecurityGroupsRequest(param0 *ec2.DescribeSecurityGroupsInput) (*request.Request, *ec2.DescribeSecurityGroupsOutput) {
	m.addCall("DescribeSecurityGroupsRequest")
	m.verifyInput("DescribeSecurityGroupsRequest", param0)
	return m.DescribeSecurityGroupsRequestFunc(param0)
}

func (m *ec2Mock) DescribeSecurityGroupsWithContext(param0 aws.Context, param1 *ec2.DescribeSecurityGroupsInput, param2 ...request.Option) (*ec2.DescribeSecurityGroupsOutput, error) {
	m.addCall("DescribeSecurityGroupsWithContext")
	m.verifyInput("DescribeSecurityGroupsWithContext", param0)
	return m.DescribeSecurityGroupsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeSnapshotAttribute(param0 *ec2.DescribeSnapshotAttributeInput) (*ec2.DescribeSnapshotAttributeOutput, error) {
	m.addCall("DescribeSnapshotAttribute")
	m.verifyInput("DescribeSnapshotAttribute", param0)
	return m.DescribeSnapshotAttributeFunc(param0)
}

func (m *ec2Mock) DescribeSnapshotAttributeRequest(param0 *ec2.DescribeSnapshotAttributeInput) (*request.Request, *ec2.DescribeSnapshotAttributeOutput) {
	m.addCall("DescribeSnapshotAttributeRequest")
	m.verifyInput("DescribeSnapshotAttributeRequest", param0)
	return m.DescribeSnapshotAttributeRequestFunc(param0)
}

func (m *ec2Mock) DescribeSnapshotAttributeWithContext(param0 aws.Context, param1 *ec2.DescribeSnapshotAttributeInput, param2 ...request.Option) (*ec2.DescribeSnapshotAttributeOutput, error) {
	m.addCall("DescribeSnapshotAttributeWithContext")
	m.verifyInput("DescribeSnapshotAttributeWithContext", param0)
	return m.DescribeSnapshotAttributeWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeSnapshots(param0 *ec2.DescribeSnapshotsInput) (*ec2.DescribeSnapshotsOutput, error) {
	m.addCall("DescribeSnapshots")
	m.verifyInput("DescribeSnapshots", param0)
	return m.DescribeSnapshotsFunc(param0)
}

func (m *ec2Mock) DescribeSnapshotsRequest(param0 *ec2.DescribeSnapshotsInput) (*request.Request, *ec2.DescribeSnapshotsOutput) {
	m.addCall("DescribeSnapshotsRequest")
	m.verifyInput("DescribeSnapshotsRequest", param0)
	return m.DescribeSnapshotsRequestFunc(param0)
}

func (m *ec2Mock) DescribeSnapshotsWithContext(param0 aws.Context, param1 *ec2.DescribeSnapshotsInput, param2 ...request.Option) (*ec2.DescribeSnapshotsOutput, error) {
	m.addCall("DescribeSnapshotsWithContext")
	m.verifyInput("DescribeSnapshotsWithContext", param0)
	return m.DescribeSnapshotsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeSpotDatafeedSubscription(param0 *ec2.DescribeSpotDatafeedSubscriptionInput) (*ec2.DescribeSpotDatafeedSubscriptionOutput, error) {
	m.addCall("DescribeSpotDatafeedSubscription")
	m.verifyInput("DescribeSpotDatafeedSubscription", param0)
	return m.DescribeSpotDatafeedSubscriptionFunc(param0)
}

func (m *ec2Mock) DescribeSpotDatafeedSubscriptionRequest(param0 *ec2.DescribeSpotDatafeedSubscriptionInput) (*request.Request, *ec2.DescribeSpotDatafeedSubscriptionOutput) {
	m.addCall("DescribeSpotDatafeedSubscriptionRequest")
	m.verifyInput("DescribeSpotDatafeedSubscriptionRequest", param0)
	return m.DescribeSpotDatafeedSubscriptionRequestFunc(param0)
}

func (m *ec2Mock) DescribeSpotDatafeedSubscriptionWithContext(param0 aws.Context, param1 *ec2.DescribeSpotDatafeedSubscriptionInput, param2 ...request.Option) (*ec2.DescribeSpotDatafeedSubscriptionOutput, error) {
	m.addCall("DescribeSpotDatafeedSubscriptionWithContext")
	m.verifyInput("DescribeSpotDatafeedSubscriptionWithContext", param0)
	return m.DescribeSpotDatafeedSubscriptionWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeSpotFleetInstances(param0 *ec2.DescribeSpotFleetInstancesInput) (*ec2.DescribeSpotFleetInstancesOutput, error) {
	m.addCall("DescribeSpotFleetInstances")
	m.verifyInput("DescribeSpotFleetInstances", param0)
	return m.DescribeSpotFleetInstancesFunc(param0)
}

func (m *ec2Mock) DescribeSpotFleetInstancesRequest(param0 *ec2.DescribeSpotFleetInstancesInput) (*request.Request, *ec2.DescribeSpotFleetInstancesOutput) {
	m.addCall("DescribeSpotFleetInstancesRequest")
	m.verifyInput("DescribeSpotFleetInstancesRequest", param0)
	return m.DescribeSpotFleetInstancesRequestFunc(param0)
}

func (m *ec2Mock) DescribeSpotFleetInstancesWithContext(param0 aws.Context, param1 *ec2.DescribeSpotFleetInstancesInput, param2 ...request.Option) (*ec2.DescribeSpotFleetInstancesOutput, error) {
	m.addCall("DescribeSpotFleetInstancesWithContext")
	m.verifyInput("DescribeSpotFleetInstancesWithContext", param0)
	return m.DescribeSpotFleetInstancesWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeSpotFleetRequestHistory(param0 *ec2.DescribeSpotFleetRequestHistoryInput) (*ec2.DescribeSpotFleetRequestHistoryOutput, error) {
	m.addCall("DescribeSpotFleetRequestHistory")
	m.verifyInput("DescribeSpotFleetRequestHistory", param0)
	return m.DescribeSpotFleetRequestHistoryFunc(param0)
}

func (m *ec2Mock) DescribeSpotFleetRequestHistoryRequest(param0 *ec2.DescribeSpotFleetRequestHistoryInput) (*request.Request, *ec2.DescribeSpotFleetRequestHistoryOutput) {
	m.addCall("DescribeSpotFleetRequestHistoryRequest")
	m.verifyInput("DescribeSpotFleetRequestHistoryRequest", param0)
	return m.DescribeSpotFleetRequestHistoryRequestFunc(param0)
}

func (m *ec2Mock) DescribeSpotFleetRequestHistoryWithContext(param0 aws.Context, param1 *ec2.DescribeSpotFleetRequestHistoryInput, param2 ...request.Option) (*ec2.DescribeSpotFleetRequestHistoryOutput, error) {
	m.addCall("DescribeSpotFleetRequestHistoryWithContext")
	m.verifyInput("DescribeSpotFleetRequestHistoryWithContext", param0)
	return m.DescribeSpotFleetRequestHistoryWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeSpotFleetRequests(param0 *ec2.DescribeSpotFleetRequestsInput) (*ec2.DescribeSpotFleetRequestsOutput, error) {
	m.addCall("DescribeSpotFleetRequests")
	m.verifyInput("DescribeSpotFleetRequests", param0)
	return m.DescribeSpotFleetRequestsFunc(param0)
}

func (m *ec2Mock) DescribeSpotFleetRequestsRequest(param0 *ec2.DescribeSpotFleetRequestsInput) (*request.Request, *ec2.DescribeSpotFleetRequestsOutput) {
	m.addCall("DescribeSpotFleetRequestsRequest")
	m.verifyInput("DescribeSpotFleetRequestsRequest", param0)
	return m.DescribeSpotFleetRequestsRequestFunc(param0)
}

func (m *ec2Mock) DescribeSpotFleetRequestsWithContext(param0 aws.Context, param1 *ec2.DescribeSpotFleetRequestsInput, param2 ...request.Option) (*ec2.DescribeSpotFleetRequestsOutput, error) {
	m.addCall("DescribeSpotFleetRequestsWithContext")
	m.verifyInput("DescribeSpotFleetRequestsWithContext", param0)
	return m.DescribeSpotFleetRequestsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeSpotInstanceRequests(param0 *ec2.DescribeSpotInstanceRequestsInput) (*ec2.DescribeSpotInstanceRequestsOutput, error) {
	m.addCall("DescribeSpotInstanceRequests")
	m.verifyInput("DescribeSpotInstanceRequests", param0)
	return m.DescribeSpotInstanceRequestsFunc(param0)
}

func (m *ec2Mock) DescribeSpotInstanceRequestsRequest(param0 *ec2.DescribeSpotInstanceRequestsInput) (*request.Request, *ec2.DescribeSpotInstanceRequestsOutput) {
	m.addCall("DescribeSpotInstanceRequestsRequest")
	m.verifyInput("DescribeSpotInstanceRequestsRequest", param0)
	return m.DescribeSpotInstanceRequestsRequestFunc(param0)
}

func (m *ec2Mock) DescribeSpotInstanceRequestsWithContext(param0 aws.Context, param1 *ec2.DescribeSpotInstanceRequestsInput, param2 ...request.Option) (*ec2.DescribeSpotInstanceRequestsOutput, error) {
	m.addCall("DescribeSpotInstanceRequestsWithContext")
	m.verifyInput("DescribeSpotInstanceRequestsWithContext", param0)
	return m.DescribeSpotInstanceRequestsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeSpotPriceHistory(param0 *ec2.DescribeSpotPriceHistoryInput) (*ec2.DescribeSpotPriceHistoryOutput, error) {
	m.addCall("DescribeSpotPriceHistory")
	m.verifyInput("DescribeSpotPriceHistory", param0)
	return m.DescribeSpotPriceHistoryFunc(param0)
}

func (m *ec2Mock) DescribeSpotPriceHistoryRequest(param0 *ec2.DescribeSpotPriceHistoryInput) (*request.Request, *ec2.DescribeSpotPriceHistoryOutput) {
	m.addCall("DescribeSpotPriceHistoryRequest")
	m.verifyInput("DescribeSpotPriceHistoryRequest", param0)
	return m.DescribeSpotPriceHistoryRequestFunc(param0)
}

func (m *ec2Mock) DescribeSpotPriceHistoryWithContext(param0 aws.Context, param1 *ec2.DescribeSpotPriceHistoryInput, param2 ...request.Option) (*ec2.DescribeSpotPriceHistoryOutput, error) {
	m.addCall("DescribeSpotPriceHistoryWithContext")
	m.verifyInput("DescribeSpotPriceHistoryWithContext", param0)
	return m.DescribeSpotPriceHistoryWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeStaleSecurityGroups(param0 *ec2.DescribeStaleSecurityGroupsInput) (*ec2.DescribeStaleSecurityGroupsOutput, error) {
	m.addCall("DescribeStaleSecurityGroups")
	m.verifyInput("DescribeStaleSecurityGroups", param0)
	return m.DescribeStaleSecurityGroupsFunc(param0)
}

func (m *ec2Mock) DescribeStaleSecurityGroupsRequest(param0 *ec2.DescribeStaleSecurityGroupsInput) (*request.Request, *ec2.DescribeStaleSecurityGroupsOutput) {
	m.addCall("DescribeStaleSecurityGroupsRequest")
	m.verifyInput("DescribeStaleSecurityGroupsRequest", param0)
	return m.DescribeStaleSecurityGroupsRequestFunc(param0)
}

func (m *ec2Mock) DescribeStaleSecurityGroupsWithContext(param0 aws.Context, param1 *ec2.DescribeStaleSecurityGroupsInput, param2 ...request.Option) (*ec2.DescribeStaleSecurityGroupsOutput, error) {
	m.addCall("DescribeStaleSecurityGroupsWithContext")
	m.verifyInput("DescribeStaleSecurityGroupsWithContext", param0)
	return m.DescribeStaleSecurityGroupsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeSubnets(param0 *ec2.DescribeSubnetsInput) (*ec2.DescribeSubnetsOutput, error) {
	m.addCall("DescribeSubnets")
	m.verifyInput("DescribeSubnets", param0)
	return m.DescribeSubnetsFunc(param0)
}

func (m *ec2Mock) DescribeSubnetsRequest(param0 *ec2.DescribeSubnetsInput) (*request.Request, *ec2.DescribeSubnetsOutput) {
	m.addCall("DescribeSubnetsRequest")
	m.verifyInput("DescribeSubnetsRequest", param0)
	return m.DescribeSubnetsRequestFunc(param0)
}

func (m *ec2Mock) DescribeSubnetsWithContext(param0 aws.Context, param1 *ec2.DescribeSubnetsInput, param2 ...request.Option) (*ec2.DescribeSubnetsOutput, error) {
	m.addCall("DescribeSubnetsWithContext")
	m.verifyInput("DescribeSubnetsWithContext", param0)
	return m.DescribeSubnetsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeTags(param0 *ec2.DescribeTagsInput) (*ec2.DescribeTagsOutput, error) {
	m.addCall("DescribeTags")
	m.verifyInput("DescribeTags", param0)
	return m.DescribeTagsFunc(param0)
}

func (m *ec2Mock) DescribeTagsRequest(param0 *ec2.DescribeTagsInput) (*request.Request, *ec2.DescribeTagsOutput) {
	m.addCall("DescribeTagsRequest")
	m.verifyInput("DescribeTagsRequest", param0)
	return m.DescribeTagsRequestFunc(param0)
}

func (m *ec2Mock) DescribeTagsWithContext(param0 aws.Context, param1 *ec2.DescribeTagsInput, param2 ...request.Option) (*ec2.DescribeTagsOutput, error) {
	m.addCall("DescribeTagsWithContext")
	m.verifyInput("DescribeTagsWithContext", param0)
	return m.DescribeTagsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeVolumeAttribute(param0 *ec2.DescribeVolumeAttributeInput) (*ec2.DescribeVolumeAttributeOutput, error) {
	m.addCall("DescribeVolumeAttribute")
	m.verifyInput("DescribeVolumeAttribute", param0)
	return m.DescribeVolumeAttributeFunc(param0)
}

func (m *ec2Mock) DescribeVolumeAttributeRequest(param0 *ec2.DescribeVolumeAttributeInput) (*request.Request, *ec2.DescribeVolumeAttributeOutput) {
	m.addCall("DescribeVolumeAttributeRequest")
	m.verifyInput("DescribeVolumeAttributeRequest", param0)
	return m.DescribeVolumeAttributeRequestFunc(param0)
}

func (m *ec2Mock) DescribeVolumeAttributeWithContext(param0 aws.Context, param1 *ec2.DescribeVolumeAttributeInput, param2 ...request.Option) (*ec2.DescribeVolumeAttributeOutput, error) {
	m.addCall("DescribeVolumeAttributeWithContext")
	m.verifyInput("DescribeVolumeAttributeWithContext", param0)
	return m.DescribeVolumeAttributeWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeVolumeStatus(param0 *ec2.DescribeVolumeStatusInput) (*ec2.DescribeVolumeStatusOutput, error) {
	m.addCall("DescribeVolumeStatus")
	m.verifyInput("DescribeVolumeStatus", param0)
	return m.DescribeVolumeStatusFunc(param0)
}

func (m *ec2Mock) DescribeVolumeStatusRequest(param0 *ec2.DescribeVolumeStatusInput) (*request.Request, *ec2.DescribeVolumeStatusOutput) {
	m.addCall("DescribeVolumeStatusRequest")
	m.verifyInput("DescribeVolumeStatusRequest", param0)
	return m.DescribeVolumeStatusRequestFunc(param0)
}

func (m *ec2Mock) DescribeVolumeStatusWithContext(param0 aws.Context, param1 *ec2.DescribeVolumeStatusInput, param2 ...request.Option) (*ec2.DescribeVolumeStatusOutput, error) {
	m.addCall("DescribeVolumeStatusWithContext")
	m.verifyInput("DescribeVolumeStatusWithContext", param0)
	return m.DescribeVolumeStatusWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeVolumes(param0 *ec2.DescribeVolumesInput) (*ec2.DescribeVolumesOutput, error) {
	m.addCall("DescribeVolumes")
	m.verifyInput("DescribeVolumes", param0)
	return m.DescribeVolumesFunc(param0)
}

func (m *ec2Mock) DescribeVolumesModifications(param0 *ec2.DescribeVolumesModificationsInput) (*ec2.DescribeVolumesModificationsOutput, error) {
	m.addCall("DescribeVolumesModifications")
	m.verifyInput("DescribeVolumesModifications", param0)
	return m.DescribeVolumesModificationsFunc(param0)
}

func (m *ec2Mock) DescribeVolumesModificationsRequest(param0 *ec2.DescribeVolumesModificationsInput) (*request.Request, *ec2.DescribeVolumesModificationsOutput) {
	m.addCall("DescribeVolumesModificationsRequest")
	m.verifyInput("DescribeVolumesModificationsRequest", param0)
	return m.DescribeVolumesModificationsRequestFunc(param0)
}

func (m *ec2Mock) DescribeVolumesModificationsWithContext(param0 aws.Context, param1 *ec2.DescribeVolumesModificationsInput, param2 ...request.Option) (*ec2.DescribeVolumesModificationsOutput, error) {
	m.addCall("DescribeVolumesModificationsWithContext")
	m.verifyInput("DescribeVolumesModificationsWithContext", param0)
	return m.DescribeVolumesModificationsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeVolumesRequest(param0 *ec2.DescribeVolumesInput) (*request.Request, *ec2.DescribeVolumesOutput) {
	m.addCall("DescribeVolumesRequest")
	m.verifyInput("DescribeVolumesRequest", param0)
	return m.DescribeVolumesRequestFunc(param0)
}

func (m *ec2Mock) DescribeVolumesWithContext(param0 aws.Context, param1 *ec2.DescribeVolumesInput, param2 ...request.Option) (*ec2.DescribeVolumesOutput, error) {
	m.addCall("DescribeVolumesWithContext")
	m.verifyInput("DescribeVolumesWithContext", param0)
	return m.DescribeVolumesWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeVpcAttribute(param0 *ec2.DescribeVpcAttributeInput) (*ec2.DescribeVpcAttributeOutput, error) {
	m.addCall("DescribeVpcAttribute")
	m.verifyInput("DescribeVpcAttribute", param0)
	return m.DescribeVpcAttributeFunc(param0)
}

func (m *ec2Mock) DescribeVpcAttributeRequest(param0 *ec2.DescribeVpcAttributeInput) (*request.Request, *ec2.DescribeVpcAttributeOutput) {
	m.addCall("DescribeVpcAttributeRequest")
	m.verifyInput("DescribeVpcAttributeRequest", param0)
	return m.DescribeVpcAttributeRequestFunc(param0)
}

func (m *ec2Mock) DescribeVpcAttributeWithContext(param0 aws.Context, param1 *ec2.DescribeVpcAttributeInput, param2 ...request.Option) (*ec2.DescribeVpcAttributeOutput, error) {
	m.addCall("DescribeVpcAttributeWithContext")
	m.verifyInput("DescribeVpcAttributeWithContext", param0)
	return m.DescribeVpcAttributeWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeVpcClassicLink(param0 *ec2.DescribeVpcClassicLinkInput) (*ec2.DescribeVpcClassicLinkOutput, error) {
	m.addCall("DescribeVpcClassicLink")
	m.verifyInput("DescribeVpcClassicLink", param0)
	return m.DescribeVpcClassicLinkFunc(param0)
}

func (m *ec2Mock) DescribeVpcClassicLinkDnsSupport(param0 *ec2.DescribeVpcClassicLinkDnsSupportInput) (*ec2.DescribeVpcClassicLinkDnsSupportOutput, error) {
	m.addCall("DescribeVpcClassicLinkDnsSupport")
	m.verifyInput("DescribeVpcClassicLinkDnsSupport", param0)
	return m.DescribeVpcClassicLinkDnsSupportFunc(param0)
}

func (m *ec2Mock) DescribeVpcClassicLinkDnsSupportRequest(param0 *ec2.DescribeVpcClassicLinkDnsSupportInput) (*request.Request, *ec2.DescribeVpcClassicLinkDnsSupportOutput) {
	m.addCall("DescribeVpcClassicLinkDnsSupportRequest")
	m.verifyInput("DescribeVpcClassicLinkDnsSupportRequest", param0)
	return m.DescribeVpcClassicLinkDnsSupportRequestFunc(param0)
}

func (m *ec2Mock) DescribeVpcClassicLinkDnsSupportWithContext(param0 aws.Context, param1 *ec2.DescribeVpcClassicLinkDnsSupportInput, param2 ...request.Option) (*ec2.DescribeVpcClassicLinkDnsSupportOutput, error) {
	m.addCall("DescribeVpcClassicLinkDnsSupportWithContext")
	m.verifyInput("DescribeVpcClassicLinkDnsSupportWithContext", param0)
	return m.DescribeVpcClassicLinkDnsSupportWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeVpcClassicLinkRequest(param0 *ec2.DescribeVpcClassicLinkInput) (*request.Request, *ec2.DescribeVpcClassicLinkOutput) {
	m.addCall("DescribeVpcClassicLinkRequest")
	m.verifyInput("DescribeVpcClassicLinkRequest", param0)
	return m.DescribeVpcClassicLinkRequestFunc(param0)
}

func (m *ec2Mock) DescribeVpcClassicLinkWithContext(param0 aws.Context, param1 *ec2.DescribeVpcClassicLinkInput, param2 ...request.Option) (*ec2.DescribeVpcClassicLinkOutput, error) {
	m.addCall("DescribeVpcClassicLinkWithContext")
	m.verifyInput("DescribeVpcClassicLinkWithContext", param0)
	return m.DescribeVpcClassicLinkWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeVpcEndpointServices(param0 *ec2.DescribeVpcEndpointServicesInput) (*ec2.DescribeVpcEndpointServicesOutput, error) {
	m.addCall("DescribeVpcEndpointServices")
	m.verifyInput("DescribeVpcEndpointServices", param0)
	return m.DescribeVpcEndpointServicesFunc(param0)
}

func (m *ec2Mock) DescribeVpcEndpointServicesRequest(param0 *ec2.DescribeVpcEndpointServicesInput) (*request.Request, *ec2.DescribeVpcEndpointServicesOutput) {
	m.addCall("DescribeVpcEndpointServicesRequest")
	m.verifyInput("DescribeVpcEndpointServicesRequest", param0)
	return m.DescribeVpcEndpointServicesRequestFunc(param0)
}

func (m *ec2Mock) DescribeVpcEndpointServicesWithContext(param0 aws.Context, param1 *ec2.DescribeVpcEndpointServicesInput, param2 ...request.Option) (*ec2.DescribeVpcEndpointServicesOutput, error) {
	m.addCall("DescribeVpcEndpointServicesWithContext")
	m.verifyInput("DescribeVpcEndpointServicesWithContext", param0)
	return m.DescribeVpcEndpointServicesWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeVpcEndpoints(param0 *ec2.DescribeVpcEndpointsInput) (*ec2.DescribeVpcEndpointsOutput, error) {
	m.addCall("DescribeVpcEndpoints")
	m.verifyInput("DescribeVpcEndpoints", param0)
	return m.DescribeVpcEndpointsFunc(param0)
}

func (m *ec2Mock) DescribeVpcEndpointsRequest(param0 *ec2.DescribeVpcEndpointsInput) (*request.Request, *ec2.DescribeVpcEndpointsOutput) {
	m.addCall("DescribeVpcEndpointsRequest")
	m.verifyInput("DescribeVpcEndpointsRequest", param0)
	return m.DescribeVpcEndpointsRequestFunc(param0)
}

func (m *ec2Mock) DescribeVpcEndpointsWithContext(param0 aws.Context, param1 *ec2.DescribeVpcEndpointsInput, param2 ...request.Option) (*ec2.DescribeVpcEndpointsOutput, error) {
	m.addCall("DescribeVpcEndpointsWithContext")
	m.verifyInput("DescribeVpcEndpointsWithContext", param0)
	return m.DescribeVpcEndpointsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeVpcPeeringConnections(param0 *ec2.DescribeVpcPeeringConnectionsInput) (*ec2.DescribeVpcPeeringConnectionsOutput, error) {
	m.addCall("DescribeVpcPeeringConnections")
	m.verifyInput("DescribeVpcPeeringConnections", param0)
	return m.DescribeVpcPeeringConnectionsFunc(param0)
}

func (m *ec2Mock) DescribeVpcPeeringConnectionsRequest(param0 *ec2.DescribeVpcPeeringConnectionsInput) (*request.Request, *ec2.DescribeVpcPeeringConnectionsOutput) {
	m.addCall("DescribeVpcPeeringConnectionsRequest")
	m.verifyInput("DescribeVpcPeeringConnectionsRequest", param0)
	return m.DescribeVpcPeeringConnectionsRequestFunc(param0)
}

func (m *ec2Mock) DescribeVpcPeeringConnectionsWithContext(param0 aws.Context, param1 *ec2.DescribeVpcPeeringConnectionsInput, param2 ...request.Option) (*ec2.DescribeVpcPeeringConnectionsOutput, error) {
	m.addCall("DescribeVpcPeeringConnectionsWithContext")
	m.verifyInput("DescribeVpcPeeringConnectionsWithContext", param0)
	return m.DescribeVpcPeeringConnectionsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeVpcs(param0 *ec2.DescribeVpcsInput) (*ec2.DescribeVpcsOutput, error) {
	m.addCall("DescribeVpcs")
	m.verifyInput("DescribeVpcs", param0)
	return m.DescribeVpcsFunc(param0)
}

func (m *ec2Mock) DescribeVpcsRequest(param0 *ec2.DescribeVpcsInput) (*request.Request, *ec2.DescribeVpcsOutput) {
	m.addCall("DescribeVpcsRequest")
	m.verifyInput("DescribeVpcsRequest", param0)
	return m.DescribeVpcsRequestFunc(param0)
}

func (m *ec2Mock) DescribeVpcsWithContext(param0 aws.Context, param1 *ec2.DescribeVpcsInput, param2 ...request.Option) (*ec2.DescribeVpcsOutput, error) {
	m.addCall("DescribeVpcsWithContext")
	m.verifyInput("DescribeVpcsWithContext", param0)
	return m.DescribeVpcsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeVpnConnections(param0 *ec2.DescribeVpnConnectionsInput) (*ec2.DescribeVpnConnectionsOutput, error) {
	m.addCall("DescribeVpnConnections")
	m.verifyInput("DescribeVpnConnections", param0)
	return m.DescribeVpnConnectionsFunc(param0)
}

func (m *ec2Mock) DescribeVpnConnectionsRequest(param0 *ec2.DescribeVpnConnectionsInput) (*request.Request, *ec2.DescribeVpnConnectionsOutput) {
	m.addCall("DescribeVpnConnectionsRequest")
	m.verifyInput("DescribeVpnConnectionsRequest", param0)
	return m.DescribeVpnConnectionsRequestFunc(param0)
}

func (m *ec2Mock) DescribeVpnConnectionsWithContext(param0 aws.Context, param1 *ec2.DescribeVpnConnectionsInput, param2 ...request.Option) (*ec2.DescribeVpnConnectionsOutput, error) {
	m.addCall("DescribeVpnConnectionsWithContext")
	m.verifyInput("DescribeVpnConnectionsWithContext", param0)
	return m.DescribeVpnConnectionsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DescribeVpnGateways(param0 *ec2.DescribeVpnGatewaysInput) (*ec2.DescribeVpnGatewaysOutput, error) {
	m.addCall("DescribeVpnGateways")
	m.verifyInput("DescribeVpnGateways", param0)
	return m.DescribeVpnGatewaysFunc(param0)
}

func (m *ec2Mock) DescribeVpnGatewaysRequest(param0 *ec2.DescribeVpnGatewaysInput) (*request.Request, *ec2.DescribeVpnGatewaysOutput) {
	m.addCall("DescribeVpnGatewaysRequest")
	m.verifyInput("DescribeVpnGatewaysRequest", param0)
	return m.DescribeVpnGatewaysRequestFunc(param0)
}

func (m *ec2Mock) DescribeVpnGatewaysWithContext(param0 aws.Context, param1 *ec2.DescribeVpnGatewaysInput, param2 ...request.Option) (*ec2.DescribeVpnGatewaysOutput, error) {
	m.addCall("DescribeVpnGatewaysWithContext")
	m.verifyInput("DescribeVpnGatewaysWithContext", param0)
	return m.DescribeVpnGatewaysWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DetachClassicLinkVpc(param0 *ec2.DetachClassicLinkVpcInput) (*ec2.DetachClassicLinkVpcOutput, error) {
	m.addCall("DetachClassicLinkVpc")
	m.verifyInput("DetachClassicLinkVpc", param0)
	return m.DetachClassicLinkVpcFunc(param0)
}

func (m *ec2Mock) DetachClassicLinkVpcRequest(param0 *ec2.DetachClassicLinkVpcInput) (*request.Request, *ec2.DetachClassicLinkVpcOutput) {
	m.addCall("DetachClassicLinkVpcRequest")
	m.verifyInput("DetachClassicLinkVpcRequest", param0)
	return m.DetachClassicLinkVpcRequestFunc(param0)
}

func (m *ec2Mock) DetachClassicLinkVpcWithContext(param0 aws.Context, param1 *ec2.DetachClassicLinkVpcInput, param2 ...request.Option) (*ec2.DetachClassicLinkVpcOutput, error) {
	m.addCall("DetachClassicLinkVpcWithContext")
	m.verifyInput("DetachClassicLinkVpcWithContext", param0)
	return m.DetachClassicLinkVpcWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DetachInternetGateway(param0 *ec2.DetachInternetGatewayInput) (*ec2.DetachInternetGatewayOutput, error) {
	m.addCall("DetachInternetGateway")
	m.verifyInput("DetachInternetGateway", param0)
	return m.DetachInternetGatewayFunc(param0)
}

func (m *ec2Mock) DetachInternetGatewayRequest(param0 *ec2.DetachInternetGatewayInput) (*request.Request, *ec2.DetachInternetGatewayOutput) {
	m.addCall("DetachInternetGatewayRequest")
	m.verifyInput("DetachInternetGatewayRequest", param0)
	return m.DetachInternetGatewayRequestFunc(param0)
}

func (m *ec2Mock) DetachInternetGatewayWithContext(param0 aws.Context, param1 *ec2.DetachInternetGatewayInput, param2 ...request.Option) (*ec2.DetachInternetGatewayOutput, error) {
	m.addCall("DetachInternetGatewayWithContext")
	m.verifyInput("DetachInternetGatewayWithContext", param0)
	return m.DetachInternetGatewayWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DetachNetworkInterface(param0 *ec2.DetachNetworkInterfaceInput) (*ec2.DetachNetworkInterfaceOutput, error) {
	m.addCall("DetachNetworkInterface")
	m.verifyInput("DetachNetworkInterface", param0)
	return m.DetachNetworkInterfaceFunc(param0)
}

func (m *ec2Mock) DetachNetworkInterfaceRequest(param0 *ec2.DetachNetworkInterfaceInput) (*request.Request, *ec2.DetachNetworkInterfaceOutput) {
	m.addCall("DetachNetworkInterfaceRequest")
	m.verifyInput("DetachNetworkInterfaceRequest", param0)
	return m.DetachNetworkInterfaceRequestFunc(param0)
}

func (m *ec2Mock) DetachNetworkInterfaceWithContext(param0 aws.Context, param1 *ec2.DetachNetworkInterfaceInput, param2 ...request.Option) (*ec2.DetachNetworkInterfaceOutput, error) {
	m.addCall("DetachNetworkInterfaceWithContext")
	m.verifyInput("DetachNetworkInterfaceWithContext", param0)
	return m.DetachNetworkInterfaceWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DetachVolume(param0 *ec2.DetachVolumeInput) (*ec2.VolumeAttachment, error) {
	m.addCall("DetachVolume")
	m.verifyInput("DetachVolume", param0)
	return m.DetachVolumeFunc(param0)
}

func (m *ec2Mock) DetachVolumeRequest(param0 *ec2.DetachVolumeInput) (*request.Request, *ec2.VolumeAttachment) {
	m.addCall("DetachVolumeRequest")
	m.verifyInput("DetachVolumeRequest", param0)
	return m.DetachVolumeRequestFunc(param0)
}

func (m *ec2Mock) DetachVolumeWithContext(param0 aws.Context, param1 *ec2.DetachVolumeInput, param2 ...request.Option) (*ec2.VolumeAttachment, error) {
	m.addCall("DetachVolumeWithContext")
	m.verifyInput("DetachVolumeWithContext", param0)
	return m.DetachVolumeWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DetachVpnGateway(param0 *ec2.DetachVpnGatewayInput) (*ec2.DetachVpnGatewayOutput, error) {
	m.addCall("DetachVpnGateway")
	m.verifyInput("DetachVpnGateway", param0)
	return m.DetachVpnGatewayFunc(param0)
}

func (m *ec2Mock) DetachVpnGatewayRequest(param0 *ec2.DetachVpnGatewayInput) (*request.Request, *ec2.DetachVpnGatewayOutput) {
	m.addCall("DetachVpnGatewayRequest")
	m.verifyInput("DetachVpnGatewayRequest", param0)
	return m.DetachVpnGatewayRequestFunc(param0)
}

func (m *ec2Mock) DetachVpnGatewayWithContext(param0 aws.Context, param1 *ec2.DetachVpnGatewayInput, param2 ...request.Option) (*ec2.DetachVpnGatewayOutput, error) {
	m.addCall("DetachVpnGatewayWithContext")
	m.verifyInput("DetachVpnGatewayWithContext", param0)
	return m.DetachVpnGatewayWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DisableVgwRoutePropagation(param0 *ec2.DisableVgwRoutePropagationInput) (*ec2.DisableVgwRoutePropagationOutput, error) {
	m.addCall("DisableVgwRoutePropagation")
	m.verifyInput("DisableVgwRoutePropagation", param0)
	return m.DisableVgwRoutePropagationFunc(param0)
}

func (m *ec2Mock) DisableVgwRoutePropagationRequest(param0 *ec2.DisableVgwRoutePropagationInput) (*request.Request, *ec2.DisableVgwRoutePropagationOutput) {
	m.addCall("DisableVgwRoutePropagationRequest")
	m.verifyInput("DisableVgwRoutePropagationRequest", param0)
	return m.DisableVgwRoutePropagationRequestFunc(param0)
}

func (m *ec2Mock) DisableVgwRoutePropagationWithContext(param0 aws.Context, param1 *ec2.DisableVgwRoutePropagationInput, param2 ...request.Option) (*ec2.DisableVgwRoutePropagationOutput, error) {
	m.addCall("DisableVgwRoutePropagationWithContext")
	m.verifyInput("DisableVgwRoutePropagationWithContext", param0)
	return m.DisableVgwRoutePropagationWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DisableVpcClassicLink(param0 *ec2.DisableVpcClassicLinkInput) (*ec2.DisableVpcClassicLinkOutput, error) {
	m.addCall("DisableVpcClassicLink")
	m.verifyInput("DisableVpcClassicLink", param0)
	return m.DisableVpcClassicLinkFunc(param0)
}

func (m *ec2Mock) DisableVpcClassicLinkDnsSupport(param0 *ec2.DisableVpcClassicLinkDnsSupportInput) (*ec2.DisableVpcClassicLinkDnsSupportOutput, error) {
	m.addCall("DisableVpcClassicLinkDnsSupport")
	m.verifyInput("DisableVpcClassicLinkDnsSupport", param0)
	return m.DisableVpcClassicLinkDnsSupportFunc(param0)
}

func (m *ec2Mock) DisableVpcClassicLinkDnsSupportRequest(param0 *ec2.DisableVpcClassicLinkDnsSupportInput) (*request.Request, *ec2.DisableVpcClassicLinkDnsSupportOutput) {
	m.addCall("DisableVpcClassicLinkDnsSupportRequest")
	m.verifyInput("DisableVpcClassicLinkDnsSupportRequest", param0)
	return m.DisableVpcClassicLinkDnsSupportRequestFunc(param0)
}

func (m *ec2Mock) DisableVpcClassicLinkDnsSupportWithContext(param0 aws.Context, param1 *ec2.DisableVpcClassicLinkDnsSupportInput, param2 ...request.Option) (*ec2.DisableVpcClassicLinkDnsSupportOutput, error) {
	m.addCall("DisableVpcClassicLinkDnsSupportWithContext")
	m.verifyInput("DisableVpcClassicLinkDnsSupportWithContext", param0)
	return m.DisableVpcClassicLinkDnsSupportWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DisableVpcClassicLinkRequest(param0 *ec2.DisableVpcClassicLinkInput) (*request.Request, *ec2.DisableVpcClassicLinkOutput) {
	m.addCall("DisableVpcClassicLinkRequest")
	m.verifyInput("DisableVpcClassicLinkRequest", param0)
	return m.DisableVpcClassicLinkRequestFunc(param0)
}

func (m *ec2Mock) DisableVpcClassicLinkWithContext(param0 aws.Context, param1 *ec2.DisableVpcClassicLinkInput, param2 ...request.Option) (*ec2.DisableVpcClassicLinkOutput, error) {
	m.addCall("DisableVpcClassicLinkWithContext")
	m.verifyInput("DisableVpcClassicLinkWithContext", param0)
	return m.DisableVpcClassicLinkWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DisassociateAddress(param0 *ec2.DisassociateAddressInput) (*ec2.DisassociateAddressOutput, error) {
	m.addCall("DisassociateAddress")
	m.verifyInput("DisassociateAddress", param0)
	return m.DisassociateAddressFunc(param0)
}

func (m *ec2Mock) DisassociateAddressRequest(param0 *ec2.DisassociateAddressInput) (*request.Request, *ec2.DisassociateAddressOutput) {
	m.addCall("DisassociateAddressRequest")
	m.verifyInput("DisassociateAddressRequest", param0)
	return m.DisassociateAddressRequestFunc(param0)
}

func (m *ec2Mock) DisassociateAddressWithContext(param0 aws.Context, param1 *ec2.DisassociateAddressInput, param2 ...request.Option) (*ec2.DisassociateAddressOutput, error) {
	m.addCall("DisassociateAddressWithContext")
	m.verifyInput("DisassociateAddressWithContext", param0)
	return m.DisassociateAddressWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DisassociateIamInstanceProfile(param0 *ec2.DisassociateIamInstanceProfileInput) (*ec2.DisassociateIamInstanceProfileOutput, error) {
	m.addCall("DisassociateIamInstanceProfile")
	m.verifyInput("DisassociateIamInstanceProfile", param0)
	return m.DisassociateIamInstanceProfileFunc(param0)
}

func (m *ec2Mock) DisassociateIamInstanceProfileRequest(param0 *ec2.DisassociateIamInstanceProfileInput) (*request.Request, *ec2.DisassociateIamInstanceProfileOutput) {
	m.addCall("DisassociateIamInstanceProfileRequest")
	m.verifyInput("DisassociateIamInstanceProfileRequest", param0)
	return m.DisassociateIamInstanceProfileRequestFunc(param0)
}

func (m *ec2Mock) DisassociateIamInstanceProfileWithContext(param0 aws.Context, param1 *ec2.DisassociateIamInstanceProfileInput, param2 ...request.Option) (*ec2.DisassociateIamInstanceProfileOutput, error) {
	m.addCall("DisassociateIamInstanceProfileWithContext")
	m.verifyInput("DisassociateIamInstanceProfileWithContext", param0)
	return m.DisassociateIamInstanceProfileWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DisassociateRouteTable(param0 *ec2.DisassociateRouteTableInput) (*ec2.DisassociateRouteTableOutput, error) {
	m.addCall("DisassociateRouteTable")
	m.verifyInput("DisassociateRouteTable", param0)
	return m.DisassociateRouteTableFunc(param0)
}

func (m *ec2Mock) DisassociateRouteTableRequest(param0 *ec2.DisassociateRouteTableInput) (*request.Request, *ec2.DisassociateRouteTableOutput) {
	m.addCall("DisassociateRouteTableRequest")
	m.verifyInput("DisassociateRouteTableRequest", param0)
	return m.DisassociateRouteTableRequestFunc(param0)
}

func (m *ec2Mock) DisassociateRouteTableWithContext(param0 aws.Context, param1 *ec2.DisassociateRouteTableInput, param2 ...request.Option) (*ec2.DisassociateRouteTableOutput, error) {
	m.addCall("DisassociateRouteTableWithContext")
	m.verifyInput("DisassociateRouteTableWithContext", param0)
	return m.DisassociateRouteTableWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DisassociateSubnetCidrBlock(param0 *ec2.DisassociateSubnetCidrBlockInput) (*ec2.DisassociateSubnetCidrBlockOutput, error) {
	m.addCall("DisassociateSubnetCidrBlock")
	m.verifyInput("DisassociateSubnetCidrBlock", param0)
	return m.DisassociateSubnetCidrBlockFunc(param0)
}

func (m *ec2Mock) DisassociateSubnetCidrBlockRequest(param0 *ec2.DisassociateSubnetCidrBlockInput) (*request.Request, *ec2.DisassociateSubnetCidrBlockOutput) {
	m.addCall("DisassociateSubnetCidrBlockRequest")
	m.verifyInput("DisassociateSubnetCidrBlockRequest", param0)
	return m.DisassociateSubnetCidrBlockRequestFunc(param0)
}

func (m *ec2Mock) DisassociateSubnetCidrBlockWithContext(param0 aws.Context, param1 *ec2.DisassociateSubnetCidrBlockInput, param2 ...request.Option) (*ec2.DisassociateSubnetCidrBlockOutput, error) {
	m.addCall("DisassociateSubnetCidrBlockWithContext")
	m.verifyInput("DisassociateSubnetCidrBlockWithContext", param0)
	return m.DisassociateSubnetCidrBlockWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) DisassociateVpcCidrBlock(param0 *ec2.DisassociateVpcCidrBlockInput) (*ec2.DisassociateVpcCidrBlockOutput, error) {
	m.addCall("DisassociateVpcCidrBlock")
	m.verifyInput("DisassociateVpcCidrBlock", param0)
	return m.DisassociateVpcCidrBlockFunc(param0)
}

func (m *ec2Mock) DisassociateVpcCidrBlockRequest(param0 *ec2.DisassociateVpcCidrBlockInput) (*request.Request, *ec2.DisassociateVpcCidrBlockOutput) {
	m.addCall("DisassociateVpcCidrBlockRequest")
	m.verifyInput("DisassociateVpcCidrBlockRequest", param0)
	return m.DisassociateVpcCidrBlockRequestFunc(param0)
}

func (m *ec2Mock) DisassociateVpcCidrBlockWithContext(param0 aws.Context, param1 *ec2.DisassociateVpcCidrBlockInput, param2 ...request.Option) (*ec2.DisassociateVpcCidrBlockOutput, error) {
	m.addCall("DisassociateVpcCidrBlockWithContext")
	m.verifyInput("DisassociateVpcCidrBlockWithContext", param0)
	return m.DisassociateVpcCidrBlockWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) EnableVgwRoutePropagation(param0 *ec2.EnableVgwRoutePropagationInput) (*ec2.EnableVgwRoutePropagationOutput, error) {
	m.addCall("EnableVgwRoutePropagation")
	m.verifyInput("EnableVgwRoutePropagation", param0)
	return m.EnableVgwRoutePropagationFunc(param0)
}

func (m *ec2Mock) EnableVgwRoutePropagationRequest(param0 *ec2.EnableVgwRoutePropagationInput) (*request.Request, *ec2.EnableVgwRoutePropagationOutput) {
	m.addCall("EnableVgwRoutePropagationRequest")
	m.verifyInput("EnableVgwRoutePropagationRequest", param0)
	return m.EnableVgwRoutePropagationRequestFunc(param0)
}

func (m *ec2Mock) EnableVgwRoutePropagationWithContext(param0 aws.Context, param1 *ec2.EnableVgwRoutePropagationInput, param2 ...request.Option) (*ec2.EnableVgwRoutePropagationOutput, error) {
	m.addCall("EnableVgwRoutePropagationWithContext")
	m.verifyInput("EnableVgwRoutePropagationWithContext", param0)
	return m.EnableVgwRoutePropagationWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) EnableVolumeIO(param0 *ec2.EnableVolumeIOInput) (*ec2.EnableVolumeIOOutput, error) {
	m.addCall("EnableVolumeIO")
	m.verifyInput("EnableVolumeIO", param0)
	return m.EnableVolumeIOFunc(param0)
}

func (m *ec2Mock) EnableVolumeIORequest(param0 *ec2.EnableVolumeIOInput) (*request.Request, *ec2.EnableVolumeIOOutput) {
	m.addCall("EnableVolumeIORequest")
	m.verifyInput("EnableVolumeIORequest", param0)
	return m.EnableVolumeIORequestFunc(param0)
}

func (m *ec2Mock) EnableVolumeIOWithContext(param0 aws.Context, param1 *ec2.EnableVolumeIOInput, param2 ...request.Option) (*ec2.EnableVolumeIOOutput, error) {
	m.addCall("EnableVolumeIOWithContext")
	m.verifyInput("EnableVolumeIOWithContext", param0)
	return m.EnableVolumeIOWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) EnableVpcClassicLink(param0 *ec2.EnableVpcClassicLinkInput) (*ec2.EnableVpcClassicLinkOutput, error) {
	m.addCall("EnableVpcClassicLink")
	m.verifyInput("EnableVpcClassicLink", param0)
	return m.EnableVpcClassicLinkFunc(param0)
}

func (m *ec2Mock) EnableVpcClassicLinkDnsSupport(param0 *ec2.EnableVpcClassicLinkDnsSupportInput) (*ec2.EnableVpcClassicLinkDnsSupportOutput, error) {
	m.addCall("EnableVpcClassicLinkDnsSupport")
	m.verifyInput("EnableVpcClassicLinkDnsSupport", param0)
	return m.EnableVpcClassicLinkDnsSupportFunc(param0)
}

func (m *ec2Mock) EnableVpcClassicLinkDnsSupportRequest(param0 *ec2.EnableVpcClassicLinkDnsSupportInput) (*request.Request, *ec2.EnableVpcClassicLinkDnsSupportOutput) {
	m.addCall("EnableVpcClassicLinkDnsSupportRequest")
	m.verifyInput("EnableVpcClassicLinkDnsSupportRequest", param0)
	return m.EnableVpcClassicLinkDnsSupportRequestFunc(param0)
}

func (m *ec2Mock) EnableVpcClassicLinkDnsSupportWithContext(param0 aws.Context, param1 *ec2.EnableVpcClassicLinkDnsSupportInput, param2 ...request.Option) (*ec2.EnableVpcClassicLinkDnsSupportOutput, error) {
	m.addCall("EnableVpcClassicLinkDnsSupportWithContext")
	m.verifyInput("EnableVpcClassicLinkDnsSupportWithContext", param0)
	return m.EnableVpcClassicLinkDnsSupportWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) EnableVpcClassicLinkRequest(param0 *ec2.EnableVpcClassicLinkInput) (*request.Request, *ec2.EnableVpcClassicLinkOutput) {
	m.addCall("EnableVpcClassicLinkRequest")
	m.verifyInput("EnableVpcClassicLinkRequest", param0)
	return m.EnableVpcClassicLinkRequestFunc(param0)
}

func (m *ec2Mock) EnableVpcClassicLinkWithContext(param0 aws.Context, param1 *ec2.EnableVpcClassicLinkInput, param2 ...request.Option) (*ec2.EnableVpcClassicLinkOutput, error) {
	m.addCall("EnableVpcClassicLinkWithContext")
	m.verifyInput("EnableVpcClassicLinkWithContext", param0)
	return m.EnableVpcClassicLinkWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) GetConsoleOutput(param0 *ec2.GetConsoleOutputInput) (*ec2.GetConsoleOutputOutput, error) {
	m.addCall("GetConsoleOutput")
	m.verifyInput("GetConsoleOutput", param0)
	return m.GetConsoleOutputFunc(param0)
}

func (m *ec2Mock) GetConsoleOutputRequest(param0 *ec2.GetConsoleOutputInput) (*request.Request, *ec2.GetConsoleOutputOutput) {
	m.addCall("GetConsoleOutputRequest")
	m.verifyInput("GetConsoleOutputRequest", param0)
	return m.GetConsoleOutputRequestFunc(param0)
}

func (m *ec2Mock) GetConsoleOutputWithContext(param0 aws.Context, param1 *ec2.GetConsoleOutputInput, param2 ...request.Option) (*ec2.GetConsoleOutputOutput, error) {
	m.addCall("GetConsoleOutputWithContext")
	m.verifyInput("GetConsoleOutputWithContext", param0)
	return m.GetConsoleOutputWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) GetConsoleScreenshot(param0 *ec2.GetConsoleScreenshotInput) (*ec2.GetConsoleScreenshotOutput, error) {
	m.addCall("GetConsoleScreenshot")
	m.verifyInput("GetConsoleScreenshot", param0)
	return m.GetConsoleScreenshotFunc(param0)
}

func (m *ec2Mock) GetConsoleScreenshotRequest(param0 *ec2.GetConsoleScreenshotInput) (*request.Request, *ec2.GetConsoleScreenshotOutput) {
	m.addCall("GetConsoleScreenshotRequest")
	m.verifyInput("GetConsoleScreenshotRequest", param0)
	return m.GetConsoleScreenshotRequestFunc(param0)
}

func (m *ec2Mock) GetConsoleScreenshotWithContext(param0 aws.Context, param1 *ec2.GetConsoleScreenshotInput, param2 ...request.Option) (*ec2.GetConsoleScreenshotOutput, error) {
	m.addCall("GetConsoleScreenshotWithContext")
	m.verifyInput("GetConsoleScreenshotWithContext", param0)
	return m.GetConsoleScreenshotWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) GetHostReservationPurchasePreview(param0 *ec2.GetHostReservationPurchasePreviewInput) (*ec2.GetHostReservationPurchasePreviewOutput, error) {
	m.addCall("GetHostReservationPurchasePreview")
	m.verifyInput("GetHostReservationPurchasePreview", param0)
	return m.GetHostReservationPurchasePreviewFunc(param0)
}

func (m *ec2Mock) GetHostReservationPurchasePreviewRequest(param0 *ec2.GetHostReservationPurchasePreviewInput) (*request.Request, *ec2.GetHostReservationPurchasePreviewOutput) {
	m.addCall("GetHostReservationPurchasePreviewRequest")
	m.verifyInput("GetHostReservationPurchasePreviewRequest", param0)
	return m.GetHostReservationPurchasePreviewRequestFunc(param0)
}

func (m *ec2Mock) GetHostReservationPurchasePreviewWithContext(param0 aws.Context, param1 *ec2.GetHostReservationPurchasePreviewInput, param2 ...request.Option) (*ec2.GetHostReservationPurchasePreviewOutput, error) {
	m.addCall("GetHostReservationPurchasePreviewWithContext")
	m.verifyInput("GetHostReservationPurchasePreviewWithContext", param0)
	return m.GetHostReservationPurchasePreviewWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) GetPasswordData(param0 *ec2.GetPasswordDataInput) (*ec2.GetPasswordDataOutput, error) {
	m.addCall("GetPasswordData")
	m.verifyInput("GetPasswordData", param0)
	return m.GetPasswordDataFunc(param0)
}

func (m *ec2Mock) GetPasswordDataRequest(param0 *ec2.GetPasswordDataInput) (*request.Request, *ec2.GetPasswordDataOutput) {
	m.addCall("GetPasswordDataRequest")
	m.verifyInput("GetPasswordDataRequest", param0)
	return m.GetPasswordDataRequestFunc(param0)
}

func (m *ec2Mock) GetPasswordDataWithContext(param0 aws.Context, param1 *ec2.GetPasswordDataInput, param2 ...request.Option) (*ec2.GetPasswordDataOutput, error) {
	m.addCall("GetPasswordDataWithContext")
	m.verifyInput("GetPasswordDataWithContext", param0)
	return m.GetPasswordDataWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) GetReservedInstancesExchangeQuote(param0 *ec2.GetReservedInstancesExchangeQuoteInput) (*ec2.GetReservedInstancesExchangeQuoteOutput, error) {
	m.addCall("GetReservedInstancesExchangeQuote")
	m.verifyInput("GetReservedInstancesExchangeQuote", param0)
	return m.GetReservedInstancesExchangeQuoteFunc(param0)
}

func (m *ec2Mock) GetReservedInstancesExchangeQuoteRequest(param0 *ec2.GetReservedInstancesExchangeQuoteInput) (*request.Request, *ec2.GetReservedInstancesExchangeQuoteOutput) {
	m.addCall("GetReservedInstancesExchangeQuoteRequest")
	m.verifyInput("GetReservedInstancesExchangeQuoteRequest", param0)
	return m.GetReservedInstancesExchangeQuoteRequestFunc(param0)
}

func (m *ec2Mock) GetReservedInstancesExchangeQuoteWithContext(param0 aws.Context, param1 *ec2.GetReservedInstancesExchangeQuoteInput, param2 ...request.Option) (*ec2.GetReservedInstancesExchangeQuoteOutput, error) {
	m.addCall("GetReservedInstancesExchangeQuoteWithContext")
	m.verifyInput("GetReservedInstancesExchangeQuoteWithContext", param0)
	return m.GetReservedInstancesExchangeQuoteWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ImportImage(param0 *ec2.ImportImageInput) (*ec2.ImportImageOutput, error) {
	m.addCall("ImportImage")
	m.verifyInput("ImportImage", param0)
	return m.ImportImageFunc(param0)
}

func (m *ec2Mock) ImportImageRequest(param0 *ec2.ImportImageInput) (*request.Request, *ec2.ImportImageOutput) {
	m.addCall("ImportImageRequest")
	m.verifyInput("ImportImageRequest", param0)
	return m.ImportImageRequestFunc(param0)
}

func (m *ec2Mock) ImportImageWithContext(param0 aws.Context, param1 *ec2.ImportImageInput, param2 ...request.Option) (*ec2.ImportImageOutput, error) {
	m.addCall("ImportImageWithContext")
	m.verifyInput("ImportImageWithContext", param0)
	return m.ImportImageWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ImportInstance(param0 *ec2.ImportInstanceInput) (*ec2.ImportInstanceOutput, error) {
	m.addCall("ImportInstance")
	m.verifyInput("ImportInstance", param0)
	return m.ImportInstanceFunc(param0)
}

func (m *ec2Mock) ImportInstanceRequest(param0 *ec2.ImportInstanceInput) (*request.Request, *ec2.ImportInstanceOutput) {
	m.addCall("ImportInstanceRequest")
	m.verifyInput("ImportInstanceRequest", param0)
	return m.ImportInstanceRequestFunc(param0)
}

func (m *ec2Mock) ImportInstanceWithContext(param0 aws.Context, param1 *ec2.ImportInstanceInput, param2 ...request.Option) (*ec2.ImportInstanceOutput, error) {
	m.addCall("ImportInstanceWithContext")
	m.verifyInput("ImportInstanceWithContext", param0)
	return m.ImportInstanceWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ImportKeyPair(param0 *ec2.ImportKeyPairInput) (*ec2.ImportKeyPairOutput, error) {
	m.addCall("ImportKeyPair")
	m.verifyInput("ImportKeyPair", param0)
	return m.ImportKeyPairFunc(param0)
}

func (m *ec2Mock) ImportKeyPairRequest(param0 *ec2.ImportKeyPairInput) (*request.Request, *ec2.ImportKeyPairOutput) {
	m.addCall("ImportKeyPairRequest")
	m.verifyInput("ImportKeyPairRequest", param0)
	return m.ImportKeyPairRequestFunc(param0)
}

func (m *ec2Mock) ImportKeyPairWithContext(param0 aws.Context, param1 *ec2.ImportKeyPairInput, param2 ...request.Option) (*ec2.ImportKeyPairOutput, error) {
	m.addCall("ImportKeyPairWithContext")
	m.verifyInput("ImportKeyPairWithContext", param0)
	return m.ImportKeyPairWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ImportSnapshot(param0 *ec2.ImportSnapshotInput) (*ec2.ImportSnapshotOutput, error) {
	m.addCall("ImportSnapshot")
	m.verifyInput("ImportSnapshot", param0)
	return m.ImportSnapshotFunc(param0)
}

func (m *ec2Mock) ImportSnapshotRequest(param0 *ec2.ImportSnapshotInput) (*request.Request, *ec2.ImportSnapshotOutput) {
	m.addCall("ImportSnapshotRequest")
	m.verifyInput("ImportSnapshotRequest", param0)
	return m.ImportSnapshotRequestFunc(param0)
}

func (m *ec2Mock) ImportSnapshotWithContext(param0 aws.Context, param1 *ec2.ImportSnapshotInput, param2 ...request.Option) (*ec2.ImportSnapshotOutput, error) {
	m.addCall("ImportSnapshotWithContext")
	m.verifyInput("ImportSnapshotWithContext", param0)
	return m.ImportSnapshotWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ImportVolume(param0 *ec2.ImportVolumeInput) (*ec2.ImportVolumeOutput, error) {
	m.addCall("ImportVolume")
	m.verifyInput("ImportVolume", param0)
	return m.ImportVolumeFunc(param0)
}

func (m *ec2Mock) ImportVolumeRequest(param0 *ec2.ImportVolumeInput) (*request.Request, *ec2.ImportVolumeOutput) {
	m.addCall("ImportVolumeRequest")
	m.verifyInput("ImportVolumeRequest", param0)
	return m.ImportVolumeRequestFunc(param0)
}

func (m *ec2Mock) ImportVolumeWithContext(param0 aws.Context, param1 *ec2.ImportVolumeInput, param2 ...request.Option) (*ec2.ImportVolumeOutput, error) {
	m.addCall("ImportVolumeWithContext")
	m.verifyInput("ImportVolumeWithContext", param0)
	return m.ImportVolumeWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ModifyFpgaImageAttribute(param0 *ec2.ModifyFpgaImageAttributeInput) (*ec2.ModifyFpgaImageAttributeOutput, error) {
	m.addCall("ModifyFpgaImageAttribute")
	m.verifyInput("ModifyFpgaImageAttribute", param0)
	return m.ModifyFpgaImageAttributeFunc(param0)
}

func (m *ec2Mock) ModifyFpgaImageAttributeRequest(param0 *ec2.ModifyFpgaImageAttributeInput) (*request.Request, *ec2.ModifyFpgaImageAttributeOutput) {
	m.addCall("ModifyFpgaImageAttributeRequest")
	m.verifyInput("ModifyFpgaImageAttributeRequest", param0)
	return m.ModifyFpgaImageAttributeRequestFunc(param0)
}

func (m *ec2Mock) ModifyFpgaImageAttributeWithContext(param0 aws.Context, param1 *ec2.ModifyFpgaImageAttributeInput, param2 ...request.Option) (*ec2.ModifyFpgaImageAttributeOutput, error) {
	m.addCall("ModifyFpgaImageAttributeWithContext")
	m.verifyInput("ModifyFpgaImageAttributeWithContext", param0)
	return m.ModifyFpgaImageAttributeWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ModifyHosts(param0 *ec2.ModifyHostsInput) (*ec2.ModifyHostsOutput, error) {
	m.addCall("ModifyHosts")
	m.verifyInput("ModifyHosts", param0)
	return m.ModifyHostsFunc(param0)
}

func (m *ec2Mock) ModifyHostsRequest(param0 *ec2.ModifyHostsInput) (*request.Request, *ec2.ModifyHostsOutput) {
	m.addCall("ModifyHostsRequest")
	m.verifyInput("ModifyHostsRequest", param0)
	return m.ModifyHostsRequestFunc(param0)
}

func (m *ec2Mock) ModifyHostsWithContext(param0 aws.Context, param1 *ec2.ModifyHostsInput, param2 ...request.Option) (*ec2.ModifyHostsOutput, error) {
	m.addCall("ModifyHostsWithContext")
	m.verifyInput("ModifyHostsWithContext", param0)
	return m.ModifyHostsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ModifyIdFormat(param0 *ec2.ModifyIdFormatInput) (*ec2.ModifyIdFormatOutput, error) {
	m.addCall("ModifyIdFormat")
	m.verifyInput("ModifyIdFormat", param0)
	return m.ModifyIdFormatFunc(param0)
}

func (m *ec2Mock) ModifyIdFormatRequest(param0 *ec2.ModifyIdFormatInput) (*request.Request, *ec2.ModifyIdFormatOutput) {
	m.addCall("ModifyIdFormatRequest")
	m.verifyInput("ModifyIdFormatRequest", param0)
	return m.ModifyIdFormatRequestFunc(param0)
}

func (m *ec2Mock) ModifyIdFormatWithContext(param0 aws.Context, param1 *ec2.ModifyIdFormatInput, param2 ...request.Option) (*ec2.ModifyIdFormatOutput, error) {
	m.addCall("ModifyIdFormatWithContext")
	m.verifyInput("ModifyIdFormatWithContext", param0)
	return m.ModifyIdFormatWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ModifyIdentityIdFormat(param0 *ec2.ModifyIdentityIdFormatInput) (*ec2.ModifyIdentityIdFormatOutput, error) {
	m.addCall("ModifyIdentityIdFormat")
	m.verifyInput("ModifyIdentityIdFormat", param0)
	return m.ModifyIdentityIdFormatFunc(param0)
}

func (m *ec2Mock) ModifyIdentityIdFormatRequest(param0 *ec2.ModifyIdentityIdFormatInput) (*request.Request, *ec2.ModifyIdentityIdFormatOutput) {
	m.addCall("ModifyIdentityIdFormatRequest")
	m.verifyInput("ModifyIdentityIdFormatRequest", param0)
	return m.ModifyIdentityIdFormatRequestFunc(param0)
}

func (m *ec2Mock) ModifyIdentityIdFormatWithContext(param0 aws.Context, param1 *ec2.ModifyIdentityIdFormatInput, param2 ...request.Option) (*ec2.ModifyIdentityIdFormatOutput, error) {
	m.addCall("ModifyIdentityIdFormatWithContext")
	m.verifyInput("ModifyIdentityIdFormatWithContext", param0)
	return m.ModifyIdentityIdFormatWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ModifyImageAttribute(param0 *ec2.ModifyImageAttributeInput) (*ec2.ModifyImageAttributeOutput, error) {
	m.addCall("ModifyImageAttribute")
	m.verifyInput("ModifyImageAttribute", param0)
	return m.ModifyImageAttributeFunc(param0)
}

func (m *ec2Mock) ModifyImageAttributeRequest(param0 *ec2.ModifyImageAttributeInput) (*request.Request, *ec2.ModifyImageAttributeOutput) {
	m.addCall("ModifyImageAttributeRequest")
	m.verifyInput("ModifyImageAttributeRequest", param0)
	return m.ModifyImageAttributeRequestFunc(param0)
}

func (m *ec2Mock) ModifyImageAttributeWithContext(param0 aws.Context, param1 *ec2.ModifyImageAttributeInput, param2 ...request.Option) (*ec2.ModifyImageAttributeOutput, error) {
	m.addCall("ModifyImageAttributeWithContext")
	m.verifyInput("ModifyImageAttributeWithContext", param0)
	return m.ModifyImageAttributeWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ModifyInstanceAttribute(param0 *ec2.ModifyInstanceAttributeInput) (*ec2.ModifyInstanceAttributeOutput, error) {
	m.addCall("ModifyInstanceAttribute")
	m.verifyInput("ModifyInstanceAttribute", param0)
	return m.ModifyInstanceAttributeFunc(param0)
}

func (m *ec2Mock) ModifyInstanceAttributeRequest(param0 *ec2.ModifyInstanceAttributeInput) (*request.Request, *ec2.ModifyInstanceAttributeOutput) {
	m.addCall("ModifyInstanceAttributeRequest")
	m.verifyInput("ModifyInstanceAttributeRequest", param0)
	return m.ModifyInstanceAttributeRequestFunc(param0)
}

func (m *ec2Mock) ModifyInstanceAttributeWithContext(param0 aws.Context, param1 *ec2.ModifyInstanceAttributeInput, param2 ...request.Option) (*ec2.ModifyInstanceAttributeOutput, error) {
	m.addCall("ModifyInstanceAttributeWithContext")
	m.verifyInput("ModifyInstanceAttributeWithContext", param0)
	return m.ModifyInstanceAttributeWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ModifyInstancePlacement(param0 *ec2.ModifyInstancePlacementInput) (*ec2.ModifyInstancePlacementOutput, error) {
	m.addCall("ModifyInstancePlacement")
	m.verifyInput("ModifyInstancePlacement", param0)
	return m.ModifyInstancePlacementFunc(param0)
}

func (m *ec2Mock) ModifyInstancePlacementRequest(param0 *ec2.ModifyInstancePlacementInput) (*request.Request, *ec2.ModifyInstancePlacementOutput) {
	m.addCall("ModifyInstancePlacementRequest")
	m.verifyInput("ModifyInstancePlacementRequest", param0)
	return m.ModifyInstancePlacementRequestFunc(param0)
}

func (m *ec2Mock) ModifyInstancePlacementWithContext(param0 aws.Context, param1 *ec2.ModifyInstancePlacementInput, param2 ...request.Option) (*ec2.ModifyInstancePlacementOutput, error) {
	m.addCall("ModifyInstancePlacementWithContext")
	m.verifyInput("ModifyInstancePlacementWithContext", param0)
	return m.ModifyInstancePlacementWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ModifyNetworkInterfaceAttribute(param0 *ec2.ModifyNetworkInterfaceAttributeInput) (*ec2.ModifyNetworkInterfaceAttributeOutput, error) {
	m.addCall("ModifyNetworkInterfaceAttribute")
	m.verifyInput("ModifyNetworkInterfaceAttribute", param0)
	return m.ModifyNetworkInterfaceAttributeFunc(param0)
}

func (m *ec2Mock) ModifyNetworkInterfaceAttributeRequest(param0 *ec2.ModifyNetworkInterfaceAttributeInput) (*request.Request, *ec2.ModifyNetworkInterfaceAttributeOutput) {
	m.addCall("ModifyNetworkInterfaceAttributeRequest")
	m.verifyInput("ModifyNetworkInterfaceAttributeRequest", param0)
	return m.ModifyNetworkInterfaceAttributeRequestFunc(param0)
}

func (m *ec2Mock) ModifyNetworkInterfaceAttributeWithContext(param0 aws.Context, param1 *ec2.ModifyNetworkInterfaceAttributeInput, param2 ...request.Option) (*ec2.ModifyNetworkInterfaceAttributeOutput, error) {
	m.addCall("ModifyNetworkInterfaceAttributeWithContext")
	m.verifyInput("ModifyNetworkInterfaceAttributeWithContext", param0)
	return m.ModifyNetworkInterfaceAttributeWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ModifyReservedInstances(param0 *ec2.ModifyReservedInstancesInput) (*ec2.ModifyReservedInstancesOutput, error) {
	m.addCall("ModifyReservedInstances")
	m.verifyInput("ModifyReservedInstances", param0)
	return m.ModifyReservedInstancesFunc(param0)
}

func (m *ec2Mock) ModifyReservedInstancesRequest(param0 *ec2.ModifyReservedInstancesInput) (*request.Request, *ec2.ModifyReservedInstancesOutput) {
	m.addCall("ModifyReservedInstancesRequest")
	m.verifyInput("ModifyReservedInstancesRequest", param0)
	return m.ModifyReservedInstancesRequestFunc(param0)
}

func (m *ec2Mock) ModifyReservedInstancesWithContext(param0 aws.Context, param1 *ec2.ModifyReservedInstancesInput, param2 ...request.Option) (*ec2.ModifyReservedInstancesOutput, error) {
	m.addCall("ModifyReservedInstancesWithContext")
	m.verifyInput("ModifyReservedInstancesWithContext", param0)
	return m.ModifyReservedInstancesWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ModifySnapshotAttribute(param0 *ec2.ModifySnapshotAttributeInput) (*ec2.ModifySnapshotAttributeOutput, error) {
	m.addCall("ModifySnapshotAttribute")
	m.verifyInput("ModifySnapshotAttribute", param0)
	return m.ModifySnapshotAttributeFunc(param0)
}

func (m *ec2Mock) ModifySnapshotAttributeRequest(param0 *ec2.ModifySnapshotAttributeInput) (*request.Request, *ec2.ModifySnapshotAttributeOutput) {
	m.addCall("ModifySnapshotAttributeRequest")
	m.verifyInput("ModifySnapshotAttributeRequest", param0)
	return m.ModifySnapshotAttributeRequestFunc(param0)
}

func (m *ec2Mock) ModifySnapshotAttributeWithContext(param0 aws.Context, param1 *ec2.ModifySnapshotAttributeInput, param2 ...request.Option) (*ec2.ModifySnapshotAttributeOutput, error) {
	m.addCall("ModifySnapshotAttributeWithContext")
	m.verifyInput("ModifySnapshotAttributeWithContext", param0)
	return m.ModifySnapshotAttributeWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ModifySpotFleetRequest(param0 *ec2.ModifySpotFleetRequestInput) (*ec2.ModifySpotFleetRequestOutput, error) {
	m.addCall("ModifySpotFleetRequest")
	m.verifyInput("ModifySpotFleetRequest", param0)
	return m.ModifySpotFleetRequestFunc(param0)
}

func (m *ec2Mock) ModifySpotFleetRequestRequest(param0 *ec2.ModifySpotFleetRequestInput) (*request.Request, *ec2.ModifySpotFleetRequestOutput) {
	m.addCall("ModifySpotFleetRequestRequest")
	m.verifyInput("ModifySpotFleetRequestRequest", param0)
	return m.ModifySpotFleetRequestRequestFunc(param0)
}

func (m *ec2Mock) ModifySpotFleetRequestWithContext(param0 aws.Context, param1 *ec2.ModifySpotFleetRequestInput, param2 ...request.Option) (*ec2.ModifySpotFleetRequestOutput, error) {
	m.addCall("ModifySpotFleetRequestWithContext")
	m.verifyInput("ModifySpotFleetRequestWithContext", param0)
	return m.ModifySpotFleetRequestWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ModifySubnetAttribute(param0 *ec2.ModifySubnetAttributeInput) (*ec2.ModifySubnetAttributeOutput, error) {
	m.addCall("ModifySubnetAttribute")
	m.verifyInput("ModifySubnetAttribute", param0)
	return m.ModifySubnetAttributeFunc(param0)
}

func (m *ec2Mock) ModifySubnetAttributeRequest(param0 *ec2.ModifySubnetAttributeInput) (*request.Request, *ec2.ModifySubnetAttributeOutput) {
	m.addCall("ModifySubnetAttributeRequest")
	m.verifyInput("ModifySubnetAttributeRequest", param0)
	return m.ModifySubnetAttributeRequestFunc(param0)
}

func (m *ec2Mock) ModifySubnetAttributeWithContext(param0 aws.Context, param1 *ec2.ModifySubnetAttributeInput, param2 ...request.Option) (*ec2.ModifySubnetAttributeOutput, error) {
	m.addCall("ModifySubnetAttributeWithContext")
	m.verifyInput("ModifySubnetAttributeWithContext", param0)
	return m.ModifySubnetAttributeWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ModifyVolume(param0 *ec2.ModifyVolumeInput) (*ec2.ModifyVolumeOutput, error) {
	m.addCall("ModifyVolume")
	m.verifyInput("ModifyVolume", param0)
	return m.ModifyVolumeFunc(param0)
}

func (m *ec2Mock) ModifyVolumeAttribute(param0 *ec2.ModifyVolumeAttributeInput) (*ec2.ModifyVolumeAttributeOutput, error) {
	m.addCall("ModifyVolumeAttribute")
	m.verifyInput("ModifyVolumeAttribute", param0)
	return m.ModifyVolumeAttributeFunc(param0)
}

func (m *ec2Mock) ModifyVolumeAttributeRequest(param0 *ec2.ModifyVolumeAttributeInput) (*request.Request, *ec2.ModifyVolumeAttributeOutput) {
	m.addCall("ModifyVolumeAttributeRequest")
	m.verifyInput("ModifyVolumeAttributeRequest", param0)
	return m.ModifyVolumeAttributeRequestFunc(param0)
}

func (m *ec2Mock) ModifyVolumeAttributeWithContext(param0 aws.Context, param1 *ec2.ModifyVolumeAttributeInput, param2 ...request.Option) (*ec2.ModifyVolumeAttributeOutput, error) {
	m.addCall("ModifyVolumeAttributeWithContext")
	m.verifyInput("ModifyVolumeAttributeWithContext", param0)
	return m.ModifyVolumeAttributeWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ModifyVolumeRequest(param0 *ec2.ModifyVolumeInput) (*request.Request, *ec2.ModifyVolumeOutput) {
	m.addCall("ModifyVolumeRequest")
	m.verifyInput("ModifyVolumeRequest", param0)
	return m.ModifyVolumeRequestFunc(param0)
}

func (m *ec2Mock) ModifyVolumeWithContext(param0 aws.Context, param1 *ec2.ModifyVolumeInput, param2 ...request.Option) (*ec2.ModifyVolumeOutput, error) {
	m.addCall("ModifyVolumeWithContext")
	m.verifyInput("ModifyVolumeWithContext", param0)
	return m.ModifyVolumeWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ModifyVpcAttribute(param0 *ec2.ModifyVpcAttributeInput) (*ec2.ModifyVpcAttributeOutput, error) {
	m.addCall("ModifyVpcAttribute")
	m.verifyInput("ModifyVpcAttribute", param0)
	return m.ModifyVpcAttributeFunc(param0)
}

func (m *ec2Mock) ModifyVpcAttributeRequest(param0 *ec2.ModifyVpcAttributeInput) (*request.Request, *ec2.ModifyVpcAttributeOutput) {
	m.addCall("ModifyVpcAttributeRequest")
	m.verifyInput("ModifyVpcAttributeRequest", param0)
	return m.ModifyVpcAttributeRequestFunc(param0)
}

func (m *ec2Mock) ModifyVpcAttributeWithContext(param0 aws.Context, param1 *ec2.ModifyVpcAttributeInput, param2 ...request.Option) (*ec2.ModifyVpcAttributeOutput, error) {
	m.addCall("ModifyVpcAttributeWithContext")
	m.verifyInput("ModifyVpcAttributeWithContext", param0)
	return m.ModifyVpcAttributeWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ModifyVpcEndpoint(param0 *ec2.ModifyVpcEndpointInput) (*ec2.ModifyVpcEndpointOutput, error) {
	m.addCall("ModifyVpcEndpoint")
	m.verifyInput("ModifyVpcEndpoint", param0)
	return m.ModifyVpcEndpointFunc(param0)
}

func (m *ec2Mock) ModifyVpcEndpointRequest(param0 *ec2.ModifyVpcEndpointInput) (*request.Request, *ec2.ModifyVpcEndpointOutput) {
	m.addCall("ModifyVpcEndpointRequest")
	m.verifyInput("ModifyVpcEndpointRequest", param0)
	return m.ModifyVpcEndpointRequestFunc(param0)
}

func (m *ec2Mock) ModifyVpcEndpointWithContext(param0 aws.Context, param1 *ec2.ModifyVpcEndpointInput, param2 ...request.Option) (*ec2.ModifyVpcEndpointOutput, error) {
	m.addCall("ModifyVpcEndpointWithContext")
	m.verifyInput("ModifyVpcEndpointWithContext", param0)
	return m.ModifyVpcEndpointWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ModifyVpcPeeringConnectionOptions(param0 *ec2.ModifyVpcPeeringConnectionOptionsInput) (*ec2.ModifyVpcPeeringConnectionOptionsOutput, error) {
	m.addCall("ModifyVpcPeeringConnectionOptions")
	m.verifyInput("ModifyVpcPeeringConnectionOptions", param0)
	return m.ModifyVpcPeeringConnectionOptionsFunc(param0)
}

func (m *ec2Mock) ModifyVpcPeeringConnectionOptionsRequest(param0 *ec2.ModifyVpcPeeringConnectionOptionsInput) (*request.Request, *ec2.ModifyVpcPeeringConnectionOptionsOutput) {
	m.addCall("ModifyVpcPeeringConnectionOptionsRequest")
	m.verifyInput("ModifyVpcPeeringConnectionOptionsRequest", param0)
	return m.ModifyVpcPeeringConnectionOptionsRequestFunc(param0)
}

func (m *ec2Mock) ModifyVpcPeeringConnectionOptionsWithContext(param0 aws.Context, param1 *ec2.ModifyVpcPeeringConnectionOptionsInput, param2 ...request.Option) (*ec2.ModifyVpcPeeringConnectionOptionsOutput, error) {
	m.addCall("ModifyVpcPeeringConnectionOptionsWithContext")
	m.verifyInput("ModifyVpcPeeringConnectionOptionsWithContext", param0)
	return m.ModifyVpcPeeringConnectionOptionsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) MonitorInstances(param0 *ec2.MonitorInstancesInput) (*ec2.MonitorInstancesOutput, error) {
	m.addCall("MonitorInstances")
	m.verifyInput("MonitorInstances", param0)
	return m.MonitorInstancesFunc(param0)
}

func (m *ec2Mock) MonitorInstancesRequest(param0 *ec2.MonitorInstancesInput) (*request.Request, *ec2.MonitorInstancesOutput) {
	m.addCall("MonitorInstancesRequest")
	m.verifyInput("MonitorInstancesRequest", param0)
	return m.MonitorInstancesRequestFunc(param0)
}

func (m *ec2Mock) MonitorInstancesWithContext(param0 aws.Context, param1 *ec2.MonitorInstancesInput, param2 ...request.Option) (*ec2.MonitorInstancesOutput, error) {
	m.addCall("MonitorInstancesWithContext")
	m.verifyInput("MonitorInstancesWithContext", param0)
	return m.MonitorInstancesWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) MoveAddressToVpc(param0 *ec2.MoveAddressToVpcInput) (*ec2.MoveAddressToVpcOutput, error) {
	m.addCall("MoveAddressToVpc")
	m.verifyInput("MoveAddressToVpc", param0)
	return m.MoveAddressToVpcFunc(param0)
}

func (m *ec2Mock) MoveAddressToVpcRequest(param0 *ec2.MoveAddressToVpcInput) (*request.Request, *ec2.MoveAddressToVpcOutput) {
	m.addCall("MoveAddressToVpcRequest")
	m.verifyInput("MoveAddressToVpcRequest", param0)
	return m.MoveAddressToVpcRequestFunc(param0)
}

func (m *ec2Mock) MoveAddressToVpcWithContext(param0 aws.Context, param1 *ec2.MoveAddressToVpcInput, param2 ...request.Option) (*ec2.MoveAddressToVpcOutput, error) {
	m.addCall("MoveAddressToVpcWithContext")
	m.verifyInput("MoveAddressToVpcWithContext", param0)
	return m.MoveAddressToVpcWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) PurchaseHostReservation(param0 *ec2.PurchaseHostReservationInput) (*ec2.PurchaseHostReservationOutput, error) {
	m.addCall("PurchaseHostReservation")
	m.verifyInput("PurchaseHostReservation", param0)
	return m.PurchaseHostReservationFunc(param0)
}

func (m *ec2Mock) PurchaseHostReservationRequest(param0 *ec2.PurchaseHostReservationInput) (*request.Request, *ec2.PurchaseHostReservationOutput) {
	m.addCall("PurchaseHostReservationRequest")
	m.verifyInput("PurchaseHostReservationRequest", param0)
	return m.PurchaseHostReservationRequestFunc(param0)
}

func (m *ec2Mock) PurchaseHostReservationWithContext(param0 aws.Context, param1 *ec2.PurchaseHostReservationInput, param2 ...request.Option) (*ec2.PurchaseHostReservationOutput, error) {
	m.addCall("PurchaseHostReservationWithContext")
	m.verifyInput("PurchaseHostReservationWithContext", param0)
	return m.PurchaseHostReservationWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) PurchaseReservedInstancesOffering(param0 *ec2.PurchaseReservedInstancesOfferingInput) (*ec2.PurchaseReservedInstancesOfferingOutput, error) {
	m.addCall("PurchaseReservedInstancesOffering")
	m.verifyInput("PurchaseReservedInstancesOffering", param0)
	return m.PurchaseReservedInstancesOfferingFunc(param0)
}

func (m *ec2Mock) PurchaseReservedInstancesOfferingRequest(param0 *ec2.PurchaseReservedInstancesOfferingInput) (*request.Request, *ec2.PurchaseReservedInstancesOfferingOutput) {
	m.addCall("PurchaseReservedInstancesOfferingRequest")
	m.verifyInput("PurchaseReservedInstancesOfferingRequest", param0)
	return m.PurchaseReservedInstancesOfferingRequestFunc(param0)
}

func (m *ec2Mock) PurchaseReservedInstancesOfferingWithContext(param0 aws.Context, param1 *ec2.PurchaseReservedInstancesOfferingInput, param2 ...request.Option) (*ec2.PurchaseReservedInstancesOfferingOutput, error) {
	m.addCall("PurchaseReservedInstancesOfferingWithContext")
	m.verifyInput("PurchaseReservedInstancesOfferingWithContext", param0)
	return m.PurchaseReservedInstancesOfferingWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) PurchaseScheduledInstances(param0 *ec2.PurchaseScheduledInstancesInput) (*ec2.PurchaseScheduledInstancesOutput, error) {
	m.addCall("PurchaseScheduledInstances")
	m.verifyInput("PurchaseScheduledInstances", param0)
	return m.PurchaseScheduledInstancesFunc(param0)
}

func (m *ec2Mock) PurchaseScheduledInstancesRequest(param0 *ec2.PurchaseScheduledInstancesInput) (*request.Request, *ec2.PurchaseScheduledInstancesOutput) {
	m.addCall("PurchaseScheduledInstancesRequest")
	m.verifyInput("PurchaseScheduledInstancesRequest", param0)
	return m.PurchaseScheduledInstancesRequestFunc(param0)
}

func (m *ec2Mock) PurchaseScheduledInstancesWithContext(param0 aws.Context, param1 *ec2.PurchaseScheduledInstancesInput, param2 ...request.Option) (*ec2.PurchaseScheduledInstancesOutput, error) {
	m.addCall("PurchaseScheduledInstancesWithContext")
	m.verifyInput("PurchaseScheduledInstancesWithContext", param0)
	return m.PurchaseScheduledInstancesWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) RebootInstances(param0 *ec2.RebootInstancesInput) (*ec2.RebootInstancesOutput, error) {
	m.addCall("RebootInstances")
	m.verifyInput("RebootInstances", param0)
	return m.RebootInstancesFunc(param0)
}

func (m *ec2Mock) RebootInstancesRequest(param0 *ec2.RebootInstancesInput) (*request.Request, *ec2.RebootInstancesOutput) {
	m.addCall("RebootInstancesRequest")
	m.verifyInput("RebootInstancesRequest", param0)
	return m.RebootInstancesRequestFunc(param0)
}

func (m *ec2Mock) RebootInstancesWithContext(param0 aws.Context, param1 *ec2.RebootInstancesInput, param2 ...request.Option) (*ec2.RebootInstancesOutput, error) {
	m.addCall("RebootInstancesWithContext")
	m.verifyInput("RebootInstancesWithContext", param0)
	return m.RebootInstancesWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) RegisterImage(param0 *ec2.RegisterImageInput) (*ec2.RegisterImageOutput, error) {
	m.addCall("RegisterImage")
	m.verifyInput("RegisterImage", param0)
	return m.RegisterImageFunc(param0)
}

func (m *ec2Mock) RegisterImageRequest(param0 *ec2.RegisterImageInput) (*request.Request, *ec2.RegisterImageOutput) {
	m.addCall("RegisterImageRequest")
	m.verifyInput("RegisterImageRequest", param0)
	return m.RegisterImageRequestFunc(param0)
}

func (m *ec2Mock) RegisterImageWithContext(param0 aws.Context, param1 *ec2.RegisterImageInput, param2 ...request.Option) (*ec2.RegisterImageOutput, error) {
	m.addCall("RegisterImageWithContext")
	m.verifyInput("RegisterImageWithContext", param0)
	return m.RegisterImageWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) RejectVpcPeeringConnection(param0 *ec2.RejectVpcPeeringConnectionInput) (*ec2.RejectVpcPeeringConnectionOutput, error) {
	m.addCall("RejectVpcPeeringConnection")
	m.verifyInput("RejectVpcPeeringConnection", param0)
	return m.RejectVpcPeeringConnectionFunc(param0)
}

func (m *ec2Mock) RejectVpcPeeringConnectionRequest(param0 *ec2.RejectVpcPeeringConnectionInput) (*request.Request, *ec2.RejectVpcPeeringConnectionOutput) {
	m.addCall("RejectVpcPeeringConnectionRequest")
	m.verifyInput("RejectVpcPeeringConnectionRequest", param0)
	return m.RejectVpcPeeringConnectionRequestFunc(param0)
}

func (m *ec2Mock) RejectVpcPeeringConnectionWithContext(param0 aws.Context, param1 *ec2.RejectVpcPeeringConnectionInput, param2 ...request.Option) (*ec2.RejectVpcPeeringConnectionOutput, error) {
	m.addCall("RejectVpcPeeringConnectionWithContext")
	m.verifyInput("RejectVpcPeeringConnectionWithContext", param0)
	return m.RejectVpcPeeringConnectionWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ReleaseAddress(param0 *ec2.ReleaseAddressInput) (*ec2.ReleaseAddressOutput, error) {
	m.addCall("ReleaseAddress")
	m.verifyInput("ReleaseAddress", param0)
	return m.ReleaseAddressFunc(param0)
}

func (m *ec2Mock) ReleaseAddressRequest(param0 *ec2.ReleaseAddressInput) (*request.Request, *ec2.ReleaseAddressOutput) {
	m.addCall("ReleaseAddressRequest")
	m.verifyInput("ReleaseAddressRequest", param0)
	return m.ReleaseAddressRequestFunc(param0)
}

func (m *ec2Mock) ReleaseAddressWithContext(param0 aws.Context, param1 *ec2.ReleaseAddressInput, param2 ...request.Option) (*ec2.ReleaseAddressOutput, error) {
	m.addCall("ReleaseAddressWithContext")
	m.verifyInput("ReleaseAddressWithContext", param0)
	return m.ReleaseAddressWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ReleaseHosts(param0 *ec2.ReleaseHostsInput) (*ec2.ReleaseHostsOutput, error) {
	m.addCall("ReleaseHosts")
	m.verifyInput("ReleaseHosts", param0)
	return m.ReleaseHostsFunc(param0)
}

func (m *ec2Mock) ReleaseHostsRequest(param0 *ec2.ReleaseHostsInput) (*request.Request, *ec2.ReleaseHostsOutput) {
	m.addCall("ReleaseHostsRequest")
	m.verifyInput("ReleaseHostsRequest", param0)
	return m.ReleaseHostsRequestFunc(param0)
}

func (m *ec2Mock) ReleaseHostsWithContext(param0 aws.Context, param1 *ec2.ReleaseHostsInput, param2 ...request.Option) (*ec2.ReleaseHostsOutput, error) {
	m.addCall("ReleaseHostsWithContext")
	m.verifyInput("ReleaseHostsWithContext", param0)
	return m.ReleaseHostsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ReplaceIamInstanceProfileAssociation(param0 *ec2.ReplaceIamInstanceProfileAssociationInput) (*ec2.ReplaceIamInstanceProfileAssociationOutput, error) {
	m.addCall("ReplaceIamInstanceProfileAssociation")
	m.verifyInput("ReplaceIamInstanceProfileAssociation", param0)
	return m.ReplaceIamInstanceProfileAssociationFunc(param0)
}

func (m *ec2Mock) ReplaceIamInstanceProfileAssociationRequest(param0 *ec2.ReplaceIamInstanceProfileAssociationInput) (*request.Request, *ec2.ReplaceIamInstanceProfileAssociationOutput) {
	m.addCall("ReplaceIamInstanceProfileAssociationRequest")
	m.verifyInput("ReplaceIamInstanceProfileAssociationRequest", param0)
	return m.ReplaceIamInstanceProfileAssociationRequestFunc(param0)
}

func (m *ec2Mock) ReplaceIamInstanceProfileAssociationWithContext(param0 aws.Context, param1 *ec2.ReplaceIamInstanceProfileAssociationInput, param2 ...request.Option) (*ec2.ReplaceIamInstanceProfileAssociationOutput, error) {
	m.addCall("ReplaceIamInstanceProfileAssociationWithContext")
	m.verifyInput("ReplaceIamInstanceProfileAssociationWithContext", param0)
	return m.ReplaceIamInstanceProfileAssociationWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ReplaceNetworkAclAssociation(param0 *ec2.ReplaceNetworkAclAssociationInput) (*ec2.ReplaceNetworkAclAssociationOutput, error) {
	m.addCall("ReplaceNetworkAclAssociation")
	m.verifyInput("ReplaceNetworkAclAssociation", param0)
	return m.ReplaceNetworkAclAssociationFunc(param0)
}

func (m *ec2Mock) ReplaceNetworkAclAssociationRequest(param0 *ec2.ReplaceNetworkAclAssociationInput) (*request.Request, *ec2.ReplaceNetworkAclAssociationOutput) {
	m.addCall("ReplaceNetworkAclAssociationRequest")
	m.verifyInput("ReplaceNetworkAclAssociationRequest", param0)
	return m.ReplaceNetworkAclAssociationRequestFunc(param0)
}

func (m *ec2Mock) ReplaceNetworkAclAssociationWithContext(param0 aws.Context, param1 *ec2.ReplaceNetworkAclAssociationInput, param2 ...request.Option) (*ec2.ReplaceNetworkAclAssociationOutput, error) {
	m.addCall("ReplaceNetworkAclAssociationWithContext")
	m.verifyInput("ReplaceNetworkAclAssociationWithContext", param0)
	return m.ReplaceNetworkAclAssociationWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ReplaceNetworkAclEntry(param0 *ec2.ReplaceNetworkAclEntryInput) (*ec2.ReplaceNetworkAclEntryOutput, error) {
	m.addCall("ReplaceNetworkAclEntry")
	m.verifyInput("ReplaceNetworkAclEntry", param0)
	return m.ReplaceNetworkAclEntryFunc(param0)
}

func (m *ec2Mock) ReplaceNetworkAclEntryRequest(param0 *ec2.ReplaceNetworkAclEntryInput) (*request.Request, *ec2.ReplaceNetworkAclEntryOutput) {
	m.addCall("ReplaceNetworkAclEntryRequest")
	m.verifyInput("ReplaceNetworkAclEntryRequest", param0)
	return m.ReplaceNetworkAclEntryRequestFunc(param0)
}

func (m *ec2Mock) ReplaceNetworkAclEntryWithContext(param0 aws.Context, param1 *ec2.ReplaceNetworkAclEntryInput, param2 ...request.Option) (*ec2.ReplaceNetworkAclEntryOutput, error) {
	m.addCall("ReplaceNetworkAclEntryWithContext")
	m.verifyInput("ReplaceNetworkAclEntryWithContext", param0)
	return m.ReplaceNetworkAclEntryWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ReplaceRoute(param0 *ec2.ReplaceRouteInput) (*ec2.ReplaceRouteOutput, error) {
	m.addCall("ReplaceRoute")
	m.verifyInput("ReplaceRoute", param0)
	return m.ReplaceRouteFunc(param0)
}

func (m *ec2Mock) ReplaceRouteRequest(param0 *ec2.ReplaceRouteInput) (*request.Request, *ec2.ReplaceRouteOutput) {
	m.addCall("ReplaceRouteRequest")
	m.verifyInput("ReplaceRouteRequest", param0)
	return m.ReplaceRouteRequestFunc(param0)
}

func (m *ec2Mock) ReplaceRouteTableAssociation(param0 *ec2.ReplaceRouteTableAssociationInput) (*ec2.ReplaceRouteTableAssociationOutput, error) {
	m.addCall("ReplaceRouteTableAssociation")
	m.verifyInput("ReplaceRouteTableAssociation", param0)
	return m.ReplaceRouteTableAssociationFunc(param0)
}

func (m *ec2Mock) ReplaceRouteTableAssociationRequest(param0 *ec2.ReplaceRouteTableAssociationInput) (*request.Request, *ec2.ReplaceRouteTableAssociationOutput) {
	m.addCall("ReplaceRouteTableAssociationRequest")
	m.verifyInput("ReplaceRouteTableAssociationRequest", param0)
	return m.ReplaceRouteTableAssociationRequestFunc(param0)
}

func (m *ec2Mock) ReplaceRouteTableAssociationWithContext(param0 aws.Context, param1 *ec2.ReplaceRouteTableAssociationInput, param2 ...request.Option) (*ec2.ReplaceRouteTableAssociationOutput, error) {
	m.addCall("ReplaceRouteTableAssociationWithContext")
	m.verifyInput("ReplaceRouteTableAssociationWithContext", param0)
	return m.ReplaceRouteTableAssociationWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ReplaceRouteWithContext(param0 aws.Context, param1 *ec2.ReplaceRouteInput, param2 ...request.Option) (*ec2.ReplaceRouteOutput, error) {
	m.addCall("ReplaceRouteWithContext")
	m.verifyInput("ReplaceRouteWithContext", param0)
	return m.ReplaceRouteWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ReportInstanceStatus(param0 *ec2.ReportInstanceStatusInput) (*ec2.ReportInstanceStatusOutput, error) {
	m.addCall("ReportInstanceStatus")
	m.verifyInput("ReportInstanceStatus", param0)
	return m.ReportInstanceStatusFunc(param0)
}

func (m *ec2Mock) ReportInstanceStatusRequest(param0 *ec2.ReportInstanceStatusInput) (*request.Request, *ec2.ReportInstanceStatusOutput) {
	m.addCall("ReportInstanceStatusRequest")
	m.verifyInput("ReportInstanceStatusRequest", param0)
	return m.ReportInstanceStatusRequestFunc(param0)
}

func (m *ec2Mock) ReportInstanceStatusWithContext(param0 aws.Context, param1 *ec2.ReportInstanceStatusInput, param2 ...request.Option) (*ec2.ReportInstanceStatusOutput, error) {
	m.addCall("ReportInstanceStatusWithContext")
	m.verifyInput("ReportInstanceStatusWithContext", param0)
	return m.ReportInstanceStatusWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) RequestSpotFleet(param0 *ec2.RequestSpotFleetInput) (*ec2.RequestSpotFleetOutput, error) {
	m.addCall("RequestSpotFleet")
	m.verifyInput("RequestSpotFleet", param0)
	return m.RequestSpotFleetFunc(param0)
}

func (m *ec2Mock) RequestSpotFleetRequest(param0 *ec2.RequestSpotFleetInput) (*request.Request, *ec2.RequestSpotFleetOutput) {
	m.addCall("RequestSpotFleetRequest")
	m.verifyInput("RequestSpotFleetRequest", param0)
	return m.RequestSpotFleetRequestFunc(param0)
}

func (m *ec2Mock) RequestSpotFleetWithContext(param0 aws.Context, param1 *ec2.RequestSpotFleetInput, param2 ...request.Option) (*ec2.RequestSpotFleetOutput, error) {
	m.addCall("RequestSpotFleetWithContext")
	m.verifyInput("RequestSpotFleetWithContext", param0)
	return m.RequestSpotFleetWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) RequestSpotInstances(param0 *ec2.RequestSpotInstancesInput) (*ec2.RequestSpotInstancesOutput, error) {
	m.addCall("RequestSpotInstances")
	m.verifyInput("RequestSpotInstances", param0)
	return m.RequestSpotInstancesFunc(param0)
}

func (m *ec2Mock) RequestSpotInstancesRequest(param0 *ec2.RequestSpotInstancesInput) (*request.Request, *ec2.RequestSpotInstancesOutput) {
	m.addCall("RequestSpotInstancesRequest")
	m.verifyInput("RequestSpotInstancesRequest", param0)
	return m.RequestSpotInstancesRequestFunc(param0)
}

func (m *ec2Mock) RequestSpotInstancesWithContext(param0 aws.Context, param1 *ec2.RequestSpotInstancesInput, param2 ...request.Option) (*ec2.RequestSpotInstancesOutput, error) {
	m.addCall("RequestSpotInstancesWithContext")
	m.verifyInput("RequestSpotInstancesWithContext", param0)
	return m.RequestSpotInstancesWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ResetFpgaImageAttribute(param0 *ec2.ResetFpgaImageAttributeInput) (*ec2.ResetFpgaImageAttributeOutput, error) {
	m.addCall("ResetFpgaImageAttribute")
	m.verifyInput("ResetFpgaImageAttribute", param0)
	return m.ResetFpgaImageAttributeFunc(param0)
}

func (m *ec2Mock) ResetFpgaImageAttributeRequest(param0 *ec2.ResetFpgaImageAttributeInput) (*request.Request, *ec2.ResetFpgaImageAttributeOutput) {
	m.addCall("ResetFpgaImageAttributeRequest")
	m.verifyInput("ResetFpgaImageAttributeRequest", param0)
	return m.ResetFpgaImageAttributeRequestFunc(param0)
}

func (m *ec2Mock) ResetFpgaImageAttributeWithContext(param0 aws.Context, param1 *ec2.ResetFpgaImageAttributeInput, param2 ...request.Option) (*ec2.ResetFpgaImageAttributeOutput, error) {
	m.addCall("ResetFpgaImageAttributeWithContext")
	m.verifyInput("ResetFpgaImageAttributeWithContext", param0)
	return m.ResetFpgaImageAttributeWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ResetImageAttribute(param0 *ec2.ResetImageAttributeInput) (*ec2.ResetImageAttributeOutput, error) {
	m.addCall("ResetImageAttribute")
	m.verifyInput("ResetImageAttribute", param0)
	return m.ResetImageAttributeFunc(param0)
}

func (m *ec2Mock) ResetImageAttributeRequest(param0 *ec2.ResetImageAttributeInput) (*request.Request, *ec2.ResetImageAttributeOutput) {
	m.addCall("ResetImageAttributeRequest")
	m.verifyInput("ResetImageAttributeRequest", param0)
	return m.ResetImageAttributeRequestFunc(param0)
}

func (m *ec2Mock) ResetImageAttributeWithContext(param0 aws.Context, param1 *ec2.ResetImageAttributeInput, param2 ...request.Option) (*ec2.ResetImageAttributeOutput, error) {
	m.addCall("ResetImageAttributeWithContext")
	m.verifyInput("ResetImageAttributeWithContext", param0)
	return m.ResetImageAttributeWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ResetInstanceAttribute(param0 *ec2.ResetInstanceAttributeInput) (*ec2.ResetInstanceAttributeOutput, error) {
	m.addCall("ResetInstanceAttribute")
	m.verifyInput("ResetInstanceAttribute", param0)
	return m.ResetInstanceAttributeFunc(param0)
}

func (m *ec2Mock) ResetInstanceAttributeRequest(param0 *ec2.ResetInstanceAttributeInput) (*request.Request, *ec2.ResetInstanceAttributeOutput) {
	m.addCall("ResetInstanceAttributeRequest")
	m.verifyInput("ResetInstanceAttributeRequest", param0)
	return m.ResetInstanceAttributeRequestFunc(param0)
}

func (m *ec2Mock) ResetInstanceAttributeWithContext(param0 aws.Context, param1 *ec2.ResetInstanceAttributeInput, param2 ...request.Option) (*ec2.ResetInstanceAttributeOutput, error) {
	m.addCall("ResetInstanceAttributeWithContext")
	m.verifyInput("ResetInstanceAttributeWithContext", param0)
	return m.ResetInstanceAttributeWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ResetNetworkInterfaceAttribute(param0 *ec2.ResetNetworkInterfaceAttributeInput) (*ec2.ResetNetworkInterfaceAttributeOutput, error) {
	m.addCall("ResetNetworkInterfaceAttribute")
	m.verifyInput("ResetNetworkInterfaceAttribute", param0)
	return m.ResetNetworkInterfaceAttributeFunc(param0)
}

func (m *ec2Mock) ResetNetworkInterfaceAttributeRequest(param0 *ec2.ResetNetworkInterfaceAttributeInput) (*request.Request, *ec2.ResetNetworkInterfaceAttributeOutput) {
	m.addCall("ResetNetworkInterfaceAttributeRequest")
	m.verifyInput("ResetNetworkInterfaceAttributeRequest", param0)
	return m.ResetNetworkInterfaceAttributeRequestFunc(param0)
}

func (m *ec2Mock) ResetNetworkInterfaceAttributeWithContext(param0 aws.Context, param1 *ec2.ResetNetworkInterfaceAttributeInput, param2 ...request.Option) (*ec2.ResetNetworkInterfaceAttributeOutput, error) {
	m.addCall("ResetNetworkInterfaceAttributeWithContext")
	m.verifyInput("ResetNetworkInterfaceAttributeWithContext", param0)
	return m.ResetNetworkInterfaceAttributeWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) ResetSnapshotAttribute(param0 *ec2.ResetSnapshotAttributeInput) (*ec2.ResetSnapshotAttributeOutput, error) {
	m.addCall("ResetSnapshotAttribute")
	m.verifyInput("ResetSnapshotAttribute", param0)
	return m.ResetSnapshotAttributeFunc(param0)
}

func (m *ec2Mock) ResetSnapshotAttributeRequest(param0 *ec2.ResetSnapshotAttributeInput) (*request.Request, *ec2.ResetSnapshotAttributeOutput) {
	m.addCall("ResetSnapshotAttributeRequest")
	m.verifyInput("ResetSnapshotAttributeRequest", param0)
	return m.ResetSnapshotAttributeRequestFunc(param0)
}

func (m *ec2Mock) ResetSnapshotAttributeWithContext(param0 aws.Context, param1 *ec2.ResetSnapshotAttributeInput, param2 ...request.Option) (*ec2.ResetSnapshotAttributeOutput, error) {
	m.addCall("ResetSnapshotAttributeWithContext")
	m.verifyInput("ResetSnapshotAttributeWithContext", param0)
	return m.ResetSnapshotAttributeWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) RestoreAddressToClassic(param0 *ec2.RestoreAddressToClassicInput) (*ec2.RestoreAddressToClassicOutput, error) {
	m.addCall("RestoreAddressToClassic")
	m.verifyInput("RestoreAddressToClassic", param0)
	return m.RestoreAddressToClassicFunc(param0)
}

func (m *ec2Mock) RestoreAddressToClassicRequest(param0 *ec2.RestoreAddressToClassicInput) (*request.Request, *ec2.RestoreAddressToClassicOutput) {
	m.addCall("RestoreAddressToClassicRequest")
	m.verifyInput("RestoreAddressToClassicRequest", param0)
	return m.RestoreAddressToClassicRequestFunc(param0)
}

func (m *ec2Mock) RestoreAddressToClassicWithContext(param0 aws.Context, param1 *ec2.RestoreAddressToClassicInput, param2 ...request.Option) (*ec2.RestoreAddressToClassicOutput, error) {
	m.addCall("RestoreAddressToClassicWithContext")
	m.verifyInput("RestoreAddressToClassicWithContext", param0)
	return m.RestoreAddressToClassicWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) RevokeSecurityGroupEgress(param0 *ec2.RevokeSecurityGroupEgressInput) (*ec2.RevokeSecurityGroupEgressOutput, error) {
	m.addCall("RevokeSecurityGroupEgress")
	m.verifyInput("RevokeSecurityGroupEgress", param0)
	return m.RevokeSecurityGroupEgressFunc(param0)
}

func (m *ec2Mock) RevokeSecurityGroupEgressRequest(param0 *ec2.RevokeSecurityGroupEgressInput) (*request.Request, *ec2.RevokeSecurityGroupEgressOutput) {
	m.addCall("RevokeSecurityGroupEgressRequest")
	m.verifyInput("RevokeSecurityGroupEgressRequest", param0)
	return m.RevokeSecurityGroupEgressRequestFunc(param0)
}

func (m *ec2Mock) RevokeSecurityGroupEgressWithContext(param0 aws.Context, param1 *ec2.RevokeSecurityGroupEgressInput, param2 ...request.Option) (*ec2.RevokeSecurityGroupEgressOutput, error) {
	m.addCall("RevokeSecurityGroupEgressWithContext")
	m.verifyInput("RevokeSecurityGroupEgressWithContext", param0)
	return m.RevokeSecurityGroupEgressWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) RevokeSecurityGroupIngress(param0 *ec2.RevokeSecurityGroupIngressInput) (*ec2.RevokeSecurityGroupIngressOutput, error) {
	m.addCall("RevokeSecurityGroupIngress")
	m.verifyInput("RevokeSecurityGroupIngress", param0)
	return m.RevokeSecurityGroupIngressFunc(param0)
}

func (m *ec2Mock) RevokeSecurityGroupIngressRequest(param0 *ec2.RevokeSecurityGroupIngressInput) (*request.Request, *ec2.RevokeSecurityGroupIngressOutput) {
	m.addCall("RevokeSecurityGroupIngressRequest")
	m.verifyInput("RevokeSecurityGroupIngressRequest", param0)
	return m.RevokeSecurityGroupIngressRequestFunc(param0)
}

func (m *ec2Mock) RevokeSecurityGroupIngressWithContext(param0 aws.Context, param1 *ec2.RevokeSecurityGroupIngressInput, param2 ...request.Option) (*ec2.RevokeSecurityGroupIngressOutput, error) {
	m.addCall("RevokeSecurityGroupIngressWithContext")
	m.verifyInput("RevokeSecurityGroupIngressWithContext", param0)
	return m.RevokeSecurityGroupIngressWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) RunInstances(param0 *ec2.RunInstancesInput) (*ec2.Reservation, error) {
	m.addCall("RunInstances")
	m.verifyInput("RunInstances", param0)
	return m.RunInstancesFunc(param0)
}

func (m *ec2Mock) RunInstancesRequest(param0 *ec2.RunInstancesInput) (*request.Request, *ec2.Reservation) {
	m.addCall("RunInstancesRequest")
	m.verifyInput("RunInstancesRequest", param0)
	return m.RunInstancesRequestFunc(param0)
}

func (m *ec2Mock) RunInstancesWithContext(param0 aws.Context, param1 *ec2.RunInstancesInput, param2 ...request.Option) (*ec2.Reservation, error) {
	m.addCall("RunInstancesWithContext")
	m.verifyInput("RunInstancesWithContext", param0)
	return m.RunInstancesWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) RunScheduledInstances(param0 *ec2.RunScheduledInstancesInput) (*ec2.RunScheduledInstancesOutput, error) {
	m.addCall("RunScheduledInstances")
	m.verifyInput("RunScheduledInstances", param0)
	return m.RunScheduledInstancesFunc(param0)
}

func (m *ec2Mock) RunScheduledInstancesRequest(param0 *ec2.RunScheduledInstancesInput) (*request.Request, *ec2.RunScheduledInstancesOutput) {
	m.addCall("RunScheduledInstancesRequest")
	m.verifyInput("RunScheduledInstancesRequest", param0)
	return m.RunScheduledInstancesRequestFunc(param0)
}

func (m *ec2Mock) RunScheduledInstancesWithContext(param0 aws.Context, param1 *ec2.RunScheduledInstancesInput, param2 ...request.Option) (*ec2.RunScheduledInstancesOutput, error) {
	m.addCall("RunScheduledInstancesWithContext")
	m.verifyInput("RunScheduledInstancesWithContext", param0)
	return m.RunScheduledInstancesWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) StartInstances(param0 *ec2.StartInstancesInput) (*ec2.StartInstancesOutput, error) {
	m.addCall("StartInstances")
	m.verifyInput("StartInstances", param0)
	return m.StartInstancesFunc(param0)
}

func (m *ec2Mock) StartInstancesRequest(param0 *ec2.StartInstancesInput) (*request.Request, *ec2.StartInstancesOutput) {
	m.addCall("StartInstancesRequest")
	m.verifyInput("StartInstancesRequest", param0)
	return m.StartInstancesRequestFunc(param0)
}

func (m *ec2Mock) StartInstancesWithContext(param0 aws.Context, param1 *ec2.StartInstancesInput, param2 ...request.Option) (*ec2.StartInstancesOutput, error) {
	m.addCall("StartInstancesWithContext")
	m.verifyInput("StartInstancesWithContext", param0)
	return m.StartInstancesWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) StopInstances(param0 *ec2.StopInstancesInput) (*ec2.StopInstancesOutput, error) {
	m.addCall("StopInstances")
	m.verifyInput("StopInstances", param0)
	return m.StopInstancesFunc(param0)
}

func (m *ec2Mock) StopInstancesRequest(param0 *ec2.StopInstancesInput) (*request.Request, *ec2.StopInstancesOutput) {
	m.addCall("StopInstancesRequest")
	m.verifyInput("StopInstancesRequest", param0)
	return m.StopInstancesRequestFunc(param0)
}

func (m *ec2Mock) StopInstancesWithContext(param0 aws.Context, param1 *ec2.StopInstancesInput, param2 ...request.Option) (*ec2.StopInstancesOutput, error) {
	m.addCall("StopInstancesWithContext")
	m.verifyInput("StopInstancesWithContext", param0)
	return m.StopInstancesWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) TerminateInstances(param0 *ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error) {
	m.addCall("TerminateInstances")
	m.verifyInput("TerminateInstances", param0)
	return m.TerminateInstancesFunc(param0)
}

func (m *ec2Mock) TerminateInstancesRequest(param0 *ec2.TerminateInstancesInput) (*request.Request, *ec2.TerminateInstancesOutput) {
	m.addCall("TerminateInstancesRequest")
	m.verifyInput("TerminateInstancesRequest", param0)
	return m.TerminateInstancesRequestFunc(param0)
}

func (m *ec2Mock) TerminateInstancesWithContext(param0 aws.Context, param1 *ec2.TerminateInstancesInput, param2 ...request.Option) (*ec2.TerminateInstancesOutput, error) {
	m.addCall("TerminateInstancesWithContext")
	m.verifyInput("TerminateInstancesWithContext", param0)
	return m.TerminateInstancesWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) UnassignIpv6Addresses(param0 *ec2.UnassignIpv6AddressesInput) (*ec2.UnassignIpv6AddressesOutput, error) {
	m.addCall("UnassignIpv6Addresses")
	m.verifyInput("UnassignIpv6Addresses", param0)
	return m.UnassignIpv6AddressesFunc(param0)
}

func (m *ec2Mock) UnassignIpv6AddressesRequest(param0 *ec2.UnassignIpv6AddressesInput) (*request.Request, *ec2.UnassignIpv6AddressesOutput) {
	m.addCall("UnassignIpv6AddressesRequest")
	m.verifyInput("UnassignIpv6AddressesRequest", param0)
	return m.UnassignIpv6AddressesRequestFunc(param0)
}

func (m *ec2Mock) UnassignIpv6AddressesWithContext(param0 aws.Context, param1 *ec2.UnassignIpv6AddressesInput, param2 ...request.Option) (*ec2.UnassignIpv6AddressesOutput, error) {
	m.addCall("UnassignIpv6AddressesWithContext")
	m.verifyInput("UnassignIpv6AddressesWithContext", param0)
	return m.UnassignIpv6AddressesWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) UnassignPrivateIpAddresses(param0 *ec2.UnassignPrivateIpAddressesInput) (*ec2.UnassignPrivateIpAddressesOutput, error) {
	m.addCall("UnassignPrivateIpAddresses")
	m.verifyInput("UnassignPrivateIpAddresses", param0)
	return m.UnassignPrivateIpAddressesFunc(param0)
}

func (m *ec2Mock) UnassignPrivateIpAddressesRequest(param0 *ec2.UnassignPrivateIpAddressesInput) (*request.Request, *ec2.UnassignPrivateIpAddressesOutput) {
	m.addCall("UnassignPrivateIpAddressesRequest")
	m.verifyInput("UnassignPrivateIpAddressesRequest", param0)
	return m.UnassignPrivateIpAddressesRequestFunc(param0)
}

func (m *ec2Mock) UnassignPrivateIpAddressesWithContext(param0 aws.Context, param1 *ec2.UnassignPrivateIpAddressesInput, param2 ...request.Option) (*ec2.UnassignPrivateIpAddressesOutput, error) {
	m.addCall("UnassignPrivateIpAddressesWithContext")
	m.verifyInput("UnassignPrivateIpAddressesWithContext", param0)
	return m.UnassignPrivateIpAddressesWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) UnmonitorInstances(param0 *ec2.UnmonitorInstancesInput) (*ec2.UnmonitorInstancesOutput, error) {
	m.addCall("UnmonitorInstances")
	m.verifyInput("UnmonitorInstances", param0)
	return m.UnmonitorInstancesFunc(param0)
}

func (m *ec2Mock) UnmonitorInstancesRequest(param0 *ec2.UnmonitorInstancesInput) (*request.Request, *ec2.UnmonitorInstancesOutput) {
	m.addCall("UnmonitorInstancesRequest")
	m.verifyInput("UnmonitorInstancesRequest", param0)
	return m.UnmonitorInstancesRequestFunc(param0)
}

func (m *ec2Mock) UnmonitorInstancesWithContext(param0 aws.Context, param1 *ec2.UnmonitorInstancesInput, param2 ...request.Option) (*ec2.UnmonitorInstancesOutput, error) {
	m.addCall("UnmonitorInstancesWithContext")
	m.verifyInput("UnmonitorInstancesWithContext", param0)
	return m.UnmonitorInstancesWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) UpdateSecurityGroupRuleDescriptionsEgress(param0 *ec2.UpdateSecurityGroupRuleDescriptionsEgressInput) (*ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput, error) {
	m.addCall("UpdateSecurityGroupRuleDescriptionsEgress")
	m.verifyInput("UpdateSecurityGroupRuleDescriptionsEgress", param0)
	return m.UpdateSecurityGroupRuleDescriptionsEgressFunc(param0)
}

func (m *ec2Mock) UpdateSecurityGroupRuleDescriptionsEgressRequest(param0 *ec2.UpdateSecurityGroupRuleDescriptionsEgressInput) (*request.Request, *ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput) {
	m.addCall("UpdateSecurityGroupRuleDescriptionsEgressRequest")
	m.verifyInput("UpdateSecurityGroupRuleDescriptionsEgressRequest", param0)
	return m.UpdateSecurityGroupRuleDescriptionsEgressRequestFunc(param0)
}

func (m *ec2Mock) UpdateSecurityGroupRuleDescriptionsEgressWithContext(param0 aws.Context, param1 *ec2.UpdateSecurityGroupRuleDescriptionsEgressInput, param2 ...request.Option) (*ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput, error) {
	m.addCall("UpdateSecurityGroupRuleDescriptionsEgressWithContext")
	m.verifyInput("UpdateSecurityGroupRuleDescriptionsEgressWithContext", param0)
	return m.UpdateSecurityGroupRuleDescriptionsEgressWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) UpdateSecurityGroupRuleDescriptionsIngress(param0 *ec2.UpdateSecurityGroupRuleDescriptionsIngressInput) (*ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput, error) {
	m.addCall("UpdateSecurityGroupRuleDescriptionsIngress")
	m.verifyInput("UpdateSecurityGroupRuleDescriptionsIngress", param0)
	return m.UpdateSecurityGroupRuleDescriptionsIngressFunc(param0)
}

func (m *ec2Mock) UpdateSecurityGroupRuleDescriptionsIngressRequest(param0 *ec2.UpdateSecurityGroupRuleDescriptionsIngressInput) (*request.Request, *ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput) {
	m.addCall("UpdateSecurityGroupRuleDescriptionsIngressRequest")
	m.verifyInput("UpdateSecurityGroupRuleDescriptionsIngressRequest", param0)
	return m.UpdateSecurityGroupRuleDescriptionsIngressRequestFunc(param0)
}

func (m *ec2Mock) UpdateSecurityGroupRuleDescriptionsIngressWithContext(param0 aws.Context, param1 *ec2.UpdateSecurityGroupRuleDescriptionsIngressInput, param2 ...request.Option) (*ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput, error) {
	m.addCall("UpdateSecurityGroupRuleDescriptionsIngressWithContext")
	m.verifyInput("UpdateSecurityGroupRuleDescriptionsIngressWithContext", param0)
	return m.UpdateSecurityGroupRuleDescriptionsIngressWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) WaitUntilBundleTaskComplete(param0 *ec2.DescribeBundleTasksInput) error {
	m.addCall("WaitUntilBundleTaskComplete")
	m.verifyInput("WaitUntilBundleTaskComplete", param0)
	return m.WaitUntilBundleTaskCompleteFunc(param0)
}

func (m *ec2Mock) WaitUntilBundleTaskCompleteWithContext(param0 aws.Context, param1 *ec2.DescribeBundleTasksInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilBundleTaskCompleteWithContext")
	m.verifyInput("WaitUntilBundleTaskCompleteWithContext", param0)
	return m.WaitUntilBundleTaskCompleteWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) WaitUntilConversionTaskCancelled(param0 *ec2.DescribeConversionTasksInput) error {
	m.addCall("WaitUntilConversionTaskCancelled")
	m.verifyInput("WaitUntilConversionTaskCancelled", param0)
	return m.WaitUntilConversionTaskCancelledFunc(param0)
}

func (m *ec2Mock) WaitUntilConversionTaskCancelledWithContext(param0 aws.Context, param1 *ec2.DescribeConversionTasksInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilConversionTaskCancelledWithContext")
	m.verifyInput("WaitUntilConversionTaskCancelledWithContext", param0)
	return m.WaitUntilConversionTaskCancelledWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) WaitUntilConversionTaskCompleted(param0 *ec2.DescribeConversionTasksInput) error {
	m.addCall("WaitUntilConversionTaskCompleted")
	m.verifyInput("WaitUntilConversionTaskCompleted", param0)
	return m.WaitUntilConversionTaskCompletedFunc(param0)
}

func (m *ec2Mock) WaitUntilConversionTaskCompletedWithContext(param0 aws.Context, param1 *ec2.DescribeConversionTasksInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilConversionTaskCompletedWithContext")
	m.verifyInput("WaitUntilConversionTaskCompletedWithContext", param0)
	return m.WaitUntilConversionTaskCompletedWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) WaitUntilConversionTaskDeleted(param0 *ec2.DescribeConversionTasksInput) error {
	m.addCall("WaitUntilConversionTaskDeleted")
	m.verifyInput("WaitUntilConversionTaskDeleted", param0)
	return m.WaitUntilConversionTaskDeletedFunc(param0)
}

func (m *ec2Mock) WaitUntilConversionTaskDeletedWithContext(param0 aws.Context, param1 *ec2.DescribeConversionTasksInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilConversionTaskDeletedWithContext")
	m.verifyInput("WaitUntilConversionTaskDeletedWithContext", param0)
	return m.WaitUntilConversionTaskDeletedWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) WaitUntilCustomerGatewayAvailable(param0 *ec2.DescribeCustomerGatewaysInput) error {
	m.addCall("WaitUntilCustomerGatewayAvailable")
	m.verifyInput("WaitUntilCustomerGatewayAvailable", param0)
	return m.WaitUntilCustomerGatewayAvailableFunc(param0)
}

func (m *ec2Mock) WaitUntilCustomerGatewayAvailableWithContext(param0 aws.Context, param1 *ec2.DescribeCustomerGatewaysInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilCustomerGatewayAvailableWithContext")
	m.verifyInput("WaitUntilCustomerGatewayAvailableWithContext", param0)
	return m.WaitUntilCustomerGatewayAvailableWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) WaitUntilExportTaskCancelled(param0 *ec2.DescribeExportTasksInput) error {
	m.addCall("WaitUntilExportTaskCancelled")
	m.verifyInput("WaitUntilExportTaskCancelled", param0)
	return m.WaitUntilExportTaskCancelledFunc(param0)
}

func (m *ec2Mock) WaitUntilExportTaskCancelledWithContext(param0 aws.Context, param1 *ec2.DescribeExportTasksInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilExportTaskCancelledWithContext")
	m.verifyInput("WaitUntilExportTaskCancelledWithContext", param0)
	return m.WaitUntilExportTaskCancelledWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) WaitUntilExportTaskCompleted(param0 *ec2.DescribeExportTasksInput) error {
	m.addCall("WaitUntilExportTaskCompleted")
	m.verifyInput("WaitUntilExportTaskCompleted", param0)
	return m.WaitUntilExportTaskCompletedFunc(param0)
}

func (m *ec2Mock) WaitUntilExportTaskCompletedWithContext(param0 aws.Context, param1 *ec2.DescribeExportTasksInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilExportTaskCompletedWithContext")
	m.verifyInput("WaitUntilExportTaskCompletedWithContext", param0)
	return m.WaitUntilExportTaskCompletedWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) WaitUntilImageAvailable(param0 *ec2.DescribeImagesInput) error {
	m.addCall("WaitUntilImageAvailable")
	m.verifyInput("WaitUntilImageAvailable", param0)
	return m.WaitUntilImageAvailableFunc(param0)
}

func (m *ec2Mock) WaitUntilImageAvailableWithContext(param0 aws.Context, param1 *ec2.DescribeImagesInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilImageAvailableWithContext")
	m.verifyInput("WaitUntilImageAvailableWithContext", param0)
	return m.WaitUntilImageAvailableWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) WaitUntilImageExists(param0 *ec2.DescribeImagesInput) error {
	m.addCall("WaitUntilImageExists")
	m.verifyInput("WaitUntilImageExists", param0)
	return m.WaitUntilImageExistsFunc(param0)
}

func (m *ec2Mock) WaitUntilImageExistsWithContext(param0 aws.Context, param1 *ec2.DescribeImagesInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilImageExistsWithContext")
	m.verifyInput("WaitUntilImageExistsWithContext", param0)
	return m.WaitUntilImageExistsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) WaitUntilInstanceExists(param0 *ec2.DescribeInstancesInput) error {
	m.addCall("WaitUntilInstanceExists")
	m.verifyInput("WaitUntilInstanceExists", param0)
	return m.WaitUntilInstanceExistsFunc(param0)
}

func (m *ec2Mock) WaitUntilInstanceExistsWithContext(param0 aws.Context, param1 *ec2.DescribeInstancesInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilInstanceExistsWithContext")
	m.verifyInput("WaitUntilInstanceExistsWithContext", param0)
	return m.WaitUntilInstanceExistsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) WaitUntilInstanceRunning(param0 *ec2.DescribeInstancesInput) error {
	m.addCall("WaitUntilInstanceRunning")
	m.verifyInput("WaitUntilInstanceRunning", param0)
	return m.WaitUntilInstanceRunningFunc(param0)
}

func (m *ec2Mock) WaitUntilInstanceRunningWithContext(param0 aws.Context, param1 *ec2.DescribeInstancesInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilInstanceRunningWithContext")
	m.verifyInput("WaitUntilInstanceRunningWithContext", param0)
	return m.WaitUntilInstanceRunningWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) WaitUntilInstanceStatusOk(param0 *ec2.DescribeInstanceStatusInput) error {
	m.addCall("WaitUntilInstanceStatusOk")
	m.verifyInput("WaitUntilInstanceStatusOk", param0)
	return m.WaitUntilInstanceStatusOkFunc(param0)
}

func (m *ec2Mock) WaitUntilInstanceStatusOkWithContext(param0 aws.Context, param1 *ec2.DescribeInstanceStatusInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilInstanceStatusOkWithContext")
	m.verifyInput("WaitUntilInstanceStatusOkWithContext", param0)
	return m.WaitUntilInstanceStatusOkWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) WaitUntilInstanceStopped(param0 *ec2.DescribeInstancesInput) error {
	m.addCall("WaitUntilInstanceStopped")
	m.verifyInput("WaitUntilInstanceStopped", param0)
	return m.WaitUntilInstanceStoppedFunc(param0)
}

func (m *ec2Mock) WaitUntilInstanceStoppedWithContext(param0 aws.Context, param1 *ec2.DescribeInstancesInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilInstanceStoppedWithContext")
	m.verifyInput("WaitUntilInstanceStoppedWithContext", param0)
	return m.WaitUntilInstanceStoppedWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) WaitUntilInstanceTerminated(param0 *ec2.DescribeInstancesInput) error {
	m.addCall("WaitUntilInstanceTerminated")
	m.verifyInput("WaitUntilInstanceTerminated", param0)
	return m.WaitUntilInstanceTerminatedFunc(param0)
}

func (m *ec2Mock) WaitUntilInstanceTerminatedWithContext(param0 aws.Context, param1 *ec2.DescribeInstancesInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilInstanceTerminatedWithContext")
	m.verifyInput("WaitUntilInstanceTerminatedWithContext", param0)
	return m.WaitUntilInstanceTerminatedWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) WaitUntilKeyPairExists(param0 *ec2.DescribeKeyPairsInput) error {
	m.addCall("WaitUntilKeyPairExists")
	m.verifyInput("WaitUntilKeyPairExists", param0)
	return m.WaitUntilKeyPairExistsFunc(param0)
}

func (m *ec2Mock) WaitUntilKeyPairExistsWithContext(param0 aws.Context, param1 *ec2.DescribeKeyPairsInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilKeyPairExistsWithContext")
	m.verifyInput("WaitUntilKeyPairExistsWithContext", param0)
	return m.WaitUntilKeyPairExistsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) WaitUntilNatGatewayAvailable(param0 *ec2.DescribeNatGatewaysInput) error {
	m.addCall("WaitUntilNatGatewayAvailable")
	m.verifyInput("WaitUntilNatGatewayAvailable", param0)
	return m.WaitUntilNatGatewayAvailableFunc(param0)
}

func (m *ec2Mock) WaitUntilNatGatewayAvailableWithContext(param0 aws.Context, param1 *ec2.DescribeNatGatewaysInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilNatGatewayAvailableWithContext")
	m.verifyInput("WaitUntilNatGatewayAvailableWithContext", param0)
	return m.WaitUntilNatGatewayAvailableWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) WaitUntilNetworkInterfaceAvailable(param0 *ec2.DescribeNetworkInterfacesInput) error {
	m.addCall("WaitUntilNetworkInterfaceAvailable")
	m.verifyInput("WaitUntilNetworkInterfaceAvailable", param0)
	return m.WaitUntilNetworkInterfaceAvailableFunc(param0)
}

func (m *ec2Mock) WaitUntilNetworkInterfaceAvailableWithContext(param0 aws.Context, param1 *ec2.DescribeNetworkInterfacesInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilNetworkInterfaceAvailableWithContext")
	m.verifyInput("WaitUntilNetworkInterfaceAvailableWithContext", param0)
	return m.WaitUntilNetworkInterfaceAvailableWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) WaitUntilPasswordDataAvailable(param0 *ec2.GetPasswordDataInput) error {
	m.addCall("WaitUntilPasswordDataAvailable")
	m.verifyInput("WaitUntilPasswordDataAvailable", param0)
	return m.WaitUntilPasswordDataAvailableFunc(param0)
}

func (m *ec2Mock) WaitUntilPasswordDataAvailableWithContext(param0 aws.Context, param1 *ec2.GetPasswordDataInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilPasswordDataAvailableWithContext")
	m.verifyInput("WaitUntilPasswordDataAvailableWithContext", param0)
	return m.WaitUntilPasswordDataAvailableWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) WaitUntilSnapshotCompleted(param0 *ec2.DescribeSnapshotsInput) error {
	m.addCall("WaitUntilSnapshotCompleted")
	m.verifyInput("WaitUntilSnapshotCompleted", param0)
	return m.WaitUntilSnapshotCompletedFunc(param0)
}

func (m *ec2Mock) WaitUntilSnapshotCompletedWithContext(param0 aws.Context, param1 *ec2.DescribeSnapshotsInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilSnapshotCompletedWithContext")
	m.verifyInput("WaitUntilSnapshotCompletedWithContext", param0)
	return m.WaitUntilSnapshotCompletedWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) WaitUntilSpotInstanceRequestFulfilled(param0 *ec2.DescribeSpotInstanceRequestsInput) error {
	m.addCall("WaitUntilSpotInstanceRequestFulfilled")
	m.verifyInput("WaitUntilSpotInstanceRequestFulfilled", param0)
	return m.WaitUntilSpotInstanceRequestFulfilledFunc(param0)
}

func (m *ec2Mock) WaitUntilSpotInstanceRequestFulfilledWithContext(param0 aws.Context, param1 *ec2.DescribeSpotInstanceRequestsInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilSpotInstanceRequestFulfilledWithContext")
	m.verifyInput("WaitUntilSpotInstanceRequestFulfilledWithContext", param0)
	return m.WaitUntilSpotInstanceRequestFulfilledWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) WaitUntilSubnetAvailable(param0 *ec2.DescribeSubnetsInput) error {
	m.addCall("WaitUntilSubnetAvailable")
	m.verifyInput("WaitUntilSubnetAvailable", param0)
	return m.WaitUntilSubnetAvailableFunc(param0)
}

func (m *ec2Mock) WaitUntilSubnetAvailableWithContext(param0 aws.Context, param1 *ec2.DescribeSubnetsInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilSubnetAvailableWithContext")
	m.verifyInput("WaitUntilSubnetAvailableWithContext", param0)
	return m.WaitUntilSubnetAvailableWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) WaitUntilSystemStatusOk(param0 *ec2.DescribeInstanceStatusInput) error {
	m.addCall("WaitUntilSystemStatusOk")
	m.verifyInput("WaitUntilSystemStatusOk", param0)
	return m.WaitUntilSystemStatusOkFunc(param0)
}

func (m *ec2Mock) WaitUntilSystemStatusOkWithContext(param0 aws.Context, param1 *ec2.DescribeInstanceStatusInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilSystemStatusOkWithContext")
	m.verifyInput("WaitUntilSystemStatusOkWithContext", param0)
	return m.WaitUntilSystemStatusOkWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) WaitUntilVolumeAvailable(param0 *ec2.DescribeVolumesInput) error {
	m.addCall("WaitUntilVolumeAvailable")
	m.verifyInput("WaitUntilVolumeAvailable", param0)
	return m.WaitUntilVolumeAvailableFunc(param0)
}

func (m *ec2Mock) WaitUntilVolumeAvailableWithContext(param0 aws.Context, param1 *ec2.DescribeVolumesInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilVolumeAvailableWithContext")
	m.verifyInput("WaitUntilVolumeAvailableWithContext", param0)
	return m.WaitUntilVolumeAvailableWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) WaitUntilVolumeDeleted(param0 *ec2.DescribeVolumesInput) error {
	m.addCall("WaitUntilVolumeDeleted")
	m.verifyInput("WaitUntilVolumeDeleted", param0)
	return m.WaitUntilVolumeDeletedFunc(param0)
}

func (m *ec2Mock) WaitUntilVolumeDeletedWithContext(param0 aws.Context, param1 *ec2.DescribeVolumesInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilVolumeDeletedWithContext")
	m.verifyInput("WaitUntilVolumeDeletedWithContext", param0)
	return m.WaitUntilVolumeDeletedWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) WaitUntilVolumeInUse(param0 *ec2.DescribeVolumesInput) error {
	m.addCall("WaitUntilVolumeInUse")
	m.verifyInput("WaitUntilVolumeInUse", param0)
	return m.WaitUntilVolumeInUseFunc(param0)
}

func (m *ec2Mock) WaitUntilVolumeInUseWithContext(param0 aws.Context, param1 *ec2.DescribeVolumesInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilVolumeInUseWithContext")
	m.verifyInput("WaitUntilVolumeInUseWithContext", param0)
	return m.WaitUntilVolumeInUseWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) WaitUntilVpcAvailable(param0 *ec2.DescribeVpcsInput) error {
	m.addCall("WaitUntilVpcAvailable")
	m.verifyInput("WaitUntilVpcAvailable", param0)
	return m.WaitUntilVpcAvailableFunc(param0)
}

func (m *ec2Mock) WaitUntilVpcAvailableWithContext(param0 aws.Context, param1 *ec2.DescribeVpcsInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilVpcAvailableWithContext")
	m.verifyInput("WaitUntilVpcAvailableWithContext", param0)
	return m.WaitUntilVpcAvailableWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) WaitUntilVpcExists(param0 *ec2.DescribeVpcsInput) error {
	m.addCall("WaitUntilVpcExists")
	m.verifyInput("WaitUntilVpcExists", param0)
	return m.WaitUntilVpcExistsFunc(param0)
}

func (m *ec2Mock) WaitUntilVpcExistsWithContext(param0 aws.Context, param1 *ec2.DescribeVpcsInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilVpcExistsWithContext")
	m.verifyInput("WaitUntilVpcExistsWithContext", param0)
	return m.WaitUntilVpcExistsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) WaitUntilVpcPeeringConnectionDeleted(param0 *ec2.DescribeVpcPeeringConnectionsInput) error {
	m.addCall("WaitUntilVpcPeeringConnectionDeleted")
	m.verifyInput("WaitUntilVpcPeeringConnectionDeleted", param0)
	return m.WaitUntilVpcPeeringConnectionDeletedFunc(param0)
}

func (m *ec2Mock) WaitUntilVpcPeeringConnectionDeletedWithContext(param0 aws.Context, param1 *ec2.DescribeVpcPeeringConnectionsInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilVpcPeeringConnectionDeletedWithContext")
	m.verifyInput("WaitUntilVpcPeeringConnectionDeletedWithContext", param0)
	return m.WaitUntilVpcPeeringConnectionDeletedWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) WaitUntilVpcPeeringConnectionExists(param0 *ec2.DescribeVpcPeeringConnectionsInput) error {
	m.addCall("WaitUntilVpcPeeringConnectionExists")
	m.verifyInput("WaitUntilVpcPeeringConnectionExists", param0)
	return m.WaitUntilVpcPeeringConnectionExistsFunc(param0)
}

func (m *ec2Mock) WaitUntilVpcPeeringConnectionExistsWithContext(param0 aws.Context, param1 *ec2.DescribeVpcPeeringConnectionsInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilVpcPeeringConnectionExistsWithContext")
	m.verifyInput("WaitUntilVpcPeeringConnectionExistsWithContext", param0)
	return m.WaitUntilVpcPeeringConnectionExistsWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) WaitUntilVpnConnectionAvailable(param0 *ec2.DescribeVpnConnectionsInput) error {
	m.addCall("WaitUntilVpnConnectionAvailable")
	m.verifyInput("WaitUntilVpnConnectionAvailable", param0)
	return m.WaitUntilVpnConnectionAvailableFunc(param0)
}

func (m *ec2Mock) WaitUntilVpnConnectionAvailableWithContext(param0 aws.Context, param1 *ec2.DescribeVpnConnectionsInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilVpnConnectionAvailableWithContext")
	m.verifyInput("WaitUntilVpnConnectionAvailableWithContext", param0)
	return m.WaitUntilVpnConnectionAvailableWithContextFunc(param0, param1, param2...)
}

func (m *ec2Mock) WaitUntilVpnConnectionDeleted(param0 *ec2.DescribeVpnConnectionsInput) error {
	m.addCall("WaitUntilVpnConnectionDeleted")
	m.verifyInput("WaitUntilVpnConnectionDeleted", param0)
	return m.WaitUntilVpnConnectionDeletedFunc(param0)
}

func (m *ec2Mock) WaitUntilVpnConnectionDeletedWithContext(param0 aws.Context, param1 *ec2.DescribeVpnConnectionsInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilVpnConnectionDeletedWithContext")
	m.verifyInput("WaitUntilVpnConnectionDeletedWithContext", param0)
	return m.WaitUntilVpnConnectionDeletedWithContextFunc(param0, param1, param2...)
}

type iamMock struct {
	basicMock
	iamiface.IAMAPI
	AddClientIDToOpenIDConnectProviderFunc                 func(param0 *iam.AddClientIDToOpenIDConnectProviderInput) (*iam.AddClientIDToOpenIDConnectProviderOutput, error)
	AddClientIDToOpenIDConnectProviderRequestFunc          func(param0 *iam.AddClientIDToOpenIDConnectProviderInput) (*request.Request, *iam.AddClientIDToOpenIDConnectProviderOutput)
	AddClientIDToOpenIDConnectProviderWithContextFunc      func(param0 aws.Context, param1 *iam.AddClientIDToOpenIDConnectProviderInput, param2 ...request.Option) (*iam.AddClientIDToOpenIDConnectProviderOutput, error)
	AddRoleToInstanceProfileFunc                           func(param0 *iam.AddRoleToInstanceProfileInput) (*iam.AddRoleToInstanceProfileOutput, error)
	AddRoleToInstanceProfileRequestFunc                    func(param0 *iam.AddRoleToInstanceProfileInput) (*request.Request, *iam.AddRoleToInstanceProfileOutput)
	AddRoleToInstanceProfileWithContextFunc                func(param0 aws.Context, param1 *iam.AddRoleToInstanceProfileInput, param2 ...request.Option) (*iam.AddRoleToInstanceProfileOutput, error)
	AddUserToGroupFunc                                     func(param0 *iam.AddUserToGroupInput) (*iam.AddUserToGroupOutput, error)
	AddUserToGroupRequestFunc                              func(param0 *iam.AddUserToGroupInput) (*request.Request, *iam.AddUserToGroupOutput)
	AddUserToGroupWithContextFunc                          func(param0 aws.Context, param1 *iam.AddUserToGroupInput, param2 ...request.Option) (*iam.AddUserToGroupOutput, error)
	AttachGroupPolicyFunc                                  func(param0 *iam.AttachGroupPolicyInput) (*iam.AttachGroupPolicyOutput, error)
	AttachGroupPolicyRequestFunc                           func(param0 *iam.AttachGroupPolicyInput) (*request.Request, *iam.AttachGroupPolicyOutput)
	AttachGroupPolicyWithContextFunc                       func(param0 aws.Context, param1 *iam.AttachGroupPolicyInput, param2 ...request.Option) (*iam.AttachGroupPolicyOutput, error)
	AttachRolePolicyFunc                                   func(param0 *iam.AttachRolePolicyInput) (*iam.AttachRolePolicyOutput, error)
	AttachRolePolicyRequestFunc                            func(param0 *iam.AttachRolePolicyInput) (*request.Request, *iam.AttachRolePolicyOutput)
	AttachRolePolicyWithContextFunc                        func(param0 aws.Context, param1 *iam.AttachRolePolicyInput, param2 ...request.Option) (*iam.AttachRolePolicyOutput, error)
	AttachUserPolicyFunc                                   func(param0 *iam.AttachUserPolicyInput) (*iam.AttachUserPolicyOutput, error)
	AttachUserPolicyRequestFunc                            func(param0 *iam.AttachUserPolicyInput) (*request.Request, *iam.AttachUserPolicyOutput)
	AttachUserPolicyWithContextFunc                        func(param0 aws.Context, param1 *iam.AttachUserPolicyInput, param2 ...request.Option) (*iam.AttachUserPolicyOutput, error)
	ChangePasswordFunc                                     func(param0 *iam.ChangePasswordInput) (*iam.ChangePasswordOutput, error)
	ChangePasswordRequestFunc                              func(param0 *iam.ChangePasswordInput) (*request.Request, *iam.ChangePasswordOutput)
	ChangePasswordWithContextFunc                          func(param0 aws.Context, param1 *iam.ChangePasswordInput, param2 ...request.Option) (*iam.ChangePasswordOutput, error)
	CreateAccessKeyFunc                                    func(param0 *iam.CreateAccessKeyInput) (*iam.CreateAccessKeyOutput, error)
	CreateAccessKeyRequestFunc                             func(param0 *iam.CreateAccessKeyInput) (*request.Request, *iam.CreateAccessKeyOutput)
	CreateAccessKeyWithContextFunc                         func(param0 aws.Context, param1 *iam.CreateAccessKeyInput, param2 ...request.Option) (*iam.CreateAccessKeyOutput, error)
	CreateAccountAliasFunc                                 func(param0 *iam.CreateAccountAliasInput) (*iam.CreateAccountAliasOutput, error)
	CreateAccountAliasRequestFunc                          func(param0 *iam.CreateAccountAliasInput) (*request.Request, *iam.CreateAccountAliasOutput)
	CreateAccountAliasWithContextFunc                      func(param0 aws.Context, param1 *iam.CreateAccountAliasInput, param2 ...request.Option) (*iam.CreateAccountAliasOutput, error)
	CreateGroupFunc                                        func(param0 *iam.CreateGroupInput) (*iam.CreateGroupOutput, error)
	CreateGroupRequestFunc                                 func(param0 *iam.CreateGroupInput) (*request.Request, *iam.CreateGroupOutput)
	CreateGroupWithContextFunc                             func(param0 aws.Context, param1 *iam.CreateGroupInput, param2 ...request.Option) (*iam.CreateGroupOutput, error)
	CreateInstanceProfileFunc                              func(param0 *iam.CreateInstanceProfileInput) (*iam.CreateInstanceProfileOutput, error)
	CreateInstanceProfileRequestFunc                       func(param0 *iam.CreateInstanceProfileInput) (*request.Request, *iam.CreateInstanceProfileOutput)
	CreateInstanceProfileWithContextFunc                   func(param0 aws.Context, param1 *iam.CreateInstanceProfileInput, param2 ...request.Option) (*iam.CreateInstanceProfileOutput, error)
	CreateLoginProfileFunc                                 func(param0 *iam.CreateLoginProfileInput) (*iam.CreateLoginProfileOutput, error)
	CreateLoginProfileRequestFunc                          func(param0 *iam.CreateLoginProfileInput) (*request.Request, *iam.CreateLoginProfileOutput)
	CreateLoginProfileWithContextFunc                      func(param0 aws.Context, param1 *iam.CreateLoginProfileInput, param2 ...request.Option) (*iam.CreateLoginProfileOutput, error)
	CreateOpenIDConnectProviderFunc                        func(param0 *iam.CreateOpenIDConnectProviderInput) (*iam.CreateOpenIDConnectProviderOutput, error)
	CreateOpenIDConnectProviderRequestFunc                 func(param0 *iam.CreateOpenIDConnectProviderInput) (*request.Request, *iam.CreateOpenIDConnectProviderOutput)
	CreateOpenIDConnectProviderWithContextFunc             func(param0 aws.Context, param1 *iam.CreateOpenIDConnectProviderInput, param2 ...request.Option) (*iam.CreateOpenIDConnectProviderOutput, error)
	CreatePolicyFunc                                       func(param0 *iam.CreatePolicyInput) (*iam.CreatePolicyOutput, error)
	CreatePolicyRequestFunc                                func(param0 *iam.CreatePolicyInput) (*request.Request, *iam.CreatePolicyOutput)
	CreatePolicyVersionFunc                                func(param0 *iam.CreatePolicyVersionInput) (*iam.CreatePolicyVersionOutput, error)
	CreatePolicyVersionRequestFunc                         func(param0 *iam.CreatePolicyVersionInput) (*request.Request, *iam.CreatePolicyVersionOutput)
	CreatePolicyVersionWithContextFunc                     func(param0 aws.Context, param1 *iam.CreatePolicyVersionInput, param2 ...request.Option) (*iam.CreatePolicyVersionOutput, error)
	CreatePolicyWithContextFunc                            func(param0 aws.Context, param1 *iam.CreatePolicyInput, param2 ...request.Option) (*iam.CreatePolicyOutput, error)
	CreateRoleFunc                                         func(param0 *iam.CreateRoleInput) (*iam.CreateRoleOutput, error)
	CreateRoleRequestFunc                                  func(param0 *iam.CreateRoleInput) (*request.Request, *iam.CreateRoleOutput)
	CreateRoleWithContextFunc                              func(param0 aws.Context, param1 *iam.CreateRoleInput, param2 ...request.Option) (*iam.CreateRoleOutput, error)
	CreateSAMLProviderFunc                                 func(param0 *iam.CreateSAMLProviderInput) (*iam.CreateSAMLProviderOutput, error)
	CreateSAMLProviderRequestFunc                          func(param0 *iam.CreateSAMLProviderInput) (*request.Request, *iam.CreateSAMLProviderOutput)
	CreateSAMLProviderWithContextFunc                      func(param0 aws.Context, param1 *iam.CreateSAMLProviderInput, param2 ...request.Option) (*iam.CreateSAMLProviderOutput, error)
	CreateServiceLinkedRoleFunc                            func(param0 *iam.CreateServiceLinkedRoleInput) (*iam.CreateServiceLinkedRoleOutput, error)
	CreateServiceLinkedRoleRequestFunc                     func(param0 *iam.CreateServiceLinkedRoleInput) (*request.Request, *iam.CreateServiceLinkedRoleOutput)
	CreateServiceLinkedRoleWithContextFunc                 func(param0 aws.Context, param1 *iam.CreateServiceLinkedRoleInput, param2 ...request.Option) (*iam.CreateServiceLinkedRoleOutput, error)
	CreateServiceSpecificCredentialFunc                    func(param0 *iam.CreateServiceSpecificCredentialInput) (*iam.CreateServiceSpecificCredentialOutput, error)
	CreateServiceSpecificCredentialRequestFunc             func(param0 *iam.CreateServiceSpecificCredentialInput) (*request.Request, *iam.CreateServiceSpecificCredentialOutput)
	CreateServiceSpecificCredentialWithContextFunc         func(param0 aws.Context, param1 *iam.CreateServiceSpecificCredentialInput, param2 ...request.Option) (*iam.CreateServiceSpecificCredentialOutput, error)
	CreateUserFunc                                         func(param0 *iam.CreateUserInput) (*iam.CreateUserOutput, error)
	CreateUserRequestFunc                                  func(param0 *iam.CreateUserInput) (*request.Request, *iam.CreateUserOutput)
	CreateUserWithContextFunc                              func(param0 aws.Context, param1 *iam.CreateUserInput, param2 ...request.Option) (*iam.CreateUserOutput, error)
	CreateVirtualMFADeviceFunc                             func(param0 *iam.CreateVirtualMFADeviceInput) (*iam.CreateVirtualMFADeviceOutput, error)
	CreateVirtualMFADeviceRequestFunc                      func(param0 *iam.CreateVirtualMFADeviceInput) (*request.Request, *iam.CreateVirtualMFADeviceOutput)
	CreateVirtualMFADeviceWithContextFunc                  func(param0 aws.Context, param1 *iam.CreateVirtualMFADeviceInput, param2 ...request.Option) (*iam.CreateVirtualMFADeviceOutput, error)
	DeactivateMFADeviceFunc                                func(param0 *iam.DeactivateMFADeviceInput) (*iam.DeactivateMFADeviceOutput, error)
	DeactivateMFADeviceRequestFunc                         func(param0 *iam.DeactivateMFADeviceInput) (*request.Request, *iam.DeactivateMFADeviceOutput)
	DeactivateMFADeviceWithContextFunc                     func(param0 aws.Context, param1 *iam.DeactivateMFADeviceInput, param2 ...request.Option) (*iam.DeactivateMFADeviceOutput, error)
	DeleteAccessKeyFunc                                    func(param0 *iam.DeleteAccessKeyInput) (*iam.DeleteAccessKeyOutput, error)
	DeleteAccessKeyRequestFunc                             func(param0 *iam.DeleteAccessKeyInput) (*request.Request, *iam.DeleteAccessKeyOutput)
	DeleteAccessKeyWithContextFunc                         func(param0 aws.Context, param1 *iam.DeleteAccessKeyInput, param2 ...request.Option) (*iam.DeleteAccessKeyOutput, error)
	DeleteAccountAliasFunc                                 func(param0 *iam.DeleteAccountAliasInput) (*iam.DeleteAccountAliasOutput, error)
	DeleteAccountAliasRequestFunc                          func(param0 *iam.DeleteAccountAliasInput) (*request.Request, *iam.DeleteAccountAliasOutput)
	DeleteAccountAliasWithContextFunc                      func(param0 aws.Context, param1 *iam.DeleteAccountAliasInput, param2 ...request.Option) (*iam.DeleteAccountAliasOutput, error)
	DeleteAccountPasswordPolicyFunc                        func(param0 *iam.DeleteAccountPasswordPolicyInput) (*iam.DeleteAccountPasswordPolicyOutput, error)
	DeleteAccountPasswordPolicyRequestFunc                 func(param0 *iam.DeleteAccountPasswordPolicyInput) (*request.Request, *iam.DeleteAccountPasswordPolicyOutput)
	DeleteAccountPasswordPolicyWithContextFunc             func(param0 aws.Context, param1 *iam.DeleteAccountPasswordPolicyInput, param2 ...request.Option) (*iam.DeleteAccountPasswordPolicyOutput, error)
	DeleteGroupFunc                                        func(param0 *iam.DeleteGroupInput) (*iam.DeleteGroupOutput, error)
	DeleteGroupPolicyFunc                                  func(param0 *iam.DeleteGroupPolicyInput) (*iam.DeleteGroupPolicyOutput, error)
	DeleteGroupPolicyRequestFunc                           func(param0 *iam.DeleteGroupPolicyInput) (*request.Request, *iam.DeleteGroupPolicyOutput)
	DeleteGroupPolicyWithContextFunc                       func(param0 aws.Context, param1 *iam.DeleteGroupPolicyInput, param2 ...request.Option) (*iam.DeleteGroupPolicyOutput, error)
	DeleteGroupRequestFunc                                 func(param0 *iam.DeleteGroupInput) (*request.Request, *iam.DeleteGroupOutput)
	DeleteGroupWithContextFunc                             func(param0 aws.Context, param1 *iam.DeleteGroupInput, param2 ...request.Option) (*iam.DeleteGroupOutput, error)
	DeleteInstanceProfileFunc                              func(param0 *iam.DeleteInstanceProfileInput) (*iam.DeleteInstanceProfileOutput, error)
	DeleteInstanceProfileRequestFunc                       func(param0 *iam.DeleteInstanceProfileInput) (*request.Request, *iam.DeleteInstanceProfileOutput)
	DeleteInstanceProfileWithContextFunc                   func(param0 aws.Context, param1 *iam.DeleteInstanceProfileInput, param2 ...request.Option) (*iam.DeleteInstanceProfileOutput, error)
	DeleteLoginProfileFunc                                 func(param0 *iam.DeleteLoginProfileInput) (*iam.DeleteLoginProfileOutput, error)
	DeleteLoginProfileRequestFunc                          func(param0 *iam.DeleteLoginProfileInput) (*request.Request, *iam.DeleteLoginProfileOutput)
	DeleteLoginProfileWithContextFunc                      func(param0 aws.Context, param1 *iam.DeleteLoginProfileInput, param2 ...request.Option) (*iam.DeleteLoginProfileOutput, error)
	DeleteOpenIDConnectProviderFunc                        func(param0 *iam.DeleteOpenIDConnectProviderInput) (*iam.DeleteOpenIDConnectProviderOutput, error)
	DeleteOpenIDConnectProviderRequestFunc                 func(param0 *iam.DeleteOpenIDConnectProviderInput) (*request.Request, *iam.DeleteOpenIDConnectProviderOutput)
	DeleteOpenIDConnectProviderWithContextFunc             func(param0 aws.Context, param1 *iam.DeleteOpenIDConnectProviderInput, param2 ...request.Option) (*iam.DeleteOpenIDConnectProviderOutput, error)
	DeletePolicyFunc                                       func(param0 *iam.DeletePolicyInput) (*iam.DeletePolicyOutput, error)
	DeletePolicyRequestFunc                                func(param0 *iam.DeletePolicyInput) (*request.Request, *iam.DeletePolicyOutput)
	DeletePolicyVersionFunc                                func(param0 *iam.DeletePolicyVersionInput) (*iam.DeletePolicyVersionOutput, error)
	DeletePolicyVersionRequestFunc                         func(param0 *iam.DeletePolicyVersionInput) (*request.Request, *iam.DeletePolicyVersionOutput)
	DeletePolicyVersionWithContextFunc                     func(param0 aws.Context, param1 *iam.DeletePolicyVersionInput, param2 ...request.Option) (*iam.DeletePolicyVersionOutput, error)
	DeletePolicyWithContextFunc                            func(param0 aws.Context, param1 *iam.DeletePolicyInput, param2 ...request.Option) (*iam.DeletePolicyOutput, error)
	DeleteRoleFunc                                         func(param0 *iam.DeleteRoleInput) (*iam.DeleteRoleOutput, error)
	DeleteRolePolicyFunc                                   func(param0 *iam.DeleteRolePolicyInput) (*iam.DeleteRolePolicyOutput, error)
	DeleteRolePolicyRequestFunc                            func(param0 *iam.DeleteRolePolicyInput) (*request.Request, *iam.DeleteRolePolicyOutput)
	DeleteRolePolicyWithContextFunc                        func(param0 aws.Context, param1 *iam.DeleteRolePolicyInput, param2 ...request.Option) (*iam.DeleteRolePolicyOutput, error)
	DeleteRoleRequestFunc                                  func(param0 *iam.DeleteRoleInput) (*request.Request, *iam.DeleteRoleOutput)
	DeleteRoleWithContextFunc                              func(param0 aws.Context, param1 *iam.DeleteRoleInput, param2 ...request.Option) (*iam.DeleteRoleOutput, error)
	DeleteSAMLProviderFunc                                 func(param0 *iam.DeleteSAMLProviderInput) (*iam.DeleteSAMLProviderOutput, error)
	DeleteSAMLProviderRequestFunc                          func(param0 *iam.DeleteSAMLProviderInput) (*request.Request, *iam.DeleteSAMLProviderOutput)
	DeleteSAMLProviderWithContextFunc                      func(param0 aws.Context, param1 *iam.DeleteSAMLProviderInput, param2 ...request.Option) (*iam.DeleteSAMLProviderOutput, error)
	DeleteSSHPublicKeyFunc                                 func(param0 *iam.DeleteSSHPublicKeyInput) (*iam.DeleteSSHPublicKeyOutput, error)
	DeleteSSHPublicKeyRequestFunc                          func(param0 *iam.DeleteSSHPublicKeyInput) (*request.Request, *iam.DeleteSSHPublicKeyOutput)
	DeleteSSHPublicKeyWithContextFunc                      func(param0 aws.Context, param1 *iam.DeleteSSHPublicKeyInput, param2 ...request.Option) (*iam.DeleteSSHPublicKeyOutput, error)
	DeleteServerCertificateFunc                            func(param0 *iam.DeleteServerCertificateInput) (*iam.DeleteServerCertificateOutput, error)
	DeleteServerCertificateRequestFunc                     func(param0 *iam.DeleteServerCertificateInput) (*request.Request, *iam.DeleteServerCertificateOutput)
	DeleteServerCertificateWithContextFunc                 func(param0 aws.Context, param1 *iam.DeleteServerCertificateInput, param2 ...request.Option) (*iam.DeleteServerCertificateOutput, error)
	DeleteServiceLinkedRoleFunc                            func(param0 *iam.DeleteServiceLinkedRoleInput) (*iam.DeleteServiceLinkedRoleOutput, error)
	DeleteServiceLinkedRoleRequestFunc                     func(param0 *iam.DeleteServiceLinkedRoleInput) (*request.Request, *iam.DeleteServiceLinkedRoleOutput)
	DeleteServiceLinkedRoleWithContextFunc                 func(param0 aws.Context, param1 *iam.DeleteServiceLinkedRoleInput, param2 ...request.Option) (*iam.DeleteServiceLinkedRoleOutput, error)
	DeleteServiceSpecificCredentialFunc                    func(param0 *iam.DeleteServiceSpecificCredentialInput) (*iam.DeleteServiceSpecificCredentialOutput, error)
	DeleteServiceSpecificCredentialRequestFunc             func(param0 *iam.DeleteServiceSpecificCredentialInput) (*request.Request, *iam.DeleteServiceSpecificCredentialOutput)
	DeleteServiceSpecificCredentialWithContextFunc         func(param0 aws.Context, param1 *iam.DeleteServiceSpecificCredentialInput, param2 ...request.Option) (*iam.DeleteServiceSpecificCredentialOutput, error)
	DeleteSigningCertificateFunc                           func(param0 *iam.DeleteSigningCertificateInput) (*iam.DeleteSigningCertificateOutput, error)
	DeleteSigningCertificateRequestFunc                    func(param0 *iam.DeleteSigningCertificateInput) (*request.Request, *iam.DeleteSigningCertificateOutput)
	DeleteSigningCertificateWithContextFunc                func(param0 aws.Context, param1 *iam.DeleteSigningCertificateInput, param2 ...request.Option) (*iam.DeleteSigningCertificateOutput, error)
	DeleteUserFunc                                         func(param0 *iam.DeleteUserInput) (*iam.DeleteUserOutput, error)
	DeleteUserPolicyFunc                                   func(param0 *iam.DeleteUserPolicyInput) (*iam.DeleteUserPolicyOutput, error)
	DeleteUserPolicyRequestFunc                            func(param0 *iam.DeleteUserPolicyInput) (*request.Request, *iam.DeleteUserPolicyOutput)
	DeleteUserPolicyWithContextFunc                        func(param0 aws.Context, param1 *iam.DeleteUserPolicyInput, param2 ...request.Option) (*iam.DeleteUserPolicyOutput, error)
	DeleteUserRequestFunc                                  func(param0 *iam.DeleteUserInput) (*request.Request, *iam.DeleteUserOutput)
	DeleteUserWithContextFunc                              func(param0 aws.Context, param1 *iam.DeleteUserInput, param2 ...request.Option) (*iam.DeleteUserOutput, error)
	DeleteVirtualMFADeviceFunc                             func(param0 *iam.DeleteVirtualMFADeviceInput) (*iam.DeleteVirtualMFADeviceOutput, error)
	DeleteVirtualMFADeviceRequestFunc                      func(param0 *iam.DeleteVirtualMFADeviceInput) (*request.Request, *iam.DeleteVirtualMFADeviceOutput)
	DeleteVirtualMFADeviceWithContextFunc                  func(param0 aws.Context, param1 *iam.DeleteVirtualMFADeviceInput, param2 ...request.Option) (*iam.DeleteVirtualMFADeviceOutput, error)
	DetachGroupPolicyFunc                                  func(param0 *iam.DetachGroupPolicyInput) (*iam.DetachGroupPolicyOutput, error)
	DetachGroupPolicyRequestFunc                           func(param0 *iam.DetachGroupPolicyInput) (*request.Request, *iam.DetachGroupPolicyOutput)
	DetachGroupPolicyWithContextFunc                       func(param0 aws.Context, param1 *iam.DetachGroupPolicyInput, param2 ...request.Option) (*iam.DetachGroupPolicyOutput, error)
	DetachRolePolicyFunc                                   func(param0 *iam.DetachRolePolicyInput) (*iam.DetachRolePolicyOutput, error)
	DetachRolePolicyRequestFunc                            func(param0 *iam.DetachRolePolicyInput) (*request.Request, *iam.DetachRolePolicyOutput)
	DetachRolePolicyWithContextFunc                        func(param0 aws.Context, param1 *iam.DetachRolePolicyInput, param2 ...request.Option) (*iam.DetachRolePolicyOutput, error)
	DetachUserPolicyFunc                                   func(param0 *iam.DetachUserPolicyInput) (*iam.DetachUserPolicyOutput, error)
	DetachUserPolicyRequestFunc                            func(param0 *iam.DetachUserPolicyInput) (*request.Request, *iam.DetachUserPolicyOutput)
	DetachUserPolicyWithContextFunc                        func(param0 aws.Context, param1 *iam.DetachUserPolicyInput, param2 ...request.Option) (*iam.DetachUserPolicyOutput, error)
	EnableMFADeviceFunc                                    func(param0 *iam.EnableMFADeviceInput) (*iam.EnableMFADeviceOutput, error)
	EnableMFADeviceRequestFunc                             func(param0 *iam.EnableMFADeviceInput) (*request.Request, *iam.EnableMFADeviceOutput)
	EnableMFADeviceWithContextFunc                         func(param0 aws.Context, param1 *iam.EnableMFADeviceInput, param2 ...request.Option) (*iam.EnableMFADeviceOutput, error)
	GenerateCredentialReportFunc                           func(param0 *iam.GenerateCredentialReportInput) (*iam.GenerateCredentialReportOutput, error)
	GenerateCredentialReportRequestFunc                    func(param0 *iam.GenerateCredentialReportInput) (*request.Request, *iam.GenerateCredentialReportOutput)
	GenerateCredentialReportWithContextFunc                func(param0 aws.Context, param1 *iam.GenerateCredentialReportInput, param2 ...request.Option) (*iam.GenerateCredentialReportOutput, error)
	GetAccessKeyLastUsedFunc                               func(param0 *iam.GetAccessKeyLastUsedInput) (*iam.GetAccessKeyLastUsedOutput, error)
	GetAccessKeyLastUsedRequestFunc                        func(param0 *iam.GetAccessKeyLastUsedInput) (*request.Request, *iam.GetAccessKeyLastUsedOutput)
	GetAccessKeyLastUsedWithContextFunc                    func(param0 aws.Context, param1 *iam.GetAccessKeyLastUsedInput, param2 ...request.Option) (*iam.GetAccessKeyLastUsedOutput, error)
	GetAccountAuthorizationDetailsFunc                     func(param0 *iam.GetAccountAuthorizationDetailsInput) (*iam.GetAccountAuthorizationDetailsOutput, error)
	GetAccountAuthorizationDetailsRequestFunc              func(param0 *iam.GetAccountAuthorizationDetailsInput) (*request.Request, *iam.GetAccountAuthorizationDetailsOutput)
	GetAccountAuthorizationDetailsWithContextFunc          func(param0 aws.Context, param1 *iam.GetAccountAuthorizationDetailsInput, param2 ...request.Option) (*iam.GetAccountAuthorizationDetailsOutput, error)
	GetAccountPasswordPolicyFunc                           func(param0 *iam.GetAccountPasswordPolicyInput) (*iam.GetAccountPasswordPolicyOutput, error)
	GetAccountPasswordPolicyRequestFunc                    func(param0 *iam.GetAccountPasswordPolicyInput) (*request.Request, *iam.GetAccountPasswordPolicyOutput)
	GetAccountPasswordPolicyWithContextFunc                func(param0 aws.Context, param1 *iam.GetAccountPasswordPolicyInput, param2 ...request.Option) (*iam.GetAccountPasswordPolicyOutput, error)
	GetAccountSummaryFunc                                  func(param0 *iam.GetAccountSummaryInput) (*iam.GetAccountSummaryOutput, error)
	GetAccountSummaryRequestFunc                           func(param0 *iam.GetAccountSummaryInput) (*request.Request, *iam.GetAccountSummaryOutput)
	GetAccountSummaryWithContextFunc                       func(param0 aws.Context, param1 *iam.GetAccountSummaryInput, param2 ...request.Option) (*iam.GetAccountSummaryOutput, error)
	GetContextKeysForCustomPolicyFunc                      func(param0 *iam.GetContextKeysForCustomPolicyInput) (*iam.GetContextKeysForPolicyResponse, error)
	GetContextKeysForCustomPolicyRequestFunc               func(param0 *iam.GetContextKeysForCustomPolicyInput) (*request.Request, *iam.GetContextKeysForPolicyResponse)
	GetContextKeysForCustomPolicyWithContextFunc           func(param0 aws.Context, param1 *iam.GetContextKeysForCustomPolicyInput, param2 ...request.Option) (*iam.GetContextKeysForPolicyResponse, error)
	GetContextKeysForPrincipalPolicyFunc                   func(param0 *iam.GetContextKeysForPrincipalPolicyInput) (*iam.GetContextKeysForPolicyResponse, error)
	GetContextKeysForPrincipalPolicyRequestFunc            func(param0 *iam.GetContextKeysForPrincipalPolicyInput) (*request.Request, *iam.GetContextKeysForPolicyResponse)
	GetContextKeysForPrincipalPolicyWithContextFunc        func(param0 aws.Context, param1 *iam.GetContextKeysForPrincipalPolicyInput, param2 ...request.Option) (*iam.GetContextKeysForPolicyResponse, error)
	GetCredentialReportFunc                                func(param0 *iam.GetCredentialReportInput) (*iam.GetCredentialReportOutput, error)
	GetCredentialReportRequestFunc                         func(param0 *iam.GetCredentialReportInput) (*request.Request, *iam.GetCredentialReportOutput)
	GetCredentialReportWithContextFunc                     func(param0 aws.Context, param1 *iam.GetCredentialReportInput, param2 ...request.Option) (*iam.GetCredentialReportOutput, error)
	GetGroupFunc                                           func(param0 *iam.GetGroupInput) (*iam.GetGroupOutput, error)
	GetGroupPolicyFunc                                     func(param0 *iam.GetGroupPolicyInput) (*iam.GetGroupPolicyOutput, error)
	GetGroupPolicyRequestFunc                              func(param0 *iam.GetGroupPolicyInput) (*request.Request, *iam.GetGroupPolicyOutput)
	GetGroupPolicyWithContextFunc                          func(param0 aws.Context, param1 *iam.GetGroupPolicyInput, param2 ...request.Option) (*iam.GetGroupPolicyOutput, error)
	GetGroupRequestFunc                                    func(param0 *iam.GetGroupInput) (*request.Request, *iam.GetGroupOutput)
	GetGroupWithContextFunc                                func(param0 aws.Context, param1 *iam.GetGroupInput, param2 ...request.Option) (*iam.GetGroupOutput, error)
	GetInstanceProfileFunc                                 func(param0 *iam.GetInstanceProfileInput) (*iam.GetInstanceProfileOutput, error)
	GetInstanceProfileRequestFunc                          func(param0 *iam.GetInstanceProfileInput) (*request.Request, *iam.GetInstanceProfileOutput)
	GetInstanceProfileWithContextFunc                      func(param0 aws.Context, param1 *iam.GetInstanceProfileInput, param2 ...request.Option) (*iam.GetInstanceProfileOutput, error)
	GetLoginProfileFunc                                    func(param0 *iam.GetLoginProfileInput) (*iam.GetLoginProfileOutput, error)
	GetLoginProfileRequestFunc                             func(param0 *iam.GetLoginProfileInput) (*request.Request, *iam.GetLoginProfileOutput)
	GetLoginProfileWithContextFunc                         func(param0 aws.Context, param1 *iam.GetLoginProfileInput, param2 ...request.Option) (*iam.GetLoginProfileOutput, error)
	GetOpenIDConnectProviderFunc                           func(param0 *iam.GetOpenIDConnectProviderInput) (*iam.GetOpenIDConnectProviderOutput, error)
	GetOpenIDConnectProviderRequestFunc                    func(param0 *iam.GetOpenIDConnectProviderInput) (*request.Request, *iam.GetOpenIDConnectProviderOutput)
	GetOpenIDConnectProviderWithContextFunc                func(param0 aws.Context, param1 *iam.GetOpenIDConnectProviderInput, param2 ...request.Option) (*iam.GetOpenIDConnectProviderOutput, error)
	GetPolicyFunc                                          func(param0 *iam.GetPolicyInput) (*iam.GetPolicyOutput, error)
	GetPolicyRequestFunc                                   func(param0 *iam.GetPolicyInput) (*request.Request, *iam.GetPolicyOutput)
	GetPolicyVersionFunc                                   func(param0 *iam.GetPolicyVersionInput) (*iam.GetPolicyVersionOutput, error)
	GetPolicyVersionRequestFunc                            func(param0 *iam.GetPolicyVersionInput) (*request.Request, *iam.GetPolicyVersionOutput)
	GetPolicyVersionWithContextFunc                        func(param0 aws.Context, param1 *iam.GetPolicyVersionInput, param2 ...request.Option) (*iam.GetPolicyVersionOutput, error)
	GetPolicyWithContextFunc                               func(param0 aws.Context, param1 *iam.GetPolicyInput, param2 ...request.Option) (*iam.GetPolicyOutput, error)
	GetRoleFunc                                            func(param0 *iam.GetRoleInput) (*iam.GetRoleOutput, error)
	GetRolePolicyFunc                                      func(param0 *iam.GetRolePolicyInput) (*iam.GetRolePolicyOutput, error)
	GetRolePolicyRequestFunc                               func(param0 *iam.GetRolePolicyInput) (*request.Request, *iam.GetRolePolicyOutput)
	GetRolePolicyWithContextFunc                           func(param0 aws.Context, param1 *iam.GetRolePolicyInput, param2 ...request.Option) (*iam.GetRolePolicyOutput, error)
	GetRoleRequestFunc                                     func(param0 *iam.GetRoleInput) (*request.Request, *iam.GetRoleOutput)
	GetRoleWithContextFunc                                 func(param0 aws.Context, param1 *iam.GetRoleInput, param2 ...request.Option) (*iam.GetRoleOutput, error)
	GetSAMLProviderFunc                                    func(param0 *iam.GetSAMLProviderInput) (*iam.GetSAMLProviderOutput, error)
	GetSAMLProviderRequestFunc                             func(param0 *iam.GetSAMLProviderInput) (*request.Request, *iam.GetSAMLProviderOutput)
	GetSAMLProviderWithContextFunc                         func(param0 aws.Context, param1 *iam.GetSAMLProviderInput, param2 ...request.Option) (*iam.GetSAMLProviderOutput, error)
	GetSSHPublicKeyFunc                                    func(param0 *iam.GetSSHPublicKeyInput) (*iam.GetSSHPublicKeyOutput, error)
	GetSSHPublicKeyRequestFunc                             func(param0 *iam.GetSSHPublicKeyInput) (*request.Request, *iam.GetSSHPublicKeyOutput)
	GetSSHPublicKeyWithContextFunc                         func(param0 aws.Context, param1 *iam.GetSSHPublicKeyInput, param2 ...request.Option) (*iam.GetSSHPublicKeyOutput, error)
	GetServerCertificateFunc                               func(param0 *iam.GetServerCertificateInput) (*iam.GetServerCertificateOutput, error)
	GetServerCertificateRequestFunc                        func(param0 *iam.GetServerCertificateInput) (*request.Request, *iam.GetServerCertificateOutput)
	GetServerCertificateWithContextFunc                    func(param0 aws.Context, param1 *iam.GetServerCertificateInput, param2 ...request.Option) (*iam.GetServerCertificateOutput, error)
	GetServiceLinkedRoleDeletionStatusFunc                 func(param0 *iam.GetServiceLinkedRoleDeletionStatusInput) (*iam.GetServiceLinkedRoleDeletionStatusOutput, error)
	GetServiceLinkedRoleDeletionStatusRequestFunc          func(param0 *iam.GetServiceLinkedRoleDeletionStatusInput) (*request.Request, *iam.GetServiceLinkedRoleDeletionStatusOutput)
	GetServiceLinkedRoleDeletionStatusWithContextFunc      func(param0 aws.Context, param1 *iam.GetServiceLinkedRoleDeletionStatusInput, param2 ...request.Option) (*iam.GetServiceLinkedRoleDeletionStatusOutput, error)
	GetUserFunc                                            func(param0 *iam.GetUserInput) (*iam.GetUserOutput, error)
	GetUserPolicyFunc                                      func(param0 *iam.GetUserPolicyInput) (*iam.GetUserPolicyOutput, error)
	GetUserPolicyRequestFunc                               func(param0 *iam.GetUserPolicyInput) (*request.Request, *iam.GetUserPolicyOutput)
	GetUserPolicyWithContextFunc                           func(param0 aws.Context, param1 *iam.GetUserPolicyInput, param2 ...request.Option) (*iam.GetUserPolicyOutput, error)
	GetUserRequestFunc                                     func(param0 *iam.GetUserInput) (*request.Request, *iam.GetUserOutput)
	GetUserWithContextFunc                                 func(param0 aws.Context, param1 *iam.GetUserInput, param2 ...request.Option) (*iam.GetUserOutput, error)
	ListAccessKeysFunc                                     func(param0 *iam.ListAccessKeysInput) (*iam.ListAccessKeysOutput, error)
	ListAccessKeysRequestFunc                              func(param0 *iam.ListAccessKeysInput) (*request.Request, *iam.ListAccessKeysOutput)
	ListAccessKeysWithContextFunc                          func(param0 aws.Context, param1 *iam.ListAccessKeysInput, param2 ...request.Option) (*iam.ListAccessKeysOutput, error)
	ListAccountAliasesFunc                                 func(param0 *iam.ListAccountAliasesInput) (*iam.ListAccountAliasesOutput, error)
	ListAccountAliasesRequestFunc                          func(param0 *iam.ListAccountAliasesInput) (*request.Request, *iam.ListAccountAliasesOutput)
	ListAccountAliasesWithContextFunc                      func(param0 aws.Context, param1 *iam.ListAccountAliasesInput, param2 ...request.Option) (*iam.ListAccountAliasesOutput, error)
	ListAttachedGroupPoliciesFunc                          func(param0 *iam.ListAttachedGroupPoliciesInput) (*iam.ListAttachedGroupPoliciesOutput, error)
	ListAttachedGroupPoliciesRequestFunc                   func(param0 *iam.ListAttachedGroupPoliciesInput) (*request.Request, *iam.ListAttachedGroupPoliciesOutput)
	ListAttachedGroupPoliciesWithContextFunc               func(param0 aws.Context, param1 *iam.ListAttachedGroupPoliciesInput, param2 ...request.Option) (*iam.ListAttachedGroupPoliciesOutput, error)
	ListAttachedRolePoliciesFunc                           func(param0 *iam.ListAttachedRolePoliciesInput) (*iam.ListAttachedRolePoliciesOutput, error)
	ListAttachedRolePoliciesRequestFunc                    func(param0 *iam.ListAttachedRolePoliciesInput) (*request.Request, *iam.ListAttachedRolePoliciesOutput)
	ListAttachedRolePoliciesWithContextFunc                func(param0 aws.Context, param1 *iam.ListAttachedRolePoliciesInput, param2 ...request.Option) (*iam.ListAttachedRolePoliciesOutput, error)
	ListAttachedUserPoliciesFunc                           func(param0 *iam.ListAttachedUserPoliciesInput) (*iam.ListAttachedUserPoliciesOutput, error)
	ListAttachedUserPoliciesRequestFunc                    func(param0 *iam.ListAttachedUserPoliciesInput) (*request.Request, *iam.ListAttachedUserPoliciesOutput)
	ListAttachedUserPoliciesWithContextFunc                func(param0 aws.Context, param1 *iam.ListAttachedUserPoliciesInput, param2 ...request.Option) (*iam.ListAttachedUserPoliciesOutput, error)
	ListEntitiesForPolicyFunc                              func(param0 *iam.ListEntitiesForPolicyInput) (*iam.ListEntitiesForPolicyOutput, error)
	ListEntitiesForPolicyRequestFunc                       func(param0 *iam.ListEntitiesForPolicyInput) (*request.Request, *iam.ListEntitiesForPolicyOutput)
	ListEntitiesForPolicyWithContextFunc                   func(param0 aws.Context, param1 *iam.ListEntitiesForPolicyInput, param2 ...request.Option) (*iam.ListEntitiesForPolicyOutput, error)
	ListGroupPoliciesFunc                                  func(param0 *iam.ListGroupPoliciesInput) (*iam.ListGroupPoliciesOutput, error)
	ListGroupPoliciesRequestFunc                           func(param0 *iam.ListGroupPoliciesInput) (*request.Request, *iam.ListGroupPoliciesOutput)
	ListGroupPoliciesWithContextFunc                       func(param0 aws.Context, param1 *iam.ListGroupPoliciesInput, param2 ...request.Option) (*iam.ListGroupPoliciesOutput, error)
	ListGroupsFunc                                         func(param0 *iam.ListGroupsInput) (*iam.ListGroupsOutput, error)
	ListGroupsForUserFunc                                  func(param0 *iam.ListGroupsForUserInput) (*iam.ListGroupsForUserOutput, error)
	ListGroupsForUserRequestFunc                           func(param0 *iam.ListGroupsForUserInput) (*request.Request, *iam.ListGroupsForUserOutput)
	ListGroupsForUserWithContextFunc                       func(param0 aws.Context, param1 *iam.ListGroupsForUserInput, param2 ...request.Option) (*iam.ListGroupsForUserOutput, error)
	ListGroupsRequestFunc                                  func(param0 *iam.ListGroupsInput) (*request.Request, *iam.ListGroupsOutput)
	ListGroupsWithContextFunc                              func(param0 aws.Context, param1 *iam.ListGroupsInput, param2 ...request.Option) (*iam.ListGroupsOutput, error)
	ListInstanceProfilesFunc                               func(param0 *iam.ListInstanceProfilesInput) (*iam.ListInstanceProfilesOutput, error)
	ListInstanceProfilesForRoleFunc                        func(param0 *iam.ListInstanceProfilesForRoleInput) (*iam.ListInstanceProfilesForRoleOutput, error)
	ListInstanceProfilesForRoleRequestFunc                 func(param0 *iam.ListInstanceProfilesForRoleInput) (*request.Request, *iam.ListInstanceProfilesForRoleOutput)
	ListInstanceProfilesForRoleWithContextFunc             func(param0 aws.Context, param1 *iam.ListInstanceProfilesForRoleInput, param2 ...request.Option) (*iam.ListInstanceProfilesForRoleOutput, error)
	ListInstanceProfilesRequestFunc                        func(param0 *iam.ListInstanceProfilesInput) (*request.Request, *iam.ListInstanceProfilesOutput)
	ListInstanceProfilesWithContextFunc                    func(param0 aws.Context, param1 *iam.ListInstanceProfilesInput, param2 ...request.Option) (*iam.ListInstanceProfilesOutput, error)
	ListMFADevicesFunc                                     func(param0 *iam.ListMFADevicesInput) (*iam.ListMFADevicesOutput, error)
	ListMFADevicesRequestFunc                              func(param0 *iam.ListMFADevicesInput) (*request.Request, *iam.ListMFADevicesOutput)
	ListMFADevicesWithContextFunc                          func(param0 aws.Context, param1 *iam.ListMFADevicesInput, param2 ...request.Option) (*iam.ListMFADevicesOutput, error)
	ListOpenIDConnectProvidersFunc                         func(param0 *iam.ListOpenIDConnectProvidersInput) (*iam.ListOpenIDConnectProvidersOutput, error)
	ListOpenIDConnectProvidersRequestFunc                  func(param0 *iam.ListOpenIDConnectProvidersInput) (*request.Request, *iam.ListOpenIDConnectProvidersOutput)
	ListOpenIDConnectProvidersWithContextFunc              func(param0 aws.Context, param1 *iam.ListOpenIDConnectProvidersInput, param2 ...request.Option) (*iam.ListOpenIDConnectProvidersOutput, error)
	ListPoliciesFunc                                       func(param0 *iam.ListPoliciesInput) (*iam.ListPoliciesOutput, error)
	ListPoliciesRequestFunc                                func(param0 *iam.ListPoliciesInput) (*request.Request, *iam.ListPoliciesOutput)
	ListPoliciesWithContextFunc                            func(param0 aws.Context, param1 *iam.ListPoliciesInput, param2 ...request.Option) (*iam.ListPoliciesOutput, error)
	ListPolicyVersionsFunc                                 func(param0 *iam.ListPolicyVersionsInput) (*iam.ListPolicyVersionsOutput, error)
	ListPolicyVersionsRequestFunc                          func(param0 *iam.ListPolicyVersionsInput) (*request.Request, *iam.ListPolicyVersionsOutput)
	ListPolicyVersionsWithContextFunc                      func(param0 aws.Context, param1 *iam.ListPolicyVersionsInput, param2 ...request.Option) (*iam.ListPolicyVersionsOutput, error)
	ListRolePoliciesFunc                                   func(param0 *iam.ListRolePoliciesInput) (*iam.ListRolePoliciesOutput, error)
	ListRolePoliciesRequestFunc                            func(param0 *iam.ListRolePoliciesInput) (*request.Request, *iam.ListRolePoliciesOutput)
	ListRolePoliciesWithContextFunc                        func(param0 aws.Context, param1 *iam.ListRolePoliciesInput, param2 ...request.Option) (*iam.ListRolePoliciesOutput, error)
	ListRolesFunc                                          func(param0 *iam.ListRolesInput) (*iam.ListRolesOutput, error)
	ListRolesRequestFunc                                   func(param0 *iam.ListRolesInput) (*request.Request, *iam.ListRolesOutput)
	ListRolesWithContextFunc                               func(param0 aws.Context, param1 *iam.ListRolesInput, param2 ...request.Option) (*iam.ListRolesOutput, error)
	ListSAMLProvidersFunc                                  func(param0 *iam.ListSAMLProvidersInput) (*iam.ListSAMLProvidersOutput, error)
	ListSAMLProvidersRequestFunc                           func(param0 *iam.ListSAMLProvidersInput) (*request.Request, *iam.ListSAMLProvidersOutput)
	ListSAMLProvidersWithContextFunc                       func(param0 aws.Context, param1 *iam.ListSAMLProvidersInput, param2 ...request.Option) (*iam.ListSAMLProvidersOutput, error)
	ListSSHPublicKeysFunc                                  func(param0 *iam.ListSSHPublicKeysInput) (*iam.ListSSHPublicKeysOutput, error)
	ListSSHPublicKeysRequestFunc                           func(param0 *iam.ListSSHPublicKeysInput) (*request.Request, *iam.ListSSHPublicKeysOutput)
	ListSSHPublicKeysWithContextFunc                       func(param0 aws.Context, param1 *iam.ListSSHPublicKeysInput, param2 ...request.Option) (*iam.ListSSHPublicKeysOutput, error)
	ListServerCertificatesFunc                             func(param0 *iam.ListServerCertificatesInput) (*iam.ListServerCertificatesOutput, error)
	ListServerCertificatesRequestFunc                      func(param0 *iam.ListServerCertificatesInput) (*request.Request, *iam.ListServerCertificatesOutput)
	ListServerCertificatesWithContextFunc                  func(param0 aws.Context, param1 *iam.ListServerCertificatesInput, param2 ...request.Option) (*iam.ListServerCertificatesOutput, error)
	ListServiceSpecificCredentialsFunc                     func(param0 *iam.ListServiceSpecificCredentialsInput) (*iam.ListServiceSpecificCredentialsOutput, error)
	ListServiceSpecificCredentialsRequestFunc              func(param0 *iam.ListServiceSpecificCredentialsInput) (*request.Request, *iam.ListServiceSpecificCredentialsOutput)
	ListServiceSpecificCredentialsWithContextFunc          func(param0 aws.Context, param1 *iam.ListServiceSpecificCredentialsInput, param2 ...request.Option) (*iam.ListServiceSpecificCredentialsOutput, error)
	ListSigningCertificatesFunc                            func(param0 *iam.ListSigningCertificatesInput) (*iam.ListSigningCertificatesOutput, error)
	ListSigningCertificatesRequestFunc                     func(param0 *iam.ListSigningCertificatesInput) (*request.Request, *iam.ListSigningCertificatesOutput)
	ListSigningCertificatesWithContextFunc                 func(param0 aws.Context, param1 *iam.ListSigningCertificatesInput, param2 ...request.Option) (*iam.ListSigningCertificatesOutput, error)
	ListUserPoliciesFunc                                   func(param0 *iam.ListUserPoliciesInput) (*iam.ListUserPoliciesOutput, error)
	ListUserPoliciesRequestFunc                            func(param0 *iam.ListUserPoliciesInput) (*request.Request, *iam.ListUserPoliciesOutput)
	ListUserPoliciesWithContextFunc                        func(param0 aws.Context, param1 *iam.ListUserPoliciesInput, param2 ...request.Option) (*iam.ListUserPoliciesOutput, error)
	ListUsersFunc                                          func(param0 *iam.ListUsersInput) (*iam.ListUsersOutput, error)
	ListUsersRequestFunc                                   func(param0 *iam.ListUsersInput) (*request.Request, *iam.ListUsersOutput)
	ListUsersWithContextFunc                               func(param0 aws.Context, param1 *iam.ListUsersInput, param2 ...request.Option) (*iam.ListUsersOutput, error)
	ListVirtualMFADevicesFunc                              func(param0 *iam.ListVirtualMFADevicesInput) (*iam.ListVirtualMFADevicesOutput, error)
	ListVirtualMFADevicesRequestFunc                       func(param0 *iam.ListVirtualMFADevicesInput) (*request.Request, *iam.ListVirtualMFADevicesOutput)
	ListVirtualMFADevicesWithContextFunc                   func(param0 aws.Context, param1 *iam.ListVirtualMFADevicesInput, param2 ...request.Option) (*iam.ListVirtualMFADevicesOutput, error)
	PutGroupPolicyFunc                                     func(param0 *iam.PutGroupPolicyInput) (*iam.PutGroupPolicyOutput, error)
	PutGroupPolicyRequestFunc                              func(param0 *iam.PutGroupPolicyInput) (*request.Request, *iam.PutGroupPolicyOutput)
	PutGroupPolicyWithContextFunc                          func(param0 aws.Context, param1 *iam.PutGroupPolicyInput, param2 ...request.Option) (*iam.PutGroupPolicyOutput, error)
	PutRolePolicyFunc                                      func(param0 *iam.PutRolePolicyInput) (*iam.PutRolePolicyOutput, error)
	PutRolePolicyRequestFunc                               func(param0 *iam.PutRolePolicyInput) (*request.Request, *iam.PutRolePolicyOutput)
	PutRolePolicyWithContextFunc                           func(param0 aws.Context, param1 *iam.PutRolePolicyInput, param2 ...request.Option) (*iam.PutRolePolicyOutput, error)
	PutUserPolicyFunc                                      func(param0 *iam.PutUserPolicyInput) (*iam.PutUserPolicyOutput, error)
	PutUserPolicyRequestFunc                               func(param0 *iam.PutUserPolicyInput) (*request.Request, *iam.PutUserPolicyOutput)
	PutUserPolicyWithContextFunc                           func(param0 aws.Context, param1 *iam.PutUserPolicyInput, param2 ...request.Option) (*iam.PutUserPolicyOutput, error)
	RemoveClientIDFromOpenIDConnectProviderFunc            func(param0 *iam.RemoveClientIDFromOpenIDConnectProviderInput) (*iam.RemoveClientIDFromOpenIDConnectProviderOutput, error)
	RemoveClientIDFromOpenIDConnectProviderRequestFunc     func(param0 *iam.RemoveClientIDFromOpenIDConnectProviderInput) (*request.Request, *iam.RemoveClientIDFromOpenIDConnectProviderOutput)
	RemoveClientIDFromOpenIDConnectProviderWithContextFunc func(param0 aws.Context, param1 *iam.RemoveClientIDFromOpenIDConnectProviderInput, param2 ...request.Option) (*iam.RemoveClientIDFromOpenIDConnectProviderOutput, error)
	RemoveRoleFromInstanceProfileFunc                      func(param0 *iam.RemoveRoleFromInstanceProfileInput) (*iam.RemoveRoleFromInstanceProfileOutput, error)
	RemoveRoleFromInstanceProfileRequestFunc               func(param0 *iam.RemoveRoleFromInstanceProfileInput) (*request.Request, *iam.RemoveRoleFromInstanceProfileOutput)
	RemoveRoleFromInstanceProfileWithContextFunc           func(param0 aws.Context, param1 *iam.RemoveRoleFromInstanceProfileInput, param2 ...request.Option) (*iam.RemoveRoleFromInstanceProfileOutput, error)
	RemoveUserFromGroupFunc                                func(param0 *iam.RemoveUserFromGroupInput) (*iam.RemoveUserFromGroupOutput, error)
	RemoveUserFromGroupRequestFunc                         func(param0 *iam.RemoveUserFromGroupInput) (*request.Request, *iam.RemoveUserFromGroupOutput)
	RemoveUserFromGroupWithContextFunc                     func(param0 aws.Context, param1 *iam.RemoveUserFromGroupInput, param2 ...request.Option) (*iam.RemoveUserFromGroupOutput, error)
	ResetServiceSpecificCredentialFunc                     func(param0 *iam.ResetServiceSpecificCredentialInput) (*iam.ResetServiceSpecificCredentialOutput, error)
	ResetServiceSpecificCredentialRequestFunc              func(param0 *iam.ResetServiceSpecificCredentialInput) (*request.Request, *iam.ResetServiceSpecificCredentialOutput)
	ResetServiceSpecificCredentialWithContextFunc          func(param0 aws.Context, param1 *iam.ResetServiceSpecificCredentialInput, param2 ...request.Option) (*iam.ResetServiceSpecificCredentialOutput, error)
	ResyncMFADeviceFunc                                    func(param0 *iam.ResyncMFADeviceInput) (*iam.ResyncMFADeviceOutput, error)
	ResyncMFADeviceRequestFunc                             func(param0 *iam.ResyncMFADeviceInput) (*request.Request, *iam.ResyncMFADeviceOutput)
	ResyncMFADeviceWithContextFunc                         func(param0 aws.Context, param1 *iam.ResyncMFADeviceInput, param2 ...request.Option) (*iam.ResyncMFADeviceOutput, error)
	SetDefaultPolicyVersionFunc                            func(param0 *iam.SetDefaultPolicyVersionInput) (*iam.SetDefaultPolicyVersionOutput, error)
	SetDefaultPolicyVersionRequestFunc                     func(param0 *iam.SetDefaultPolicyVersionInput) (*request.Request, *iam.SetDefaultPolicyVersionOutput)
	SetDefaultPolicyVersionWithContextFunc                 func(param0 aws.Context, param1 *iam.SetDefaultPolicyVersionInput, param2 ...request.Option) (*iam.SetDefaultPolicyVersionOutput, error)
	SimulateCustomPolicyFunc                               func(param0 *iam.SimulateCustomPolicyInput) (*iam.SimulatePolicyResponse, error)
	SimulateCustomPolicyRequestFunc                        func(param0 *iam.SimulateCustomPolicyInput) (*request.Request, *iam.SimulatePolicyResponse)
	SimulateCustomPolicyWithContextFunc                    func(param0 aws.Context, param1 *iam.SimulateCustomPolicyInput, param2 ...request.Option) (*iam.SimulatePolicyResponse, error)
	SimulatePrincipalPolicyFunc                            func(param0 *iam.SimulatePrincipalPolicyInput) (*iam.SimulatePolicyResponse, error)
	SimulatePrincipalPolicyRequestFunc                     func(param0 *iam.SimulatePrincipalPolicyInput) (*request.Request, *iam.SimulatePolicyResponse)
	SimulatePrincipalPolicyWithContextFunc                 func(param0 aws.Context, param1 *iam.SimulatePrincipalPolicyInput, param2 ...request.Option) (*iam.SimulatePolicyResponse, error)
	UpdateAccessKeyFunc                                    func(param0 *iam.UpdateAccessKeyInput) (*iam.UpdateAccessKeyOutput, error)
	UpdateAccessKeyRequestFunc                             func(param0 *iam.UpdateAccessKeyInput) (*request.Request, *iam.UpdateAccessKeyOutput)
	UpdateAccessKeyWithContextFunc                         func(param0 aws.Context, param1 *iam.UpdateAccessKeyInput, param2 ...request.Option) (*iam.UpdateAccessKeyOutput, error)
	UpdateAccountPasswordPolicyFunc                        func(param0 *iam.UpdateAccountPasswordPolicyInput) (*iam.UpdateAccountPasswordPolicyOutput, error)
	UpdateAccountPasswordPolicyRequestFunc                 func(param0 *iam.UpdateAccountPasswordPolicyInput) (*request.Request, *iam.UpdateAccountPasswordPolicyOutput)
	UpdateAccountPasswordPolicyWithContextFunc             func(param0 aws.Context, param1 *iam.UpdateAccountPasswordPolicyInput, param2 ...request.Option) (*iam.UpdateAccountPasswordPolicyOutput, error)
	UpdateAssumeRolePolicyFunc                             func(param0 *iam.UpdateAssumeRolePolicyInput) (*iam.UpdateAssumeRolePolicyOutput, error)
	UpdateAssumeRolePolicyRequestFunc                      func(param0 *iam.UpdateAssumeRolePolicyInput) (*request.Request, *iam.UpdateAssumeRolePolicyOutput)
	UpdateAssumeRolePolicyWithContextFunc                  func(param0 aws.Context, param1 *iam.UpdateAssumeRolePolicyInput, param2 ...request.Option) (*iam.UpdateAssumeRolePolicyOutput, error)
	UpdateGroupFunc                                        func(param0 *iam.UpdateGroupInput) (*iam.UpdateGroupOutput, error)
	UpdateGroupRequestFunc                                 func(param0 *iam.UpdateGroupInput) (*request.Request, *iam.UpdateGroupOutput)
	UpdateGroupWithContextFunc                             func(param0 aws.Context, param1 *iam.UpdateGroupInput, param2 ...request.Option) (*iam.UpdateGroupOutput, error)
	UpdateLoginProfileFunc                                 func(param0 *iam.UpdateLoginProfileInput) (*iam.UpdateLoginProfileOutput, error)
	UpdateLoginProfileRequestFunc                          func(param0 *iam.UpdateLoginProfileInput) (*request.Request, *iam.UpdateLoginProfileOutput)
	UpdateLoginProfileWithContextFunc                      func(param0 aws.Context, param1 *iam.UpdateLoginProfileInput, param2 ...request.Option) (*iam.UpdateLoginProfileOutput, error)
	UpdateOpenIDConnectProviderThumbprintFunc              func(param0 *iam.UpdateOpenIDConnectProviderThumbprintInput) (*iam.UpdateOpenIDConnectProviderThumbprintOutput, error)
	UpdateOpenIDConnectProviderThumbprintRequestFunc       func(param0 *iam.UpdateOpenIDConnectProviderThumbprintInput) (*request.Request, *iam.UpdateOpenIDConnectProviderThumbprintOutput)
	UpdateOpenIDConnectProviderThumbprintWithContextFunc   func(param0 aws.Context, param1 *iam.UpdateOpenIDConnectProviderThumbprintInput, param2 ...request.Option) (*iam.UpdateOpenIDConnectProviderThumbprintOutput, error)
	UpdateRoleDescriptionFunc                              func(param0 *iam.UpdateRoleDescriptionInput) (*iam.UpdateRoleDescriptionOutput, error)
	UpdateRoleDescriptionRequestFunc                       func(param0 *iam.UpdateRoleDescriptionInput) (*request.Request, *iam.UpdateRoleDescriptionOutput)
	UpdateRoleDescriptionWithContextFunc                   func(param0 aws.Context, param1 *iam.UpdateRoleDescriptionInput, param2 ...request.Option) (*iam.UpdateRoleDescriptionOutput, error)
	UpdateSAMLProviderFunc                                 func(param0 *iam.UpdateSAMLProviderInput) (*iam.UpdateSAMLProviderOutput, error)
	UpdateSAMLProviderRequestFunc                          func(param0 *iam.UpdateSAMLProviderInput) (*request.Request, *iam.UpdateSAMLProviderOutput)
	UpdateSAMLProviderWithContextFunc                      func(param0 aws.Context, param1 *iam.UpdateSAMLProviderInput, param2 ...request.Option) (*iam.UpdateSAMLProviderOutput, error)
	UpdateSSHPublicKeyFunc                                 func(param0 *iam.UpdateSSHPublicKeyInput) (*iam.UpdateSSHPublicKeyOutput, error)
	UpdateSSHPublicKeyRequestFunc                          func(param0 *iam.UpdateSSHPublicKeyInput) (*request.Request, *iam.UpdateSSHPublicKeyOutput)
	UpdateSSHPublicKeyWithContextFunc                      func(param0 aws.Context, param1 *iam.UpdateSSHPublicKeyInput, param2 ...request.Option) (*iam.UpdateSSHPublicKeyOutput, error)
	UpdateServerCertificateFunc                            func(param0 *iam.UpdateServerCertificateInput) (*iam.UpdateServerCertificateOutput, error)
	UpdateServerCertificateRequestFunc                     func(param0 *iam.UpdateServerCertificateInput) (*request.Request, *iam.UpdateServerCertificateOutput)
	UpdateServerCertificateWithContextFunc                 func(param0 aws.Context, param1 *iam.UpdateServerCertificateInput, param2 ...request.Option) (*iam.UpdateServerCertificateOutput, error)
	UpdateServiceSpecificCredentialFunc                    func(param0 *iam.UpdateServiceSpecificCredentialInput) (*iam.UpdateServiceSpecificCredentialOutput, error)
	UpdateServiceSpecificCredentialRequestFunc             func(param0 *iam.UpdateServiceSpecificCredentialInput) (*request.Request, *iam.UpdateServiceSpecificCredentialOutput)
	UpdateServiceSpecificCredentialWithContextFunc         func(param0 aws.Context, param1 *iam.UpdateServiceSpecificCredentialInput, param2 ...request.Option) (*iam.UpdateServiceSpecificCredentialOutput, error)
	UpdateSigningCertificateFunc                           func(param0 *iam.UpdateSigningCertificateInput) (*iam.UpdateSigningCertificateOutput, error)
	UpdateSigningCertificateRequestFunc                    func(param0 *iam.UpdateSigningCertificateInput) (*request.Request, *iam.UpdateSigningCertificateOutput)
	UpdateSigningCertificateWithContextFunc                func(param0 aws.Context, param1 *iam.UpdateSigningCertificateInput, param2 ...request.Option) (*iam.UpdateSigningCertificateOutput, error)
	UpdateUserFunc                                         func(param0 *iam.UpdateUserInput) (*iam.UpdateUserOutput, error)
	UpdateUserRequestFunc                                  func(param0 *iam.UpdateUserInput) (*request.Request, *iam.UpdateUserOutput)
	UpdateUserWithContextFunc                              func(param0 aws.Context, param1 *iam.UpdateUserInput, param2 ...request.Option) (*iam.UpdateUserOutput, error)
	UploadSSHPublicKeyFunc                                 func(param0 *iam.UploadSSHPublicKeyInput) (*iam.UploadSSHPublicKeyOutput, error)
	UploadSSHPublicKeyRequestFunc                          func(param0 *iam.UploadSSHPublicKeyInput) (*request.Request, *iam.UploadSSHPublicKeyOutput)
	UploadSSHPublicKeyWithContextFunc                      func(param0 aws.Context, param1 *iam.UploadSSHPublicKeyInput, param2 ...request.Option) (*iam.UploadSSHPublicKeyOutput, error)
	UploadServerCertificateFunc                            func(param0 *iam.UploadServerCertificateInput) (*iam.UploadServerCertificateOutput, error)
	UploadServerCertificateRequestFunc                     func(param0 *iam.UploadServerCertificateInput) (*request.Request, *iam.UploadServerCertificateOutput)
	UploadServerCertificateWithContextFunc                 func(param0 aws.Context, param1 *iam.UploadServerCertificateInput, param2 ...request.Option) (*iam.UploadServerCertificateOutput, error)
	UploadSigningCertificateFunc                           func(param0 *iam.UploadSigningCertificateInput) (*iam.UploadSigningCertificateOutput, error)
	UploadSigningCertificateRequestFunc                    func(param0 *iam.UploadSigningCertificateInput) (*request.Request, *iam.UploadSigningCertificateOutput)
	UploadSigningCertificateWithContextFunc                func(param0 aws.Context, param1 *iam.UploadSigningCertificateInput, param2 ...request.Option) (*iam.UploadSigningCertificateOutput, error)
	WaitUntilInstanceProfileExistsFunc                     func(param0 *iam.GetInstanceProfileInput) error
	WaitUntilInstanceProfileExistsWithContextFunc          func(param0 aws.Context, param1 *iam.GetInstanceProfileInput, param2 ...request.WaiterOption) error
	WaitUntilUserExistsFunc                                func(param0 *iam.GetUserInput) error
	WaitUntilUserExistsWithContextFunc                     func(param0 aws.Context, param1 *iam.GetUserInput, param2 ...request.WaiterOption) error
}

func (m *iamMock) AddClientIDToOpenIDConnectProvider(param0 *iam.AddClientIDToOpenIDConnectProviderInput) (*iam.AddClientIDToOpenIDConnectProviderOutput, error) {
	m.addCall("AddClientIDToOpenIDConnectProvider")
	m.verifyInput("AddClientIDToOpenIDConnectProvider", param0)
	return m.AddClientIDToOpenIDConnectProviderFunc(param0)
}

func (m *iamMock) AddClientIDToOpenIDConnectProviderRequest(param0 *iam.AddClientIDToOpenIDConnectProviderInput) (*request.Request, *iam.AddClientIDToOpenIDConnectProviderOutput) {
	m.addCall("AddClientIDToOpenIDConnectProviderRequest")
	m.verifyInput("AddClientIDToOpenIDConnectProviderRequest", param0)
	return m.AddClientIDToOpenIDConnectProviderRequestFunc(param0)
}

func (m *iamMock) AddClientIDToOpenIDConnectProviderWithContext(param0 aws.Context, param1 *iam.AddClientIDToOpenIDConnectProviderInput, param2 ...request.Option) (*iam.AddClientIDToOpenIDConnectProviderOutput, error) {
	m.addCall("AddClientIDToOpenIDConnectProviderWithContext")
	m.verifyInput("AddClientIDToOpenIDConnectProviderWithContext", param0)
	return m.AddClientIDToOpenIDConnectProviderWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) AddRoleToInstanceProfile(param0 *iam.AddRoleToInstanceProfileInput) (*iam.AddRoleToInstanceProfileOutput, error) {
	m.addCall("AddRoleToInstanceProfile")
	m.verifyInput("AddRoleToInstanceProfile", param0)
	return m.AddRoleToInstanceProfileFunc(param0)
}

func (m *iamMock) AddRoleToInstanceProfileRequest(param0 *iam.AddRoleToInstanceProfileInput) (*request.Request, *iam.AddRoleToInstanceProfileOutput) {
	m.addCall("AddRoleToInstanceProfileRequest")
	m.verifyInput("AddRoleToInstanceProfileRequest", param0)
	return m.AddRoleToInstanceProfileRequestFunc(param0)
}

func (m *iamMock) AddRoleToInstanceProfileWithContext(param0 aws.Context, param1 *iam.AddRoleToInstanceProfileInput, param2 ...request.Option) (*iam.AddRoleToInstanceProfileOutput, error) {
	m.addCall("AddRoleToInstanceProfileWithContext")
	m.verifyInput("AddRoleToInstanceProfileWithContext", param0)
	return m.AddRoleToInstanceProfileWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) AddUserToGroup(param0 *iam.AddUserToGroupInput) (*iam.AddUserToGroupOutput, error) {
	m.addCall("AddUserToGroup")
	m.verifyInput("AddUserToGroup", param0)
	return m.AddUserToGroupFunc(param0)
}

func (m *iamMock) AddUserToGroupRequest(param0 *iam.AddUserToGroupInput) (*request.Request, *iam.AddUserToGroupOutput) {
	m.addCall("AddUserToGroupRequest")
	m.verifyInput("AddUserToGroupRequest", param0)
	return m.AddUserToGroupRequestFunc(param0)
}

func (m *iamMock) AddUserToGroupWithContext(param0 aws.Context, param1 *iam.AddUserToGroupInput, param2 ...request.Option) (*iam.AddUserToGroupOutput, error) {
	m.addCall("AddUserToGroupWithContext")
	m.verifyInput("AddUserToGroupWithContext", param0)
	return m.AddUserToGroupWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) AttachGroupPolicy(param0 *iam.AttachGroupPolicyInput) (*iam.AttachGroupPolicyOutput, error) {
	m.addCall("AttachGroupPolicy")
	m.verifyInput("AttachGroupPolicy", param0)
	return m.AttachGroupPolicyFunc(param0)
}

func (m *iamMock) AttachGroupPolicyRequest(param0 *iam.AttachGroupPolicyInput) (*request.Request, *iam.AttachGroupPolicyOutput) {
	m.addCall("AttachGroupPolicyRequest")
	m.verifyInput("AttachGroupPolicyRequest", param0)
	return m.AttachGroupPolicyRequestFunc(param0)
}

func (m *iamMock) AttachGroupPolicyWithContext(param0 aws.Context, param1 *iam.AttachGroupPolicyInput, param2 ...request.Option) (*iam.AttachGroupPolicyOutput, error) {
	m.addCall("AttachGroupPolicyWithContext")
	m.verifyInput("AttachGroupPolicyWithContext", param0)
	return m.AttachGroupPolicyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) AttachRolePolicy(param0 *iam.AttachRolePolicyInput) (*iam.AttachRolePolicyOutput, error) {
	m.addCall("AttachRolePolicy")
	m.verifyInput("AttachRolePolicy", param0)
	return m.AttachRolePolicyFunc(param0)
}

func (m *iamMock) AttachRolePolicyRequest(param0 *iam.AttachRolePolicyInput) (*request.Request, *iam.AttachRolePolicyOutput) {
	m.addCall("AttachRolePolicyRequest")
	m.verifyInput("AttachRolePolicyRequest", param0)
	return m.AttachRolePolicyRequestFunc(param0)
}

func (m *iamMock) AttachRolePolicyWithContext(param0 aws.Context, param1 *iam.AttachRolePolicyInput, param2 ...request.Option) (*iam.AttachRolePolicyOutput, error) {
	m.addCall("AttachRolePolicyWithContext")
	m.verifyInput("AttachRolePolicyWithContext", param0)
	return m.AttachRolePolicyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) AttachUserPolicy(param0 *iam.AttachUserPolicyInput) (*iam.AttachUserPolicyOutput, error) {
	m.addCall("AttachUserPolicy")
	m.verifyInput("AttachUserPolicy", param0)
	return m.AttachUserPolicyFunc(param0)
}

func (m *iamMock) AttachUserPolicyRequest(param0 *iam.AttachUserPolicyInput) (*request.Request, *iam.AttachUserPolicyOutput) {
	m.addCall("AttachUserPolicyRequest")
	m.verifyInput("AttachUserPolicyRequest", param0)
	return m.AttachUserPolicyRequestFunc(param0)
}

func (m *iamMock) AttachUserPolicyWithContext(param0 aws.Context, param1 *iam.AttachUserPolicyInput, param2 ...request.Option) (*iam.AttachUserPolicyOutput, error) {
	m.addCall("AttachUserPolicyWithContext")
	m.verifyInput("AttachUserPolicyWithContext", param0)
	return m.AttachUserPolicyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) ChangePassword(param0 *iam.ChangePasswordInput) (*iam.ChangePasswordOutput, error) {
	m.addCall("ChangePassword")
	m.verifyInput("ChangePassword", param0)
	return m.ChangePasswordFunc(param0)
}

func (m *iamMock) ChangePasswordRequest(param0 *iam.ChangePasswordInput) (*request.Request, *iam.ChangePasswordOutput) {
	m.addCall("ChangePasswordRequest")
	m.verifyInput("ChangePasswordRequest", param0)
	return m.ChangePasswordRequestFunc(param0)
}

func (m *iamMock) ChangePasswordWithContext(param0 aws.Context, param1 *iam.ChangePasswordInput, param2 ...request.Option) (*iam.ChangePasswordOutput, error) {
	m.addCall("ChangePasswordWithContext")
	m.verifyInput("ChangePasswordWithContext", param0)
	return m.ChangePasswordWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) CreateAccessKey(param0 *iam.CreateAccessKeyInput) (*iam.CreateAccessKeyOutput, error) {
	m.addCall("CreateAccessKey")
	m.verifyInput("CreateAccessKey", param0)
	return m.CreateAccessKeyFunc(param0)
}

func (m *iamMock) CreateAccessKeyRequest(param0 *iam.CreateAccessKeyInput) (*request.Request, *iam.CreateAccessKeyOutput) {
	m.addCall("CreateAccessKeyRequest")
	m.verifyInput("CreateAccessKeyRequest", param0)
	return m.CreateAccessKeyRequestFunc(param0)
}

func (m *iamMock) CreateAccessKeyWithContext(param0 aws.Context, param1 *iam.CreateAccessKeyInput, param2 ...request.Option) (*iam.CreateAccessKeyOutput, error) {
	m.addCall("CreateAccessKeyWithContext")
	m.verifyInput("CreateAccessKeyWithContext", param0)
	return m.CreateAccessKeyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) CreateAccountAlias(param0 *iam.CreateAccountAliasInput) (*iam.CreateAccountAliasOutput, error) {
	m.addCall("CreateAccountAlias")
	m.verifyInput("CreateAccountAlias", param0)
	return m.CreateAccountAliasFunc(param0)
}

func (m *iamMock) CreateAccountAliasRequest(param0 *iam.CreateAccountAliasInput) (*request.Request, *iam.CreateAccountAliasOutput) {
	m.addCall("CreateAccountAliasRequest")
	m.verifyInput("CreateAccountAliasRequest", param0)
	return m.CreateAccountAliasRequestFunc(param0)
}

func (m *iamMock) CreateAccountAliasWithContext(param0 aws.Context, param1 *iam.CreateAccountAliasInput, param2 ...request.Option) (*iam.CreateAccountAliasOutput, error) {
	m.addCall("CreateAccountAliasWithContext")
	m.verifyInput("CreateAccountAliasWithContext", param0)
	return m.CreateAccountAliasWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) CreateGroup(param0 *iam.CreateGroupInput) (*iam.CreateGroupOutput, error) {
	m.addCall("CreateGroup")
	m.verifyInput("CreateGroup", param0)
	return m.CreateGroupFunc(param0)
}

func (m *iamMock) CreateGroupRequest(param0 *iam.CreateGroupInput) (*request.Request, *iam.CreateGroupOutput) {
	m.addCall("CreateGroupRequest")
	m.verifyInput("CreateGroupRequest", param0)
	return m.CreateGroupRequestFunc(param0)
}

func (m *iamMock) CreateGroupWithContext(param0 aws.Context, param1 *iam.CreateGroupInput, param2 ...request.Option) (*iam.CreateGroupOutput, error) {
	m.addCall("CreateGroupWithContext")
	m.verifyInput("CreateGroupWithContext", param0)
	return m.CreateGroupWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) CreateInstanceProfile(param0 *iam.CreateInstanceProfileInput) (*iam.CreateInstanceProfileOutput, error) {
	m.addCall("CreateInstanceProfile")
	m.verifyInput("CreateInstanceProfile", param0)
	return m.CreateInstanceProfileFunc(param0)
}

func (m *iamMock) CreateInstanceProfileRequest(param0 *iam.CreateInstanceProfileInput) (*request.Request, *iam.CreateInstanceProfileOutput) {
	m.addCall("CreateInstanceProfileRequest")
	m.verifyInput("CreateInstanceProfileRequest", param0)
	return m.CreateInstanceProfileRequestFunc(param0)
}

func (m *iamMock) CreateInstanceProfileWithContext(param0 aws.Context, param1 *iam.CreateInstanceProfileInput, param2 ...request.Option) (*iam.CreateInstanceProfileOutput, error) {
	m.addCall("CreateInstanceProfileWithContext")
	m.verifyInput("CreateInstanceProfileWithContext", param0)
	return m.CreateInstanceProfileWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) CreateLoginProfile(param0 *iam.CreateLoginProfileInput) (*iam.CreateLoginProfileOutput, error) {
	m.addCall("CreateLoginProfile")
	m.verifyInput("CreateLoginProfile", param0)
	return m.CreateLoginProfileFunc(param0)
}

func (m *iamMock) CreateLoginProfileRequest(param0 *iam.CreateLoginProfileInput) (*request.Request, *iam.CreateLoginProfileOutput) {
	m.addCall("CreateLoginProfileRequest")
	m.verifyInput("CreateLoginProfileRequest", param0)
	return m.CreateLoginProfileRequestFunc(param0)
}

func (m *iamMock) CreateLoginProfileWithContext(param0 aws.Context, param1 *iam.CreateLoginProfileInput, param2 ...request.Option) (*iam.CreateLoginProfileOutput, error) {
	m.addCall("CreateLoginProfileWithContext")
	m.verifyInput("CreateLoginProfileWithContext", param0)
	return m.CreateLoginProfileWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) CreateOpenIDConnectProvider(param0 *iam.CreateOpenIDConnectProviderInput) (*iam.CreateOpenIDConnectProviderOutput, error) {
	m.addCall("CreateOpenIDConnectProvider")
	m.verifyInput("CreateOpenIDConnectProvider", param0)
	return m.CreateOpenIDConnectProviderFunc(param0)
}

func (m *iamMock) CreateOpenIDConnectProviderRequest(param0 *iam.CreateOpenIDConnectProviderInput) (*request.Request, *iam.CreateOpenIDConnectProviderOutput) {
	m.addCall("CreateOpenIDConnectProviderRequest")
	m.verifyInput("CreateOpenIDConnectProviderRequest", param0)
	return m.CreateOpenIDConnectProviderRequestFunc(param0)
}

func (m *iamMock) CreateOpenIDConnectProviderWithContext(param0 aws.Context, param1 *iam.CreateOpenIDConnectProviderInput, param2 ...request.Option) (*iam.CreateOpenIDConnectProviderOutput, error) {
	m.addCall("CreateOpenIDConnectProviderWithContext")
	m.verifyInput("CreateOpenIDConnectProviderWithContext", param0)
	return m.CreateOpenIDConnectProviderWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) CreatePolicy(param0 *iam.CreatePolicyInput) (*iam.CreatePolicyOutput, error) {
	m.addCall("CreatePolicy")
	m.verifyInput("CreatePolicy", param0)
	return m.CreatePolicyFunc(param0)
}

func (m *iamMock) CreatePolicyRequest(param0 *iam.CreatePolicyInput) (*request.Request, *iam.CreatePolicyOutput) {
	m.addCall("CreatePolicyRequest")
	m.verifyInput("CreatePolicyRequest", param0)
	return m.CreatePolicyRequestFunc(param0)
}

func (m *iamMock) CreatePolicyVersion(param0 *iam.CreatePolicyVersionInput) (*iam.CreatePolicyVersionOutput, error) {
	m.addCall("CreatePolicyVersion")
	m.verifyInput("CreatePolicyVersion", param0)
	return m.CreatePolicyVersionFunc(param0)
}

func (m *iamMock) CreatePolicyVersionRequest(param0 *iam.CreatePolicyVersionInput) (*request.Request, *iam.CreatePolicyVersionOutput) {
	m.addCall("CreatePolicyVersionRequest")
	m.verifyInput("CreatePolicyVersionRequest", param0)
	return m.CreatePolicyVersionRequestFunc(param0)
}

func (m *iamMock) CreatePolicyVersionWithContext(param0 aws.Context, param1 *iam.CreatePolicyVersionInput, param2 ...request.Option) (*iam.CreatePolicyVersionOutput, error) {
	m.addCall("CreatePolicyVersionWithContext")
	m.verifyInput("CreatePolicyVersionWithContext", param0)
	return m.CreatePolicyVersionWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) CreatePolicyWithContext(param0 aws.Context, param1 *iam.CreatePolicyInput, param2 ...request.Option) (*iam.CreatePolicyOutput, error) {
	m.addCall("CreatePolicyWithContext")
	m.verifyInput("CreatePolicyWithContext", param0)
	return m.CreatePolicyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) CreateRole(param0 *iam.CreateRoleInput) (*iam.CreateRoleOutput, error) {
	m.addCall("CreateRole")
	m.verifyInput("CreateRole", param0)
	return m.CreateRoleFunc(param0)
}

func (m *iamMock) CreateRoleRequest(param0 *iam.CreateRoleInput) (*request.Request, *iam.CreateRoleOutput) {
	m.addCall("CreateRoleRequest")
	m.verifyInput("CreateRoleRequest", param0)
	return m.CreateRoleRequestFunc(param0)
}

func (m *iamMock) CreateRoleWithContext(param0 aws.Context, param1 *iam.CreateRoleInput, param2 ...request.Option) (*iam.CreateRoleOutput, error) {
	m.addCall("CreateRoleWithContext")
	m.verifyInput("CreateRoleWithContext", param0)
	return m.CreateRoleWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) CreateSAMLProvider(param0 *iam.CreateSAMLProviderInput) (*iam.CreateSAMLProviderOutput, error) {
	m.addCall("CreateSAMLProvider")
	m.verifyInput("CreateSAMLProvider", param0)
	return m.CreateSAMLProviderFunc(param0)
}

func (m *iamMock) CreateSAMLProviderRequest(param0 *iam.CreateSAMLProviderInput) (*request.Request, *iam.CreateSAMLProviderOutput) {
	m.addCall("CreateSAMLProviderRequest")
	m.verifyInput("CreateSAMLProviderRequest", param0)
	return m.CreateSAMLProviderRequestFunc(param0)
}

func (m *iamMock) CreateSAMLProviderWithContext(param0 aws.Context, param1 *iam.CreateSAMLProviderInput, param2 ...request.Option) (*iam.CreateSAMLProviderOutput, error) {
	m.addCall("CreateSAMLProviderWithContext")
	m.verifyInput("CreateSAMLProviderWithContext", param0)
	return m.CreateSAMLProviderWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) CreateServiceLinkedRole(param0 *iam.CreateServiceLinkedRoleInput) (*iam.CreateServiceLinkedRoleOutput, error) {
	m.addCall("CreateServiceLinkedRole")
	m.verifyInput("CreateServiceLinkedRole", param0)
	return m.CreateServiceLinkedRoleFunc(param0)
}

func (m *iamMock) CreateServiceLinkedRoleRequest(param0 *iam.CreateServiceLinkedRoleInput) (*request.Request, *iam.CreateServiceLinkedRoleOutput) {
	m.addCall("CreateServiceLinkedRoleRequest")
	m.verifyInput("CreateServiceLinkedRoleRequest", param0)
	return m.CreateServiceLinkedRoleRequestFunc(param0)
}

func (m *iamMock) CreateServiceLinkedRoleWithContext(param0 aws.Context, param1 *iam.CreateServiceLinkedRoleInput, param2 ...request.Option) (*iam.CreateServiceLinkedRoleOutput, error) {
	m.addCall("CreateServiceLinkedRoleWithContext")
	m.verifyInput("CreateServiceLinkedRoleWithContext", param0)
	return m.CreateServiceLinkedRoleWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) CreateServiceSpecificCredential(param0 *iam.CreateServiceSpecificCredentialInput) (*iam.CreateServiceSpecificCredentialOutput, error) {
	m.addCall("CreateServiceSpecificCredential")
	m.verifyInput("CreateServiceSpecificCredential", param0)
	return m.CreateServiceSpecificCredentialFunc(param0)
}

func (m *iamMock) CreateServiceSpecificCredentialRequest(param0 *iam.CreateServiceSpecificCredentialInput) (*request.Request, *iam.CreateServiceSpecificCredentialOutput) {
	m.addCall("CreateServiceSpecificCredentialRequest")
	m.verifyInput("CreateServiceSpecificCredentialRequest", param0)
	return m.CreateServiceSpecificCredentialRequestFunc(param0)
}

func (m *iamMock) CreateServiceSpecificCredentialWithContext(param0 aws.Context, param1 *iam.CreateServiceSpecificCredentialInput, param2 ...request.Option) (*iam.CreateServiceSpecificCredentialOutput, error) {
	m.addCall("CreateServiceSpecificCredentialWithContext")
	m.verifyInput("CreateServiceSpecificCredentialWithContext", param0)
	return m.CreateServiceSpecificCredentialWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) CreateUser(param0 *iam.CreateUserInput) (*iam.CreateUserOutput, error) {
	m.addCall("CreateUser")
	m.verifyInput("CreateUser", param0)
	return m.CreateUserFunc(param0)
}

func (m *iamMock) CreateUserRequest(param0 *iam.CreateUserInput) (*request.Request, *iam.CreateUserOutput) {
	m.addCall("CreateUserRequest")
	m.verifyInput("CreateUserRequest", param0)
	return m.CreateUserRequestFunc(param0)
}

func (m *iamMock) CreateUserWithContext(param0 aws.Context, param1 *iam.CreateUserInput, param2 ...request.Option) (*iam.CreateUserOutput, error) {
	m.addCall("CreateUserWithContext")
	m.verifyInput("CreateUserWithContext", param0)
	return m.CreateUserWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) CreateVirtualMFADevice(param0 *iam.CreateVirtualMFADeviceInput) (*iam.CreateVirtualMFADeviceOutput, error) {
	m.addCall("CreateVirtualMFADevice")
	m.verifyInput("CreateVirtualMFADevice", param0)
	return m.CreateVirtualMFADeviceFunc(param0)
}

func (m *iamMock) CreateVirtualMFADeviceRequest(param0 *iam.CreateVirtualMFADeviceInput) (*request.Request, *iam.CreateVirtualMFADeviceOutput) {
	m.addCall("CreateVirtualMFADeviceRequest")
	m.verifyInput("CreateVirtualMFADeviceRequest", param0)
	return m.CreateVirtualMFADeviceRequestFunc(param0)
}

func (m *iamMock) CreateVirtualMFADeviceWithContext(param0 aws.Context, param1 *iam.CreateVirtualMFADeviceInput, param2 ...request.Option) (*iam.CreateVirtualMFADeviceOutput, error) {
	m.addCall("CreateVirtualMFADeviceWithContext")
	m.verifyInput("CreateVirtualMFADeviceWithContext", param0)
	return m.CreateVirtualMFADeviceWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) DeactivateMFADevice(param0 *iam.DeactivateMFADeviceInput) (*iam.DeactivateMFADeviceOutput, error) {
	m.addCall("DeactivateMFADevice")
	m.verifyInput("DeactivateMFADevice", param0)
	return m.DeactivateMFADeviceFunc(param0)
}

func (m *iamMock) DeactivateMFADeviceRequest(param0 *iam.DeactivateMFADeviceInput) (*request.Request, *iam.DeactivateMFADeviceOutput) {
	m.addCall("DeactivateMFADeviceRequest")
	m.verifyInput("DeactivateMFADeviceRequest", param0)
	return m.DeactivateMFADeviceRequestFunc(param0)
}

func (m *iamMock) DeactivateMFADeviceWithContext(param0 aws.Context, param1 *iam.DeactivateMFADeviceInput, param2 ...request.Option) (*iam.DeactivateMFADeviceOutput, error) {
	m.addCall("DeactivateMFADeviceWithContext")
	m.verifyInput("DeactivateMFADeviceWithContext", param0)
	return m.DeactivateMFADeviceWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) DeleteAccessKey(param0 *iam.DeleteAccessKeyInput) (*iam.DeleteAccessKeyOutput, error) {
	m.addCall("DeleteAccessKey")
	m.verifyInput("DeleteAccessKey", param0)
	return m.DeleteAccessKeyFunc(param0)
}

func (m *iamMock) DeleteAccessKeyRequest(param0 *iam.DeleteAccessKeyInput) (*request.Request, *iam.DeleteAccessKeyOutput) {
	m.addCall("DeleteAccessKeyRequest")
	m.verifyInput("DeleteAccessKeyRequest", param0)
	return m.DeleteAccessKeyRequestFunc(param0)
}

func (m *iamMock) DeleteAccessKeyWithContext(param0 aws.Context, param1 *iam.DeleteAccessKeyInput, param2 ...request.Option) (*iam.DeleteAccessKeyOutput, error) {
	m.addCall("DeleteAccessKeyWithContext")
	m.verifyInput("DeleteAccessKeyWithContext", param0)
	return m.DeleteAccessKeyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) DeleteAccountAlias(param0 *iam.DeleteAccountAliasInput) (*iam.DeleteAccountAliasOutput, error) {
	m.addCall("DeleteAccountAlias")
	m.verifyInput("DeleteAccountAlias", param0)
	return m.DeleteAccountAliasFunc(param0)
}

func (m *iamMock) DeleteAccountAliasRequest(param0 *iam.DeleteAccountAliasInput) (*request.Request, *iam.DeleteAccountAliasOutput) {
	m.addCall("DeleteAccountAliasRequest")
	m.verifyInput("DeleteAccountAliasRequest", param0)
	return m.DeleteAccountAliasRequestFunc(param0)
}

func (m *iamMock) DeleteAccountAliasWithContext(param0 aws.Context, param1 *iam.DeleteAccountAliasInput, param2 ...request.Option) (*iam.DeleteAccountAliasOutput, error) {
	m.addCall("DeleteAccountAliasWithContext")
	m.verifyInput("DeleteAccountAliasWithContext", param0)
	return m.DeleteAccountAliasWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) DeleteAccountPasswordPolicy(param0 *iam.DeleteAccountPasswordPolicyInput) (*iam.DeleteAccountPasswordPolicyOutput, error) {
	m.addCall("DeleteAccountPasswordPolicy")
	m.verifyInput("DeleteAccountPasswordPolicy", param0)
	return m.DeleteAccountPasswordPolicyFunc(param0)
}

func (m *iamMock) DeleteAccountPasswordPolicyRequest(param0 *iam.DeleteAccountPasswordPolicyInput) (*request.Request, *iam.DeleteAccountPasswordPolicyOutput) {
	m.addCall("DeleteAccountPasswordPolicyRequest")
	m.verifyInput("DeleteAccountPasswordPolicyRequest", param0)
	return m.DeleteAccountPasswordPolicyRequestFunc(param0)
}

func (m *iamMock) DeleteAccountPasswordPolicyWithContext(param0 aws.Context, param1 *iam.DeleteAccountPasswordPolicyInput, param2 ...request.Option) (*iam.DeleteAccountPasswordPolicyOutput, error) {
	m.addCall("DeleteAccountPasswordPolicyWithContext")
	m.verifyInput("DeleteAccountPasswordPolicyWithContext", param0)
	return m.DeleteAccountPasswordPolicyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) DeleteGroup(param0 *iam.DeleteGroupInput) (*iam.DeleteGroupOutput, error) {
	m.addCall("DeleteGroup")
	m.verifyInput("DeleteGroup", param0)
	return m.DeleteGroupFunc(param0)
}

func (m *iamMock) DeleteGroupPolicy(param0 *iam.DeleteGroupPolicyInput) (*iam.DeleteGroupPolicyOutput, error) {
	m.addCall("DeleteGroupPolicy")
	m.verifyInput("DeleteGroupPolicy", param0)
	return m.DeleteGroupPolicyFunc(param0)
}

func (m *iamMock) DeleteGroupPolicyRequest(param0 *iam.DeleteGroupPolicyInput) (*request.Request, *iam.DeleteGroupPolicyOutput) {
	m.addCall("DeleteGroupPolicyRequest")
	m.verifyInput("DeleteGroupPolicyRequest", param0)
	return m.DeleteGroupPolicyRequestFunc(param0)
}

func (m *iamMock) DeleteGroupPolicyWithContext(param0 aws.Context, param1 *iam.DeleteGroupPolicyInput, param2 ...request.Option) (*iam.DeleteGroupPolicyOutput, error) {
	m.addCall("DeleteGroupPolicyWithContext")
	m.verifyInput("DeleteGroupPolicyWithContext", param0)
	return m.DeleteGroupPolicyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) DeleteGroupRequest(param0 *iam.DeleteGroupInput) (*request.Request, *iam.DeleteGroupOutput) {
	m.addCall("DeleteGroupRequest")
	m.verifyInput("DeleteGroupRequest", param0)
	return m.DeleteGroupRequestFunc(param0)
}

func (m *iamMock) DeleteGroupWithContext(param0 aws.Context, param1 *iam.DeleteGroupInput, param2 ...request.Option) (*iam.DeleteGroupOutput, error) {
	m.addCall("DeleteGroupWithContext")
	m.verifyInput("DeleteGroupWithContext", param0)
	return m.DeleteGroupWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) DeleteInstanceProfile(param0 *iam.DeleteInstanceProfileInput) (*iam.DeleteInstanceProfileOutput, error) {
	m.addCall("DeleteInstanceProfile")
	m.verifyInput("DeleteInstanceProfile", param0)
	return m.DeleteInstanceProfileFunc(param0)
}

func (m *iamMock) DeleteInstanceProfileRequest(param0 *iam.DeleteInstanceProfileInput) (*request.Request, *iam.DeleteInstanceProfileOutput) {
	m.addCall("DeleteInstanceProfileRequest")
	m.verifyInput("DeleteInstanceProfileRequest", param0)
	return m.DeleteInstanceProfileRequestFunc(param0)
}

func (m *iamMock) DeleteInstanceProfileWithContext(param0 aws.Context, param1 *iam.DeleteInstanceProfileInput, param2 ...request.Option) (*iam.DeleteInstanceProfileOutput, error) {
	m.addCall("DeleteInstanceProfileWithContext")
	m.verifyInput("DeleteInstanceProfileWithContext", param0)
	return m.DeleteInstanceProfileWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) DeleteLoginProfile(param0 *iam.DeleteLoginProfileInput) (*iam.DeleteLoginProfileOutput, error) {
	m.addCall("DeleteLoginProfile")
	m.verifyInput("DeleteLoginProfile", param0)
	return m.DeleteLoginProfileFunc(param0)
}

func (m *iamMock) DeleteLoginProfileRequest(param0 *iam.DeleteLoginProfileInput) (*request.Request, *iam.DeleteLoginProfileOutput) {
	m.addCall("DeleteLoginProfileRequest")
	m.verifyInput("DeleteLoginProfileRequest", param0)
	return m.DeleteLoginProfileRequestFunc(param0)
}

func (m *iamMock) DeleteLoginProfileWithContext(param0 aws.Context, param1 *iam.DeleteLoginProfileInput, param2 ...request.Option) (*iam.DeleteLoginProfileOutput, error) {
	m.addCall("DeleteLoginProfileWithContext")
	m.verifyInput("DeleteLoginProfileWithContext", param0)
	return m.DeleteLoginProfileWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) DeleteOpenIDConnectProvider(param0 *iam.DeleteOpenIDConnectProviderInput) (*iam.DeleteOpenIDConnectProviderOutput, error) {
	m.addCall("DeleteOpenIDConnectProvider")
	m.verifyInput("DeleteOpenIDConnectProvider", param0)
	return m.DeleteOpenIDConnectProviderFunc(param0)
}

func (m *iamMock) DeleteOpenIDConnectProviderRequest(param0 *iam.DeleteOpenIDConnectProviderInput) (*request.Request, *iam.DeleteOpenIDConnectProviderOutput) {
	m.addCall("DeleteOpenIDConnectProviderRequest")
	m.verifyInput("DeleteOpenIDConnectProviderRequest", param0)
	return m.DeleteOpenIDConnectProviderRequestFunc(param0)
}

func (m *iamMock) DeleteOpenIDConnectProviderWithContext(param0 aws.Context, param1 *iam.DeleteOpenIDConnectProviderInput, param2 ...request.Option) (*iam.DeleteOpenIDConnectProviderOutput, error) {
	m.addCall("DeleteOpenIDConnectProviderWithContext")
	m.verifyInput("DeleteOpenIDConnectProviderWithContext", param0)
	return m.DeleteOpenIDConnectProviderWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) DeletePolicy(param0 *iam.DeletePolicyInput) (*iam.DeletePolicyOutput, error) {
	m.addCall("DeletePolicy")
	m.verifyInput("DeletePolicy", param0)
	return m.DeletePolicyFunc(param0)
}

func (m *iamMock) DeletePolicyRequest(param0 *iam.DeletePolicyInput) (*request.Request, *iam.DeletePolicyOutput) {
	m.addCall("DeletePolicyRequest")
	m.verifyInput("DeletePolicyRequest", param0)
	return m.DeletePolicyRequestFunc(param0)
}

func (m *iamMock) DeletePolicyVersion(param0 *iam.DeletePolicyVersionInput) (*iam.DeletePolicyVersionOutput, error) {
	m.addCall("DeletePolicyVersion")
	m.verifyInput("DeletePolicyVersion", param0)
	return m.DeletePolicyVersionFunc(param0)
}

func (m *iamMock) DeletePolicyVersionRequest(param0 *iam.DeletePolicyVersionInput) (*request.Request, *iam.DeletePolicyVersionOutput) {
	m.addCall("DeletePolicyVersionRequest")
	m.verifyInput("DeletePolicyVersionRequest", param0)
	return m.DeletePolicyVersionRequestFunc(param0)
}

func (m *iamMock) DeletePolicyVersionWithContext(param0 aws.Context, param1 *iam.DeletePolicyVersionInput, param2 ...request.Option) (*iam.DeletePolicyVersionOutput, error) {
	m.addCall("DeletePolicyVersionWithContext")
	m.verifyInput("DeletePolicyVersionWithContext", param0)
	return m.DeletePolicyVersionWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) DeletePolicyWithContext(param0 aws.Context, param1 *iam.DeletePolicyInput, param2 ...request.Option) (*iam.DeletePolicyOutput, error) {
	m.addCall("DeletePolicyWithContext")
	m.verifyInput("DeletePolicyWithContext", param0)
	return m.DeletePolicyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) DeleteRole(param0 *iam.DeleteRoleInput) (*iam.DeleteRoleOutput, error) {
	m.addCall("DeleteRole")
	m.verifyInput("DeleteRole", param0)
	return m.DeleteRoleFunc(param0)
}

func (m *iamMock) DeleteRolePolicy(param0 *iam.DeleteRolePolicyInput) (*iam.DeleteRolePolicyOutput, error) {
	m.addCall("DeleteRolePolicy")
	m.verifyInput("DeleteRolePolicy", param0)
	return m.DeleteRolePolicyFunc(param0)
}

func (m *iamMock) DeleteRolePolicyRequest(param0 *iam.DeleteRolePolicyInput) (*request.Request, *iam.DeleteRolePolicyOutput) {
	m.addCall("DeleteRolePolicyRequest")
	m.verifyInput("DeleteRolePolicyRequest", param0)
	return m.DeleteRolePolicyRequestFunc(param0)
}

func (m *iamMock) DeleteRolePolicyWithContext(param0 aws.Context, param1 *iam.DeleteRolePolicyInput, param2 ...request.Option) (*iam.DeleteRolePolicyOutput, error) {
	m.addCall("DeleteRolePolicyWithContext")
	m.verifyInput("DeleteRolePolicyWithContext", param0)
	return m.DeleteRolePolicyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) DeleteRoleRequest(param0 *iam.DeleteRoleInput) (*request.Request, *iam.DeleteRoleOutput) {
	m.addCall("DeleteRoleRequest")
	m.verifyInput("DeleteRoleRequest", param0)
	return m.DeleteRoleRequestFunc(param0)
}

func (m *iamMock) DeleteRoleWithContext(param0 aws.Context, param1 *iam.DeleteRoleInput, param2 ...request.Option) (*iam.DeleteRoleOutput, error) {
	m.addCall("DeleteRoleWithContext")
	m.verifyInput("DeleteRoleWithContext", param0)
	return m.DeleteRoleWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) DeleteSAMLProvider(param0 *iam.DeleteSAMLProviderInput) (*iam.DeleteSAMLProviderOutput, error) {
	m.addCall("DeleteSAMLProvider")
	m.verifyInput("DeleteSAMLProvider", param0)
	return m.DeleteSAMLProviderFunc(param0)
}

func (m *iamMock) DeleteSAMLProviderRequest(param0 *iam.DeleteSAMLProviderInput) (*request.Request, *iam.DeleteSAMLProviderOutput) {
	m.addCall("DeleteSAMLProviderRequest")
	m.verifyInput("DeleteSAMLProviderRequest", param0)
	return m.DeleteSAMLProviderRequestFunc(param0)
}

func (m *iamMock) DeleteSAMLProviderWithContext(param0 aws.Context, param1 *iam.DeleteSAMLProviderInput, param2 ...request.Option) (*iam.DeleteSAMLProviderOutput, error) {
	m.addCall("DeleteSAMLProviderWithContext")
	m.verifyInput("DeleteSAMLProviderWithContext", param0)
	return m.DeleteSAMLProviderWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) DeleteSSHPublicKey(param0 *iam.DeleteSSHPublicKeyInput) (*iam.DeleteSSHPublicKeyOutput, error) {
	m.addCall("DeleteSSHPublicKey")
	m.verifyInput("DeleteSSHPublicKey", param0)
	return m.DeleteSSHPublicKeyFunc(param0)
}

func (m *iamMock) DeleteSSHPublicKeyRequest(param0 *iam.DeleteSSHPublicKeyInput) (*request.Request, *iam.DeleteSSHPublicKeyOutput) {
	m.addCall("DeleteSSHPublicKeyRequest")
	m.verifyInput("DeleteSSHPublicKeyRequest", param0)
	return m.DeleteSSHPublicKeyRequestFunc(param0)
}

func (m *iamMock) DeleteSSHPublicKeyWithContext(param0 aws.Context, param1 *iam.DeleteSSHPublicKeyInput, param2 ...request.Option) (*iam.DeleteSSHPublicKeyOutput, error) {
	m.addCall("DeleteSSHPublicKeyWithContext")
	m.verifyInput("DeleteSSHPublicKeyWithContext", param0)
	return m.DeleteSSHPublicKeyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) DeleteServerCertificate(param0 *iam.DeleteServerCertificateInput) (*iam.DeleteServerCertificateOutput, error) {
	m.addCall("DeleteServerCertificate")
	m.verifyInput("DeleteServerCertificate", param0)
	return m.DeleteServerCertificateFunc(param0)
}

func (m *iamMock) DeleteServerCertificateRequest(param0 *iam.DeleteServerCertificateInput) (*request.Request, *iam.DeleteServerCertificateOutput) {
	m.addCall("DeleteServerCertificateRequest")
	m.verifyInput("DeleteServerCertificateRequest", param0)
	return m.DeleteServerCertificateRequestFunc(param0)
}

func (m *iamMock) DeleteServerCertificateWithContext(param0 aws.Context, param1 *iam.DeleteServerCertificateInput, param2 ...request.Option) (*iam.DeleteServerCertificateOutput, error) {
	m.addCall("DeleteServerCertificateWithContext")
	m.verifyInput("DeleteServerCertificateWithContext", param0)
	return m.DeleteServerCertificateWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) DeleteServiceLinkedRole(param0 *iam.DeleteServiceLinkedRoleInput) (*iam.DeleteServiceLinkedRoleOutput, error) {
	m.addCall("DeleteServiceLinkedRole")
	m.verifyInput("DeleteServiceLinkedRole", param0)
	return m.DeleteServiceLinkedRoleFunc(param0)
}

func (m *iamMock) DeleteServiceLinkedRoleRequest(param0 *iam.DeleteServiceLinkedRoleInput) (*request.Request, *iam.DeleteServiceLinkedRoleOutput) {
	m.addCall("DeleteServiceLinkedRoleRequest")
	m.verifyInput("DeleteServiceLinkedRoleRequest", param0)
	return m.DeleteServiceLinkedRoleRequestFunc(param0)
}

func (m *iamMock) DeleteServiceLinkedRoleWithContext(param0 aws.Context, param1 *iam.DeleteServiceLinkedRoleInput, param2 ...request.Option) (*iam.DeleteServiceLinkedRoleOutput, error) {
	m.addCall("DeleteServiceLinkedRoleWithContext")
	m.verifyInput("DeleteServiceLinkedRoleWithContext", param0)
	return m.DeleteServiceLinkedRoleWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) DeleteServiceSpecificCredential(param0 *iam.DeleteServiceSpecificCredentialInput) (*iam.DeleteServiceSpecificCredentialOutput, error) {
	m.addCall("DeleteServiceSpecificCredential")
	m.verifyInput("DeleteServiceSpecificCredential", param0)
	return m.DeleteServiceSpecificCredentialFunc(param0)
}

func (m *iamMock) DeleteServiceSpecificCredentialRequest(param0 *iam.DeleteServiceSpecificCredentialInput) (*request.Request, *iam.DeleteServiceSpecificCredentialOutput) {
	m.addCall("DeleteServiceSpecificCredentialRequest")
	m.verifyInput("DeleteServiceSpecificCredentialRequest", param0)
	return m.DeleteServiceSpecificCredentialRequestFunc(param0)
}

func (m *iamMock) DeleteServiceSpecificCredentialWithContext(param0 aws.Context, param1 *iam.DeleteServiceSpecificCredentialInput, param2 ...request.Option) (*iam.DeleteServiceSpecificCredentialOutput, error) {
	m.addCall("DeleteServiceSpecificCredentialWithContext")
	m.verifyInput("DeleteServiceSpecificCredentialWithContext", param0)
	return m.DeleteServiceSpecificCredentialWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) DeleteSigningCertificate(param0 *iam.DeleteSigningCertificateInput) (*iam.DeleteSigningCertificateOutput, error) {
	m.addCall("DeleteSigningCertificate")
	m.verifyInput("DeleteSigningCertificate", param0)
	return m.DeleteSigningCertificateFunc(param0)
}

func (m *iamMock) DeleteSigningCertificateRequest(param0 *iam.DeleteSigningCertificateInput) (*request.Request, *iam.DeleteSigningCertificateOutput) {
	m.addCall("DeleteSigningCertificateRequest")
	m.verifyInput("DeleteSigningCertificateRequest", param0)
	return m.DeleteSigningCertificateRequestFunc(param0)
}

func (m *iamMock) DeleteSigningCertificateWithContext(param0 aws.Context, param1 *iam.DeleteSigningCertificateInput, param2 ...request.Option) (*iam.DeleteSigningCertificateOutput, error) {
	m.addCall("DeleteSigningCertificateWithContext")
	m.verifyInput("DeleteSigningCertificateWithContext", param0)
	return m.DeleteSigningCertificateWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) DeleteUser(param0 *iam.DeleteUserInput) (*iam.DeleteUserOutput, error) {
	m.addCall("DeleteUser")
	m.verifyInput("DeleteUser", param0)
	return m.DeleteUserFunc(param0)
}

func (m *iamMock) DeleteUserPolicy(param0 *iam.DeleteUserPolicyInput) (*iam.DeleteUserPolicyOutput, error) {
	m.addCall("DeleteUserPolicy")
	m.verifyInput("DeleteUserPolicy", param0)
	return m.DeleteUserPolicyFunc(param0)
}

func (m *iamMock) DeleteUserPolicyRequest(param0 *iam.DeleteUserPolicyInput) (*request.Request, *iam.DeleteUserPolicyOutput) {
	m.addCall("DeleteUserPolicyRequest")
	m.verifyInput("DeleteUserPolicyRequest", param0)
	return m.DeleteUserPolicyRequestFunc(param0)
}

func (m *iamMock) DeleteUserPolicyWithContext(param0 aws.Context, param1 *iam.DeleteUserPolicyInput, param2 ...request.Option) (*iam.DeleteUserPolicyOutput, error) {
	m.addCall("DeleteUserPolicyWithContext")
	m.verifyInput("DeleteUserPolicyWithContext", param0)
	return m.DeleteUserPolicyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) DeleteUserRequest(param0 *iam.DeleteUserInput) (*request.Request, *iam.DeleteUserOutput) {
	m.addCall("DeleteUserRequest")
	m.verifyInput("DeleteUserRequest", param0)
	return m.DeleteUserRequestFunc(param0)
}

func (m *iamMock) DeleteUserWithContext(param0 aws.Context, param1 *iam.DeleteUserInput, param2 ...request.Option) (*iam.DeleteUserOutput, error) {
	m.addCall("DeleteUserWithContext")
	m.verifyInput("DeleteUserWithContext", param0)
	return m.DeleteUserWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) DeleteVirtualMFADevice(param0 *iam.DeleteVirtualMFADeviceInput) (*iam.DeleteVirtualMFADeviceOutput, error) {
	m.addCall("DeleteVirtualMFADevice")
	m.verifyInput("DeleteVirtualMFADevice", param0)
	return m.DeleteVirtualMFADeviceFunc(param0)
}

func (m *iamMock) DeleteVirtualMFADeviceRequest(param0 *iam.DeleteVirtualMFADeviceInput) (*request.Request, *iam.DeleteVirtualMFADeviceOutput) {
	m.addCall("DeleteVirtualMFADeviceRequest")
	m.verifyInput("DeleteVirtualMFADeviceRequest", param0)
	return m.DeleteVirtualMFADeviceRequestFunc(param0)
}

func (m *iamMock) DeleteVirtualMFADeviceWithContext(param0 aws.Context, param1 *iam.DeleteVirtualMFADeviceInput, param2 ...request.Option) (*iam.DeleteVirtualMFADeviceOutput, error) {
	m.addCall("DeleteVirtualMFADeviceWithContext")
	m.verifyInput("DeleteVirtualMFADeviceWithContext", param0)
	return m.DeleteVirtualMFADeviceWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) DetachGroupPolicy(param0 *iam.DetachGroupPolicyInput) (*iam.DetachGroupPolicyOutput, error) {
	m.addCall("DetachGroupPolicy")
	m.verifyInput("DetachGroupPolicy", param0)
	return m.DetachGroupPolicyFunc(param0)
}

func (m *iamMock) DetachGroupPolicyRequest(param0 *iam.DetachGroupPolicyInput) (*request.Request, *iam.DetachGroupPolicyOutput) {
	m.addCall("DetachGroupPolicyRequest")
	m.verifyInput("DetachGroupPolicyRequest", param0)
	return m.DetachGroupPolicyRequestFunc(param0)
}

func (m *iamMock) DetachGroupPolicyWithContext(param0 aws.Context, param1 *iam.DetachGroupPolicyInput, param2 ...request.Option) (*iam.DetachGroupPolicyOutput, error) {
	m.addCall("DetachGroupPolicyWithContext")
	m.verifyInput("DetachGroupPolicyWithContext", param0)
	return m.DetachGroupPolicyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) DetachRolePolicy(param0 *iam.DetachRolePolicyInput) (*iam.DetachRolePolicyOutput, error) {
	m.addCall("DetachRolePolicy")
	m.verifyInput("DetachRolePolicy", param0)
	return m.DetachRolePolicyFunc(param0)
}

func (m *iamMock) DetachRolePolicyRequest(param0 *iam.DetachRolePolicyInput) (*request.Request, *iam.DetachRolePolicyOutput) {
	m.addCall("DetachRolePolicyRequest")
	m.verifyInput("DetachRolePolicyRequest", param0)
	return m.DetachRolePolicyRequestFunc(param0)
}

func (m *iamMock) DetachRolePolicyWithContext(param0 aws.Context, param1 *iam.DetachRolePolicyInput, param2 ...request.Option) (*iam.DetachRolePolicyOutput, error) {
	m.addCall("DetachRolePolicyWithContext")
	m.verifyInput("DetachRolePolicyWithContext", param0)
	return m.DetachRolePolicyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) DetachUserPolicy(param0 *iam.DetachUserPolicyInput) (*iam.DetachUserPolicyOutput, error) {
	m.addCall("DetachUserPolicy")
	m.verifyInput("DetachUserPolicy", param0)
	return m.DetachUserPolicyFunc(param0)
}

func (m *iamMock) DetachUserPolicyRequest(param0 *iam.DetachUserPolicyInput) (*request.Request, *iam.DetachUserPolicyOutput) {
	m.addCall("DetachUserPolicyRequest")
	m.verifyInput("DetachUserPolicyRequest", param0)
	return m.DetachUserPolicyRequestFunc(param0)
}

func (m *iamMock) DetachUserPolicyWithContext(param0 aws.Context, param1 *iam.DetachUserPolicyInput, param2 ...request.Option) (*iam.DetachUserPolicyOutput, error) {
	m.addCall("DetachUserPolicyWithContext")
	m.verifyInput("DetachUserPolicyWithContext", param0)
	return m.DetachUserPolicyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) EnableMFADevice(param0 *iam.EnableMFADeviceInput) (*iam.EnableMFADeviceOutput, error) {
	m.addCall("EnableMFADevice")
	m.verifyInput("EnableMFADevice", param0)
	return m.EnableMFADeviceFunc(param0)
}

func (m *iamMock) EnableMFADeviceRequest(param0 *iam.EnableMFADeviceInput) (*request.Request, *iam.EnableMFADeviceOutput) {
	m.addCall("EnableMFADeviceRequest")
	m.verifyInput("EnableMFADeviceRequest", param0)
	return m.EnableMFADeviceRequestFunc(param0)
}

func (m *iamMock) EnableMFADeviceWithContext(param0 aws.Context, param1 *iam.EnableMFADeviceInput, param2 ...request.Option) (*iam.EnableMFADeviceOutput, error) {
	m.addCall("EnableMFADeviceWithContext")
	m.verifyInput("EnableMFADeviceWithContext", param0)
	return m.EnableMFADeviceWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) GenerateCredentialReport(param0 *iam.GenerateCredentialReportInput) (*iam.GenerateCredentialReportOutput, error) {
	m.addCall("GenerateCredentialReport")
	m.verifyInput("GenerateCredentialReport", param0)
	return m.GenerateCredentialReportFunc(param0)
}

func (m *iamMock) GenerateCredentialReportRequest(param0 *iam.GenerateCredentialReportInput) (*request.Request, *iam.GenerateCredentialReportOutput) {
	m.addCall("GenerateCredentialReportRequest")
	m.verifyInput("GenerateCredentialReportRequest", param0)
	return m.GenerateCredentialReportRequestFunc(param0)
}

func (m *iamMock) GenerateCredentialReportWithContext(param0 aws.Context, param1 *iam.GenerateCredentialReportInput, param2 ...request.Option) (*iam.GenerateCredentialReportOutput, error) {
	m.addCall("GenerateCredentialReportWithContext")
	m.verifyInput("GenerateCredentialReportWithContext", param0)
	return m.GenerateCredentialReportWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) GetAccessKeyLastUsed(param0 *iam.GetAccessKeyLastUsedInput) (*iam.GetAccessKeyLastUsedOutput, error) {
	m.addCall("GetAccessKeyLastUsed")
	m.verifyInput("GetAccessKeyLastUsed", param0)
	return m.GetAccessKeyLastUsedFunc(param0)
}

func (m *iamMock) GetAccessKeyLastUsedRequest(param0 *iam.GetAccessKeyLastUsedInput) (*request.Request, *iam.GetAccessKeyLastUsedOutput) {
	m.addCall("GetAccessKeyLastUsedRequest")
	m.verifyInput("GetAccessKeyLastUsedRequest", param0)
	return m.GetAccessKeyLastUsedRequestFunc(param0)
}

func (m *iamMock) GetAccessKeyLastUsedWithContext(param0 aws.Context, param1 *iam.GetAccessKeyLastUsedInput, param2 ...request.Option) (*iam.GetAccessKeyLastUsedOutput, error) {
	m.addCall("GetAccessKeyLastUsedWithContext")
	m.verifyInput("GetAccessKeyLastUsedWithContext", param0)
	return m.GetAccessKeyLastUsedWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) GetAccountAuthorizationDetails(param0 *iam.GetAccountAuthorizationDetailsInput) (*iam.GetAccountAuthorizationDetailsOutput, error) {
	m.addCall("GetAccountAuthorizationDetails")
	m.verifyInput("GetAccountAuthorizationDetails", param0)
	return m.GetAccountAuthorizationDetailsFunc(param0)
}

func (m *iamMock) GetAccountAuthorizationDetailsRequest(param0 *iam.GetAccountAuthorizationDetailsInput) (*request.Request, *iam.GetAccountAuthorizationDetailsOutput) {
	m.addCall("GetAccountAuthorizationDetailsRequest")
	m.verifyInput("GetAccountAuthorizationDetailsRequest", param0)
	return m.GetAccountAuthorizationDetailsRequestFunc(param0)
}

func (m *iamMock) GetAccountAuthorizationDetailsWithContext(param0 aws.Context, param1 *iam.GetAccountAuthorizationDetailsInput, param2 ...request.Option) (*iam.GetAccountAuthorizationDetailsOutput, error) {
	m.addCall("GetAccountAuthorizationDetailsWithContext")
	m.verifyInput("GetAccountAuthorizationDetailsWithContext", param0)
	return m.GetAccountAuthorizationDetailsWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) GetAccountPasswordPolicy(param0 *iam.GetAccountPasswordPolicyInput) (*iam.GetAccountPasswordPolicyOutput, error) {
	m.addCall("GetAccountPasswordPolicy")
	m.verifyInput("GetAccountPasswordPolicy", param0)
	return m.GetAccountPasswordPolicyFunc(param0)
}

func (m *iamMock) GetAccountPasswordPolicyRequest(param0 *iam.GetAccountPasswordPolicyInput) (*request.Request, *iam.GetAccountPasswordPolicyOutput) {
	m.addCall("GetAccountPasswordPolicyRequest")
	m.verifyInput("GetAccountPasswordPolicyRequest", param0)
	return m.GetAccountPasswordPolicyRequestFunc(param0)
}

func (m *iamMock) GetAccountPasswordPolicyWithContext(param0 aws.Context, param1 *iam.GetAccountPasswordPolicyInput, param2 ...request.Option) (*iam.GetAccountPasswordPolicyOutput, error) {
	m.addCall("GetAccountPasswordPolicyWithContext")
	m.verifyInput("GetAccountPasswordPolicyWithContext", param0)
	return m.GetAccountPasswordPolicyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) GetAccountSummary(param0 *iam.GetAccountSummaryInput) (*iam.GetAccountSummaryOutput, error) {
	m.addCall("GetAccountSummary")
	m.verifyInput("GetAccountSummary", param0)
	return m.GetAccountSummaryFunc(param0)
}

func (m *iamMock) GetAccountSummaryRequest(param0 *iam.GetAccountSummaryInput) (*request.Request, *iam.GetAccountSummaryOutput) {
	m.addCall("GetAccountSummaryRequest")
	m.verifyInput("GetAccountSummaryRequest", param0)
	return m.GetAccountSummaryRequestFunc(param0)
}

func (m *iamMock) GetAccountSummaryWithContext(param0 aws.Context, param1 *iam.GetAccountSummaryInput, param2 ...request.Option) (*iam.GetAccountSummaryOutput, error) {
	m.addCall("GetAccountSummaryWithContext")
	m.verifyInput("GetAccountSummaryWithContext", param0)
	return m.GetAccountSummaryWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) GetContextKeysForCustomPolicy(param0 *iam.GetContextKeysForCustomPolicyInput) (*iam.GetContextKeysForPolicyResponse, error) {
	m.addCall("GetContextKeysForCustomPolicy")
	m.verifyInput("GetContextKeysForCustomPolicy", param0)
	return m.GetContextKeysForCustomPolicyFunc(param0)
}

func (m *iamMock) GetContextKeysForCustomPolicyRequest(param0 *iam.GetContextKeysForCustomPolicyInput) (*request.Request, *iam.GetContextKeysForPolicyResponse) {
	m.addCall("GetContextKeysForCustomPolicyRequest")
	m.verifyInput("GetContextKeysForCustomPolicyRequest", param0)
	return m.GetContextKeysForCustomPolicyRequestFunc(param0)
}

func (m *iamMock) GetContextKeysForCustomPolicyWithContext(param0 aws.Context, param1 *iam.GetContextKeysForCustomPolicyInput, param2 ...request.Option) (*iam.GetContextKeysForPolicyResponse, error) {
	m.addCall("GetContextKeysForCustomPolicyWithContext")
	m.verifyInput("GetContextKeysForCustomPolicyWithContext", param0)
	return m.GetContextKeysForCustomPolicyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) GetContextKeysForPrincipalPolicy(param0 *iam.GetContextKeysForPrincipalPolicyInput) (*iam.GetContextKeysForPolicyResponse, error) {
	m.addCall("GetContextKeysForPrincipalPolicy")
	m.verifyInput("GetContextKeysForPrincipalPolicy", param0)
	return m.GetContextKeysForPrincipalPolicyFunc(param0)
}

func (m *iamMock) GetContextKeysForPrincipalPolicyRequest(param0 *iam.GetContextKeysForPrincipalPolicyInput) (*request.Request, *iam.GetContextKeysForPolicyResponse) {
	m.addCall("GetContextKeysForPrincipalPolicyRequest")
	m.verifyInput("GetContextKeysForPrincipalPolicyRequest", param0)
	return m.GetContextKeysForPrincipalPolicyRequestFunc(param0)
}

func (m *iamMock) GetContextKeysForPrincipalPolicyWithContext(param0 aws.Context, param1 *iam.GetContextKeysForPrincipalPolicyInput, param2 ...request.Option) (*iam.GetContextKeysForPolicyResponse, error) {
	m.addCall("GetContextKeysForPrincipalPolicyWithContext")
	m.verifyInput("GetContextKeysForPrincipalPolicyWithContext", param0)
	return m.GetContextKeysForPrincipalPolicyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) GetCredentialReport(param0 *iam.GetCredentialReportInput) (*iam.GetCredentialReportOutput, error) {
	m.addCall("GetCredentialReport")
	m.verifyInput("GetCredentialReport", param0)
	return m.GetCredentialReportFunc(param0)
}

func (m *iamMock) GetCredentialReportRequest(param0 *iam.GetCredentialReportInput) (*request.Request, *iam.GetCredentialReportOutput) {
	m.addCall("GetCredentialReportRequest")
	m.verifyInput("GetCredentialReportRequest", param0)
	return m.GetCredentialReportRequestFunc(param0)
}

func (m *iamMock) GetCredentialReportWithContext(param0 aws.Context, param1 *iam.GetCredentialReportInput, param2 ...request.Option) (*iam.GetCredentialReportOutput, error) {
	m.addCall("GetCredentialReportWithContext")
	m.verifyInput("GetCredentialReportWithContext", param0)
	return m.GetCredentialReportWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) GetGroup(param0 *iam.GetGroupInput) (*iam.GetGroupOutput, error) {
	m.addCall("GetGroup")
	m.verifyInput("GetGroup", param0)
	return m.GetGroupFunc(param0)
}

func (m *iamMock) GetGroupPolicy(param0 *iam.GetGroupPolicyInput) (*iam.GetGroupPolicyOutput, error) {
	m.addCall("GetGroupPolicy")
	m.verifyInput("GetGroupPolicy", param0)
	return m.GetGroupPolicyFunc(param0)
}

func (m *iamMock) GetGroupPolicyRequest(param0 *iam.GetGroupPolicyInput) (*request.Request, *iam.GetGroupPolicyOutput) {
	m.addCall("GetGroupPolicyRequest")
	m.verifyInput("GetGroupPolicyRequest", param0)
	return m.GetGroupPolicyRequestFunc(param0)
}

func (m *iamMock) GetGroupPolicyWithContext(param0 aws.Context, param1 *iam.GetGroupPolicyInput, param2 ...request.Option) (*iam.GetGroupPolicyOutput, error) {
	m.addCall("GetGroupPolicyWithContext")
	m.verifyInput("GetGroupPolicyWithContext", param0)
	return m.GetGroupPolicyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) GetGroupRequest(param0 *iam.GetGroupInput) (*request.Request, *iam.GetGroupOutput) {
	m.addCall("GetGroupRequest")
	m.verifyInput("GetGroupRequest", param0)
	return m.GetGroupRequestFunc(param0)
}

func (m *iamMock) GetGroupWithContext(param0 aws.Context, param1 *iam.GetGroupInput, param2 ...request.Option) (*iam.GetGroupOutput, error) {
	m.addCall("GetGroupWithContext")
	m.verifyInput("GetGroupWithContext", param0)
	return m.GetGroupWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) GetInstanceProfile(param0 *iam.GetInstanceProfileInput) (*iam.GetInstanceProfileOutput, error) {
	m.addCall("GetInstanceProfile")
	m.verifyInput("GetInstanceProfile", param0)
	return m.GetInstanceProfileFunc(param0)
}

func (m *iamMock) GetInstanceProfileRequest(param0 *iam.GetInstanceProfileInput) (*request.Request, *iam.GetInstanceProfileOutput) {
	m.addCall("GetInstanceProfileRequest")
	m.verifyInput("GetInstanceProfileRequest", param0)
	return m.GetInstanceProfileRequestFunc(param0)
}

func (m *iamMock) GetInstanceProfileWithContext(param0 aws.Context, param1 *iam.GetInstanceProfileInput, param2 ...request.Option) (*iam.GetInstanceProfileOutput, error) {
	m.addCall("GetInstanceProfileWithContext")
	m.verifyInput("GetInstanceProfileWithContext", param0)
	return m.GetInstanceProfileWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) GetLoginProfile(param0 *iam.GetLoginProfileInput) (*iam.GetLoginProfileOutput, error) {
	m.addCall("GetLoginProfile")
	m.verifyInput("GetLoginProfile", param0)
	return m.GetLoginProfileFunc(param0)
}

func (m *iamMock) GetLoginProfileRequest(param0 *iam.GetLoginProfileInput) (*request.Request, *iam.GetLoginProfileOutput) {
	m.addCall("GetLoginProfileRequest")
	m.verifyInput("GetLoginProfileRequest", param0)
	return m.GetLoginProfileRequestFunc(param0)
}

func (m *iamMock) GetLoginProfileWithContext(param0 aws.Context, param1 *iam.GetLoginProfileInput, param2 ...request.Option) (*iam.GetLoginProfileOutput, error) {
	m.addCall("GetLoginProfileWithContext")
	m.verifyInput("GetLoginProfileWithContext", param0)
	return m.GetLoginProfileWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) GetOpenIDConnectProvider(param0 *iam.GetOpenIDConnectProviderInput) (*iam.GetOpenIDConnectProviderOutput, error) {
	m.addCall("GetOpenIDConnectProvider")
	m.verifyInput("GetOpenIDConnectProvider", param0)
	return m.GetOpenIDConnectProviderFunc(param0)
}

func (m *iamMock) GetOpenIDConnectProviderRequest(param0 *iam.GetOpenIDConnectProviderInput) (*request.Request, *iam.GetOpenIDConnectProviderOutput) {
	m.addCall("GetOpenIDConnectProviderRequest")
	m.verifyInput("GetOpenIDConnectProviderRequest", param0)
	return m.GetOpenIDConnectProviderRequestFunc(param0)
}

func (m *iamMock) GetOpenIDConnectProviderWithContext(param0 aws.Context, param1 *iam.GetOpenIDConnectProviderInput, param2 ...request.Option) (*iam.GetOpenIDConnectProviderOutput, error) {
	m.addCall("GetOpenIDConnectProviderWithContext")
	m.verifyInput("GetOpenIDConnectProviderWithContext", param0)
	return m.GetOpenIDConnectProviderWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) GetPolicy(param0 *iam.GetPolicyInput) (*iam.GetPolicyOutput, error) {
	m.addCall("GetPolicy")
	m.verifyInput("GetPolicy", param0)
	return m.GetPolicyFunc(param0)
}

func (m *iamMock) GetPolicyRequest(param0 *iam.GetPolicyInput) (*request.Request, *iam.GetPolicyOutput) {
	m.addCall("GetPolicyRequest")
	m.verifyInput("GetPolicyRequest", param0)
	return m.GetPolicyRequestFunc(param0)
}

func (m *iamMock) GetPolicyVersion(param0 *iam.GetPolicyVersionInput) (*iam.GetPolicyVersionOutput, error) {
	m.addCall("GetPolicyVersion")
	m.verifyInput("GetPolicyVersion", param0)
	return m.GetPolicyVersionFunc(param0)
}

func (m *iamMock) GetPolicyVersionRequest(param0 *iam.GetPolicyVersionInput) (*request.Request, *iam.GetPolicyVersionOutput) {
	m.addCall("GetPolicyVersionRequest")
	m.verifyInput("GetPolicyVersionRequest", param0)
	return m.GetPolicyVersionRequestFunc(param0)
}

func (m *iamMock) GetPolicyVersionWithContext(param0 aws.Context, param1 *iam.GetPolicyVersionInput, param2 ...request.Option) (*iam.GetPolicyVersionOutput, error) {
	m.addCall("GetPolicyVersionWithContext")
	m.verifyInput("GetPolicyVersionWithContext", param0)
	return m.GetPolicyVersionWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) GetPolicyWithContext(param0 aws.Context, param1 *iam.GetPolicyInput, param2 ...request.Option) (*iam.GetPolicyOutput, error) {
	m.addCall("GetPolicyWithContext")
	m.verifyInput("GetPolicyWithContext", param0)
	return m.GetPolicyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) GetRole(param0 *iam.GetRoleInput) (*iam.GetRoleOutput, error) {
	m.addCall("GetRole")
	m.verifyInput("GetRole", param0)
	return m.GetRoleFunc(param0)
}

func (m *iamMock) GetRolePolicy(param0 *iam.GetRolePolicyInput) (*iam.GetRolePolicyOutput, error) {
	m.addCall("GetRolePolicy")
	m.verifyInput("GetRolePolicy", param0)
	return m.GetRolePolicyFunc(param0)
}

func (m *iamMock) GetRolePolicyRequest(param0 *iam.GetRolePolicyInput) (*request.Request, *iam.GetRolePolicyOutput) {
	m.addCall("GetRolePolicyRequest")
	m.verifyInput("GetRolePolicyRequest", param0)
	return m.GetRolePolicyRequestFunc(param0)
}

func (m *iamMock) GetRolePolicyWithContext(param0 aws.Context, param1 *iam.GetRolePolicyInput, param2 ...request.Option) (*iam.GetRolePolicyOutput, error) {
	m.addCall("GetRolePolicyWithContext")
	m.verifyInput("GetRolePolicyWithContext", param0)
	return m.GetRolePolicyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) GetRoleRequest(param0 *iam.GetRoleInput) (*request.Request, *iam.GetRoleOutput) {
	m.addCall("GetRoleRequest")
	m.verifyInput("GetRoleRequest", param0)
	return m.GetRoleRequestFunc(param0)
}

func (m *iamMock) GetRoleWithContext(param0 aws.Context, param1 *iam.GetRoleInput, param2 ...request.Option) (*iam.GetRoleOutput, error) {
	m.addCall("GetRoleWithContext")
	m.verifyInput("GetRoleWithContext", param0)
	return m.GetRoleWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) GetSAMLProvider(param0 *iam.GetSAMLProviderInput) (*iam.GetSAMLProviderOutput, error) {
	m.addCall("GetSAMLProvider")
	m.verifyInput("GetSAMLProvider", param0)
	return m.GetSAMLProviderFunc(param0)
}

func (m *iamMock) GetSAMLProviderRequest(param0 *iam.GetSAMLProviderInput) (*request.Request, *iam.GetSAMLProviderOutput) {
	m.addCall("GetSAMLProviderRequest")
	m.verifyInput("GetSAMLProviderRequest", param0)
	return m.GetSAMLProviderRequestFunc(param0)
}

func (m *iamMock) GetSAMLProviderWithContext(param0 aws.Context, param1 *iam.GetSAMLProviderInput, param2 ...request.Option) (*iam.GetSAMLProviderOutput, error) {
	m.addCall("GetSAMLProviderWithContext")
	m.verifyInput("GetSAMLProviderWithContext", param0)
	return m.GetSAMLProviderWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) GetSSHPublicKey(param0 *iam.GetSSHPublicKeyInput) (*iam.GetSSHPublicKeyOutput, error) {
	m.addCall("GetSSHPublicKey")
	m.verifyInput("GetSSHPublicKey", param0)
	return m.GetSSHPublicKeyFunc(param0)
}

func (m *iamMock) GetSSHPublicKeyRequest(param0 *iam.GetSSHPublicKeyInput) (*request.Request, *iam.GetSSHPublicKeyOutput) {
	m.addCall("GetSSHPublicKeyRequest")
	m.verifyInput("GetSSHPublicKeyRequest", param0)
	return m.GetSSHPublicKeyRequestFunc(param0)
}

func (m *iamMock) GetSSHPublicKeyWithContext(param0 aws.Context, param1 *iam.GetSSHPublicKeyInput, param2 ...request.Option) (*iam.GetSSHPublicKeyOutput, error) {
	m.addCall("GetSSHPublicKeyWithContext")
	m.verifyInput("GetSSHPublicKeyWithContext", param0)
	return m.GetSSHPublicKeyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) GetServerCertificate(param0 *iam.GetServerCertificateInput) (*iam.GetServerCertificateOutput, error) {
	m.addCall("GetServerCertificate")
	m.verifyInput("GetServerCertificate", param0)
	return m.GetServerCertificateFunc(param0)
}

func (m *iamMock) GetServerCertificateRequest(param0 *iam.GetServerCertificateInput) (*request.Request, *iam.GetServerCertificateOutput) {
	m.addCall("GetServerCertificateRequest")
	m.verifyInput("GetServerCertificateRequest", param0)
	return m.GetServerCertificateRequestFunc(param0)
}

func (m *iamMock) GetServerCertificateWithContext(param0 aws.Context, param1 *iam.GetServerCertificateInput, param2 ...request.Option) (*iam.GetServerCertificateOutput, error) {
	m.addCall("GetServerCertificateWithContext")
	m.verifyInput("GetServerCertificateWithContext", param0)
	return m.GetServerCertificateWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) GetServiceLinkedRoleDeletionStatus(param0 *iam.GetServiceLinkedRoleDeletionStatusInput) (*iam.GetServiceLinkedRoleDeletionStatusOutput, error) {
	m.addCall("GetServiceLinkedRoleDeletionStatus")
	m.verifyInput("GetServiceLinkedRoleDeletionStatus", param0)
	return m.GetServiceLinkedRoleDeletionStatusFunc(param0)
}

func (m *iamMock) GetServiceLinkedRoleDeletionStatusRequest(param0 *iam.GetServiceLinkedRoleDeletionStatusInput) (*request.Request, *iam.GetServiceLinkedRoleDeletionStatusOutput) {
	m.addCall("GetServiceLinkedRoleDeletionStatusRequest")
	m.verifyInput("GetServiceLinkedRoleDeletionStatusRequest", param0)
	return m.GetServiceLinkedRoleDeletionStatusRequestFunc(param0)
}

func (m *iamMock) GetServiceLinkedRoleDeletionStatusWithContext(param0 aws.Context, param1 *iam.GetServiceLinkedRoleDeletionStatusInput, param2 ...request.Option) (*iam.GetServiceLinkedRoleDeletionStatusOutput, error) {
	m.addCall("GetServiceLinkedRoleDeletionStatusWithContext")
	m.verifyInput("GetServiceLinkedRoleDeletionStatusWithContext", param0)
	return m.GetServiceLinkedRoleDeletionStatusWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) GetUser(param0 *iam.GetUserInput) (*iam.GetUserOutput, error) {
	m.addCall("GetUser")
	m.verifyInput("GetUser", param0)
	return m.GetUserFunc(param0)
}

func (m *iamMock) GetUserPolicy(param0 *iam.GetUserPolicyInput) (*iam.GetUserPolicyOutput, error) {
	m.addCall("GetUserPolicy")
	m.verifyInput("GetUserPolicy", param0)
	return m.GetUserPolicyFunc(param0)
}

func (m *iamMock) GetUserPolicyRequest(param0 *iam.GetUserPolicyInput) (*request.Request, *iam.GetUserPolicyOutput) {
	m.addCall("GetUserPolicyRequest")
	m.verifyInput("GetUserPolicyRequest", param0)
	return m.GetUserPolicyRequestFunc(param0)
}

func (m *iamMock) GetUserPolicyWithContext(param0 aws.Context, param1 *iam.GetUserPolicyInput, param2 ...request.Option) (*iam.GetUserPolicyOutput, error) {
	m.addCall("GetUserPolicyWithContext")
	m.verifyInput("GetUserPolicyWithContext", param0)
	return m.GetUserPolicyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) GetUserRequest(param0 *iam.GetUserInput) (*request.Request, *iam.GetUserOutput) {
	m.addCall("GetUserRequest")
	m.verifyInput("GetUserRequest", param0)
	return m.GetUserRequestFunc(param0)
}

func (m *iamMock) GetUserWithContext(param0 aws.Context, param1 *iam.GetUserInput, param2 ...request.Option) (*iam.GetUserOutput, error) {
	m.addCall("GetUserWithContext")
	m.verifyInput("GetUserWithContext", param0)
	return m.GetUserWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) ListAccessKeys(param0 *iam.ListAccessKeysInput) (*iam.ListAccessKeysOutput, error) {
	m.addCall("ListAccessKeys")
	m.verifyInput("ListAccessKeys", param0)
	return m.ListAccessKeysFunc(param0)
}

func (m *iamMock) ListAccessKeysRequest(param0 *iam.ListAccessKeysInput) (*request.Request, *iam.ListAccessKeysOutput) {
	m.addCall("ListAccessKeysRequest")
	m.verifyInput("ListAccessKeysRequest", param0)
	return m.ListAccessKeysRequestFunc(param0)
}

func (m *iamMock) ListAccessKeysWithContext(param0 aws.Context, param1 *iam.ListAccessKeysInput, param2 ...request.Option) (*iam.ListAccessKeysOutput, error) {
	m.addCall("ListAccessKeysWithContext")
	m.verifyInput("ListAccessKeysWithContext", param0)
	return m.ListAccessKeysWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) ListAccountAliases(param0 *iam.ListAccountAliasesInput) (*iam.ListAccountAliasesOutput, error) {
	m.addCall("ListAccountAliases")
	m.verifyInput("ListAccountAliases", param0)
	return m.ListAccountAliasesFunc(param0)
}

func (m *iamMock) ListAccountAliasesRequest(param0 *iam.ListAccountAliasesInput) (*request.Request, *iam.ListAccountAliasesOutput) {
	m.addCall("ListAccountAliasesRequest")
	m.verifyInput("ListAccountAliasesRequest", param0)
	return m.ListAccountAliasesRequestFunc(param0)
}

func (m *iamMock) ListAccountAliasesWithContext(param0 aws.Context, param1 *iam.ListAccountAliasesInput, param2 ...request.Option) (*iam.ListAccountAliasesOutput, error) {
	m.addCall("ListAccountAliasesWithContext")
	m.verifyInput("ListAccountAliasesWithContext", param0)
	return m.ListAccountAliasesWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) ListAttachedGroupPolicies(param0 *iam.ListAttachedGroupPoliciesInput) (*iam.ListAttachedGroupPoliciesOutput, error) {
	m.addCall("ListAttachedGroupPolicies")
	m.verifyInput("ListAttachedGroupPolicies", param0)
	return m.ListAttachedGroupPoliciesFunc(param0)
}

func (m *iamMock) ListAttachedGroupPoliciesRequest(param0 *iam.ListAttachedGroupPoliciesInput) (*request.Request, *iam.ListAttachedGroupPoliciesOutput) {
	m.addCall("ListAttachedGroupPoliciesRequest")
	m.verifyInput("ListAttachedGroupPoliciesRequest", param0)
	return m.ListAttachedGroupPoliciesRequestFunc(param0)
}

func (m *iamMock) ListAttachedGroupPoliciesWithContext(param0 aws.Context, param1 *iam.ListAttachedGroupPoliciesInput, param2 ...request.Option) (*iam.ListAttachedGroupPoliciesOutput, error) {
	m.addCall("ListAttachedGroupPoliciesWithContext")
	m.verifyInput("ListAttachedGroupPoliciesWithContext", param0)
	return m.ListAttachedGroupPoliciesWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) ListAttachedRolePolicies(param0 *iam.ListAttachedRolePoliciesInput) (*iam.ListAttachedRolePoliciesOutput, error) {
	m.addCall("ListAttachedRolePolicies")
	m.verifyInput("ListAttachedRolePolicies", param0)
	return m.ListAttachedRolePoliciesFunc(param0)
}

func (m *iamMock) ListAttachedRolePoliciesRequest(param0 *iam.ListAttachedRolePoliciesInput) (*request.Request, *iam.ListAttachedRolePoliciesOutput) {
	m.addCall("ListAttachedRolePoliciesRequest")
	m.verifyInput("ListAttachedRolePoliciesRequest", param0)
	return m.ListAttachedRolePoliciesRequestFunc(param0)
}

func (m *iamMock) ListAttachedRolePoliciesWithContext(param0 aws.Context, param1 *iam.ListAttachedRolePoliciesInput, param2 ...request.Option) (*iam.ListAttachedRolePoliciesOutput, error) {
	m.addCall("ListAttachedRolePoliciesWithContext")
	m.verifyInput("ListAttachedRolePoliciesWithContext", param0)
	return m.ListAttachedRolePoliciesWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) ListAttachedUserPolicies(param0 *iam.ListAttachedUserPoliciesInput) (*iam.ListAttachedUserPoliciesOutput, error) {
	m.addCall("ListAttachedUserPolicies")
	m.verifyInput("ListAttachedUserPolicies", param0)
	return m.ListAttachedUserPoliciesFunc(param0)
}

func (m *iamMock) ListAttachedUserPoliciesRequest(param0 *iam.ListAttachedUserPoliciesInput) (*request.Request, *iam.ListAttachedUserPoliciesOutput) {
	m.addCall("ListAttachedUserPoliciesRequest")
	m.verifyInput("ListAttachedUserPoliciesRequest", param0)
	return m.ListAttachedUserPoliciesRequestFunc(param0)
}

func (m *iamMock) ListAttachedUserPoliciesWithContext(param0 aws.Context, param1 *iam.ListAttachedUserPoliciesInput, param2 ...request.Option) (*iam.ListAttachedUserPoliciesOutput, error) {
	m.addCall("ListAttachedUserPoliciesWithContext")
	m.verifyInput("ListAttachedUserPoliciesWithContext", param0)
	return m.ListAttachedUserPoliciesWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) ListEntitiesForPolicy(param0 *iam.ListEntitiesForPolicyInput) (*iam.ListEntitiesForPolicyOutput, error) {
	m.addCall("ListEntitiesForPolicy")
	m.verifyInput("ListEntitiesForPolicy", param0)
	return m.ListEntitiesForPolicyFunc(param0)
}

func (m *iamMock) ListEntitiesForPolicyRequest(param0 *iam.ListEntitiesForPolicyInput) (*request.Request, *iam.ListEntitiesForPolicyOutput) {
	m.addCall("ListEntitiesForPolicyRequest")
	m.verifyInput("ListEntitiesForPolicyRequest", param0)
	return m.ListEntitiesForPolicyRequestFunc(param0)
}

func (m *iamMock) ListEntitiesForPolicyWithContext(param0 aws.Context, param1 *iam.ListEntitiesForPolicyInput, param2 ...request.Option) (*iam.ListEntitiesForPolicyOutput, error) {
	m.addCall("ListEntitiesForPolicyWithContext")
	m.verifyInput("ListEntitiesForPolicyWithContext", param0)
	return m.ListEntitiesForPolicyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) ListGroupPolicies(param0 *iam.ListGroupPoliciesInput) (*iam.ListGroupPoliciesOutput, error) {
	m.addCall("ListGroupPolicies")
	m.verifyInput("ListGroupPolicies", param0)
	return m.ListGroupPoliciesFunc(param0)
}

func (m *iamMock) ListGroupPoliciesRequest(param0 *iam.ListGroupPoliciesInput) (*request.Request, *iam.ListGroupPoliciesOutput) {
	m.addCall("ListGroupPoliciesRequest")
	m.verifyInput("ListGroupPoliciesRequest", param0)
	return m.ListGroupPoliciesRequestFunc(param0)
}

func (m *iamMock) ListGroupPoliciesWithContext(param0 aws.Context, param1 *iam.ListGroupPoliciesInput, param2 ...request.Option) (*iam.ListGroupPoliciesOutput, error) {
	m.addCall("ListGroupPoliciesWithContext")
	m.verifyInput("ListGroupPoliciesWithContext", param0)
	return m.ListGroupPoliciesWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) ListGroups(param0 *iam.ListGroupsInput) (*iam.ListGroupsOutput, error) {
	m.addCall("ListGroups")
	m.verifyInput("ListGroups", param0)
	return m.ListGroupsFunc(param0)
}

func (m *iamMock) ListGroupsForUser(param0 *iam.ListGroupsForUserInput) (*iam.ListGroupsForUserOutput, error) {
	m.addCall("ListGroupsForUser")
	m.verifyInput("ListGroupsForUser", param0)
	return m.ListGroupsForUserFunc(param0)
}

func (m *iamMock) ListGroupsForUserRequest(param0 *iam.ListGroupsForUserInput) (*request.Request, *iam.ListGroupsForUserOutput) {
	m.addCall("ListGroupsForUserRequest")
	m.verifyInput("ListGroupsForUserRequest", param0)
	return m.ListGroupsForUserRequestFunc(param0)
}

func (m *iamMock) ListGroupsForUserWithContext(param0 aws.Context, param1 *iam.ListGroupsForUserInput, param2 ...request.Option) (*iam.ListGroupsForUserOutput, error) {
	m.addCall("ListGroupsForUserWithContext")
	m.verifyInput("ListGroupsForUserWithContext", param0)
	return m.ListGroupsForUserWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) ListGroupsRequest(param0 *iam.ListGroupsInput) (*request.Request, *iam.ListGroupsOutput) {
	m.addCall("ListGroupsRequest")
	m.verifyInput("ListGroupsRequest", param0)
	return m.ListGroupsRequestFunc(param0)
}

func (m *iamMock) ListGroupsWithContext(param0 aws.Context, param1 *iam.ListGroupsInput, param2 ...request.Option) (*iam.ListGroupsOutput, error) {
	m.addCall("ListGroupsWithContext")
	m.verifyInput("ListGroupsWithContext", param0)
	return m.ListGroupsWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) ListInstanceProfiles(param0 *iam.ListInstanceProfilesInput) (*iam.ListInstanceProfilesOutput, error) {
	m.addCall("ListInstanceProfiles")
	m.verifyInput("ListInstanceProfiles", param0)
	return m.ListInstanceProfilesFunc(param0)
}

func (m *iamMock) ListInstanceProfilesForRole(param0 *iam.ListInstanceProfilesForRoleInput) (*iam.ListInstanceProfilesForRoleOutput, error) {
	m.addCall("ListInstanceProfilesForRole")
	m.verifyInput("ListInstanceProfilesForRole", param0)
	return m.ListInstanceProfilesForRoleFunc(param0)
}

func (m *iamMock) ListInstanceProfilesForRoleRequest(param0 *iam.ListInstanceProfilesForRoleInput) (*request.Request, *iam.ListInstanceProfilesForRoleOutput) {
	m.addCall("ListInstanceProfilesForRoleRequest")
	m.verifyInput("ListInstanceProfilesForRoleRequest", param0)
	return m.ListInstanceProfilesForRoleRequestFunc(param0)
}

func (m *iamMock) ListInstanceProfilesForRoleWithContext(param0 aws.Context, param1 *iam.ListInstanceProfilesForRoleInput, param2 ...request.Option) (*iam.ListInstanceProfilesForRoleOutput, error) {
	m.addCall("ListInstanceProfilesForRoleWithContext")
	m.verifyInput("ListInstanceProfilesForRoleWithContext", param0)
	return m.ListInstanceProfilesForRoleWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) ListInstanceProfilesRequest(param0 *iam.ListInstanceProfilesInput) (*request.Request, *iam.ListInstanceProfilesOutput) {
	m.addCall("ListInstanceProfilesRequest")
	m.verifyInput("ListInstanceProfilesRequest", param0)
	return m.ListInstanceProfilesRequestFunc(param0)
}

func (m *iamMock) ListInstanceProfilesWithContext(param0 aws.Context, param1 *iam.ListInstanceProfilesInput, param2 ...request.Option) (*iam.ListInstanceProfilesOutput, error) {
	m.addCall("ListInstanceProfilesWithContext")
	m.verifyInput("ListInstanceProfilesWithContext", param0)
	return m.ListInstanceProfilesWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) ListMFADevices(param0 *iam.ListMFADevicesInput) (*iam.ListMFADevicesOutput, error) {
	m.addCall("ListMFADevices")
	m.verifyInput("ListMFADevices", param0)
	return m.ListMFADevicesFunc(param0)
}

func (m *iamMock) ListMFADevicesRequest(param0 *iam.ListMFADevicesInput) (*request.Request, *iam.ListMFADevicesOutput) {
	m.addCall("ListMFADevicesRequest")
	m.verifyInput("ListMFADevicesRequest", param0)
	return m.ListMFADevicesRequestFunc(param0)
}

func (m *iamMock) ListMFADevicesWithContext(param0 aws.Context, param1 *iam.ListMFADevicesInput, param2 ...request.Option) (*iam.ListMFADevicesOutput, error) {
	m.addCall("ListMFADevicesWithContext")
	m.verifyInput("ListMFADevicesWithContext", param0)
	return m.ListMFADevicesWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) ListOpenIDConnectProviders(param0 *iam.ListOpenIDConnectProvidersInput) (*iam.ListOpenIDConnectProvidersOutput, error) {
	m.addCall("ListOpenIDConnectProviders")
	m.verifyInput("ListOpenIDConnectProviders", param0)
	return m.ListOpenIDConnectProvidersFunc(param0)
}

func (m *iamMock) ListOpenIDConnectProvidersRequest(param0 *iam.ListOpenIDConnectProvidersInput) (*request.Request, *iam.ListOpenIDConnectProvidersOutput) {
	m.addCall("ListOpenIDConnectProvidersRequest")
	m.verifyInput("ListOpenIDConnectProvidersRequest", param0)
	return m.ListOpenIDConnectProvidersRequestFunc(param0)
}

func (m *iamMock) ListOpenIDConnectProvidersWithContext(param0 aws.Context, param1 *iam.ListOpenIDConnectProvidersInput, param2 ...request.Option) (*iam.ListOpenIDConnectProvidersOutput, error) {
	m.addCall("ListOpenIDConnectProvidersWithContext")
	m.verifyInput("ListOpenIDConnectProvidersWithContext", param0)
	return m.ListOpenIDConnectProvidersWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) ListPolicies(param0 *iam.ListPoliciesInput) (*iam.ListPoliciesOutput, error) {
	m.addCall("ListPolicies")
	m.verifyInput("ListPolicies", param0)
	return m.ListPoliciesFunc(param0)
}

func (m *iamMock) ListPoliciesRequest(param0 *iam.ListPoliciesInput) (*request.Request, *iam.ListPoliciesOutput) {
	m.addCall("ListPoliciesRequest")
	m.verifyInput("ListPoliciesRequest", param0)
	return m.ListPoliciesRequestFunc(param0)
}

func (m *iamMock) ListPoliciesWithContext(param0 aws.Context, param1 *iam.ListPoliciesInput, param2 ...request.Option) (*iam.ListPoliciesOutput, error) {
	m.addCall("ListPoliciesWithContext")
	m.verifyInput("ListPoliciesWithContext", param0)
	return m.ListPoliciesWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) ListPolicyVersions(param0 *iam.ListPolicyVersionsInput) (*iam.ListPolicyVersionsOutput, error) {
	m.addCall("ListPolicyVersions")
	m.verifyInput("ListPolicyVersions", param0)
	return m.ListPolicyVersionsFunc(param0)
}

func (m *iamMock) ListPolicyVersionsRequest(param0 *iam.ListPolicyVersionsInput) (*request.Request, *iam.ListPolicyVersionsOutput) {
	m.addCall("ListPolicyVersionsRequest")
	m.verifyInput("ListPolicyVersionsRequest", param0)
	return m.ListPolicyVersionsRequestFunc(param0)
}

func (m *iamMock) ListPolicyVersionsWithContext(param0 aws.Context, param1 *iam.ListPolicyVersionsInput, param2 ...request.Option) (*iam.ListPolicyVersionsOutput, error) {
	m.addCall("ListPolicyVersionsWithContext")
	m.verifyInput("ListPolicyVersionsWithContext", param0)
	return m.ListPolicyVersionsWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) ListRolePolicies(param0 *iam.ListRolePoliciesInput) (*iam.ListRolePoliciesOutput, error) {
	m.addCall("ListRolePolicies")
	m.verifyInput("ListRolePolicies", param0)
	return m.ListRolePoliciesFunc(param0)
}

func (m *iamMock) ListRolePoliciesRequest(param0 *iam.ListRolePoliciesInput) (*request.Request, *iam.ListRolePoliciesOutput) {
	m.addCall("ListRolePoliciesRequest")
	m.verifyInput("ListRolePoliciesRequest", param0)
	return m.ListRolePoliciesRequestFunc(param0)
}

func (m *iamMock) ListRolePoliciesWithContext(param0 aws.Context, param1 *iam.ListRolePoliciesInput, param2 ...request.Option) (*iam.ListRolePoliciesOutput, error) {
	m.addCall("ListRolePoliciesWithContext")
	m.verifyInput("ListRolePoliciesWithContext", param0)
	return m.ListRolePoliciesWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) ListRoles(param0 *iam.ListRolesInput) (*iam.ListRolesOutput, error) {
	m.addCall("ListRoles")
	m.verifyInput("ListRoles", param0)
	return m.ListRolesFunc(param0)
}

func (m *iamMock) ListRolesRequest(param0 *iam.ListRolesInput) (*request.Request, *iam.ListRolesOutput) {
	m.addCall("ListRolesRequest")
	m.verifyInput("ListRolesRequest", param0)
	return m.ListRolesRequestFunc(param0)
}

func (m *iamMock) ListRolesWithContext(param0 aws.Context, param1 *iam.ListRolesInput, param2 ...request.Option) (*iam.ListRolesOutput, error) {
	m.addCall("ListRolesWithContext")
	m.verifyInput("ListRolesWithContext", param0)
	return m.ListRolesWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) ListSAMLProviders(param0 *iam.ListSAMLProvidersInput) (*iam.ListSAMLProvidersOutput, error) {
	m.addCall("ListSAMLProviders")
	m.verifyInput("ListSAMLProviders", param0)
	return m.ListSAMLProvidersFunc(param0)
}

func (m *iamMock) ListSAMLProvidersRequest(param0 *iam.ListSAMLProvidersInput) (*request.Request, *iam.ListSAMLProvidersOutput) {
	m.addCall("ListSAMLProvidersRequest")
	m.verifyInput("ListSAMLProvidersRequest", param0)
	return m.ListSAMLProvidersRequestFunc(param0)
}

func (m *iamMock) ListSAMLProvidersWithContext(param0 aws.Context, param1 *iam.ListSAMLProvidersInput, param2 ...request.Option) (*iam.ListSAMLProvidersOutput, error) {
	m.addCall("ListSAMLProvidersWithContext")
	m.verifyInput("ListSAMLProvidersWithContext", param0)
	return m.ListSAMLProvidersWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) ListSSHPublicKeys(param0 *iam.ListSSHPublicKeysInput) (*iam.ListSSHPublicKeysOutput, error) {
	m.addCall("ListSSHPublicKeys")
	m.verifyInput("ListSSHPublicKeys", param0)
	return m.ListSSHPublicKeysFunc(param0)
}

func (m *iamMock) ListSSHPublicKeysRequest(param0 *iam.ListSSHPublicKeysInput) (*request.Request, *iam.ListSSHPublicKeysOutput) {
	m.addCall("ListSSHPublicKeysRequest")
	m.verifyInput("ListSSHPublicKeysRequest", param0)
	return m.ListSSHPublicKeysRequestFunc(param0)
}

func (m *iamMock) ListSSHPublicKeysWithContext(param0 aws.Context, param1 *iam.ListSSHPublicKeysInput, param2 ...request.Option) (*iam.ListSSHPublicKeysOutput, error) {
	m.addCall("ListSSHPublicKeysWithContext")
	m.verifyInput("ListSSHPublicKeysWithContext", param0)
	return m.ListSSHPublicKeysWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) ListServerCertificates(param0 *iam.ListServerCertificatesInput) (*iam.ListServerCertificatesOutput, error) {
	m.addCall("ListServerCertificates")
	m.verifyInput("ListServerCertificates", param0)
	return m.ListServerCertificatesFunc(param0)
}

func (m *iamMock) ListServerCertificatesRequest(param0 *iam.ListServerCertificatesInput) (*request.Request, *iam.ListServerCertificatesOutput) {
	m.addCall("ListServerCertificatesRequest")
	m.verifyInput("ListServerCertificatesRequest", param0)
	return m.ListServerCertificatesRequestFunc(param0)
}

func (m *iamMock) ListServerCertificatesWithContext(param0 aws.Context, param1 *iam.ListServerCertificatesInput, param2 ...request.Option) (*iam.ListServerCertificatesOutput, error) {
	m.addCall("ListServerCertificatesWithContext")
	m.verifyInput("ListServerCertificatesWithContext", param0)
	return m.ListServerCertificatesWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) ListServiceSpecificCredentials(param0 *iam.ListServiceSpecificCredentialsInput) (*iam.ListServiceSpecificCredentialsOutput, error) {
	m.addCall("ListServiceSpecificCredentials")
	m.verifyInput("ListServiceSpecificCredentials", param0)
	return m.ListServiceSpecificCredentialsFunc(param0)
}

func (m *iamMock) ListServiceSpecificCredentialsRequest(param0 *iam.ListServiceSpecificCredentialsInput) (*request.Request, *iam.ListServiceSpecificCredentialsOutput) {
	m.addCall("ListServiceSpecificCredentialsRequest")
	m.verifyInput("ListServiceSpecificCredentialsRequest", param0)
	return m.ListServiceSpecificCredentialsRequestFunc(param0)
}

func (m *iamMock) ListServiceSpecificCredentialsWithContext(param0 aws.Context, param1 *iam.ListServiceSpecificCredentialsInput, param2 ...request.Option) (*iam.ListServiceSpecificCredentialsOutput, error) {
	m.addCall("ListServiceSpecificCredentialsWithContext")
	m.verifyInput("ListServiceSpecificCredentialsWithContext", param0)
	return m.ListServiceSpecificCredentialsWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) ListSigningCertificates(param0 *iam.ListSigningCertificatesInput) (*iam.ListSigningCertificatesOutput, error) {
	m.addCall("ListSigningCertificates")
	m.verifyInput("ListSigningCertificates", param0)
	return m.ListSigningCertificatesFunc(param0)
}

func (m *iamMock) ListSigningCertificatesRequest(param0 *iam.ListSigningCertificatesInput) (*request.Request, *iam.ListSigningCertificatesOutput) {
	m.addCall("ListSigningCertificatesRequest")
	m.verifyInput("ListSigningCertificatesRequest", param0)
	return m.ListSigningCertificatesRequestFunc(param0)
}

func (m *iamMock) ListSigningCertificatesWithContext(param0 aws.Context, param1 *iam.ListSigningCertificatesInput, param2 ...request.Option) (*iam.ListSigningCertificatesOutput, error) {
	m.addCall("ListSigningCertificatesWithContext")
	m.verifyInput("ListSigningCertificatesWithContext", param0)
	return m.ListSigningCertificatesWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) ListUserPolicies(param0 *iam.ListUserPoliciesInput) (*iam.ListUserPoliciesOutput, error) {
	m.addCall("ListUserPolicies")
	m.verifyInput("ListUserPolicies", param0)
	return m.ListUserPoliciesFunc(param0)
}

func (m *iamMock) ListUserPoliciesRequest(param0 *iam.ListUserPoliciesInput) (*request.Request, *iam.ListUserPoliciesOutput) {
	m.addCall("ListUserPoliciesRequest")
	m.verifyInput("ListUserPoliciesRequest", param0)
	return m.ListUserPoliciesRequestFunc(param0)
}

func (m *iamMock) ListUserPoliciesWithContext(param0 aws.Context, param1 *iam.ListUserPoliciesInput, param2 ...request.Option) (*iam.ListUserPoliciesOutput, error) {
	m.addCall("ListUserPoliciesWithContext")
	m.verifyInput("ListUserPoliciesWithContext", param0)
	return m.ListUserPoliciesWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) ListUsers(param0 *iam.ListUsersInput) (*iam.ListUsersOutput, error) {
	m.addCall("ListUsers")
	m.verifyInput("ListUsers", param0)
	return m.ListUsersFunc(param0)
}

func (m *iamMock) ListUsersRequest(param0 *iam.ListUsersInput) (*request.Request, *iam.ListUsersOutput) {
	m.addCall("ListUsersRequest")
	m.verifyInput("ListUsersRequest", param0)
	return m.ListUsersRequestFunc(param0)
}

func (m *iamMock) ListUsersWithContext(param0 aws.Context, param1 *iam.ListUsersInput, param2 ...request.Option) (*iam.ListUsersOutput, error) {
	m.addCall("ListUsersWithContext")
	m.verifyInput("ListUsersWithContext", param0)
	return m.ListUsersWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) ListVirtualMFADevices(param0 *iam.ListVirtualMFADevicesInput) (*iam.ListVirtualMFADevicesOutput, error) {
	m.addCall("ListVirtualMFADevices")
	m.verifyInput("ListVirtualMFADevices", param0)
	return m.ListVirtualMFADevicesFunc(param0)
}

func (m *iamMock) ListVirtualMFADevicesRequest(param0 *iam.ListVirtualMFADevicesInput) (*request.Request, *iam.ListVirtualMFADevicesOutput) {
	m.addCall("ListVirtualMFADevicesRequest")
	m.verifyInput("ListVirtualMFADevicesRequest", param0)
	return m.ListVirtualMFADevicesRequestFunc(param0)
}

func (m *iamMock) ListVirtualMFADevicesWithContext(param0 aws.Context, param1 *iam.ListVirtualMFADevicesInput, param2 ...request.Option) (*iam.ListVirtualMFADevicesOutput, error) {
	m.addCall("ListVirtualMFADevicesWithContext")
	m.verifyInput("ListVirtualMFADevicesWithContext", param0)
	return m.ListVirtualMFADevicesWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) PutGroupPolicy(param0 *iam.PutGroupPolicyInput) (*iam.PutGroupPolicyOutput, error) {
	m.addCall("PutGroupPolicy")
	m.verifyInput("PutGroupPolicy", param0)
	return m.PutGroupPolicyFunc(param0)
}

func (m *iamMock) PutGroupPolicyRequest(param0 *iam.PutGroupPolicyInput) (*request.Request, *iam.PutGroupPolicyOutput) {
	m.addCall("PutGroupPolicyRequest")
	m.verifyInput("PutGroupPolicyRequest", param0)
	return m.PutGroupPolicyRequestFunc(param0)
}

func (m *iamMock) PutGroupPolicyWithContext(param0 aws.Context, param1 *iam.PutGroupPolicyInput, param2 ...request.Option) (*iam.PutGroupPolicyOutput, error) {
	m.addCall("PutGroupPolicyWithContext")
	m.verifyInput("PutGroupPolicyWithContext", param0)
	return m.PutGroupPolicyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) PutRolePolicy(param0 *iam.PutRolePolicyInput) (*iam.PutRolePolicyOutput, error) {
	m.addCall("PutRolePolicy")
	m.verifyInput("PutRolePolicy", param0)
	return m.PutRolePolicyFunc(param0)
}

func (m *iamMock) PutRolePolicyRequest(param0 *iam.PutRolePolicyInput) (*request.Request, *iam.PutRolePolicyOutput) {
	m.addCall("PutRolePolicyRequest")
	m.verifyInput("PutRolePolicyRequest", param0)
	return m.PutRolePolicyRequestFunc(param0)
}

func (m *iamMock) PutRolePolicyWithContext(param0 aws.Context, param1 *iam.PutRolePolicyInput, param2 ...request.Option) (*iam.PutRolePolicyOutput, error) {
	m.addCall("PutRolePolicyWithContext")
	m.verifyInput("PutRolePolicyWithContext", param0)
	return m.PutRolePolicyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) PutUserPolicy(param0 *iam.PutUserPolicyInput) (*iam.PutUserPolicyOutput, error) {
	m.addCall("PutUserPolicy")
	m.verifyInput("PutUserPolicy", param0)
	return m.PutUserPolicyFunc(param0)
}

func (m *iamMock) PutUserPolicyRequest(param0 *iam.PutUserPolicyInput) (*request.Request, *iam.PutUserPolicyOutput) {
	m.addCall("PutUserPolicyRequest")
	m.verifyInput("PutUserPolicyRequest", param0)
	return m.PutUserPolicyRequestFunc(param0)
}

func (m *iamMock) PutUserPolicyWithContext(param0 aws.Context, param1 *iam.PutUserPolicyInput, param2 ...request.Option) (*iam.PutUserPolicyOutput, error) {
	m.addCall("PutUserPolicyWithContext")
	m.verifyInput("PutUserPolicyWithContext", param0)
	return m.PutUserPolicyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) RemoveClientIDFromOpenIDConnectProvider(param0 *iam.RemoveClientIDFromOpenIDConnectProviderInput) (*iam.RemoveClientIDFromOpenIDConnectProviderOutput, error) {
	m.addCall("RemoveClientIDFromOpenIDConnectProvider")
	m.verifyInput("RemoveClientIDFromOpenIDConnectProvider", param0)
	return m.RemoveClientIDFromOpenIDConnectProviderFunc(param0)
}

func (m *iamMock) RemoveClientIDFromOpenIDConnectProviderRequest(param0 *iam.RemoveClientIDFromOpenIDConnectProviderInput) (*request.Request, *iam.RemoveClientIDFromOpenIDConnectProviderOutput) {
	m.addCall("RemoveClientIDFromOpenIDConnectProviderRequest")
	m.verifyInput("RemoveClientIDFromOpenIDConnectProviderRequest", param0)
	return m.RemoveClientIDFromOpenIDConnectProviderRequestFunc(param0)
}

func (m *iamMock) RemoveClientIDFromOpenIDConnectProviderWithContext(param0 aws.Context, param1 *iam.RemoveClientIDFromOpenIDConnectProviderInput, param2 ...request.Option) (*iam.RemoveClientIDFromOpenIDConnectProviderOutput, error) {
	m.addCall("RemoveClientIDFromOpenIDConnectProviderWithContext")
	m.verifyInput("RemoveClientIDFromOpenIDConnectProviderWithContext", param0)
	return m.RemoveClientIDFromOpenIDConnectProviderWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) RemoveRoleFromInstanceProfile(param0 *iam.RemoveRoleFromInstanceProfileInput) (*iam.RemoveRoleFromInstanceProfileOutput, error) {
	m.addCall("RemoveRoleFromInstanceProfile")
	m.verifyInput("RemoveRoleFromInstanceProfile", param0)
	return m.RemoveRoleFromInstanceProfileFunc(param0)
}

func (m *iamMock) RemoveRoleFromInstanceProfileRequest(param0 *iam.RemoveRoleFromInstanceProfileInput) (*request.Request, *iam.RemoveRoleFromInstanceProfileOutput) {
	m.addCall("RemoveRoleFromInstanceProfileRequest")
	m.verifyInput("RemoveRoleFromInstanceProfileRequest", param0)
	return m.RemoveRoleFromInstanceProfileRequestFunc(param0)
}

func (m *iamMock) RemoveRoleFromInstanceProfileWithContext(param0 aws.Context, param1 *iam.RemoveRoleFromInstanceProfileInput, param2 ...request.Option) (*iam.RemoveRoleFromInstanceProfileOutput, error) {
	m.addCall("RemoveRoleFromInstanceProfileWithContext")
	m.verifyInput("RemoveRoleFromInstanceProfileWithContext", param0)
	return m.RemoveRoleFromInstanceProfileWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) RemoveUserFromGroup(param0 *iam.RemoveUserFromGroupInput) (*iam.RemoveUserFromGroupOutput, error) {
	m.addCall("RemoveUserFromGroup")
	m.verifyInput("RemoveUserFromGroup", param0)
	return m.RemoveUserFromGroupFunc(param0)
}

func (m *iamMock) RemoveUserFromGroupRequest(param0 *iam.RemoveUserFromGroupInput) (*request.Request, *iam.RemoveUserFromGroupOutput) {
	m.addCall("RemoveUserFromGroupRequest")
	m.verifyInput("RemoveUserFromGroupRequest", param0)
	return m.RemoveUserFromGroupRequestFunc(param0)
}

func (m *iamMock) RemoveUserFromGroupWithContext(param0 aws.Context, param1 *iam.RemoveUserFromGroupInput, param2 ...request.Option) (*iam.RemoveUserFromGroupOutput, error) {
	m.addCall("RemoveUserFromGroupWithContext")
	m.verifyInput("RemoveUserFromGroupWithContext", param0)
	return m.RemoveUserFromGroupWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) ResetServiceSpecificCredential(param0 *iam.ResetServiceSpecificCredentialInput) (*iam.ResetServiceSpecificCredentialOutput, error) {
	m.addCall("ResetServiceSpecificCredential")
	m.verifyInput("ResetServiceSpecificCredential", param0)
	return m.ResetServiceSpecificCredentialFunc(param0)
}

func (m *iamMock) ResetServiceSpecificCredentialRequest(param0 *iam.ResetServiceSpecificCredentialInput) (*request.Request, *iam.ResetServiceSpecificCredentialOutput) {
	m.addCall("ResetServiceSpecificCredentialRequest")
	m.verifyInput("ResetServiceSpecificCredentialRequest", param0)
	return m.ResetServiceSpecificCredentialRequestFunc(param0)
}

func (m *iamMock) ResetServiceSpecificCredentialWithContext(param0 aws.Context, param1 *iam.ResetServiceSpecificCredentialInput, param2 ...request.Option) (*iam.ResetServiceSpecificCredentialOutput, error) {
	m.addCall("ResetServiceSpecificCredentialWithContext")
	m.verifyInput("ResetServiceSpecificCredentialWithContext", param0)
	return m.ResetServiceSpecificCredentialWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) ResyncMFADevice(param0 *iam.ResyncMFADeviceInput) (*iam.ResyncMFADeviceOutput, error) {
	m.addCall("ResyncMFADevice")
	m.verifyInput("ResyncMFADevice", param0)
	return m.ResyncMFADeviceFunc(param0)
}

func (m *iamMock) ResyncMFADeviceRequest(param0 *iam.ResyncMFADeviceInput) (*request.Request, *iam.ResyncMFADeviceOutput) {
	m.addCall("ResyncMFADeviceRequest")
	m.verifyInput("ResyncMFADeviceRequest", param0)
	return m.ResyncMFADeviceRequestFunc(param0)
}

func (m *iamMock) ResyncMFADeviceWithContext(param0 aws.Context, param1 *iam.ResyncMFADeviceInput, param2 ...request.Option) (*iam.ResyncMFADeviceOutput, error) {
	m.addCall("ResyncMFADeviceWithContext")
	m.verifyInput("ResyncMFADeviceWithContext", param0)
	return m.ResyncMFADeviceWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) SetDefaultPolicyVersion(param0 *iam.SetDefaultPolicyVersionInput) (*iam.SetDefaultPolicyVersionOutput, error) {
	m.addCall("SetDefaultPolicyVersion")
	m.verifyInput("SetDefaultPolicyVersion", param0)
	return m.SetDefaultPolicyVersionFunc(param0)
}

func (m *iamMock) SetDefaultPolicyVersionRequest(param0 *iam.SetDefaultPolicyVersionInput) (*request.Request, *iam.SetDefaultPolicyVersionOutput) {
	m.addCall("SetDefaultPolicyVersionRequest")
	m.verifyInput("SetDefaultPolicyVersionRequest", param0)
	return m.SetDefaultPolicyVersionRequestFunc(param0)
}

func (m *iamMock) SetDefaultPolicyVersionWithContext(param0 aws.Context, param1 *iam.SetDefaultPolicyVersionInput, param2 ...request.Option) (*iam.SetDefaultPolicyVersionOutput, error) {
	m.addCall("SetDefaultPolicyVersionWithContext")
	m.verifyInput("SetDefaultPolicyVersionWithContext", param0)
	return m.SetDefaultPolicyVersionWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) SimulateCustomPolicy(param0 *iam.SimulateCustomPolicyInput) (*iam.SimulatePolicyResponse, error) {
	m.addCall("SimulateCustomPolicy")
	m.verifyInput("SimulateCustomPolicy", param0)
	return m.SimulateCustomPolicyFunc(param0)
}

func (m *iamMock) SimulateCustomPolicyRequest(param0 *iam.SimulateCustomPolicyInput) (*request.Request, *iam.SimulatePolicyResponse) {
	m.addCall("SimulateCustomPolicyRequest")
	m.verifyInput("SimulateCustomPolicyRequest", param0)
	return m.SimulateCustomPolicyRequestFunc(param0)
}

func (m *iamMock) SimulateCustomPolicyWithContext(param0 aws.Context, param1 *iam.SimulateCustomPolicyInput, param2 ...request.Option) (*iam.SimulatePolicyResponse, error) {
	m.addCall("SimulateCustomPolicyWithContext")
	m.verifyInput("SimulateCustomPolicyWithContext", param0)
	return m.SimulateCustomPolicyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) SimulatePrincipalPolicy(param0 *iam.SimulatePrincipalPolicyInput) (*iam.SimulatePolicyResponse, error) {
	m.addCall("SimulatePrincipalPolicy")
	m.verifyInput("SimulatePrincipalPolicy", param0)
	return m.SimulatePrincipalPolicyFunc(param0)
}

func (m *iamMock) SimulatePrincipalPolicyRequest(param0 *iam.SimulatePrincipalPolicyInput) (*request.Request, *iam.SimulatePolicyResponse) {
	m.addCall("SimulatePrincipalPolicyRequest")
	m.verifyInput("SimulatePrincipalPolicyRequest", param0)
	return m.SimulatePrincipalPolicyRequestFunc(param0)
}

func (m *iamMock) SimulatePrincipalPolicyWithContext(param0 aws.Context, param1 *iam.SimulatePrincipalPolicyInput, param2 ...request.Option) (*iam.SimulatePolicyResponse, error) {
	m.addCall("SimulatePrincipalPolicyWithContext")
	m.verifyInput("SimulatePrincipalPolicyWithContext", param0)
	return m.SimulatePrincipalPolicyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) UpdateAccessKey(param0 *iam.UpdateAccessKeyInput) (*iam.UpdateAccessKeyOutput, error) {
	m.addCall("UpdateAccessKey")
	m.verifyInput("UpdateAccessKey", param0)
	return m.UpdateAccessKeyFunc(param0)
}

func (m *iamMock) UpdateAccessKeyRequest(param0 *iam.UpdateAccessKeyInput) (*request.Request, *iam.UpdateAccessKeyOutput) {
	m.addCall("UpdateAccessKeyRequest")
	m.verifyInput("UpdateAccessKeyRequest", param0)
	return m.UpdateAccessKeyRequestFunc(param0)
}

func (m *iamMock) UpdateAccessKeyWithContext(param0 aws.Context, param1 *iam.UpdateAccessKeyInput, param2 ...request.Option) (*iam.UpdateAccessKeyOutput, error) {
	m.addCall("UpdateAccessKeyWithContext")
	m.verifyInput("UpdateAccessKeyWithContext", param0)
	return m.UpdateAccessKeyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) UpdateAccountPasswordPolicy(param0 *iam.UpdateAccountPasswordPolicyInput) (*iam.UpdateAccountPasswordPolicyOutput, error) {
	m.addCall("UpdateAccountPasswordPolicy")
	m.verifyInput("UpdateAccountPasswordPolicy", param0)
	return m.UpdateAccountPasswordPolicyFunc(param0)
}

func (m *iamMock) UpdateAccountPasswordPolicyRequest(param0 *iam.UpdateAccountPasswordPolicyInput) (*request.Request, *iam.UpdateAccountPasswordPolicyOutput) {
	m.addCall("UpdateAccountPasswordPolicyRequest")
	m.verifyInput("UpdateAccountPasswordPolicyRequest", param0)
	return m.UpdateAccountPasswordPolicyRequestFunc(param0)
}

func (m *iamMock) UpdateAccountPasswordPolicyWithContext(param0 aws.Context, param1 *iam.UpdateAccountPasswordPolicyInput, param2 ...request.Option) (*iam.UpdateAccountPasswordPolicyOutput, error) {
	m.addCall("UpdateAccountPasswordPolicyWithContext")
	m.verifyInput("UpdateAccountPasswordPolicyWithContext", param0)
	return m.UpdateAccountPasswordPolicyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) UpdateAssumeRolePolicy(param0 *iam.UpdateAssumeRolePolicyInput) (*iam.UpdateAssumeRolePolicyOutput, error) {
	m.addCall("UpdateAssumeRolePolicy")
	m.verifyInput("UpdateAssumeRolePolicy", param0)
	return m.UpdateAssumeRolePolicyFunc(param0)
}

func (m *iamMock) UpdateAssumeRolePolicyRequest(param0 *iam.UpdateAssumeRolePolicyInput) (*request.Request, *iam.UpdateAssumeRolePolicyOutput) {
	m.addCall("UpdateAssumeRolePolicyRequest")
	m.verifyInput("UpdateAssumeRolePolicyRequest", param0)
	return m.UpdateAssumeRolePolicyRequestFunc(param0)
}

func (m *iamMock) UpdateAssumeRolePolicyWithContext(param0 aws.Context, param1 *iam.UpdateAssumeRolePolicyInput, param2 ...request.Option) (*iam.UpdateAssumeRolePolicyOutput, error) {
	m.addCall("UpdateAssumeRolePolicyWithContext")
	m.verifyInput("UpdateAssumeRolePolicyWithContext", param0)
	return m.UpdateAssumeRolePolicyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) UpdateGroup(param0 *iam.UpdateGroupInput) (*iam.UpdateGroupOutput, error) {
	m.addCall("UpdateGroup")
	m.verifyInput("UpdateGroup", param0)
	return m.UpdateGroupFunc(param0)
}

func (m *iamMock) UpdateGroupRequest(param0 *iam.UpdateGroupInput) (*request.Request, *iam.UpdateGroupOutput) {
	m.addCall("UpdateGroupRequest")
	m.verifyInput("UpdateGroupRequest", param0)
	return m.UpdateGroupRequestFunc(param0)
}

func (m *iamMock) UpdateGroupWithContext(param0 aws.Context, param1 *iam.UpdateGroupInput, param2 ...request.Option) (*iam.UpdateGroupOutput, error) {
	m.addCall("UpdateGroupWithContext")
	m.verifyInput("UpdateGroupWithContext", param0)
	return m.UpdateGroupWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) UpdateLoginProfile(param0 *iam.UpdateLoginProfileInput) (*iam.UpdateLoginProfileOutput, error) {
	m.addCall("UpdateLoginProfile")
	m.verifyInput("UpdateLoginProfile", param0)
	return m.UpdateLoginProfileFunc(param0)
}

func (m *iamMock) UpdateLoginProfileRequest(param0 *iam.UpdateLoginProfileInput) (*request.Request, *iam.UpdateLoginProfileOutput) {
	m.addCall("UpdateLoginProfileRequest")
	m.verifyInput("UpdateLoginProfileRequest", param0)
	return m.UpdateLoginProfileRequestFunc(param0)
}

func (m *iamMock) UpdateLoginProfileWithContext(param0 aws.Context, param1 *iam.UpdateLoginProfileInput, param2 ...request.Option) (*iam.UpdateLoginProfileOutput, error) {
	m.addCall("UpdateLoginProfileWithContext")
	m.verifyInput("UpdateLoginProfileWithContext", param0)
	return m.UpdateLoginProfileWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) UpdateOpenIDConnectProviderThumbprint(param0 *iam.UpdateOpenIDConnectProviderThumbprintInput) (*iam.UpdateOpenIDConnectProviderThumbprintOutput, error) {
	m.addCall("UpdateOpenIDConnectProviderThumbprint")
	m.verifyInput("UpdateOpenIDConnectProviderThumbprint", param0)
	return m.UpdateOpenIDConnectProviderThumbprintFunc(param0)
}

func (m *iamMock) UpdateOpenIDConnectProviderThumbprintRequest(param0 *iam.UpdateOpenIDConnectProviderThumbprintInput) (*request.Request, *iam.UpdateOpenIDConnectProviderThumbprintOutput) {
	m.addCall("UpdateOpenIDConnectProviderThumbprintRequest")
	m.verifyInput("UpdateOpenIDConnectProviderThumbprintRequest", param0)
	return m.UpdateOpenIDConnectProviderThumbprintRequestFunc(param0)
}

func (m *iamMock) UpdateOpenIDConnectProviderThumbprintWithContext(param0 aws.Context, param1 *iam.UpdateOpenIDConnectProviderThumbprintInput, param2 ...request.Option) (*iam.UpdateOpenIDConnectProviderThumbprintOutput, error) {
	m.addCall("UpdateOpenIDConnectProviderThumbprintWithContext")
	m.verifyInput("UpdateOpenIDConnectProviderThumbprintWithContext", param0)
	return m.UpdateOpenIDConnectProviderThumbprintWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) UpdateRoleDescription(param0 *iam.UpdateRoleDescriptionInput) (*iam.UpdateRoleDescriptionOutput, error) {
	m.addCall("UpdateRoleDescription")
	m.verifyInput("UpdateRoleDescription", param0)
	return m.UpdateRoleDescriptionFunc(param0)
}

func (m *iamMock) UpdateRoleDescriptionRequest(param0 *iam.UpdateRoleDescriptionInput) (*request.Request, *iam.UpdateRoleDescriptionOutput) {
	m.addCall("UpdateRoleDescriptionRequest")
	m.verifyInput("UpdateRoleDescriptionRequest", param0)
	return m.UpdateRoleDescriptionRequestFunc(param0)
}

func (m *iamMock) UpdateRoleDescriptionWithContext(param0 aws.Context, param1 *iam.UpdateRoleDescriptionInput, param2 ...request.Option) (*iam.UpdateRoleDescriptionOutput, error) {
	m.addCall("UpdateRoleDescriptionWithContext")
	m.verifyInput("UpdateRoleDescriptionWithContext", param0)
	return m.UpdateRoleDescriptionWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) UpdateSAMLProvider(param0 *iam.UpdateSAMLProviderInput) (*iam.UpdateSAMLProviderOutput, error) {
	m.addCall("UpdateSAMLProvider")
	m.verifyInput("UpdateSAMLProvider", param0)
	return m.UpdateSAMLProviderFunc(param0)
}

func (m *iamMock) UpdateSAMLProviderRequest(param0 *iam.UpdateSAMLProviderInput) (*request.Request, *iam.UpdateSAMLProviderOutput) {
	m.addCall("UpdateSAMLProviderRequest")
	m.verifyInput("UpdateSAMLProviderRequest", param0)
	return m.UpdateSAMLProviderRequestFunc(param0)
}

func (m *iamMock) UpdateSAMLProviderWithContext(param0 aws.Context, param1 *iam.UpdateSAMLProviderInput, param2 ...request.Option) (*iam.UpdateSAMLProviderOutput, error) {
	m.addCall("UpdateSAMLProviderWithContext")
	m.verifyInput("UpdateSAMLProviderWithContext", param0)
	return m.UpdateSAMLProviderWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) UpdateSSHPublicKey(param0 *iam.UpdateSSHPublicKeyInput) (*iam.UpdateSSHPublicKeyOutput, error) {
	m.addCall("UpdateSSHPublicKey")
	m.verifyInput("UpdateSSHPublicKey", param0)
	return m.UpdateSSHPublicKeyFunc(param0)
}

func (m *iamMock) UpdateSSHPublicKeyRequest(param0 *iam.UpdateSSHPublicKeyInput) (*request.Request, *iam.UpdateSSHPublicKeyOutput) {
	m.addCall("UpdateSSHPublicKeyRequest")
	m.verifyInput("UpdateSSHPublicKeyRequest", param0)
	return m.UpdateSSHPublicKeyRequestFunc(param0)
}

func (m *iamMock) UpdateSSHPublicKeyWithContext(param0 aws.Context, param1 *iam.UpdateSSHPublicKeyInput, param2 ...request.Option) (*iam.UpdateSSHPublicKeyOutput, error) {
	m.addCall("UpdateSSHPublicKeyWithContext")
	m.verifyInput("UpdateSSHPublicKeyWithContext", param0)
	return m.UpdateSSHPublicKeyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) UpdateServerCertificate(param0 *iam.UpdateServerCertificateInput) (*iam.UpdateServerCertificateOutput, error) {
	m.addCall("UpdateServerCertificate")
	m.verifyInput("UpdateServerCertificate", param0)
	return m.UpdateServerCertificateFunc(param0)
}

func (m *iamMock) UpdateServerCertificateRequest(param0 *iam.UpdateServerCertificateInput) (*request.Request, *iam.UpdateServerCertificateOutput) {
	m.addCall("UpdateServerCertificateRequest")
	m.verifyInput("UpdateServerCertificateRequest", param0)
	return m.UpdateServerCertificateRequestFunc(param0)
}

func (m *iamMock) UpdateServerCertificateWithContext(param0 aws.Context, param1 *iam.UpdateServerCertificateInput, param2 ...request.Option) (*iam.UpdateServerCertificateOutput, error) {
	m.addCall("UpdateServerCertificateWithContext")
	m.verifyInput("UpdateServerCertificateWithContext", param0)
	return m.UpdateServerCertificateWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) UpdateServiceSpecificCredential(param0 *iam.UpdateServiceSpecificCredentialInput) (*iam.UpdateServiceSpecificCredentialOutput, error) {
	m.addCall("UpdateServiceSpecificCredential")
	m.verifyInput("UpdateServiceSpecificCredential", param0)
	return m.UpdateServiceSpecificCredentialFunc(param0)
}

func (m *iamMock) UpdateServiceSpecificCredentialRequest(param0 *iam.UpdateServiceSpecificCredentialInput) (*request.Request, *iam.UpdateServiceSpecificCredentialOutput) {
	m.addCall("UpdateServiceSpecificCredentialRequest")
	m.verifyInput("UpdateServiceSpecificCredentialRequest", param0)
	return m.UpdateServiceSpecificCredentialRequestFunc(param0)
}

func (m *iamMock) UpdateServiceSpecificCredentialWithContext(param0 aws.Context, param1 *iam.UpdateServiceSpecificCredentialInput, param2 ...request.Option) (*iam.UpdateServiceSpecificCredentialOutput, error) {
	m.addCall("UpdateServiceSpecificCredentialWithContext")
	m.verifyInput("UpdateServiceSpecificCredentialWithContext", param0)
	return m.UpdateServiceSpecificCredentialWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) UpdateSigningCertificate(param0 *iam.UpdateSigningCertificateInput) (*iam.UpdateSigningCertificateOutput, error) {
	m.addCall("UpdateSigningCertificate")
	m.verifyInput("UpdateSigningCertificate", param0)
	return m.UpdateSigningCertificateFunc(param0)
}

func (m *iamMock) UpdateSigningCertificateRequest(param0 *iam.UpdateSigningCertificateInput) (*request.Request, *iam.UpdateSigningCertificateOutput) {
	m.addCall("UpdateSigningCertificateRequest")
	m.verifyInput("UpdateSigningCertificateRequest", param0)
	return m.UpdateSigningCertificateRequestFunc(param0)
}

func (m *iamMock) UpdateSigningCertificateWithContext(param0 aws.Context, param1 *iam.UpdateSigningCertificateInput, param2 ...request.Option) (*iam.UpdateSigningCertificateOutput, error) {
	m.addCall("UpdateSigningCertificateWithContext")
	m.verifyInput("UpdateSigningCertificateWithContext", param0)
	return m.UpdateSigningCertificateWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) UpdateUser(param0 *iam.UpdateUserInput) (*iam.UpdateUserOutput, error) {
	m.addCall("UpdateUser")
	m.verifyInput("UpdateUser", param0)
	return m.UpdateUserFunc(param0)
}

func (m *iamMock) UpdateUserRequest(param0 *iam.UpdateUserInput) (*request.Request, *iam.UpdateUserOutput) {
	m.addCall("UpdateUserRequest")
	m.verifyInput("UpdateUserRequest", param0)
	return m.UpdateUserRequestFunc(param0)
}

func (m *iamMock) UpdateUserWithContext(param0 aws.Context, param1 *iam.UpdateUserInput, param2 ...request.Option) (*iam.UpdateUserOutput, error) {
	m.addCall("UpdateUserWithContext")
	m.verifyInput("UpdateUserWithContext", param0)
	return m.UpdateUserWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) UploadSSHPublicKey(param0 *iam.UploadSSHPublicKeyInput) (*iam.UploadSSHPublicKeyOutput, error) {
	m.addCall("UploadSSHPublicKey")
	m.verifyInput("UploadSSHPublicKey", param0)
	return m.UploadSSHPublicKeyFunc(param0)
}

func (m *iamMock) UploadSSHPublicKeyRequest(param0 *iam.UploadSSHPublicKeyInput) (*request.Request, *iam.UploadSSHPublicKeyOutput) {
	m.addCall("UploadSSHPublicKeyRequest")
	m.verifyInput("UploadSSHPublicKeyRequest", param0)
	return m.UploadSSHPublicKeyRequestFunc(param0)
}

func (m *iamMock) UploadSSHPublicKeyWithContext(param0 aws.Context, param1 *iam.UploadSSHPublicKeyInput, param2 ...request.Option) (*iam.UploadSSHPublicKeyOutput, error) {
	m.addCall("UploadSSHPublicKeyWithContext")
	m.verifyInput("UploadSSHPublicKeyWithContext", param0)
	return m.UploadSSHPublicKeyWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) UploadServerCertificate(param0 *iam.UploadServerCertificateInput) (*iam.UploadServerCertificateOutput, error) {
	m.addCall("UploadServerCertificate")
	m.verifyInput("UploadServerCertificate", param0)
	return m.UploadServerCertificateFunc(param0)
}

func (m *iamMock) UploadServerCertificateRequest(param0 *iam.UploadServerCertificateInput) (*request.Request, *iam.UploadServerCertificateOutput) {
	m.addCall("UploadServerCertificateRequest")
	m.verifyInput("UploadServerCertificateRequest", param0)
	return m.UploadServerCertificateRequestFunc(param0)
}

func (m *iamMock) UploadServerCertificateWithContext(param0 aws.Context, param1 *iam.UploadServerCertificateInput, param2 ...request.Option) (*iam.UploadServerCertificateOutput, error) {
	m.addCall("UploadServerCertificateWithContext")
	m.verifyInput("UploadServerCertificateWithContext", param0)
	return m.UploadServerCertificateWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) UploadSigningCertificate(param0 *iam.UploadSigningCertificateInput) (*iam.UploadSigningCertificateOutput, error) {
	m.addCall("UploadSigningCertificate")
	m.verifyInput("UploadSigningCertificate", param0)
	return m.UploadSigningCertificateFunc(param0)
}

func (m *iamMock) UploadSigningCertificateRequest(param0 *iam.UploadSigningCertificateInput) (*request.Request, *iam.UploadSigningCertificateOutput) {
	m.addCall("UploadSigningCertificateRequest")
	m.verifyInput("UploadSigningCertificateRequest", param0)
	return m.UploadSigningCertificateRequestFunc(param0)
}

func (m *iamMock) UploadSigningCertificateWithContext(param0 aws.Context, param1 *iam.UploadSigningCertificateInput, param2 ...request.Option) (*iam.UploadSigningCertificateOutput, error) {
	m.addCall("UploadSigningCertificateWithContext")
	m.verifyInput("UploadSigningCertificateWithContext", param0)
	return m.UploadSigningCertificateWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) WaitUntilInstanceProfileExists(param0 *iam.GetInstanceProfileInput) error {
	m.addCall("WaitUntilInstanceProfileExists")
	m.verifyInput("WaitUntilInstanceProfileExists", param0)
	return m.WaitUntilInstanceProfileExistsFunc(param0)
}

func (m *iamMock) WaitUntilInstanceProfileExistsWithContext(param0 aws.Context, param1 *iam.GetInstanceProfileInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilInstanceProfileExistsWithContext")
	m.verifyInput("WaitUntilInstanceProfileExistsWithContext", param0)
	return m.WaitUntilInstanceProfileExistsWithContextFunc(param0, param1, param2...)
}

func (m *iamMock) WaitUntilUserExists(param0 *iam.GetUserInput) error {
	m.addCall("WaitUntilUserExists")
	m.verifyInput("WaitUntilUserExists", param0)
	return m.WaitUntilUserExistsFunc(param0)
}

func (m *iamMock) WaitUntilUserExistsWithContext(param0 aws.Context, param1 *iam.GetUserInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilUserExistsWithContext")
	m.verifyInput("WaitUntilUserExistsWithContext", param0)
	return m.WaitUntilUserExistsWithContextFunc(param0, param1, param2...)
}

type route53Mock struct {
	basicMock
	route53iface.Route53API
	AssociateVPCWithHostedZoneFunc                        func(param0 *route53.AssociateVPCWithHostedZoneInput) (*route53.AssociateVPCWithHostedZoneOutput, error)
	AssociateVPCWithHostedZoneRequestFunc                 func(param0 *route53.AssociateVPCWithHostedZoneInput) (*request.Request, *route53.AssociateVPCWithHostedZoneOutput)
	AssociateVPCWithHostedZoneWithContextFunc             func(param0 aws.Context, param1 *route53.AssociateVPCWithHostedZoneInput, param2 ...request.Option) (*route53.AssociateVPCWithHostedZoneOutput, error)
	ChangeResourceRecordSetsFunc                          func(param0 *route53.ChangeResourceRecordSetsInput) (*route53.ChangeResourceRecordSetsOutput, error)
	ChangeResourceRecordSetsRequestFunc                   func(param0 *route53.ChangeResourceRecordSetsInput) (*request.Request, *route53.ChangeResourceRecordSetsOutput)
	ChangeResourceRecordSetsWithContextFunc               func(param0 aws.Context, param1 *route53.ChangeResourceRecordSetsInput, param2 ...request.Option) (*route53.ChangeResourceRecordSetsOutput, error)
	ChangeTagsForResourceFunc                             func(param0 *route53.ChangeTagsForResourceInput) (*route53.ChangeTagsForResourceOutput, error)
	ChangeTagsForResourceRequestFunc                      func(param0 *route53.ChangeTagsForResourceInput) (*request.Request, *route53.ChangeTagsForResourceOutput)
	ChangeTagsForResourceWithContextFunc                  func(param0 aws.Context, param1 *route53.ChangeTagsForResourceInput, param2 ...request.Option) (*route53.ChangeTagsForResourceOutput, error)
	CreateHealthCheckFunc                                 func(param0 *route53.CreateHealthCheckInput) (*route53.CreateHealthCheckOutput, error)
	CreateHealthCheckRequestFunc                          func(param0 *route53.CreateHealthCheckInput) (*request.Request, *route53.CreateHealthCheckOutput)
	CreateHealthCheckWithContextFunc                      func(param0 aws.Context, param1 *route53.CreateHealthCheckInput, param2 ...request.Option) (*route53.CreateHealthCheckOutput, error)
	CreateHostedZoneFunc                                  func(param0 *route53.CreateHostedZoneInput) (*route53.CreateHostedZoneOutput, error)
	CreateHostedZoneRequestFunc                           func(param0 *route53.CreateHostedZoneInput) (*request.Request, *route53.CreateHostedZoneOutput)
	CreateHostedZoneWithContextFunc                       func(param0 aws.Context, param1 *route53.CreateHostedZoneInput, param2 ...request.Option) (*route53.CreateHostedZoneOutput, error)
	CreateQueryLoggingConfigFunc                          func(param0 *route53.CreateQueryLoggingConfigInput) (*route53.CreateQueryLoggingConfigOutput, error)
	CreateQueryLoggingConfigRequestFunc                   func(param0 *route53.CreateQueryLoggingConfigInput) (*request.Request, *route53.CreateQueryLoggingConfigOutput)
	CreateQueryLoggingConfigWithContextFunc               func(param0 aws.Context, param1 *route53.CreateQueryLoggingConfigInput, param2 ...request.Option) (*route53.CreateQueryLoggingConfigOutput, error)
	CreateReusableDelegationSetFunc                       func(param0 *route53.CreateReusableDelegationSetInput) (*route53.CreateReusableDelegationSetOutput, error)
	CreateReusableDelegationSetRequestFunc                func(param0 *route53.CreateReusableDelegationSetInput) (*request.Request, *route53.CreateReusableDelegationSetOutput)
	CreateReusableDelegationSetWithContextFunc            func(param0 aws.Context, param1 *route53.CreateReusableDelegationSetInput, param2 ...request.Option) (*route53.CreateReusableDelegationSetOutput, error)
	CreateTrafficPolicyFunc                               func(param0 *route53.CreateTrafficPolicyInput) (*route53.CreateTrafficPolicyOutput, error)
	CreateTrafficPolicyInstanceFunc                       func(param0 *route53.CreateTrafficPolicyInstanceInput) (*route53.CreateTrafficPolicyInstanceOutput, error)
	CreateTrafficPolicyInstanceRequestFunc                func(param0 *route53.CreateTrafficPolicyInstanceInput) (*request.Request, *route53.CreateTrafficPolicyInstanceOutput)
	CreateTrafficPolicyInstanceWithContextFunc            func(param0 aws.Context, param1 *route53.CreateTrafficPolicyInstanceInput, param2 ...request.Option) (*route53.CreateTrafficPolicyInstanceOutput, error)
	CreateTrafficPolicyRequestFunc                        func(param0 *route53.CreateTrafficPolicyInput) (*request.Request, *route53.CreateTrafficPolicyOutput)
	CreateTrafficPolicyVersionFunc                        func(param0 *route53.CreateTrafficPolicyVersionInput) (*route53.CreateTrafficPolicyVersionOutput, error)
	CreateTrafficPolicyVersionRequestFunc                 func(param0 *route53.CreateTrafficPolicyVersionInput) (*request.Request, *route53.CreateTrafficPolicyVersionOutput)
	CreateTrafficPolicyVersionWithContextFunc             func(param0 aws.Context, param1 *route53.CreateTrafficPolicyVersionInput, param2 ...request.Option) (*route53.CreateTrafficPolicyVersionOutput, error)
	CreateTrafficPolicyWithContextFunc                    func(param0 aws.Context, param1 *route53.CreateTrafficPolicyInput, param2 ...request.Option) (*route53.CreateTrafficPolicyOutput, error)
	CreateVPCAssociationAuthorizationFunc                 func(param0 *route53.CreateVPCAssociationAuthorizationInput) (*route53.CreateVPCAssociationAuthorizationOutput, error)
	CreateVPCAssociationAuthorizationRequestFunc          func(param0 *route53.CreateVPCAssociationAuthorizationInput) (*request.Request, *route53.CreateVPCAssociationAuthorizationOutput)
	CreateVPCAssociationAuthorizationWithContextFunc      func(param0 aws.Context, param1 *route53.CreateVPCAssociationAuthorizationInput, param2 ...request.Option) (*route53.CreateVPCAssociationAuthorizationOutput, error)
	DeleteHealthCheckFunc                                 func(param0 *route53.DeleteHealthCheckInput) (*route53.DeleteHealthCheckOutput, error)
	DeleteHealthCheckRequestFunc                          func(param0 *route53.DeleteHealthCheckInput) (*request.Request, *route53.DeleteHealthCheckOutput)
	DeleteHealthCheckWithContextFunc                      func(param0 aws.Context, param1 *route53.DeleteHealthCheckInput, param2 ...request.Option) (*route53.DeleteHealthCheckOutput, error)
	DeleteHostedZoneFunc                                  func(param0 *route53.DeleteHostedZoneInput) (*route53.DeleteHostedZoneOutput, error)
	DeleteHostedZoneRequestFunc                           func(param0 *route53.DeleteHostedZoneInput) (*request.Request, *route53.DeleteHostedZoneOutput)
	DeleteHostedZoneWithContextFunc                       func(param0 aws.Context, param1 *route53.DeleteHostedZoneInput, param2 ...request.Option) (*route53.DeleteHostedZoneOutput, error)
	DeleteQueryLoggingConfigFunc                          func(param0 *route53.DeleteQueryLoggingConfigInput) (*route53.DeleteQueryLoggingConfigOutput, error)
	DeleteQueryLoggingConfigRequestFunc                   func(param0 *route53.DeleteQueryLoggingConfigInput) (*request.Request, *route53.DeleteQueryLoggingConfigOutput)
	DeleteQueryLoggingConfigWithContextFunc               func(param0 aws.Context, param1 *route53.DeleteQueryLoggingConfigInput, param2 ...request.Option) (*route53.DeleteQueryLoggingConfigOutput, error)
	DeleteReusableDelegationSetFunc                       func(param0 *route53.DeleteReusableDelegationSetInput) (*route53.DeleteReusableDelegationSetOutput, error)
	DeleteReusableDelegationSetRequestFunc                func(param0 *route53.DeleteReusableDelegationSetInput) (*request.Request, *route53.DeleteReusableDelegationSetOutput)
	DeleteReusableDelegationSetWithContextFunc            func(param0 aws.Context, param1 *route53.DeleteReusableDelegationSetInput, param2 ...request.Option) (*route53.DeleteReusableDelegationSetOutput, error)
	DeleteTrafficPolicyFunc                               func(param0 *route53.DeleteTrafficPolicyInput) (*route53.DeleteTrafficPolicyOutput, error)
	DeleteTrafficPolicyInstanceFunc                       func(param0 *route53.DeleteTrafficPolicyInstanceInput) (*route53.DeleteTrafficPolicyInstanceOutput, error)
	DeleteTrafficPolicyInstanceRequestFunc                func(param0 *route53.DeleteTrafficPolicyInstanceInput) (*request.Request, *route53.DeleteTrafficPolicyInstanceOutput)
	DeleteTrafficPolicyInstanceWithContextFunc            func(param0 aws.Context, param1 *route53.DeleteTrafficPolicyInstanceInput, param2 ...request.Option) (*route53.DeleteTrafficPolicyInstanceOutput, error)
	DeleteTrafficPolicyRequestFunc                        func(param0 *route53.DeleteTrafficPolicyInput) (*request.Request, *route53.DeleteTrafficPolicyOutput)
	DeleteTrafficPolicyWithContextFunc                    func(param0 aws.Context, param1 *route53.DeleteTrafficPolicyInput, param2 ...request.Option) (*route53.DeleteTrafficPolicyOutput, error)
	DeleteVPCAssociationAuthorizationFunc                 func(param0 *route53.DeleteVPCAssociationAuthorizationInput) (*route53.DeleteVPCAssociationAuthorizationOutput, error)
	DeleteVPCAssociationAuthorizationRequestFunc          func(param0 *route53.DeleteVPCAssociationAuthorizationInput) (*request.Request, *route53.DeleteVPCAssociationAuthorizationOutput)
	DeleteVPCAssociationAuthorizationWithContextFunc      func(param0 aws.Context, param1 *route53.DeleteVPCAssociationAuthorizationInput, param2 ...request.Option) (*route53.DeleteVPCAssociationAuthorizationOutput, error)
	DisassociateVPCFromHostedZoneFunc                     func(param0 *route53.DisassociateVPCFromHostedZoneInput) (*route53.DisassociateVPCFromHostedZoneOutput, error)
	DisassociateVPCFromHostedZoneRequestFunc              func(param0 *route53.DisassociateVPCFromHostedZoneInput) (*request.Request, *route53.DisassociateVPCFromHostedZoneOutput)
	DisassociateVPCFromHostedZoneWithContextFunc          func(param0 aws.Context, param1 *route53.DisassociateVPCFromHostedZoneInput, param2 ...request.Option) (*route53.DisassociateVPCFromHostedZoneOutput, error)
	GetChangeFunc                                         func(param0 *route53.GetChangeInput) (*route53.GetChangeOutput, error)
	GetChangeRequestFunc                                  func(param0 *route53.GetChangeInput) (*request.Request, *route53.GetChangeOutput)
	GetChangeWithContextFunc                              func(param0 aws.Context, param1 *route53.GetChangeInput, param2 ...request.Option) (*route53.GetChangeOutput, error)
	GetCheckerIpRangesFunc                                func(param0 *route53.GetCheckerIpRangesInput) (*route53.GetCheckerIpRangesOutput, error)
	GetCheckerIpRangesRequestFunc                         func(param0 *route53.GetCheckerIpRangesInput) (*request.Request, *route53.GetCheckerIpRangesOutput)
	GetCheckerIpRangesWithContextFunc                     func(param0 aws.Context, param1 *route53.GetCheckerIpRangesInput, param2 ...request.Option) (*route53.GetCheckerIpRangesOutput, error)
	GetGeoLocationFunc                                    func(param0 *route53.GetGeoLocationInput) (*route53.GetGeoLocationOutput, error)
	GetGeoLocationRequestFunc                             func(param0 *route53.GetGeoLocationInput) (*request.Request, *route53.GetGeoLocationOutput)
	GetGeoLocationWithContextFunc                         func(param0 aws.Context, param1 *route53.GetGeoLocationInput, param2 ...request.Option) (*route53.GetGeoLocationOutput, error)
	GetHealthCheckFunc                                    func(param0 *route53.GetHealthCheckInput) (*route53.GetHealthCheckOutput, error)
	GetHealthCheckCountFunc                               func(param0 *route53.GetHealthCheckCountInput) (*route53.GetHealthCheckCountOutput, error)
	GetHealthCheckCountRequestFunc                        func(param0 *route53.GetHealthCheckCountInput) (*request.Request, *route53.GetHealthCheckCountOutput)
	GetHealthCheckCountWithContextFunc                    func(param0 aws.Context, param1 *route53.GetHealthCheckCountInput, param2 ...request.Option) (*route53.GetHealthCheckCountOutput, error)
	GetHealthCheckLastFailureReasonFunc                   func(param0 *route53.GetHealthCheckLastFailureReasonInput) (*route53.GetHealthCheckLastFailureReasonOutput, error)
	GetHealthCheckLastFailureReasonRequestFunc            func(param0 *route53.GetHealthCheckLastFailureReasonInput) (*request.Request, *route53.GetHealthCheckLastFailureReasonOutput)
	GetHealthCheckLastFailureReasonWithContextFunc        func(param0 aws.Context, param1 *route53.GetHealthCheckLastFailureReasonInput, param2 ...request.Option) (*route53.GetHealthCheckLastFailureReasonOutput, error)
	GetHealthCheckRequestFunc                             func(param0 *route53.GetHealthCheckInput) (*request.Request, *route53.GetHealthCheckOutput)
	GetHealthCheckStatusFunc                              func(param0 *route53.GetHealthCheckStatusInput) (*route53.GetHealthCheckStatusOutput, error)
	GetHealthCheckStatusRequestFunc                       func(param0 *route53.GetHealthCheckStatusInput) (*request.Request, *route53.GetHealthCheckStatusOutput)
	GetHealthCheckStatusWithContextFunc                   func(param0 aws.Context, param1 *route53.GetHealthCheckStatusInput, param2 ...request.Option) (*route53.GetHealthCheckStatusOutput, error)
	GetHealthCheckWithContextFunc                         func(param0 aws.Context, param1 *route53.GetHealthCheckInput, param2 ...request.Option) (*route53.GetHealthCheckOutput, error)
	GetHostedZoneFunc                                     func(param0 *route53.GetHostedZoneInput) (*route53.GetHostedZoneOutput, error)
	GetHostedZoneCountFunc                                func(param0 *route53.GetHostedZoneCountInput) (*route53.GetHostedZoneCountOutput, error)
	GetHostedZoneCountRequestFunc                         func(param0 *route53.GetHostedZoneCountInput) (*request.Request, *route53.GetHostedZoneCountOutput)
	GetHostedZoneCountWithContextFunc                     func(param0 aws.Context, param1 *route53.GetHostedZoneCountInput, param2 ...request.Option) (*route53.GetHostedZoneCountOutput, error)
	GetHostedZoneRequestFunc                              func(param0 *route53.GetHostedZoneInput) (*request.Request, *route53.GetHostedZoneOutput)
	GetHostedZoneWithContextFunc                          func(param0 aws.Context, param1 *route53.GetHostedZoneInput, param2 ...request.Option) (*route53.GetHostedZoneOutput, error)
	GetQueryLoggingConfigFunc                             func(param0 *route53.GetQueryLoggingConfigInput) (*route53.GetQueryLoggingConfigOutput, error)
	GetQueryLoggingConfigRequestFunc                      func(param0 *route53.GetQueryLoggingConfigInput) (*request.Request, *route53.GetQueryLoggingConfigOutput)
	GetQueryLoggingConfigWithContextFunc                  func(param0 aws.Context, param1 *route53.GetQueryLoggingConfigInput, param2 ...request.Option) (*route53.GetQueryLoggingConfigOutput, error)
	GetReusableDelegationSetFunc                          func(param0 *route53.GetReusableDelegationSetInput) (*route53.GetReusableDelegationSetOutput, error)
	GetReusableDelegationSetRequestFunc                   func(param0 *route53.GetReusableDelegationSetInput) (*request.Request, *route53.GetReusableDelegationSetOutput)
	GetReusableDelegationSetWithContextFunc               func(param0 aws.Context, param1 *route53.GetReusableDelegationSetInput, param2 ...request.Option) (*route53.GetReusableDelegationSetOutput, error)
	GetTrafficPolicyFunc                                  func(param0 *route53.GetTrafficPolicyInput) (*route53.GetTrafficPolicyOutput, error)
	GetTrafficPolicyInstanceFunc                          func(param0 *route53.GetTrafficPolicyInstanceInput) (*route53.GetTrafficPolicyInstanceOutput, error)
	GetTrafficPolicyInstanceCountFunc                     func(param0 *route53.GetTrafficPolicyInstanceCountInput) (*route53.GetTrafficPolicyInstanceCountOutput, error)
	GetTrafficPolicyInstanceCountRequestFunc              func(param0 *route53.GetTrafficPolicyInstanceCountInput) (*request.Request, *route53.GetTrafficPolicyInstanceCountOutput)
	GetTrafficPolicyInstanceCountWithContextFunc          func(param0 aws.Context, param1 *route53.GetTrafficPolicyInstanceCountInput, param2 ...request.Option) (*route53.GetTrafficPolicyInstanceCountOutput, error)
	GetTrafficPolicyInstanceRequestFunc                   func(param0 *route53.GetTrafficPolicyInstanceInput) (*request.Request, *route53.GetTrafficPolicyInstanceOutput)
	GetTrafficPolicyInstanceWithContextFunc               func(param0 aws.Context, param1 *route53.GetTrafficPolicyInstanceInput, param2 ...request.Option) (*route53.GetTrafficPolicyInstanceOutput, error)
	GetTrafficPolicyRequestFunc                           func(param0 *route53.GetTrafficPolicyInput) (*request.Request, *route53.GetTrafficPolicyOutput)
	GetTrafficPolicyWithContextFunc                       func(param0 aws.Context, param1 *route53.GetTrafficPolicyInput, param2 ...request.Option) (*route53.GetTrafficPolicyOutput, error)
	ListGeoLocationsFunc                                  func(param0 *route53.ListGeoLocationsInput) (*route53.ListGeoLocationsOutput, error)
	ListGeoLocationsRequestFunc                           func(param0 *route53.ListGeoLocationsInput) (*request.Request, *route53.ListGeoLocationsOutput)
	ListGeoLocationsWithContextFunc                       func(param0 aws.Context, param1 *route53.ListGeoLocationsInput, param2 ...request.Option) (*route53.ListGeoLocationsOutput, error)
	ListHealthChecksFunc                                  func(param0 *route53.ListHealthChecksInput) (*route53.ListHealthChecksOutput, error)
	ListHealthChecksRequestFunc                           func(param0 *route53.ListHealthChecksInput) (*request.Request, *route53.ListHealthChecksOutput)
	ListHealthChecksWithContextFunc                       func(param0 aws.Context, param1 *route53.ListHealthChecksInput, param2 ...request.Option) (*route53.ListHealthChecksOutput, error)
	ListHostedZonesFunc                                   func(param0 *route53.ListHostedZonesInput) (*route53.ListHostedZonesOutput, error)
	ListHostedZonesByNameFunc                             func(param0 *route53.ListHostedZonesByNameInput) (*route53.ListHostedZonesByNameOutput, error)
	ListHostedZonesByNameRequestFunc                      func(param0 *route53.ListHostedZonesByNameInput) (*request.Request, *route53.ListHostedZonesByNameOutput)
	ListHostedZonesByNameWithContextFunc                  func(param0 aws.Context, param1 *route53.ListHostedZonesByNameInput, param2 ...request.Option) (*route53.ListHostedZonesByNameOutput, error)
	ListHostedZonesRequestFunc                            func(param0 *route53.ListHostedZonesInput) (*request.Request, *route53.ListHostedZonesOutput)
	ListHostedZonesWithContextFunc                        func(param0 aws.Context, param1 *route53.ListHostedZonesInput, param2 ...request.Option) (*route53.ListHostedZonesOutput, error)
	ListQueryLoggingConfigsFunc                           func(param0 *route53.ListQueryLoggingConfigsInput) (*route53.ListQueryLoggingConfigsOutput, error)
	ListQueryLoggingConfigsRequestFunc                    func(param0 *route53.ListQueryLoggingConfigsInput) (*request.Request, *route53.ListQueryLoggingConfigsOutput)
	ListQueryLoggingConfigsWithContextFunc                func(param0 aws.Context, param1 *route53.ListQueryLoggingConfigsInput, param2 ...request.Option) (*route53.ListQueryLoggingConfigsOutput, error)
	ListResourceRecordSetsFunc                            func(param0 *route53.ListResourceRecordSetsInput) (*route53.ListResourceRecordSetsOutput, error)
	ListResourceRecordSetsRequestFunc                     func(param0 *route53.ListResourceRecordSetsInput) (*request.Request, *route53.ListResourceRecordSetsOutput)
	ListResourceRecordSetsWithContextFunc                 func(param0 aws.Context, param1 *route53.ListResourceRecordSetsInput, param2 ...request.Option) (*route53.ListResourceRecordSetsOutput, error)
	ListReusableDelegationSetsFunc                        func(param0 *route53.ListReusableDelegationSetsInput) (*route53.ListReusableDelegationSetsOutput, error)
	ListReusableDelegationSetsRequestFunc                 func(param0 *route53.ListReusableDelegationSetsInput) (*request.Request, *route53.ListReusableDelegationSetsOutput)
	ListReusableDelegationSetsWithContextFunc             func(param0 aws.Context, param1 *route53.ListReusableDelegationSetsInput, param2 ...request.Option) (*route53.ListReusableDelegationSetsOutput, error)
	ListTagsForResourceFunc                               func(param0 *route53.ListTagsForResourceInput) (*route53.ListTagsForResourceOutput, error)
	ListTagsForResourceRequestFunc                        func(param0 *route53.ListTagsForResourceInput) (*request.Request, *route53.ListTagsForResourceOutput)
	ListTagsForResourceWithContextFunc                    func(param0 aws.Context, param1 *route53.ListTagsForResourceInput, param2 ...request.Option) (*route53.ListTagsForResourceOutput, error)
	ListTagsForResourcesFunc                              func(param0 *route53.ListTagsForResourcesInput) (*route53.ListTagsForResourcesOutput, error)
	ListTagsForResourcesRequestFunc                       func(param0 *route53.ListTagsForResourcesInput) (*request.Request, *route53.ListTagsForResourcesOutput)
	ListTagsForResourcesWithContextFunc                   func(param0 aws.Context, param1 *route53.ListTagsForResourcesInput, param2 ...request.Option) (*route53.ListTagsForResourcesOutput, error)
	ListTrafficPoliciesFunc                               func(param0 *route53.ListTrafficPoliciesInput) (*route53.ListTrafficPoliciesOutput, error)
	ListTrafficPoliciesRequestFunc                        func(param0 *route53.ListTrafficPoliciesInput) (*request.Request, *route53.ListTrafficPoliciesOutput)
	ListTrafficPoliciesWithContextFunc                    func(param0 aws.Context, param1 *route53.ListTrafficPoliciesInput, param2 ...request.Option) (*route53.ListTrafficPoliciesOutput, error)
	ListTrafficPolicyInstancesFunc                        func(param0 *route53.ListTrafficPolicyInstancesInput) (*route53.ListTrafficPolicyInstancesOutput, error)
	ListTrafficPolicyInstancesByHostedZoneFunc            func(param0 *route53.ListTrafficPolicyInstancesByHostedZoneInput) (*route53.ListTrafficPolicyInstancesByHostedZoneOutput, error)
	ListTrafficPolicyInstancesByHostedZoneRequestFunc     func(param0 *route53.ListTrafficPolicyInstancesByHostedZoneInput) (*request.Request, *route53.ListTrafficPolicyInstancesByHostedZoneOutput)
	ListTrafficPolicyInstancesByHostedZoneWithContextFunc func(param0 aws.Context, param1 *route53.ListTrafficPolicyInstancesByHostedZoneInput, param2 ...request.Option) (*route53.ListTrafficPolicyInstancesByHostedZoneOutput, error)
	ListTrafficPolicyInstancesByPolicyFunc                func(param0 *route53.ListTrafficPolicyInstancesByPolicyInput) (*route53.ListTrafficPolicyInstancesByPolicyOutput, error)
	ListTrafficPolicyInstancesByPolicyRequestFunc         func(param0 *route53.ListTrafficPolicyInstancesByPolicyInput) (*request.Request, *route53.ListTrafficPolicyInstancesByPolicyOutput)
	ListTrafficPolicyInstancesByPolicyWithContextFunc     func(param0 aws.Context, param1 *route53.ListTrafficPolicyInstancesByPolicyInput, param2 ...request.Option) (*route53.ListTrafficPolicyInstancesByPolicyOutput, error)
	ListTrafficPolicyInstancesRequestFunc                 func(param0 *route53.ListTrafficPolicyInstancesInput) (*request.Request, *route53.ListTrafficPolicyInstancesOutput)
	ListTrafficPolicyInstancesWithContextFunc             func(param0 aws.Context, param1 *route53.ListTrafficPolicyInstancesInput, param2 ...request.Option) (*route53.ListTrafficPolicyInstancesOutput, error)
	ListTrafficPolicyVersionsFunc                         func(param0 *route53.ListTrafficPolicyVersionsInput) (*route53.ListTrafficPolicyVersionsOutput, error)
	ListTrafficPolicyVersionsRequestFunc                  func(param0 *route53.ListTrafficPolicyVersionsInput) (*request.Request, *route53.ListTrafficPolicyVersionsOutput)
	ListTrafficPolicyVersionsWithContextFunc              func(param0 aws.Context, param1 *route53.ListTrafficPolicyVersionsInput, param2 ...request.Option) (*route53.ListTrafficPolicyVersionsOutput, error)
	ListVPCAssociationAuthorizationsFunc                  func(param0 *route53.ListVPCAssociationAuthorizationsInput) (*route53.ListVPCAssociationAuthorizationsOutput, error)
	ListVPCAssociationAuthorizationsRequestFunc           func(param0 *route53.ListVPCAssociationAuthorizationsInput) (*request.Request, *route53.ListVPCAssociationAuthorizationsOutput)
	ListVPCAssociationAuthorizationsWithContextFunc       func(param0 aws.Context, param1 *route53.ListVPCAssociationAuthorizationsInput, param2 ...request.Option) (*route53.ListVPCAssociationAuthorizationsOutput, error)
	TestDNSAnswerFunc                                     func(param0 *route53.TestDNSAnswerInput) (*route53.TestDNSAnswerOutput, error)
	TestDNSAnswerRequestFunc                              func(param0 *route53.TestDNSAnswerInput) (*request.Request, *route53.TestDNSAnswerOutput)
	TestDNSAnswerWithContextFunc                          func(param0 aws.Context, param1 *route53.TestDNSAnswerInput, param2 ...request.Option) (*route53.TestDNSAnswerOutput, error)
	UpdateHealthCheckFunc                                 func(param0 *route53.UpdateHealthCheckInput) (*route53.UpdateHealthCheckOutput, error)
	UpdateHealthCheckRequestFunc                          func(param0 *route53.UpdateHealthCheckInput) (*request.Request, *route53.UpdateHealthCheckOutput)
	UpdateHealthCheckWithContextFunc                      func(param0 aws.Context, param1 *route53.UpdateHealthCheckInput, param2 ...request.Option) (*route53.UpdateHealthCheckOutput, error)
	UpdateHostedZoneCommentFunc                           func(param0 *route53.UpdateHostedZoneCommentInput) (*route53.UpdateHostedZoneCommentOutput, error)
	UpdateHostedZoneCommentRequestFunc                    func(param0 *route53.UpdateHostedZoneCommentInput) (*request.Request, *route53.UpdateHostedZoneCommentOutput)
	UpdateHostedZoneCommentWithContextFunc                func(param0 aws.Context, param1 *route53.UpdateHostedZoneCommentInput, param2 ...request.Option) (*route53.UpdateHostedZoneCommentOutput, error)
	UpdateTrafficPolicyCommentFunc                        func(param0 *route53.UpdateTrafficPolicyCommentInput) (*route53.UpdateTrafficPolicyCommentOutput, error)
	UpdateTrafficPolicyCommentRequestFunc                 func(param0 *route53.UpdateTrafficPolicyCommentInput) (*request.Request, *route53.UpdateTrafficPolicyCommentOutput)
	UpdateTrafficPolicyCommentWithContextFunc             func(param0 aws.Context, param1 *route53.UpdateTrafficPolicyCommentInput, param2 ...request.Option) (*route53.UpdateTrafficPolicyCommentOutput, error)
	UpdateTrafficPolicyInstanceFunc                       func(param0 *route53.UpdateTrafficPolicyInstanceInput) (*route53.UpdateTrafficPolicyInstanceOutput, error)
	UpdateTrafficPolicyInstanceRequestFunc                func(param0 *route53.UpdateTrafficPolicyInstanceInput) (*request.Request, *route53.UpdateTrafficPolicyInstanceOutput)
	UpdateTrafficPolicyInstanceWithContextFunc            func(param0 aws.Context, param1 *route53.UpdateTrafficPolicyInstanceInput, param2 ...request.Option) (*route53.UpdateTrafficPolicyInstanceOutput, error)
	WaitUntilResourceRecordSetsChangedFunc                func(param0 *route53.GetChangeInput) error
	WaitUntilResourceRecordSetsChangedWithContextFunc     func(param0 aws.Context, param1 *route53.GetChangeInput, param2 ...request.WaiterOption) error
}

func (m *route53Mock) AssociateVPCWithHostedZone(param0 *route53.AssociateVPCWithHostedZoneInput) (*route53.AssociateVPCWithHostedZoneOutput, error) {
	m.addCall("AssociateVPCWithHostedZone")
	m.verifyInput("AssociateVPCWithHostedZone", param0)
	return m.AssociateVPCWithHostedZoneFunc(param0)
}

func (m *route53Mock) AssociateVPCWithHostedZoneRequest(param0 *route53.AssociateVPCWithHostedZoneInput) (*request.Request, *route53.AssociateVPCWithHostedZoneOutput) {
	m.addCall("AssociateVPCWithHostedZoneRequest")
	m.verifyInput("AssociateVPCWithHostedZoneRequest", param0)
	return m.AssociateVPCWithHostedZoneRequestFunc(param0)
}

func (m *route53Mock) AssociateVPCWithHostedZoneWithContext(param0 aws.Context, param1 *route53.AssociateVPCWithHostedZoneInput, param2 ...request.Option) (*route53.AssociateVPCWithHostedZoneOutput, error) {
	m.addCall("AssociateVPCWithHostedZoneWithContext")
	m.verifyInput("AssociateVPCWithHostedZoneWithContext", param0)
	return m.AssociateVPCWithHostedZoneWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) ChangeResourceRecordSets(param0 *route53.ChangeResourceRecordSetsInput) (*route53.ChangeResourceRecordSetsOutput, error) {
	m.addCall("ChangeResourceRecordSets")
	m.verifyInput("ChangeResourceRecordSets", param0)
	return m.ChangeResourceRecordSetsFunc(param0)
}

func (m *route53Mock) ChangeResourceRecordSetsRequest(param0 *route53.ChangeResourceRecordSetsInput) (*request.Request, *route53.ChangeResourceRecordSetsOutput) {
	m.addCall("ChangeResourceRecordSetsRequest")
	m.verifyInput("ChangeResourceRecordSetsRequest", param0)
	return m.ChangeResourceRecordSetsRequestFunc(param0)
}

func (m *route53Mock) ChangeResourceRecordSetsWithContext(param0 aws.Context, param1 *route53.ChangeResourceRecordSetsInput, param2 ...request.Option) (*route53.ChangeResourceRecordSetsOutput, error) {
	m.addCall("ChangeResourceRecordSetsWithContext")
	m.verifyInput("ChangeResourceRecordSetsWithContext", param0)
	return m.ChangeResourceRecordSetsWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) ChangeTagsForResource(param0 *route53.ChangeTagsForResourceInput) (*route53.ChangeTagsForResourceOutput, error) {
	m.addCall("ChangeTagsForResource")
	m.verifyInput("ChangeTagsForResource", param0)
	return m.ChangeTagsForResourceFunc(param0)
}

func (m *route53Mock) ChangeTagsForResourceRequest(param0 *route53.ChangeTagsForResourceInput) (*request.Request, *route53.ChangeTagsForResourceOutput) {
	m.addCall("ChangeTagsForResourceRequest")
	m.verifyInput("ChangeTagsForResourceRequest", param0)
	return m.ChangeTagsForResourceRequestFunc(param0)
}

func (m *route53Mock) ChangeTagsForResourceWithContext(param0 aws.Context, param1 *route53.ChangeTagsForResourceInput, param2 ...request.Option) (*route53.ChangeTagsForResourceOutput, error) {
	m.addCall("ChangeTagsForResourceWithContext")
	m.verifyInput("ChangeTagsForResourceWithContext", param0)
	return m.ChangeTagsForResourceWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) CreateHealthCheck(param0 *route53.CreateHealthCheckInput) (*route53.CreateHealthCheckOutput, error) {
	m.addCall("CreateHealthCheck")
	m.verifyInput("CreateHealthCheck", param0)
	return m.CreateHealthCheckFunc(param0)
}

func (m *route53Mock) CreateHealthCheckRequest(param0 *route53.CreateHealthCheckInput) (*request.Request, *route53.CreateHealthCheckOutput) {
	m.addCall("CreateHealthCheckRequest")
	m.verifyInput("CreateHealthCheckRequest", param0)
	return m.CreateHealthCheckRequestFunc(param0)
}

func (m *route53Mock) CreateHealthCheckWithContext(param0 aws.Context, param1 *route53.CreateHealthCheckInput, param2 ...request.Option) (*route53.CreateHealthCheckOutput, error) {
	m.addCall("CreateHealthCheckWithContext")
	m.verifyInput("CreateHealthCheckWithContext", param0)
	return m.CreateHealthCheckWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) CreateHostedZone(param0 *route53.CreateHostedZoneInput) (*route53.CreateHostedZoneOutput, error) {
	m.addCall("CreateHostedZone")
	m.verifyInput("CreateHostedZone", param0)
	return m.CreateHostedZoneFunc(param0)
}

func (m *route53Mock) CreateHostedZoneRequest(param0 *route53.CreateHostedZoneInput) (*request.Request, *route53.CreateHostedZoneOutput) {
	m.addCall("CreateHostedZoneRequest")
	m.verifyInput("CreateHostedZoneRequest", param0)
	return m.CreateHostedZoneRequestFunc(param0)
}

func (m *route53Mock) CreateHostedZoneWithContext(param0 aws.Context, param1 *route53.CreateHostedZoneInput, param2 ...request.Option) (*route53.CreateHostedZoneOutput, error) {
	m.addCall("CreateHostedZoneWithContext")
	m.verifyInput("CreateHostedZoneWithContext", param0)
	return m.CreateHostedZoneWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) CreateQueryLoggingConfig(param0 *route53.CreateQueryLoggingConfigInput) (*route53.CreateQueryLoggingConfigOutput, error) {
	m.addCall("CreateQueryLoggingConfig")
	m.verifyInput("CreateQueryLoggingConfig", param0)
	return m.CreateQueryLoggingConfigFunc(param0)
}

func (m *route53Mock) CreateQueryLoggingConfigRequest(param0 *route53.CreateQueryLoggingConfigInput) (*request.Request, *route53.CreateQueryLoggingConfigOutput) {
	m.addCall("CreateQueryLoggingConfigRequest")
	m.verifyInput("CreateQueryLoggingConfigRequest", param0)
	return m.CreateQueryLoggingConfigRequestFunc(param0)
}

func (m *route53Mock) CreateQueryLoggingConfigWithContext(param0 aws.Context, param1 *route53.CreateQueryLoggingConfigInput, param2 ...request.Option) (*route53.CreateQueryLoggingConfigOutput, error) {
	m.addCall("CreateQueryLoggingConfigWithContext")
	m.verifyInput("CreateQueryLoggingConfigWithContext", param0)
	return m.CreateQueryLoggingConfigWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) CreateReusableDelegationSet(param0 *route53.CreateReusableDelegationSetInput) (*route53.CreateReusableDelegationSetOutput, error) {
	m.addCall("CreateReusableDelegationSet")
	m.verifyInput("CreateReusableDelegationSet", param0)
	return m.CreateReusableDelegationSetFunc(param0)
}

func (m *route53Mock) CreateReusableDelegationSetRequest(param0 *route53.CreateReusableDelegationSetInput) (*request.Request, *route53.CreateReusableDelegationSetOutput) {
	m.addCall("CreateReusableDelegationSetRequest")
	m.verifyInput("CreateReusableDelegationSetRequest", param0)
	return m.CreateReusableDelegationSetRequestFunc(param0)
}

func (m *route53Mock) CreateReusableDelegationSetWithContext(param0 aws.Context, param1 *route53.CreateReusableDelegationSetInput, param2 ...request.Option) (*route53.CreateReusableDelegationSetOutput, error) {
	m.addCall("CreateReusableDelegationSetWithContext")
	m.verifyInput("CreateReusableDelegationSetWithContext", param0)
	return m.CreateReusableDelegationSetWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) CreateTrafficPolicy(param0 *route53.CreateTrafficPolicyInput) (*route53.CreateTrafficPolicyOutput, error) {
	m.addCall("CreateTrafficPolicy")
	m.verifyInput("CreateTrafficPolicy", param0)
	return m.CreateTrafficPolicyFunc(param0)
}

func (m *route53Mock) CreateTrafficPolicyInstance(param0 *route53.CreateTrafficPolicyInstanceInput) (*route53.CreateTrafficPolicyInstanceOutput, error) {
	m.addCall("CreateTrafficPolicyInstance")
	m.verifyInput("CreateTrafficPolicyInstance", param0)
	return m.CreateTrafficPolicyInstanceFunc(param0)
}

func (m *route53Mock) CreateTrafficPolicyInstanceRequest(param0 *route53.CreateTrafficPolicyInstanceInput) (*request.Request, *route53.CreateTrafficPolicyInstanceOutput) {
	m.addCall("CreateTrafficPolicyInstanceRequest")
	m.verifyInput("CreateTrafficPolicyInstanceRequest", param0)
	return m.CreateTrafficPolicyInstanceRequestFunc(param0)
}

func (m *route53Mock) CreateTrafficPolicyInstanceWithContext(param0 aws.Context, param1 *route53.CreateTrafficPolicyInstanceInput, param2 ...request.Option) (*route53.CreateTrafficPolicyInstanceOutput, error) {
	m.addCall("CreateTrafficPolicyInstanceWithContext")
	m.verifyInput("CreateTrafficPolicyInstanceWithContext", param0)
	return m.CreateTrafficPolicyInstanceWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) CreateTrafficPolicyRequest(param0 *route53.CreateTrafficPolicyInput) (*request.Request, *route53.CreateTrafficPolicyOutput) {
	m.addCall("CreateTrafficPolicyRequest")
	m.verifyInput("CreateTrafficPolicyRequest", param0)
	return m.CreateTrafficPolicyRequestFunc(param0)
}

func (m *route53Mock) CreateTrafficPolicyVersion(param0 *route53.CreateTrafficPolicyVersionInput) (*route53.CreateTrafficPolicyVersionOutput, error) {
	m.addCall("CreateTrafficPolicyVersion")
	m.verifyInput("CreateTrafficPolicyVersion", param0)
	return m.CreateTrafficPolicyVersionFunc(param0)
}

func (m *route53Mock) CreateTrafficPolicyVersionRequest(param0 *route53.CreateTrafficPolicyVersionInput) (*request.Request, *route53.CreateTrafficPolicyVersionOutput) {
	m.addCall("CreateTrafficPolicyVersionRequest")
	m.verifyInput("CreateTrafficPolicyVersionRequest", param0)
	return m.CreateTrafficPolicyVersionRequestFunc(param0)
}

func (m *route53Mock) CreateTrafficPolicyVersionWithContext(param0 aws.Context, param1 *route53.CreateTrafficPolicyVersionInput, param2 ...request.Option) (*route53.CreateTrafficPolicyVersionOutput, error) {
	m.addCall("CreateTrafficPolicyVersionWithContext")
	m.verifyInput("CreateTrafficPolicyVersionWithContext", param0)
	return m.CreateTrafficPolicyVersionWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) CreateTrafficPolicyWithContext(param0 aws.Context, param1 *route53.CreateTrafficPolicyInput, param2 ...request.Option) (*route53.CreateTrafficPolicyOutput, error) {
	m.addCall("CreateTrafficPolicyWithContext")
	m.verifyInput("CreateTrafficPolicyWithContext", param0)
	return m.CreateTrafficPolicyWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) CreateVPCAssociationAuthorization(param0 *route53.CreateVPCAssociationAuthorizationInput) (*route53.CreateVPCAssociationAuthorizationOutput, error) {
	m.addCall("CreateVPCAssociationAuthorization")
	m.verifyInput("CreateVPCAssociationAuthorization", param0)
	return m.CreateVPCAssociationAuthorizationFunc(param0)
}

func (m *route53Mock) CreateVPCAssociationAuthorizationRequest(param0 *route53.CreateVPCAssociationAuthorizationInput) (*request.Request, *route53.CreateVPCAssociationAuthorizationOutput) {
	m.addCall("CreateVPCAssociationAuthorizationRequest")
	m.verifyInput("CreateVPCAssociationAuthorizationRequest", param0)
	return m.CreateVPCAssociationAuthorizationRequestFunc(param0)
}

func (m *route53Mock) CreateVPCAssociationAuthorizationWithContext(param0 aws.Context, param1 *route53.CreateVPCAssociationAuthorizationInput, param2 ...request.Option) (*route53.CreateVPCAssociationAuthorizationOutput, error) {
	m.addCall("CreateVPCAssociationAuthorizationWithContext")
	m.verifyInput("CreateVPCAssociationAuthorizationWithContext", param0)
	return m.CreateVPCAssociationAuthorizationWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) DeleteHealthCheck(param0 *route53.DeleteHealthCheckInput) (*route53.DeleteHealthCheckOutput, error) {
	m.addCall("DeleteHealthCheck")
	m.verifyInput("DeleteHealthCheck", param0)
	return m.DeleteHealthCheckFunc(param0)
}

func (m *route53Mock) DeleteHealthCheckRequest(param0 *route53.DeleteHealthCheckInput) (*request.Request, *route53.DeleteHealthCheckOutput) {
	m.addCall("DeleteHealthCheckRequest")
	m.verifyInput("DeleteHealthCheckRequest", param0)
	return m.DeleteHealthCheckRequestFunc(param0)
}

func (m *route53Mock) DeleteHealthCheckWithContext(param0 aws.Context, param1 *route53.DeleteHealthCheckInput, param2 ...request.Option) (*route53.DeleteHealthCheckOutput, error) {
	m.addCall("DeleteHealthCheckWithContext")
	m.verifyInput("DeleteHealthCheckWithContext", param0)
	return m.DeleteHealthCheckWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) DeleteHostedZone(param0 *route53.DeleteHostedZoneInput) (*route53.DeleteHostedZoneOutput, error) {
	m.addCall("DeleteHostedZone")
	m.verifyInput("DeleteHostedZone", param0)
	return m.DeleteHostedZoneFunc(param0)
}

func (m *route53Mock) DeleteHostedZoneRequest(param0 *route53.DeleteHostedZoneInput) (*request.Request, *route53.DeleteHostedZoneOutput) {
	m.addCall("DeleteHostedZoneRequest")
	m.verifyInput("DeleteHostedZoneRequest", param0)
	return m.DeleteHostedZoneRequestFunc(param0)
}

func (m *route53Mock) DeleteHostedZoneWithContext(param0 aws.Context, param1 *route53.DeleteHostedZoneInput, param2 ...request.Option) (*route53.DeleteHostedZoneOutput, error) {
	m.addCall("DeleteHostedZoneWithContext")
	m.verifyInput("DeleteHostedZoneWithContext", param0)
	return m.DeleteHostedZoneWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) DeleteQueryLoggingConfig(param0 *route53.DeleteQueryLoggingConfigInput) (*route53.DeleteQueryLoggingConfigOutput, error) {
	m.addCall("DeleteQueryLoggingConfig")
	m.verifyInput("DeleteQueryLoggingConfig", param0)
	return m.DeleteQueryLoggingConfigFunc(param0)
}

func (m *route53Mock) DeleteQueryLoggingConfigRequest(param0 *route53.DeleteQueryLoggingConfigInput) (*request.Request, *route53.DeleteQueryLoggingConfigOutput) {
	m.addCall("DeleteQueryLoggingConfigRequest")
	m.verifyInput("DeleteQueryLoggingConfigRequest", param0)
	return m.DeleteQueryLoggingConfigRequestFunc(param0)
}

func (m *route53Mock) DeleteQueryLoggingConfigWithContext(param0 aws.Context, param1 *route53.DeleteQueryLoggingConfigInput, param2 ...request.Option) (*route53.DeleteQueryLoggingConfigOutput, error) {
	m.addCall("DeleteQueryLoggingConfigWithContext")
	m.verifyInput("DeleteQueryLoggingConfigWithContext", param0)
	return m.DeleteQueryLoggingConfigWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) DeleteReusableDelegationSet(param0 *route53.DeleteReusableDelegationSetInput) (*route53.DeleteReusableDelegationSetOutput, error) {
	m.addCall("DeleteReusableDelegationSet")
	m.verifyInput("DeleteReusableDelegationSet", param0)
	return m.DeleteReusableDelegationSetFunc(param0)
}

func (m *route53Mock) DeleteReusableDelegationSetRequest(param0 *route53.DeleteReusableDelegationSetInput) (*request.Request, *route53.DeleteReusableDelegationSetOutput) {
	m.addCall("DeleteReusableDelegationSetRequest")
	m.verifyInput("DeleteReusableDelegationSetRequest", param0)
	return m.DeleteReusableDelegationSetRequestFunc(param0)
}

func (m *route53Mock) DeleteReusableDelegationSetWithContext(param0 aws.Context, param1 *route53.DeleteReusableDelegationSetInput, param2 ...request.Option) (*route53.DeleteReusableDelegationSetOutput, error) {
	m.addCall("DeleteReusableDelegationSetWithContext")
	m.verifyInput("DeleteReusableDelegationSetWithContext", param0)
	return m.DeleteReusableDelegationSetWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) DeleteTrafficPolicy(param0 *route53.DeleteTrafficPolicyInput) (*route53.DeleteTrafficPolicyOutput, error) {
	m.addCall("DeleteTrafficPolicy")
	m.verifyInput("DeleteTrafficPolicy", param0)
	return m.DeleteTrafficPolicyFunc(param0)
}

func (m *route53Mock) DeleteTrafficPolicyInstance(param0 *route53.DeleteTrafficPolicyInstanceInput) (*route53.DeleteTrafficPolicyInstanceOutput, error) {
	m.addCall("DeleteTrafficPolicyInstance")
	m.verifyInput("DeleteTrafficPolicyInstance", param0)
	return m.DeleteTrafficPolicyInstanceFunc(param0)
}

func (m *route53Mock) DeleteTrafficPolicyInstanceRequest(param0 *route53.DeleteTrafficPolicyInstanceInput) (*request.Request, *route53.DeleteTrafficPolicyInstanceOutput) {
	m.addCall("DeleteTrafficPolicyInstanceRequest")
	m.verifyInput("DeleteTrafficPolicyInstanceRequest", param0)
	return m.DeleteTrafficPolicyInstanceRequestFunc(param0)
}

func (m *route53Mock) DeleteTrafficPolicyInstanceWithContext(param0 aws.Context, param1 *route53.DeleteTrafficPolicyInstanceInput, param2 ...request.Option) (*route53.DeleteTrafficPolicyInstanceOutput, error) {
	m.addCall("DeleteTrafficPolicyInstanceWithContext")
	m.verifyInput("DeleteTrafficPolicyInstanceWithContext", param0)
	return m.DeleteTrafficPolicyInstanceWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) DeleteTrafficPolicyRequest(param0 *route53.DeleteTrafficPolicyInput) (*request.Request, *route53.DeleteTrafficPolicyOutput) {
	m.addCall("DeleteTrafficPolicyRequest")
	m.verifyInput("DeleteTrafficPolicyRequest", param0)
	return m.DeleteTrafficPolicyRequestFunc(param0)
}

func (m *route53Mock) DeleteTrafficPolicyWithContext(param0 aws.Context, param1 *route53.DeleteTrafficPolicyInput, param2 ...request.Option) (*route53.DeleteTrafficPolicyOutput, error) {
	m.addCall("DeleteTrafficPolicyWithContext")
	m.verifyInput("DeleteTrafficPolicyWithContext", param0)
	return m.DeleteTrafficPolicyWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) DeleteVPCAssociationAuthorization(param0 *route53.DeleteVPCAssociationAuthorizationInput) (*route53.DeleteVPCAssociationAuthorizationOutput, error) {
	m.addCall("DeleteVPCAssociationAuthorization")
	m.verifyInput("DeleteVPCAssociationAuthorization", param0)
	return m.DeleteVPCAssociationAuthorizationFunc(param0)
}

func (m *route53Mock) DeleteVPCAssociationAuthorizationRequest(param0 *route53.DeleteVPCAssociationAuthorizationInput) (*request.Request, *route53.DeleteVPCAssociationAuthorizationOutput) {
	m.addCall("DeleteVPCAssociationAuthorizationRequest")
	m.verifyInput("DeleteVPCAssociationAuthorizationRequest", param0)
	return m.DeleteVPCAssociationAuthorizationRequestFunc(param0)
}

func (m *route53Mock) DeleteVPCAssociationAuthorizationWithContext(param0 aws.Context, param1 *route53.DeleteVPCAssociationAuthorizationInput, param2 ...request.Option) (*route53.DeleteVPCAssociationAuthorizationOutput, error) {
	m.addCall("DeleteVPCAssociationAuthorizationWithContext")
	m.verifyInput("DeleteVPCAssociationAuthorizationWithContext", param0)
	return m.DeleteVPCAssociationAuthorizationWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) DisassociateVPCFromHostedZone(param0 *route53.DisassociateVPCFromHostedZoneInput) (*route53.DisassociateVPCFromHostedZoneOutput, error) {
	m.addCall("DisassociateVPCFromHostedZone")
	m.verifyInput("DisassociateVPCFromHostedZone", param0)
	return m.DisassociateVPCFromHostedZoneFunc(param0)
}

func (m *route53Mock) DisassociateVPCFromHostedZoneRequest(param0 *route53.DisassociateVPCFromHostedZoneInput) (*request.Request, *route53.DisassociateVPCFromHostedZoneOutput) {
	m.addCall("DisassociateVPCFromHostedZoneRequest")
	m.verifyInput("DisassociateVPCFromHostedZoneRequest", param0)
	return m.DisassociateVPCFromHostedZoneRequestFunc(param0)
}

func (m *route53Mock) DisassociateVPCFromHostedZoneWithContext(param0 aws.Context, param1 *route53.DisassociateVPCFromHostedZoneInput, param2 ...request.Option) (*route53.DisassociateVPCFromHostedZoneOutput, error) {
	m.addCall("DisassociateVPCFromHostedZoneWithContext")
	m.verifyInput("DisassociateVPCFromHostedZoneWithContext", param0)
	return m.DisassociateVPCFromHostedZoneWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) GetChange(param0 *route53.GetChangeInput) (*route53.GetChangeOutput, error) {
	m.addCall("GetChange")
	m.verifyInput("GetChange", param0)
	return m.GetChangeFunc(param0)
}

func (m *route53Mock) GetChangeRequest(param0 *route53.GetChangeInput) (*request.Request, *route53.GetChangeOutput) {
	m.addCall("GetChangeRequest")
	m.verifyInput("GetChangeRequest", param0)
	return m.GetChangeRequestFunc(param0)
}

func (m *route53Mock) GetChangeWithContext(param0 aws.Context, param1 *route53.GetChangeInput, param2 ...request.Option) (*route53.GetChangeOutput, error) {
	m.addCall("GetChangeWithContext")
	m.verifyInput("GetChangeWithContext", param0)
	return m.GetChangeWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) GetCheckerIpRanges(param0 *route53.GetCheckerIpRangesInput) (*route53.GetCheckerIpRangesOutput, error) {
	m.addCall("GetCheckerIpRanges")
	m.verifyInput("GetCheckerIpRanges", param0)
	return m.GetCheckerIpRangesFunc(param0)
}

func (m *route53Mock) GetCheckerIpRangesRequest(param0 *route53.GetCheckerIpRangesInput) (*request.Request, *route53.GetCheckerIpRangesOutput) {
	m.addCall("GetCheckerIpRangesRequest")
	m.verifyInput("GetCheckerIpRangesRequest", param0)
	return m.GetCheckerIpRangesRequestFunc(param0)
}

func (m *route53Mock) GetCheckerIpRangesWithContext(param0 aws.Context, param1 *route53.GetCheckerIpRangesInput, param2 ...request.Option) (*route53.GetCheckerIpRangesOutput, error) {
	m.addCall("GetCheckerIpRangesWithContext")
	m.verifyInput("GetCheckerIpRangesWithContext", param0)
	return m.GetCheckerIpRangesWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) GetGeoLocation(param0 *route53.GetGeoLocationInput) (*route53.GetGeoLocationOutput, error) {
	m.addCall("GetGeoLocation")
	m.verifyInput("GetGeoLocation", param0)
	return m.GetGeoLocationFunc(param0)
}

func (m *route53Mock) GetGeoLocationRequest(param0 *route53.GetGeoLocationInput) (*request.Request, *route53.GetGeoLocationOutput) {
	m.addCall("GetGeoLocationRequest")
	m.verifyInput("GetGeoLocationRequest", param0)
	return m.GetGeoLocationRequestFunc(param0)
}

func (m *route53Mock) GetGeoLocationWithContext(param0 aws.Context, param1 *route53.GetGeoLocationInput, param2 ...request.Option) (*route53.GetGeoLocationOutput, error) {
	m.addCall("GetGeoLocationWithContext")
	m.verifyInput("GetGeoLocationWithContext", param0)
	return m.GetGeoLocationWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) GetHealthCheck(param0 *route53.GetHealthCheckInput) (*route53.GetHealthCheckOutput, error) {
	m.addCall("GetHealthCheck")
	m.verifyInput("GetHealthCheck", param0)
	return m.GetHealthCheckFunc(param0)
}

func (m *route53Mock) GetHealthCheckCount(param0 *route53.GetHealthCheckCountInput) (*route53.GetHealthCheckCountOutput, error) {
	m.addCall("GetHealthCheckCount")
	m.verifyInput("GetHealthCheckCount", param0)
	return m.GetHealthCheckCountFunc(param0)
}

func (m *route53Mock) GetHealthCheckCountRequest(param0 *route53.GetHealthCheckCountInput) (*request.Request, *route53.GetHealthCheckCountOutput) {
	m.addCall("GetHealthCheckCountRequest")
	m.verifyInput("GetHealthCheckCountRequest", param0)
	return m.GetHealthCheckCountRequestFunc(param0)
}

func (m *route53Mock) GetHealthCheckCountWithContext(param0 aws.Context, param1 *route53.GetHealthCheckCountInput, param2 ...request.Option) (*route53.GetHealthCheckCountOutput, error) {
	m.addCall("GetHealthCheckCountWithContext")
	m.verifyInput("GetHealthCheckCountWithContext", param0)
	return m.GetHealthCheckCountWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) GetHealthCheckLastFailureReason(param0 *route53.GetHealthCheckLastFailureReasonInput) (*route53.GetHealthCheckLastFailureReasonOutput, error) {
	m.addCall("GetHealthCheckLastFailureReason")
	m.verifyInput("GetHealthCheckLastFailureReason", param0)
	return m.GetHealthCheckLastFailureReasonFunc(param0)
}

func (m *route53Mock) GetHealthCheckLastFailureReasonRequest(param0 *route53.GetHealthCheckLastFailureReasonInput) (*request.Request, *route53.GetHealthCheckLastFailureReasonOutput) {
	m.addCall("GetHealthCheckLastFailureReasonRequest")
	m.verifyInput("GetHealthCheckLastFailureReasonRequest", param0)
	return m.GetHealthCheckLastFailureReasonRequestFunc(param0)
}

func (m *route53Mock) GetHealthCheckLastFailureReasonWithContext(param0 aws.Context, param1 *route53.GetHealthCheckLastFailureReasonInput, param2 ...request.Option) (*route53.GetHealthCheckLastFailureReasonOutput, error) {
	m.addCall("GetHealthCheckLastFailureReasonWithContext")
	m.verifyInput("GetHealthCheckLastFailureReasonWithContext", param0)
	return m.GetHealthCheckLastFailureReasonWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) GetHealthCheckRequest(param0 *route53.GetHealthCheckInput) (*request.Request, *route53.GetHealthCheckOutput) {
	m.addCall("GetHealthCheckRequest")
	m.verifyInput("GetHealthCheckRequest", param0)
	return m.GetHealthCheckRequestFunc(param0)
}

func (m *route53Mock) GetHealthCheckStatus(param0 *route53.GetHealthCheckStatusInput) (*route53.GetHealthCheckStatusOutput, error) {
	m.addCall("GetHealthCheckStatus")
	m.verifyInput("GetHealthCheckStatus", param0)
	return m.GetHealthCheckStatusFunc(param0)
}

func (m *route53Mock) GetHealthCheckStatusRequest(param0 *route53.GetHealthCheckStatusInput) (*request.Request, *route53.GetHealthCheckStatusOutput) {
	m.addCall("GetHealthCheckStatusRequest")
	m.verifyInput("GetHealthCheckStatusRequest", param0)
	return m.GetHealthCheckStatusRequestFunc(param0)
}

func (m *route53Mock) GetHealthCheckStatusWithContext(param0 aws.Context, param1 *route53.GetHealthCheckStatusInput, param2 ...request.Option) (*route53.GetHealthCheckStatusOutput, error) {
	m.addCall("GetHealthCheckStatusWithContext")
	m.verifyInput("GetHealthCheckStatusWithContext", param0)
	return m.GetHealthCheckStatusWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) GetHealthCheckWithContext(param0 aws.Context, param1 *route53.GetHealthCheckInput, param2 ...request.Option) (*route53.GetHealthCheckOutput, error) {
	m.addCall("GetHealthCheckWithContext")
	m.verifyInput("GetHealthCheckWithContext", param0)
	return m.GetHealthCheckWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) GetHostedZone(param0 *route53.GetHostedZoneInput) (*route53.GetHostedZoneOutput, error) {
	m.addCall("GetHostedZone")
	m.verifyInput("GetHostedZone", param0)
	return m.GetHostedZoneFunc(param0)
}

func (m *route53Mock) GetHostedZoneCount(param0 *route53.GetHostedZoneCountInput) (*route53.GetHostedZoneCountOutput, error) {
	m.addCall("GetHostedZoneCount")
	m.verifyInput("GetHostedZoneCount", param0)
	return m.GetHostedZoneCountFunc(param0)
}

func (m *route53Mock) GetHostedZoneCountRequest(param0 *route53.GetHostedZoneCountInput) (*request.Request, *route53.GetHostedZoneCountOutput) {
	m.addCall("GetHostedZoneCountRequest")
	m.verifyInput("GetHostedZoneCountRequest", param0)
	return m.GetHostedZoneCountRequestFunc(param0)
}

func (m *route53Mock) GetHostedZoneCountWithContext(param0 aws.Context, param1 *route53.GetHostedZoneCountInput, param2 ...request.Option) (*route53.GetHostedZoneCountOutput, error) {
	m.addCall("GetHostedZoneCountWithContext")
	m.verifyInput("GetHostedZoneCountWithContext", param0)
	return m.GetHostedZoneCountWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) GetHostedZoneRequest(param0 *route53.GetHostedZoneInput) (*request.Request, *route53.GetHostedZoneOutput) {
	m.addCall("GetHostedZoneRequest")
	m.verifyInput("GetHostedZoneRequest", param0)
	return m.GetHostedZoneRequestFunc(param0)
}

func (m *route53Mock) GetHostedZoneWithContext(param0 aws.Context, param1 *route53.GetHostedZoneInput, param2 ...request.Option) (*route53.GetHostedZoneOutput, error) {
	m.addCall("GetHostedZoneWithContext")
	m.verifyInput("GetHostedZoneWithContext", param0)
	return m.GetHostedZoneWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) GetQueryLoggingConfig(param0 *route53.GetQueryLoggingConfigInput) (*route53.GetQueryLoggingConfigOutput, error) {
	m.addCall("GetQueryLoggingConfig")
	m.verifyInput("GetQueryLoggingConfig", param0)
	return m.GetQueryLoggingConfigFunc(param0)
}

func (m *route53Mock) GetQueryLoggingConfigRequest(param0 *route53.GetQueryLoggingConfigInput) (*request.Request, *route53.GetQueryLoggingConfigOutput) {
	m.addCall("GetQueryLoggingConfigRequest")
	m.verifyInput("GetQueryLoggingConfigRequest", param0)
	return m.GetQueryLoggingConfigRequestFunc(param0)
}

func (m *route53Mock) GetQueryLoggingConfigWithContext(param0 aws.Context, param1 *route53.GetQueryLoggingConfigInput, param2 ...request.Option) (*route53.GetQueryLoggingConfigOutput, error) {
	m.addCall("GetQueryLoggingConfigWithContext")
	m.verifyInput("GetQueryLoggingConfigWithContext", param0)
	return m.GetQueryLoggingConfigWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) GetReusableDelegationSet(param0 *route53.GetReusableDelegationSetInput) (*route53.GetReusableDelegationSetOutput, error) {
	m.addCall("GetReusableDelegationSet")
	m.verifyInput("GetReusableDelegationSet", param0)
	return m.GetReusableDelegationSetFunc(param0)
}

func (m *route53Mock) GetReusableDelegationSetRequest(param0 *route53.GetReusableDelegationSetInput) (*request.Request, *route53.GetReusableDelegationSetOutput) {
	m.addCall("GetReusableDelegationSetRequest")
	m.verifyInput("GetReusableDelegationSetRequest", param0)
	return m.GetReusableDelegationSetRequestFunc(param0)
}

func (m *route53Mock) GetReusableDelegationSetWithContext(param0 aws.Context, param1 *route53.GetReusableDelegationSetInput, param2 ...request.Option) (*route53.GetReusableDelegationSetOutput, error) {
	m.addCall("GetReusableDelegationSetWithContext")
	m.verifyInput("GetReusableDelegationSetWithContext", param0)
	return m.GetReusableDelegationSetWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) GetTrafficPolicy(param0 *route53.GetTrafficPolicyInput) (*route53.GetTrafficPolicyOutput, error) {
	m.addCall("GetTrafficPolicy")
	m.verifyInput("GetTrafficPolicy", param0)
	return m.GetTrafficPolicyFunc(param0)
}

func (m *route53Mock) GetTrafficPolicyInstance(param0 *route53.GetTrafficPolicyInstanceInput) (*route53.GetTrafficPolicyInstanceOutput, error) {
	m.addCall("GetTrafficPolicyInstance")
	m.verifyInput("GetTrafficPolicyInstance", param0)
	return m.GetTrafficPolicyInstanceFunc(param0)
}

func (m *route53Mock) GetTrafficPolicyInstanceCount(param0 *route53.GetTrafficPolicyInstanceCountInput) (*route53.GetTrafficPolicyInstanceCountOutput, error) {
	m.addCall("GetTrafficPolicyInstanceCount")
	m.verifyInput("GetTrafficPolicyInstanceCount", param0)
	return m.GetTrafficPolicyInstanceCountFunc(param0)
}

func (m *route53Mock) GetTrafficPolicyInstanceCountRequest(param0 *route53.GetTrafficPolicyInstanceCountInput) (*request.Request, *route53.GetTrafficPolicyInstanceCountOutput) {
	m.addCall("GetTrafficPolicyInstanceCountRequest")
	m.verifyInput("GetTrafficPolicyInstanceCountRequest", param0)
	return m.GetTrafficPolicyInstanceCountRequestFunc(param0)
}

func (m *route53Mock) GetTrafficPolicyInstanceCountWithContext(param0 aws.Context, param1 *route53.GetTrafficPolicyInstanceCountInput, param2 ...request.Option) (*route53.GetTrafficPolicyInstanceCountOutput, error) {
	m.addCall("GetTrafficPolicyInstanceCountWithContext")
	m.verifyInput("GetTrafficPolicyInstanceCountWithContext", param0)
	return m.GetTrafficPolicyInstanceCountWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) GetTrafficPolicyInstanceRequest(param0 *route53.GetTrafficPolicyInstanceInput) (*request.Request, *route53.GetTrafficPolicyInstanceOutput) {
	m.addCall("GetTrafficPolicyInstanceRequest")
	m.verifyInput("GetTrafficPolicyInstanceRequest", param0)
	return m.GetTrafficPolicyInstanceRequestFunc(param0)
}

func (m *route53Mock) GetTrafficPolicyInstanceWithContext(param0 aws.Context, param1 *route53.GetTrafficPolicyInstanceInput, param2 ...request.Option) (*route53.GetTrafficPolicyInstanceOutput, error) {
	m.addCall("GetTrafficPolicyInstanceWithContext")
	m.verifyInput("GetTrafficPolicyInstanceWithContext", param0)
	return m.GetTrafficPolicyInstanceWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) GetTrafficPolicyRequest(param0 *route53.GetTrafficPolicyInput) (*request.Request, *route53.GetTrafficPolicyOutput) {
	m.addCall("GetTrafficPolicyRequest")
	m.verifyInput("GetTrafficPolicyRequest", param0)
	return m.GetTrafficPolicyRequestFunc(param0)
}

func (m *route53Mock) GetTrafficPolicyWithContext(param0 aws.Context, param1 *route53.GetTrafficPolicyInput, param2 ...request.Option) (*route53.GetTrafficPolicyOutput, error) {
	m.addCall("GetTrafficPolicyWithContext")
	m.verifyInput("GetTrafficPolicyWithContext", param0)
	return m.GetTrafficPolicyWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) ListGeoLocations(param0 *route53.ListGeoLocationsInput) (*route53.ListGeoLocationsOutput, error) {
	m.addCall("ListGeoLocations")
	m.verifyInput("ListGeoLocations", param0)
	return m.ListGeoLocationsFunc(param0)
}

func (m *route53Mock) ListGeoLocationsRequest(param0 *route53.ListGeoLocationsInput) (*request.Request, *route53.ListGeoLocationsOutput) {
	m.addCall("ListGeoLocationsRequest")
	m.verifyInput("ListGeoLocationsRequest", param0)
	return m.ListGeoLocationsRequestFunc(param0)
}

func (m *route53Mock) ListGeoLocationsWithContext(param0 aws.Context, param1 *route53.ListGeoLocationsInput, param2 ...request.Option) (*route53.ListGeoLocationsOutput, error) {
	m.addCall("ListGeoLocationsWithContext")
	m.verifyInput("ListGeoLocationsWithContext", param0)
	return m.ListGeoLocationsWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) ListHealthChecks(param0 *route53.ListHealthChecksInput) (*route53.ListHealthChecksOutput, error) {
	m.addCall("ListHealthChecks")
	m.verifyInput("ListHealthChecks", param0)
	return m.ListHealthChecksFunc(param0)
}

func (m *route53Mock) ListHealthChecksRequest(param0 *route53.ListHealthChecksInput) (*request.Request, *route53.ListHealthChecksOutput) {
	m.addCall("ListHealthChecksRequest")
	m.verifyInput("ListHealthChecksRequest", param0)
	return m.ListHealthChecksRequestFunc(param0)
}

func (m *route53Mock) ListHealthChecksWithContext(param0 aws.Context, param1 *route53.ListHealthChecksInput, param2 ...request.Option) (*route53.ListHealthChecksOutput, error) {
	m.addCall("ListHealthChecksWithContext")
	m.verifyInput("ListHealthChecksWithContext", param0)
	return m.ListHealthChecksWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) ListHostedZones(param0 *route53.ListHostedZonesInput) (*route53.ListHostedZonesOutput, error) {
	m.addCall("ListHostedZones")
	m.verifyInput("ListHostedZones", param0)
	return m.ListHostedZonesFunc(param0)
}

func (m *route53Mock) ListHostedZonesByName(param0 *route53.ListHostedZonesByNameInput) (*route53.ListHostedZonesByNameOutput, error) {
	m.addCall("ListHostedZonesByName")
	m.verifyInput("ListHostedZonesByName", param0)
	return m.ListHostedZonesByNameFunc(param0)
}

func (m *route53Mock) ListHostedZonesByNameRequest(param0 *route53.ListHostedZonesByNameInput) (*request.Request, *route53.ListHostedZonesByNameOutput) {
	m.addCall("ListHostedZonesByNameRequest")
	m.verifyInput("ListHostedZonesByNameRequest", param0)
	return m.ListHostedZonesByNameRequestFunc(param0)
}

func (m *route53Mock) ListHostedZonesByNameWithContext(param0 aws.Context, param1 *route53.ListHostedZonesByNameInput, param2 ...request.Option) (*route53.ListHostedZonesByNameOutput, error) {
	m.addCall("ListHostedZonesByNameWithContext")
	m.verifyInput("ListHostedZonesByNameWithContext", param0)
	return m.ListHostedZonesByNameWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) ListHostedZonesRequest(param0 *route53.ListHostedZonesInput) (*request.Request, *route53.ListHostedZonesOutput) {
	m.addCall("ListHostedZonesRequest")
	m.verifyInput("ListHostedZonesRequest", param0)
	return m.ListHostedZonesRequestFunc(param0)
}

func (m *route53Mock) ListHostedZonesWithContext(param0 aws.Context, param1 *route53.ListHostedZonesInput, param2 ...request.Option) (*route53.ListHostedZonesOutput, error) {
	m.addCall("ListHostedZonesWithContext")
	m.verifyInput("ListHostedZonesWithContext", param0)
	return m.ListHostedZonesWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) ListQueryLoggingConfigs(param0 *route53.ListQueryLoggingConfigsInput) (*route53.ListQueryLoggingConfigsOutput, error) {
	m.addCall("ListQueryLoggingConfigs")
	m.verifyInput("ListQueryLoggingConfigs", param0)
	return m.ListQueryLoggingConfigsFunc(param0)
}

func (m *route53Mock) ListQueryLoggingConfigsRequest(param0 *route53.ListQueryLoggingConfigsInput) (*request.Request, *route53.ListQueryLoggingConfigsOutput) {
	m.addCall("ListQueryLoggingConfigsRequest")
	m.verifyInput("ListQueryLoggingConfigsRequest", param0)
	return m.ListQueryLoggingConfigsRequestFunc(param0)
}

func (m *route53Mock) ListQueryLoggingConfigsWithContext(param0 aws.Context, param1 *route53.ListQueryLoggingConfigsInput, param2 ...request.Option) (*route53.ListQueryLoggingConfigsOutput, error) {
	m.addCall("ListQueryLoggingConfigsWithContext")
	m.verifyInput("ListQueryLoggingConfigsWithContext", param0)
	return m.ListQueryLoggingConfigsWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) ListResourceRecordSets(param0 *route53.ListResourceRecordSetsInput) (*route53.ListResourceRecordSetsOutput, error) {
	m.addCall("ListResourceRecordSets")
	m.verifyInput("ListResourceRecordSets", param0)
	return m.ListResourceRecordSetsFunc(param0)
}

func (m *route53Mock) ListResourceRecordSetsRequest(param0 *route53.ListResourceRecordSetsInput) (*request.Request, *route53.ListResourceRecordSetsOutput) {
	m.addCall("ListResourceRecordSetsRequest")
	m.verifyInput("ListResourceRecordSetsRequest", param0)
	return m.ListResourceRecordSetsRequestFunc(param0)
}

func (m *route53Mock) ListResourceRecordSetsWithContext(param0 aws.Context, param1 *route53.ListResourceRecordSetsInput, param2 ...request.Option) (*route53.ListResourceRecordSetsOutput, error) {
	m.addCall("ListResourceRecordSetsWithContext")
	m.verifyInput("ListResourceRecordSetsWithContext", param0)
	return m.ListResourceRecordSetsWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) ListReusableDelegationSets(param0 *route53.ListReusableDelegationSetsInput) (*route53.ListReusableDelegationSetsOutput, error) {
	m.addCall("ListReusableDelegationSets")
	m.verifyInput("ListReusableDelegationSets", param0)
	return m.ListReusableDelegationSetsFunc(param0)
}

func (m *route53Mock) ListReusableDelegationSetsRequest(param0 *route53.ListReusableDelegationSetsInput) (*request.Request, *route53.ListReusableDelegationSetsOutput) {
	m.addCall("ListReusableDelegationSetsRequest")
	m.verifyInput("ListReusableDelegationSetsRequest", param0)
	return m.ListReusableDelegationSetsRequestFunc(param0)
}

func (m *route53Mock) ListReusableDelegationSetsWithContext(param0 aws.Context, param1 *route53.ListReusableDelegationSetsInput, param2 ...request.Option) (*route53.ListReusableDelegationSetsOutput, error) {
	m.addCall("ListReusableDelegationSetsWithContext")
	m.verifyInput("ListReusableDelegationSetsWithContext", param0)
	return m.ListReusableDelegationSetsWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) ListTagsForResource(param0 *route53.ListTagsForResourceInput) (*route53.ListTagsForResourceOutput, error) {
	m.addCall("ListTagsForResource")
	m.verifyInput("ListTagsForResource", param0)
	return m.ListTagsForResourceFunc(param0)
}

func (m *route53Mock) ListTagsForResourceRequest(param0 *route53.ListTagsForResourceInput) (*request.Request, *route53.ListTagsForResourceOutput) {
	m.addCall("ListTagsForResourceRequest")
	m.verifyInput("ListTagsForResourceRequest", param0)
	return m.ListTagsForResourceRequestFunc(param0)
}

func (m *route53Mock) ListTagsForResourceWithContext(param0 aws.Context, param1 *route53.ListTagsForResourceInput, param2 ...request.Option) (*route53.ListTagsForResourceOutput, error) {
	m.addCall("ListTagsForResourceWithContext")
	m.verifyInput("ListTagsForResourceWithContext", param0)
	return m.ListTagsForResourceWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) ListTagsForResources(param0 *route53.ListTagsForResourcesInput) (*route53.ListTagsForResourcesOutput, error) {
	m.addCall("ListTagsForResources")
	m.verifyInput("ListTagsForResources", param0)
	return m.ListTagsForResourcesFunc(param0)
}

func (m *route53Mock) ListTagsForResourcesRequest(param0 *route53.ListTagsForResourcesInput) (*request.Request, *route53.ListTagsForResourcesOutput) {
	m.addCall("ListTagsForResourcesRequest")
	m.verifyInput("ListTagsForResourcesRequest", param0)
	return m.ListTagsForResourcesRequestFunc(param0)
}

func (m *route53Mock) ListTagsForResourcesWithContext(param0 aws.Context, param1 *route53.ListTagsForResourcesInput, param2 ...request.Option) (*route53.ListTagsForResourcesOutput, error) {
	m.addCall("ListTagsForResourcesWithContext")
	m.verifyInput("ListTagsForResourcesWithContext", param0)
	return m.ListTagsForResourcesWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) ListTrafficPolicies(param0 *route53.ListTrafficPoliciesInput) (*route53.ListTrafficPoliciesOutput, error) {
	m.addCall("ListTrafficPolicies")
	m.verifyInput("ListTrafficPolicies", param0)
	return m.ListTrafficPoliciesFunc(param0)
}

func (m *route53Mock) ListTrafficPoliciesRequest(param0 *route53.ListTrafficPoliciesInput) (*request.Request, *route53.ListTrafficPoliciesOutput) {
	m.addCall("ListTrafficPoliciesRequest")
	m.verifyInput("ListTrafficPoliciesRequest", param0)
	return m.ListTrafficPoliciesRequestFunc(param0)
}

func (m *route53Mock) ListTrafficPoliciesWithContext(param0 aws.Context, param1 *route53.ListTrafficPoliciesInput, param2 ...request.Option) (*route53.ListTrafficPoliciesOutput, error) {
	m.addCall("ListTrafficPoliciesWithContext")
	m.verifyInput("ListTrafficPoliciesWithContext", param0)
	return m.ListTrafficPoliciesWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) ListTrafficPolicyInstances(param0 *route53.ListTrafficPolicyInstancesInput) (*route53.ListTrafficPolicyInstancesOutput, error) {
	m.addCall("ListTrafficPolicyInstances")
	m.verifyInput("ListTrafficPolicyInstances", param0)
	return m.ListTrafficPolicyInstancesFunc(param0)
}

func (m *route53Mock) ListTrafficPolicyInstancesByHostedZone(param0 *route53.ListTrafficPolicyInstancesByHostedZoneInput) (*route53.ListTrafficPolicyInstancesByHostedZoneOutput, error) {
	m.addCall("ListTrafficPolicyInstancesByHostedZone")
	m.verifyInput("ListTrafficPolicyInstancesByHostedZone", param0)
	return m.ListTrafficPolicyInstancesByHostedZoneFunc(param0)
}

func (m *route53Mock) ListTrafficPolicyInstancesByHostedZoneRequest(param0 *route53.ListTrafficPolicyInstancesByHostedZoneInput) (*request.Request, *route53.ListTrafficPolicyInstancesByHostedZoneOutput) {
	m.addCall("ListTrafficPolicyInstancesByHostedZoneRequest")
	m.verifyInput("ListTrafficPolicyInstancesByHostedZoneRequest", param0)
	return m.ListTrafficPolicyInstancesByHostedZoneRequestFunc(param0)
}

func (m *route53Mock) ListTrafficPolicyInstancesByHostedZoneWithContext(param0 aws.Context, param1 *route53.ListTrafficPolicyInstancesByHostedZoneInput, param2 ...request.Option) (*route53.ListTrafficPolicyInstancesByHostedZoneOutput, error) {
	m.addCall("ListTrafficPolicyInstancesByHostedZoneWithContext")
	m.verifyInput("ListTrafficPolicyInstancesByHostedZoneWithContext", param0)
	return m.ListTrafficPolicyInstancesByHostedZoneWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) ListTrafficPolicyInstancesByPolicy(param0 *route53.ListTrafficPolicyInstancesByPolicyInput) (*route53.ListTrafficPolicyInstancesByPolicyOutput, error) {
	m.addCall("ListTrafficPolicyInstancesByPolicy")
	m.verifyInput("ListTrafficPolicyInstancesByPolicy", param0)
	return m.ListTrafficPolicyInstancesByPolicyFunc(param0)
}

func (m *route53Mock) ListTrafficPolicyInstancesByPolicyRequest(param0 *route53.ListTrafficPolicyInstancesByPolicyInput) (*request.Request, *route53.ListTrafficPolicyInstancesByPolicyOutput) {
	m.addCall("ListTrafficPolicyInstancesByPolicyRequest")
	m.verifyInput("ListTrafficPolicyInstancesByPolicyRequest", param0)
	return m.ListTrafficPolicyInstancesByPolicyRequestFunc(param0)
}

func (m *route53Mock) ListTrafficPolicyInstancesByPolicyWithContext(param0 aws.Context, param1 *route53.ListTrafficPolicyInstancesByPolicyInput, param2 ...request.Option) (*route53.ListTrafficPolicyInstancesByPolicyOutput, error) {
	m.addCall("ListTrafficPolicyInstancesByPolicyWithContext")
	m.verifyInput("ListTrafficPolicyInstancesByPolicyWithContext", param0)
	return m.ListTrafficPolicyInstancesByPolicyWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) ListTrafficPolicyInstancesRequest(param0 *route53.ListTrafficPolicyInstancesInput) (*request.Request, *route53.ListTrafficPolicyInstancesOutput) {
	m.addCall("ListTrafficPolicyInstancesRequest")
	m.verifyInput("ListTrafficPolicyInstancesRequest", param0)
	return m.ListTrafficPolicyInstancesRequestFunc(param0)
}

func (m *route53Mock) ListTrafficPolicyInstancesWithContext(param0 aws.Context, param1 *route53.ListTrafficPolicyInstancesInput, param2 ...request.Option) (*route53.ListTrafficPolicyInstancesOutput, error) {
	m.addCall("ListTrafficPolicyInstancesWithContext")
	m.verifyInput("ListTrafficPolicyInstancesWithContext", param0)
	return m.ListTrafficPolicyInstancesWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) ListTrafficPolicyVersions(param0 *route53.ListTrafficPolicyVersionsInput) (*route53.ListTrafficPolicyVersionsOutput, error) {
	m.addCall("ListTrafficPolicyVersions")
	m.verifyInput("ListTrafficPolicyVersions", param0)
	return m.ListTrafficPolicyVersionsFunc(param0)
}

func (m *route53Mock) ListTrafficPolicyVersionsRequest(param0 *route53.ListTrafficPolicyVersionsInput) (*request.Request, *route53.ListTrafficPolicyVersionsOutput) {
	m.addCall("ListTrafficPolicyVersionsRequest")
	m.verifyInput("ListTrafficPolicyVersionsRequest", param0)
	return m.ListTrafficPolicyVersionsRequestFunc(param0)
}

func (m *route53Mock) ListTrafficPolicyVersionsWithContext(param0 aws.Context, param1 *route53.ListTrafficPolicyVersionsInput, param2 ...request.Option) (*route53.ListTrafficPolicyVersionsOutput, error) {
	m.addCall("ListTrafficPolicyVersionsWithContext")
	m.verifyInput("ListTrafficPolicyVersionsWithContext", param0)
	return m.ListTrafficPolicyVersionsWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) ListVPCAssociationAuthorizations(param0 *route53.ListVPCAssociationAuthorizationsInput) (*route53.ListVPCAssociationAuthorizationsOutput, error) {
	m.addCall("ListVPCAssociationAuthorizations")
	m.verifyInput("ListVPCAssociationAuthorizations", param0)
	return m.ListVPCAssociationAuthorizationsFunc(param0)
}

func (m *route53Mock) ListVPCAssociationAuthorizationsRequest(param0 *route53.ListVPCAssociationAuthorizationsInput) (*request.Request, *route53.ListVPCAssociationAuthorizationsOutput) {
	m.addCall("ListVPCAssociationAuthorizationsRequest")
	m.verifyInput("ListVPCAssociationAuthorizationsRequest", param0)
	return m.ListVPCAssociationAuthorizationsRequestFunc(param0)
}

func (m *route53Mock) ListVPCAssociationAuthorizationsWithContext(param0 aws.Context, param1 *route53.ListVPCAssociationAuthorizationsInput, param2 ...request.Option) (*route53.ListVPCAssociationAuthorizationsOutput, error) {
	m.addCall("ListVPCAssociationAuthorizationsWithContext")
	m.verifyInput("ListVPCAssociationAuthorizationsWithContext", param0)
	return m.ListVPCAssociationAuthorizationsWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) TestDNSAnswer(param0 *route53.TestDNSAnswerInput) (*route53.TestDNSAnswerOutput, error) {
	m.addCall("TestDNSAnswer")
	m.verifyInput("TestDNSAnswer", param0)
	return m.TestDNSAnswerFunc(param0)
}

func (m *route53Mock) TestDNSAnswerRequest(param0 *route53.TestDNSAnswerInput) (*request.Request, *route53.TestDNSAnswerOutput) {
	m.addCall("TestDNSAnswerRequest")
	m.verifyInput("TestDNSAnswerRequest", param0)
	return m.TestDNSAnswerRequestFunc(param0)
}

func (m *route53Mock) TestDNSAnswerWithContext(param0 aws.Context, param1 *route53.TestDNSAnswerInput, param2 ...request.Option) (*route53.TestDNSAnswerOutput, error) {
	m.addCall("TestDNSAnswerWithContext")
	m.verifyInput("TestDNSAnswerWithContext", param0)
	return m.TestDNSAnswerWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) UpdateHealthCheck(param0 *route53.UpdateHealthCheckInput) (*route53.UpdateHealthCheckOutput, error) {
	m.addCall("UpdateHealthCheck")
	m.verifyInput("UpdateHealthCheck", param0)
	return m.UpdateHealthCheckFunc(param0)
}

func (m *route53Mock) UpdateHealthCheckRequest(param0 *route53.UpdateHealthCheckInput) (*request.Request, *route53.UpdateHealthCheckOutput) {
	m.addCall("UpdateHealthCheckRequest")
	m.verifyInput("UpdateHealthCheckRequest", param0)
	return m.UpdateHealthCheckRequestFunc(param0)
}

func (m *route53Mock) UpdateHealthCheckWithContext(param0 aws.Context, param1 *route53.UpdateHealthCheckInput, param2 ...request.Option) (*route53.UpdateHealthCheckOutput, error) {
	m.addCall("UpdateHealthCheckWithContext")
	m.verifyInput("UpdateHealthCheckWithContext", param0)
	return m.UpdateHealthCheckWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) UpdateHostedZoneComment(param0 *route53.UpdateHostedZoneCommentInput) (*route53.UpdateHostedZoneCommentOutput, error) {
	m.addCall("UpdateHostedZoneComment")
	m.verifyInput("UpdateHostedZoneComment", param0)
	return m.UpdateHostedZoneCommentFunc(param0)
}

func (m *route53Mock) UpdateHostedZoneCommentRequest(param0 *route53.UpdateHostedZoneCommentInput) (*request.Request, *route53.UpdateHostedZoneCommentOutput) {
	m.addCall("UpdateHostedZoneCommentRequest")
	m.verifyInput("UpdateHostedZoneCommentRequest", param0)
	return m.UpdateHostedZoneCommentRequestFunc(param0)
}

func (m *route53Mock) UpdateHostedZoneCommentWithContext(param0 aws.Context, param1 *route53.UpdateHostedZoneCommentInput, param2 ...request.Option) (*route53.UpdateHostedZoneCommentOutput, error) {
	m.addCall("UpdateHostedZoneCommentWithContext")
	m.verifyInput("UpdateHostedZoneCommentWithContext", param0)
	return m.UpdateHostedZoneCommentWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) UpdateTrafficPolicyComment(param0 *route53.UpdateTrafficPolicyCommentInput) (*route53.UpdateTrafficPolicyCommentOutput, error) {
	m.addCall("UpdateTrafficPolicyComment")
	m.verifyInput("UpdateTrafficPolicyComment", param0)
	return m.UpdateTrafficPolicyCommentFunc(param0)
}

func (m *route53Mock) UpdateTrafficPolicyCommentRequest(param0 *route53.UpdateTrafficPolicyCommentInput) (*request.Request, *route53.UpdateTrafficPolicyCommentOutput) {
	m.addCall("UpdateTrafficPolicyCommentRequest")
	m.verifyInput("UpdateTrafficPolicyCommentRequest", param0)
	return m.UpdateTrafficPolicyCommentRequestFunc(param0)
}

func (m *route53Mock) UpdateTrafficPolicyCommentWithContext(param0 aws.Context, param1 *route53.UpdateTrafficPolicyCommentInput, param2 ...request.Option) (*route53.UpdateTrafficPolicyCommentOutput, error) {
	m.addCall("UpdateTrafficPolicyCommentWithContext")
	m.verifyInput("UpdateTrafficPolicyCommentWithContext", param0)
	return m.UpdateTrafficPolicyCommentWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) UpdateTrafficPolicyInstance(param0 *route53.UpdateTrafficPolicyInstanceInput) (*route53.UpdateTrafficPolicyInstanceOutput, error) {
	m.addCall("UpdateTrafficPolicyInstance")
	m.verifyInput("UpdateTrafficPolicyInstance", param0)
	return m.UpdateTrafficPolicyInstanceFunc(param0)
}

func (m *route53Mock) UpdateTrafficPolicyInstanceRequest(param0 *route53.UpdateTrafficPolicyInstanceInput) (*request.Request, *route53.UpdateTrafficPolicyInstanceOutput) {
	m.addCall("UpdateTrafficPolicyInstanceRequest")
	m.verifyInput("UpdateTrafficPolicyInstanceRequest", param0)
	return m.UpdateTrafficPolicyInstanceRequestFunc(param0)
}

func (m *route53Mock) UpdateTrafficPolicyInstanceWithContext(param0 aws.Context, param1 *route53.UpdateTrafficPolicyInstanceInput, param2 ...request.Option) (*route53.UpdateTrafficPolicyInstanceOutput, error) {
	m.addCall("UpdateTrafficPolicyInstanceWithContext")
	m.verifyInput("UpdateTrafficPolicyInstanceWithContext", param0)
	return m.UpdateTrafficPolicyInstanceWithContextFunc(param0, param1, param2...)
}

func (m *route53Mock) WaitUntilResourceRecordSetsChanged(param0 *route53.GetChangeInput) error {
	m.addCall("WaitUntilResourceRecordSetsChanged")
	m.verifyInput("WaitUntilResourceRecordSetsChanged", param0)
	return m.WaitUntilResourceRecordSetsChangedFunc(param0)
}

func (m *route53Mock) WaitUntilResourceRecordSetsChangedWithContext(param0 aws.Context, param1 *route53.GetChangeInput, param2 ...request.WaiterOption) error {
	m.addCall("WaitUntilResourceRecordSetsChangedWithContext")
	m.verifyInput("WaitUntilResourceRecordSetsChangedWithContext", param0)
	return m.WaitUntilResourceRecordSetsChangedWithContextFunc(param0, param1, param2...)
}
