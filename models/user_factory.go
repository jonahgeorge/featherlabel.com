package models

import (
	"database/sql"
	"log"
)

type UserFactory struct {
}

// Retrieve a slice of all users
func (u UserFactory) GetUsers() []UserModel {

	sql := `
	SELECT 
		id, username, email, password	
	FROM 
		Users`

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
func (u UserFactory) GetUsersByName(name string) []UserModel {

	sql := `
	SELECT 
		id, username, email, password	
	FROM 
		Users 
	WHERE 
		username LIKE %?%`

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
func (u UserFactory) GetUsersFeatured() []UserModel {

	sql := `
	SELECT 
		id, username, email
	FROM 
		Users 
	ORDER BY 
		RAND() 
	LIMIT 
		4`

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

func (u UserFactory) GetUserByEmail(email string) UserModel {

	sql := `
	SELECT 
		id, username, email, password, timestamp
	FROM 
		Users 
	WHERE 
		email = ?`

	row := db.QueryRow(sql, email)

	var user UserModel

	err := row.Scan(
		&user.Id, &user.Username, &user.Email,
		&user.Password, &user.Timestamp)

	if err != nil {
		log.Println(err)
	}

	return user
}

func (u UserFactory) GetUserById(id string) UserModel {

	sql := `
	SELECT 
		id, username, email, password, timestamp 
	FROM 
		Users 
	WHERE 
		id = ?`

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

func (u UserFactory) Create(params map[string]interface{}) (sql.Result, error) {
	sql := `INSERT INTO Users (username, email, password) VALUES (?, ?, ?)`
	return db.Exec(sql, params["username"], params["email"], params["password"])
}
