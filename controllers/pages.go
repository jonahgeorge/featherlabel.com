package controllers

import (
	"net/http"

	"github.com/jonahgeorge/featherlabel.com/models"
)

type Page struct{}

func (p Page) Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// init session
		session, _ := store.Get(r, "user")

		// Retrieve featured songs
		songs, err := models.Song{}.RetrieveAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// Retrieve featured users
		users, err := models.User{}.Featured()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if err := t.ExecuteTemplate(w, "index", map[string]interface{}{
			"Songs":   songs,
			"Users":   users,
			"Session": session,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}
	}
}

func (p Page) Explore() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// init session
		session, _ := store.Get(r, "user")

		// Retrieve all songs
		songs, err := models.Song{}.RetrieveAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		err = t.ExecuteTemplate(w, "index", map[string]interface{}{
			"Title":   "Explore",
			"Songs":   songs,
			"Session": session,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}
}

func (p Page) About() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// init session
		session, _ := store.Get(r, "user")

		err := t.ExecuteTemplate(w, "about", map[string]interface{}{
			"Title":   "About",
			"Session": session,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
