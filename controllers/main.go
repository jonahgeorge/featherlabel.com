package controllers

import (
	"log"
	"net/http"

	"github.com/jonahgeorge/featherlabel.com/models"
)

type MainController struct{}

func (m MainController) Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// retrieve user session
		session, err := store.Get(r, "user")

		songs, err := models.SongFactory{}.GetSongsTrending()
		users := models.UserFactory{}.GetUsersFeatured()

		// catch retrieval errors
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// render home page template
		err = t.ExecuteTemplate(w, "index", map[string]interface{}{
			"Songs":   songs,
			"Users":   users,
			"Session": session,
		})

		// catch template rendering errors
		if err != nil {
			log.Println(err)
		}
	}
}

func (m MainController) Explore() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// retrieve user session
		session, err := store.Get(r, "user")

		songs, err := models.SongFactory{}.GetSongs()

		// catch retrieval errors
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = t.ExecuteTemplate(w, "index", map[string]interface{}{
			"Title":   "Explore",
			"Songs":   songs,
			"Session": session,
		})

		// catch template rendering errors
		if err != nil {
			log.Println(err)
		}
	}
}

func (m MainController) About() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// retrieve user session
		session, _ := store.Get(r, "user")

		// render about page template
		err := t.ExecuteTemplate(w, "about", map[string]interface{}{
			"Title":   "About",
			"Session": session,
		})

		// catch template rendering errors
		if err != nil {
			log.Println(err)
		}
	}
}
