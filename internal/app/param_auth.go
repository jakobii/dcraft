package app

import (
	"fmt"
	"net/http"
)

// requireAuth will write an error to the
func (s *Server) requireAuth(w http.ResponseWriter, r *http.Request) bool {
	if s.Authorize(r) {
		return true
	}
	s.Err(w, r, http.StatusUnauthorized, fmt.Errorf("Unauthorized"))
	return false
}
