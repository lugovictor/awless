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
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/acm"
	"github.com/aws/aws-sdk-go/service/acm/acmiface"
	"github.com/aws/aws-sdk-go/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go/service/applicationautoscaling/applicationautoscalingiface"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/autoscaling/autoscalingiface"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/aws/aws-sdk-go/service/cloudfront/cloudfrontiface"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/ecs/ecsiface"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/elbv2/elbv2iface"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/lambda/lambdaiface"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/rds/rdsiface"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/aws/aws-sdk-go/service/route53/route53iface"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/wallix/awless/logger"
)

func NewAttachAlarm(sess *session.Session, l ...*logger.Logger) *AttachAlarm {
	cmd := new(AttachAlarm)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = cloudwatch.New(sess)
	}
	return cmd
}

func (cmd *AttachAlarm) SetApi(api cloudwatchiface.CloudWatchAPI) {
	cmd.api = api
}

func (cmd *AttachAlarm) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx)
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *AttachAlarm) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *AttachAlarm) ParamsHelp() string {
	return generateParamsHelp("attachalarm", structListParamsKeys(cmd))
}

func (cmd *AttachAlarm) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewAttachContainertask(sess *session.Session, l ...*logger.Logger) *AttachContainertask {
	cmd := new(AttachContainertask)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ecs.New(sess)
	}
	return cmd
}

func (cmd *AttachContainertask) SetApi(api ecsiface.ECSAPI) {
	cmd.api = api
}

func (cmd *AttachContainertask) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx)
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *AttachContainertask) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *AttachContainertask) ParamsHelp() string {
	return generateParamsHelp("attachcontainertask", structListParamsKeys(cmd))
}

func (cmd *AttachContainertask) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewAttachElasticip(sess *session.Session, l ...*logger.Logger) *AttachElasticip {
	cmd := new(AttachElasticip)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *AttachElasticip) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *AttachElasticip) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.AssociateAddressInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.AssociateAddressInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.AssociateAddress(input)
	cmd.logger.ExtraVerbosef("ec2.AssociateAddress call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *AttachElasticip) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *AttachElasticip) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.AssociateAddressInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.AssociateAddressInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.AssociateAddress(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.AssociateAddress call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: attach elasticip ok")
			return fakeDryRunId("elasticip"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *AttachElasticip) ParamsHelp() string {
	return generateParamsHelp("attachelasticip", structListParamsKeys(cmd))
}

func (cmd *AttachElasticip) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewAttachInstance(sess *session.Session, l ...*logger.Logger) *AttachInstance {
	cmd := new(AttachInstance)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = elbv2.New(sess)
	}
	return cmd
}

func (cmd *AttachInstance) SetApi(api elbv2iface.ELBV2API) {
	cmd.api = api
}

func (cmd *AttachInstance) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &elbv2.RegisterTargetsInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in elbv2.RegisterTargetsInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.RegisterTargets(input)
	cmd.logger.ExtraVerbosef("elbv2.RegisterTargets call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *AttachInstance) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *AttachInstance) ParamsHelp() string {
	return generateParamsHelp("attachinstance", structListParamsKeys(cmd))
}

func (cmd *AttachInstance) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewAttachInternetgateway(sess *session.Session, l ...*logger.Logger) *AttachInternetgateway {
	cmd := new(AttachInternetgateway)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *AttachInternetgateway) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *AttachInternetgateway) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.AttachInternetGatewayInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.AttachInternetGatewayInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.AttachInternetGateway(input)
	cmd.logger.ExtraVerbosef("ec2.AttachInternetGateway call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *AttachInternetgateway) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *AttachInternetgateway) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.AttachInternetGatewayInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.AttachInternetGatewayInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.AttachInternetGateway(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.AttachInternetGateway call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: attach internetgateway ok")
			return fakeDryRunId("internetgateway"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *AttachInternetgateway) ParamsHelp() string {
	return generateParamsHelp("attachinternetgateway", structListParamsKeys(cmd))
}

func (cmd *AttachInternetgateway) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewAttachPolicy(sess *session.Session, l ...*logger.Logger) *AttachPolicy {
	cmd := new(AttachPolicy)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = iam.New(sess)
	}
	return cmd
}

func (cmd *AttachPolicy) SetApi(api iamiface.IAMAPI) {
	cmd.api = api
}

func (cmd *AttachPolicy) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx)
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *AttachPolicy) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *AttachPolicy) ParamsHelp() string {
	return generateParamsHelp("attachpolicy", structListParamsKeys(cmd))
}

func (cmd *AttachPolicy) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewAttachRoutetable(sess *session.Session, l ...*logger.Logger) *AttachRoutetable {
	cmd := new(AttachRoutetable)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *AttachRoutetable) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *AttachRoutetable) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.AssociateRouteTableInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.AssociateRouteTableInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.AssociateRouteTable(input)
	cmd.logger.ExtraVerbosef("ec2.AssociateRouteTable call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *AttachRoutetable) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *AttachRoutetable) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.AssociateRouteTableInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.AssociateRouteTableInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.AssociateRouteTable(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.AssociateRouteTable call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: attach routetable ok")
			return fakeDryRunId("routetable"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *AttachRoutetable) ParamsHelp() string {
	return generateParamsHelp("attachroutetable", structListParamsKeys(cmd))
}

func (cmd *AttachRoutetable) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewAttachSecuritygroup(sess *session.Session, l ...*logger.Logger) *AttachSecuritygroup {
	cmd := new(AttachSecuritygroup)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *AttachSecuritygroup) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *AttachSecuritygroup) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx)
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *AttachSecuritygroup) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *AttachSecuritygroup) ParamsHelp() string {
	return generateParamsHelp("attachsecuritygroup", structListParamsKeys(cmd))
}

func (cmd *AttachSecuritygroup) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewAttachUser(sess *session.Session, l ...*logger.Logger) *AttachUser {
	cmd := new(AttachUser)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = iam.New(sess)
	}
	return cmd
}

func (cmd *AttachUser) SetApi(api iamiface.IAMAPI) {
	cmd.api = api
}

func (cmd *AttachUser) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &iam.AddUserToGroupInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in iam.AddUserToGroupInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.AddUserToGroup(input)
	cmd.logger.ExtraVerbosef("iam.AddUserToGroup call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *AttachUser) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *AttachUser) ParamsHelp() string {
	return generateParamsHelp("attachuser", structListParamsKeys(cmd))
}

func (cmd *AttachUser) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewAttachVolume(sess *session.Session, l ...*logger.Logger) *AttachVolume {
	cmd := new(AttachVolume)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *AttachVolume) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *AttachVolume) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.AttachVolumeInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.AttachVolumeInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.AttachVolume(input)
	cmd.logger.ExtraVerbosef("ec2.AttachVolume call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *AttachVolume) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *AttachVolume) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.AttachVolumeInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.AttachVolumeInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.AttachVolume(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.AttachVolume call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: attach volume ok")
			return fakeDryRunId("volume"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *AttachVolume) ParamsHelp() string {
	return generateParamsHelp("attachvolume", structListParamsKeys(cmd))
}

func (cmd *AttachVolume) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCheckCertificate(sess *session.Session, l ...*logger.Logger) *CheckCertificate {
	cmd := new(CheckCertificate)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = acm.New(sess)
	}
	return cmd
}

func (cmd *CheckCertificate) SetApi(api acmiface.ACMAPI) {
	cmd.api = api
}

func (cmd *CheckCertificate) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx)
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CheckCertificate) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CheckCertificate) ParamsHelp() string {
	return generateParamsHelp("checkcertificate", structListParamsKeys(cmd))
}

func (cmd *CheckCertificate) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCheckDatabase(sess *session.Session, l ...*logger.Logger) *CheckDatabase {
	cmd := new(CheckDatabase)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = rds.New(sess)
	}
	return cmd
}

func (cmd *CheckDatabase) SetApi(api rdsiface.RDSAPI) {
	cmd.api = api
}

func (cmd *CheckDatabase) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx)
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CheckDatabase) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CheckDatabase) ParamsHelp() string {
	return generateParamsHelp("checkdatabase", structListParamsKeys(cmd))
}

func (cmd *CheckDatabase) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCheckDistribution(sess *session.Session, l ...*logger.Logger) *CheckDistribution {
	cmd := new(CheckDistribution)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = cloudfront.New(sess)
	}
	return cmd
}

func (cmd *CheckDistribution) SetApi(api cloudfrontiface.CloudFrontAPI) {
	cmd.api = api
}

func (cmd *CheckDistribution) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx)
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CheckDistribution) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CheckDistribution) ParamsHelp() string {
	return generateParamsHelp("checkdistribution", structListParamsKeys(cmd))
}

func (cmd *CheckDistribution) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCheckInstance(sess *session.Session, l ...*logger.Logger) *CheckInstance {
	cmd := new(CheckInstance)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *CheckInstance) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *CheckInstance) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx)
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CheckInstance) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CheckInstance) ParamsHelp() string {
	return generateParamsHelp("checkinstance", structListParamsKeys(cmd))
}

func (cmd *CheckInstance) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCheckSecuritygroup(sess *session.Session, l ...*logger.Logger) *CheckSecuritygroup {
	cmd := new(CheckSecuritygroup)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *CheckSecuritygroup) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *CheckSecuritygroup) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx)
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CheckSecuritygroup) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CheckSecuritygroup) ParamsHelp() string {
	return generateParamsHelp("checksecuritygroup", structListParamsKeys(cmd))
}

