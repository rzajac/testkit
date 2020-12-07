package testkit

import (
	"io"
)

// errReader represents a reader which reads at most n bytes
// and returns error err.
type errReader struct {
	r   io.Reader // Underlying error.
	n   int       // At most bytes to read without error.
	off int       // Number of bytes read.
	err error     // Error to return after reading n bytes.
}

// ErrReader wraps reader r. The returned  reader reads at most n bytes
// and returns error err. If err is set to nil ErrTestError will be used.
// Error other then err will be returned if underlying reader returns an
// error before n bytes are read.
func ErrReader(r io.Reader, n int, err error) io.Reader {
	if err == nil {
		err = ErrTestError
	}
	return &errReader{
		r:   r,
		n:   n,
		err: err,
	}
}

// Read implements io.Reader which returns error after reading n bytes.
func (r *errReader) Read(p []byte) (int, error) {
	// Read only up to the limit.
	if r.off+len(p) > r.n {
		p = p[:r.n-r.off]
	}
	n, err := r.r.Read(p)
	r.off += n
	if err != nil {
		return n, err
	}
	if r.off >= r.n {
		return n, r.err
	}
	return n, nil
}
