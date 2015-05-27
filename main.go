package main

// TODO: Test

import (
	"net/http"
	"log"
	"fmt"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
)

func main() {
	vars := GetVars(flagUtil)
	setGomniAuth(vars.secKey)
	providers := []string{"google", "github", "facebook"}
	for _, provider := range providers {
		// TODO: Change URI to param
		http.HandleFunc(fmt.Sprintf("/auth/%s/login", provider), loginHandler(provider))
		// TODO: Change URI to param
		// TODO: Change URL (google.com) to param
		http.HandleFunc(fmt.Sprintf("/auth/%s/callback", provider), callbackHandler(provider, "http://www.google.com"))
	}
	log.Println("Starting the server", vars.host)
	if err := http.ListenAndServe(vars.host, nil); err != nil {
		log.Fatalln("Could not initiate the server", vars.host, " - ", err)
	}
}

func setGomniAuth(secKey string) {
	gomniauth.SetSecurityKey(secKey)
	// TODO: Change to params
	// TODO: Add the rest of providers
	// TODO: Setup and test all providers
	gomniauth.WithProviders(
		google.New(
			"472858977716-ej3ca5dtmq4krl7m085rpfno3cjp2ogp.apps.googleusercontent.com",
			"OnkptU4BTdE45mi-b3hACdAY",
			"http://localhost:8080/auth/google/callback"),
		facebook.New(
			"",
			"",
			""),
		github.New(
			"472858977716-ej3ca5dtmq4krl7m085rpfno3cjp2ogp.apps.googleusercontent.com",
			"",
			""))
}

// TODO: Add API GET
// TODO: Create Dockerfile
// TODO: Publish to hub.docker.com
// TODO: Promote