func (cmd *CheckSecuritygroup) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCheckVolume(sess *session.Session, l ...*logger.Logger) *CheckVolume {
	cmd := new(CheckVolume)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *CheckVolume) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *CheckVolume) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx)
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CheckVolume) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CheckVolume) ParamsHelp() string {
	return generateParamsHelp("checkvolume", structListParamsKeys(cmd))
}

func (cmd *CheckVolume) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCopyImage(sess *session.Session, l ...*logger.Logger) *CopyImage {
	cmd := new(CopyImage)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *CopyImage) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *CopyImage) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.CopyImageInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.CopyImageInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.CopyImage(input)
	cmd.logger.ExtraVerbosef("ec2.CopyImage call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CopyImage) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CopyImage) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.CopyImageInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.CopyImageInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.CopyImage(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.CopyImage call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: copy image ok")
			return fakeDryRunId("image"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *CopyImage) ParamsHelp() string {
	return generateParamsHelp("copyimage", structListParamsKeys(cmd))
}

func (cmd *CopyImage) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCopySnapshot(sess *session.Session, l ...*logger.Logger) *CopySnapshot {
	cmd := new(CopySnapshot)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *CopySnapshot) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *CopySnapshot) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.CopySnapshotInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.CopySnapshotInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.CopySnapshot(input)
	cmd.logger.ExtraVerbosef("ec2.CopySnapshot call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CopySnapshot) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CopySnapshot) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.CopySnapshotInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.CopySnapshotInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.CopySnapshot(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.CopySnapshot call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: copy snapshot ok")
			return fakeDryRunId("snapshot"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *CopySnapshot) ParamsHelp() string {
	return generateParamsHelp("copysnapshot", structListParamsKeys(cmd))
}

func (cmd *CopySnapshot) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateAccesskey(sess *session.Session, l ...*logger.Logger) *CreateAccesskey {
	cmd := new(CreateAccesskey)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = iam.New(sess)
	}
	return cmd
}

func (cmd *CreateAccesskey) SetApi(api iamiface.IAMAPI) {
	cmd.api = api
}

func (cmd *CreateAccesskey) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &iam.CreateAccessKeyInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in iam.CreateAccessKeyInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.CreateAccessKey(input)
	cmd.logger.ExtraVerbosef("iam.CreateAccessKey call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreateAccesskey) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreateAccesskey) ParamsHelp() string {
	return generateParamsHelp("createaccesskey", structListParamsKeys(cmd))
}

func (cmd *CreateAccesskey) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateAlarm(sess *session.Session, l ...*logger.Logger) *CreateAlarm {
	cmd := new(CreateAlarm)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = cloudwatch.New(sess)
	}
	return cmd
}

func (cmd *CreateAlarm) SetApi(api cloudwatchiface.CloudWatchAPI) {
	cmd.api = api
}

func (cmd *CreateAlarm) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &cloudwatch.PutMetricAlarmInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in cloudwatch.PutMetricAlarmInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.PutMetricAlarm(input)
	cmd.logger.ExtraVerbosef("cloudwatch.PutMetricAlarm call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreateAlarm) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreateAlarm) ParamsHelp() string {
	return generateParamsHelp("createalarm", structListParamsKeys(cmd))
}

func (cmd *CreateAlarm) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateAppscalingpolicy(sess *session.Session, l ...*logger.Logger) *CreateAppscalingpolicy {
	cmd := new(CreateAppscalingpolicy)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = applicationautoscaling.New(sess)
	}
	return cmd
}

func (cmd *CreateAppscalingpolicy) SetApi(api applicationautoscalingiface.ApplicationAutoScalingAPI) {
	cmd.api = api
}

func (cmd *CreateAppscalingpolicy) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &applicationautoscaling.PutScalingPolicyInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in applicationautoscaling.PutScalingPolicyInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.PutScalingPolicy(input)
	cmd.logger.ExtraVerbosef("applicationautoscaling.PutScalingPolicy call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreateAppscalingpolicy) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreateAppscalingpolicy) ParamsHelp() string {
	return generateParamsHelp("createappscalingpolicy", structListParamsKeys(cmd))
}

func (cmd *CreateAppscalingpolicy) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateAppscalingtarget(sess *session.Session, l ...*logger.Logger) *CreateAppscalingtarget {
	cmd := new(CreateAppscalingtarget)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = applicationautoscaling.New(sess)
	}
	return cmd
}

func (cmd *CreateAppscalingtarget) SetApi(api applicationautoscalingiface.ApplicationAutoScalingAPI) {
	cmd.api = api
}

func (cmd *CreateAppscalingtarget) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &applicationautoscaling.RegisterScalableTargetInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in applicationautoscaling.RegisterScalableTargetInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.RegisterScalableTarget(input)
	cmd.logger.ExtraVerbosef("applicationautoscaling.RegisterScalableTarget call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreateAppscalingtarget) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreateAppscalingtarget) ParamsHelp() string {
	return generateParamsHelp("createappscalingtarget", structListParamsKeys(cmd))
}

func (cmd *CreateAppscalingtarget) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateBucket(sess *session.Session, l ...*logger.Logger) *CreateBucket {
	cmd := new(CreateBucket)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = s3.New(sess)
	}
	return cmd
}

func (cmd *CreateBucket) SetApi(api s3iface.S3API) {
	cmd.api = api
}

func (cmd *CreateBucket) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &s3.CreateBucketInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in s3.CreateBucketInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.CreateBucket(input)
	cmd.logger.ExtraVerbosef("s3.CreateBucket call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreateBucket) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreateBucket) ParamsHelp() string {
	return generateParamsHelp("createbucket", structListParamsKeys(cmd))
}

func (cmd *CreateBucket) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateCertificate(sess *session.Session, l ...*logger.Logger) *CreateCertificate {
	cmd := new(CreateCertificate)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = acm.New(sess)
	}
	return cmd
}

func (cmd *CreateCertificate) SetApi(api acmiface.ACMAPI) {
	cmd.api = api
}

func (cmd *CreateCertificate) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx)
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreateCertificate) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreateCertificate) ParamsHelp() string {
	return generateParamsHelp("createcertificate", structListParamsKeys(cmd))
}

func (cmd *CreateCertificate) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateContainercluster(sess *session.Session, l ...*logger.Logger) *CreateContainercluster {
	cmd := new(CreateContainercluster)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ecs.New(sess)
	}
	return cmd
}

func (cmd *CreateContainercluster) SetApi(api ecsiface.ECSAPI) {
	cmd.api = api
}

func (cmd *CreateContainercluster) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ecs.CreateClusterInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ecs.CreateClusterInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.CreateCluster(input)
	cmd.logger.ExtraVerbosef("ecs.CreateCluster call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreateContainercluster) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreateContainercluster) ParamsHelp() string {
	return generateParamsHelp("createcontainercluster", structListParamsKeys(cmd))
}

func (cmd *CreateContainercluster) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateDatabase(sess *session.Session, l ...*logger.Logger) *CreateDatabase {
	cmd := new(CreateDatabase)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = rds.New(sess)
	}
	return cmd
}

func (cmd *CreateDatabase) SetApi(api rdsiface.RDSAPI) {
	cmd.api = api
}

func (cmd *CreateDatabase) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &rds.CreateDBInstanceInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in rds.CreateDBInstanceInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.CreateDBInstance(input)
	cmd.logger.ExtraVerbosef("rds.CreateDBInstance call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreateDatabase) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreateDatabase) ParamsHelp() string {
	return generateParamsHelp("createdatabase", structListParamsKeys(cmd))
}

func (cmd *CreateDatabase) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateDbsubnetgroup(sess *session.Session, l ...*logger.Logger) *CreateDbsubnetgroup {
	cmd := new(CreateDbsubnetgroup)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = rds.New(sess)
	}
	return cmd
}

func (cmd *CreateDbsubnetgroup) SetApi(api rdsiface.RDSAPI) {
	cmd.api = api
}

func (cmd *CreateDbsubnetgroup) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &rds.CreateDBSubnetGroupInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in rds.CreateDBSubnetGroupInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.CreateDBSubnetGroup(input)
	cmd.logger.ExtraVerbosef("rds.CreateDBSubnetGroup call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreateDbsubnetgroup) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreateDbsubnetgroup) ParamsHelp() string {
	return generateParamsHelp("createdbsubnetgroup", structListParamsKeys(cmd))
}

func (cmd *CreateDbsubnetgroup) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateDistribution(sess *session.Session, l ...*logger.Logger) *CreateDistribution {
	cmd := new(CreateDistribution)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = cloudfront.New(sess)
	}
	return cmd
}

func (cmd *CreateDistribution) SetApi(api cloudfrontiface.CloudFrontAPI) {
	cmd.api = api
}

func (cmd *CreateDistribution) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx)
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreateDistribution) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreateDistribution) ParamsHelp() string {
	return generateParamsHelp("createdistribution", structListParamsKeys(cmd))
}

func (cmd *CreateDistribution) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateElasticip(sess *session.Session, l ...*logger.Logger) *CreateElasticip {
	cmd := new(CreateElasticip)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *CreateElasticip) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *CreateElasticip) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.AllocateAddressInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.AllocateAddressInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.AllocateAddress(input)
	cmd.logger.ExtraVerbosef("ec2.AllocateAddress call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreateElasticip) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreateElasticip) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.AllocateAddressInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.AllocateAddressInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.AllocateAddress(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.AllocateAddress call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: create elasticip ok")
			return fakeDryRunId("elasticip"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *CreateElasticip) ParamsHelp() string {
	return generateParamsHelp("createelasticip", structListParamsKeys(cmd))
}

