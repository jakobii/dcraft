package auth

import "github.com/google/uuid"

// User ...
type User struct {
	ID uuid.UUID
	Credential
}

// Credential ..
type Credential struct {
	Username string
	password string // TODO: replace with secure hash
}

// NewCredential ..
func NewCredential(u, p string) Credential {
	return Credential{
		Username: u,
		password: p,
	}
}

// Validate ...
func (c Credential) Validate(password string) bool {
	if c.password == password {
		return true
	}
	return false
}
