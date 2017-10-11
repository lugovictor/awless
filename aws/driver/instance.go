package awsdriver

import (
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/wallix/awless/logger"
)

type CreateInstance struct {
	_              string `awsAPI:"ec2" awsCall:"RunInstances" awsInput:"ec2.RunInstancesInput" awsOutput:"ec2.RunInstancesOutput"`
	result         string
	logger         *logger.Logger
	api            ec2iface.EC2API
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
}

func (cmd *CreateInstance) Inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func (cmd *CreateInstance) Validate() error {
	return validateStruct(cmd)
}

func (cmd *CreateInstance) AfterRun(ctx map[string]interface{}) (interface{}, error) {
	//	_, err = d.Create_Tag(ctx, map[string]interface{}{"key": "Name", "value": params["name"], "resource": id})
	//	if err != nil {
	//		return nil, fmt.Errorf("create instance: adding tags: %s", err)
	//	}
	//	d.logger.Infof("create instance '%s' done", id)
	return nil, nil
}
