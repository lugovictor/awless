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
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/wallix/awless/gen/aws"
)

func generateNewDriversTests() {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, SPEC_DIR, func(os.FileInfo) bool { return true }, 0)
	if err != nil {
		panic(err)
	}

	finder := &findStructs{}
	for _, pkg := range pkgs {
		for _, f := range pkg.Files {
			ast.Walk(finder, f)
		}
	}

	templ, err := template.New("cmdTests").Funcs(template.FuncMap{
		"Title":          strings.Title,
		"ToLower":        strings.ToLower,
		"ApiToInterface": aws.ApiToInterface,
	}).Parse(cmdTests)
	if err != nil {
		panic(err)
	}

	cmdsByApi := make(map[string]map[string]cmdData)
	for k, cmd := range finder.result {
		if cmd.API == "" {
			continue
		}
		if len(cmdsByApi[cmd.API]) == 0 {
			cmdsByApi[cmd.API] = make(map[string]cmdData)
		}
		cmdsByApi[cmd.API][k] = cmd
	}

	var buff bytes.Buffer
	err = templ.Execute(&buff, cmdsByApi)
	if err != nil {
		panic(err)
	}

	if err = ioutil.WriteFile(filepath.Join(SPEC_DIR, "gen_test.go"), buff.Bytes(), 0666); err != nil {
		panic(err)
	}
}

const cmdTests = `/* Copyright 2017 WALLIX

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

var (
	genTestsParams = make(map[string]map[string]interface{})
	genTestsContext = make(map[string]map[string]interface{})
	genTestsExpected = make(map[string]interface{})
	genTestsOutput = make(map[string]interface{})
	genTestsCleanupFunc = make(map[string]func())
	genTestsOutputExtractFunc = make(map[string]func() interface{})
	{{ range $api, $cmds := . }}
		{{- range $cmdName, $cmd := . }}
			new{{ $cmdName }} = func() *{{ $cmdName }} {return &{{ $cmdName }}{api: &mock{{ Title $api}}{}, logger: logger.DiscardLogger}}
		{{- end }}
	{{- end }}
)

func init() {
	{{ range $api, $cmds := . }}
		{{- range $cmdName, $cmd := . }}
			NewCommandFuncs["{{ ToLower $cmdName }}"] = func() interface{} { return new{{ $cmdName }}() }
		{{- end }}
	{{- end }}
}

{{ range $api, $cmds := . }}
type mock{{ Title $api}} struct {
	{{$api}}iface.{{ ApiToInterface $api}}
}

	{{- range $cmdName, $cmd := . }}
		func TestGen{{ $cmdName }}(t *testing.T) {
			if cleanFn, ok := genTestsCleanupFunc["{{ ToLower $cmdName }}"]; ok {
				defer cleanFn()
			}
			params := genTestsParams["{{ ToLower $cmdName }}"]
			missings, err := new{{ $cmdName }}().ValidateParams(keys(params))
			if err != nil {
				t.Fatal(err)
			}
			if got, want := len(missings), 0; got != want {
				t.Fatalf("got %d, want %d", got, want)
			}
			if params, err = convertParamsIfAvailable(new{{ $cmdName }}(), params); err != nil {
				t.Fatal(err)
			}
			if errs := new{{ $cmdName }}().ValidateCommand(genTestsParams["{{ ToLower $cmdName }}"]); len(errs) > 0 {
				t.Fatalf("%v", errs)
			}
			res, err := new{{ $cmdName }}().Run(genTestsContext["{{ ToLower $cmdName }}"], genTestsParams["{{ ToLower $cmdName }}"])
			if err != nil {
				t.Fatal(err)
			}
			if got, want := res, genTestsOutput["{{ ToLower $cmdName }}"]; ! reflect.DeepEqual(got, want) {
				t.Fatalf("got %+v, want %+v", got, want)
			}
		}

	{{- end }}
{{- end }}

{{ range $api, $cmds := . }}
	{{- range $cmdName, $cmd := . }}
		{{ if $cmd.Call }}
		func (m *mock{{ Title $api}}) {{ $cmd.Call }}(input *{{ $cmd.Input }}) (*{{ $cmd.Output }}, error) {
			if got, want := input, genTestsExpected["{{ ToLower $cmdName }}"]; !reflect.DeepEqual(got, want) {
				return nil, fmt.Errorf("got %#v, want %#v", got, want)
			}
			if outFunc, ok := genTestsOutputExtractFunc["{{ ToLower $cmdName }}"];ok{
				return outFunc().(*{{ $cmd.Output }}), nil
			}
			return nil, nil
		}
		{{ end }}
	{{- end }}
{{- end }}
`
