package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/jonahgeorge/featherlabel.com/models"
)

type UserController struct{}

func (u UserController) Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// init session
		session, _ := store.Get(r, "user")

		// Retrieve all users
		users := UserFactory{}.GetUsers()

		// Render users/index template
		err := t.ExecuteTemplate(w, "users/index", map[string]interface{}{
			"title":   "Users",
			"songs":   users,
			"Session": session,
		})

		if err != nil {
			log.Println(err)
		}
	}
}

func (u UserController) Retrieve() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// init session
		session, _ := store.Get(r, "user")

		// get url params
		params := mux.Vars(r)

		user := UserFactory{}.GetUserById(params["id"])
		songs := user.GetSongs()

		// render template
		err := t.ExecuteTemplate(w, "users/show", map[string]interface{}{
			"User":    user,
			"Songs":   songs,
			"Session": session,
		})

		if err != nil {
			log.Println(err)
		}
	}
}
