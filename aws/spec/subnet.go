package awsspec

import (
	"fmt"
	"net"
	"strings"

	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/wallix/awless/logger"
)

type CreateSubnet struct {
	_                string `action:"create" entity:"subnet" awsAPI:"ec2" awsCall:"CreateSubnet" awsInput:"ec2.CreateSubnetInput" awsOutput:"ec2.CreateSubnetOutput" awsDryRun:""`
	logger           *logger.Logger
	api              ec2iface.EC2API
	CIDR             *string `awsName:"CidrBlock" awsType:"awsstr" templateName:"cidr" required:""`
	VPC              *string `awsName:"VpcId" awsType:"awsstr" templateName:"vpc" required:""`
	AvailabilityZone *string `awsName:"AvailabilityZone" awsType:"awsstr" templateName:"availabilityzone"`
	Name             *string `templateName:"name"`
}

func (cmd *CreateSubnet) ValidateCIDR() error {
	_, _, err := net.ParseCIDR(StringValue(cmd.CIDR))
	return err
}

func (cmd *CreateSubnet) ValidateParams(params []string) ([]string, error) {
	result := structListParamsKeys(cmd)

	var extras, required, missing []string
	for n, isRequired := range result {
		if isRequired {
			required = append(required, n)
			if !contains(params, n) {
				missing = append(missing, n)
			}
		} else {
			extras = append(extras, n)
		}
	}

	var extraParams, requiredParams string
	if len(extras) > 0 {
		extraParams = fmt.Sprintf("\n\t- extra params: %s", strings.Join(extras, ", "))
	}
	if len(required) > 0 {
		requiredParams = fmt.Sprintf("\n\t- required params: %s", strings.Join(required, ", "))
	}

	for _, p := range params {
		_, ok := result[p]
		if !ok {
			return missing, fmt.Errorf("create subnet: unexpected param key '%s'%s%s\n", p, requiredParams, extraParams)
		}
	}

	return missing, nil
}

func (cmd *CreateSubnet) ExtractResultString(r *ec2.CreateSubnetOutput) string {
	return awssdk.StringValue(r.Subnet.SubnetId)
}

func (cmd *CreateSubnet) AfterRun(ctx map[string]interface{}, output interface{}) error {
	createTag := NewCommandFuncs["createtag"]().(*CreateTag)
	createTag.Key = awssdk.String("Name")
	createTag.Value = cmd.Name
	createTag.Resource = awssdk.String(cmd.ExtractResultString(output.(*ec2.CreateSubnetOutput)))
	if errs := createTag.ValidateCommand(nil); len(errs) > 0 {
		return fmt.Errorf("%v", errs)
	}
	if _, err := createTag.Run(nil, nil); err != nil {
		return err
	}
	return nil
}
