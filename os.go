package testkit

import (
	"os"
)

// Getwd is a wrapper around os.Getwd() which calls t.Fatal() on error.
func Getwd(t T) string {
	t.Helper()
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
		return ""
	}
	return wd
}