func (cmd *CreateElasticip) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateFunction(sess *session.Session, l ...*logger.Logger) *CreateFunction {
	cmd := new(CreateFunction)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = lambda.New(sess)
	}
	return cmd
}

func (cmd *CreateFunction) SetApi(api lambdaiface.LambdaAPI) {
	cmd.api = api
}

func (cmd *CreateFunction) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &lambda.CreateFunctionInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in lambda.CreateFunctionInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.CreateFunction(input)
	cmd.logger.ExtraVerbosef("lambda.CreateFunction call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreateFunction) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreateFunction) ParamsHelp() string {
	return generateParamsHelp("createfunction", structListParamsKeys(cmd))
}

func (cmd *CreateFunction) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateGroup(sess *session.Session, l ...*logger.Logger) *CreateGroup {
	cmd := new(CreateGroup)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = iam.New(sess)
	}
	return cmd
}

func (cmd *CreateGroup) SetApi(api iamiface.IAMAPI) {
	cmd.api = api
}

func (cmd *CreateGroup) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &iam.CreateGroupInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in iam.CreateGroupInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.CreateGroup(input)
	cmd.logger.ExtraVerbosef("iam.CreateGroup call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreateGroup) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreateGroup) ParamsHelp() string {
	return generateParamsHelp("creategroup", structListParamsKeys(cmd))
}

func (cmd *CreateGroup) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateInstance(sess *session.Session, l ...*logger.Logger) *CreateInstance {
	cmd := new(CreateInstance)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *CreateInstance) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *CreateInstance) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.RunInstancesInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.RunInstancesInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.RunInstances(input)
	cmd.logger.ExtraVerbosef("ec2.RunInstances call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreateInstance) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreateInstance) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.RunInstancesInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.RunInstancesInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.RunInstances(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.RunInstances call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: create instance ok")
			return fakeDryRunId("instance"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *CreateInstance) ParamsHelp() string {
	return generateParamsHelp("createinstance", structListParamsKeys(cmd))
}

func (cmd *CreateInstance) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateInternetgateway(sess *session.Session, l ...*logger.Logger) *CreateInternetgateway {
	cmd := new(CreateInternetgateway)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *CreateInternetgateway) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *CreateInternetgateway) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.CreateInternetGatewayInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.CreateInternetGatewayInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.CreateInternetGateway(input)
	cmd.logger.ExtraVerbosef("ec2.CreateInternetGateway call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreateInternetgateway) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreateInternetgateway) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.CreateInternetGatewayInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.CreateInternetGatewayInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.CreateInternetGateway(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.CreateInternetGateway call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: create internetgateway ok")
			return fakeDryRunId("internetgateway"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *CreateInternetgateway) ParamsHelp() string {
	return generateParamsHelp("createinternetgateway", structListParamsKeys(cmd))
}

func (cmd *CreateInternetgateway) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateKeypair(sess *session.Session, l ...*logger.Logger) *CreateKeypair {
	cmd := new(CreateKeypair)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *CreateKeypair) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *CreateKeypair) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.ImportKeyPairInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.ImportKeyPairInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.ImportKeyPair(input)
	cmd.logger.ExtraVerbosef("ec2.ImportKeyPair call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreateKeypair) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreateKeypair) ParamsHelp() string {
	return generateParamsHelp("createkeypair", structListParamsKeys(cmd))
}

func (cmd *CreateKeypair) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreatePolicy(sess *session.Session, l ...*logger.Logger) *CreatePolicy {
	cmd := new(CreatePolicy)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = iam.New(sess)
	}
	return cmd
}

func (cmd *CreatePolicy) SetApi(api iamiface.IAMAPI) {
	cmd.api = api
}

func (cmd *CreatePolicy) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &iam.CreatePolicyInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in iam.CreatePolicyInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.CreatePolicy(input)
	cmd.logger.ExtraVerbosef("iam.CreatePolicy call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreatePolicy) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreatePolicy) ParamsHelp() string {
	return generateParamsHelp("createpolicy", structListParamsKeys(cmd))
}

func (cmd *CreatePolicy) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateRoute(sess *session.Session, l ...*logger.Logger) *CreateRoute {
	cmd := new(CreateRoute)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *CreateRoute) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *CreateRoute) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.CreateRouteInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.CreateRouteInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.CreateRoute(input)
	cmd.logger.ExtraVerbosef("ec2.CreateRoute call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreateRoute) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreateRoute) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.CreateRouteInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.CreateRouteInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.CreateRoute(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.CreateRoute call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: create route ok")
			return fakeDryRunId("route"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *CreateRoute) ParamsHelp() string {
	return generateParamsHelp("createroute", structListParamsKeys(cmd))
}

func (cmd *CreateRoute) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateRoutetable(sess *session.Session, l ...*logger.Logger) *CreateRoutetable {
	cmd := new(CreateRoutetable)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *CreateRoutetable) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *CreateRoutetable) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.CreateRouteTableInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.CreateRouteTableInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.CreateRouteTable(input)
	cmd.logger.ExtraVerbosef("ec2.CreateRouteTable call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreateRoutetable) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreateRoutetable) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.CreateRouteTableInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.CreateRouteTableInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.CreateRouteTable(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.CreateRouteTable call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: create routetable ok")
			return fakeDryRunId("routetable"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *CreateRoutetable) ParamsHelp() string {
	return generateParamsHelp("createroutetable", structListParamsKeys(cmd))
}

func (cmd *CreateRoutetable) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateScalingpolicy(sess *session.Session, l ...*logger.Logger) *CreateScalingpolicy {
	cmd := new(CreateScalingpolicy)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = autoscaling.New(sess)
	}
	return cmd
}

func (cmd *CreateScalingpolicy) SetApi(api autoscalingiface.AutoScalingAPI) {
	cmd.api = api
}

func (cmd *CreateScalingpolicy) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &autoscaling.PutScalingPolicyInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in autoscaling.PutScalingPolicyInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.PutScalingPolicy(input)
	cmd.logger.ExtraVerbosef("autoscaling.PutScalingPolicy call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreateScalingpolicy) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreateScalingpolicy) ParamsHelp() string {
	return generateParamsHelp("createscalingpolicy", structListParamsKeys(cmd))
}

func (cmd *CreateScalingpolicy) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateSecuritygroup(sess *session.Session, l ...*logger.Logger) *CreateSecuritygroup {
	cmd := new(CreateSecuritygroup)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *CreateSecuritygroup) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *CreateSecuritygroup) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.CreateSecurityGroupInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.CreateSecurityGroupInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.CreateSecurityGroup(input)
	cmd.logger.ExtraVerbosef("ec2.CreateSecurityGroup call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreateSecuritygroup) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreateSecuritygroup) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.CreateSecurityGroupInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.CreateSecurityGroupInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.CreateSecurityGroup(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.CreateSecurityGroup call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: create securitygroup ok")
			return fakeDryRunId("securitygroup"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *CreateSecuritygroup) ParamsHelp() string {
	return generateParamsHelp("createsecuritygroup", structListParamsKeys(cmd))
}

func (cmd *CreateSecuritygroup) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateSnapshot(sess *session.Session, l ...*logger.Logger) *CreateSnapshot {
	cmd := new(CreateSnapshot)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *CreateSnapshot) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *CreateSnapshot) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.CreateSnapshotInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.CreateSnapshotInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.CreateSnapshot(input)
	cmd.logger.ExtraVerbosef("ec2.CreateSnapshot call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreateSnapshot) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreateSnapshot) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.CreateSnapshotInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.CreateSnapshotInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.CreateSnapshot(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.CreateSnapshot call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: create snapshot ok")
			return fakeDryRunId("snapshot"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *CreateSnapshot) ParamsHelp() string {
	return generateParamsHelp("createsnapshot", structListParamsKeys(cmd))
}

func (cmd *CreateSnapshot) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateStack(sess *session.Session, l ...*logger.Logger) *CreateStack {
	cmd := new(CreateStack)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = cloudformation.New(sess)
	}
	return cmd
}

func (cmd *CreateStack) SetApi(api cloudformationiface.CloudFormationAPI) {
	cmd.api = api
}

func (cmd *CreateStack) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &cloudformation.CreateStackInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in cloudformation.CreateStackInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.CreateStack(input)
	cmd.logger.ExtraVerbosef("cloudformation.CreateStack call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreateStack) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreateStack) ParamsHelp() string {
	return generateParamsHelp("createstack", structListParamsKeys(cmd))
}

func (cmd *CreateStack) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateSubnet(sess *session.Session, l ...*logger.Logger) *CreateSubnet {
	cmd := new(CreateSubnet)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *CreateSubnet) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *CreateSubnet) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.CreateSubnetInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.CreateSubnetInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.CreateSubnet(input)
	cmd.logger.ExtraVerbosef("ec2.CreateSubnet call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreateSubnet) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreateSubnet) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.CreateSubnetInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.CreateSubnetInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.CreateSubnet(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.CreateSubnet call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: create subnet ok")
			return fakeDryRunId("subnet"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *CreateSubnet) ParamsHelp() string {
	return generateParamsHelp("createsubnet", structListParamsKeys(cmd))
}

