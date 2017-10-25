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

type CreateUser struct {
	_      string `action: "create" entity: "user" awsAPI: "iam" awsCall: "CreateUser" awsInput: "CreateUserInput" awsOutput: "CreateUserOutput"`
	logger *logger.Logger
	api    iamiface.IAMAPI
	Name   *string `awsName: "UserName" awsType: "awsstr" templateName: "name" required: ""`
}
type DeleteUser struct {
	_      string `action: "delete" entity: "user" awsAPI: "iam" awsCall: "DeleteUser" awsInput: "DeleteUserInput" awsOutput: "DeleteUserOutput"`
	logger *logger.Logger
	api    iamiface.IAMAPI
	Name   *string `awsName: "UserName" awsType: "awsstr" templateName: "name" required: ""`
}
type AttachUser struct {
	_      string `action: "attach" entity: "user" awsAPI: "iam" awsCall: "AddUserToGroup" awsInput: "AddUserToGroupInput" awsOutput: "AddUserToGroupOutput"`
	logger *logger.Logger
	api    iamiface.IAMAPI
	Group  *string `awsName: "GroupName" awsType: "awsstr" templateName: "group" required: ""`
	Name   *string `awsName: "UserName" awsType: "awsstr" templateName: "name" required: ""`
}
type DetachUser struct {
	_      string `action: "detach" entity: "user" awsAPI: "iam" awsCall: "RemoveUserFromGroup" awsInput: "RemoveUserFromGroupInput" awsOutput: "RemoveUserFromGroupOutput"`
	logger *logger.Logger
	api    iamiface.IAMAPI
	Group  *string `awsName: "GroupName" awsType: "awsstr" templateName: "group" required: ""`
	Name   *string `awsName: "UserName" awsType: "awsstr" templateName: "name" required: ""`
}
