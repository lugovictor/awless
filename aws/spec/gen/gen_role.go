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
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
	"github.com/wallix/awless/logger"
)

type CreateRole struct {
	_                string `action:"create" entity:"role" awsAPI:"iam"`
	logger           *logger.Logger
	api              iamiface.IAMAPI
	Name             *string   `awsName:"RoleName" awsType:"awsstr" templateName:"name" required:""`
	PrincipalAccount *struct{} `templateName:"principal-account"`
	PrincipalUser    *struct{} `templateName:"principal-user"`
	PrincipalService *struct{} `templateName:"principal-service"`
	Conditions       *struct{} `templateName:"conditions"`
	SleepAfter       *struct{} `templateName:"sleep-after"`
}

type DeleteRole struct {
	_      string `action:"delete" entity:"role" awsAPI:"iam"`
	logger *logger.Logger
	api    iamiface.IAMAPI
	Name   *string `awsName:"RoleName" awsType:"awsstr" templateName:"name" required:""`
}

type AttachRole struct {
	_               string `action:"attach" entity:"role" awsAPI:"iam" awsCall:"AddRoleToInstanceProfile" awsInput:"iam.AddRoleToInstanceProfileInput" awsOutput:"iam.AddRoleToInstanceProfileOutput"`
	logger          *logger.Logger
	api             iamiface.IAMAPI
	Instanceprofile *string `awsName:"InstanceProfileName" awsType:"awsstr" templateName:"instanceprofile" required:""`
	Name            *string `awsName:"RoleName" awsType:"awsstr" templateName:"name" required:""`
}

func (cmd *AttachRole) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}

type DetachRole struct {
	_               string `action:"detach" entity:"role" awsAPI:"iam" awsCall:"RemoveRoleFromInstanceProfile" awsInput:"iam.RemoveRoleFromInstanceProfileInput" awsOutput:"iam.RemoveRoleFromInstanceProfileOutput"`
	logger          *logger.Logger
	api             iamiface.IAMAPI
	Instanceprofile *string `awsName:"InstanceProfileName" awsType:"awsstr" templateName:"instanceprofile" required:""`
	Name            *string `awsName:"RoleName" awsType:"awsstr" templateName:"name" required:""`
}

func (cmd *DetachRole) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}
