/*
Copyright 2017 WALLIX

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

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/wallix/awless/gen/aws"
)

func generateNewCommands() {
	cmdsByEntity := aws.BuildNewCommandsFromDriversDeprecatedDefs(aws.DriversDefs)
	for entity, cmds := range cmdsByEntity {
		templ, err := template.New("new_commands").Funcs(template.FuncMap{
			"GenerateCmdTag":           GenerateCmdTag,
			"GenerateParamTag":         GenerateParamTag,
			"ApiToInterface":           aws.ApiToInterface,
			"CommandName":              CommandName,
			"AwsTypeToGoType":          AwsTypeToGoType,
			"ToParamName":              ToParamName,
			"GenerateResultExtraction": GenerateResultExtraction,
		}).Parse(generatedNewCommands)
		if err != nil {
			panic(err)
		}

		var buff bytes.Buffer
		err = templ.Execute(&buff, cmds)
		if err != nil {
			panic(err)
		}

		if err := ioutil.WriteFile(filepath.Join(GEN_SPEC_DIR, fmt.Sprintf("gen_%s.go.txt", entity)), buff.Bytes(), 0666); err != nil {
			panic(err)
		}
	}
}

func ToParamName(name string) string {
	clean := strings.Replace(strings.Title(name), "-", "", -1)
	if strings.Contains(clean, ".") {
		splits := strings.Split(clean, ".")
		return splits[len(splits)-1]
	}
	return clean
}

func CommandName(c aws.NewCommand) string {
	return strings.Title(c.Action) + strings.Title(c.Entity)
}

func GenerateCmdTag(c aws.NewCommand) (string, error) {
	var tags []string
	if c.Action != "" {
		tags = append(tags, fmt.Sprintf("action:\"%s\"", c.Action))
	}
	if c.Entity != "" {
		tags = append(tags, fmt.Sprintf("entity:\"%s\"", c.Entity))
	}
	if c.API != "" {
		tags = append(tags, fmt.Sprintf("awsAPI:\"%s\"", c.API))
	}
	if c.Call != "" {
		tags = append(tags, fmt.Sprintf("awsCall:\"%s\"", c.Call))
	}
	if c.Input != "" {
		tags = append(tags, fmt.Sprintf("awsInput:\"%s.%s\"", c.API, c.Input))
	}
	if c.Output != "" {
		tags = append(tags, fmt.Sprintf("awsOutput:\"%s.%s\"", c.API, c.Output))
	}
	if c.DryRun {
		tags = append(tags, "awsDryRun:\"\"")
	}
	if len(tags) > 0 {
		return "`" + strings.Join(tags, " ") + "`", nil
	}
	return "", nil
}

func GenerateParamTag(p aws.NewCommandParam) (string, error) {
	var tags []string
	if p.AwsName != "" {
		tags = append(tags, fmt.Sprintf("awsName:\"%s\"", p.AwsName))
	}
	if p.AwsType != "" {
		tags = append(tags, fmt.Sprintf("awsType:\"%s\"", p.AwsType))
	}
	if p.Name != "" {
		tags = append(tags, fmt.Sprintf("templateName:\"%s\"", p.Name))
	}
	if p.IsRequired {
		tags = append(tags, "required:\"\"")
	}
	if len(tags) > 0 {
		return "`" + strings.Join(tags, " ") + "`", nil
	}
	return "", nil
}

func AwsTypeToGoType(awstype string) (string, error) {
	switch awstype {
	case "", "awsslicestruct", "awsstringpointermap", "awsstepadjustments", "awsparameterslice", "awssubnetmappings", "awsdimensionslice":
		return "*struct{}", nil
	case "awsstr", "awsfiletostring", "aws6digitsstring":
		return "*string", nil
	case "awsboolattribute", "awsbool":
		return "*bool", nil
	case "awsfiletobase64":
		return "*string", nil
	case "awsfiletobyteslice":
		return "*[]byte", nil
	case "awsstringslice", "awscsvstr":
		return "*[]string", nil
	case "awsint64", "awsslicestructint64":
		return "*int64", nil
	case "awsfloat":
		return "*float64", nil

	default:
		return "", fmt.Errorf("unknown awstype '%s'", awstype)
	}
}

func GenerateResultExtraction(cmd aws.NewCommand) string {
	out := strings.Replace(cmd.OutputExtractor, "aws.", "awssdk.", 1)
	out = strings.Replace(out, "output.", fmt.Sprintf("i.(*%s.%s).", cmd.API, cmd.Output), 1)

	return out
}

const generatedNewCommands = `// +build ignore
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
	"github.com/wallix/awless/template"
	awssdk "github.com/aws/aws-sdk-go/aws"
)

{{- range $, $cmd := . }}
type {{ CommandName $cmd }} struct {
  _              string {{ GenerateCmdTag $cmd }}
  logger         *logger.Logger
  api            {{ $cmd.API }}iface.{{ ApiToInterface $cmd.API }}
  {{- range $, $param := $cmd.Params }}
    {{ToParamName $param.Name }} {{ AwsTypeToGoType $param.AwsType }} {{ GenerateParamTag $param }}
  {{- end }}
}
{{ if $cmd.Call }}
func (cmd *{{ CommandName $cmd }}) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}
{{ end }}
{{ if $cmd.OutputExtractor }}
func (cmd *{{ CommandName $cmd }}) ExtractResult(i interface{}) string {
	return {{ GenerateResultExtraction $cmd }}
}
{{ end }}

{{- end }}
`
