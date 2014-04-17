package models

import (
	"database/sql"
	"log"
)

type Song struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	Artist    Artist `json:"artist"`
	SignedUrl string `json:"url,omitempty"`
}

type Artist struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// Retrieve all songs
func (s Song) RetrieveAll(db *sql.DB) ([]Song, error) {

	rows, err := db.Query("SELECT Songs.`id`, Songs.`title`, Artists.`id`, Artists.`name` FROM Songs LEFT JOIN Artists ON Artists.`id` = Songs.`artist_id`")

	if err != nil {
		log.Printf("%s", err)
	}

	var songs []Song

	for rows.Next() {
		var song Song
		err := rows.Scan(&song.Id, &song.Title, &song.Artist.Id, &song.Artist.Name)
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

	row := db.QueryRow("SELECT Songs.`id`, Songs.`title`, Artists.`id`, Artists.`name` FROM Songs LEFT JOIN Artists ON Artists.`id` = Songs.`artist_id` WHERE Songs.`id` = ?", id)

	err := row.Scan(&song.Id, &song.Title, &song.Artist.Id, &song.Artist.Name)
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

// Insert a song record into db
func (s Song) Create(db *sql.DB, data map[string]interface{}) (int64, error) {

	result, err := db.Exec("INSERT INTO Songs (title, artist_id) VALUES (?, ?)", data["title"], data["artist_id"])
	if err != nil {
		log.Printf("%s", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("%s", err)
	}

	return id, err
}
