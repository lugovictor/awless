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
	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/rds/rdsiface"
	"github.com/wallix/awless/logger"
)

type CreateDatabase struct {
	_                 string `action:"create" entity:"database" awsAPI:"rds" awsCall:"CreateDBInstance" awsInput:"rds.CreateDBInstanceInput" awsOutput:"rds.CreateDBInstanceOutput"`
	logger            *logger.Logger
	api               rdsiface.RDSAPI
	Type              *string   `awsName:"DBInstanceClass" awsType:"awsstr" templateName:"type" required:""`
	Id                *string   `awsName:"DBInstanceIdentifier" awsType:"awsstr" templateName:"id" required:""`
	Engine            *string   `awsName:"Engine" awsType:"awsstr" templateName:"engine" required:""`
	Password          *string   `awsName:"MasterUserPassword" awsType:"awsstr" templateName:"password" required:""`
	Username          *string   `awsName:"MasterUsername" awsType:"awsstr" templateName:"username" required:""`
	Size              *int64    `awsName:"AllocatedStorage" awsType:"awsint64" templateName:"size" required:""`
	Autoupgrade       *bool     `awsName:"AutoMinorVersionUpgrade" awsType:"awsbool" templateName:"autoupgrade"`
	Availabilityzone  *string   `awsName:"AvailabilityZone" awsType:"awsstr" templateName:"availabilityzone"`
	Backupretention   *int64    `awsName:"BackupRetentionPeriod" awsType:"awsint64" templateName:"backupretention"`
	Cluster           *string   `awsName:"DBClusterIdentifier" awsType:"awsstr" templateName:"cluster"`
	Dbname            *string   `awsName:"DBName" awsType:"awsstr" templateName:"dbname"`
	Parametergroup    *string   `awsName:"DBParameterGroupName" awsType:"awsstr" templateName:"parametergroup"`
	Dbsecuritygroups  *[]string `awsName:"DBSecurityGroups" awsType:"awsstringslice" templateName:"dbsecuritygroups"`
	Subnetgroup       *string   `awsName:"DBSubnetGroupName" awsType:"awsstr" templateName:"subnetgroup"`
	Domain            *string   `awsName:"Domain" awsType:"awsstr" templateName:"domain"`
	Iamrole           *string   `awsName:"DomainIAMRoleName" awsType:"awsstr" templateName:"iamrole"`
	Version           *string   `awsName:"EngineVersion" awsType:"awsstr" templateName:"version"`
	Iops              *int64    `awsName:"Iops" awsType:"awsint64" templateName:"iops"`
	License           *string   `awsName:"LicenseModel" awsType:"awsstr" templateName:"license"`
	Multiaz           *bool     `awsName:"MultiAZ" awsType:"awsbool" templateName:"multiaz"`
	Optiongroup       *string   `awsName:"OptionGroupName" awsType:"awsstr" templateName:"optiongroup"`
	Port              *int64    `awsName:"Port" awsType:"awsint64" templateName:"port"`
	Backupwindow      *string   `awsName:"PreferredBackupWindow" awsType:"awsstr" templateName:"backupwindow"`
	Maintenancewindow *string   `awsName:"PreferredMaintenanceWindow" awsType:"awsstr" templateName:"maintenancewindow"`
	Public            *bool     `awsName:"PubliclyAccessible" awsType:"awsbool" templateName:"public"`
	Encrypted         *bool     `awsName:"StorageEncrypted" awsType:"awsbool" templateName:"encrypted"`
	Storagetype       *string   `awsName:"StorageType" awsType:"awsstr" templateName:"storagetype"`
	Timezone          *string   `awsName:"Timezone" awsType:"awsstr" templateName:"timezone"`
	Vpcsecuritygroups *[]string `awsName:"VpcSecurityGroupIds" awsType:"awsstringslice" templateName:"vpcsecuritygroups"`
}

func (cmd *CreateDatabase) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}

func (cmd *CreateDatabase) ExtractResult(i interface{}) string {
	return awssdk.StringValue(i.(*rds.CreateDBInstanceOutput).DBInstance.DBInstanceIdentifier)
}

type DeleteDatabase struct {
	_            string `action:"delete" entity:"database" awsAPI:"rds" awsCall:"DeleteDBInstance" awsInput:"rds.DeleteDBInstanceInput" awsOutput:"rds.DeleteDBInstanceOutput"`
	logger       *logger.Logger
	api          rdsiface.RDSAPI
	Id           *string `awsName:"DBInstanceIdentifier" awsType:"awsstr" templateName:"id" required:""`
	SkipSnapshot *bool   `awsName:"SkipFinalSnapshot" awsType:"awsbool" templateName:"skip-snapshot"`
	Snapshot     *bool   `awsName:"FinalDBSnapshotIdentifier" awsType:"awsbool" templateName:"snapshot"`
}

func (cmd *DeleteDatabase) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}

type CheckDatabase struct {
	_       string `action:"check" entity:"database" awsAPI:"rds"`
	logger  *logger.Logger
	api     rdsiface.RDSAPI
	Id      *struct{} `templateName:"id" required:""`
	State   *struct{} `templateName:"state" required:""`
	Timeout *struct{} `templateName:"timeout" required:""`
}
