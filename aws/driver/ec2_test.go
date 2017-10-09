package awsdriver

import (
	"reflect"
	"testing"

	awssdk "github.com/aws/aws-sdk-go/aws"
)

type TestStruct struct {
	FieldString      *string   `templateName:"fstring"`
	FieldInt64       *int64    `templateName:"fint"`
	FieldBool        *bool     `templateName:"fbool"`
	FieldStringSlice []*string `templateName:"fstrslice"`
}

func TestParamsFillers(t *testing.T) {
	params := map[string]interface{}{
		"fstring":   "jdoe",
		"fint":      "345",
		"fbool":     "true",
		"fstrslice": []interface{}{"one", "two", 3},
	}

	in := &TestStruct{}
	err := structSetter(in, params)
	if err != nil {
		t.Fatal(err)
	}

	exp := &TestStruct{
		FieldString:      awssdk.String("jdoe"),
		FieldInt64:       awssdk.Int64(345),
		FieldBool:        awssdk.Bool(true),
		FieldStringSlice: awssdk.StringSlice([]string{"one", "two", "3"}),
	}

	if got, want := in, exp; !reflect.DeepEqual(got, want) {
		t.Fatalf("\ngot %#v\n\nwant %#v\n", got, want)
	}
}
