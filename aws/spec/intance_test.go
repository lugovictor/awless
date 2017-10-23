package awsspec

import (
	"encoding/base64"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func init() {
	userdataFile := generateTmpFile("this is my content with {{ .GoTemplateVar }} content")
	genTestsContext["createinstance"] = map[string]interface{}{"GoTemplateVar": "awesome"}
	genTestsParams["createinstance"] = map[string]interface{}{
		"count":         3,
		"image":         "ami-1234",
		"name":          "myinstance",
		"subnet":        "sub_1",
		"type":          "t2.nano",
		"keypair":       "mykp",
		"ip":            "10.2.3.4",
		"userdata":      userdataFile,
		"securitygroup": "sg-1234",
		"lock":          "true",
		"role":          "myrole",
	}
	genTestsCleanupFunc["createinstance"] = func() { os.Remove(userdataFile) }
	genTestsExpected["createinstance"] = &ec2.RunInstancesInput{
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
	}
	genTestsOutputExtractFunc["createinstance"] = func() interface{} {
		return &ec2.Reservation{Instances: []*ec2.Instance{{InstanceId: String("id-my-instance")}}}
	}
	genTestsOutput["createinstance"] = "id-my-instance"

	genTestsParams["deleteinstance"] = map[string]interface{}{
		"id": []interface{}{"my-instance-id1", "my-instance-id2"},
	}
	genTestsExpected["deleteinstance"] = &ec2.TerminateInstancesInput{
		InstanceIds: []*string{String("my-instance-id1"), String("my-instance-id2")},
	}
}

func generateTmpFile(content string) (path string) {
	file, err := ioutil.TempFile(".", "tmpfile")
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(file.Name(), []byte(content), 0644)
	return file.Name()
}

//Not tested
func (m *mockEc2) CreateTagsRequest(got *ec2.CreateTagsInput) (req *request.Request, output *ec2.CreateTagsOutput) {
	output = &ec2.CreateTagsOutput{}
	req = request.New(aws.Config{}, metadata.ClientInfo{}, request.Handlers{}, nil, &request.Operation{}, got, output)
	return
}
