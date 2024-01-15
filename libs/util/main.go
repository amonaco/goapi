package util

import (
	"net/http"

	"github.com/go-chi/chi"
)

// We use this middleware to rewrite the URL
// Keeps compatibility with WebRPC's routing
func RouteToPath(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		r.URL.Path = rctx.RoutePath
		next.ServeHTTP(w, r)
	})
}
