package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

const pathIntegrations = "/api/integrations"

func (s *Server) routeIntegrations(r *mux.Router) {
	r.HandleFunc(pathIntegrations, s.handleIntegrations).Methods(http.MethodPost)
}

func (s *Server) handleIntegrations(w http.ResponseWriter, r *http.Request) {
	s.logRequest(r)
}
