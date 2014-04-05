package models

import (
	"database/sql"
)

type Song struct {
	Id    int64  `db:"id"`
	Title string `db:"title"`
	Key   string `db:"aws_key"`
}

// Retrieve a single song by its id (primary key)
func (s Song) RetrieveById() (Song, error) {
	return nil, nil
}

// Retrieve a slice of songs by fuzzy matching name
func (s Song) RetrieveByName() ([]Song, error) {
	return nil, nil
}

// Retrieve a slice of songs by matching the tag
func (s Song) RetrieveByTag() ([]Song, error) {
	return nil, nil
}

// Insert a song record into db
func (s Song) Create() (Song, error) {
	return nil, nil
}
