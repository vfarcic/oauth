package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/gomniauth/test"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/gomniauth/common"
	"errors"
)

var googleUrl = "http://google.com/auth"
var redirectUrl = "http://example.com/redirect"
var callbackUrl = "http://example.com/foo?code=4/klISKMqMfj2ErEykXTEI94kyhspflKAzbTih1eheJ4I.AlMYdYp5IQ4YEnp6UAPFm0GKMKMymwI"
var apiUserUrl = "http://example.com/api/v1/user/123456789"


// loginHandler

func TestLoginHandlerShouldRedirectToProviderUrl(t *testing.T) {
	testProvider := new(test.TestProvider)
	testProvider.On("GetBeginAuthURL", mock.Anything, mock.Anything).Return(googleUrl, nil)
	expected, _ := testProvider.GetBeginAuthURL(nil, nil)

	req, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	loginHandler(testProvider)(w, req)

	assert.Equal(t, expected, w.Header().Get("Location"))
}

// callbackHandler

func TestCallbackHandlerShouldRedirectToUrl(t *testing.T) {
	testDB := new(TestDB)
	testDB.On("Save", mock.Anything).Return(nil)
	mockedProvider := getMockedCallbackProvider(nil, nil)

	req, _ := http.NewRequest("GET", callbackUrl, nil)
	w := httptest.NewRecorder()
	callbackHandler(mockedProvider, redirectUrl, testDB)(w, req)

	assert.Equal(t, redirectUrl + "?authID=", w.Header().Get("Location"))
}

func TestCallbackHandlerShouldReturnInternalServerErrorWhenCompleteAuthFails(t *testing.T) {
	testDB := new(TestDB)
	mockedProvider := getMockedCallbackProvider(errors.New("This is an CompleteAuth error"), nil)

	req, _ := http.NewRequest("GET", callbackUrl, nil)
	w := httptest.NewRecorder()
	callbackHandler(mockedProvider, redirectUrl, testDB)(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestCallbackHandlerShouldReturnInternalServerErrorWhenGetUserFails(t *testing.T) {
	testDB := new(TestDB)
	mockedProvider := getMockedCallbackProvider(nil, errors.New("This is an GetUser error"))

	req, _ := http.NewRequest("GET", callbackUrl, nil)
	w := httptest.NewRecorder()
	callbackHandler(mockedProvider, redirectUrl, testDB)(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestCallbackHandlerShouldSetAuthNameCookie(t *testing.T) {
	testDB := new(TestDB)
	testDB.On("Save", mock.Anything).Return(nil)
	mockedProvider := getMockedCallbackProvider(nil, nil)
	expectedUser, _ := mockedProvider.GetUser(nil)

	req, _ := http.NewRequest("GET", callbackUrl, nil)
	w := httptest.NewRecorder()
	callbackHandler(mockedProvider, redirectUrl, testDB)(w, req)
	assert.Contains(t, w.Header().Get("Set-Cookie"), "authName=" + expectedUser.Name())
}

// TODO: Figure out how to test multiple cookies (works only with the first one authName).
// authAvatarURL and authID are missing tests

// logoutHandler

func TestLogoutHandlerShouldRemoveNameCookie(t *testing.T) {
	req, _ := http.NewRequest("GET", "/auth/logout", nil)
	w := httptest.NewRecorder()

	logoutHandler("http://example.com/logout/redirect")(w, req)
	assert.Contains(t, w.Header().Get("Set-Cookie"), "authName=;")
}

// TODO: Figure out how to test multiple cookies (works only with the first one authName).
// authAvatarURL and authID are missing tests

// userApi

func TestUserApiHandlerShouldReturnInternalServerErrorWhenDbReturnsAnError(t *testing.T) {
	testDB := new(TestDB)
	testDB.On("GetByAuthID", mock.Anything).Return(testUser, errors.New("SOME ERROR"))
	req, _ := http.NewRequest("GET", apiUserUrl, nil)
	w := httptest.NewRecorder()
	userApiHandler(testDB)(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestUserApiHandlerShouldReturnJson(t *testing.T) {
	w := doTestUserApiRequest()
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}

//func TestUserApiHandlerShouldReturnUserFromDB(t *testing.T) {
//	w := doTestUserApiRequest()
//	json, _ := json.Marshal(testUser)
//	assert.Equal(t, string(json) + "\n", w.Body.String())
//}
//

// Helper

func getMockedCallbackProvider(completeAuthError error, getUserError error) common.Provider {
	creds := &common.Credentials{
		make(map[string]interface{}),
	}
	testProvider := new(test.TestProvider)
	testProvider.On("CompleteAuth", mock.Anything).Return(creds, completeAuthError)
	testProvider.On("GetUser", mock.Anything).Return(GetTestUser(), getUserError)
	return testProvider
}

// TODO: Remove
//func getMockedBowerComponents(completeAuthError error, getUserError error) common.Provider {
//	creds := &common.Credentials{
//		make(map[string]interface{}),
//	}
//	testProvider := new(test.TestProvider)
//	testProvider.On("CompleteAuth", mock.Anything).Return(creds, completeAuthError)
//	testProvider.On("GetUser", mock.Anything).Return(GetTestUser(), getUserError)
//	return testProvider
//}

func doTestUserApiRequest() *httptest.ResponseRecorder {
	testDB := new(TestDB)
	testDB.On("GetByAuthID", mock.Anything).Return(testUser, nil)
	req, _ := http.NewRequest("GET", apiUserUrl, nil)
	w := httptest.NewRecorder()
	userApiHandler(testDB)(w, req)
	return w
}