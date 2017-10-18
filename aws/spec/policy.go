package awsspec

import (
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
	"github.com/wallix/awless/aws/driver"
	"github.com/wallix/awless/logger"
)

type AttachPolicy struct {
	_      string `action:"attach" entity:"policy" awsAPI:"iam"`
	logger *logger.Logger
	api    iamiface.IAMAPI
	Arn    *string `awsName:"PolicyArn" awsType:"awsstr" templateName:"arn" required:""`
	User   *string `awsName:"UserName" awsType:"awsstr" templateName:"user"`
	Group  *string `awsName:"GroupName" awsType:"awsstr" templateName:"group"`
	Role   *string `awsName:"RoleName" awsType:"awsstr" templateName:"role"`
}

func (cmd *AttachPolicy) ValidateParams(params []string) ([]string, error) {
	allParams := map[string]bool{
		"arn": false, "user": false, "group": false, "role": false, "service": false, "access": false,
	}

	paramString := " - params: ((service and access) or arn) and (user or group or role)"

	for _, p := range params {
		_, ok := allParams[p]
		if !ok {
			return nil, fmt.Errorf("attach policy: unexpected param key '%s'%s\n", p, paramString)
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

func (cmd *AttachPolicy) ConvertParams() ([]string, func(values map[string]interface{}) (map[string]interface{}, error)) {
	return []string{"access", "service"},
		func(values map[string]interface{}) (map[string]interface{}, error) {
			service, hasService := values["service"].(string)
			access, hasAccess := values["access"].(string)

			if hasService && hasAccess {
				pol, err := awsdriver.LookupAWSPolicy(service, access)
				if err != nil {
					return values, err
				}
				return map[string]interface{}{"arn": pol.Arn}, nil
			} else {
				return nil, nil
			}
		}
}

func (cmd *AttachPolicy) ManualValidateCommand(params map[string]interface{}) (errs []error) {
	if cmd.User == nil && cmd.Group == nil && cmd.Role == nil {
		errs = append(errs, fmt.Errorf("attach policy: missing required field 'user', 'group' or 'role'"))
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
