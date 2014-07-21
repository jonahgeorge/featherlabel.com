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

type SongController struct{}

// Retrieve list of songs without authenticated urls
func (s SongController) Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// retrieve user session
		session, err := store.Get(r, "user")

		// retrieve all songs
		songs, err := models.SongFactory{}.GetSongs()

		// render songs index template
		err = t.ExecuteTemplate(w, "songs/index", map[string]interface{}{
			"Title":   "Songs",
			"Songs":   songs,
			"Session": session,
		})

		// catch template rendering errors
		if err != nil {
			log.Println(err)
		}
	}
}

// Create a song, add details into db and upload file to aws
func (s SongController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// retrieve user session
		session, err := store.Get(r, "user")
		user := session.Values["User"].(*models.UserModel)

		// check if user agreed to terms and copyright
		if r.FormValue("copyright") != "on" || r.FormValue("terms") != "on" {
			// redirect to upload page
			http.Redirect(w, r, "/upload", http.StatusFound)
			return
		}

		// get file from form
		file, _, err := r.FormFile("file")
		if err != nil {
			log.Println(err)

			// redirect to upload page
			http.Redirect(w, r, "/upload", http.StatusFound)
			return
		}

		// read file into memory
		data, err := ioutil.ReadAll(file)
		if err != nil {
			log.Println(err)

			// redirect to upload page
			http.Redirect(w, r, "/upload", http.StatusFound)
			return
		}

		// insert song data into database
		id, err := models.SongFactory{}.Create(map[string]interface{}{
			"title":     r.FormValue("title"),
			"artist_id": user.Id,
		})

		if err != nil {
			log.Println(err)

			// redirect to upload page
			http.Redirect(w, r, "/upload", http.StatusFound)
			return
		}

		// upload to Amazon S3
		err = bucket.Put(
			"songs/"+string(id)+".mp3", data,
			"audio/mpeg", s3.ACL("authenticated-read"))

		if err != nil {
			// [todo] - Remove song from database here
			log.Println(err)
			http.Redirect(w, r, "/upload", http.StatusFound)
			return
		}

		// redirect to /songs
		http.Redirect(w, r, "/songs/"+string(id), http.StatusFound)
	}
}

// Retrieve a song, details from db and authenticate url with goamz
func (s SongController) Retrieve() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// retrieve user session
		session, _ := store.Get(r, "user")

		// get url parameters
		params := mux.Vars(r)

		// retrieve song by id
		song, err := models.SongFactory{}.GetSongById(params["id"])

		// song retrieval error
		if err != nil {
			log.Println(err)

			session.AddFlash("There was an error retrieving that song.")
			session.Save(r, w)

			// redirect to songs index page
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		uri := bucket.SignedURL(
			fmt.Sprintf("songs/%d/%d.mp3", song.User.Id, song.Id),
			time.Now().Add(time.Duration(10)*time.Minute))

		song.SignedUrl = &uri

		// render songs template
		err = t.ExecuteTemplate(w, "songs/show", map[string]interface{}{
			"Title":   song.Title,
			"Song":    song,
			"Session": session,
		})

		// catch template rendering errors
		if err != nil {
			log.Println(err)
		}

	}
}

// Update a song, details to db overwrite aws key if new file
func (s SongController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Delete a song, remove details from db and delete key from aws
func (s SongController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s SongController) Form() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// retrieve user session
		session, _ := store.Get(r, "user")

		// rendering song upload form
		err := t.ExecuteTemplate(w, "songs/form", map[string]interface{}{
			"Title":   "Upload",
			"Session": session,
		})

		// catch template rendering errors
		if err != nil {
			log.Println(err)
		}
	}
}
