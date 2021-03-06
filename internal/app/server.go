package app

import (
	"crypto/ed25519"
	"encoding/hex"
)

// Server of the apps
type Server struct {
	port             uint64
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
