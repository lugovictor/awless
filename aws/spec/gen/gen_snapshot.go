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

type CreateSnapshot struct {
	_           string `action: "create" entity: "snapshot" awsAPI: "ec2" awsCall: "CreateSnapshot" awsInput: "CreateSnapshotInput" awsOutput: "Snapshot" awsDryRun: ""`
	logger      *logger.Logger
	api         ec2iface.EC2API
	Volume      *string `awsName: "VolumeId" awsType: "awsstr" templateName: "volume" required: ""`
	Description *string `awsName: "Description" awsType: "awsstr" templateName: "description"`
}
type DeleteSnapshot struct {
	_      string `action: "delete" entity: "snapshot" awsAPI: "ec2" awsCall: "DeleteSnapshot" awsInput: "DeleteSnapshotInput" awsOutput: "DeleteSnapshotOutput" awsDryRun: ""`
	logger *logger.Logger
	api    ec2iface.EC2API
	Id     *string `awsName: "SnapshotId" awsType: "awsstr" templateName: "id" required: ""`
}
type CopySnapshot struct {
	_            string `action: "copy" entity: "snapshot" awsAPI: "ec2" awsCall: "CopySnapshot" awsInput: "CopySnapshotInput" awsOutput: "CopySnapshotOutput" awsDryRun: ""`
	logger       *logger.Logger
	api          ec2iface.EC2API
	SourceId     *string `awsName: "SourceSnapshotId" awsType: "awsstr" templateName: "source-id" required: ""`
	SourceRegion *string `awsName: "SourceRegion" awsType: "awsstr" templateName: "source-region" required: ""`
	Encrypted    *bool   `awsName: "Encrypted" awsType: "awsbool" templateName: "encrypted"`
	Description  *string `awsName: "Description" awsType: "awsstr" templateName: "description"`
}
