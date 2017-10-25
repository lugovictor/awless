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

type CreateVolume struct {
	_                string `action: "create" entity: "volume" awsAPI: "ec2" awsCall: "CreateVolume" awsInput: "CreateVolumeInput" awsOutput: "Volume" awsDryRun: ""`
	logger           *logger.Logger
	api              ec2iface.EC2API
	Availabilityzone *string `awsName: "AvailabilityZone" awsType: "awsstr" templateName: "availabilityzone" required: ""`
	Size             *int64  `awsName: "Size" awsType: "awsint64" templateName: "size" required: ""`
}
type CheckVolume struct {
	_       string `action: "check" entity: "volume" awsAPI: "ec2"`
	logger  *logger.Logger
	api     ec2iface.EC2API
	Id      *struct{} `templateName: "id" required: ""`
	State   *struct{} `templateName: "state" required: ""`
	Timeout *struct{} `templateName: "timeout" required: ""`
}
type DeleteVolume struct {
	_      string `action: "delete" entity: "volume" awsAPI: "ec2" awsCall: "DeleteVolume" awsInput: "DeleteVolumeInput" awsOutput: "DeleteVolumeOutput" awsDryRun: ""`
	logger *logger.Logger
	api    ec2iface.EC2API
	Id     *string `awsName: "VolumeId" awsType: "awsstr" templateName: "id" required: ""`
}
type AttachVolume struct {
	_        string `action: "attach" entity: "volume" awsAPI: "ec2" awsCall: "AttachVolume" awsInput: "AttachVolumeInput" awsOutput: "VolumeAttachment" awsDryRun: ""`
	logger   *logger.Logger
	api      ec2iface.EC2API
	Device   *string `awsName: "Device" awsType: "awsstr" templateName: "device" required: ""`
	Id       *string `awsName: "VolumeId" awsType: "awsstr" templateName: "id" required: ""`
	Instance *string `awsName: "InstanceId" awsType: "awsstr" templateName: "instance" required: ""`
}
type DetachVolume struct {
	_        string `action: "detach" entity: "volume" awsAPI: "ec2" awsCall: "DetachVolume" awsInput: "DetachVolumeInput" awsOutput: "VolumeAttachment" awsDryRun: ""`
	logger   *logger.Logger
	api      ec2iface.EC2API
	Device   *string `awsName: "Device" awsType: "awsstr" templateName: "device" required: ""`
	Id       *string `awsName: "VolumeId" awsType: "awsstr" templateName: "id" required: ""`
	Instance *string `awsName: "InstanceId" awsType: "awsstr" templateName: "instance" required: ""`
	Force    *bool   `awsName: "Force" awsType: "awsbool" templateName: "force"`
}
