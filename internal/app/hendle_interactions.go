package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

const pathInteractions = "/api/interactions"

func (s *Server) routeInteractions(r *mux.Router) {
	r.HandleFunc(pathInteractions, s.handleInteractions)
}

func (s *Server) handleInteractions(w http.ResponseWriter, r *http.Request) {
	s.logRequest(r)
}
