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
	"github.com/aws/aws-sdk-go/service/iam"
)

func init() {
	genTestsParams["creategroup"] = map[string]interface{}{
		"name": "my-group",
	}
	genTestsExpected["creategroup"] = &iam.CreateGroupInput{
		GroupName: String("my-group"),
	}
	genTestsOutputExtractFunc["creategroup"] = func() interface{} {
		return &iam.CreateGroupOutput{Group: &iam.Group{GroupId: String("group-id")}}
	}
	genTestsOutput["creategroup"] = "group-id"

	genTestsParams["deletegroup"] = map[string]interface{}{
		"name": "my-group",
	}
	genTestsExpected["deletegroup"] = &iam.DeleteGroupInput{
		GroupName: String("my-group"),
	}
}
