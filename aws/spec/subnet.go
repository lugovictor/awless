package awsspec

import (
	"fmt"
	"net"

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

func (cmd *CreateSubnet) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}

func (cmd *CreateSubnet) ValidateCIDR() error {
	_, _, err := net.ParseCIDR(StringValue(cmd.CIDR))
	return err
}

func (cmd *CreateSubnet) ExtractResult(i interface{}) string {
	return awssdk.StringValue(i.(*ec2.CreateSubnetOutput).Subnet.SubnetId)
}

func (cmd *CreateSubnet) AfterRun(ctx map[string]interface{}, output interface{}) error {
	createTag := NewCommandFuncs["createtag"]().(*CreateTag)
	createTag.Key = awssdk.String("Name")
	createTag.Value = cmd.Name
	createTag.Resource = awssdk.String(cmd.ExtractResult(output))
	if errs := createTag.ValidateCommand(nil); len(errs) > 0 {
		return fmt.Errorf("%v", errs)
	}
	if _, err := createTag.Run(nil, nil); err != nil {
		return err
	}
	return nil
}
