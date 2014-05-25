package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	. "github.com/jonahgeorge/featherlabel.com/models"

	"code.google.com/p/go.crypto/bcrypt"
)

type AuthenticationController struct{}

// [GET /signin] - Render user sign in form
func (a AuthenticationController) RenderSignInForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// retrieve session
		session, err := store.Get(r, "user")

		// render template
		err = t.ExecuteTemplate(w, "auth/signin", map[string]interface{}{
			"Title":   "Sign In",
			"Session": session,
		})

		// catch template rendering errors
		if err != nil {
			log.Println(err)
		}
	}
}

// process user login
func (a AuthenticationController) Authenticate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// retrieve session
		session, _ := store.Get(r, "user")

		// retrieve user by email
		user := UserFactory{}.GetUserByEmail(r.FormValue("email"))

		// if no user is found
		if user.Email == nil {
			session.Values["Error"] = "No user with that email was found."
			session.Save(r, w)
			http.Redirect(w, r, "/signin", http.StatusFound)
			return
		}

		// compare saved password and submitted password
		err := bcrypt.CompareHashAndPassword(
			[]byte(*user.Password), []byte(r.FormValue("password")))

		// if passwords do not match
		if err != nil {
			session.Values["Error"] = "Invalid password."
			session.Save(r, w)
			http.Redirect(w, r, "/signin", http.StatusFound)
			return
		}

		// set sesion vars
		session.Values["User"] = &user
		session.Values["Success"] = "Welcome back, " + *user.Username
		session.Save(r, w)

		// redirect to homepage
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func (a AuthenticationController) Signout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "user")
		// Kill session
		session.Options.MaxAge = -1
		sessions.Save(r, w)

		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func (a AuthenticationController) RenderSignUpForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "user")

		t.ExecuteTemplate(w, "auth/signup",
			map[string]interface{}{
				"Title":   "Sign Up",
				"Session": session,
			})
	}
}

func (a AuthenticationController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// init session
		session, _ := store.Get(r, "user")

		// [todo] - Check if user exists
		user := UserFactory{}.GetUserByEmail(r.FormValue("email"))
		if user.Email != nil {

			// Set error message
			session.Values["Flash"] = "A user with that email already exists, please use another."
			session.Save(r, w)

			http.Redirect(w, r, "/signup", http.StatusFound)
			return
		}

		// Create hashed password
		hashed_password, err := bcrypt.GenerateFromPassword(
			[]byte(r.FormValue("password")), 9)

		if err != nil {
			log.Println(err)
		}

		// Insert user into database
		_, err = UserFactory{}.Create(map[string]interface{}{
			"username": r.FormValue("username"),
			"email":    r.FormValue("email"),
			"password": hashed_password,
		})

		if err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/", http.StatusFound)
	}
}
