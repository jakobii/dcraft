package app

import (
	"crypto/ed25519"
	"encoding/hex"
)

// ServerOption is an option server parameter
type ServerOption func(*Server)

// WithDiscordPublicKey is a ServerOption
func WithDiscordPublicKey(key string) (ServerOption, error) {
	b, err := hex.DecodeString(key)
	if err != nil {
		return nil, err
	}
	return func(s *Server) {
		s.discordPublicKey = ed25519.PublicKey(b)
	}, nil
}

// WithWhitelist add a whitelist.json file path the server
func WithWhitelist(w Whitelister) ServerOption {
	return func(s *Server) {
		s.WhitelistGetter = w
	}
}

// WithPort server http port
func WithPort(port uint64) ServerOption {
	return func(s *Server) {
		s.port = port
	}
}

// WithAuthorizer  ..
func WithAuthorizer(a Authorizer) ServerOption {
	return func(s *Server) {
		s.Authorizer = a
	}
}

// WithAuthenticator ..
func WithAuthenticator(a Authenticator) ServerOption {
	return func(s *Server) {
		s.Authenticator = a
	}
}
