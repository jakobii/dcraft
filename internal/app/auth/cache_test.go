package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCache_IsSuper(t *testing.T) {
	pw := NewSecret(10)
	a := NewCache(WithSuper(pw))

	r := httptest.NewRequest(http.MethodPost, "/foo?super="+pw, nil)

	if !a.IsSuper(r) {
		t.Fatal(a.super)
	}
}
func TestCache_Authorize(t *testing.T) {
	pw := NewSecret(10)
	a := NewCache(WithSuper(pw))

	r := httptest.NewRequest(http.MethodPost, "/foo?super="+pw, nil)

	if !a.Authorize(r) {
		t.Fatal(a.super)
	}
}

func TestCache_Authenticate(t *testing.T) {
	pw := NewSecret(10)
	a := NewCache(WithSuper(pw))

	r := httptest.NewRequest(http.MethodPost, "/foo?super="+pw, nil)

	if _, ok := a.Authenticate(r); !ok {
		t.Fatal(a.super)
	}
}
