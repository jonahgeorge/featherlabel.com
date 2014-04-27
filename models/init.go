package models

import (
	"database/sql"
	"log"

	_ "github.com/Go-SQL-Driver/MySQL"

	"github.com/gosexy/to"
	"github.com/gosexy/yaml"
)

var db *sql.DB

func init() {
	// Load config file
	conf, err := yaml.Open("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve credentials from config file
	username := to.String(conf.Get("database", "username"))
	password := to.String(conf.Get("database", "password"))
	name := to.String(conf.Get("database", "name"))

	// Open mysql connection
	db, err = sql.Open("mysql", username+":"+password+"@/"+name)
	if err != nil {
		log.Fatal(err)
	}

	// Sets the maximum number of connections in the idle connection pool
	db.SetMaxIdleConns(100)
}
