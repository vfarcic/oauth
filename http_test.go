package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"log"
	"io/ioutil"
	"github.com/stretchr/gomniauth/test"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/gomniauth/common"
)

var googleUrl = "http://google.com/auth"
var redirectUrl = "http://example.com/redirect"
var callbackUrl = "http://example.com/foo?code=4/klISKMqMfj2ErEykXTEI94kyhspflKAzbTih1eheJ4I.AlMYdYp5IQ4YEnp6UAPFm0GKMKMymwI"

var testVars = Vars {
	host: "MY_DOMAIN",
	secKey: "MY_SECURITY_KEY",
	googleProvider: provider {
		clientId: "MY_GOOGLE_CLIENT_ID",
		clientSecret: "MY_GOOGLE_CLIENT_SECRET",
		redirectUrl: "MY_REDIRECT_URL_AFTER_GOOGLE",
	},
}

func TestLoginHandlerShouldRedirectToProviderUrl(t *testing.T) {
	testProvider := new(test.TestProvider)
	testProvider.On("GetBeginAuthURL", mock.Anything, mock.Anything).Return(googleUrl, nil)
	expected, _ := testProvider.GetBeginAuthURL(nil, nil)

	req, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	loginHandler(testProvider)(w, req)

	assert.Equal(t, expected, w.Header().Get("Location"))
}

func TestCallbackHandlerShouldRedirectToUrl(t *testing.T) {
	creds := &common.Credentials{
		make(map[string]interface{}),
	}
	testProvider := new(test.TestProvider)
	testProvider.On("CompleteAuth", mock.Anything).Return(creds, nil)

	req, _ := http.NewRequest("GET", callbackUrl, nil)
	w := httptest.NewRecorder()
	callbackHandler(testProvider, redirectUrl)(w, req)

	assert.Equal(t, redirectUrl, w.Header().Get("Location"))
}

//func TestCallbackHandlerShouldReturnInternalServerErrorWhenCompleteAuthIsUnsuccessful(t *testing.T) {
//	req, _ := http.NewRequest("GET", callbackUrl, nil)
//	req.URL.RawQuery = "INCORRECT_QUERY"
//	httptest.NewRecorder()
////	w := httptest.NewRecorder()
////	assert.Panics(t, func() { callbackHandler("unknown", redirectUrl)(w, req) })
//}

func init() {
	log.SetOutput(ioutil.Discard)
	setGomniAuth(testVars)
}
