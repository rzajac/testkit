package testkit

import (
	"os"
)

// SetEnv sets environment variable key to value and registers cleanup
// function with t to set the environment variable to what it was.
func SetEnv(t T, key, value string) {
	t.Helper()
	prev, set := os.LookupEnv(key)
	if err := os.Setenv(key, value); err != nil {
		t.Fatal(err)
		return
	}
	t.Cleanup(func() {
		if set {
			if err := os.Setenv(key, prev); err != nil {
				t.Error(err)
			}
		} else {
			if err := os.Unsetenv(key); err != nil {
				t.Error(err)
			}
		}
	})
}
