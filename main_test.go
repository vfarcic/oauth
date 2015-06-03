package main

import (
	"testing"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/testify/assert"
)

func TestGetGoogleProviderShouldSetClientId(t *testing.T) {
	provider := getGoogleProvider(TestVars)

	assert.Equal(t, "google", provider.Name())
}

func TestSetGomniAuthShouldSetSecurityKey(t *testing.T) {
	getProviders(TestVars)

	assert.Equal(t, TestVars.secKey, gomniauth.GetSecurityKey())
}

func TestSetGomniAuthShouldHaveGoogleProvider(t *testing.T) {
	getProviders(TestVars)

	assert.Contains(t, gomniauth.SharedProviderList.Providers(), getGoogleProvider(TestVars))
}
