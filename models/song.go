package models

import "log"

type SongModel struct {
	Id        *int
	Title     *string
	User      UserModel
	SignedUrl *string
}

// Delete the given song
func (s SongModel) Delete() error {
	return nil
}

// Retrieve all songs
func GetSongs() ([]*SongModel, error) {

	sql := `
		SELECT Songs.id, Songs.title, Users.id, Users.username
	    FROM Songs
	    LEFT JOIN Users ON Users.id = Songs.artist_id`

	rows, err := db.Query(sql)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var songs []*SongModel

	for rows.Next() {
		song := new(SongModel)
		err := rows.Scan(&song.Id, &song.Title, &song.User.Id, &song.User.Username)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		songs = append(songs, song)
	}

	return songs, err
}

// Retrieve all songs
func GetSongsTrending() ([]*SongModel, error) {

	sql := `
		SELECT Songs.id, Songs.title, Users.id, Users.username
	    FROM Songs
	    LEFT JOIN Users ON Users.id = Songs.artist_id
		LIMIT 20`

	rows, err := db.Query(sql)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var songs []*SongModel

	for rows.Next() {
		song := new(SongModel)
		err := rows.Scan(&song.Id, &song.Title, &song.User.Id, &song.User.Username)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		songs = append(songs, song)
	}

	return songs, err
}

// Retrieve a single song by its id (primary key)
func GetSongById(id string) (*SongModel, error) {

	sql := `
		SELECT Songs.id, Songs.title, Users.id, Users.username
		FROM Songs 
		LEFT JOIN Users ON Users.id = Songs.artist_id 
		WHERE Songs.id = ?`

	row := db.QueryRow(sql, id)
	song := new(SongModel)
	err := row.Scan(&song.Id, &song.Title, &song.User.Id, &song.User.Username)
	log.Println(err)
	return song, err
}

// Retrieve a slice of songs by fuzzy matching name
func GetSongsByName() ([]*SongModel, error) {
	var songs []*SongModel
	return songs, nil
}

// Retrieve a slice of songs by matching the tag
func GetSongsByGenre() ([]*SongModel, error) {
	var songs []*SongModel
	return songs, nil
}

// Insert a song record into db
func Create(data map[string]interface{}) (int64, error) {

	sql := `
		INSERT INTO Songs (title, artist_id) 
		VALUES (?, ?)`

	result, err := db.Exec(sql, data["title"], data["artist_id"])
	id, err := result.LastInsertId()
	return id, err
}
