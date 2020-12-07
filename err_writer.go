package testkit

import (
	"io"
)

// errWriter implements io.Writer interface but allows to write only n
// number of bytes then returns given error.
type errWriter struct {
	w   io.Writer // Underlying writer.
	off int       // Number of written bytes.
	n   int       // At most bytes to write without error.
	err error     // Error to return.
}

// ErrWriter wraps writer w in a writer which writes at most n bytes and
// returns error err. If err is set to nil then ErrTestError will be used.
// Error other then err will be returned if underlying writer returns
// an error before n bytes are written.
func ErrWriter(w io.Writer, n int, err error) io.Writer {
	if err == nil {
		err = ErrTestError
	}
	return &errWriter{
		w:   w,
		n:   n,
		err: err,
	}
}

// Write writes to underlying buffer and returns error if number of written
// bytes is equal to predefined limit.
func (w *errWriter) Write(p []byte) (int, error) {
	// Write no more then n bytes.
	if w.off+len(p) > w.n {
		p = p[:w.n-w.off]
	}
	n, err := w.w.Write(p)
	w.off += n
	if err != nil {
		return n, err
	}
	if w.off >= w.n {
		return n, w.err
	}
	return n, nil
}
