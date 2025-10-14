package main

import (
	"groupie-tracker/models"
    "groupie-tracker/handlers"
    "fmt"
    "log"
    "net/http"

    
)

func main() {
    handlers.InitTemplates()

    if err := models.FetchAPIData(); err != nil {
        log.Fatal("Error fetching API data:", err)
    }

    handlers.RegisterRoutes()

    fmt.Println("Server started at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
