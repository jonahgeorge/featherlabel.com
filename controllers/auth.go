package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/jonahgeorge/featherlabel.com/models"

	"code.google.com/p/go.crypto/bcrypt"
)

type Auth struct{}

func (a Auth) Signin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := t.ExecuteTemplate(w, "auth/signin", map[string]interface{}{
			"Title": "Sign In",
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// Process user login
func (a Auth) Authenticate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// retrieve form values
		email := r.FormValue("email")
		password := r.FormValue("password")

		// retrieve user by email
		user, err := models.User{}.RetrieveByEmail(email)
		if err != nil {
			fmt.Fprint(w, http.StatusNotAcceptable)
			return
		}

		// compare saved password and submitted password
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			fmt.Fprint(w, http.StatusUnauthorized)
			return
		}

		// init session
		session, _ := store.Get(r, "user")

		// set sesion vars
		session.Values["User"] = &user
		/*
			session.Values["Id"] = user.Id
			session.Values["Email"] = user.Email
			session.Values["Name"] = user.Name
		*/

		// save session
		session.Save(r, w)

		// return 202
		fmt.Fprint(w, http.StatusAccepted)
	}
}

func (a Auth) Signout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "user")
		session.Options.MaxAge = -1
		sessions.Save(r, w)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}

func (a Auth) Signup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := t.ExecuteTemplate(w, "auth/signup", map[string]interface{}{
			"Title": "Sign Up",
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (a Auth) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
