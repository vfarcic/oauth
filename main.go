package main

/*
go build -o oauth && ./oauth \
	-sec-key=bla \
	-google-client-id='472858977716-ej3ca5dtmq4krl7m085rpfno3cjp2ogp.apps.googleusercontent.com' \
	-google-secret='OnkptU4BTdE45mi-b3hACdAY' \
	-google-redirect-url='http://localhost:8080/auth/google/callback'
*/

// TODO: Test

import (
	"net/http"
	"log"
	"fmt"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
)

func main() {
	vars := GetVars(flagUtil)
	setGomniAuth(vars)
	providerNames := []string{ "google" }
	for _, providerName := range providerNames {
		provider, err := gomniauth.Provider(providerName)
		if err != nil {
			panic(err)
		}
		// TODO: Change URI to param
		http.HandleFunc(fmt.Sprintf("/auth/%s/login", providerName), loginHandler(provider))
		// TODO: Change URI to param
		// TODO: Change URL (google.com) to param
		http.HandleFunc(fmt.Sprintf("/auth/%s/callback", providerName), callbackHandler(provider, "http://www.google.com"))
	}
	log.Println("Starting the server", vars.host)
	if err := http.ListenAndServe(vars.host, nil); err != nil {
		log.Fatalln("Could not initiate the server", vars.host, " - ", err)
	}
}

func setGomniAuth(vars Vars) {
	gomniauth.SetSecurityKey(vars.secKey)
	// TODO: Change to params
	// TODO: Add the rest of providers
	// TODO: Setup and test all providers
	gomniauth.WithProviders(
		google.New(
			vars.googleProvider.clientId,
			vars.googleProvider.clientSecret,
			vars.googleProvider.redirectUrl,
		),
	)
}

// TODO: Add API GET
// TODO: Create Dockerfile
// TODO: Publish to hub.docker.com
// TODO: Promote
