package awsspec

import (
	"fmt"
	"strings"
	"time"

	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
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

const (
	notFoundState = "not-found"
)

type CheckInstance struct {
	_       string `action:"check" entity:"instance" awsAPI:"ec2"`
	logger  *logger.Logger
	api     ec2iface.EC2API
	Id      *string `templateName:"id" required:""`
	State   *string `templateName:"state" required:""`
	Timeout *int64  `templateName:"timeout" required:""`
}

func (cmd *CheckInstance) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}

func (cmd *CheckInstance) ValidateState() error {
	switch strings.ToLower(StringValue(cmd.State)) {
	case "pending", "running", "shutting-down", "terminated", "stopping", "stopped", "not-found":
		return nil
	default:
		return fmt.Errorf("invalid value '%s' for 'state' parameter: expect 'pending', 'running', 'shutting-down', 'terminated', 'stopping', 'stopped' or 'not-found'", StringValue(cmd.State))
	}
}

func (cmd *CheckInstance) ManualRun(ctx, params map[string]interface{}) (interface{}, error) {
	input := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{cmd.Id},
	}

	c := &checker{
		description: fmt.Sprintf("instance %s", StringValue(cmd.Id)),
		timeout:     time.Duration(Int64AsIntValue(cmd.Timeout)) * time.Second,
		frequency:   5 * time.Second,
		fetchFunc: func() (string, error) {
			output, err := cmd.api.DescribeInstances(input)
			if err != nil {
				if awserr, ok := err.(awserr.Error); ok {
					if awserr.Code() == "InstanceNotFound" {
						return notFoundState, nil
					}
				} else {
					return "", err
				}
			} else {
				if res := output.Reservations; len(res) > 0 {
					if instances := output.Reservations[0].Instances; len(instances) > 0 {
						for _, inst := range instances {
							if StringValue(inst.InstanceId) == params["id"] {
								return StringValue(inst.State.Name), nil
							}
						}
					}
				}
			}
			return notFoundState, nil
		},
		expect: fmt.Sprint(params["state"]),
		logger: cmd.logger,
	}
	return nil, c.check()
}
