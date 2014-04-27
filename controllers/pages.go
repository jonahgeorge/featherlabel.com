package controllers

import (
	"net/http"

	"github.com/jonahgeorge/featherlabel.com/models"
)

type Page struct{}

func (p Page) Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Retrieve all songs
		songs, err := models.Song{}.RetrieveAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if err := t.ExecuteTemplate(w, "index", map[string]interface{}{
			"Songs": songs,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}
	}
}

func (p Page) Explore() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve all songs
		songs, err := models.Song{}.RetrieveAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if err := t.ExecuteTemplate(w, "index", map[string]interface{}{
			"Title": "Explore",
			"Songs": songs,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}
}

func (p Page) About() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := t.ExecuteTemplate(w, "about", map[string]interface{}{
			"Title": "About",
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
