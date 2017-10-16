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
package awsdriver

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/wallix/awless/logger"
)

var NewCommandFuncs = map[string]func() interface{}{}

func InitCommands(l *logger.Logger, sess *session.Session) {
	NewCommandFuncs["createinstance"] = func() interface{} { return NewCreateInstance(l, sess) }
	NewCommandFuncs["createsubnet"] = func() interface{} { return NewCreateSubnet(l, sess) }
	NewCommandFuncs["createtag"] = func() interface{} { return NewCreateTag(l, sess) }
}
