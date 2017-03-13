// Package basic provides HTTP Basic Auth services.
package basic

import (
	"net/http"

	"github.com/flimzy/kivik"
	"github.com/flimzy/kivik/auth"
	"github.com/flimzy/kivik/authdb"
	"github.com/flimzy/kivik/serve"
)

// HTTPBasicAuth provides HTTP Basic Auth
type HTTPBasicAuth struct{}

var _ auth.Handler = &HTTPBasicAuth{}

// MethodName returns "default"
func (a *HTTPBasicAuth) MethodName() string {
	return "default" // For compatibility with the name used by CouchDB
}

// Authenticate authenticates a request against a user store using HTTP Basic
// Auth.
func (a *HTTPBasicAuth) Authenticate(w http.ResponseWriter, r *http.Request) (*authdb.UserContext, error) {
	store := serve.GetService(r)
	username, password, ok := r.BasicAuth()
	if !ok {
		return nil, kivik.ErrUnauthorized
	}
	return store.Validate(r.Context(), username, password)
}
