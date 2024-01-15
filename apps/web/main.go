package web

// This file just ties together the service with
// the endpoints, it should not exist after the
// services are deployed independently

import (
	"net/http"

	"github.com/amonaco/goapi/apps/web/services/company"
	"github.com/amonaco/goapi/apps/web/services/user"
)

func Company() http.Handler {
	return company.New()
}

func User() http.Handler {
	return user.New()
}
