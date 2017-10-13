package awsdriver

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/service/iam/iamiface"
	"github.com/wallix/awless/cloud"
	"github.com/wallix/awless/logger"
)

type AttachPolicy struct {
	_      string `awsAPI:"iam"`
	logger *logger.Logger
	api    iamiface.IAMAPI
	Arn    *string `awsName:"PolicyArn" awsType:"awsstr" templateName:"arn" required:""`
	User   *string `awsName:"UserName" awsType:"awsstr" templateName:"user"`
	Group  *string `awsName:"GroupName" awsType:"awsstr" templateName:"group"`
	Role   *string `awsName:"RoleName" awsType:"awsstr" templateName:"role"`
	//	Service *string `templateName:"service"`
	//	Access  *string `templateName:"access"`
}

func (cmd *AttachPolicy) Action() string { return "attach" }
func (cmd *AttachPolicy) Entity() string { return cloud.Policy }

func (cmd *AttachPolicy) CheckParams(params []string) ([]string, error) {
	allParams := map[string]bool{
		"arn": false, "user": false, "group": false, "role": false, "service": false, "access": false,
	}

	paramString := " - params: ((service and access) or arn) and (user or group or role)"

	for _, p := range params {
		_, ok := allParams[p]
		if !ok {
			return nil, fmt.Errorf("%s %s: unexpected param key '%s'%s\n", cmd.Action(), cmd.Entity(), p, paramString)
		}
	}

	hasArn := contains(params, "arn")
	hasService := contains(params, "service")
	hasAccess := contains(params, "access")

	hasUser := contains(params, "user")
	hasGroup := contains(params, "group")
	hasRole := contains(params, "role")

	if !hasUser && !hasGroup && !hasRole {
		return nil, errors.New("missing param 'user', 'group' or 'role'")
	}

	if !hasArn && !hasService && !hasAccess {
		return nil, errors.New("missing param 'arn', 'service' or 'access'")
	} else if hasArn {
		return nil, nil
	} else if hasService && hasAccess {
		return nil, nil
	} else if hasService {
		return []string{"access"}, nil
	} else {
		return []string{"service"}, nil
	}
}

func (cmd *AttachPolicy) Inject(params map[string]interface{}) error {
	if err := structSetter(cmd, params); err != nil {
		return err
	}

	arn, _ := params["arn"]
	service, hasService := params["service"].(string)
	access, hasAccess := params["access"].(string)

	if hasService && hasAccess {
		pol, err := LookupAWSPolicy(service, access)
		if err != nil {
			return err
		}
		arn = pol.Arn
	}
	if err := setFieldWithType(arn, cmd, "Arn", awsstr); err != nil {
		return err
	}
	return nil
}

func (cmd *AttachPolicy) Validate() error {
	if err := validateStruct(cmd); err != nil {
		return err
	}
	if cmd.User == nil && cmd.Group == nil && cmd.Role == nil {
		return errors.New("AttachPolicy: missing required field 'User', 'Group' or 'Role'")
	}
	return nil
}

func (cmd *AttachPolicy) ExtractResultString(i interface{}) string {
	return ""
}
