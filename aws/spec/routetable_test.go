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
	genTestsParams["createroutetable"] = map[string]interface{}{
		"vpc": "my-vpc-id",
	}
	genTestsExpected["createroutetable"] = &ec2.CreateRouteTableInput{
		VpcId: String("my-vpc-id"),
	}
	genTestsOutputExtractFunc["createroutetable"] = func() interface{} {
		return &ec2.CreateRouteTableOutput{RouteTable: &ec2.RouteTable{RouteTableId: String("id-my-rt")}}
	}
	genTestsOutput["createroutetable"] = "id-my-rt"

	genTestsParams["deleteroutetable"] = map[string]interface{}{
		"id": "my-rt-id",
	}
	genTestsExpected["deleteroutetable"] = &ec2.DeleteRouteTableInput{
		RouteTableId: String("my-rt-id"),
	}

	genTestsParams["attachroutetable"] = map[string]interface{}{
		"id":     "my-rt-id",
		"subnet": "my-subnet-id",
	}
	genTestsExpected["attachroutetable"] = &ec2.AssociateRouteTableInput{
		RouteTableId: String("my-rt-id"),
		SubnetId:     String("my-subnet-id"),
	}
	genTestsOutputExtractFunc["attachroutetable"] = func() interface{} {
		return &ec2.AssociateRouteTableOutput{AssociationId: String("id-my-assoc")}
	}
	genTestsOutput["attachroutetable"] = "id-my-assoc"

	genTestsParams["detachroutetable"] = map[string]interface{}{
		"association": "my-assoc-id",
	}
	genTestsExpected["detachroutetable"] = &ec2.DisassociateRouteTableInput{
		AssociationId: String("my-assoc-id"),
	}
}
