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

func NewAttachPolicy(l *logger.Logger, sess *session.Session) *AttachPolicy {
	cmd := new(AttachPolicy)
	cmd.api = iam.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *AttachPolicy) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("AttachPolicy: BeforeRun: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("AttachPolicy: cannot set params on command struct: %s", err)
	}

	output, err := cmd.ManualRun(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("AttachPolicy: %s", err)
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("AttachPolicy: AfterRun: %s", brErr)
		}
	}

	return cmd.ExtractResultString(output), nil
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

func NewCreateInstance(l *logger.Logger, sess *session.Session) *CreateInstance {
	cmd := new(CreateInstance)
	cmd.api = ec2.New(sess)
	cmd.logger = l
	return cmd
}

func (cmd *CreateInstance) Run(ctx, params map[string]interface{}) (interface{}, error) {
	if v, ok := implementsBeforeRun(cmd); ok {
		if brErr := v.BeforeRun(ctx, params); brErr != nil {
			return nil, fmt.Errorf("CreateInstance: BeforeRun: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("CreateInstance: cannot set params on command struct: %s", err)
	}

	input := &ec2.RunInstancesInput{}
	if err := structInjector(cmd, input); err != nil {
		return nil, fmt.Errorf("CreateInstance: cannot inject in ec2.RunInstancesInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.RunInstances(input)
	cmd.logger.ExtraVerbosef("ec2.RunInstances call took %s", time.Since(start))
	if err != nil {
		return nil, fmt.Errorf("CreateInstance: %s", err)
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("CreateInstance: AfterRun: %s", brErr)
		}
	}

	return cmd.ExtractResultString(output), nil
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
		return nil, fmt.Errorf("dry run: CreateInstance: cannot set params on command struct: %s", err)
	}

	input := &ec2.RunInstancesInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input); err != nil {
		return nil, fmt.Errorf("dry run: CreateInstance: cannot inject in ec2.RunInstancesInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.RunInstances(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.RunInstances call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: CreateInstance ok")
			return fakeDryRunId(cmd.Entity()), nil
		}
	}

	return nil, fmt.Errorf("dry run: CreateInstance : %s", err)
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
			return nil, fmt.Errorf("CreateSubnet: BeforeRun: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("CreateSubnet: cannot set params on command struct: %s", err)
	}

	input := &ec2.CreateSubnetInput{}
	if err := structInjector(cmd, input); err != nil {
		return nil, fmt.Errorf("CreateSubnet: cannot inject in ec2.CreateSubnetInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.CreateSubnet(input)
	cmd.logger.ExtraVerbosef("ec2.CreateSubnet call took %s", time.Since(start))
	if err != nil {
		return nil, fmt.Errorf("CreateSubnet: %s", err)
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("CreateSubnet: AfterRun: %s", brErr)
		}
	}

	return cmd.ExtractResultString(output), nil
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
		return nil, fmt.Errorf("dry run: CreateSubnet: cannot set params on command struct: %s", err)
	}

	input := &ec2.CreateSubnetInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input); err != nil {
		return nil, fmt.Errorf("dry run: CreateSubnet: cannot inject in ec2.CreateSubnetInput: %s", err)
	}

	start := time.Now()
	_, err := cmd.api.CreateSubnet(input)
	if awsErr, ok := err.(awserr.Error); ok {
		switch code := awsErr.Code(); {
		case code == dryRunOperation, strings.HasSuffix(code, notFound), strings.Contains(awsErr.Message(), "Invalid IAM Instance Profile name"):
			cmd.logger.ExtraVerbosef("dry run: ec2.CreateSubnet call took %s", time.Since(start))
			cmd.logger.Verbose("dry run: CreateSubnet ok")
			return fakeDryRunId(cmd.Entity()), nil
		}
	}

	return nil, fmt.Errorf("dry run: CreateSubnet : %s", err)
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
			return nil, fmt.Errorf("CreateTag: BeforeRun: %s", brErr)
		}
	}

	if err := cmd.inject(params); err != nil {
		return nil, fmt.Errorf("CreateTag: cannot set params on command struct: %s", err)
	}

	output, err := cmd.ManualRun(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("CreateTag: %s", err)
	}

	if v, ok := implementsAfterRun(cmd); ok {
		if brErr := v.AfterRun(ctx, output); brErr != nil {
			return nil, fmt.Errorf("CreateTag: AfterRun: %s", brErr)
		}
	}

	return cmd.ExtractResultString(output), nil
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
