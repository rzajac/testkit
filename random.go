package testkit

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"time"
)

// Make sure random number generator is seeded.
func init() { rand.Seed(time.Now().UnixNano()) }

// letters list of valid alphabet characters.
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandStr returns random string of length n.
func RandStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// RandFileName returns random file name. If prefix is empty string it will be
// set to "file-". The extension wll be set to ".txt" if ext is empty string.
func RandFileName(dir, prefix, ext string) string {
	if prefix == "" {
		prefix = "file-"
	}
	if ext == "" {
		ext = ".txt"
	}

	b := make([]byte, 7)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return filepath.Join(dir, fmt.Sprintf("%s%s%s", prefix, string(b), ext))
}
