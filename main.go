package main

import (
	"net/http"
	"log"
	"flag"
	"fmt"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/objx"
	"io"
)

func main() {
	var host = flag.String("host", ":8080", "The application host.")
	flag.Parse()
	// TODO: Change to param
	gomniauth.SetSecurityKey("THIS_IS_A_TOP_SECRET_KEY")
	// TODO: Change to params
	// TODO: Add the rest of providers
	gomniauth.WithProviders(
		google.New(
			"472858977716-ej3ca5dtmq4krl7m085rpfno3cjp2ogp.apps.googleusercontent.com",
			"OnkptU4BTdE45mi-b3hACdAY",
			"http://localhost:8080/auth/google/callback"),
		facebook.New(
			"472858977716-ej3ca5dtmq4krl7m085rpfno3cjp2ogp.apps.googleusercontent.com",
			"OnkptU4BTdE45mi-b3hACdAY",
			"http://localhost:8080/auth/facebook/callback"),
		github.New(
			"472858977716-ej3ca5dtmq4krl7m085rpfno3cjp2ogp.apps.googleusercontent.com",
			"OnkptU4BTdE45mi-b3hACdAY",
			"http://localhost:8080/auth/github/callback"))
	providers := []string{"google", "github", "facebook"}
	for _, provider := range providers {
		// TODO: Change URI to param
		http.HandleFunc(fmt.Sprintf("/auth/%s/login", provider), loginHandler(provider))
		http.HandleFunc(fmt.Sprintf("/auth/%s/callback", provider), callbackHandler(provider))
	}
	log.Println("Starting the server", *host)
	if err := http.ListenAndServe(*host, nil); err != nil {
		log.Fatalln("Could not initiate the server", host, " - ", err)
	}
}

func loginHandler(providerName string) http.HandlerFunc {
	provider, err := gomniauth.Provider(providerName)
	if err != nil {
		log.Fatalln("Could not get the provider", providerName)
		panic(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		url, err := provider.GetBeginAuthURL(nil, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, url, http.StatusFound)
	}
}

func callbackHandler(providerName string) http.HandlerFunc {
	provider, err := gomniauth.Provider(providerName)
	if err != nil {
		log.Fatalln("Could not get the provider", providerName)
		panic(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		omap, err := objx.FromURLQuery(r.URL.RawQuery)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		creds, err := provider.CompleteAuth(omap)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user, err := provider.GetUser(creds)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// TODO: Store user to MongoDB
		// TODO: Remove
		data := fmt.Sprintf("%#v", user)
		// TODO: Remove
		io.WriteString(w, data)
		// TODO: Redirect to the final screen
	}
}

// TODO: Add API GET
