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
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
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
	fmt.Printf("%#v\n", finder.result)
}

type tag struct {
	call, input, output string
}

type findStructs struct {
	result map[string][]tag
}

func (v *findStructs) Visit(node ast.Node) (w ast.Visitor) {
	if v.result == nil {
		v.result = make(map[string][]tag)
	}
	if typ, ok := node.(*ast.TypeSpec); ok {
		if s, isStruct := typ.Type.(*ast.StructType); isStruct {
			for _, f := range s.Fields.List {
				if tag := f.Tag; tag != nil && strings.Contains(tag.Value, "awsCall") {
					key := typ.Name.Name
					fmt.Println(key)
					v.result[key] = append(v.result[key], extractTag(tag.Value))
				}
			}
		}
	}
	return v
}

func extractTag(s string) (t tag) {
	splits := strings.Split(s[1:len(s)-1], " ")
	for i, e := range splits {
		el := strings.Split(e, ":")
		ell := el[1][1 : len(el[1])-1]
		switch {
		case i == 0:
			t.call = ell
		case i == 1:
			t.input = ell
		case i == 2:
			t.output = ell
		}
	}
	return
}

const cmdRun = `/* Copyright 2017 WALLIX

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


var APIPerTemplateDefName = map[string]string {
{{- range $, $service := . }}
  {{- range $, $def := $service.Drivers }}
  "{{ $def.Action }}{{ $def.Entity }}": "{{ $service.Api }}",
  {{- end }}
{{- end }}
}

var AWSTemplatesDefinitions = map[string]template.Definition{
{{- range $, $service := . }}
{{- range $index, $def := $service.Drivers }}
	"{{ $def.Action }}{{ $def.Entity }}": template.Definition{
			Action: "{{ $def.Action }}",
			Entity: "{{ $def.Entity }}",
			Api: "{{ $service.Api }}",
			RequiredParams: []string{ {{- range $key := $def.RequiredKeys }}"{{ $key }}", {{- end}} },
			ExtraParams: []string{ {{- range $key := $def.ExtraKeys }}"{{ $key }}", {{- end}} },
		},
{{- end }}
{{- end }}
}

func DriverSupportedActions() map[string][]string { 
	supported := make(map[string][]string)
{{- range $, $service := . }}
{{- range $index, $def := $service.Drivers }}
	supported["{{ $def.Action }}"] = append(supported["{{ $def.Action }}"], "{{ $def.Entity }}")
{{- end }}
{{- end }}
	return supported
}
`
