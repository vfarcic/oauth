package main

import (
	"net/http"
	"github.com/stretchr/objx"
	"github.com/stretchr/gomniauth/common"
	"strings"
	"encoding/json"
	"log"
)

func loginHandler(provider common.Provider) http.HandlerFunc {
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
		url := redirectURL + "?authid=" + user.Data().Get("id").Str()
		http.Redirect(w, r, url, http.StatusFound)
	}
}

func userApiHandler(dbHandler func(authID string) (MongoUser, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		segs := strings.Split(r.URL.Path, "/")
		authID := segs[4]
		users, err := dbHandler(authID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("222")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(users)
	}
}
