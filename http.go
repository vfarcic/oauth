package main

import (
	"net/http"
	"github.com/stretchr/objx"
	"github.com/stretchr/gomniauth/common"
	"encoding/json"
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