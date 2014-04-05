package models

import (
	"database/sql"
)

type User struct {
	Id       int64  `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
}
