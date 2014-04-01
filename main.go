package main

import (
	//   "log"
	//   "fmt"
	"github.com/gorilla/mux"
	"github.com/gosexy/to"
	"github.com/gosexy/yaml"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Landing())
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}
