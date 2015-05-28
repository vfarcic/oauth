package main

import (
	"net/http"
	"github.com/stretchr/objx"
	"fmt"
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
		_, err := provider.CompleteAuth(query)
		fmt.Println(query)
		fmt.Println(err)
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
