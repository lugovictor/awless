package awsat

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/wallix/awless/aws/spec"
	"github.com/wallix/awless/logger"
)

var MockCommands map[string]func(interface{}) interface{}

func init() {
	MockCommands = make(map[string]func(interface{}) interface{})
	MockCommands["createinstance"] = func(api interface{}) interface{} {
		cmd := new(awsspec.CreateInstance)
		cmd.SetApi(api.(ec2iface.EC2API))
		cmd.SetLogger(logger.DiscardLogger)
		return cmd
	}
}

type ec2Mock struct {
	ec2iface.EC2API
	RunInstancesFunc      func(input *ec2.RunInstancesInput) (*ec2.Reservation, error)
	CreateTagsRequestFunc func(input *ec2.CreateTagsInput) (req *request.Request, output *ec2.CreateTagsOutput)
}

func (m *ec2Mock) RunInstances(input *ec2.RunInstancesInput) (*ec2.Reservation, error) {
	return m.RunInstancesFunc(input)
}

func (m *ec2Mock) CreateTagsRequest(input *ec2.CreateTagsInput) (*request.Request, *ec2.CreateTagsOutput) {
	return m.CreateTagsRequestFunc(input)
}
