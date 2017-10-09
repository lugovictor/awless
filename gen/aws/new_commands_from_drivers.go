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

package aws

type NewCommandParam struct {
	Name, AwsName, AwsType string
	IsRequired             bool
}

type NewCommand struct {
	Action, Entity, API, Call, Input, Output, OutputExtractor string
	DryRun                                                    bool
	Params                                                    []NewCommandParam
}

func BuildNewCommandsFromDriversDeprecatedDefs(defs []driversDef) map[string][]NewCommand {
	commands := make(map[string][]NewCommand)
	for _, def := range defs {
		for _, driver := range def.Drivers {
			var params []NewCommandParam
			for _, p := range driver.RequiredParams {
				params = append(params, NewCommandParam{
					Name:       p.TemplateName,
					AwsName:    p.AwsField,
					AwsType:    p.AwsType,
					IsRequired: true,
				})
			}
			for _, p := range driver.ExtraParams {
				params = append(params, NewCommandParam{
					Name:       p.TemplateName,
					AwsName:    p.AwsField,
					AwsType:    p.AwsType,
					IsRequired: false,
				})
			}
			if driver.ManualFuncDefinition {
				commands[driver.Entity] = append(commands[driver.Entity], NewCommand{
					Action: driver.Action,
					Entity: driver.Entity,
					API:    def.Api,
					Params: params,
				})
			} else {
				commands[driver.Entity] = append(commands[driver.Entity], NewCommand{
					Action:          driver.Action,
					Entity:          driver.Entity,
					API:             def.Api,
					Call:            driver.ApiMethod,
					Input:           driver.Input,
					Output:          driver.Output,
					OutputExtractor: driver.OutputExtractor,
					DryRun:          !driver.DryRunUnsupported,
					Params:          params,
				})
			}
		}
	}
	return commands
}
