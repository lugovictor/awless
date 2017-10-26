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
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/wallix/awless/logger"
)

func NewAttachInternetgateway(l *logger.Logger, sess *session.Session) *AttachInternetgateway {
	cmd := new(AttachInternetgateway)
	cmd.api = ec2.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *AttachInternetgateway) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
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

func (cmd *AttachInternetgateway) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
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

func NewAttachPolicy(l *logger.Logger, sess *session.Session) *AttachPolicy {
	cmd := new(AttachPolicy)
	cmd.api = iam.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *AttachPolicy) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
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

func (cmd *AttachPolicy) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
	}

	return
}

func (cmd *AttachPolicy) ParamsHelp() string {
	return generateParamsHelp("attachpolicy", structListParamsKeys(cmd))
}

func (cmd *AttachPolicy) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewAttachRoutetable(l *logger.Logger, sess *session.Session) *AttachRoutetable {
	cmd := new(AttachRoutetable)
	cmd.api = ec2.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *AttachRoutetable) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
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

func (cmd *AttachRoutetable) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
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

func NewAttachSecuritygroup(l *logger.Logger, sess *session.Session) *AttachSecuritygroup {
	cmd := new(AttachSecuritygroup)
	cmd.api = ec2.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *AttachSecuritygroup) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
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

func (cmd *AttachSecuritygroup) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
	}

	return
}

func (cmd *AttachSecuritygroup) ParamsHelp() string {
	return generateParamsHelp("attachsecuritygroup", structListParamsKeys(cmd))
}

func (cmd *AttachSecuritygroup) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCheckSecuritygroup(l *logger.Logger, sess *session.Session) *CheckSecuritygroup {
	cmd := new(CheckSecuritygroup)
	cmd.api = ec2.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *CheckSecuritygroup) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
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

func (cmd *CheckSecuritygroup) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
	}

	return
}

func (cmd *CheckSecuritygroup) ParamsHelp() string {
	return generateParamsHelp("checksecuritygroup", structListParamsKeys(cmd))
}

func (cmd *CheckSecuritygroup) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateInstance(l *logger.Logger, sess *session.Session) *CreateInstance {
	cmd := new(CreateInstance)
	cmd.api = ec2.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *CreateInstance) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
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

func (cmd *CreateInstance) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
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

func NewCreateInternetgateway(l *logger.Logger, sess *session.Session) *CreateInternetgateway {
	cmd := new(CreateInternetgateway)
	cmd.api = ec2.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *CreateInternetgateway) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
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

func (cmd *CreateInternetgateway) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
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

func NewCreateKeypair(l *logger.Logger, sess *session.Session) *CreateKeypair {
	cmd := new(CreateKeypair)
	cmd.api = ec2.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *CreateKeypair) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
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

func (cmd *CreateKeypair) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
	}

	return
}

func (cmd *CreateKeypair) ParamsHelp() string {
	return generateParamsHelp("createkeypair", structListParamsKeys(cmd))
}

func (cmd *CreateKeypair) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreatePolicy(l *logger.Logger, sess *session.Session) *CreatePolicy {
	cmd := new(CreatePolicy)
	cmd.api = iam.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *CreatePolicy) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
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

func (cmd *CreatePolicy) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
	}

	return
}

func (cmd *CreatePolicy) ParamsHelp() string {
	return generateParamsHelp("createpolicy", structListParamsKeys(cmd))
}

func (cmd *CreatePolicy) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateRoute(l *logger.Logger, sess *session.Session) *CreateRoute {
	cmd := new(CreateRoute)
	cmd.api = ec2.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *CreateRoute) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
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

func (cmd *CreateRoute) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
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

func NewCreateRoutetable(l *logger.Logger, sess *session.Session) *CreateRoutetable {
	cmd := new(CreateRoutetable)
	cmd.api = ec2.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *CreateRoutetable) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
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

func (cmd *CreateRoutetable) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
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

func NewCreateSecuritygroup(l *logger.Logger, sess *session.Session) *CreateSecuritygroup {
	cmd := new(CreateSecuritygroup)
	cmd.api = ec2.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *CreateSecuritygroup) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
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

func (cmd *CreateSecuritygroup) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
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

func NewCreateSubnet(l *logger.Logger, sess *session.Session) *CreateSubnet {
	cmd := new(CreateSubnet)
	cmd.api = ec2.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *CreateSubnet) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
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

func (cmd *CreateSubnet) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
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

func NewCreateTag(l *logger.Logger, sess *session.Session) *CreateTag {
	cmd := new(CreateTag)
	cmd.api = ec2.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *CreateTag) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
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

func (cmd *CreateTag) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
	}

	return
}

func (cmd *CreateTag) ParamsHelp() string {
	return generateParamsHelp("createtag", structListParamsKeys(cmd))
}

func (cmd *CreateTag) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewCreateVpc(l *logger.Logger, sess *session.Session) *CreateVpc {
	cmd := new(CreateVpc)
	cmd.api = ec2.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *CreateVpc) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
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

func (cmd *CreateVpc) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
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

func NewDeleteInstance(l *logger.Logger, sess *session.Session) *DeleteInstance {
	cmd := new(DeleteInstance)
	cmd.api = ec2.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *DeleteInstance) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
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

func (cmd *DeleteInstance) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
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

