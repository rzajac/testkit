package testkit

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// MD5File returns MD5 hash of the file. Calls t.Fatal() on error.
func MD5File(t T, pth string) string {
	fil, err := os.Open(pth)
	if err != nil {
		t.Fatal(err)
		return ""
	}
	defer fil.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, fil); err != nil {
		t.Fatal(err)
		return ""
	}
	return hex.EncodeToString(hash.Sum(nil))
}

// MD5Reader returns MD5 hash off everything in the reader. Calls t.Fatal()
// on error.
func MD5Reader(t T, r io.Reader) string {
	hash := md5.New()
	if _, err := io.Copy(hash, r); err != nil {
		t.Fatal(err)
		return ""
	}
	return hex.EncodeToString(hash.Sum(nil))
}
