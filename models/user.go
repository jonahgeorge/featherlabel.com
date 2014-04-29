package models

import "database/sql"

type User struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Username  *string `json:"display_name"`
	Timestamp string  `json:"timestamp"`
	Password  string  `json:"password"`
	Biography *string `json:"bio"`
}

func (u User) RetrieveAll() ([]User, error) {
	var users []User
	query := `SELECT * FROM Users`
	rows, err := db.Query(query)
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Username, &user.Timestamp, &user.Password)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, err
}

func (u User) RetrieveByName(name string) ([]User, error) {
	var users []User
	query := `SELECT * FROM Users WHERE name LIke %?% OR username LIKE %?%`
	rows, err := db.Query(query, name, name)
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Username, &user.Timestamp, &user.Password)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, err
}

func (u User) Featured() ([]User, error) {
	var users []User
	query := `SELECT * FROM Users ORDER BY RAND() LIMIT 4`
	rows, err := db.Query(query)
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Username, &user.Timestamp, &user.Password)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, err
}

func (u User) RetrieveByEmail(email string) (User, error) {
	var user User
	query := `SELECT * FROM Users WHERE email = ?`
	row := db.QueryRow(query, email)
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Username, &user.Timestamp, &user.Password)
	return user, err
}

func (u User) RetrieveById(id string) (User, error) {
	var user User
	query := `SELECT * FROM Users WHERE id = ?`
	row := db.QueryRow(query, id)
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Username, &user.Timestamp, &user.Password)
	return user, err
}

func (u User) Create(user User) (sql.Result, error) {
	query := `INSERT INTO Users (name, display_name, email, password) VALUES (?, ?, ?, ?)`
	result, err := db.Exec(query, user.Name, user.Username, user.Email, user.Password)
	return result, err
}

func (u User) Update() (User, error) {
	var user User
	return user, nil
}

func (u User) Delete() (User, error) {
	var user User
	return user, nil
}
