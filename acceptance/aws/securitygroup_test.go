package awsat

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func TestSecuritygroup(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		Template(`create securitygroup name=my-sg-name vpc=my-vpc-id description="security group description"`).Mock(&ec2Mock{
			CreateSecurityGroupFunc: func(input *ec2.CreateSecurityGroupInput) (*ec2.CreateSecurityGroupOutput, error) {
				return &ec2.CreateSecurityGroupOutput{GroupId: String("new-secgroup-id")}, nil
			}}).
			ExpectInput("CreateSecurityGroup", &ec2.CreateSecurityGroupInput{
				GroupName:   String("my-sg-name"),
				VpcId:       String("my-vpc-id"),
				Description: String("security group description"),
			}).ExpectCommandResult("new-secgroup-id").ExpectCalls("CreateSecurityGroup").Run(t)
	})

	t.Run("update", func(t *testing.T) {
		Template("update securitygroup id=my-secgroup-id inbound=authorize protocol=tcp cidr=10.10.10.0/24 portrange=10-22").Mock(&ec2Mock{
			AuthorizeSecurityGroupIngressFunc: func(input *ec2.AuthorizeSecurityGroupIngressInput) (*ec2.AuthorizeSecurityGroupIngressOutput, error) {
				return nil, nil
			}}).
			ExpectInput("AuthorizeSecurityGroupIngress", &ec2.AuthorizeSecurityGroupIngressInput{
				GroupId: String("my-secgroup-id"),
				IpPermissions: []*ec2.IpPermission{
					{
						IpProtocol: String("tcp"),
						IpRanges: []*ec2.IpRange{
							{CidrIp: String("10.10.10.0/24")},
						},
						FromPort: Int64(10),
						ToPort:   Int64(22),
					},
				},
			}).ExpectCalls("AuthorizeSecurityGroupIngress").Run(t)
	})

	t.Run("delete", func(t *testing.T) {
		Template("delete securitygroup id=my-secgroup-id").Mock(&ec2Mock{
			DeleteSecurityGroupFunc: func(input *ec2.DeleteSecurityGroupInput) (*ec2.DeleteSecurityGroupOutput, error) {
				return nil, nil
			}}).
			ExpectInput("DeleteSecurityGroup", &ec2.DeleteSecurityGroupInput{
				GroupId: String("my-secgroup-id"),
			}).ExpectCalls("DeleteSecurityGroup").Run(t)
	})

	t.Run("attach", func(t *testing.T) {
		Template("attach securitygroup id=my-secgroup-id instance=secgroup-instance-id").Mock(&ec2Mock{
			DescribeInstanceAttributeFunc: func(input *ec2.DescribeInstanceAttributeInput) (*ec2.DescribeInstanceAttributeOutput, error) {
				return &ec2.DescribeInstanceAttributeOutput{Groups: []*ec2.GroupIdentifier{
					{GroupId: String("secgroup-1")}, {GroupId: String("secgroup-2")},
				}}, nil
			},
			ModifyInstanceAttributeFunc: func(input *ec2.ModifyInstanceAttributeInput) (*ec2.ModifyInstanceAttributeOutput, error) {
				return nil, nil
			}}).
			ExpectInput("DescribeInstanceAttribute", &ec2.DescribeInstanceAttributeInput{
				Attribute:  String("groupSet"),
				InstanceId: String("secgroup-instance-id"),
			}).
			ExpectInput("ModifyInstanceAttribute", &ec2.ModifyInstanceAttributeInput{
				InstanceId: String("secgroup-instance-id"),
				Groups:     []*string{String("secgroup-1"), String("secgroup-2"), String("my-secgroup-id")},
			}).ExpectCalls("DescribeInstanceAttribute", "ModifyInstanceAttribute").Run(t)
	})

	t.Run("detach", func(t *testing.T) {
		Template("detach securitygroup id=my-secgroup-id instance=secgroup-instance-id").Mock(&ec2Mock{
			DescribeInstanceAttributeFunc: func(input *ec2.DescribeInstanceAttributeInput) (*ec2.DescribeInstanceAttributeOutput, error) {
				return &ec2.DescribeInstanceAttributeOutput{Groups: []*ec2.GroupIdentifier{
					{GroupId: String("secgroup-1")}, {GroupId: String("my-secgroup-id")},
				}}, nil
			},
			ModifyInstanceAttributeFunc: func(input *ec2.ModifyInstanceAttributeInput) (*ec2.ModifyInstanceAttributeOutput, error) {
				return nil, nil
			}}).
			ExpectInput("DescribeInstanceAttribute", &ec2.DescribeInstanceAttributeInput{
				Attribute:  String("groupSet"),
				InstanceId: String("secgroup-instance-id"),
			}).
			ExpectInput("ModifyInstanceAttribute", &ec2.ModifyInstanceAttributeInput{
				InstanceId: String("secgroup-instance-id"),
				Groups:     []*string{String("secgroup-1")},
			}).ExpectCalls("DescribeInstanceAttribute", "ModifyInstanceAttribute").Run(t)
	})
}
