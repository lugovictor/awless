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
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func generateNewDrivers() {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, DRIVERS_DIR, func(os.FileInfo) bool { return true }, 0)
	if err != nil {
		panic(err)
	}

	finder := &findStructs{}
	for _, pkg := range pkgs {
		for _, f := range pkg.Files {
			ast.Walk(finder, f)
		}
	}

	templ, err := template.New("cmdRuns").Parse(cmdRuns)
	if err != nil {
		panic(err)
	}

	var buff bytes.Buffer
	err = templ.Execute(&buff, finder.result)
	if err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(filepath.Join(DRIVERS_DIR, "gen_cmd_runs.go"), buff.Bytes(), 0666); err != nil {
		panic(err)
	}
}

type cmdData struct {
	Call, Input, Output string
}

type findStructs struct {
	result map[string]cmdData
}

func (v *findStructs) Visit(node ast.Node) (w ast.Visitor) {
	if v.result == nil {
		v.result = make(map[string]cmdData)
	}
	if typ, ok := node.(*ast.TypeSpec); ok {
		if s, isStruct := typ.Type.(*ast.StructType); isStruct {
			for _, f := range s.Fields.List {
				if tag := f.Tag; tag != nil && strings.Contains(tag.Value, "awsCall") {
					key := typ.Name.Name
					v.result[key] = extractTag(tag.Value)
					break
				}
			}
		}
	}
	return v
}

func extractTag(s string) (t cmdData) {
	splits := strings.Split(s[1:len(s)-1], " ")
	for i, e := range splits {
		el := strings.Split(e, ":")
		ell := el[1][1 : len(el[1])-1]
		switch {
		case i == 0:
			t.Call = ell
		case i == 1:
			t.Input = ell
		case i == 2:
			t.Output = ell
		}
	}
	return
}

const cmdRuns = `/* Copyright 2017 WALLIX

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
	"github.com/wallix/awless/template"
)

{{- range $cmdName, $tag := . }}
func (cmd *{{ $cmdName }}) Run() (interface{}, error) {
	input := &{{ $tag.Input }}{}
	start := time.Now()
	output, err := cmd.api.{{ $tag.Call }}(input)
	if err != nil {
		return nil, fmt.Errorf("{{ $cmdName }}: %s", err)
	}
	cmd.logger.ExtraVerbosef("{{ $tag.Call }} call took %s", time.Since(start))
	cmd.result = aws.StringValue(output.Instances[0].InstanceId)
	return cmd.result, nil
}
{{- end }}
`
