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

package awsspec

import (
	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/elbv2/elbv2iface"
	"github.com/wallix/awless/logger"
)

type CreateListener struct {
	_            string `action:"create" entity:"listener" awsAPI:"elbv2" awsCall:"CreateListener" awsInput:"elbv2.CreateListenerInput" awsOutput:"elbv2.CreateListenerOutput"`
	logger       *logger.Logger
	api          elbv2iface.ELBV2API
	Actiontype   *string `awsName:"DefaultActions[0]Type" awsType:"awsslicestruct" templateName:"actiontype" required:""`
	Targetgroup  *string `awsName:"DefaultActions[0]TargetGroupArn" awsType:"awsslicestruct" templateName:"targetgroup" required:""`
	Loadbalancer *string `awsName:"LoadBalancerArn" awsType:"awsstr" templateName:"loadbalancer" required:""`
	Port         *int64  `awsName:"Port" awsType:"awsint64" templateName:"port" required:""`
	Protocol     *string `awsName:"Protocol" awsType:"awsstr" templateName:"protocol" required:""`
	Certificate  *string `awsName:"Certificates[0]CertificateArn" awsType:"awsslicestruct" templateName:"certificate"`
	Sslpolicy    *string `awsName:"SslPolicy" awsType:"awsstr" templateName:"sslpolicy"`
}

func (cmd *CreateListener) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}

func (cmd *CreateListener) ExtractResult(i interface{}) string {
	return awssdk.StringValue(i.(*elbv2.CreateListenerOutput).Listeners[0].ListenerArn)
}

type DeleteListener struct {
	_      string `action:"delete" entity:"listener" awsAPI:"elbv2" awsCall:"DeleteListener" awsInput:"elbv2.DeleteListenerInput" awsOutput:"elbv2.DeleteListenerOutput"`
	logger *logger.Logger
	api    elbv2iface.ELBV2API
	Id     *string `awsName:"ListenerArn" awsType:"awsstr" templateName:"id" required:""`
}

func (cmd *DeleteListener) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}
