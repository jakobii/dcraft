package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const pathToken = "/tokens"

func (s *Server) routeToken(r *mux.Router) {
	r.HandleFunc(pathToken, s.handleToken).Methods(http.MethodPost)
}

func (s *Server) handleToken(w http.ResponseWriter, r *http.Request) {
	token, ok := s.Authenticate(r)
	if !ok {
		s.Err(w, r, http.StatusUnauthorized, fmt.Errorf("missing credential"))
		return
	}
	fmt.Fprint(w, token)
}
