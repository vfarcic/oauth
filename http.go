package main

import (
	"net/http"
	"github.com/stretchr/objx"
	"github.com/stretchr/gomniauth/common"
	"encoding/json"
	"log"
)

func loginHandler(provider common.Provider) http.HandlerFunc {
	log.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	return func(w http.ResponseWriter, r *http.Request) {
		url, _ := provider.GetBeginAuthURL(nil, nil)
		http.Redirect(w, r, url, http.StatusFound)
	}
}

func callbackHandler(provider common.Provider, redirectURL string, dbHandler func(user MongoUser) error) http.HandlerFunc {
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
		dbHandler(getMongoUser(user))
		url := redirectURL + "?authID=" + user.Data().Get("id").Str()
		http.Redirect(w, r, url, http.StatusFound)
	}
}

func userApiHandler(dbHandler func(authID string) (MongoUser, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authID := r.URL.Query().Get(":id")
		users, err := dbHandler(authID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

//func allowOrigin() http.Handler {
//	return func(w http.ResponseWriter, r *http.Request) {
//
//	}
//}