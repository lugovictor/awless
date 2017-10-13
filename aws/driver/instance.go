package awsdriver

import (
	"fmt"
	"strings"

	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/wallix/awless/logger"
)

type CreateInstance struct {
	_              string `awsAPI:"ec2" awsCall:"RunInstances" awsInput:"ec2.RunInstancesInput" awsOutput:"ec2.Reservation" awsDryRun:""`
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

func (cmd *CreateInstance) Action() string { return "create" }
func (cmd *CreateInstance) Entity() string { return "instance" }

func (cmd *CreateInstance) CheckParams(params []string) ([]string, error) {
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
			return missing, fmt.Errorf("%s %s: unexpected param key '%s'%s%s\n", cmd.Action(), cmd.Entity(), p, requiredParams, extraParams)
		}
	}

	return missing, nil
}

func (cmd *CreateInstance) Inject(params map[string]interface{}) error {
	return structSetter(cmd, params)
}

func (cmd *CreateInstance) Validate() error {
	return validateStruct(cmd)
}

func (cmd *CreateInstance) ExtractResultString(r *ec2.Reservation) string {
	return awssdk.StringValue(r.Instances[0].InstanceId)
}

func (cmd *CreateInstance) AfterRun(ctx map[string]interface{}, output interface{}) error {
	createTag := NewCommandFuncs["createtag"]().(*CreateTag)
	createTag.Key = awssdk.String("Name")
	createTag.Value = cmd.Name
	createTag.Resource = awssdk.String(cmd.ExtractResultString(output.(*ec2.Reservation)))
	if err := createTag.Validate(); err != nil {
		return err
	}
	if _, err := createTag.Run(); err != nil {
		return err
	}
	return nil
}

type BeforeRunner interface {
	BeforeRun(ctx, params map[string]interface{}) error
}

type AfterRunner interface {
	AfterRun(ctx map[string]interface{}, output interface{}) error
}

func implementsBeforeRun(i interface{}) (BeforeRunner, bool) {
	v, ok := i.(BeforeRunner)
	return v, ok
}

func implementsAfterRun(i interface{}) (AfterRunner, bool) {
	v, ok := i.(AfterRunner)
	return v, ok
}
