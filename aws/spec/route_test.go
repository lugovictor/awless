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
	genTestsParams["createroute"] = map[string]interface{}{
		"table":   "table-id",
		"cidr":    "10.0.0.0/16",
		"gateway": "igw-id",
	}
	genTestsExpected["createroute"] = &ec2.CreateRouteInput{
		RouteTableId:         String("table-id"),
		DestinationCidrBlock: String("10.0.0.0/16"),
		GatewayId:            String("igw-id"),
	}

	genTestsParams["deleteroute"] = map[string]interface{}{
		"table": "table-id",
		"cidr":  "10.0.0.0/16",
	}
	genTestsExpected["deleteroute"] = &ec2.DeleteRouteInput{
		RouteTableId:         String("table-id"),
		DestinationCidrBlock: String("10.0.0.0/16"),
	}
}
