package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

const pathPing = "/ping"

func (s *Server) routePing(r *mux.Router) {
	r.HandleFunc(pathPing, s.handlePing).Methods(http.MethodGet)
}

func (s *Server) handlePing(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
