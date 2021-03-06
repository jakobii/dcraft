package auth

import (
	"net/http"
	"sync"
	"time"
)

// Cache ...
type Cache struct {
	sync.RWMutex
	super  string
	tokens map[string]Token
	creds  map[string]Credential
}

// NewCache ...
func NewCache(opts ...CacheOption) *Cache {
	c := &Cache{
		tokens: make(map[string]Token),
		creds:  make(map[string]Credential),
	}
	c.Set(opts...)
	return c
}

// Set a servers configuration
func (c *Cache) Set(opts ...CacheOption) {
	for _, opt := range opts {
		opt(c)
	}
}

// IsSuper ...
func (c *Cache) IsSuper(r *http.Request) bool {
	if c.super != "" && r.URL.Query().Get("super") == c.super {
		return true
	}
	return false
}

// Authorize implements app.Authorizer
func (c *Cache) Authorize(r *http.Request) bool {
	if c.IsSuper(r) || c.AuthorizeToken(r) || c.AuthorizeCredential(r) {
		return true
	}
	return false
}

// Authenticate implements app.Authenticator
func (c *Cache) Authenticate(r *http.Request) (token string, ok bool) {
	if c.Authorize(r) {
		t := NewToken(time.Hour * 720)
		c.PutToken(t)
		return t.secret, true
	}
	return "", false
}
