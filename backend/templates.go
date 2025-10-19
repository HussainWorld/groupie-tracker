package main

import "html/template"

var tmpl = template.Must(template.ParseFiles(
	"../templates/index.html",
	"../templates/artist.html",
	"../templates/error.html",
))
