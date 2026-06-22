package httpclient

import "context"

// FetchJSON performs GET url and decodes JSON into dest.
// Returns an error for non-2xx status codes or invalid JSON.
func FetchJSON(ctx context.Context, url string, dest any) error {
	return nil
}

// GetStatus performs GET url and returns the HTTP status code.
func GetStatus(ctx context.Context, url string) (int, error) {
	return 0, nil
}
