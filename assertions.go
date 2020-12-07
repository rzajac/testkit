package testkit

import (
	"regexp"
)

// AssertErrPrefix asserts error message has prefix.
func AssertErrPrefix(t T, err error, prefix string) {
	t.Helper()
	rs := "(?m)^" + prefix
	r := regexp.MustCompile(rs)

	if err == nil {
		t.Error("expected error not to be nil")
		return
	}

	msg := err.Error()
	if r.FindStringIndex(msg) == nil {
		t.Errorf("expect error %q to match %q", msg, rs)
	}
}
