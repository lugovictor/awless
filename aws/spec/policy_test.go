package awsspec

import (
	"testing"
)

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
