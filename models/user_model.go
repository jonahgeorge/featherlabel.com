package models

import "log"

type UserModel struct {
	Id        *int
	Username  *string
	Email     *string
	Password  *string
	Biography *string
	Timestamp *string
}

// Retrieve all songs for the given user
func (u UserModel) GetSongs() []SongModel {

	sql := `
	SELECT 
        Songs.id, Songs.title, Users.id, Users.username
	FROM 
		Songs 
	LEFT JOIN 
		Users ON Users.id = Songs.artist_id
	WHERE
		Songs.artist_id = ?`

	rows, err := db.Query(sql, *u.Id)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var songs []SongModel

	for rows.Next() {
		var song SongModel
		err := rows.Scan(&song.Id, &song.Title, &song.User.Id, &song.User.Username)
		if err != nil {
			log.Println(err)
		}
		songs = append(songs, song)
	}

	return songs
}

func (u UserModel) Update() error {
	return nil
}

func (u UserModel) Delete() error {
	return nil
}
