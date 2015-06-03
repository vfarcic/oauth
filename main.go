package main

// TODO: Test

import (
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/gomniauth/common"
	"fmt"
	"net/http"
	"log"
)

func main() {
	vars := GetVars(flagUtil)
	providerNames := getProviders(vars)
	for _, providerName := range providerNames {
		provider, err := gomniauth.Provider(providerName)
		if err != nil {
			panic(err)
		}
		// TODO: Change URI to param
		http.HandleFunc(fmt.Sprintf("/auth/%s/login", providerName), loginHandler(provider))
		// TODO: Change URI to param
		http.HandleFunc(
			fmt.Sprintf("/auth/%s/callback", providerName),
			callbackHandler(provider, vars.redirectUrl, SaveToMongoDB))
	}
	// TODO: Test
	http.HandleFunc("/api/v1/user/", userApiHandler(GetFromDBByAuthID))
	if err := http.ListenAndServe(vars.host, nil); err != nil {
		log.Fatalln("Could not initiate the server", vars.host, " - ", err)
	}
	log.Println("Started the server on", vars.host)
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
