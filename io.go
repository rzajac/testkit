package testkit

import (
	"bytes"
	"io"
	"io/ioutil"
)

// ReadAll is a wrapper around ioutil.ReadAll which calls t.Fatal() on error.
func ReadAll(t T, r io.Reader) []byte {
	t.Helper()
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
		return nil
	}
	return bs
}

// ReadAllString is a wrapper around ioutil.ReadAll which returns string
// instead of byte slice. Calls t.Fatal() on error.
func ReadAllString(t T, r io.Reader) string {
	t.Helper()
	return string(ReadAll(t, r))
}

// ReadAllFromStart seeks to the beginning of rs and reads it till gets
// io.EOF or any other error. Then seeks back to the position where rs
// was before the call. Calls t.Fatal() on error.
func ReadAllFromStart(t T, rs io.ReadSeeker) []byte {
	t.Helper()
	cur, err := rs.Seek(0, io.SeekCurrent)
	if err != nil {
		t.Fatal(err)
		return nil
	}

	if _, err := rs.Seek(0, io.SeekStart); err != nil {
		t.Fatal(err)
		return nil
	}

	defer func() { _, _ = rs.Seek(cur, io.SeekStart) }()

	ret := &bytes.Buffer{}
	if _, err := ret.ReadFrom(rs); err != nil {
		t.Fatal(err)
		return nil
	}

	return ret.Bytes()
}
