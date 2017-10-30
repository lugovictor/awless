package awsat

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func TestInstance(t *testing.T) {
	userdataFile := generateTmpFile("this is my content with {{ .Variables.GoTemplateVar }} content")
	defer os.Remove(userdataFile)

	builder := Command("GoTemplateVar=awesome\ncreate instance count=3 image=ami-1234 "+
		"name=myinstance subnet=sub_1 type=t2.nano keypair=mykp ip=10.2.3.4 "+
		"userdata="+userdataFile+" securitygroup=sg-1234 lock=true role=myrole").
		Input("RunInstances", &ec2.RunInstancesInput{
			SubnetId:              String("sub_1"),
			ImageId:               String("ami-1234"),
			InstanceType:          String("t2.nano"),
			MinCount:              Int64(3),
			MaxCount:              Int64(3),
			KeyName:               String("mykp"),
			PrivateIpAddress:      String("10.2.3.4"),
			SecurityGroupIds:      []*string{String("sg-1234")},
			DisableApiTermination: Bool(true),
			IamInstanceProfile:    &ec2.IamInstanceProfileSpecification{Name: String("myrole")},
			UserData:              String(base64.StdEncoding.EncodeToString([]byte("this is my content with awesome content"))),
		})
	builder.Input("CreateTagsRequest", &ec2.CreateTagsInput{
		Resources: []*string{String("res.createinstance")},
		Tags: []*ec2.Tag{
			{Key: String("Name"), Value: String("myinstance")},
		},
	}).Mock(
		&ec2Mock{
			RunInstancesFunc: func(input *ec2.RunInstancesInput) (*ec2.Reservation, error) {
				builder.VerifyInput("RunInstances", input)
				return &ec2.Reservation{Instances: []*ec2.Instance{{InstanceId: String("res.createinstance")}}}, nil
			},
			CreateTagsRequestFunc: func(input *ec2.CreateTagsInput) (req *request.Request, output *ec2.CreateTagsOutput) {
				output = &ec2.CreateTagsOutput{}
				req = request.New(aws.Config{}, metadata.ClientInfo{}, request.Handlers{}, nil, &request.Operation{}, input, output)
				builder.VerifyInput("CreateTagsRequest", input)
				return
			}}).VerifyTemplateResult("res.createinstance").Run(t)
}

func generateTmpFile(content string) (path string) {
	file, err := ioutil.TempFile("", "tmpfile")
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(file.Name(), []byte(content), 0644)
	return file.Name()
}
