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
	"fmt"
	"net/http"
	"log"
	phttp "github.com/pikanezi/http"
)

func main() {
	vars := GetVars(flagUtil)
	providerNames := getProviders(vars)
	addr := vars.addr
	r := phttp.NewRouter()
	r.SetCustomHeader(phttp.Header{
		"Access-Control-Allow-Origin": "*",
	})
	for _, providerName := range providerNames {
		provider, err := gomniauth.Provider(providerName)
		if err != nil {
			panic(err)
		}
		// TODO: Change URI to param
		r.HandleFunc(fmt.Sprintf("/auth/%s/login", providerName), loginHandler(provider))
		// TODO: Change URI to param
		r.HandleFunc(
			fmt.Sprintf("/auth/%s/callback", providerName),
			callbackHandler(provider, vars.redirectUrl, MongoDB{}))
	}
	r.HandleFunc("/auth/api/v1/user/{id}", userApiHandler(MongoDB{}))
	r.HandleFunc("/auth/logout", logoutHandler(vars.redirectUrl))
	r.PathPrefix("/components/").Handler(
		http.StripPrefix("/components/", http.FileServer(http.Dir("components"))))
	r.PathPrefix("/component_tests/").Handler(
		http.StripPrefix("/component_tests/", http.FileServer(http.Dir("component_tests"))))
	log.Println("Starting the server on", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalln("Could not initiate the server", addr, " - ", err)
	}
}

func getGoogleProvider(vars Vars) common.Provider {
	return google.New(
		vars.googleProvider.clientId,
		vars.googleProvider.clientSecret,
		vars.googleProvider.redirectUrl,
	)
}

// TODO: Add the rest of providers
func getProviders(vars Vars) []string {
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
