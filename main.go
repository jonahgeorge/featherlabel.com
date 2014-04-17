package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/feather-label/api/controllers"
	"github.com/gorilla/mux"
	"github.com/gosexy/to"
	"github.com/gosexy/yaml"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
)

func main() {

	// load settings from config file
	conf, err := yaml.Open("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// initalize database connection
	db := InitializeDb(conf)
	defer db.Close()

	// initialize s3 client
	bucket := InitializeAWS(conf)

	// initialize router
	router := InitializeRouter(db, bucket, conf)

	// pass router into http server
	http.Handle("/", router)

	// grab port from conf file
	port := fmt.Sprintf(":%s", to.String(conf.Get("server", "port")))

	// spin 'er up
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func InitializeDb(conf *yaml.Yaml) *sql.DB {

	// retrieve credentials from config file
	username := to.String(conf.Get("database", "username"))
	password := to.String(conf.Get("database", "password"))
	name := to.String(conf.Get("database", "name"))

	// start up mysql connection
	db, err := sql.Open("mysql", username+":"+password+"@/"+name)
	if err != nil {
		log.Fatal(err)
	}

	// set maximum number of connections in the idle connection pool
	db.SetMaxIdleConns(100)

	return db
}

func InitializeRouter(db *sql.DB, bucket *s3.Bucket, conf *yaml.Yaml) *mux.Router {

	secret := []byte(to.String(conf.Get("server", "secret")))

	// setup new router
	router := mux.NewRouter()

	// song routes
	router.HandleFunc("/songs", controllers.Song{}.Index(db)).Methods("GET")
	router.HandleFunc("/songs", controllers.Song{}.Create(db, bucket)).Methods("POST")
	router.HandleFunc("/songs/{id:[0-9]+}", controllers.Song{}.Retrieve(db, bucket)).Methods("GET")
	router.HandleFunc("/songs/{id:[0-9]+}", controllers.Song{}.Update(db, bucket)).Methods("PUT")
	router.HandleFunc("/songs/{id:[0-9]+}", controllers.Song{}.Delete(db, bucket)).Methods("DELETE")

	// user routes
	router.HandleFunc("/users", controllers.User{}.Index(db)).Methods("GET")
	router.HandleFunc("/users", controllers.User{}.Create(db)).Methods("POST")
	router.HandleFunc("/authenticate", controllers.User{}.Authenticate(db)).Methods("POST")
	router.HandleFunc("/credentials", controllers.User{}.Credentials(db, secret)).Methods("POST")

	return router
}

func InitializeAWS(conf *yaml.Yaml) *s3.Bucket {

	// load variables from config
	access := to.String(conf.Get("amazon", "access"))
	secret := to.String(conf.Get("amazon", "secret"))
	bucket := to.String(conf.Get("amazon", "bucket"))

	// configure aws authentication
	auth, err := aws.GetAuth(access, secret)
	if err != nil {
		log.Fatal(err)
	}

	// [todo] - load region from config file
	client := s3.New(auth, aws.USWest2)

	return client.Bucket(bucket)
}
