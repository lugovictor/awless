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
	"github.com/aws/aws-sdk-go/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go/service/applicationautoscaling/applicationautoscalingiface"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx, params)
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx, params)
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx, params)
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx, params)
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx, params)
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx, params)
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx, params)
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx, params)
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx, params)
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx, params)
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx, params)
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx, params)
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	output, err := cmd.ManualRun(ctx, params)
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
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
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
