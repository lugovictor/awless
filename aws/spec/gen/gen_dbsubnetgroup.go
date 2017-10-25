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
	"github.com/aws/aws-sdk-go/service/rds/rdsiface"
	"github.com/wallix/awless/logger"
)

type CreateDbsubnetgroup struct {
	_           string `action: "create" entity: "dbsubnetgroup" awsAPI: "rds" awsCall: "CreateDBSubnetGroup" awsInput: "CreateDBSubnetGroupInput" awsOutput: "CreateDBSubnetGroupOutput"`
	logger      *logger.Logger
	api         rdsiface.RDSAPI
	Description *string   `awsName: "DBSubnetGroupDescription" awsType: "awsstr" templateName: "description" required: ""`
	Name        *string   `awsName: "DBSubnetGroupName" awsType: "awsstr" templateName: "name" required: ""`
	Subnets     *[]string `awsName: "SubnetIds" awsType: "awsstringslice" templateName: "subnets" required: ""`
}
type DeleteDbsubnetgroup struct {
	_      string `action: "delete" entity: "dbsubnetgroup" awsAPI: "rds" awsCall: "DeleteDBSubnetGroup" awsInput: "DeleteDBSubnetGroupInput" awsOutput: "DeleteDBSubnetGroupOutput"`
	logger *logger.Logger
	api    rdsiface.RDSAPI
	Name   *string `awsName: "DBSubnetGroupName" awsType: "awsstr" templateName: "name" required: ""`
}
