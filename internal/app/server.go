package app

import (
	"crypto/ed25519"
)

// Server of the apps
type Server struct {
	port uint64

	Authorizer
	Authenticator

	discordPublicKey ed25519.PublicKey
	WhitelistGetter  WhitelistGetter
}

// NewServer creates a new server
func NewServer(opts ...ServerOption) Server {
	s := Server{}
	s.Set(opts...)
	return s
}

// IsValid checks a servers configuration
func (s *Server) IsValid() error {
	return nil
}

// Set a servers configuration
func (s *Server) Set(opts ...ServerOption) {
	for _, opt := range opts {
		opt(s)
	}
}
