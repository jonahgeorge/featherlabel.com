package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"code.google.com/p/go.crypto/bcrypt"

	"github.com/dchest/authcookie"
	"github.com/gorilla/mux"
	"github.com/jonahgeorge/featherlabel.com/libraries"
	"github.com/jonahgeorge/featherlabel.com/models"
)

type User struct{}

func (u User) Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Retrieve all users
		users, err := models.User{}.RetrieveAll()
		if err != nil {
			// lib.QuickResponse(w, "failure", err.Error())
			return
		}

		// Render users/index template
		if err = t.ExecuteTemplate(w, "users/index", map[string]interface{}{
			"title": "Users",
			"songs": users,
			//"session": session,
		}); err != nil {
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

func (u User) Authenticate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Set response headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		var credentials struct {
			Email    *string `json:"email"`
			Password *string `json:"password"`
		}
		decoder := json.NewDecoder(r.Body)

		// Decode json payload
		err := decoder.Decode(&credentials)
		if err != nil {
			lib.QuickResponse(w, "failure", err.Error())
			return
		}

		// Retrieve user with email
		user, err := models.User{}.RetrieveByEmail(*credentials.Email)
		if err != nil {
			lib.QuickResponse(w, "failure", err.Error())
			return
		}

		// Check password here
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(*credentials.Password))
		if err != nil {
			// Password do not match
			lib.QuickResponse(w, "failure", err.Error())
			return
		}

		// How long token will be valid for
		// [todo] - Load time from config file
		time := 2 * time.Hour

		// [todo] - Load secret from config file
		secret := "keyboardcat"

		// Create token
		token := authcookie.NewSinceNow(*credentials.Email, time, []byte(secret))

		// Send token in response
		lib.QuickResponse(w, "success", token)
	}
}

func (u User) Retrieve() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		user, err := models.User{}.RetrieveById(params["id"])
		if err != nil {
			// send error
		}

		if err = t.ExecuteTemplate(w, "users/show", map[string]interface{}{
			"User": user,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (u User) Credentials(secret []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.FormValue("token") == "" {
			lib.QuickResponse(w, "failure", "Invalid token.")
			return
		}

		// contains email
		login := authcookie.Login(r.FormValue("token"), secret)
		log.Printf("%s\n", login)

		if login != "" {
			lib.QuickResponse(w, "success", "Access granted")
		} else {
			lib.QuickResponse(w, "failure", "Access denied")
		}
	}
}
