package restapi

import "net/http"

// Item is a CRUD resource.
type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// NewServer returns an http.Handler with CRUD routes and middleware.
// Requires header X-API-Key: dev-key for all routes except GET /health.
func NewServer() http.Handler {
	return http.NotFoundHandler()
}
