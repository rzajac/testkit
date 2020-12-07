package testkit_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	kit "testkit"
	"testkit/internal"
)

func Test_HTTPServer_SmokeTest(t *testing.T) {
	// --- Given ---
	mck := &internal.TMock{}
	mck.On("Cleanup", mock.Anything)
	mck.On("Helper")

	srv := kit.NewHTTPServer(mck).Rsp(http.StatusOK, []byte("response"))
	defer srv.Close()

	// --- When ---
	rsp, err := http.Post(srv.URL()+"/?k0=v0", "", bytes.NewReader([]byte("req body")))

	// --- Then ---
	mck.AssertExpectations(t)
	require.NoError(t, err)
	assert.Exactly(t, http.StatusOK, rsp.StatusCode)
	assert.Exactly(t, []byte("response"), getResponseBody(t, rsp))
	assert.Exactly(t, 1, srv.ReqCount())
	assert.Exactly(t, "req body", string(srv.Body(0)))
	require.Len(t, srv.Values(0), 1)
	assert.Exactly(t, "v0", srv.Values(0).Get("k0"))

	req := srv.Request(0)
	assert.Exactly(t, srv.URL()+"/?k0=v0", req.URL.String())
	assert.Exactly(t, "req body", string(getRequestBody(t, req)))
}

// getResponseBody reads response body, closes it and returns as byte slice. Calls
// t.Fatal() on error.
func getRequestBody(t *testing.T, r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}
	if err := r.Body.Close(); err != nil {
		t.Fatal(err)
	}
	return body
}

// getResponseBody reads response body, closes it and returns as byte slice. Calls
// t.Fatal() on error.
func getResponseBody(t *testing.T, r *http.Response) []byte {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}
	if err := r.Body.Close(); err != nil {
		t.Fatal(err)
	}
	return body
}
