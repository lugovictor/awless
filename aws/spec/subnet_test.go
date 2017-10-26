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
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func init() {
	genTestsParams["createsubnet"] = map[string]interface{}{
		"cidr":             "10.10.10.10/24",
		"vpc":              "my-vpc-id",
		"availabilityzone": "eu-west-1a",
		"name":             "mysubnet",
	}
	genTestsExpected["createsubnet"] = &ec2.CreateSubnetInput{
		AvailabilityZone: String("eu-west-1a"),
		CidrBlock:        String("10.10.10.10/24"),
		VpcId:            String("my-vpc-id"),
	}
	genTestsOutputExtractFunc["createsubnet"] = func() interface{} {
		return &ec2.CreateSubnetOutput{Subnet: &ec2.Subnet{SubnetId: String("id-my-subnet")}}
	}
	genTestsOutput["createsubnet"] = "id-my-subnet"

	genTestsParams["updatesubnet"] = map[string]interface{}{
		"id":     "my-subnet-id",
		"public": "true",
	}
	genTestsExpected["updatesubnet"] = &ec2.ModifySubnetAttributeInput{
		MapPublicIpOnLaunch: &ec2.AttributeBooleanValue{Value: Bool(true)},
		SubnetId:            String("my-subnet-id"),
	}

	genTestsParams["deletesubnet"] = map[string]interface{}{
		"id": "my-subnet-id",
	}
	genTestsExpected["deletesubnet"] = &ec2.DeleteSubnetInput{
		SubnetId: String("my-subnet-id"),
	}
}

func TestCreateSubnet(t *testing.T) {
	create := &CreateSubnet{}
	params := map[string]interface{}{
		"cidr": "10.10.10.10/24",
		"vpc":  "vpc-1234",
	}
	t.Run("Validate", func(t *testing.T) {
		params["cidr"] = "10.10.10.10/128"
		errs := create.ValidateCommand(params)
		checkErrs(t, errs, 1, "cidr")

	})
}
