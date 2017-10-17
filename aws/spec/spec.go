package awsspec

import (
	"fmt"
	"math/rand"

	"github.com/wallix/awless/cloud"
)

const (
	dryRunOperation = "DryRunOperation"
	notFound        = "NotFound"
)

type BeforeRunner interface {
	BeforeRun(ctx, params map[string]interface{}) error
}

type AfterRunner interface {
	AfterRun(ctx map[string]interface{}, output interface{}) error
}

type command interface {
	Run(ctx map[string]interface{}, params map[string]interface{}) (interface{}, error)
	DryRun(ctx map[string]interface{}, params map[string]interface{}) (interface{}, error)
	Action() string
	Entity() string
}

func implementsBeforeRun(i interface{}) (BeforeRunner, bool) {
	v, ok := i.(BeforeRunner)
	return v, ok
}

func implementsAfterRun(i interface{}) (AfterRunner, bool) {
	v, ok := i.(AfterRunner)
	return v, ok
}

func fakeDryRunId(entity string) string {
	suffix := rand.Intn(1e6)
	switch entity {
	case cloud.Instance:
		return fmt.Sprintf("i-%d", suffix)
	case cloud.Subnet:
		return fmt.Sprintf("subnet-%d", suffix)
	case cloud.Vpc:
		return fmt.Sprintf("vpc-%d", suffix)
	case cloud.Volume:
		return fmt.Sprintf("vol-%d", suffix)
	case cloud.SecurityGroup:
		return fmt.Sprintf("sg-%d", suffix)
	case cloud.InternetGateway:
		return fmt.Sprintf("igw-%d", suffix)
	case cloud.NatGateway:
		return fmt.Sprintf("nat-%d", suffix)
	case cloud.RouteTable:
		return fmt.Sprintf("rtb-%d", suffix)
	default:
		return fmt.Sprintf("dryrunid-%d", suffix)
	}
}