func NewDeleteInternetgateway(l *logger.Logger, sess *session.Session) *DeleteInternetgateway {
	cmd := new(DeleteInternetgateway)
	cmd.api = ec2.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *DeleteInternetgateway) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
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

func (cmd *DeleteInternetgateway) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
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

func NewDeleteKeypair(l *logger.Logger, sess *session.Session) *DeleteKeypair {
	cmd := new(DeleteKeypair)
	cmd.api = ec2.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *DeleteKeypair) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
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

func (cmd *DeleteKeypair) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
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

func NewDeletePolicy(l *logger.Logger, sess *session.Session) *DeletePolicy {
	cmd := new(DeletePolicy)
	cmd.api = iam.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *DeletePolicy) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
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

func (cmd *DeletePolicy) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
	}

	return
}

func (cmd *DeletePolicy) ParamsHelp() string {
	return generateParamsHelp("deletepolicy", structListParamsKeys(cmd))
}

func (cmd *DeletePolicy) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDeleteRoute(l *logger.Logger, sess *session.Session) *DeleteRoute {
	cmd := new(DeleteRoute)
	cmd.api = ec2.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *DeleteRoute) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
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

func (cmd *DeleteRoute) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
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

func NewDeleteRoutetable(l *logger.Logger, sess *session.Session) *DeleteRoutetable {
	cmd := new(DeleteRoutetable)
	cmd.api = ec2.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *DeleteRoutetable) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
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

func (cmd *DeleteRoutetable) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
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

func NewDeleteSecuritygroup(l *logger.Logger, sess *session.Session) *DeleteSecuritygroup {
	cmd := new(DeleteSecuritygroup)
	cmd.api = ec2.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *DeleteSecuritygroup) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
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

func (cmd *DeleteSecuritygroup) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
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

func NewDeleteSubnet(l *logger.Logger, sess *session.Session) *DeleteSubnet {
	cmd := new(DeleteSubnet)
	cmd.api = ec2.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *DeleteSubnet) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
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

func (cmd *DeleteSubnet) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
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

func NewDeleteVpc(l *logger.Logger, sess *session.Session) *DeleteVpc {
	cmd := new(DeleteVpc)
	cmd.api = ec2.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *DeleteVpc) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
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

func (cmd *DeleteVpc) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
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

func NewDetachInternetgateway(l *logger.Logger, sess *session.Session) *DetachInternetgateway {
	cmd := new(DetachInternetgateway)
	cmd.api = ec2.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *DetachInternetgateway) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
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

func (cmd *DetachInternetgateway) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
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

func NewDetachPolicy(l *logger.Logger, sess *session.Session) *DetachPolicy {
	cmd := new(DetachPolicy)
	cmd.api = iam.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *DetachPolicy) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
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

func (cmd *DetachPolicy) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
	}

	return
}

func (cmd *DetachPolicy) ParamsHelp() string {
	return generateParamsHelp("detachpolicy", structListParamsKeys(cmd))
}

func (cmd *DetachPolicy) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewDetachRoutetable(l *logger.Logger, sess *session.Session) *DetachRoutetable {
	cmd := new(DetachRoutetable)
	cmd.api = ec2.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *DetachRoutetable) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
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

func (cmd *DetachRoutetable) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
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

func NewDetachSecuritygroup(l *logger.Logger, sess *session.Session) *DetachSecuritygroup {
	cmd := new(DetachSecuritygroup)
	cmd.api = ec2.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *DetachSecuritygroup) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
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

func (cmd *DetachSecuritygroup) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
	}

	return
}

func (cmd *DetachSecuritygroup) ParamsHelp() string {
	return generateParamsHelp("detachsecuritygroup", structListParamsKeys(cmd))
}

func (cmd *DetachSecuritygroup) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewUpdatePolicy(l *logger.Logger, sess *session.Session) *UpdatePolicy {
	cmd := new(UpdatePolicy)
	cmd.api = iam.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *UpdatePolicy) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
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

func (cmd *UpdatePolicy) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
	}

	return
}

func (cmd *UpdatePolicy) ParamsHelp() string {
	return generateParamsHelp("updatepolicy", structListParamsKeys(cmd))
}

func (cmd *UpdatePolicy) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewUpdateSecuritygroup(l *logger.Logger, sess *session.Session) *UpdateSecuritygroup {
	cmd := new(UpdateSecuritygroup)
	cmd.api = ec2.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *UpdateSecuritygroup) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
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

func (cmd *UpdateSecuritygroup) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
	}

	return
}

func (cmd *UpdateSecuritygroup) ParamsHelp() string {
	return generateParamsHelp("updatesecuritygroup", structListParamsKeys(cmd))
}

func (cmd *UpdateSecuritygroup) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func NewUpdateSubnet(l *logger.Logger, sess *session.Session) *UpdateSubnet {
	cmd := new(UpdateSubnet)
	cmd.api = ec2.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *UpdateSubnet) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("before run: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("cannot set params on command struct: %s", err)
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

func (cmd *UpdateSubnet) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}

	if mv, ok := implementsManualValidator(cmd); ok {
		errs = append(errs, mv.ManualValidateCommand(params)...)
	}

	return
}

func (cmd *UpdateSubnet) ParamsHelp() string {
	return generateParamsHelp("updatesubnet", structListParamsKeys(cmd))
}

func (cmd *UpdateSubnet) inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}
