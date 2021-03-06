package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jakobii/dcraft/internal/app/auth"
)

func TestServer_handleToken(t *testing.T) {
	pw := auth.NewSecret(10)
	s := NewServer(
		WithAuthenticator(auth.NewCache(auth.WithSuper(pw))),
	)

	router := mux.NewRouter()
	s.routeToken(router)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, pathToken+"?super="+pw, nil)
	router.ServeHTTP(w, r)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatal(resp.StatusCode)
	}
}
