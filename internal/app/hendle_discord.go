package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

const pathDiscord = "/discord"

func (s *Server) routeDiscord(r *mux.Router) {
	r.HandleFunc(pathDiscord, s.handleDiscord).Methods(http.MethodPost)
}

func (s *Server) handleDiscord(w http.ResponseWriter, r *http.Request) {
	s.logRequest(r)
}
