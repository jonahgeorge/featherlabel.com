package models

import (
	"database/sql"
	"log"
)

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

// Retrieve a slice of all users
func GetUsers() []UserModel {

	sql := `
		SELECT id, username, email, password	
		FROM Users`

	rows, err := db.Query(sql)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var users []UserModel

	for rows.Next() {
		var user UserModel
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
		if err != nil {
			log.Println(err)
		}
		users = append(users, user)
	}

	return users
}

// Fuzzy match usernames against the provided string
func GetUsersByName(name string) []UserModel {

	sql := `
		SELECT id, username, email, password	
		FROM Users 
		WHERE username LIKE %?%`

	rows, err := db.Query(sql, name)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var users []UserModel

	for rows.Next() {
		var user UserModel
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
		if err != nil {
			log.Println(err)
		}
		users = append(users, user)
	}

	return users
}

// Placeholder function to retrieve "featured" users
func GetUsersFeatured() []UserModel {

	sql := `
		SELECT id, username, email
		FROM Users 
		ORDER BY RAND() 
		LIMIT 4`

	rows, err := db.Query(sql)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var users []UserModel

	for rows.Next() {
		var user UserModel
		err := rows.Scan(&user.Id, &user.Username, &user.Email)
		if err != nil {
			log.Println(err)
		}
		users = append(users, user)
	}

	return users
}

func GetUserByEmail(email string) (*UserModel, error) {

	sql := `
		SELECT id, username, email, password, timestamp
		FROM Users 
		WHERE email = ?`

	row := db.QueryRow(sql, email)

	user := new(UserModel)

	err := row.Scan(&user.Id, &user.Username,
		&user.Email, &user.Password, &user.Timestamp)

	return user, err
}

func GetUserById(id string) UserModel {

	sql := `
		SELECT id, username, email, password, timestamp 
		FROM Users 
		WHERE id = ?`

	row := db.QueryRow(sql, id)

	var user UserModel

	err := row.Scan(
		&user.Id, &user.Username, &user.Email,
		&user.Password, &user.Timestamp)

	if err != nil {
		log.Println(err)
	}

	return user
}

func Create(params map[string]interface{}) (sql.Result, error) {

	sql := `
		INSERT INTO Users (username, email, password) 
		VALUES (?, ?, ?)`

	return db.Exec(sql, params["username"], params["email"], params["password"])
}
