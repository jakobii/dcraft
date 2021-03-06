package auth

import (
	"net/http"
	"strings"
)

// GetToken ...
func (c *Cache) GetToken(secret string) (Token, bool) {
	c.RLock()
	defer c.RUnlock()
	token, ok := c.tokens[secret]
	if !ok {
		return token, ok
	}
	return token, true
}

// PutToken ...
func (c *Cache) PutToken(tokens ...Token) {
	c.Lock()
	defer c.Unlock()
	for _, t := range tokens {
		c.tokens[t.secret] = t
	}
}

// DelToken ...
func (c *Cache) DelToken(secret string) {
	c.Lock()
	defer c.Unlock()
	delete(c.tokens, secret)
}

// CheckToken ..
func (c *Cache) CheckToken(secret string) bool {
	if token, ok := c.GetToken(secret); ok {
		if token.IsValid() {
			return true
		}
		c.DelToken(secret)
	}
	return false
}

// AuthorizeToken ...
func (c *Cache) AuthorizeToken(r *http.Request) bool {
	if t, ok := parseTokenQuery(r); ok {
		return c.CheckToken(t)
	}
	if t, ok := parseTokenHeader(r); ok {
		return c.CheckToken(t)
	}
	return false
}

func parseTokenHeader(r *http.Request) (string, bool) {
	header := r.Header.Get("Authorization")
	parts := strings.Split(header, " ")
	if len(parts) != 2 {
		return "", false
	}
	if parts[0] != "Bearer" {
		return "", false
	}
	return parts[1], true
}

func parseTokenQuery(r *http.Request) (string, bool) {
	q := r.URL.Query()
	if t, ok := q["token"]; ok {
		return t[0], true
	}
	return "", false
}
