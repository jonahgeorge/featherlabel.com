package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/Go-SQL-Driver/MySQL"

	"github.com/gorilla/mux"
	"github.com/gosexy/to"
	"github.com/gosexy/yaml"
	"github.com/jonahgeorge/featherlabel.com/controllers"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
)

func main() {
	// load settings from config file
	conf, err := yaml.Open("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// initialize s3 client
	bucket := InitializeAWS(conf)

	// initialize router
	router := InitializeRouter(bucket, conf)

	// register router
	http.Handle("/", router)

	// grab port from config and serve
	port := fmt.Sprintf(":%s", to.String(conf.Get("server", "port")))
	if err = http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func InitializeRouter(bucket *s3.Bucket, conf *yaml.Yaml) *mux.Router {
	secret := []byte(to.String(conf.Get("server", "secret")))

	// setup new router
	router := mux.NewRouter()

	// song routes
	router.HandleFunc("/songs", controllers.Song{}.Index()).Methods("GET")
	router.HandleFunc("/songs", controllers.Song{}.Create(bucket)).Methods("POST")
	router.HandleFunc("/songs/{id:[0-9]+}", controllers.Song{}.Retrieve(bucket)).Methods("GET")
	router.HandleFunc("/songs/{id:[0-9]+}", controllers.Song{}.Update(bucket)).Methods("PUT")
	router.HandleFunc("/songs/{id:[0-9]+}", controllers.Song{}.Delete(bucket)).Methods("DELETE")

	// user routes
	router.HandleFunc("/users", controllers.User{}.Index()).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", controllers.User{}.Retrieve()).Methods("GET")
	router.HandleFunc("/users", controllers.User{}.Create()).Methods("POST")
	router.HandleFunc("/authenticate", controllers.User{}.Authenticate()).Methods("POST")
	router.HandleFunc("/credentials", controllers.User{}.Credentials(secret)).Methods("POST")

	// pages
	router.HandleFunc("/about", controllers.Page{}.About()).Methods("GET")
	router.HandleFunc("/explore", controllers.Page{}.Explore()).Methods("GET")
	router.HandleFunc("/", controllers.Page{}.Index()).Methods("GET")

	// resource files
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	http.Handle("/vendor/", http.StripPrefix("/vendor/", http.FileServer(http.Dir("vendor"))))

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
