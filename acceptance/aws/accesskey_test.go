package awsat

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/iam"
)

func TestAccesskey(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		Template("create accesskey user=jdoe no-prompt=true").
			Mock(&iamMock{
				CreateAccessKeyFunc: func(*iam.CreateAccessKeyInput) (*iam.CreateAccessKeyOutput, error) {
					return &iam.CreateAccessKeyOutput{AccessKey: &iam.AccessKey{AccessKeyId: String("new-keypair-id")}}, nil
				},
			}).ExpectInput("CreateAccessKey", &iam.CreateAccessKeyInput{UserName: String("jdoe")}).
			ExpectCommandResult("new-keypair-id").ExpectCalls("CreateAccessKey").Run(t)
	})

	t.Run("delete", func(t *testing.T) {
		Template("delete accesskey id=ACCESSKEYID user=jdoe").
			Mock(&iamMock{
				DeleteAccessKeyFunc: func(param0 *iam.DeleteAccessKeyInput) (*iam.DeleteAccessKeyOutput, error) {
					return nil, nil
				},
			}).ExpectInput("DeleteAccessKey", &iam.DeleteAccessKeyInput{
			UserName:    String("jdoe"),
			AccessKeyId: String("ACCESSKEYID"),
		}).ExpectCalls("DeleteAccessKey").Run(t)
	})
}
