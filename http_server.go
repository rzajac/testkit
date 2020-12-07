package testkit

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
)

// response represents a response returned by HTTPServer.
type response struct {
	status int    // HTTP status code.
	body   []byte // Response body.
}

// HTTPServer represents very simple HTTP server recording all the requests it
// receives, responding with responses added with Rsp() method.
//
// The server instance provides methods to analyse the received requests.
type HTTPServer struct {
	requests  []*http.Request  // Received requests.
	srv       *httptest.Server // The test server.
	host      string           // Test server host:port.
	scheme    string           // Test server scheme.
	responses []response       // Responses to return.
	t         T                // Test state manager.
}

// NewHTTPServer returns new instance of HTTPServer.
func NewHTTPServer(t T) *HTTPServer {
	t.Helper()
	tst := &HTTPServer{
		t: t,
	}

	// Cleanup after the test is done.
	t.Cleanup(func() { tst.Close() })

	// Handler for all incoming requests.
	handler := func(w http.ResponseWriter, req *http.Request) {
		rsp := tst.next()
		w.WriteHeader(rsp.status)

		c := CloneHTTPRequest(t, req)
		c.URL.Scheme = tst.scheme

		tst.requests = append(tst.requests, c)
		if _, err := w.Write(rsp.body); err != nil {
			t.Fatal(err)
		}
	}

	tst.srv = httptest.NewServer(http.HandlerFunc(handler))
	u, err := url.Parse(tst.srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	tst.host = u.Host
	tst.scheme = u.Scheme

	return tst
}

// Rsp adds response with status and body to the list of responses to return.
//
// Every time server receives a request it returns predefined response. The
// responses are returned in the order they were added. If there is no more
// responses the t.Fatal() will be called.
func (tst *HTTPServer) Rsp(status int, rsp []byte) *HTTPServer {
	tst.responses = append(tst.responses, response{
		status: status,
		body:   rsp,
	})
	return tst
}

// URL returns URL for the test server.
func (tst *HTTPServer) URL() string { return tst.srv.URL }

// Request returns clone of the nth received request. Calls t.Fatal() if n is
// greater then number or received requests.
func (tst *HTTPServer) Request(n int) *http.Request {
	tst.t.Helper()
	if n >= 0 && n < len(tst.requests) {
		return CloneHTTPRequest(tst.t, tst.requests[n])
	}
	tst.t.Fatalf("no request with index %d recorded", n)
	return nil
}

// ReqCount returns number of requests recorded by test server.
func (tst *HTTPServer) ReqCount() int { return len(tst.requests) }

// Values returns URL query values of the nth received request. Calls t.Fatal()
// if n is greater then number or received requests.
func (tst *HTTPServer) Values(n int) url.Values {
	tst.t.Helper()
	if n >= 0 && n < len(tst.requests) {
		return tst.requests[n].URL.Query()
	}
	tst.t.Fatalf("no request with index %d recorded", n)
	return url.Values{}
}

// Body returns body of the nth received request. Calls t.Fatal() if n is
// greater then number or received requests.
func (tst *HTTPServer) Body(n int) []byte {
	tst.t.Helper()
	if n >= 0 && n < len(tst.requests) {
		req := tst.requests[n]
		var buf bytes.Buffer
		body, err := ioutil.ReadAll(io.TeeReader(req.Body, &buf))
		if err != nil {
			tst.t.Fatal(err)
			return nil
		}
		if err = req.Body.Close(); err != nil {
			tst.t.Fatal(err)
			return nil
		}
		req.Body = ioutil.NopCloser(bytes.NewReader(buf.Bytes()))
		return body
	}
	tst.t.Fatalf("no request with index %d recorded", n)
	return nil
}

// Headers returns headers for given request index. Calls t.Fatal() if n is
// greater then number or received requests.
func (tst *HTTPServer) Headers(n int) http.Header {
	tst.t.Helper()
	if n >= 0 && n < len(tst.requests) {
		return tst.requests[n].Header
	}
	tst.t.Fatalf("no request with index %d recorded", n)
	return nil
}

// next returns the next response to return. Calls t.Fatal() if no more
// responses to give.
func (tst *HTTPServer) next() response {
	tst.t.Helper()
	var rsp response
	if len(tst.responses) == 0 {
		tst.t.Fatal("no more responses to give")
		return rsp
	}
	rsp, tst.responses = tst.responses[0], tst.responses[1:]
	return rsp
}

// Close stops the test server and does cleanup. May be called multiple times.
func (tst *HTTPServer) Close() error {
	tst.srv.Close()
	for _, req := range tst.requests {
		_ = req.Body.Close()
	}
	tst.requests = tst.requests[:0]
	tst.responses = tst.responses[:0]
	return nil
}
