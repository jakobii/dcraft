package auth

import (
	"encoding/base64"
	"net/http"
	"strings"
)

// GetCredential ...
func (c *Cache) GetCredential(username string) (Credential, bool) {
	c.RLock()
	defer c.RUnlock()
	cred, ok := c.creds[username]
	if !ok {
		return cred, ok
	}
	return cred, true
}

// PutCredential ...
func (c *Cache) PutCredential(creds ...Credential) {
	c.Lock()
	defer c.Unlock()
	for _, cred := range creds {
		c.creds[cred.Username] = cred
	}
}

// DelCredential ...
func (c *Cache) DelCredential(username string) {
	c.Lock()
	defer c.Unlock()
	delete(c.creds, username)
}

// CheckCredential ..
func (c *Cache) CheckCredential(username, password string) bool {
	if cred, ok := c.GetCredential(username); ok {
		if cred.Validate(password) {
			return true
		}
	}
	return false
}

// AuthorizeCredential ..
func (c *Cache) AuthorizeCredential(r *http.Request) bool {
	if u, p, ok := parseCredentialQuery(r); ok {
		return c.CheckCredential(u, p)
	}
	if u, p, ok := parseCredentialHeader(r); ok {
		return c.CheckCredential(u, p)
	}
	return false
}

func parseCredentialQuery(r *http.Request) (username string, password string, ok bool) {
	q := r.URL.Query()
	u := q.Get("u")
	p := q.Get("p")
	if u != "" && p != "" {
		return u, p, true
	}
	return "", "", false
}

func parseCredentialHeader(r *http.Request) (username string, password string, ok bool) {
	header := r.Header.Get("Authorization")
	parts := strings.Split(header, " ")
	if len(parts) != 2 {
		return "", "", false
	}
	if parts[0] != "Basic" {
		return "", "", false
	}
	b, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return "", "", false
	}
	cred := strings.Split(string(b), ":")
	if len(cred) != 2 {
		return "", "", false
	}
	return cred[0], cred[1], true
}
