package controllers

import (
	"encoding/gob"
	"html/template"
	"log"

	"github.com/gorilla/sessions"
	"github.com/gosexy/to"
	"github.com/gosexy/yaml"
	"github.com/jonahgeorge/featherlabel.com/models"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
)

var (
	t      *template.Template
	store  *sessions.CookieStore
	bucket *s3.Bucket
)

func init() {
	// Load config file
	conf, err := yaml.Open("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// compile templates
	t = template.Must(t.ParseGlob("views/_templates/*.html"))
	t = template.Must(t.ParseGlob("views/users/*.html"))
	t = template.Must(t.ParseGlob("views/songs/*.html"))
	t = template.Must(t.ParseGlob("views/auth/*.html"))
	t = template.Must(t.ParseGlob("views/*.html"))

	// initialize session storage
	store = sessions.NewCookieStore([]byte(to.String(conf.Get("server", "secret"))))

	gob.Register(&models.UserModel{})

	// load variables from config
	access := to.String(conf.Get("amazon", "s3", "access"))
	secret := to.String(conf.Get("amazon", "s3", "secret"))
	name := to.String(conf.Get("amazon", "s3", "name"))

	// configure aws authentication
	auth, err := aws.GetAuth(access, secret)
	if err != nil {
		log.Fatal(err)
	}

	// create s3 client
	client := s3.New(auth, aws.USWest2)

	// retrieve bucket from name
	bucket = client.Bucket(name)
}
