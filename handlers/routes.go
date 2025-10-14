package handlers

import (
    "net/http"
)

func RegisterRoutes() {
    http.HandleFunc("/", HomeHandler)
    http.HandleFunc("/artist/", ArtistHandler)
    http.HandleFunc("/search", SearchHandler)
    http.HandleFunc("/filter", FilterHandler)

    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
}
