package controllers

import "html/template"

var t *template.Template

func init() {
	t = template.Must(t.ParseGlob("views/shared/*.html"))
	t = template.Must(t.ParseGlob("views/users/*.html"))
	t = template.Must(t.ParseGlob("views/songs/*.html"))
	t = template.Must(t.ParseGlob("views/*.html"))
}
