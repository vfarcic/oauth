package main

import (
	"net/http"
	"github.com/stretchr/objx"
	"github.com/stretchr/gomniauth/common"
)

func loginHandler(provider common.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url, _ := provider.GetBeginAuthURL(nil, nil)
		http.Redirect(w, r, url, http.StatusFound)
	}
}

func callbackHandler(provider common.Provider, redirectURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query, _ := objx.FromURLQuery(r.URL.RawQuery)
		creds, err := provider.CompleteAuth(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user, err := provider.GetUser(creds)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// TODO: Test
		SaveToDB(getMongoUser(user))
		http.Redirect(w, r, redirectURL, http.StatusFound)
	}
}
