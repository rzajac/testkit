package testkit

import (
	"net"
	"sync"
)

// Makes sure Internet connection is checked only once.
var chk sync.Once

// isConnected is true if the Internet connection is available.
var isConnected bool

// HasNetConn returns true if the Internet connection is present.
func HasNetConn() bool {
	chk.Do(func() {
		if _, err := net.LookupIP("www.google.com"); err == nil {
			isConnected = true
		}
	})
	return isConnected
}

// NoNetworkSkip skips test if there is no Internet connection.
func NoNetworkSkip(t T) {
	t.Helper()
	if !HasNetConn() {
		t.Skip("skipping test: no Internet connection")
	}
}
