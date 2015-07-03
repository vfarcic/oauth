package main

import (
	"net/http"
	"github.com/stretchr/objx"
	"github.com/stretchr/gomniauth/common"
	"encoding/json"
	"log"
	phttp "github.com/pikanezi/http"
	"github.com/stretchr/gomniauth"
	"fmt"
)

func loginHandler(provider common.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url, _ := provider.GetBeginAuthURL(nil, nil)
		http.Redirect(w, r, url, http.StatusFound)
	}
}

func callbackHandler(provider common.Provider, redirectURL string, db DB) http.HandlerFunc {
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
		authID := user.Data().Get("id").Str()
		db.Save(GetMongoUser(user))
		http.SetCookie(w, &http.Cookie{
			Name: "authName",
			Value: user.Name(),
			Path: "/"})
		http.SetCookie(w, &http.Cookie{
			Name: "authAvatarURL",
			Value: user.AvatarURL(),
			Path: "/"})
		http.SetCookie(w, &http.Cookie{
			Name: "authID",
			Value: authID,
			Path: "/"})
		url := redirectURL + "?authID=" + authID
		http.Redirect(w, r, url, http.StatusFound)
	}
}

func userApiHandler(db DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authID := r.URL.Query().Get(":id")
		users, err := db.GetByAuthID(authID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

func logoutHandler(redirectURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		removeCookie(w, "authName")
		removeCookie(w, "authAvatarURL")
		removeCookie(w, "authID")
		w.Header()["Location"] = []string{redirectURL}
		w.WriteHeader(http.StatusTemporaryRedirect)
	}
}

func removeCookie(w http.ResponseWriter, cookieName string) {
	http.SetCookie(w, &http.Cookie{
		Name: cookieName,
		Value: "",
		Path: "/",
		MaxAge: -1,
	})
}

func StartServer(providerNames []string, redirectUrl string, addr string) {
	r := phttp.NewRouter()
	r.SetCustomHeader(phttp.Header{
		"Access-Control-Allow-Origin": "*",
	})
	for _, providerName := range providerNames {
		provider, err := gomniauth.Provider(providerName)
		if err != nil {
			panic(err)
		}
		// TODO: Change URI to param
		r.HandleFunc(fmt.Sprintf("/auth/%s/login", providerName), loginHandler(provider))
		// TODO: Change URI to param
		r.HandleFunc(
			fmt.Sprintf("/auth/%s/callback", providerName),
			callbackHandler(provider, redirectUrl, MongoDB{}))
	}
	r.HandleFunc("/auth/api/v1/user/{id}", userApiHandler(MongoDB{}))
	r.HandleFunc("/auth/logout", logoutHandler(redirectUrl))
	r.PathPrefix("/components/").Handler(
		http.StripPrefix("/components/", http.FileServer(http.Dir("components"))))
	r.PathPrefix("/bower_components/").Handler(
		http.StripPrefix("/bower_components/", http.FileServer(http.Dir("bower_components"))))
	r.PathPrefix("/component_tests/").Handler(
		http.StripPrefix("/component_tests/", http.FileServer(http.Dir("component_tests"))))
	log.Println("Starting the server on", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalln("Could not initiate the server", addr, " - ", err)
	}
}