package main

import (
	"net/http"
	"github.com/stretchr/gomniauth"
)

func loginHandler(providerName string) http.HandlerFunc {
	provider, err := gomniauth.Provider(providerName)
	if err != nil {
		panic(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		url, _ := provider.GetBeginAuthURL(nil, nil)
		http.Redirect(w, r, url, http.StatusFound)
	}
}

func callbackHandler(providerName string, redirectURL string) http.HandlerFunc {
//	provider, err := gomniauth.Provider(providerName)
//	if err != nil {
//		panic(err)
//	}
	return func(w http.ResponseWriter, r *http.Request) {
//		omap, err := objx.FromURLQuery(r.URL.RawQuery)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		creds, err := provider.CompleteAuth(omap)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		user, err := provider.GetUser(creds)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		// TODO: Store user to MongoDB
//		log.Print(user)
		http.Redirect(w, r, redirectURL, http.StatusFound)
	}
}
