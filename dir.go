package testkit

import (
	"io/ioutil"
	"os"
)

// TempDir is a wrapper around ioutil.TempDir() which calls t.Fatal() on error.
func TempDir(t T, dir string, prefix string) string {
	t.Helper()
	pth, err := ioutil.TempDir(dir, prefix)
	if err != nil {
		t.Fatal(err)
		return ""
	}
	return pth
}

// CreateDir creates directory with name returns passed name, so it can be used
// in place.
//
//     MyFunction(CreateDir(t, "name"))
//
// Calls t.Fatal() on error.
func CreateDir(t T, name string) string {
	t.Helper()
	if err := os.Mkdir(name, 0700); err != nil {
		t.Fatal(err)
		return ""
	}
	return name
}
