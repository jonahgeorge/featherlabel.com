package models

import "database/sql"

type User struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	DisplayName *string `json:"display_name"`
	Timestamp   string  `json:"timestamp"`
	Password    string  `json:"password"`
}

func (u User) RetrieveAll() ([]User, error) {

	var users []User

	rows, err := db.Query(`
    SELECT 
      id, name, email, display_name, timestamp, password 
    FROM 
      Users
    `)
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.DisplayName,
			&user.Timestamp, &user.Password)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, err
}

func (u User) RetrieveByName(name string) ([]User, error) {

	var users []User

	rows, err := db.Query(`
    SELECT 
      id, name, email, display_name, timestamp 
    FROM 
      Users 
    WHERE 
      name LIKE %?% 
    OR 
      display_name LIKE %?%
    `, name, name)
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.DisplayName,
			&user.Timestamp)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, err
}

func (u User) Featured() ([]User, error) {

	var users []User

	rows, err := db.Query(`
    SELECT 
      id, IF(display_name IS NOT NULL, display_name, name)
    FROM 
      Users 
    ORDER BY 
      RAND()
    LIMIT 
      4
    `)
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, err
}

func (u User) RetrieveByEmail(email string) (User, error) {
	var user User

	row := db.QueryRow("SELECT * FROM Users WHERE email = ?", email)
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.DisplayName, &user.Timestamp, &user.Password)
	if err != nil {
		return user, err
	}

	return user, err
}

func (u User) RetrieveById(id string) (User, error) {
	var user User

	query := `SELECT * FROM Users WHERE id = ?`

	row := db.QueryRow(query, id)
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.DisplayName, &user.Timestamp, &user.Password)
	return user, err
}

func (u User) Create(user User) (sql.Result, error) {

	result, err := db.Exec("INSERT INTO Users (name, display_name, email, password) VALUES (?, ?, ?, ?)", user.Name, user.DisplayName, user.Email, user.Password)
	if err != nil {
		return result, err
	}

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
