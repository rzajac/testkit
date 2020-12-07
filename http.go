package testkit

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
)

// CloneHTTPRequest clones HTTP request with body, URL.Host and URL.Scheme.
// Calls t.Fatal() on error.
func CloneHTTPRequest(t T, req *http.Request) *http.Request {
	t.Helper()
	c := req.Clone(context.Background())
	if req.Body != nil {
		var buf bytes.Buffer
		body, err := ioutil.ReadAll(io.TeeReader(req.Body, &buf))
		if err != nil {
			t.Fatal(err)
			return nil
		}
		if err := req.Body.Close(); err != nil {
			t.Fatal(err)
			return nil
		}
		req.Body = ioutil.NopCloser(bytes.NewReader(buf.Bytes()))
		c.Body = ioutil.NopCloser(bytes.NewReader(body))
	}
	c.URL.Host = req.Host
	c.URL.Scheme = req.URL.Scheme

	return c
}
