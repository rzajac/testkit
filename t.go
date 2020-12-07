package testkit

// T is a subset of testing.TB interface.
// It's primarily used to test the testkit package but can be used to
// implement custom actions to be taken on errors.
type T interface {
	// Cleanup registers a function to be called when the test and all its
	// subtests complete. Cleanup functions will be called in last added,
	// first called order.
	Cleanup(func())

	// Error is equivalent to Log followed by Fail.
	Error(args ...interface{})

	// Errorf is equivalent to Logf followed by Fail.
	Errorf(format string, args ...interface{})

	// Fatal is equivalent to Log followed by FailNow.
	Fatal(args ...interface{})

	// Fatalf is equivalent to Logf followed by FailNow.
	Fatalf(format string, args ...interface{})

	// Helper marks the calling function as a test helper function.
	// When printing file and line information, that function will be skipped.
	// Helper may be called simultaneously from multiple goroutines.
	Helper()

	// Skip is equivalent to Log followed by SkipNow.
	Skip(args ...interface{})
}
