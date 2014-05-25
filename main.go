package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/Go-SQL-Driver/MySQL"
	. "github.com/jonahgeorge/featherlabel.com/controllers"

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

	// initialize s3 client
	bucket := InitializeAWS(conf)

	// initialize router
	router := InitializeRouter(bucket, conf)

	// register router
	http.Handle("/", router)

	// grab port from config and serve
	port := fmt.Sprintf(":%s", to.String(conf.Get("server", "port")))
	log.Printf("<== Listening on port %s\n", port)

	if err = http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func InitializeRouter(bucket *s3.Bucket, conf *yaml.Yaml) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/songs", SongController{}.Index()).Methods("GET")
	router.HandleFunc("/songs", SongController{}.Create(bucket)).Methods("POST")
	router.HandleFunc("/songs/{id:[0-9]+}", SongController{}.Retrieve(bucket)).Methods("GET")
	router.HandleFunc("/songs/{id:[0-9]+}/update", SongController{}.Update(bucket)).Methods("POST")
	router.HandleFunc("/songs/{id:[0-9]+}/delete", SongController{}.Delete(bucket)).Methods("POST")
	router.HandleFunc("/upload", SongController{}.Form()).Methods("GET")

	router.HandleFunc("/users", UserController{}.Index()).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", UserController{}.Retrieve()).Methods("GET")
	router.HandleFunc("/users", UserController{}.Create()).Methods("POST")

	router.HandleFunc("/signin", AuthenticationController{}.RenderSignInForm()).Methods("GET")
	router.HandleFunc("/signin", AuthenticationController{}.Authenticate()).Methods("POST")
	router.HandleFunc("/signout", AuthenticationController{}.Signout()).Methods("GET")
	router.HandleFunc("/signup", AuthenticationController{}.RenderSignUpForm()).Methods("GET")
	router.HandleFunc("/signup", AuthenticationController{}.Create()).Methods("POST")

	router.HandleFunc("/about", MainController{}.About()).Methods("GET")
	router.HandleFunc("/explore", MainController{}.Explore()).Methods("GET")
	router.HandleFunc("/", MainController{}.Index()).Methods("GET")

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	http.Handle("/vendor/", http.StripPrefix("/vendor/", http.FileServer(http.Dir("vendor"))))

	log.Printf("<== Router initialized\n")
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

	log.Printf("<== AWS Configured\n")
	return client.Bucket(bucket)
}
