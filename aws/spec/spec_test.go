package awsspec

import (
	"reflect"
	"strings"
	"testing"
)

func checkErrs(t *testing.T, errs []error, length int, expected ...string) {
	t.Helper()
	if got, want := len(errs), length; got != want {
		t.Fatalf("got %d want %d", got, want)
	}
	for _, exp := range expected {
		var found bool
		for _, err := range errs {
			if strings.Contains(err.Error(), exp) {
				found = true
			}
		}
		if !found {
			t.Fatalf("'%s' should be in errors %v", exp, errs)
		}
	}
}

func TestValidationDSL(t *testing.T) {
	tcases := []struct {
		rule           ruleNode
		extras         []string
		userInput      []string
		expErrContains []string
		expMissing     []string
	}{
		{nil, nil, []string{"one", "two"}, nil, nil},
		{allOf(), nil, []string{"one", "two"}, []string{"unexpected", "one", "two"}, nil},
		{allOf(), []string{"one", "two", "three"}, []string{"one", "two"}, nil, nil},
		{allOf(node("one"), node("two")), nil, []string{"one", "two"}, nil, nil},
		{allOf(node("one"), node("two")), nil, []string{"one", "other"}, []string{"unexpected", "other"}, nil},
		{allOf(node("one"), node("two")), nil, []string{"one"}, nil, []string{"two"}},
		{allOf(node("one"), node("two")), []string{"three"}, []string{"one", "three"}, nil, []string{"two"}},
		{allOf(oneOf(node("role"), node("user"), node("group")), oneOf(node("arn"), allOf(node("access"), node("service")))), nil, []string{"role", "arn"}, nil, nil},
		{allOf(oneOf(node("role"), node("user"), node("group")), oneOf(node("arn"), allOf(node("access"), node("service")))), nil, []string{"role", "service"}, nil, []string{"access"}},
		{allOf(oneOf(node("role"), node("user"), node("group")), oneOf(node("arn"), allOf(node("access"), node("service")))), nil, []string{}, nil, []string{"role", "arn"}},
	}
	for i, tcase := range tcases {
		missing, err := paramRule{tree: tcase.rule, extras: tcase.extras}.verify(tcase.userInput)
		if len(tcase.expErrContains) == 0 {
			if err != nil {
				t.Fatalf("%d: %s", i+1, err)
			}
			if got, want := missing, tcase.expMissing; !reflect.DeepEqual(got, want) {
				t.Fatalf("%d: got %#v, want %#v", i+1, got, want)
			}
		} else {
			if err == nil {
				t.Fatalf("%d: expected err, got nil", i+1)
			}
			for _, each := range tcase.expErrContains {
				if got, want := err.Error(), each; !strings.Contains(got, want) {
					t.Fatalf("%d: got '%s', want '%s'", i+1, got, want)
				}
			}
		}
	}

}
