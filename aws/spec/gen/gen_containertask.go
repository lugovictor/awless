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

type StartContainertask struct {
	_              string `action: "start" entity: "containertask" awsAPI: "ecs"`
	logger         *logger.Logger
	api            ecsiface.ECSAPI
	Cluster        *struct{} `templateName: "cluster" required: ""`
	DesiredCount   *struct{} `templateName: "desired-count" required: ""`
	Name           *struct{} `templateName: "name" required: ""`
	Type           *struct{} `templateName: "type" required: ""`
	Role           *struct{} `templateName: "role"`
	DeploymentName *struct{} `templateName: "deployment-name"`
	ContainerName  *struct{} `templateName: "loadbalancer.container-name"`
	ContainerPort  *struct{} `templateName: "loadbalancer.container-port"`
	Targetgroup    *struct{} `templateName: "loadbalancer.targetgroup"`
}
type StopContainertask struct {
	_              string `action: "stop" entity: "containertask" awsAPI: "ecs"`
	logger         *logger.Logger
	api            ecsiface.ECSAPI
	Cluster        *struct{} `templateName: "cluster" required: ""`
	Type           *struct{} `templateName: "type" required: ""`
	DeploymentName *struct{} `templateName: "deployment-name"`
	RunArn         *struct{} `templateName: "run-arn"`
}
type UpdateContainertask struct {
	_              string `action: "update" entity: "containertask" awsAPI: "ecs" awsCall: "UpdateService" awsInput: "UpdateServiceInput" awsOutput: "UpdateServiceOutput"`
	logger         *logger.Logger
	api            ecsiface.ECSAPI
	Cluster        *string `awsName: "Cluster" awsType: "awsstr" templateName: "cluster" required: ""`
	DeploymentName *string `awsName: "Service" awsType: "awsstr" templateName: "deployment-name" required: ""`
	DesiredCount   *int64  `awsName: "DesiredCount" awsType: "awsint64" templateName: "desired-count"`
	Name           *string `awsName: "TaskDefinition" awsType: "awsstr" templateName: "name"`
}
type AttachContainertask struct {
	_               string `action: "attach" entity: "containertask" awsAPI: "ecs"`
	logger          *logger.Logger
	api             ecsiface.ECSAPI
	Name            *struct{} `templateName: "name" required: ""`
	ContainerName   *struct{} `templateName: "container-name" required: ""`
	Image           *struct{} `templateName: "image" required: ""`
	MemoryHardLimit *struct{} `templateName: "memory-hard-limit" required: ""`
	Command         *struct{} `templateName: "command"`
	Env             *struct{} `templateName: "env"`
	Privileged      *struct{} `templateName: "privileged"`
	Workdir         *struct{} `templateName: "workdir"`
	Ports           *struct{} `templateName: "ports"`
}
type DetachContainertask struct {
	_             string `action: "detach" entity: "containertask" awsAPI: "ecs"`
	logger        *logger.Logger
	api           ecsiface.ECSAPI
	Name          *struct{} `templateName: "name" required: ""`
	ContainerName *struct{} `templateName: "container-name" required: ""`
}
type DeleteContainertask struct {
	_           string `action: "delete" entity: "containertask" awsAPI: "ecs"`
	logger      *logger.Logger
	api         ecsiface.ECSAPI
	Name        *struct{} `templateName: "name" required: ""`
	AllVersions *struct{} `templateName: "all-versions"`
}
