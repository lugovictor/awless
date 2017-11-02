package awsspec

import (
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func (m *mockEc2) CreateTagsRequest(input *ec2.CreateTagsInput) (req *request.Request, output *ec2.CreateTagsOutput) {
	output = &ec2.CreateTagsOutput{}
	req = request.New(aws.Config{}, metadata.ClientInfo{}, request.Handlers{}, nil, &request.Operation{}, input, output)
	if len(input.Resources) >= 1 {
		switch StringValue(input.Resources[0]) {
		case "res.createtag":
			if got, want := input, genTestsExpected["createtag"]; !reflect.DeepEqual(got, want) {
				req.Error = fmt.Errorf("got %#v, want %#v", got, want)
			}
		case "res.createinstance":
			if got, want := input, genTestsExpected["createinstance.createtag"]; !reflect.DeepEqual(got, want) {
				req.Error = fmt.Errorf("got %#v, want %#v", got, want)
			}
		case "res.createsubnet":
			if got, want := input, genTestsExpected["createsubnet.createtag"]; !reflect.DeepEqual(got, want) {
				req.Error = fmt.Errorf("got %#v, want %#v", got, want)
			}
		case "res.createvpc":
			if got, want := input, genTestsExpected["createvpc.createtag"]; !reflect.DeepEqual(got, want) {
				req.Error = fmt.Errorf("got %#v, want %#v", got, want)
			}
		}
	}
	return
}

func (m *mockEc2) DeleteTagsRequest(input *ec2.DeleteTagsInput) (req *request.Request, output *ec2.DeleteTagsOutput) {
	output = &ec2.DeleteTagsOutput{}
	req = request.New(aws.Config{}, metadata.ClientInfo{}, request.Handlers{}, nil, &request.Operation{}, input, output)
	if got, want := input, genTestsExpected["deletetag"]; !reflect.DeepEqual(got, want) {
		req.Error = fmt.Errorf("got %#v, want %#v", got, want)
	}
	return
}
