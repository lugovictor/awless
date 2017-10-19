package awsspec

import (
	"errors"
	"fmt"
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
		rules      *paramRule
		keys       []string
		expErr     error
		expMissing []string
	}{
		{allOf(node("one"), node("two")), []string{"one", "two"}, nil, nil},
		{allOf(node("one"), node("two")), []string{"one", "other"}, errors.New("unexpected param key(s) 'other'"), nil},
		{allOf(node("one"), node("two")), []string{"one"}, nil, []string{"two"}},
		{allOf(oneOf(node("role"), node("user"), node("group")), oneOf(node("arn"), allOf(node("access"), node("service")))), []string{"role", "arn"}, nil, nil},
		{allOf(oneOf(node("role"), node("user"), node("group")), oneOf(node("arn"), allOf(node("access"), node("service")))), []string{"role", "service"}, nil, []string{"access"}},
		{allOf(oneOf(node("role"), node("user"), node("group")), oneOf(node("arn"), allOf(node("access"), node("service")))), []string{}, nil, []string{"role", "arn"}},
	}
	for i, tcase := range tcases {
		fmt.Println(i, ":", tcase.rules.hint())
		missing, err := tcase.rules.verify(tcase.keys)
		if tcase.expErr == nil {
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
			if got, want := err.Error(), tcase.expErr.Error(); !strings.Contains(got, want) {
				t.Fatalf("%d: got %s, want %s", i+1, got, want)
			}
		}
	}
}
