package contextlesson

import (
	"context"
	"time"
)

// FetchWithTimeout GETs url and returns the response body.
// When the request exceeds d, returns an error wrapping context.DeadlineExceeded.
func FetchWithTimeout(url string, d time.Duration) ([]byte, error) {
	return nil, nil
}

// FetchWithContext GETs url using ctx for cancellation.
func FetchWithContext(ctx context.Context, url string) ([]byte, error) {
	return nil, nil
}
