package awsspec

import (
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
	"github.com/wallix/awless/aws/driver"
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

func (cmd *AttachPolicy) ValidateParams(params []string) ([]string, error) {
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
		return []string{"arn"}, nil
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
		pol, err := awsdriver.LookupAWSPolicy(service, access)
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

func (cmd *AttachPolicy) ValidateCommand(params map[string]interface{}) (errs []error) {
	if err := cmd.Inject(params); err != nil {
		return []error{err}
	}
	if err := validateStruct(cmd); err != nil {
		errs = append(errs, err)
	}
	if cmd.User == nil && cmd.Group == nil && cmd.Role == nil {
		errs = append(errs, fmt.Errorf("AttachPolicy: missing required field 'User', 'Group' or 'Role'"))
	}
	return
}
func (cmd *AttachPolicy) ExtractResultString(i interface{}) string {
	return ""
}

func (cmd *AttachPolicy) ManualRun(ctx, params map[string]interface{}) (interface{}, error) {
	start := time.Now()
	switch {
	case cmd.User != nil:
		input := &iam.AttachUserPolicyInput{}
		input.PolicyArn = cmd.Arn
		input.UserName = cmd.User
		_, err := cmd.api.AttachUserPolicy(input)
		if err != nil {
			return nil, fmt.Errorf("AttachPolicy: %s", err)
		}
		cmd.logger.ExtraVerbosef("ec2.AttachUserPolicy call took %s", time.Since(start))
	case cmd.Group != nil:
		input := &iam.AttachGroupPolicyInput{}
		input.PolicyArn = cmd.Arn
		input.GroupName = cmd.Group
		_, err := cmd.api.AttachGroupPolicy(input)
		if err != nil {
			return nil, fmt.Errorf("AttachPolicy: %s", err)
		}
		cmd.logger.ExtraVerbosef("ec2.AttachGroupPolicy call took %s", time.Since(start))
	case cmd.Role != nil:
		input := &iam.AttachRolePolicyInput{}
		input.PolicyArn = cmd.Arn
		input.RoleName = cmd.Role
		_, err := cmd.api.AttachRolePolicy(input)
		if err != nil {
			return nil, fmt.Errorf("AttachPolicy: %s", err)
		}
		cmd.logger.ExtraVerbosef("ec2.AttachRolePolicy call took %s", time.Since(start))
	default:
		return nil, errors.New("missing one of 'user, group, role' param")
	}
	return "", nil
}

func (cmd *AttachPolicy) DryRun(ctx, params map[string]interface{}) (interface{}, error) {
	return nil, nil
}
