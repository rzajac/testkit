## Test Kit

[![Go Report Card](https://goreportcard.com/badge/github.com/rzajac/testkit)](https://goreportcard.com/report/github.com/rzajac/testkit)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/rzajac/testkit)](https://pkg.go.dev/github.com/rzajac/testkit)

Package testkit provides all sorts of test helpers which take `*testing.T` 
instance and report errors if something goes wrong. See GoDoc for the
documentation and examples.

## Motivation

In my opinion constantly checking errors while setting up a test reduces its
readability.

Example test:

```
func Test_MyFunction(t *testing.T) {
	// --- Given ---
	fil, err := ioutil.TempFile(t.TempDir(), "test_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer fil.Close()

	content := []byte("line1\nline2\nend")
	if _, err := fil.Write(content); err != nil {
		t.Fatal(err)
	}

	// --- When ---

	// Do some operations on the test file.

	// --- Then ---
	chk, err := os.Open(fil.Name())
	if err != nil {
		t.Fatal(err)
	}

	got, err := ioutil.ReadAll(chk)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0, 1, 2, 3}
	if !bytes.Equal(expected, got) {
		t.Errorf("expected %v got %v", expected, got)
	}
}
```

The same test using testkit:

```

import (
    kit "github.com/rzajac/testkit"
)

func Test_MyFunction(t *testing.T) {
	// --- Given ---
	content := []byte("line1\nline2\nend")
	pth := kit.WriteTempFile(t, t.TempDir(), bytes.NewReader(content))

	// --- When ---

	// Do some operations on the test file.

	// --- Then ---
	got := kit.ReadFile(t, pth)
	expected := []byte{0, 1, 2, 3}
	if !bytes.Equal(expected, got) {
		t.Errorf("expected %v got %v", expected, got)
	}
}
```

`WriteTempFile` creates and writes to temporary file from a reader r. Returns
path to created file. It registers cleanup function with t removing the created
file. Calls t.Fatal() on error. Similarly `ReadFile` is a wrapper
around `ioutil.ReadFile()` which calls t.Fatal() on error.

## License

BSD-2-Clause