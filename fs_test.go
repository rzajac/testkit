package testkit

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ModTime(t *testing.T) {
	// --- Given ---
	pth := filepath.Join(t.TempDir(), "test_file.txt")
	_, err := os.Create(pth)
	require.NoError(t, err)

	mck := &TMock{}
	mck.On("Helper")

	// --- When ---
	mt := ModTime(mck, pth)

	// --- Then ---
	mck.AssertExpectations(t)
	assert.True(t, time.Now().Sub(mt) < 10*time.Millisecond)
}
