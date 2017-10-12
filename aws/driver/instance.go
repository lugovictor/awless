package awsdriver

import (
	"fmt"

	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/wallix/awless/logger"
)

type CreateInstance struct {
	_              string `awsAPI:"ec2" awsCall:"RunInstances" awsInput:"ec2.RunInstancesInput" awsOutput:"ec2.Reservation" awsDryRun:""`
	result         string
	logger         *logger.Logger
	api            ec2iface.EC2API
	sess           *session.Session
	Image          *string   `awsName:"ImageId" awsType:"awsstr" templateName:"image" required:""`
	Count          *int64    `awsName:"MaxCount,MinCount" awsType:"awsin64" templateName:"count" required:""`
	Type           *string   `awsName:"InstanceType" awsType:"awsstr" templateName:"type" required:""`
	Subnet         *string   `awsName:"SubnetId" awsType:"awsstr" templateName:"subnet" required:""`
	Keypair        *string   `awsName:"KeyName" awsType:"awsstr" templateName:"keypair"`
	PrivateIP      *string   `awsName:"PrivateIpAddress" awsType:"awsstr" templateName:"ip"`
	UserData       *string   `awsName:"UserData" awsType:"awsfiletobase64" templateName:"userdata"`
	SecurityGroups []*string `awsName:"SecurityGroupIds" awsType:"awsstringslice" templateName:"securitygroup"`
	Lock           *bool     `awsName:"DisableApiTermination" awsType:"awsbool" templateName:"lock"`
	Role           *string   `awsName:"IamInstanceProfile.Name" awsType:"awsstr" templateName:"role"`
	Name           *string   `templateName:"name"`
}

func (cmd *CreateInstance) Inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func (cmd *CreateInstance) Validate() error {
	return validateStruct(cmd)
}

func (cmd *CreateInstance) Action() string { return "create" }
func (cmd *CreateInstance) Entity() string { return "instance" }

func (cmd *CreateInstance) SetResult(r *ec2.Reservation) {
	cmd.result = awssdk.StringValue(r.Instances[0].InstanceId)
}

func (cmd *CreateInstance) AfterRun() error {
	createTag := NewCreateTag(cmd.logger, cmd.sess)
	createTag.Key = awssdk.String("Name")
	createTag.Value = cmd.Name
	createTag.Resource = awssdk.String(cmd.result)
	if err := createTag.Validate(); err != nil {
		return fmt.Errorf("CreateInstance: CreateTag: %s", err)
	}
	if _, err := createTag.Run(); err != nil {
		return fmt.Errorf("CreateInstance: CreateTag: %s", err)
	}
	return nil
}