func (cmd *CreateSubnet) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateSubscription(sess *session.Session, l ...*logger.Logger) *CreateSubscription {
	cmd := new(CreateSubscription)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = sns.New(sess)
	}
	return cmd
}

func (cmd *CreateSubscription) SetApi(api snsiface.SNSAPI) {
	cmd.api = api
}

func (cmd *CreateSubscription) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &sns.SubscribeInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in sns.SubscribeInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.Subscribe(input)
	cmd.logger.ExtraVerbosef("sns.Subscribe call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreateSubscription) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreateSubscription) ParamsHelp() string {
	return generateParamsHelp("createsubscription", structListParamsKeys(cmd))
}

func (cmd *CreateSubscription) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateTag(sess *session.Session, l ...*logger.Logger) *CreateTag {
	cmd := new(CreateTag)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *CreateTag) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *CreateTag) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx)
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreateTag) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreateTag) ParamsHelp() string {
	return generateParamsHelp("createtag", structListParamsKeys(cmd))
}

func (cmd *CreateTag) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateTargetgroup(sess *session.Session, l ...*logger.Logger) *CreateTargetgroup {
	cmd := new(CreateTargetgroup)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = elbv2.New(sess)
	}
	return cmd
}

func (cmd *CreateTargetgroup) SetApi(api elbv2iface.ELBV2API) {
	cmd.api = api
}

func (cmd *CreateTargetgroup) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &elbv2.CreateTargetGroupInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in elbv2.CreateTargetGroupInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.CreateTargetGroup(input)
	cmd.logger.ExtraVerbosef("elbv2.CreateTargetGroup call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreateTargetgroup) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreateTargetgroup) ParamsHelp() string {
	return generateParamsHelp("createtargetgroup", structListParamsKeys(cmd))
}

func (cmd *CreateTargetgroup) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateTopic(sess *session.Session, l ...*logger.Logger) *CreateTopic {
	cmd := new(CreateTopic)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = sns.New(sess)
	}
	return cmd
}

func (cmd *CreateTopic) SetApi(api snsiface.SNSAPI) {
	cmd.api = api
}

func (cmd *CreateTopic) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &sns.CreateTopicInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in sns.CreateTopicInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.CreateTopic(input)
	cmd.logger.ExtraVerbosef("sns.CreateTopic call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreateTopic) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreateTopic) ParamsHelp() string {
	return generateParamsHelp("createtopic", structListParamsKeys(cmd))
}

func (cmd *CreateTopic) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateUser(sess *session.Session, l ...*logger.Logger) *CreateUser {
	cmd := new(CreateUser)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = iam.New(sess)
	}
	return cmd
}

func (cmd *CreateUser) SetApi(api iamiface.IAMAPI) {
	cmd.api = api
}

func (cmd *CreateUser) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &iam.CreateUserInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in iam.CreateUserInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.CreateUser(input)
	cmd.logger.ExtraVerbosef("iam.CreateUser call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreateUser) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreateUser) ParamsHelp() string {
	return generateParamsHelp("createuser", structListParamsKeys(cmd))
}

func (cmd *CreateUser) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateVolume(sess *session.Session, l ...*logger.Logger) *CreateVolume {
	cmd := new(CreateVolume)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *CreateVolume) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *CreateVolume) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.CreateVolumeInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.CreateVolumeInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.CreateVolume(input)
	cmd.logger.ExtraVerbosef("ec2.CreateVolume call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreateVolume) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreateVolume) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.CreateVolumeInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.CreateVolumeInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.CreateVolume(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.CreateVolume call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: create volume ok")
			return fakeDryRunId("volume"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *CreateVolume) ParamsHelp() string {
	return generateParamsHelp("createvolume", structListParamsKeys(cmd))
}

func (cmd *CreateVolume) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateVpc(sess *session.Session, l ...*logger.Logger) *CreateVpc {
	cmd := new(CreateVpc)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *CreateVpc) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *CreateVpc) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.CreateVpcInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.CreateVpcInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.CreateVpc(input)
	cmd.logger.ExtraVerbosef("ec2.CreateVpc call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreateVpc) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreateVpc) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.CreateVpcInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.CreateVpcInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.CreateVpc(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.CreateVpc call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: create vpc ok")
			return fakeDryRunId("vpc"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *CreateVpc) ParamsHelp() string {
	return generateParamsHelp("createvpc", structListParamsKeys(cmd))
}

func (cmd *CreateVpc) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateZone(sess *session.Session, l ...*logger.Logger) *CreateZone {
	cmd := new(CreateZone)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = route53.New(sess)
	}
	return cmd
}

func (cmd *CreateZone) SetApi(api route53iface.Route53API) {
	cmd.api = api
}

func (cmd *CreateZone) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &route53.CreateHostedZoneInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in route53.CreateHostedZoneInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.CreateHostedZone(input)
	cmd.logger.ExtraVerbosef("route53.CreateHostedZone call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *CreateZone) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *CreateZone) ParamsHelp() string {
	return generateParamsHelp("createzone", structListParamsKeys(cmd))
}

func (cmd *CreateZone) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteAccesskey(sess *session.Session, l ...*logger.Logger) *DeleteAccesskey {
	cmd := new(DeleteAccesskey)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = iam.New(sess)
	}
	return cmd
}

func (cmd *DeleteAccesskey) SetApi(api iamiface.IAMAPI) {
	cmd.api = api
}

func (cmd *DeleteAccesskey) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &iam.DeleteAccessKeyInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in iam.DeleteAccessKeyInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DeleteAccessKey(input)
	cmd.logger.ExtraVerbosef("iam.DeleteAccessKey call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteAccesskey) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteAccesskey) ParamsHelp() string {
	return generateParamsHelp("deleteaccesskey", structListParamsKeys(cmd))
}

func (cmd *DeleteAccesskey) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteAlarm(sess *session.Session, l ...*logger.Logger) *DeleteAlarm {
	cmd := new(DeleteAlarm)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = cloudwatch.New(sess)
	}
	return cmd
}

func (cmd *DeleteAlarm) SetApi(api cloudwatchiface.CloudWatchAPI) {
	cmd.api = api
}

func (cmd *DeleteAlarm) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &cloudwatch.DeleteAlarmsInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in cloudwatch.DeleteAlarmsInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DeleteAlarms(input)
	cmd.logger.ExtraVerbosef("cloudwatch.DeleteAlarms call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteAlarm) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteAlarm) ParamsHelp() string {
	return generateParamsHelp("deletealarm", structListParamsKeys(cmd))
}

func (cmd *DeleteAlarm) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteAppscalingpolicy(sess *session.Session, l ...*logger.Logger) *DeleteAppscalingpolicy {
	cmd := new(DeleteAppscalingpolicy)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = applicationautoscaling.New(sess)
	}
	return cmd
}

func (cmd *DeleteAppscalingpolicy) SetApi(api applicationautoscalingiface.ApplicationAutoScalingAPI) {
	cmd.api = api
}

func (cmd *DeleteAppscalingpolicy) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &applicationautoscaling.DeleteScalingPolicyInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in applicationautoscaling.DeleteScalingPolicyInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DeleteScalingPolicy(input)
	cmd.logger.ExtraVerbosef("applicationautoscaling.DeleteScalingPolicy call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteAppscalingpolicy) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteAppscalingpolicy) ParamsHelp() string {
	return generateParamsHelp("deleteappscalingpolicy", structListParamsKeys(cmd))
}

func (cmd *DeleteAppscalingpolicy) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteAppscalingtarget(sess *session.Session, l ...*logger.Logger) *DeleteAppscalingtarget {
	cmd := new(DeleteAppscalingtarget)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = applicationautoscaling.New(sess)
	}
	return cmd
}

func (cmd *DeleteAppscalingtarget) SetApi(api applicationautoscalingiface.ApplicationAutoScalingAPI) {
	cmd.api = api
}

func (cmd *DeleteAppscalingtarget) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &applicationautoscaling.DeregisterScalableTargetInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in applicationautoscaling.DeregisterScalableTargetInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DeregisterScalableTarget(input)
	cmd.logger.ExtraVerbosef("applicationautoscaling.DeregisterScalableTarget call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteAppscalingtarget) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteAppscalingtarget) ParamsHelp() string {
	return generateParamsHelp("deleteappscalingtarget", structListParamsKeys(cmd))
}

func (cmd *DeleteAppscalingtarget) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteBucket(sess *session.Session, l ...*logger.Logger) *DeleteBucket {
	cmd := new(DeleteBucket)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = s3.New(sess)
	}
	return cmd
}

func (cmd *DeleteBucket) SetApi(api s3iface.S3API) {
	cmd.api = api
}

func (cmd *DeleteBucket) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &s3.DeleteBucketInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in s3.DeleteBucketInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DeleteBucket(input)
	cmd.logger.ExtraVerbosef("s3.DeleteBucket call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteBucket) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteBucket) ParamsHelp() string {
	return generateParamsHelp("deletebucket", structListParamsKeys(cmd))
}

func (cmd *DeleteBucket) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteCertificate(sess *session.Session, l ...*logger.Logger) *DeleteCertificate {
	cmd := new(DeleteCertificate)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = acm.New(sess)
	}
	return cmd
}

func (cmd *DeleteCertificate) SetApi(api acmiface.ACMAPI) {
	cmd.api = api
}

