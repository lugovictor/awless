package awsspec

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/wallix/awless/logger"
)

func TestCreateInstance(t *testing.T) {
	newCmd := func() *CreateInstance {
		return &CreateInstance{api: &mockInstance{t: t}, logger: logger.DiscardLogger}
	}
	NewCommandFuncs["createinstance"] = func() interface{} { return newCmd() }
	NewCommandFuncs["createtag"] = func() interface{} { return &CreateTag{api: &mockInstance{t: t}, logger: logger.DiscardLogger} }

	params := map[string]interface{}{
		"name":   "myinstance",
		"subnet": "sub_1",
		"image":  "ami-1234",
		"type":   "t2.nano",
		"count":  3,
	}
	missings, err := newCmd().ValidateParams(keys(params))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(missings), 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if errs := newCmd().ValidateCommand(params); len(errs) > 0 {
		t.Fatalf("%v", errs)
	}
	res, err := newCmd().Run(nil, params)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res, "id-my-instance"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

type mockInstance struct {
	ec2iface.EC2API
	t *testing.T
}

func (m *mockInstance) RunInstances(got *ec2.RunInstancesInput) (*ec2.Reservation, error) {
	want := &ec2.RunInstancesInput{
		SubnetId:     String("sub_1"),
		ImageId:      String("ami-1234"),
		InstanceType: String("t2.nano"),
		MinCount:     Int64(3),
		MaxCount:     Int64(3),
	}
	if !reflect.DeepEqual(got, want) {
		m.t.Fatalf("got %#v, want %#v", got, want)
	}
	return &ec2.Reservation{Instances: []*ec2.Instance{{InstanceId: String("id-my-instance")}}}, nil
}

func (m *mockInstance) CreateTagsRequest(got *ec2.CreateTagsInput) (req *request.Request, output *ec2.CreateTagsOutput) {
	want := &ec2.CreateTagsInput{
		Resources: []*string{String("id-my-instance")},
		Tags:      []*ec2.Tag{{Key: String("Name"), Value: String("myinstance")}},
	}
	if !reflect.DeepEqual(got, want) {
		m.t.Fatalf("got %#v, want %#v", got, want)
	}

	output = &ec2.CreateTagsOutput{}
	req = request.New(aws.Config{}, metadata.ClientInfo{}, request.Handlers{}, nil, &request.Operation{}, got, output)
	return
}
