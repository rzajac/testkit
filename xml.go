package testkit

import (
	"encoding/xml"
)

// FromXML is a wrapper around xml.Unmarshal() which calls t.Fatal() on error.
func FromXML(t T, data []byte, v interface{}) {
	t.Helper()
	if err := xml.Unmarshal(data, v); err != nil {
		t.Fatal(err)
	}
}

// ToXML is a wrapper around xml.Marshal() which calls t.Fatal() on error.
func ToXML(t T, v interface{}) []byte {
	t.Helper()
	data, err := xml.Marshal(v)
	if err != nil {
		t.Fatal(err)
		return nil
	}
	return data
}
