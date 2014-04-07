package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jonahgeorge/featherlabel.com/models"
	"github.com/mitchellh/goamz/s3"
	"log"
	"net/http"
)

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
		// 		file, header, err := r.FormFile("file")
		// 		if err != nil {
		// 			log.Printf("%s", err)
		// 		}

		// 		fmt.Printf("%s\n%+v\n", header.Filename, file)

		// 		b, err := ioutil.ReadAll(file)
		// 		if err != nil {
		// 			log.Printf("%s", err)
		// 		}

		// 		ioutil.WriteFile(header.Filename, b, 0x777)
	}
}

// Retrieve a song, details from db and authenticate url with goamz
func (s Song) Retrieve(db *sql.DB, bucket *s3.Bucket) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		song, err := models.Song{}.RetrieveById(db, bucket, params["id"])
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
