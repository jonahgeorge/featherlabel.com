package main

import "database/sql"
import "log"
import "net/http"

import _ "github.com/Go-SQL-Driver/MySQL"
import "github.com/gorilla/mux"
import "github.com/jonahgeorge/featherlabel.com/controllers"
import "github.com/mitchellh/goamz/aws"
import "github.com/mitchellh/goamz/s3"
import "github.com/gosexy/yaml"
import "github.com/gosexy/to"

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
	accessKey := to.String(conf.Get("amazon", "access"))
	secretKey := to.String(conf.Get("amazon", "secret"))
	bucketKey := to.String(conf.Get("amazon", "bucket"))

	// open database connection
	db, err := sql.Open("mysql", user+":"+pass+"@/"+name)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.SetMaxIdleConns(100)

	// configure aws authentication
	auth, err := aws.GetAuth(accessKey, secretKey)
	if err != nil {
		log.Printf("%s", err)
	}

	// Try loading this from conf file?
	client := s3.New(auth, aws.USWest2)
	bucket := client.Bucket(bucketKey)

	// setup router
	r := mux.NewRouter()

	// register routes
	r.HandleFunc("/song", controllers.Song{}.Index(db)).Methods("GET")
	r.HandleFunc("/song", controllers.Song{}.Create(db, bucket)).Methods("POST")
	r.HandleFunc("/song/{id:[0-9]+}", controllers.Song{}.Retrieve(db, bucket)).Methods("GET")
	r.HandleFunc("/song/{id:[0-9]+}", controllers.Song{}.Update(db, bucket)).Methods("PUT")
	r.HandleFunc("/song/{id:[0-9]+}", controllers.Song{}.Delete(db, bucket)).Methods("DELETE")

	// pass router into http server
	http.Handle("/", r)

	// spin 'er up
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
