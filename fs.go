package testkit

import (
	"os"
	"time"
)

// ModTime calls os.Stat() and returns modification time. Calls t.Fatal()
// on error.
func ModTime(t T, name string) time.Time {
	t.Helper()
	fi, err := os.Stat(name)
	if err != nil {
		t.Fatal(err)
		return time.Time{}
	}
	return fi.ModTime()
}
