package main

/*
sudo docker run -d --name mongo -p 27017:27017 mongo
go build -o oauth && ./oauth \
	-sec-key=bla \
	-google-client-id='472858977716-ej3ca5dtmq4krl7m085rpfno3cjp2ogp.apps.googleusercontent.com' \
	-google-secret='OnkptU4BTdE45mi-b3hACdAY' \
	-google-redirect-url='http://localhost:8080/auth/google/callback' \
	-redirect-url='http://www.wikipedia.org'
*/

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
	setGomniAuth(vars)
	// TODO: Add the rest of providers
	providerNames := []string{ "google" }
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
	log.Println("Starting the server", vars.host)
	if err := http.ListenAndServe(vars.host, nil); err != nil {
		log.Fatalln("Could not initiate the server", vars.host, " - ", err)
	}
}

func getGoogleProvider(vars Vars) common.Provider {
	return google.New(
		vars.googleProvider.clientId,
		vars.googleProvider.clientSecret,
		vars.googleProvider.redirectUrl,
	)
}

func setGomniAuth(vars Vars) {
	gomniauth.SetSecurityKey(vars.secKey)
	// TODO: Add the rest of providers
	// TODO: Manually test all providers
	gomniauth.WithProviders(
		getGoogleProvider(vars),
	)
}

// TODO: Add GET API
// TODO: Change Dockerfile FROM to Alpine Linux
// TODO: Add to Travis
// TODO: Add to Docker Hub
// TODO: Publish to hub.docker.com
// TODO: Promote
