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

type CreateLoginprofile struct {
	_             string `action: "create" entity: "loginprofile" awsAPI: "iam" awsCall: "CreateLoginProfile" awsInput: "CreateLoginProfileInput" awsOutput: "CreateLoginProfileOutput"`
	logger        *logger.Logger
	api           iamiface.IAMAPI
	Username      *string `awsName: "UserName" awsType: "awsstr" templateName: "username" required: ""`
	Password      *string `awsName: "Password" awsType: "awsstr" templateName: "password" required: ""`
	PasswordReset *bool   `awsName: "PasswordResetRequired" awsType: "awsbool" templateName: "password-reset"`
}
type UpdateLoginprofile struct {
	_             string `action: "update" entity: "loginprofile" awsAPI: "iam" awsCall: "UpdateLoginProfile" awsInput: "UpdateLoginProfileInput" awsOutput: "UpdateLoginProfileOutput"`
	logger        *logger.Logger
	api           iamiface.IAMAPI
	Username      *string `awsName: "UserName" awsType: "awsstr" templateName: "username" required: ""`
	Password      *string `awsName: "Password" awsType: "awsstr" templateName: "password" required: ""`
	PasswordReset *bool   `awsName: "PasswordResetRequired" awsType: "awsbool" templateName: "password-reset"`
}
type DeleteLoginprofile struct {
	_        string `action: "delete" entity: "loginprofile" awsAPI: "iam" awsCall: "DeleteLoginProfile" awsInput: "DeleteLoginProfileInput" awsOutput: "DeleteLoginProfileOutput"`
	logger   *logger.Logger
	api      iamiface.IAMAPI
	Username *string `awsName: "UserName" awsType: "awsstr" templateName: "username" required: ""`
}
