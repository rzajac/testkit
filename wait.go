package testkit

import (
	"context"
	"time"
)

// WaitTh waits for fn to return true but no longer then timeout max. The
// calls to fn are throttled with th.
func WaitTh(max, th time.Duration, fn func() bool) {
	ctx, cxl := context.WithTimeout(context.Background(), max)
	defer cxl()

	throttle := time.NewTimer(th)
	defer throttle.Stop()

	for {
		throttle.Reset(th)
		select {
		case <-ctx.Done():
			return
		default:
			if fn() {
				return
			}
			<-throttle.C
		}
	}
}

// Wait waits till fn returns true but no longer then timeout max.
func Wait(max time.Duration, fn func() bool) {
	ctx, cxl := context.WithTimeout(context.Background(), max)
	defer cxl()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if fn() {
				return
			}
		}
	}
}
