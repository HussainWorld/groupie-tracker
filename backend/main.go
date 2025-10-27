package main

import (
	"fmt"
	"groupie-tracker/models"
	"log"
	"net/http"
	"os"
)

func main() {
	initTemplates()

	if err := models.LoadArtistsFromAPI(); err != nil {
		log.Fatal(err)
	}
	if err := models.LoadDatesFromAPI(); err != nil {
		log.Fatal("dates", err)
	}
	if err := models.LoadLocationFromAPI(); err != nil {
		log.Fatal("locations", err)
	}
	if err := models.LoadRelationsFromAPI(); err != nil {
		log.Fatal("relations", err)
	}

	fs := http.FileServer(http.Dir("../assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/artist/", artistHandler)

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting working directory:", err)
	}
	fmt.Println("Working directory:", dir)
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
