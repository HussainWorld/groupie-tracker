package main

import (
    "html/template"
    "log"
)

var tmpl *template.Template

func initTemplates() {
    var err error
tmpl, err = template.ParseGlob("../templates/*.html")
    if err != nil {
        log.Fatal("Error loading templates:", err)
    }
}
