package awsdriver

import (
	"errors"
	"reflect"
	"strings"
	"testing"

	awssdk "github.com/aws/aws-sdk-go/aws"
)

type TestStruct struct {
	FieldString      *string   `templateName:"fstring"`
	FieldInt64       *int64    `templateName:"fint"`
	FieldBool        *bool     `templateName:"fbool" required:""`
	FieldStringSlice []*string `templateName:"fstrslice"`
}

func (ts *TestStruct) ValidateFieldString() (err error) {
	if len(*ts.FieldString) == 0 {
		err = errors.New("fstring should not be empty")
	}
	return
}

func (ts *TestStruct) ValidateFieldInt64() (err error) {
	if *ts.FieldInt64 > 10 {
		err = errors.New("fint should not exceed 10")
	}
	return
}

func TestValidateStruct(t *testing.T) {
	tcases := []struct {
		stru        *TestStruct
		contains    []string
		notContains []string
		err         bool
	}{
		{
			stru:     &TestStruct{FieldString: awssdk.String(""), FieldInt64: awssdk.Int64(12)},
			contains: []string{"fstring should not be empty", "fint should not exceed 10", "missing required field FieldBool"},
			err:      true,
		},
		{
			stru:        &TestStruct{FieldString: awssdk.String("not empty"), FieldInt64: awssdk.Int64(12)},
			contains:    []string{"fint should not exceed 10", "missing required field FieldBool"},
			notContains: []string{"fstring"},
			err:         true,
		},
		{
			stru:        &TestStruct{FieldString: awssdk.String(""), FieldInt64: awssdk.Int64(9), FieldBool: awssdk.Bool(true)},
			contains:    []string{"fstring should not be empty"},
			notContains: []string{"fint"},
			err:         true,
		},
		{
			stru:        &TestStruct{FieldString: awssdk.String("non empty"), FieldInt64: awssdk.Int64(8)},
			contains:    []string{"missing required field FieldBool"},
			notContains: []string{"fint", "fstring"},
			err:         true,
		},
		{
			stru: &TestStruct{FieldString: awssdk.String("non empty"), FieldInt64: awssdk.Int64(8), FieldBool: awssdk.Bool(true)},
			err:  false,
		},
	}

	for i, tcase := range tcases {
		err := validateStruct(tcase.stru)
		if tcase.err && err == nil {
			t.Fatalf("%d. expected err got none", i+1)
		}
		if !tcase.err && err != nil {
			t.Fatalf("%d. expected no err got one", i+1)
		}
		for _, msg := range tcase.contains {
			if got, want := err.Error(), msg; !strings.Contains(got, want) {
				t.Fatalf("%d. %q should contains %q", i+1, got, want)
			}
		}
		for _, msg := range tcase.notContains {
			if err != nil {
				if got, want := err.Error(), msg; strings.Contains(got, want) {
					t.Fatalf("%d. %q should not contains %q", i+1, got, want)
				}
			}
		}
	}
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
