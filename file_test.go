package testkit

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rzajac/testkit/internal"
)

func Test_ReplaceInFile(t *testing.T) {
	// --- Given ---
	pth := filepath.Join(t.TempDir(), "test_file.txt")
	content := `line1
	line2
	end
	`
	require.NoError(t, ioutil.WriteFile(pth, []byte(content), 0777))

	mck := &internal.TMock{}
	mck.On("Helper")

	// --- When ---
	ReplaceAllInFile(mck, pth, "line", "test")

	// --- Then ---
	mck.AssertExpectations(t)
	got, err := ioutil.ReadFile(pth)
	assert.NoError(t, err)

	exp := `test1
	test2
	end
	`
	assert.Exactly(t, exp, string(got))
}

func Test_CurrentOffset(t *testing.T) {
	// --- Given ---
	pth := filepath.Join(t.TempDir(), "test_file.txt")
	content := "line1\nline2\nend"
	require.NoError(t, ioutil.WriteFile(pth, []byte(content), 0777))

	fil, err := os.Open(pth)
	require.NoError(t, err)
	_, err = fil.Read(make([]byte, 3))
	require.NoError(t, err)

	mck := &internal.TMock{}
	mck.On("Helper")

	// --- When ---
	off := CurrOffset(mck, fil)

	// --- Then ---
	assert.Exactly(t, int64(3), off)
}

func Test_Seek(t *testing.T) {
	// --- Given ---
	pth := filepath.Join(t.TempDir(), "test_file.txt")
	content := "line1\nline2\nend"
	require.NoError(t, ioutil.WriteFile(pth, []byte(content), 0777))

	fil, err := os.Open(pth)
	require.NoError(t, err)

	mck := &internal.TMock{}
	mck.On("Helper")

	// --- When ---
	off := Seek(mck, fil, 4, io.SeekStart)

	// --- Then ---
	assert.Exactly(t, int64(4), off)

	got, err := ioutil.ReadAll(fil)
	require.NoError(t, err)
	assert.Exactly(t, "1\nline2\nend", string(got))
}
