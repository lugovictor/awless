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
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/wallix/awless/logger"
)

type CreateTag struct {
	_        string `action:"create" entity:"tag" awsAPI:"ec2"`
	logger   *logger.Logger
	api      ec2iface.EC2API
	Resource *struct{} `templateName:"resource" required:""`
	Key      *struct{} `templateName:"key" required:""`
	Value    *struct{} `templateName:"value" required:""`
}

type DeleteTag struct {
	_        string `action:"delete" entity:"tag" awsAPI:"ec2"`
	logger   *logger.Logger
	api      ec2iface.EC2API
	Resource *struct{} `templateName:"resource" required:""`
	Key      *struct{} `templateName:"key" required:""`
	Value    *struct{} `templateName:"value" required:""`
}