func (cmd *DeleteCertificate) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &acm.DeleteCertificateInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in acm.DeleteCertificateInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DeleteCertificate(input)
	cmd.logger.ExtraVerbosef("acm.DeleteCertificate call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteCertificate) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteCertificate) ParamsHelp() string {
	return generateParamsHelp("deletecertificate", structListParamsKeys(cmd))
}

func (cmd *DeleteCertificate) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteContainercluster(sess *session.Session, l ...*logger.Logger) *DeleteContainercluster {
	cmd := new(DeleteContainercluster)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ecs.New(sess)
	}
	return cmd
}

func (cmd *DeleteContainercluster) SetApi(api ecsiface.ECSAPI) {
	cmd.api = api
}

func (cmd *DeleteContainercluster) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ecs.DeleteClusterInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ecs.DeleteClusterInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DeleteCluster(input)
	cmd.logger.ExtraVerbosef("ecs.DeleteCluster call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteContainercluster) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteContainercluster) ParamsHelp() string {
	return generateParamsHelp("deletecontainercluster", structListParamsKeys(cmd))
}

func (cmd *DeleteContainercluster) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteContainertask(sess *session.Session, l ...*logger.Logger) *DeleteContainertask {
	cmd := new(DeleteContainertask)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ecs.New(sess)
	}
	return cmd
}

func (cmd *DeleteContainertask) SetApi(api ecsiface.ECSAPI) {
	cmd.api = api
}

func (cmd *DeleteContainertask) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx)
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteContainertask) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteContainertask) ParamsHelp() string {
	return generateParamsHelp("deletecontainertask", structListParamsKeys(cmd))
}

func (cmd *DeleteContainertask) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteDatabase(sess *session.Session, l ...*logger.Logger) *DeleteDatabase {
	cmd := new(DeleteDatabase)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = rds.New(sess)
	}
	return cmd
}

func (cmd *DeleteDatabase) SetApi(api rdsiface.RDSAPI) {
	cmd.api = api
}

func (cmd *DeleteDatabase) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &rds.DeleteDBInstanceInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in rds.DeleteDBInstanceInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DeleteDBInstance(input)
	cmd.logger.ExtraVerbosef("rds.DeleteDBInstance call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteDatabase) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteDatabase) ParamsHelp() string {
	return generateParamsHelp("deletedatabase", structListParamsKeys(cmd))
}

func (cmd *DeleteDatabase) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteDbsubnetgroup(sess *session.Session, l ...*logger.Logger) *DeleteDbsubnetgroup {
	cmd := new(DeleteDbsubnetgroup)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = rds.New(sess)
	}
	return cmd
}

func (cmd *DeleteDbsubnetgroup) SetApi(api rdsiface.RDSAPI) {
	cmd.api = api
}

func (cmd *DeleteDbsubnetgroup) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &rds.DeleteDBSubnetGroupInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in rds.DeleteDBSubnetGroupInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DeleteDBSubnetGroup(input)
	cmd.logger.ExtraVerbosef("rds.DeleteDBSubnetGroup call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteDbsubnetgroup) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteDbsubnetgroup) ParamsHelp() string {
	return generateParamsHelp("deletedbsubnetgroup", structListParamsKeys(cmd))
}

func (cmd *DeleteDbsubnetgroup) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteDistribution(sess *session.Session, l ...*logger.Logger) *DeleteDistribution {
	cmd := new(DeleteDistribution)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = cloudfront.New(sess)
	}
	return cmd
}

func (cmd *DeleteDistribution) SetApi(api cloudfrontiface.CloudFrontAPI) {
	cmd.api = api
}

func (cmd *DeleteDistribution) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx)
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteDistribution) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteDistribution) ParamsHelp() string {
	return generateParamsHelp("deletedistribution", structListParamsKeys(cmd))
}

func (cmd *DeleteDistribution) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteElasticip(sess *session.Session, l ...*logger.Logger) *DeleteElasticip {
	cmd := new(DeleteElasticip)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *DeleteElasticip) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *DeleteElasticip) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.ReleaseAddressInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.ReleaseAddressInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.ReleaseAddress(input)
	cmd.logger.ExtraVerbosef("ec2.ReleaseAddress call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteElasticip) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteElasticip) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.ReleaseAddressInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.ReleaseAddressInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.ReleaseAddress(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.ReleaseAddress call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: delete elasticip ok")
			return fakeDryRunId("elasticip"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *DeleteElasticip) ParamsHelp() string {
	return generateParamsHelp("deleteelasticip", structListParamsKeys(cmd))
}

func (cmd *DeleteElasticip) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteFunction(sess *session.Session, l ...*logger.Logger) *DeleteFunction {
	cmd := new(DeleteFunction)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = lambda.New(sess)
	}
	return cmd
}

func (cmd *DeleteFunction) SetApi(api lambdaiface.LambdaAPI) {
	cmd.api = api
}

func (cmd *DeleteFunction) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &lambda.DeleteFunctionInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in lambda.DeleteFunctionInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DeleteFunction(input)
	cmd.logger.ExtraVerbosef("lambda.DeleteFunction call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteFunction) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteFunction) ParamsHelp() string {
	return generateParamsHelp("deletefunction", structListParamsKeys(cmd))
}

func (cmd *DeleteFunction) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteGroup(sess *session.Session, l ...*logger.Logger) *DeleteGroup {
	cmd := new(DeleteGroup)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = iam.New(sess)
	}
	return cmd
}

func (cmd *DeleteGroup) SetApi(api iamiface.IAMAPI) {
	cmd.api = api
}

func (cmd *DeleteGroup) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &iam.DeleteGroupInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in iam.DeleteGroupInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DeleteGroup(input)
	cmd.logger.ExtraVerbosef("iam.DeleteGroup call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteGroup) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteGroup) ParamsHelp() string {
	return generateParamsHelp("deletegroup", structListParamsKeys(cmd))
}

func (cmd *DeleteGroup) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteImage(sess *session.Session, l ...*logger.Logger) *DeleteImage {
	cmd := new(DeleteImage)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *DeleteImage) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *DeleteImage) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx)
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteImage) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteImage) ParamsHelp() string {
	return generateParamsHelp("deleteimage", structListParamsKeys(cmd))
}

func (cmd *DeleteImage) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteInstance(sess *session.Session, l ...*logger.Logger) *DeleteInstance {
	cmd := new(DeleteInstance)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *DeleteInstance) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *DeleteInstance) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.TerminateInstancesInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.TerminateInstancesInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.TerminateInstances(input)
	cmd.logger.ExtraVerbosef("ec2.TerminateInstances call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteInstance) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteInstance) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.TerminateInstancesInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.TerminateInstancesInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.TerminateInstances(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.TerminateInstances call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: delete instance ok")
			return fakeDryRunId("instance"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *DeleteInstance) ParamsHelp() string {
	return generateParamsHelp("deleteinstance", structListParamsKeys(cmd))
}

func (cmd *DeleteInstance) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteInternetgateway(sess *session.Session, l ...*logger.Logger) *DeleteInternetgateway {
	cmd := new(DeleteInternetgateway)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *DeleteInternetgateway) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *DeleteInternetgateway) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.DeleteInternetGatewayInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.DeleteInternetGatewayInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DeleteInternetGateway(input)
	cmd.logger.ExtraVerbosef("ec2.DeleteInternetGateway call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteInternetgateway) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteInternetgateway) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.DeleteInternetGatewayInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.DeleteInternetGatewayInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.DeleteInternetGateway(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.DeleteInternetGateway call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: delete internetgateway ok")
			return fakeDryRunId("internetgateway"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *DeleteInternetgateway) ParamsHelp() string {
	return generateParamsHelp("deleteinternetgateway", structListParamsKeys(cmd))
}

func (cmd *DeleteInternetgateway) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteKeypair(sess *session.Session, l ...*logger.Logger) *DeleteKeypair {
	cmd := new(DeleteKeypair)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *DeleteKeypair) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *DeleteKeypair) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.DeleteKeyPairInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.DeleteKeyPairInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DeleteKeyPair(input)
	cmd.logger.ExtraVerbosef("ec2.DeleteKeyPair call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteKeypair) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteKeypair) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.DeleteKeyPairInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.DeleteKeyPairInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.DeleteKeyPair(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.DeleteKeyPair call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: delete keypair ok")
			return fakeDryRunId("keypair"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *DeleteKeypair) ParamsHelp() string {
	return generateParamsHelp("deletekeypair", structListParamsKeys(cmd))
}

func (cmd *DeleteKeypair) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeletePolicy(sess *session.Session, l ...*logger.Logger) *DeletePolicy {
	cmd := new(DeletePolicy)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = iam.New(sess)
	}
	return cmd
}

func (cmd *DeletePolicy) SetApi(api iamiface.IAMAPI) {
	cmd.api = api
}

func (cmd *DeletePolicy) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &iam.DeletePolicyInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in iam.DeletePolicyInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DeletePolicy(input)
	cmd.logger.ExtraVerbosef("iam.DeletePolicy call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeletePolicy) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeletePolicy) ParamsHelp() string {
	return generateParamsHelp("deletepolicy", structListParamsKeys(cmd))
}

