package controllers

import (
	"encoding/json"
	"net/http"

	"code.google.com/p/go.crypto/bcrypt"
	"github.com/gorilla/mux"
	"github.com/jonahgeorge/featherlabel.com/libraries"
	"github.com/jonahgeorge/featherlabel.com/models"
)

type User struct{}

func (u User) Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// init session
		session, _ := store.Get(r, "user")

		// Retrieve all users
		users, err := models.User{}.RetrieveAll()
		if err != nil {
			// lib.QuickResponse(w, "failure", err.Error())
			return
		}

		// Render users/index template
		err = t.ExecuteTemplate(w, "users/index", map[string]interface{}{
			"title":   "Users",
			"songs":   users,
			"Session": session,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}
}

func (u User) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// set response headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		var creds models.User
		decoder := json.NewDecoder(r.Body)

		// Decode json payload
		err := decoder.Decode(&creds)
		if err != nil {
			lib.QuickResponse(w, "failure", err.Error())
			return
		}

		// [todo] - Check if user exists

		// Create hashed password
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 9)
		if err != nil {
			lib.QuickResponse(w, "failure", err.Error())
			return
		}

		// Construct user struct with valid password
		creds.Password = string(hashedPass)

		// Insert user into database
		_, err = models.User{}.Create(creds)
		if err != nil {
			lib.QuickResponse(w, "failure", err.Error())
			return
		}

		// Send response
		// [todo] - Send token?
		lib.QuickResponse(w, "success", "You have been registered")
	}
}

func (u User) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := t.ExecuteTemplate(w, "auth/login", map[string]interface{}{
			"Title": "Login",
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (u User) Retrieve() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// init session
		session, _ := store.Get(r, "user")

		// get url params
		params := mux.Vars(r)

		// retrieve user
		user, err := models.User{}.RetrieveById(params["id"])

		// retrieve songs
		songs, err := user.GetSongs()

		// render template
		err = t.ExecuteTemplate(w, "users/show", map[string]interface{}{
			"User":    user,
			"Songs":   songs,
			"Session": session,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
