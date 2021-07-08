package middleware

import "github.com/harmony-development/legato/server/db"

// Middlewares contains all dependencies for the middleware.
type Middlewares struct {
	AuthDB db.IAuthDB
}

// New creates a new instance of the middleware.
func New() *Middlewares {
	return &Middlewares{}
}
