package awsspec

import (
	"fmt"

	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/wallix/awless/logger"
)

type CreateInstance struct {
	_              string `action:"create" entity:"instance" awsAPI:"ec2" awsCall:"RunInstances" awsInput:"ec2.RunInstancesInput" awsOutput:"ec2.Reservation" awsDryRun:""`
	logger         *logger.Logger
	api            ec2iface.EC2API
	Image          *string   `awsName:"ImageId" awsType:"awsstr" templateName:"image" required:""`
	Count          *int64    `awsName:"MaxCount,MinCount" awsType:"awsin64" templateName:"count" required:""`
	Type           *string   `awsName:"InstanceType" awsType:"awsstr" templateName:"type" required:""`
	Name           *string   `templateName:"name" required:""`
	Subnet         *string   `awsName:"SubnetId" awsType:"awsstr" templateName:"subnet" required:""`
	Keypair        *string   `awsName:"KeyName" awsType:"awsstr" templateName:"keypair"`
	PrivateIP      *string   `awsName:"PrivateIpAddress" awsType:"awsstr" templateName:"ip"`
	UserData       *string   `awsName:"UserData" awsType:"awsfiletobase64" templateName:"userdata"`
	SecurityGroups []*string `awsName:"SecurityGroupIds" awsType:"awsstringslice" templateName:"securitygroup"`
	Lock           *bool     `awsName:"DisableApiTermination" awsType:"awsbool" templateName:"lock"`
	Role           *string   `awsName:"IamInstanceProfile.Name" awsType:"awsstr" templateName:"role"`
}

func (cmd *CreateInstance) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}

func (cmd *CreateInstance) ExtractResult(i interface{}) string {
	return awssdk.StringValue(i.(*ec2.Reservation).Instances[0].InstanceId)
}

func (cmd *CreateInstance) AfterRun(ctx map[string]interface{}, output interface{}) error {
	createTag := NewCommandFuncs["createtag"]().(*CreateTag)
	createTag.Key = awssdk.String("Name")
	createTag.Value = cmd.Name
	createTag.Resource = awssdk.String(cmd.ExtractResult(output))
	if errs := createTag.ValidateCommand(nil, nil); len(errs) > 0 {
		return fmt.Errorf("%v", errs)
	}
	if _, err := createTag.Run(ctx, nil); err != nil {
		return err
	}
	return nil
}

type DeleteInstance struct {
	_      string `action:"delete" entity:"instance" awsAPI:"ec2" awsCall:"TerminateInstances" awsInput:"ec2.TerminateInstancesInput" awsOutput:"ec2.TerminateInstancesOutput" awsDryRun:""`
	logger *logger.Logger
	api    ec2iface.EC2API
	IDs    []*string `awsName:"InstanceIds" awsType:"awsstringslice" templateName:"ids"`
}

func (cmd *DeleteInstance) ConvertParams() ([]string, func(values map[string]interface{}) (map[string]interface{}, error)) {
	return []string{"id"},
		func(values map[string]interface{}) (map[string]interface{}, error) {
			if id, hasID := values["id"]; hasID {
				return map[string]interface{}{"ids": id}, nil
			} else {
				return nil, nil
			}
		}
}

func (cmd *DeleteInstance) ValidateParams(params []string) ([]string, error) {
	return paramRule{tree: oneOf(node("ids"), node("id"))}.verify(params)
}
