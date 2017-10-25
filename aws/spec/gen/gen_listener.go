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
	"github.com/aws/aws-sdk-go/service/elbv2/elbv2iface"
	"github.com/wallix/awless/logger"
)

type CreateListener struct {
	_            string `action: "create" entity: "listener" awsAPI: "elbv2" awsCall: "CreateListener" awsInput: "CreateListenerInput" awsOutput: "CreateListenerOutput"`
	logger       *logger.Logger
	api          elbv2iface.ELBV2API
	Actiontype   *struct{} `awsName: "DefaultActions[0]Type" awsType: "awsslicestruct" templateName: "actiontype" required: ""`
	Targetgroup  *struct{} `awsName: "DefaultActions[0]TargetGroupArn" awsType: "awsslicestruct" templateName: "targetgroup" required: ""`
	Loadbalancer *string   `awsName: "LoadBalancerArn" awsType: "awsstr" templateName: "loadbalancer" required: ""`
	Port         *int64    `awsName: "Port" awsType: "awsint64" templateName: "port" required: ""`
	Protocol     *string   `awsName: "Protocol" awsType: "awsstr" templateName: "protocol" required: ""`
	Certificate  *struct{} `awsName: "Certificates[0]CertificateArn" awsType: "awsslicestruct" templateName: "certificate"`
	Sslpolicy    *string   `awsName: "SslPolicy" awsType: "awsstr" templateName: "sslpolicy"`
}
type DeleteListener struct {
	_      string `action: "delete" entity: "listener" awsAPI: "elbv2" awsCall: "DeleteListener" awsInput: "DeleteListenerInput" awsOutput: "DeleteListenerOutput"`
	logger *logger.Logger
	api    elbv2iface.ELBV2API
	Id     *string `awsName: "ListenerArn" awsType: "awsstr" templateName: "id" required: ""`
}
