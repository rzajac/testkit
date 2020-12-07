package testkit

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// OpenFile is a wrapper around os.Open() which calls t.Fatal() on error.
func OpenFile(t T, name string) *os.File {
	t.Helper()
	fil, err := os.Open(name)
	if err != nil {
		t.Fatal(err)
	}
	return fil
}

// CreateFile is a wrapper around os.Create() which calls t.Fatal() on error.
func CreateFile(t T, name string) *os.File {
	t.Helper()
	fil, err := os.Create(name)
	if err != nil {
		t.Fatal(err)
	}
	return fil
}

// TempFile is a wrapper around ioutil.TempFile() which calls t.Fatal()
// on error. It registers cleanup function with t removing the
// created file.
func TempFile(t T, dir, pattern string) *os.File {
	t.Helper()
	fil, err := ioutil.TempFile(dir, pattern)
	if err != nil {
		t.Fatal(err)
	}
	name := fil.Name()
	t.Cleanup(func() { _ = os.Remove(name) })
	return fil
}

// ReadFile is a wrapper around ioutil.ReadFile() which calls t.Fatal()
// on error.
func ReadFile(t T, filename string) []byte {
	t.Helper()
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	return buf
}

// FileSize returns file size. Calls t.Fatal() on error.
func FileSize(t T, fil *os.File) int64 {
	t.Helper()
	s, err := fil.Stat()
	if err != nil {
		t.Fatal(err)
	}
	return s.Size()
}

// ReplaceAllInFile replaces all occurrences of old sting with new
// in the filename. Calls t.Fatal() on error.
func ReplaceAllInFile(t T, filename, old, new string) {
	t.Helper()
	data := ReadFile(t, filename)
	str := strings.ReplaceAll(string(data), old, new)
	err := ioutil.WriteFile(filename, []byte(str), 0)
	if err != nil {
		t.Fatal(err)
	}
}

// Readdirnames returns slice returned by Readdirnames(0) called on fil
// instance. Calls t.Fatal() on error.
func Readdirnames(t T, fil *os.File) []string {
	t.Helper()
	names, err := fil.Readdirnames(0)
	if err != nil {
		t.Fatal(err)
	}
	return names
}

// WriteTempFile creates and writes to temporary file from reader r. Returns
// path to created file. It registers cleanup function with t removing the
// created file. Calls t.Fatal() on error.
func WriteTempFile(t T, dir string, r io.Reader) string {
	f, err := ioutil.TempFile(dir, "")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	pth := f.Name()
	t.Cleanup(func() {
		if err := os.Remove(pth); err != nil {
			t.Fatal(err)
		}
	})
	if _, err := f.ReadFrom(r); err != nil {
		t.Fatal(err)
	}
	return pth
}
