package app

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

const pathWhitelist = "/whitelist"

func (s *Server) routeWhitelist(r *mux.Router) {
	r.HandleFunc(pathWhitelist, s.handleWhitelistGet).Methods(http.MethodGet)
}

func (s *Server) handleWhitelistGet(w http.ResponseWriter, r *http.Request) {
	ok := s.requireAuth(w, r)
	if !ok {
		return
	}

	users, err := s.WhitelistGetter.Get(context.Background())
	if err != nil {
		s.ErrInternal(w, r, err)
	}
	usernames := make([]string, 0, len(users))
	for _, u := range users {
		usernames = append(usernames, u.Name)
	}
	WriteJSON(w, http.StatusOK, usernames)
}
