package awsspec

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/service/iam"
)

func init() {
	genTestsParams["createpolicy"] = map[string]interface{}{
		"name":        "AwlessInfraReadonlyPolicy",
		"effect":      "Allow",
		"action":      []interface{}{"ec2:Describe*", "autoscaling:Describe*", "elasticloadbalancing:Describe*"},
		"resource":    []interface{}{"arn:aws:iam::0123456789:mfa/${aws:username}", "arn:aws:iam::0123456789:user/${aws:username}"},
		"conditions":  []interface{}{"aws:MultiFactorAuthPresent==true", "aws:TokenIssueTime!=Null"},
		"description": "Readonly access to infra resources",
	}
	genTestsExpected["createpolicy"] = &iam.CreatePolicyInput{
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
}`),
	}
	genTestsOutputExtractFunc["createpolicy"] = func() interface{} {
		return &iam.CreatePolicyOutput{Policy: &iam.Policy{Arn: String("mynewpolicy")}}
	}
	genTestsOutput["createpolicy"] = "mynewpolicy"

	genTestsParams["updatepolicy"] = map[string]interface{}{
		"arn":        "updatepolicy-arn-policy",
		"effect":     "Deny",
		"action":     []interface{}{"ec2:AttachVolume", "DescribeVolumeAttribute"},
		"resource":   "arn:aws:ec2:eu-west-1:0123456789:volume/*",
		"conditions": "aws:MultiFactorAuthPresent==true",
	}
	genTestsExpected["updatepolicy"] = &iam.CreatePolicyVersionInput{
		PolicyArn:    String("updatepolicy-arn-policy"),
		SetAsDefault: Bool(true),
		PolicyDocument: String(`{
 "Version": "2012-10-17",
 "Statement": [
  {
   "Effect": "Allow",
   "Action": [
    "ec2:AttachVolume",
    "ec2:DetachVolume"
   ],
   "Resource": "arn:aws:ec2:eu-west-1:0123456789:instance/*",
   "Condition": {
    "StringEquals": {
     "ec2:ResourceTag/department": "dev"
    },
    "Null": {
     "aws:TokenIssueTime": "false"
    }
   }
  },
  {
   "Effect": "Allow",
   "Action": [
    "ec2:AttachVolume",
    "ec2:DetachVolume"
   ],
   "Resource": "arn:aws:ec2:eu-west-1:0123456789:volume/*",
   "Condition": {
    "StringEquals": {
     "ec2:ResourceTag/volume_user": "${aws:username}"
    }
   }
  },
  {
   "Effect": "Deny",
   "Action": [
    "ec2:AttachVolume",
    "DescribeVolumeAttribute"
   ],
   "Resource": [
    "arn:aws:ec2:eu-west-1:0123456789:volume/*"
   ],
   "Condition": {
    "Bool": {
     "aws:MultiFactorAuthPresent": "true"
    }
   }
  }
 ]
}`),
	}

	genTestsParams["deletepolicy"] = map[string]interface{}{
		"arn":          "arn-policy-to-delete",
		"all-versions": "true",
	}
	genTestsExpected["deletepolicy"] = &iam.DeletePolicyInput{
		PolicyArn: String("arn-policy-to-delete"),
	}
	genTestsExpected["deletepolicy.v1"] = &iam.DeletePolicyVersionInput{
		VersionId: String("v1"),
		PolicyArn: String("arn-policy-to-delete"),
	}

	genTestsParams["attachpolicy"] = map[string]interface{}{
		"group":   "administrators",
		"access":  "readonly",
		"service": "ec2",
	}
	genTestsExpected["attachpolicy"] = &iam.AttachGroupPolicyInput{
		GroupName: String("administrators"),
		PolicyArn: String("arn:aws:iam::aws:policy/AmazonEC2ReadOnlyAccess"),
	}

	genTestsParams["detachpolicy"] = map[string]interface{}{
		"group":   "administrators",
		"access":  "readonly",
		"service": "ec2",
	}
	genTestsExpected["detachpolicy"] = &iam.DetachGroupPolicyInput{
		GroupName: String("administrators"),
		PolicyArn: String("arn:aws:iam::aws:policy/AmazonEC2ReadOnlyAccess"),
	}
}

func TestAttachPolicy(t *testing.T) {
	attach := &AttachPolicy{}
	params := map[string]interface{}{
		"arn": "arn:1234:2345:3456",
	}
	t.Run("Validate", func(t *testing.T) {
		errs := attach.ValidateCommand(params, nil)
		checkErrs(t, errs, 1, "user", "role", "group")
	})
}

func (m *mockIam) CreatePolicy(input *iam.CreatePolicyInput) (*iam.CreatePolicyOutput, error) {
	if got, want := input, genTestsExpected["createpolicy"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["createpolicy"]; ok {
		return outFunc().(*iam.CreatePolicyOutput), nil
	}
	return nil, nil
}

func (m *mockIam) ListPolicyVersions(input *iam.ListPolicyVersionsInput) (*iam.ListPolicyVersionsOutput, error) {
	if StringValue(input.PolicyArn) == "updatepolicy-arn-policy" {
		return &iam.ListPolicyVersionsOutput{Versions: []*iam.PolicyVersion{
			{VersionId: String("v2"), IsDefaultVersion: Bool(true)},
		}}, nil
	}
	if StringValue(input.PolicyArn) == "arn-policy-to-delete" {
		return &iam.ListPolicyVersionsOutput{Versions: []*iam.PolicyVersion{
			{VersionId: String("v1"), IsDefaultVersion: Bool(false)},
			{VersionId: String("v2"), IsDefaultVersion: Bool(true)},
		}}, nil
	}
	return nil, fmt.Errorf("ListPolicyVersions mock: can not find policy with arn '%s'", StringValue(input.PolicyArn))
}

func (m *mockIam) GetPolicyVersion(input *iam.GetPolicyVersionInput) (*iam.GetPolicyVersionOutput, error) {
	if StringValue(input.PolicyArn) == "updatepolicy-arn-policy" {
		return &iam.GetPolicyVersionOutput{PolicyVersion: &iam.PolicyVersion{
			Document: String(url.QueryEscape(`{
		 "Version": "2012-10-17",
		 "Statement": [
		  {
		   "Effect": "Allow",
		   "Action": [
		    "ec2:AttachVolume",
		    "ec2:DetachVolume"
		   ],
		   "Resource": "arn:aws:ec2:eu-west-1:0123456789:instance/*",
		   "Condition": {
		    "StringEquals": {
		     "ec2:ResourceTag/department": "dev"
		    },
		    "Null": {
		     "aws:TokenIssueTime": "false"
		    }
		   }
		  },
		  {
		   "Effect": "Allow",
		   "Action": [
		    "ec2:AttachVolume",
		    "ec2:DetachVolume"
		   ],
		   "Resource": "arn:aws:ec2:eu-west-1:0123456789:volume/*",
		   "Condition": {
		    "StringEquals": {
		     "ec2:ResourceTag/volume_user": "${aws:username}"}
		    }
		  }
		 ]
		}`)),
			IsDefaultVersion: Bool(true),
			VersionId:        String("v2"),
		}}, nil
	}
	return nil, fmt.Errorf("GetPolicyVersion mock: can not find policy with arn '%s'", StringValue(input.PolicyArn))
}

func (m *mockIam) CreatePolicyVersion(input *iam.CreatePolicyVersionInput) (*iam.CreatePolicyVersionOutput, error) {
	if got, want := input, genTestsExpected["updatepolicy"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	return nil, nil
}

func (m *mockIam) DeletePolicyVersion(input *iam.DeletePolicyVersionInput) (*iam.DeletePolicyVersionOutput, error) {
	if got, want := input, genTestsExpected["deletepolicy.v1"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	return nil, nil
}

func (m *mockIam) DeletePolicy(input *iam.DeletePolicyInput) (*iam.DeletePolicyOutput, error) {
	if got, want := input, genTestsExpected["deletepolicy"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	return nil, nil
}

func (m *mockIam) AttachGroupPolicy(input *iam.AttachGroupPolicyInput) (*iam.AttachGroupPolicyOutput, error) {
	if got, want := input, genTestsExpected["attachpolicy"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	return nil, nil
}

func (m *mockIam) DetachGroupPolicy(input *iam.DetachGroupPolicyInput) (*iam.DetachGroupPolicyOutput, error) {
	if got, want := input, genTestsExpected["detachpolicy"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	return nil, nil
}
