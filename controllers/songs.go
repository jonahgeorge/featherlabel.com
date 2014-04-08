package controllers

import "database/sql"
import "encoding/json"
import "fmt"
import "log"
import "net/http"
import "io/ioutil"
import "time"

import "github.com/gorilla/mux"
import "github.com/jonahgeorge/featherlabel.com/models"
import "github.com/mitchellh/goamz/s3"

type Song struct{}

// Retrieve list of songs without authenticated urls
func (s Song) Index(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		songs, err := models.Song{}.RetrieveAll(db)
		bytes, err := json.MarshalIndent(songs, "", "  ")
		if err != nil {
			log.Printf("%s", err)
		}

		fmt.Fprintf(w, "%s", bytes)
	}
}

// Create a song, add details into db and upload file to aws
func (s Song) Create(db *sql.DB, bucket *s3.Bucket) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// insert into db
		id, err := models.Song{}.Create(db, map[string]interface{}{
			"title":     r.FormValue("title"),
			"artist_id": r.FormValue("artist"),
		})

		if err != nil {
			log.Printf("%s", err)
		}

		file, _, err := r.FormFile("file")
		if err != nil {
			log.Printf("%s", err)
		}

		// upload to aws
		data, err := ioutil.ReadAll(file)
		key := fmt.Sprintf("songs/%s/%d.mp3", r.FormValue("artist"), id)
		fmt.Println(key)

		err = bucket.Put(key, data, "audio/mpeg", s3.ACL("authenticated-read"))
		if err != nil {
			log.Printf("%s", err)
		}

		// Send response

	}
}

// Retrieve a song, details from db and authenticate url with goamz
func (s Song) Retrieve(db *sql.DB, bucket *s3.Bucket) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		song, err := models.Song{}.RetrieveById(db, params["id"])

		expires := time.Now().Add(time.Duration(10) * time.Minute)
		key := fmt.Sprintf("songs/%d/%d.mp3", song.Artist.Id, song.Id)

		uri := bucket.SignedURL(key, expires)
		song.SignedUrl = uri

		bytes, err := json.MarshalIndent(song, "", "  ")
		if err != nil {
			log.Printf("%s", err)
		}

		fmt.Fprintf(w, "%s", bytes)
	}
}

// Update a song, details to db overwrite aws key if new file
func (s Song) Update(db *sql.DB, bucket *s3.Bucket) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Delete a song, remove details from db and delete key from aws
func (s Song) Delete(db *sql.DB, bucket *s3.Bucket) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
