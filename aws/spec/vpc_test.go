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
	genTestsParams["createvpc"] = map[string]interface{}{
		"name": "myvpc",
		"cidr": "10.0.0.0/16",
	}
	genTestsExpected["createvpc"] = &ec2.CreateVpcInput{
		CidrBlock: String("10.0.0.0/16"),
	}
	genTestsOutputExtractFunc["createvpc"] = func() interface{} {
		return &ec2.CreateVpcOutput{Vpc: &ec2.Vpc{VpcId: String("res.createvpc")}}
	}
	genTestsOutput["createvpc"] = "res.createvpc"
	genTestsExpected["createvpc.createtag"] = &ec2.CreateTagsInput{
		Resources: []*string{String("res.createvpc")},
		Tags: []*ec2.Tag{
			{Key: String("Name"), Value: String("myvpc")},
		},
	}

	genTestsParams["deletevpc"] = map[string]interface{}{
		"id": "my-vpc-id",
	}
	genTestsExpected["deletevpc"] = &ec2.DeleteVpcInput{
		VpcId: String("my-vpc-id"),
	}
}
