package models

import "log"

type Song struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	User      User   `json:"user"`
	SignedUrl string `json:"url,omitempty"`
}

// Retrieve all songs
func (s Song) RetrieveAll() ([]Song, error) {

	var songs []Song

	rows, err := db.Query(`
    SELECT 
      Songs.id, Songs.title, Users.id, 
      IF(Users.display_name, Users.display_name, Users.name)
    FROM 
      Songs 
    LEFT JOIN 
      Users ON Users.id = Songs.artist_id
    `)

	if err != nil {
		return songs, err
	}

	for rows.Next() {
		var song Song
		err := rows.Scan(&song.Id, &song.Title, &song.User.Id, &song.User.Name)
		if err != nil {
			return songs, err
		}
		songs = append(songs, song)
	}

	return songs, err
}

func (u User) GetSongs() ([]Song, error) {

	var songs []Song

	query := `SELECT 
        Songs.id, Songs.title, Users.id, 
        IF(Users.display_name, Users.display_name, Users.name)
      FROM 
        Songs 
      LEFT JOIN 
        Users ON Users.id = Songs.artist_id
      WHERE
        Songs.artist_id = ?`

	rows, err := db.Query(query, u.Id)
	if err != nil {
		return songs, err
	}

	for rows.Next() {
		var song Song
		err := rows.Scan(&song.Id, &song.Title, &song.User.Id, &song.User.Name)
		if err != nil {
			return songs, err
		}
		songs = append(songs, song)
	}

	return songs, err
}

// Retrieve a single song by its id (primary key)
func (s Song) RetrieveById(id string) (Song, error) {

	var song Song

	row := db.QueryRow(`
      SELECT 
        Songs.id, Songs.title, Users.id, 
        IF(Users.display_name, Users.display_name, Users.name)
      FROM 
        Songs 
      LEFT JOIN 
        Users ON Users.id = Songs.artist_id 
      WHERE 
        Songs.id = ?
    `, id)

	err := row.Scan(&song.Id, &song.Title, &song.User.Id, &song.User.Name)
	if err != nil {
		return song, err
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
func (s Song) Create(data map[string]interface{}) (int64, error) {

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
