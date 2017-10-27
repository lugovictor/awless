package awsspec

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

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
		return &ec2.Reservation{Instances: []*ec2.Instance{{InstanceId: String("res.createinstance")}}}
	}
	genTestsOutput["createinstance"] = "res.createinstance"
	genTestsExpected["createinstance.createtag"] = &ec2.CreateTagsInput{
		Resources: []*string{String("res.createinstance")},
		Tags: []*ec2.Tag{
			{Key: String("Name"), Value: String("myinstance")},
		},
	}

	genTestsParams["deleteinstance"] = map[string]interface{}{
		"id": []interface{}{"my-instance-id1", "my-instance-id2"},
	}
	genTestsExpected["deleteinstance"] = &ec2.TerminateInstancesInput{
		InstanceIds: []*string{String("my-instance-id1"), String("my-instance-id2")},
	}
	genTestsParams["checkinstance"] = map[string]interface{}{
		"id":      "my-check-instance-id",
		"state":   "running",
		"timeout": 0,
	}
	genTestsExpected["checkinstance"] = &ec2.DescribeInstancesInput{
		InstanceIds: []*string{String("my-check-instance-id")},
	}
}

func (m *mockEc2) DescribeInstances(input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	if len(input.InstanceIds) > 0 && StringValue(input.InstanceIds[0]) == "my-check-instance-id" {
		if got, want := input, genTestsExpected["checkinstance"]; !reflect.DeepEqual(got, want) {
			return nil, fmt.Errorf("got %#v, want %#v", got, want)
		}
		return &ec2.DescribeInstancesOutput{Reservations: []*ec2.Reservation{
			{Instances: []*ec2.Instance{{InstanceId: input.InstanceIds[0], State: &ec2.InstanceState{Name: String("running")}}}},
		}}, nil
	}
	return nil, fmt.Errorf("DescribeInstances mock: invalid value for 'InstanceIds': %#v", input)
}

func generateTmpFile(content string) (path string) {
	file, err := ioutil.TempFile("", "tmpfile")
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(file.Name(), []byte(content), 0644)
	return file.Name()
}
