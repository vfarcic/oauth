package main

import (
	"testing"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/testify/assert"
)

var testClVars = CLVars{
	domain: "MY_DOMAIN",
	port: "1234",
	secKey: "MY_SECURITY_KEY",
	googleProvider: provider {
		clientId: "MY_GOOGLE_CLIENT_ID",
		clientSecret: "MY_GOOGLE_CLIENT_SECRET",
		redirectUrl: "MY_REDIRECT_URL_AFTER_GOOGLE",
	},
}

// getGoogleProvider

func TestGetGoogleProviderShouldSetClientId(t *testing.T) {
	provider := getGoogleProvider(testClVars)

	assert.Equal(t, "google", provider.Name())
}

// getProviders

func TestGetProvidersShouldSetSecurityKey(t *testing.T) {
	getProviders(testClVars)

	assert.Equal(t, testClVars.secKey, gomniauth.GetSecurityKey())
}

func TestGetProvidersShouldHaveGoogleProvider(t *testing.T) {
	getProviders(testClVars)

	assert.Contains(t, gomniauth.SharedProviderList.Providers(), getGoogleProvider(testClVars))
}

// main

//func TestMainShould