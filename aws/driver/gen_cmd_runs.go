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
package awsdriver

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/wallix/awless/logger"
)

func NewCreateInstance(l *logger.Logger, sess *session.Session) *CreateInstance {
	cmd := new(CreateInstance)
	cmd.api = ec2.New(sess)
	cmd.sess = sess
	cmd.logger = l
	return cmd
}

func (cmd *CreateInstance) Run() (interface{}, error) {
	input := &ec2.RunInstancesInput{}
	if err := structInjector(cmd, input); err != nil {
		return nil, fmt.Errorf("CreateInstance: cannot inject in ec2.RunInstancesInput: %s", err)
	}
	start := time.Now()
	output, err := cmd.api.RunInstances(input)
	if err != nil {
		return nil, fmt.Errorf("CreateInstance: %s", err)
	}
	cmd.logger.ExtraVerbosef("ec2.RunInstances call took %s", time.Since(start))
	cmd.SetResult(output)
	return cmd.result, nil
}

func (cmd *CreateInstance) DryRun() (interface{}, error) {
	input := &ec2.RunInstancesInput{}
	input.SetDryRun(true)
	if err := structInjector(cmd, input); err != nil {
		return nil, fmt.Errorf("dry run: CreateInstance: cannot inject in ec2.RunInstancesInput: %s", err)
	}
	start := time.Now()
	if _, err := cmd.api.RunInstances(input); err != nil {
		return nil, fmt.Errorf("dry run: CreateInstance: %s", err)
	}
	cmd.logger.ExtraVerbosef("dry run: ec2.RunInstances call took %s", time.Since(start))
	cmd.result = fakeDryRunId(cmd.Entity())
	return cmd.result, nil
}

func NewCreateTag(l *logger.Logger, sess *session.Session) *CreateTag {
	cmd := new(CreateTag)
	cmd.api = ec2.New(sess)
	cmd.sess = sess
	cmd.logger = l
	return cmd
}
