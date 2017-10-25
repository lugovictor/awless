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
	"github.com/aws/aws-sdk-go/service/ecs/ecsiface"
	"github.com/wallix/awless/logger"
)

type CreateContainercluster struct {
	_      string `action: "create" entity: "containercluster" awsAPI: "ecs" awsCall: "CreateCluster" awsInput: "CreateClusterInput" awsOutput: "CreateClusterOutput"`
	logger *logger.Logger
	api    ecsiface.ECSAPI
	Name   *string `awsName: "ClusterName" awsType: "awsstr" templateName: "name" required: ""`
}
type DeleteContainercluster struct {
	_      string `action: "delete" entity: "containercluster" awsAPI: "ecs" awsCall: "DeleteCluster" awsInput: "DeleteClusterInput" awsOutput: "DeleteClusterOutput"`
	logger *logger.Logger
	api    ecsiface.ECSAPI
	Id     *string `awsName: "Cluster" awsType: "awsstr" templateName: "id" required: ""`
}
