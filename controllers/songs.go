package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jonahgeorge/featherlabel.com/libraries"
	"github.com/jonahgeorge/featherlabel.com/models"
	"github.com/mitchellh/goamz/s3"
)

type Song struct{}

// Retrieve list of songs without authenticated urls
func (s Song) Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Retrieve all songs
		songs, err := models.Song{}.RetrieveAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// Render songs template
		if err = t.ExecuteTemplate(w, "songs/index", map[string]interface{}{
			"title": "Songs",
			"songs": songs,
			//"session": session,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// Create a song, add details into db and upload file to aws
func (s Song) Create(bucket *s3.Bucket) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// insert into db
		id, err := models.Song{}.Create(map[string]interface{}{
			"title":     r.FormValue("title"),
			"artist_id": r.FormValue("artist"),
		})

		if err != nil {
			lib.QuickResponse(w, "failure", err.Error())
			return
		}

		file, _, err := r.FormFile("file")
		if err != nil {
			lib.QuickResponse(w, "failure", err.Error())
			return
		}

		// upload to aws
		data, err := ioutil.ReadAll(file)
		key := fmt.Sprintf("songs/%s/%d.mp3", r.FormValue("artist"), id)
		fmt.Println(key)

		err = bucket.Put(key, data, "audio/mpeg", s3.ACL("authenticated-read"))
		if err != nil {
			lib.QuickResponse(w, "failure", err.Error())
			return
		}

		// Send response

	}
}

// Retrieve a song, details from db and authenticate url with goamz
func (s Song) Retrieve(bucket *s3.Bucket) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// get url params
		params := mux.Vars(r)

		// retrieve song by id
		song, err := models.Song{}.RetrieveById(params["id"])
		if err != nil {
			// lib.QuickResponse(w, "failure", err.Error())
			return
		}

		expires := time.Now().Add(time.Duration(10) * time.Minute)
		key := fmt.Sprintf("songs/%d/%d.mp3", song.User.Id, song.Id)

		uri := bucket.SignedURL(key, expires)
		song.SignedUrl = uri

		//bytes, err := json.MarshalIndent(song, "", "  ")
		//if err != nil {
		//	lib.QuickResponse(w, "failure", err.Error())
		//	return
		//}
		//
		//fmt.Fprintf(w, "%s", bytes)

		// Render songs template
		if err = t.ExecuteTemplate(w, "songs/show", map[string]interface{}{
			"Title": song.Title,
			"Song":  song,
			//"session": session,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}
}

// Update a song, details to db overwrite aws key if new file
func (s Song) Update(bucket *s3.Bucket) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Delete a song, remove details from db and delete key from aws
func (s Song) Delete(bucket *s3.Bucket) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