func (cmd *DeletePolicy) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteRoute(sess *session.Session, l ...*logger.Logger) *DeleteRoute {
	cmd := new(DeleteRoute)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *DeleteRoute) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *DeleteRoute) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.DeleteRouteInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.DeleteRouteInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DeleteRoute(input)
	cmd.logger.ExtraVerbosef("ec2.DeleteRoute call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteRoute) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteRoute) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.DeleteRouteInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.DeleteRouteInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.DeleteRoute(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.DeleteRoute call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: delete route ok")
			return fakeDryRunId("route"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *DeleteRoute) ParamsHelp() string {
	return generateParamsHelp("deleteroute", structListParamsKeys(cmd))
}

func (cmd *DeleteRoute) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteRoutetable(sess *session.Session, l ...*logger.Logger) *DeleteRoutetable {
	cmd := new(DeleteRoutetable)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *DeleteRoutetable) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *DeleteRoutetable) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.DeleteRouteTableInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.DeleteRouteTableInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DeleteRouteTable(input)
	cmd.logger.ExtraVerbosef("ec2.DeleteRouteTable call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteRoutetable) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteRoutetable) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.DeleteRouteTableInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.DeleteRouteTableInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.DeleteRouteTable(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.DeleteRouteTable call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: delete routetable ok")
			return fakeDryRunId("routetable"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *DeleteRoutetable) ParamsHelp() string {
	return generateParamsHelp("deleteroutetable", structListParamsKeys(cmd))
}

func (cmd *DeleteRoutetable) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteScalingpolicy(sess *session.Session, l ...*logger.Logger) *DeleteScalingpolicy {
	cmd := new(DeleteScalingpolicy)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = autoscaling.New(sess)
	}
	return cmd
}

func (cmd *DeleteScalingpolicy) SetApi(api autoscalingiface.AutoScalingAPI) {
	cmd.api = api
}

func (cmd *DeleteScalingpolicy) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &autoscaling.DeletePolicyInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in autoscaling.DeletePolicyInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DeletePolicy(input)
	cmd.logger.ExtraVerbosef("autoscaling.DeletePolicy call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteScalingpolicy) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteScalingpolicy) ParamsHelp() string {
	return generateParamsHelp("deletescalingpolicy", structListParamsKeys(cmd))
}

func (cmd *DeleteScalingpolicy) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteSecuritygroup(sess *session.Session, l ...*logger.Logger) *DeleteSecuritygroup {
	cmd := new(DeleteSecuritygroup)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *DeleteSecuritygroup) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *DeleteSecuritygroup) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.DeleteSecurityGroupInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.DeleteSecurityGroupInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DeleteSecurityGroup(input)
	cmd.logger.ExtraVerbosef("ec2.DeleteSecurityGroup call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteSecuritygroup) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteSecuritygroup) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.DeleteSecurityGroupInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.DeleteSecurityGroupInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.DeleteSecurityGroup(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.DeleteSecurityGroup call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: delete securitygroup ok")
			return fakeDryRunId("securitygroup"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *DeleteSecuritygroup) ParamsHelp() string {
	return generateParamsHelp("deletesecuritygroup", structListParamsKeys(cmd))
}

func (cmd *DeleteSecuritygroup) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteSnapshot(sess *session.Session, l ...*logger.Logger) *DeleteSnapshot {
	cmd := new(DeleteSnapshot)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *DeleteSnapshot) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *DeleteSnapshot) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.DeleteSnapshotInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.DeleteSnapshotInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DeleteSnapshot(input)
	cmd.logger.ExtraVerbosef("ec2.DeleteSnapshot call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteSnapshot) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteSnapshot) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.DeleteSnapshotInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.DeleteSnapshotInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.DeleteSnapshot(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.DeleteSnapshot call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: delete snapshot ok")
			return fakeDryRunId("snapshot"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *DeleteSnapshot) ParamsHelp() string {
	return generateParamsHelp("deletesnapshot", structListParamsKeys(cmd))
}

func (cmd *DeleteSnapshot) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteStack(sess *session.Session, l ...*logger.Logger) *DeleteStack {
	cmd := new(DeleteStack)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = cloudformation.New(sess)
	}
	return cmd
}

func (cmd *DeleteStack) SetApi(api cloudformationiface.CloudFormationAPI) {
	cmd.api = api
}

func (cmd *DeleteStack) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &cloudformation.DeleteStackInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in cloudformation.DeleteStackInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DeleteStack(input)
	cmd.logger.ExtraVerbosef("cloudformation.DeleteStack call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteStack) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteStack) ParamsHelp() string {
	return generateParamsHelp("deletestack", structListParamsKeys(cmd))
}

func (cmd *DeleteStack) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteSubnet(sess *session.Session, l ...*logger.Logger) *DeleteSubnet {
	cmd := new(DeleteSubnet)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *DeleteSubnet) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *DeleteSubnet) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.DeleteSubnetInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.DeleteSubnetInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DeleteSubnet(input)
	cmd.logger.ExtraVerbosef("ec2.DeleteSubnet call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteSubnet) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteSubnet) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.DeleteSubnetInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.DeleteSubnetInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.DeleteSubnet(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.DeleteSubnet call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: delete subnet ok")
			return fakeDryRunId("subnet"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *DeleteSubnet) ParamsHelp() string {
	return generateParamsHelp("deletesubnet", structListParamsKeys(cmd))
}

func (cmd *DeleteSubnet) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteSubscription(sess *session.Session, l ...*logger.Logger) *DeleteSubscription {
	cmd := new(DeleteSubscription)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = sns.New(sess)
	}
	return cmd
}

func (cmd *DeleteSubscription) SetApi(api snsiface.SNSAPI) {
	cmd.api = api
}

func (cmd *DeleteSubscription) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &sns.UnsubscribeInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in sns.UnsubscribeInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.Unsubscribe(input)
	cmd.logger.ExtraVerbosef("sns.Unsubscribe call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteSubscription) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteSubscription) ParamsHelp() string {
	return generateParamsHelp("deletesubscription", structListParamsKeys(cmd))
}

func (cmd *DeleteSubscription) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteTag(sess *session.Session, l ...*logger.Logger) *DeleteTag {
	cmd := new(DeleteTag)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *DeleteTag) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *DeleteTag) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx)
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteTag) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteTag) ParamsHelp() string {
	return generateParamsHelp("deletetag", structListParamsKeys(cmd))
}

func (cmd *DeleteTag) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteTargetgroup(sess *session.Session, l ...*logger.Logger) *DeleteTargetgroup {
	cmd := new(DeleteTargetgroup)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = elbv2.New(sess)
	}
	return cmd
}

func (cmd *DeleteTargetgroup) SetApi(api elbv2iface.ELBV2API) {
	cmd.api = api
}

func (cmd *DeleteTargetgroup) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &elbv2.DeleteTargetGroupInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in elbv2.DeleteTargetGroupInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DeleteTargetGroup(input)
	cmd.logger.ExtraVerbosef("elbv2.DeleteTargetGroup call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteTargetgroup) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteTargetgroup) ParamsHelp() string {
	return generateParamsHelp("deletetargetgroup", structListParamsKeys(cmd))
}

func (cmd *DeleteTargetgroup) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteTopic(sess *session.Session, l ...*logger.Logger) *DeleteTopic {
	cmd := new(DeleteTopic)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = sns.New(sess)
	}
	return cmd
}

func (cmd *DeleteTopic) SetApi(api snsiface.SNSAPI) {
	cmd.api = api
}

func (cmd *DeleteTopic) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &sns.DeleteTopicInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in sns.DeleteTopicInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DeleteTopic(input)
	cmd.logger.ExtraVerbosef("sns.DeleteTopic call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteTopic) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteTopic) ParamsHelp() string {
	return generateParamsHelp("deletetopic", structListParamsKeys(cmd))
}

func (cmd *DeleteTopic) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteUser(sess *session.Session, l ...*logger.Logger) *DeleteUser {
	cmd := new(DeleteUser)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = iam.New(sess)
	}
	return cmd
}

func (cmd *DeleteUser) SetApi(api iamiface.IAMAPI) {
	cmd.api = api
}

func (cmd *DeleteUser) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &iam.DeleteUserInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in iam.DeleteUserInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DeleteUser(input)
	cmd.logger.ExtraVerbosef("iam.DeleteUser call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteUser) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteUser) ParamsHelp() string {
	return generateParamsHelp("deleteuser", structListParamsKeys(cmd))
}

func (cmd *DeleteUser) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteVolume(sess *session.Session, l ...*logger.Logger) *DeleteVolume {
	cmd := new(DeleteVolume)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *DeleteVolume) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *DeleteVolume) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.DeleteVolumeInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.DeleteVolumeInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DeleteVolume(input)
	cmd.logger.ExtraVerbosef("ec2.DeleteVolume call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteVolume) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteVolume) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.DeleteVolumeInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.DeleteVolumeInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.DeleteVolume(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.DeleteVolume call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: delete volume ok")
			return fakeDryRunId("volume"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *DeleteVolume) ParamsHelp() string {
	return generateParamsHelp("deletevolume", structListParamsKeys(cmd))
}

func (cmd *DeleteVolume) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteVpc(sess *session.Session, l ...*logger.Logger) *DeleteVpc {
	cmd := new(DeleteVpc)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *DeleteVpc) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *DeleteVpc) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.DeleteVpcInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.DeleteVpcInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DeleteVpc(input)
	cmd.logger.ExtraVerbosef("ec2.DeleteVpc call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteVpc) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteVpc) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.DeleteVpcInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.DeleteVpcInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.DeleteVpc(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.DeleteVpc call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: delete vpc ok")
			return fakeDryRunId("vpc"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *DeleteVpc) ParamsHelp() string {
	return generateParamsHelp("deletevpc", structListParamsKeys(cmd))
}

