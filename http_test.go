package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/gomniauth"
	"log"
	"io/ioutil"
)

func TestLoginHandlerShouldRedirectToProviderUrl(t *testing.T) {
	provider, _ := gomniauth.Provider("google")
	expected, _ := provider.GetBeginAuthURL(nil, nil)

	req, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	loginHandler("google")(w, req)

	assert.Equal(t, expected, w.Header().Get("Location"))
}

func TestLoginHandlerShouldPanicWhenUnknownProvider(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()

	assert.Panics(t, func() { loginHandler("unknown")(w, req) })
}

func TestCallbackHandlerShouldRedirectToUrl(t *testing.T) {
	expected := "http://example.com/redirect"
	req, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	callbackHandler("google", expected)(w, req)

	assert.Equal(t, expected, w.Header().Get("Location"))
}

func init() {
	log.SetOutput(ioutil.Discard)
	setGomniAuth("MY_SECRET_KEY")
}
