package testkit

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ErrReader(t *testing.T) {
	// --- Given ---
	src := bytes.NewReader([]byte{0, 1, 2, 3})
	dst := make([]byte, 3)

	// --- When ---
	er := ErrReader(src, 3, nil)
	n, err := er.Read(dst)

	// --- Then ---
	assert.ErrorIs(t, err, ErrTestError)
	assert.Exactly(t, 3, n)
	assert.Exactly(t, []byte{0, 1, 2}, dst)
}

func Test_ErrReader_NoError(t *testing.T) {
	// --- Given ---
	src := bytes.NewReader([]byte{0, 1, 2, 3})
	dst := make([]byte, 2)

	// --- When ---
	er := ErrReader(src, 3, nil)
	n, err := er.Read(dst)

	// --- Then ---
	assert.NoError(t, err)
	assert.Exactly(t, 2, n)
	assert.Exactly(t, []byte{0, 1}, dst)
}

func Test_ErrReader_CustomError(t *testing.T) {
	// --- Given ---
	src := bytes.NewReader([]byte{0, 1, 2, 3})
	dst := make([]byte, 3)
	ce := errors.New("my error")

	// --- When ---
	er := ErrReader(src, 3, ce)
	n, err := er.Read(dst)

	// --- Then ---
	assert.ErrorIs(t, err, ce)
	assert.Exactly(t, 3, n)
	assert.Exactly(t, []byte{0, 1, 2}, dst)
}

func ExampleErrReader() {
	r := bytes.NewReader([]byte{0, 1, 2, 3})
	er := ErrReader(r, 2, nil)

	got, err := ioutil.ReadAll(er)

	fmt.Println(got)
	fmt.Println(err)

	// Output:
	// [0 1]
	// testkit test error
}