func (cmd *DeleteVpc) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteZone(sess *session.Session, l ...*logger.Logger) *DeleteZone {
	cmd := new(DeleteZone)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = route53.New(sess)
	}
	return cmd
}

func (cmd *DeleteZone) SetApi(api route53iface.Route53API) {
	cmd.api = api
}

func (cmd *DeleteZone) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &route53.DeleteHostedZoneInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in route53.DeleteHostedZoneInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DeleteHostedZone(input)
	cmd.logger.ExtraVerbosef("route53.DeleteHostedZone call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DeleteZone) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DeleteZone) ParamsHelp() string {
	return generateParamsHelp("deletezone", structListParamsKeys(cmd))
}

func (cmd *DeleteZone) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDetachAlarm(sess *session.Session, l ...*logger.Logger) *DetachAlarm {
	cmd := new(DetachAlarm)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = cloudwatch.New(sess)
	}
	return cmd
}

func (cmd *DetachAlarm) SetApi(api cloudwatchiface.CloudWatchAPI) {
	cmd.api = api
}

func (cmd *DetachAlarm) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx)
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DetachAlarm) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DetachAlarm) ParamsHelp() string {
	return generateParamsHelp("detachalarm", structListParamsKeys(cmd))
}

func (cmd *DetachAlarm) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDetachContainertask(sess *session.Session, l ...*logger.Logger) *DetachContainertask {
	cmd := new(DetachContainertask)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ecs.New(sess)
	}
	return cmd
}

func (cmd *DetachContainertask) SetApi(api ecsiface.ECSAPI) {
	cmd.api = api
}

func (cmd *DetachContainertask) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx)
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DetachContainertask) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DetachContainertask) ParamsHelp() string {
	return generateParamsHelp("detachcontainertask", structListParamsKeys(cmd))
}

func (cmd *DetachContainertask) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDetachElasticip(sess *session.Session, l ...*logger.Logger) *DetachElasticip {
	cmd := new(DetachElasticip)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *DetachElasticip) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *DetachElasticip) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.DisassociateAddressInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.DisassociateAddressInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DisassociateAddress(input)
	cmd.logger.ExtraVerbosef("ec2.DisassociateAddress call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DetachElasticip) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DetachElasticip) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.DisassociateAddressInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.DisassociateAddressInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.DisassociateAddress(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.DisassociateAddress call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: detach elasticip ok")
			return fakeDryRunId("elasticip"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *DetachElasticip) ParamsHelp() string {
	return generateParamsHelp("detachelasticip", structListParamsKeys(cmd))
}

func (cmd *DetachElasticip) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDetachInstance(sess *session.Session, l ...*logger.Logger) *DetachInstance {
	cmd := new(DetachInstance)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = elbv2.New(sess)
	}
	return cmd
}

func (cmd *DetachInstance) SetApi(api elbv2iface.ELBV2API) {
	cmd.api = api
}

func (cmd *DetachInstance) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &elbv2.DeregisterTargetsInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in elbv2.DeregisterTargetsInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DeregisterTargets(input)
	cmd.logger.ExtraVerbosef("elbv2.DeregisterTargets call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DetachInstance) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DetachInstance) ParamsHelp() string {
	return generateParamsHelp("detachinstance", structListParamsKeys(cmd))
}

func (cmd *DetachInstance) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDetachInternetgateway(sess *session.Session, l ...*logger.Logger) *DetachInternetgateway {
	cmd := new(DetachInternetgateway)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *DetachInternetgateway) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *DetachInternetgateway) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.DetachInternetGatewayInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.DetachInternetGatewayInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DetachInternetGateway(input)
	cmd.logger.ExtraVerbosef("ec2.DetachInternetGateway call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DetachInternetgateway) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DetachInternetgateway) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.DetachInternetGatewayInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.DetachInternetGatewayInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.DetachInternetGateway(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.DetachInternetGateway call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: detach internetgateway ok")
			return fakeDryRunId("internetgateway"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *DetachInternetgateway) ParamsHelp() string {
	return generateParamsHelp("detachinternetgateway", structListParamsKeys(cmd))
}

func (cmd *DetachInternetgateway) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDetachPolicy(sess *session.Session, l ...*logger.Logger) *DetachPolicy {
	cmd := new(DetachPolicy)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = iam.New(sess)
	}
	return cmd
}

func (cmd *DetachPolicy) SetApi(api iamiface.IAMAPI) {
	cmd.api = api
}

func (cmd *DetachPolicy) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx)
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DetachPolicy) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DetachPolicy) ParamsHelp() string {
	return generateParamsHelp("detachpolicy", structListParamsKeys(cmd))
}

func (cmd *DetachPolicy) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDetachRoutetable(sess *session.Session, l ...*logger.Logger) *DetachRoutetable {
	cmd := new(DetachRoutetable)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *DetachRoutetable) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *DetachRoutetable) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.DisassociateRouteTableInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.DisassociateRouteTableInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DisassociateRouteTable(input)
	cmd.logger.ExtraVerbosef("ec2.DisassociateRouteTable call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DetachRoutetable) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DetachRoutetable) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.DisassociateRouteTableInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.DisassociateRouteTableInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.DisassociateRouteTable(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.DisassociateRouteTable call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: detach routetable ok")
			return fakeDryRunId("routetable"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *DetachRoutetable) ParamsHelp() string {
	return generateParamsHelp("detachroutetable", structListParamsKeys(cmd))
}

func (cmd *DetachRoutetable) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDetachSecuritygroup(sess *session.Session, l ...*logger.Logger) *DetachSecuritygroup {
	cmd := new(DetachSecuritygroup)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *DetachSecuritygroup) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *DetachSecuritygroup) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx)
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DetachSecuritygroup) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DetachSecuritygroup) ParamsHelp() string {
	return generateParamsHelp("detachsecuritygroup", structListParamsKeys(cmd))
}

func (cmd *DetachSecuritygroup) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDetachUser(sess *session.Session, l ...*logger.Logger) *DetachUser {
	cmd := new(DetachUser)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = iam.New(sess)
	}
	return cmd
}

func (cmd *DetachUser) SetApi(api iamiface.IAMAPI) {
	cmd.api = api
}

func (cmd *DetachUser) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &iam.RemoveUserFromGroupInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in iam.RemoveUserFromGroupInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.RemoveUserFromGroup(input)
	cmd.logger.ExtraVerbosef("iam.RemoveUserFromGroup call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DetachUser) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DetachUser) ParamsHelp() string {
	return generateParamsHelp("detachuser", structListParamsKeys(cmd))
}

func (cmd *DetachUser) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDetachVolume(sess *session.Session, l ...*logger.Logger) *DetachVolume {
	cmd := new(DetachVolume)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *DetachVolume) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *DetachVolume) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.DetachVolumeInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.DetachVolumeInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DetachVolume(input)
	cmd.logger.ExtraVerbosef("ec2.DetachVolume call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *DetachVolume) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *DetachVolume) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.DetachVolumeInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.DetachVolumeInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.DetachVolume(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.DetachVolume call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: detach volume ok")
			return fakeDryRunId("volume"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *DetachVolume) ParamsHelp() string {
	return generateParamsHelp("detachvolume", structListParamsKeys(cmd))
}

func (cmd *DetachVolume) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewImportImage(sess *session.Session, l ...*logger.Logger) *ImportImage {
	cmd := new(ImportImage)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *ImportImage) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *ImportImage) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.ImportImageInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.ImportImageInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.ImportImage(input)
	cmd.logger.ExtraVerbosef("ec2.ImportImage call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *ImportImage) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *ImportImage) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.ImportImageInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.ImportImageInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.ImportImage(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.ImportImage call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: import image ok")
			return fakeDryRunId("image"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *ImportImage) ParamsHelp() string {
	return generateParamsHelp("importimage", structListParamsKeys(cmd))
}

func (cmd *ImportImage) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewStartAlarm(sess *session.Session, l ...*logger.Logger) *StartAlarm {
	cmd := new(StartAlarm)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = cloudwatch.New(sess)
	}
	return cmd
}

func (cmd *StartAlarm) SetApi(api cloudwatchiface.CloudWatchAPI) {
	cmd.api = api
}

func (cmd *StartAlarm) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &cloudwatch.EnableAlarmActionsInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in cloudwatch.EnableAlarmActionsInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.EnableAlarmActions(input)
	cmd.logger.ExtraVerbosef("cloudwatch.EnableAlarmActions call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *StartAlarm) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *StartAlarm) ParamsHelp() string {
	return generateParamsHelp("startalarm", structListParamsKeys(cmd))
}

func (cmd *StartAlarm) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewStartContainertask(sess *session.Session, l ...*logger.Logger) *StartContainertask {
	cmd := new(StartContainertask)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ecs.New(sess)
	}
	return cmd
}

func (cmd *StartContainertask) SetApi(api ecsiface.ECSAPI) {
	cmd.api = api
}

func (cmd *StartContainertask) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx)
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *StartContainertask) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *StartContainertask) ParamsHelp() string {
	return generateParamsHelp("startcontainertask", structListParamsKeys(cmd))
}

