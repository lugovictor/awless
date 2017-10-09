package awsat

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/iam"
)

func TestPolicy(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		Template(
			"create policy name=AwlessInfraReadonlyPolicy effect=Allow action=ec2:Describe*,autoscaling:Describe*,elasticloadbalancing:Describe* "+
				"resource=\"arn:aws:iam::0123456789:mfa/${aws:username}\",\"arn:aws:iam::0123456789:user/${aws:username}\" "+
				"conditions=\"aws:MultiFactorAuthPresent==true\", \"aws:TokenIssueTime!=Null\" description=\"Readonly access to infra resources\"").
			Mock(&iamMock{
				CreatePolicyFunc: func(input *iam.CreatePolicyInput) (*iam.CreatePolicyOutput, error) {
					return &iam.CreatePolicyOutput{Policy: &iam.Policy{Arn: String("new-policy-arn")}}, nil
				},
			}).ExpectInput("CreatePolicy", &iam.CreatePolicyInput{
			PolicyName:  String("AwlessInfraReadonlyPolicy"),
			Description: String("Readonly access to infra resources"),
			PolicyDocument: String(`{
 "Version": "2012-10-17",
 "Statement": [
  {
   "Effect": "Allow",
   "Action": [
    "ec2:Describe*",
    "autoscaling:Describe*",
    "elasticloadbalancing:Describe*"
   ],
   "Resource": [
    "arn:aws:iam::0123456789:mfa/${aws:username}",
    "arn:aws:iam::0123456789:user/${aws:username}"
   ],
   "Condition": {
    "Bool": {
     "aws:MultiFactorAuthPresent": "true"
    },
    "Null": {
     "aws:TokenIssueTime": "false"
    }
   }
  }
 ]
}`)}).ExpectCommandResult("new-policy-arn").ExpectCalls("CreatePolicy").Run(t)
	})
}
