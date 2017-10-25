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
	"github.com/aws/aws-sdk-go/service/ec2"
)

func init() {
	genTestsParams["createinternetgateway"] = map[string]interface{}{}
	genTestsExpected["createinternetgateway"] = &ec2.CreateInternetGatewayInput{}
	genTestsOutputExtractFunc["createinternetgateway"] = func() interface{} {
		return &ec2.CreateInternetGatewayOutput{InternetGateway: &ec2.InternetGateway{InternetGatewayId: String("id-my-igw")}}
	}
	genTestsOutput["createinternetgateway"] = "id-my-igw"

	genTestsParams["deleteinternetgateway"] = map[string]interface{}{
		"id": "my-igw-id",
	}
	genTestsExpected["deleteinternetgateway"] = &ec2.DeleteInternetGatewayInput{
		InternetGatewayId: String("my-igw-id"),
	}

	genTestsParams["attachinternetgateway"] = map[string]interface{}{
		"id":  "my-igw-id",
		"vpc": "my-vpc-id",
	}
	genTestsExpected["attachinternetgateway"] = &ec2.AttachInternetGatewayInput{
		InternetGatewayId: String("my-igw-id"),
		VpcId:             String("my-vpc-id"),
	}

	genTestsParams["detachinternetgateway"] = map[string]interface{}{
		"id":  "my-igw-id",
		"vpc": "my-vpc-id",
	}
	genTestsExpected["detachinternetgateway"] = &ec2.DetachInternetGatewayInput{
		InternetGatewayId: String("my-igw-id"),
		VpcId:             String("my-vpc-id"),
	}
}
