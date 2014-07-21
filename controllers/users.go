package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jonahgeorge/featherlabel.com/models"
)

type UserController struct{}

// Render a list of all users. To be used as a user search/browse page
func (u UserController) Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// retrieve user session
		session, _ := store.Get(r, "user")

		// retrieve all users
		users := models.UserFactory{}.GetUsers()

		// render users/index template
		err := t.ExecuteTemplate(w, "users/index", map[string]interface{}{
			"title":   "Users",
			"Users":   users,
			"Session": session,
		})
		
		// catch template rendering errors
		if err != nil {
			log.Println(err)
		}
	}
}

// Render an individual user page
func (u UserController) Retrieve() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// retrieve user session
		session, _ := store.Get(r, "user")

		// get url parameters
		params := mux.Vars(r)

		user := models.UserFactory{}.GetUserById(params["id"])
		songs := user.GetSongs()

		// render template
		err := t.ExecuteTemplate(w, "users/show", map[string]interface{}{
			"User":    user,
			"Songs":   songs,
			"Session": session,
		})

		// catch template rendering errors
		if err != nil {
			log.Println(err)
		}
	}
}
