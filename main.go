package main

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/gorilla/mux"
	"github.com/gosexy/to"
	"github.com/gosexy/yaml"
	"github.com/jonahgeorge/featherlabel.com/controllers"
	"log"
	"net/http"
)

func main() {

	// load settings from config file
	conf, err := yaml.Open("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	user := to.String(conf.Get("database", "user"))
	pass := to.String(conf.Get("database", "pass"))
	name := to.String(conf.Get("database", "name"))
	// secret := to.String(conf.Get("session", "secret"))
	port := to.String(conf.Get("port"))

	// open database connection
	db, err := sql.Open("mysql", user+":"+pass+"@/"+name)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.SetMaxIdleConns(100)

	// setup router
	r := mux.NewRouter()

	// register routes
	r.HandleFunc("/song", controllers.Song{}.Index(db)).Methods("GET")
	r.HandleFunc("/song", controllers.Song{}.Create(db)).Methods("POST")
	r.HandleFunc("/song/{id:[0-9]+}", controllers.Song{}.Retrieve(db)).Methods("GET")
	r.HandleFunc("/song/{id:[0-9]+}", controllers.Song{}.Update(db)).Methods("PUT")
	r.HandleFunc("/song/{id:[0-9]+}", controllers.Song{}.Delete(db)).Methods("DELETE")

	// pass router into http server
	http.Handle("/", r)

	// spin 'er up
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
