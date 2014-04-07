package models

import (
	"database/sql"
	"log"
)

type Song struct {
	Id     int64  `db:"id"      json:"id"`
	Title  string `db:"title"   json:"title"`
	Key    string `db:"aws_key" json:"key,omitempty"`
	Artist string `db:"artist"  json:"artist"`
}

// Retrieve all songs
func (s Song) RetrieveAll(db *sql.DB) ([]Song, error) {

	rows, err := db.Query("SELECT id, title, artist FROM Songs")
	if err != nil {
		log.Printf("%s", err)
	}

	var songs []Song

	for rows.Next() {
		var song Song

		err := rows.Scan(&song.Id, &song.Title, &song.Artist)
		if err != nil {
			log.Printf("%s", err)
		}

		songs = append(songs, song)
	}

	return songs, err
}

// Retrieve a single song by its id (primary key)
func (s Song) RetrieveById(db *sql.DB, id string) (Song, error) {

	var song Song

	row := db.QueryRow("SELECT id, title, aws_key, artist FROM Songs WHERE id = ?", id)
	err := row.Scan(&song.Id, &song.Title, &song.Key, &song.Artist)
	if err != nil {
		log.Printf("%s", err)
	}

	return song, err
}

// // Retrieve a slice of songs by fuzzy matching name
// func (s Song) RetrieveByName() ([]Song, error) {
// 	return nil, nil
// }

// // Retrieve a slice of songs by matching the tag
// func (s Song) RetrieveByTag() ([]Song, error) {
// 	return nil, nil
// }

// // Insert a song record into db
// func (s Song) Create() (Song, error) {
// 	return nil, nil
// }
