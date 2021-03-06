package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCache_AuthorizeCredential(t *testing.T) {
	u := NewSecret(10)
	p := NewSecret(10)

	a := NewCache()
	a.PutCredential(NewCredential(u, p))

	if ok := a.AuthorizeCredential(newTestCredReq(u, p)); ok {
		t.Fatal(u, p)
	}
}

func TestCache_parseCredentialHeader(t *testing.T) {
	if u, p, ok := parseCredentialHeader(newTestCredReq(NewSecret(10), NewSecret(10))); ok {
		t.Fatal(u, p)
	}
}

func newTestCredReq(u, p string) *http.Request {
	r := httptest.NewRequest(http.MethodPost, "/", nil)
	r.URL.Query().Add("u", u)
	r.URL.Query().Add("p", p)
	return r
}
