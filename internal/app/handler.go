package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Handler is the servers main Handler
func (s *Server) Handler() http.Handler {
	r := mux.NewRouter()

	routes := []routeHandler{
		s.routeDiscord,
		s.routeIntegrations,
	}
	for _, route := range routes {
		route(r)
	}

	return s.Middleware(r)
}

type routeHandler func(*mux.Router)
