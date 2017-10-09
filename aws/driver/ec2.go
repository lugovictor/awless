package awsdriver

import (
	"fmt"
	"reflect"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/wallix/awless/logger"
)

type CreateInstance struct {
	logger         logger.Logger
	api            ec2iface.EC2API
	Image          *string   `awsName:"ImageId" templateName:"image" required:"true"`
	Count          *int64    `awsName:"MaxCount,MinCount" templateName:"count" required:"true"`
	Type           *string   `awsName:"InstanceType" templateName:"type" required:"true"`
	Subnet         *string   `awsName:"SubnetId" templateName:"subnet" required:"true"`
	Keypair        *string   `awsName:"KeyName" templateName:"keypair"`
	PrivateIP      *string   `awsName:"PrivateIpAddress" templateName:"ip"`
	UserData       *string   `awsName:"UserData" templateName:"userdata" setter:"filetobase64"`
	SecurityGroups []*string `awsName:"SecurityGroupIds" templateName:"securitygroup"`
	Lock           *bool     `awsName:"DisableApiTermination" templateName:"lock"`
	Role           *string   `awsName:"IamInstanceProfile.Name" templateName:"role"`
}

func (cmd *CreateInstance) ProcessParams(ctx map[string]interface{}, params map[string]interface{}) error {
	return nil
}

func (cmd *CreateInstance) Run(ctx map[string]interface{}) (interface{}, error) {
	input := &ec2.RunInstancesInput{}

	start := time.Now()
	output, err := cmd.api.RunInstances(input)
	if err != nil {
		return nil, fmt.Errorf("create instance: %s", err)
	}
	cmd.logger.ExtraVerbosef("ec2.RunInstances call took %s", time.Since(start))
	id := aws.StringValue(output.Instances[0].InstanceId)
	//	_, err = d.Create_Tag(ctx, map[string]interface{}{"key": "Name", "value": params["name"], "resource": id})
	//	if err != nil {
	//		return nil, fmt.Errorf("create instance: adding tags: %s", err)
	//	}

	//	d.logger.Infof("create instance '%s' done", id)
	return id, nil
}

func structSetter(s interface{}, params map[string]interface{}) error {
	val := reflect.ValueOf(s).Elem()
	stru := val.Type()

	for i := 0; i < stru.NumField(); i++ {
		field := stru.Field(i)
		tplName := field.Tag.Get("templateName")
		fieldType := -1
		if v, ok := params[tplName]; ok {
			switch field.Type.String() {
			case "*string":
				fieldType = awsstr
			case "*int64":
				fieldType = awsint64
			case "[]*string":
				fieldType = awsstringslice
			case "*bool":
				fieldType = awsbool
			default:
				return fmt.Errorf("unknown type %s", field.Type.String())
			}
			if err := setFieldWithType(v, s, field.Name, fieldType); err != nil {
				return err
			}
		}
	}

	return nil
}
