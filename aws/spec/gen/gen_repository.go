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
	"github.com/aws/aws-sdk-go/service/ecr/ecriface"
	"github.com/wallix/awless/logger"
)

type CreateRepository struct {
	_      string `action: "create" entity: "repository" awsAPI: "ecr" awsCall: "CreateRepository" awsInput: "CreateRepositoryInput" awsOutput: "CreateRepositoryOutput"`
	logger *logger.Logger
	api    ecriface.ECRAPI
	Name   *string `awsName: "RepositoryName" awsType: "awsstr" templateName: "name" required: ""`
}
type DeleteRepository struct {
	_       string `action: "delete" entity: "repository" awsAPI: "ecr" awsCall: "DeleteRepository" awsInput: "DeleteRepositoryInput" awsOutput: "DeleteRepositoryOutput"`
	logger  *logger.Logger
	api     ecriface.ECRAPI
	Name    *string `awsName: "RepositoryName" awsType: "awsstr" templateName: "name" required: ""`
	Force   *bool   `awsName: "Force" awsType: "awsbool" templateName: "force"`
	Account *string `awsName: "RegistryId" awsType: "awsstr" templateName: "account"`
}
