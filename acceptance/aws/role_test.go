package awsat

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/wallix/awless/logger"
)

func TestRole(t *testing.T) {
	/*
			Name             *string   `awsName:"RoleName" awsType:"awsstr" templateName:"name" required:""`
		PrincipalAccount *struct{} `templateName:"principal-account"`
		PrincipalUser    *struct{} `templateName:"principal-user"`
		PrincipalService *struct{} `templateName:"principal-service"`
		Conditions       *struct{} `templateName:"conditions"`
		SleepAfter       *struct{} `templateName:"sleep-after"`

	*/
	t.Run("create", func(t *testing.T) {
		Template(`create role name=president principal-account=account principal-user=user principal-service=service conditions="aws:SecureTransport==true","s3:max-keys==10" sleep-after=0`).Mock(&iamMock{
			CreateRoleFunc: func(input *iam.CreateRoleInput) (*iam.CreateRoleOutput, error) {
				return &iam.CreateRoleOutput{Role: &iam.Role{Arn: String("new-role-arn"), RoleId: String("new-role-id"), RoleName: String("president")}}, nil
			},
			CreateInstanceProfileFunc: func(input *iam.CreateInstanceProfileInput) (*iam.CreateInstanceProfileOutput, error) {
				return nil, nil
			},
			AddRoleToInstanceProfileFunc: func(input *iam.AddRoleToInstanceProfileInput) (*iam.AddRoleToInstanceProfileOutput, error) {
				return &iam.AddRoleToInstanceProfileOutput{}, nil
			}}).ExpectInput("AddRoleToInstanceProfile", &iam.AddRoleToInstanceProfileInput{
			InstanceProfileName: String("president"),
			RoleName:            String("president"),
		}).ExpectInput("CreateInstanceProfile", &iam.CreateInstanceProfileInput{
			InstanceProfileName: String("president"),
		}).ExpectInput("CreateRole", &iam.CreateRoleInput{
			RoleName: String("president"),
			AssumeRolePolicyDocument: String(`{
 "Version": "2012-10-17",
 "Statement": [
  {
   "Effect": "Allow",
   "Action": [
    "sts:AssumeRole"
   ],
   "Principal": {
    "AWS": "account"
   },
   "Condition": {
    "Bool": {
     "aws:SecureTransport": "true"
    },
    "NumericEquals": {
     "s3:max-keys": "10"
    }
   }
  }
 ]
}`),
		}).ExpectCommandResult("new-role-arn").ExpectCalls("CreateRole", "CreateInstanceProfile", "AddRoleToInstanceProfile").Run(t, logger.DefaultLogger)
	})

	t.Run("delete", func(t *testing.T) {
		Template(`delete role name=president`).Mock(&iamMock{
			RemoveRoleFromInstanceProfileFunc: func(input *iam.RemoveRoleFromInstanceProfileInput) (*iam.RemoveRoleFromInstanceProfileOutput, error) {
				return &iam.RemoveRoleFromInstanceProfileOutput{}, nil
			},
			DeleteInstanceProfileFunc: func(input *iam.DeleteInstanceProfileInput) (*iam.DeleteInstanceProfileOutput, error) {
				return &iam.DeleteInstanceProfileOutput{}, nil
			},
			DeleteRoleFunc: func(input *iam.DeleteRoleInput) (*iam.DeleteRoleOutput, error) {
				return nil, nil
			}}).ExpectInput("RemoveRoleFromInstanceProfile", &iam.RemoveRoleFromInstanceProfileInput{
			InstanceProfileName: String("president"),
			RoleName:            String("president"),
		}).ExpectInput("DeleteInstanceProfile", &iam.DeleteInstanceProfileInput{
			InstanceProfileName: String("president"),
		}).ExpectInput("DeleteRole", &iam.DeleteRoleInput{
			RoleName: String("president"),
		}).ExpectCommandResult("new-role-arn").ExpectCalls("RemoveRoleFromInstanceProfile", "DeleteInstanceProfile", "DeleteRole").Run(t, logger.DefaultLogger)
	})
}