func (cmd *StartContainertask) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewStartInstance(sess *session.Session, l ...*logger.Logger) *StartInstance {
	cmd := new(StartInstance)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *StartInstance) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *StartInstance) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.StartInstancesInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.StartInstancesInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.StartInstances(input)
	cmd.logger.ExtraVerbosef("ec2.StartInstances call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *StartInstance) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *StartInstance) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.StartInstancesInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.StartInstancesInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.StartInstances(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.StartInstances call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: start instance ok")
			return fakeDryRunId("instance"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *StartInstance) ParamsHelp() string {
	return generateParamsHelp("startinstance", structListParamsKeys(cmd))
}

func (cmd *StartInstance) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewStopAlarm(sess *session.Session, l ...*logger.Logger) *StopAlarm {
	cmd := new(StopAlarm)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = cloudwatch.New(sess)
	}
	return cmd
}

func (cmd *StopAlarm) SetApi(api cloudwatchiface.CloudWatchAPI) {
	cmd.api = api
}

func (cmd *StopAlarm) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &cloudwatch.DisableAlarmActionsInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in cloudwatch.DisableAlarmActionsInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.DisableAlarmActions(input)
	cmd.logger.ExtraVerbosef("cloudwatch.DisableAlarmActions call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *StopAlarm) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *StopAlarm) ParamsHelp() string {
	return generateParamsHelp("stopalarm", structListParamsKeys(cmd))
}

func (cmd *StopAlarm) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewStopContainertask(sess *session.Session, l ...*logger.Logger) *StopContainertask {
	cmd := new(StopContainertask)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ecs.New(sess)
	}
	return cmd
}

func (cmd *StopContainertask) SetApi(api ecsiface.ECSAPI) {
	cmd.api = api
}

func (cmd *StopContainertask) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx)
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *StopContainertask) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *StopContainertask) ParamsHelp() string {
	return generateParamsHelp("stopcontainertask", structListParamsKeys(cmd))
}

func (cmd *StopContainertask) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewStopInstance(sess *session.Session, l ...*logger.Logger) *StopInstance {
	cmd := new(StopInstance)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *StopInstance) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *StopInstance) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.StopInstancesInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.StopInstancesInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.StopInstances(input)
	cmd.logger.ExtraVerbosef("ec2.StopInstances call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *StopInstance) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *StopInstance) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.StopInstancesInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.StopInstancesInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.StopInstances(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.StopInstances call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: stop instance ok")
			return fakeDryRunId("instance"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *StopInstance) ParamsHelp() string {
	return generateParamsHelp("stopinstance", structListParamsKeys(cmd))
}

func (cmd *StopInstance) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewUpdateBucket(sess *session.Session, l ...*logger.Logger) *UpdateBucket {
	cmd := new(UpdateBucket)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = s3.New(sess)
	}
	return cmd
}

func (cmd *UpdateBucket) SetApi(api s3iface.S3API) {
	cmd.api = api
}

func (cmd *UpdateBucket) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx)
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *UpdateBucket) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *UpdateBucket) ParamsHelp() string {
	return generateParamsHelp("updatebucket", structListParamsKeys(cmd))
}

func (cmd *UpdateBucket) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewUpdateContainertask(sess *session.Session, l ...*logger.Logger) *UpdateContainertask {
	cmd := new(UpdateContainertask)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ecs.New(sess)
	}
	return cmd
}

func (cmd *UpdateContainertask) SetApi(api ecsiface.ECSAPI) {
	cmd.api = api
}

func (cmd *UpdateContainertask) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ecs.UpdateServiceInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ecs.UpdateServiceInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.UpdateService(input)
	cmd.logger.ExtraVerbosef("ecs.UpdateService call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *UpdateContainertask) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *UpdateContainertask) ParamsHelp() string {
	return generateParamsHelp("updatecontainertask", structListParamsKeys(cmd))
}

func (cmd *UpdateContainertask) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewUpdateDistribution(sess *session.Session, l ...*logger.Logger) *UpdateDistribution {
	cmd := new(UpdateDistribution)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = cloudfront.New(sess)
	}
	return cmd
}

func (cmd *UpdateDistribution) SetApi(api cloudfrontiface.CloudFrontAPI) {
	cmd.api = api
}

func (cmd *UpdateDistribution) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx)
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *UpdateDistribution) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *UpdateDistribution) ParamsHelp() string {
	return generateParamsHelp("updatedistribution", structListParamsKeys(cmd))
}

func (cmd *UpdateDistribution) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewUpdateInstance(sess *session.Session, l ...*logger.Logger) *UpdateInstance {
	cmd := new(UpdateInstance)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *UpdateInstance) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *UpdateInstance) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.ModifyInstanceAttributeInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.ModifyInstanceAttributeInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.ModifyInstanceAttribute(input)
	cmd.logger.ExtraVerbosef("ec2.ModifyInstanceAttribute call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *UpdateInstance) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *UpdateInstance) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("dry run: cannot set params on command struct: %s", err)
	}

	input := &ec2.ModifyInstanceAttributeInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("dry run: cannot inject in ec2.ModifyInstanceAttributeInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.ModifyInstanceAttribute(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.ModifyInstanceAttribute call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: update instance ok")
			return fakeDryRunId("instance"), nil
		}
	}

	return nil, fmt.Errorf("dry run: %s", err)
}

func (cmd *UpdateInstance) ParamsHelp() string {
	return generateParamsHelp("updateinstance", structListParamsKeys(cmd))
}

func (cmd *UpdateInstance) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewUpdatePolicy(sess *session.Session, l ...*logger.Logger) *UpdatePolicy {
	cmd := new(UpdatePolicy)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = iam.New(sess)
	}
	return cmd
}

func (cmd *UpdatePolicy) SetApi(api iamiface.IAMAPI) {
	cmd.api = api
}

func (cmd *UpdatePolicy) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &iam.CreatePolicyVersionInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in iam.CreatePolicyVersionInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.CreatePolicyVersion(input)
	cmd.logger.ExtraVerbosef("iam.CreatePolicyVersion call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *UpdatePolicy) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *UpdatePolicy) ParamsHelp() string {
	return generateParamsHelp("updatepolicy", structListParamsKeys(cmd))
}

func (cmd *UpdatePolicy) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewUpdateSecuritygroup(sess *session.Session, l ...*logger.Logger) *UpdateSecuritygroup {
	cmd := new(UpdateSecuritygroup)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *UpdateSecuritygroup) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *UpdateSecuritygroup) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx)
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *UpdateSecuritygroup) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *UpdateSecuritygroup) ParamsHelp() string {
	return generateParamsHelp("updatesecuritygroup", structListParamsKeys(cmd))
}

func (cmd *UpdateSecuritygroup) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewUpdateStack(sess *session.Session, l ...*logger.Logger) *UpdateStack {
	cmd := new(UpdateStack)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = cloudformation.New(sess)
	}
	return cmd
}

func (cmd *UpdateStack) SetApi(api cloudformationiface.CloudFormationAPI) {
	cmd.api = api
}

func (cmd *UpdateStack) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &cloudformation.UpdateStackInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in cloudformation.UpdateStackInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.UpdateStack(input)
	cmd.logger.ExtraVerbosef("cloudformation.UpdateStack call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *UpdateStack) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *UpdateStack) ParamsHelp() string {
	return generateParamsHelp("updatestack", structListParamsKeys(cmd))
}

func (cmd *UpdateStack) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewUpdateSubnet(sess *session.Session, l ...*logger.Logger) *UpdateSubnet {
	cmd := new(UpdateSubnet)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = ec2.New(sess)
	}
	return cmd
}

func (cmd *UpdateSubnet) SetApi(api ec2iface.EC2API) {
	cmd.api = api
}

func (cmd *UpdateSubnet) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	input := &ec2.ModifySubnetAttributeInput{}
	if err := structInjector(cmd, input, ctx); err != nil {
		return nil, fmt.Errorf("cannot inject in ec2.ModifySubnetAttributeInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.ModifySubnetAttribute(input)
	cmd.logger.ExtraVerbosef("ec2.ModifySubnetAttribute call took %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *UpdateSubnet) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *UpdateSubnet) ParamsHelp() string {
	return generateParamsHelp("updatesubnet", structListParamsKeys(cmd))
}

func (cmd *UpdateSubnet) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewUpdateTargetgroup(sess *session.Session, l ...*logger.Logger) *UpdateTargetgroup {
	cmd := new(UpdateTargetgroup)
	if len(l) > 0 {
		cmd.logger = l[0]
	} else {
		cmd.logger = logger.DiscardLogger
	}
	if sess != nil {
		cmd.api = elbv2.New(sess)
	}
	return cmd
}

func (cmd *UpdateTargetgroup) SetApi(api elbv2iface.ELBV2API) {
	cmd.api = api
}

func (cmd *UpdateTargetgroup) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
	}

	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx)
	if err != nil {
		return nil, err
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("after run: %s", brErr)
		}
	}

	if v, ok := implementsResultExtractor(cmd); ok {
		return v.ExtractResult(output), nil
	}
	return nil, nil
}

func (cmd *UpdateTargetgroup) ValidateCommand(params map[string]interface{}, refs []string) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd, refs); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params, refs)...)
	}

	return
}

func (cmd *UpdateTargetgroup) ParamsHelp() string {
	return generateParamsHelp("updatetargetgroup", structListParamsKeys(cmd))
}

func (cmd *UpdateTargetgroup) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}
