package awsspec

import (
	"testing"
)

func TestCreateSubnet(t *testing.T) {
	create := &CreateSubnet{}
	params := map[string]interface{}{
		"cidr": "10.10.10.10/24",
		"vpc":  "vpc-1234",
	}
	t.Run("Validate", func(t *testing.T) {
		params["cidr"] = "10.10.10.10/128"
		errs := create.ValidateCommand(params)
		checkErrs(t, errs, 1, "cidr")

	})
}
