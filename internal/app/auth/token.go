package auth

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
	"time"
)

// Token ...
type Token struct {
	secret string
	expire time.Time
}

// NewToken ...
func NewToken(d time.Duration) Token {
	return Token{
		secret: NewSecret(20),
		expire: time.Now().Add(d),
	}
}

// IsValid ...
func (t *Token) IsValid() bool {
	return t.expire.After(time.Now())
}

// NewSecret ..
func NewSecret(n uint64) string {
	b := make([]byte, n)
	rand.Read(b)
	return strings.ToLower(hex.EncodeToString(b))
}
