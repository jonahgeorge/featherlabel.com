package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"code.google.com/p/go.crypto/bcrypt"

	"github.com/dchest/authcookie"
	"github.com/jonahgeorge/featherlabel.com/models"
	"github.com/mecop/mecop-api/libraries"
)

type User struct{}

func (u User) Index(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// set response headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		users, err := models.User{}.RetrieveAll(db)
		if err != nil {
			lib.QuickResponse(w, "failure", err.Error())
			return
		}

		bytes, err := json.Marshal(users)
		if err != nil {
			lib.QuickResponse(w, "failure", err.Error())
			return
		}

		fmt.Fprintf(w, "%s\n", bytes)
	}
}

func (u User) Create(db *sql.DB) http.HandlerFunc {
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
		_, err = models.User{}.Create(db, creds)
		if err != nil {
			lib.QuickResponse(w, "failure", err.Error())
			return
		}

		// Send response
		// [todo] - Send token?
		lib.QuickResponse(w, "success", "You have been registered")
	}
}

func (u User) Authenticate(db *sql.DB) http.HandlerFunc {
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
		user, err := models.User{}.RetrieveByEmail(db, *credentials.Email)
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

func (u User) Credentials(db *sql.DB, secret []byte) http.HandlerFunc {
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
