package controllers

import (
	"encoding/gob"
	"html/template"
	"log"

	"github.com/gorilla/sessions"
	. "github.com/jonahgeorge/featherlabel.com/models"
)

var (
	t     *template.Template
	store *sessions.CookieStore
)

func init() {
	// compile templates
	t = template.Must(t.ParseGlob("views/_templates/*.html"))
	t = template.Must(t.ParseGlob("views/users/*.html"))
	t = template.Must(t.ParseGlob("views/songs/*.html"))
	t = template.Must(t.ParseGlob("views/auth/*.html"))
	t = template.Must(t.ParseGlob("views/*.html"))
	log.Printf("<== Templates compiled\n")

	// initialize session storage
	store = sessions.NewCookieStore([]byte("octocat"))
	log.Printf("<== Session store configured\n")

	gob.Register(&UserModel{})
	log.Printf("<== UserModel registered\n")
}
