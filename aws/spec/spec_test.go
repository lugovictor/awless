package awsspec

import (
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
