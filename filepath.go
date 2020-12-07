package testkit

import (
	"path/filepath"
)

// AbsPath is a wrapper around filepath.Abs() which calls t.Fatal() on error.
func AbsPath(t T, pth string) string {
	t.Helper()
	pth, err := filepath.Abs(pth)
	if err != nil {
		t.Fatal(err)
	}
	return pth
}
