package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/jonahgeorge/featherlabel.com/models"
	"github.com/jonahgeorge/featherlabel.com/libraries"

	"code.google.com/p/go.crypto/bcrypt"
)

type AuthenticationController struct{}

// Render user log in form
func (a AuthenticationController) RenderLogInForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		// retrieve user session
		session, _ := store.Get(r, "user")
		flashes := session.Flashes()
		session.Save(r, w)

		// render template
		err := t.ExecuteTemplate(w, "auth/login", map[string]interface{}{
			"Title":   "Log In",
			"Session": session,
			"Flashes": flashes,
		})
		
		// catch template rendering errors
		if err != nil {
			log.Println(err)
		}
	}
}

// Process user login
func (a AuthenticationController) Authenticate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// retrieve user session
		session, _ := store.Get(r, "user")

		// retrieve user by email
		user, err := models.UserFactory{}.GetUserByEmail(r.FormValue("email"))
		
		// if user does not exist
		if err != nil {
			session.AddFlash("Invalid email or password")
			session.Save(r, w)
			
			// redirect to signin
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		// compare saved password and submitted password
		err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(r.FormValue("password")))
		
		// if passwords do not match
		if err != nil {
			session.AddFlash("Invalid email or password")
			session.Save(r, w)
			
			// redirect to signin
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		// login successfull, set sesion vars
		session.Values["User"] = &user
		session.AddFlash("Welcome back, " + *user.Username)
		session.Save(r, w)

		// redirect to homepage
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

// End the current user session
func (a AuthenticationController) LogOut() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		// retrieve user session
		session, _ := store.Get(r, "user")
		
		// kill session
		session.Options.MaxAge = -1
		sessions.Save(r, w)

		// redirect to home page
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func (a AuthenticationController) RenderSignUpForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		// retrieve user session
		session, _ := store.Get(r, "user")
		flashes := session.Flashes()
		session.Save(r, w)
		
		// render sign up page
		err := t.ExecuteTemplate(w, "auth/signup", map[string]interface{}{
			"Title":   "Sign Up",
			"Session": session,
			"Flashes": flashes,
		})

		// catch template rendering errors
		if err != nil {
			log.Println(err)
		}
	}
}

// Create new user record
func (a AuthenticationController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// retrieve user session
		session, _ := store.Get(r, "user")

		// check if user exists
		user, err := models.UserFactory{}.GetUserByEmail(r.FormValue("email"))
		
		// if user already exists
		if user.Id != nil {
			session.AddFlash("This email address is already in use")
			session.Save(r, w)
			
			// redirect to sign up page
			http.Redirect(w, r, "/signup", http.StatusFound)
			return
		}

		// create hashed password
		hashed_password, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 9)
		if err != nil {
			log.Println(err)
			
			session.AddFlash("An error occured during the sign up process")
			session.Save(r, w)
			
			// redirect to sign up page
			http.Redirect(w, r, "/signup", http.StatusFound)
			return
		}

		// insert user into database
		_, err = models.UserFactory{}.Create(map[string]interface{}{
			"username": r.FormValue("username"),
			"email":    r.FormValue("email"),
			"password": hashed_password,
		})
		
		if err != nil {
			log.Println(err)
			
			session.AddFlash("An error occured during the sign up process")
			session.Save(r, w)
			
			// redirect to sign up page
			http.Redirect(w, r, "/signup", http.StatusFound)
			return
		}
		
		// [todo] - Send confirmation email here
		mailer.SendConfirmationEmail(r.FormValue("email"))

		// redirect to home page
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
