package auth

// This file just ties together the service with
// the endpoints, it should not exist after the
// services are deployed independently

import (
	"net/http"

	"github.com/amonaco/goapi/apps/auth/services/auth"
)

func New() http.Handler {
	return auth.New()
}
