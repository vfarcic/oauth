package main

// TODO: Test Web Components UI
// TODO: Hash (MD5) cookie values
// TODO: Add base tag to components
// TODO: Add gulp with tests
// TODO: Create Dockerfile for all tests

// TODO: Test

import (
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/gomniauth/common"
)

var clVars = CLVars{}

func main() {
	vars := clVars.GetCLVariables(flagUtil)
	providerNames := getProviders(vars)
	StartServer(providerNames, vars.redirectUrl, vars.addr)
}

func getGoogleProvider(vars CLVars) common.Provider {
	return google.New(
		vars.googleProvider.clientId,
		vars.googleProvider.clientSecret,
		vars.googleProvider.redirectUrl,
	)
}

// TODO: Add the rest of providers
func getProviders(vars CLVars) []string {
	gomniauth.SetSecurityKey(vars.secKey)
	// TODO: Add the rest of providers
	// TODO: Manually test all providers
	gomniauth.WithProviders(
		getGoogleProvider(vars),
	)
	return []string{ "google" }
}

// TODO: Change Dockerfile FROM to Alpine Linux
// TODO: Add to Travis
// TODO: Add to Docker Hub
// TODO: Publish to hub.docker.com
// TODO: Promote
