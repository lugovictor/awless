package awsdriver

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/wallix/awless/logger"
)

type CreateInstance struct {
	logger         logger.Logger
	api            ec2iface.EC2API
	Image          *string   `awsName:"ImageId" templateName:"image" required:""`
	Count          *int64    `awsName:"MaxCount,MinCount" templateName:"count" required:""`
	Type           *string   `awsName:"InstanceType" templateName:"type" required:""`
	Subnet         *string   `awsName:"SubnetId" templateName:"subnet" required:""`
	Keypair        *string   `awsName:"KeyName" templateName:"keypair"`
	PrivateIP      *string   `awsName:"PrivateIpAddress" templateName:"ip"`
	UserData       *string   `awsName:"UserData" templateName:"userdata" setter:"filetobase64"`
	SecurityGroups []*string `awsName:"SecurityGroupIds" templateName:"securitygroup"`
	Lock           *bool     `awsName:"DisableApiTermination" templateName:"lock"`
	Role           *string   `awsName:"IamInstanceProfile.Name" templateName:"role"`
}

func (cmd *CreateInstance) Inject(params map[string]interface{}) error {
	return nil
}

func (cmd *CreateInstance) Validate() error {
	return validateStruct(cmd)
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
			kind := field.Type.Kind()
			if kind == reflect.Ptr {
				switch field.Type.Elem().Kind() {
				case reflect.String:
					fieldType = awsstr
				case reflect.Int64:
					fieldType = awsint64
				case reflect.Bool:
					fieldType = awsbool
				default:
					return fmt.Errorf("unknown type in pointer %s", field.Type.String())
				}
			} else if kind == reflect.Slice && field.Type.Elem().Kind() == reflect.Ptr {
				switch field.Type.Elem().Elem().Kind() {
				case reflect.String:
					fieldType = awsstringslice
				case reflect.Int64:
					fieldType = awsint64slice
				default:
					return fmt.Errorf("unknown type in slice %s", field.Type.String())
				}
			}
			if err := setFieldWithType(v, s, field.Name, fieldType); err != nil {
				return err
			}
		}
	}
	return nil
}

func validateStruct(s interface{}) error {
	val := reflect.ValueOf(s)
	stru := val.Elem().Type()

	var messages []string
	for i := 0; i < stru.NumField(); i++ {
		field := stru.Field(i)
		if _, ok := field.Tag.Lookup("required"); ok {
			if val.Elem().Field(i).IsNil() {
				messages = append(messages, fmt.Sprintf("missing required field %s", field.Name))
			}
			continue
		}

		if _, ok := field.Tag.Lookup("templateName"); ok {
			methName := fmt.Sprintf("Validate%s", field.Name)
			meth := val.MethodByName(methName)
			if meth != (reflect.Value{}) {
				results := meth.Call(nil)
				if len(results) == 1 {
					if iface := results[0].Interface(); iface != nil {
						messages = append(messages, iface.(error).Error())
					}
				}
			}
		}
	}
	if len(messages) > 0 {
		return errors.New(strings.Join(messages, "; "))
	}
	return nil
}
