package testkit

import (
	"bytes"
	"encoding/json"
	"io"
)

// ToJSON is a wrapper around json.MarshalIndent() which calls t.Fatal()
// on error.
func ToJSON(t T, v interface{}) []byte {
	t.Helper()
	data, err := json.MarshalIndent(v, "  ", "  ")
	if err != nil {
		t.Fatal(err)
	}
	return data
}

// ToJSONReader marshals v to JSON and returns io.Reader for it.
// Calls t.Fatal() on error.
func ToJSONReader(t T, v interface{}) io.Reader {
	t.Helper()
	return bytes.NewReader(ToJSON(t, v))
}

// FromJSON is a wrapper around json.Unmarshal() which calls t.Fatal() on error.
func FromJSON(t T, data []byte, v interface{}) {
	t.Helper()
	if err := json.Unmarshal(data, v); err != nil {
		t.Fatal(err)
	}
}

// FromJSONReader reads JSON from r and unmarshalls them to v. Calls t.Fatal()
// on error.
func FromJSONReader(t T, r io.Reader, v interface{}) {
	t.Helper()
	if err := json.NewDecoder(r).Decode(v); err != nil {
		t.Fatal(err)
	}
}
