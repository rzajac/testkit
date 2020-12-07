package testkit

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ErrWriter(t *testing.T) {
	// --- Given ---
	dst := &bytes.Buffer{}

	// --- When ---
	ew := ErrWriter(dst, 3, nil)
	n, err := ew.Write([]byte{0, 1, 2, 3, 4})

	// --- Then ---
	assert.ErrorIs(t, err, ErrTestError)
	assert.Exactly(t, 3, n)
	assert.Exactly(t, []byte{0, 1, 2}, dst.Bytes())
}

func Test_ErrWriter_NoError(t *testing.T) {
	// --- Given ---
	dst := &bytes.Buffer{}

	// --- When ---
	ew := ErrWriter(dst, 3, nil)
	n, err := ew.Write([]byte{0, 1})

	// --- Then ---
	assert.NoError(t, err)
	assert.Exactly(t, 2, n)
	assert.Exactly(t, []byte{0, 1}, dst.Bytes())
}

func Test_ErrWriter_CustomError(t *testing.T) {
	// --- Given ---
	dst := &bytes.Buffer{}
	ce := errors.New("my error")

	// --- When ---
	ew := ErrWriter(dst, 3, ce)
	n, err := ew.Write([]byte{0, 1, 2})

	// --- Then ---
	assert.ErrorIs(t, err, ce)
	assert.Exactly(t, 3, n)
	assert.Exactly(t, []byte{0, 1, 2}, dst.Bytes())
}

func ExampleErrWriter() {
	dst := &bytes.Buffer{}
	ew := ErrWriter(dst, 3, errors.New("my error"))

	n, err := ew.Write([]byte{0, 1, 2})

	fmt.Println(n)
	fmt.Println(err)
	fmt.Println(dst.Bytes())

	// Output:
	// 3
	// my error
	// [0 1 2]
}
