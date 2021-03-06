package app

import "net/http"

// Authorizer of http requests
type Authorizer interface {
	Authorize(r *http.Request) (ok bool)
}

// Authenticator of http requests
type Authenticator interface {
	Authenticate(r *http.Request) (token string, ok bool)
}
