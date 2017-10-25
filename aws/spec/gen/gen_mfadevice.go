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

type CreateMfadevice struct {
	_      string `action: "create" entity: "mfadevice" awsAPI: "iam"`
	logger *logger.Logger
	api    iamiface.IAMAPI
	Name   *struct{} `templateName: "name" required: ""`
}
type DeleteMfadevice struct {
	_      string `action: "delete" entity: "mfadevice" awsAPI: "iam" awsCall: "DeleteVirtualMFADevice" awsInput: "DeleteVirtualMFADeviceInput" awsOutput: "DeleteVirtualMFADeviceOutput"`
	logger *logger.Logger
	api    iamiface.IAMAPI
	Id     *string `awsName: "SerialNumber" awsType: "awsstr" templateName: "id" required: ""`
}
type AttachMfadevice struct {
	_        string `action: "attach" entity: "mfadevice" awsAPI: "iam" awsCall: "EnableMFADevice" awsInput: "EnableMFADeviceInput" awsOutput: "EnableMFADeviceOutput"`
	logger   *logger.Logger
	api      iamiface.IAMAPI
	Id       *string `awsName: "SerialNumber" awsType: "awsstr" templateName: "id" required: ""`
	User     *string `awsName: "UserName" awsType: "awsstr" templateName: "user" required: ""`
	MfaCode1 *string `awsName: "AuthenticationCode1" awsType: "aws6digitsstring" templateName: "mfa-code-1" required: ""`
	MfaCode2 *string `awsName: "AuthenticationCode2" awsType: "aws6digitsstring" templateName: "mfa-code-2" required: ""`
}
type DetachMfadevice struct {
	_      string `action: "detach" entity: "mfadevice" awsAPI: "iam" awsCall: "DeactivateMFADevice" awsInput: "DeactivateMFADeviceInput" awsOutput: "DeactivateMFADeviceOutput"`
	logger *logger.Logger
	api    iamiface.IAMAPI
	Id     *string `awsName: "SerialNumber" awsType: "awsstr" templateName: "id" required: ""`
	User   *string `awsName: "UserName" awsType: "awsstr" templateName: "user" required: ""`
}
