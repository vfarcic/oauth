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
	"errors"
	"encoding/json"
)

var googleUrl = "http://google.com/auth"
var redirectUrl = "http://example.com/redirect"
var callbackUrl = "http://example.com/foo?code=4/klISKMqMfj2ErEykXTEI94kyhspflKAzbTih1eheJ4I.AlMYdYp5IQ4YEnp6UAPFm0GKMKMymwI"
var apiUserUrl = "http://example.com/api/v1/user/123456789"

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
	mockedProvider := getMockedCallbackProvider(nil, nil)

	req, _ := http.NewRequest("GET", callbackUrl, nil)
	w := httptest.NewRecorder()
	callbackHandler(mockedProvider, redirectUrl, SaveToMockedDB)(w, req)

	assert.Equal(t, redirectUrl + "?authID=", w.Header().Get("Location"))
}

func TestCallbackHandlerShouldReturnInternalServerErrorWhenCompleteAuthFails(t *testing.T) {
	mockedProvider := getMockedCallbackProvider(errors.New("This is an CompleteAuth error"), nil)

	req, _ := http.NewRequest("GET", callbackUrl, nil)
	w := httptest.NewRecorder()
	callbackHandler(mockedProvider, redirectUrl, SaveToMockedDB)(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestUserApiHandlerShouldReturnInternalServerErrorWhenDbReturnsAnError(t *testing.T) {
	req, _ := http.NewRequest("GET", apiUserUrl, nil)
	w := httptest.NewRecorder()
	userApiHandler(GetFromDBByAuthIDMockedWithError)(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestUserApiHandlerShouldReturnJson(t *testing.T) {
	w := doTestUserApiRequest()
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}

func TestUserApiHandlerShouldReturnUserFromDB(t *testing.T) {
	w := doTestUserApiRequest()
	json, _ := json.Marshal(testUser)
	assert.Equal(t, string(json) + "\n", w.Body.String())
}

func doTestUserApiRequest() *httptest.ResponseRecorder {
	req, _ := http.NewRequest("GET", apiUserUrl, nil)
	w := httptest.NewRecorder()
	userApiHandler(GetFromDBByAuthIDMocked)(w, req)
	return w
}

func TestCallbackHandlerShouldReturnInternalServerErrorWhenGetUserFails(t *testing.T) {
	mockedProvider := getMockedCallbackProvider(nil, errors.New("This is an GetUser error"))

	req, _ := http.NewRequest("GET", callbackUrl, nil)
	w := httptest.NewRecorder()
	callbackHandler(mockedProvider, redirectUrl, SaveToMockedDB)(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func getMockedCallbackProvider(completeAuthError error, getUserError error) common.Provider {
	creds := &common.Credentials{
		make(map[string]interface{}),
	}
	testProvider := new(test.TestProvider)
	testProvider.On("CompleteAuth", mock.Anything).Return(creds, completeAuthError)
	testProvider.On("GetUser", mock.Anything).Return(GetTestUser(), getUserError)
	return testProvider
}

func init() {
	log.SetOutput(ioutil.Discard)
	getProviders(TestVars)
}

func SaveToMockedDB(user MongoUser) error {
	return nil
}

func GetFromDBByAuthIDMocked(authID string) (MongoUser, error) {
	return testUser, nil
}

func GetFromDBByAuthIDMockedWithError(authID string) (MongoUser, error) {
	return testUser, errors.New("This is an error")
}
