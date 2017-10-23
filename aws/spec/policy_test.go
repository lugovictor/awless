package awsspec

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/service/iam"
)

func init() {
	genTestsParams["attachpolicy"] = map[string]interface{}{
		"group":   "administrators",
		"access":  "readonly",
		"service": "ec2",
	}
	genTestsExpected["attachpolicy"] = &iam.AttachGroupPolicyInput{
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
		errs := attach.ValidateCommand(params)
		checkErrs(t, errs, 1, "user", "role", "group")
	})
}

func (m *mockIam) AttachGroupPolicy(input *iam.AttachGroupPolicyInput) (*iam.AttachGroupPolicyOutput, error) {
	if got, want := input, genTestsExpected["attachpolicy"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	return nil, nil
}
