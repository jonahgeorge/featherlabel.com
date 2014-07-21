package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/jonahgeorge/featherlabel.com/controllers"

	"github.com/gorilla/mux"
	"github.com/gosexy/to"
	"github.com/gosexy/yaml"
)

func main() {

	// load settings from config file
	conf, err := yaml.Open("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// initialize router
	router := InitializeRouter()

	// register router
	http.Handle("/", router)

	// get port number from config file
	port := to.String(conf.Get("server", "port"))

	fmt.Printf("Serving application from port %s\n", port)

	// serve application
	if err = http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func InitializeRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/songs", controllers.SongController{}.Index()).Methods("GET")
	router.HandleFunc("/songs", controllers.SongController{}.Create()).Methods("POST")
	router.HandleFunc("/songs/{id:[0-9]+}", controllers.SongController{}.Retrieve()).Methods("GET")
	router.HandleFunc("/songs/{id:[0-9]+}/update", controllers.SongController{}.Update()).Methods("POST")
	router.HandleFunc("/songs/{id:[0-9]+}/delete", controllers.SongController{}.Delete()).Methods("POST")
	router.HandleFunc("/upload", controllers.SongController{}.Form()).Methods("GET")

	router.HandleFunc("/users", controllers.UserController{}.Index()).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", controllers.UserController{}.Retrieve()).Methods("GET")
	// router.HandleFunc("/users", UserController{}.Create()).Methods("POST")

	router.HandleFunc("/login", controllers.AuthenticationController{}.RenderLogInForm()).Methods("GET")
	router.HandleFunc("/login", controllers.AuthenticationController{}.Authenticate()).Methods("POST")
	router.HandleFunc("/logout", controllers.AuthenticationController{}.LogOut()).Methods("GET")
	router.HandleFunc("/signup", controllers.AuthenticationController{}.RenderSignUpForm()).Methods("GET")
	router.HandleFunc("/signup", controllers.AuthenticationController{}.Create()).Methods("POST")

	router.HandleFunc("/about", controllers.MainController{}.About()).Methods("GET")
	router.HandleFunc("/explore", controllers.MainController{}.Explore()).Methods("GET")
	router.HandleFunc("/", controllers.MainController{}.Index()).Methods("GET")

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("resources"))))

	return router
}
