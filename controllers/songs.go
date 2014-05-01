package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jonahgeorge/featherlabel.com/models"
	"github.com/mitchellh/goamz/s3"
)

type Song struct{}

// Retrieve list of songs without authenticated urls
func (s Song) Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// init session
		session, _ := store.Get(r, "user")

		// Retrieve all songs
		songs, err := models.Song{}.RetrieveAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// Render songs template
		err = t.ExecuteTemplate(w, "songs/index", map[string]interface{}{
			"Title":   "Songs",
			"Songs":   songs,
			"Session": session,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// Create a song, add details into db and upload file to aws
func (s Song) Create(bucket *s3.Bucket) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// init session
		session, _ := store.Get(r, "user")
		user := session.Values["User"].(*models.User)

		file, _, err := r.FormFile("file")
		if err != nil {
			log.Println(err)
			log.Println("Error happened on form parse")
			http.Redirect(w, r, "/songs", http.StatusInternalServerError)
			return
		}

		// insert into db
		id, err := models.Song{}.Create(map[string]interface{}{
			"title":     r.FormValue("title"),
			"artist_id": user.Id,
		})

		if err != nil {
			log.Println(err)
			log.Println("Error happened on db insert")
			http.Redirect(w, r, "/songs", http.StatusInternalServerError)
			return
		}

		// read file
		data, err := ioutil.ReadAll(file)
		if err != nil {
			log.Println(err)
			log.Println("Error happened on file read")
			http.Redirect(w, r, "/songs", http.StatusInternalServerError)
			return
		}

		// upload to aws
		key := fmt.Sprintf("songs/%d/%d.mp3", user.Id, id)
		err = bucket.Put(key, data, "audio/mpeg", s3.ACL("authenticated-read"))
		if err != nil {
			log.Println(err)
			log.Println("Error happened on file upload")
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		// Send response
		// Retrieve all songs
		songs, err := models.Song{}.RetrieveAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// Render songs template
		err = t.ExecuteTemplate(w, "songs/index", map[string]interface{}{
			"Title":   "Songs",
			"Songs":   songs,
			"Session": session,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// Retrieve a song, details from db and authenticate url with goamz
func (s Song) Retrieve(bucket *s3.Bucket) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// init session
		session, _ := store.Get(r, "user")

		// get url params
		params := mux.Vars(r)

		// retrieve song by id
		song, err := models.Song{}.RetrieveById(params["id"])

		expires := time.Now().Add(time.Duration(10) * time.Minute)
		key := fmt.Sprintf("songs/%d/%d.mp3", song.User.Id, song.Id)

		uri := bucket.SignedURL(key, expires)
		song.SignedUrl = uri

		// Render songs template
		err = t.ExecuteTemplate(w, "songs/show", map[string]interface{}{
			"Title":   song.Title,
			"Song":    song,
			"Session": session,
		})

		if err != nil {
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

func (s Song) Form() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// init session
		session, _ := store.Get(r, "user")

		// Render songs template
		err := t.ExecuteTemplate(w, "songs/form", map[string]interface{}{
			"Title":   "Upload",
			"Session": session,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
