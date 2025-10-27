package main

import (
    "groupie-tracker/models"
    "net/http"
    "strconv"
    "strings"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        renderError(w, http.StatusNotFound, "Page not found", nil)
        return
    }

    if r.Method != http.MethodGet {
        renderError(w, http.StatusBadRequest, "Bad Request: only GET is allowed", nil)
        return
    }

    if len(models.Data.Artists) == 0 {
        renderError(w, http.StatusNotFound, "No artists found", nil)
        return
    }

    view := models.IndexData{
        Artists: models.Data.Artists,
    }

    err := tmpl.ExecuteTemplate(w, "index.html", view)
    if err != nil {
        renderError(w, http.StatusInternalServerError, "Internal Server Error", err)
    }
}

func artistHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        renderError(w, http.StatusBadRequest, "Bad Request: only GET is allowed", nil)
        return
    }

    path := strings.TrimSuffix(r.URL.Path, "/")
    parts := strings.Split(path, "/")
    if len(parts) != 3 || parts[1] != "artist" || parts[2] == "" {
        renderError(w, http.StatusNotFound, "Page not found", nil)
        return
    }

    idStr := parts[2]
    id, err := strconv.Atoi(idStr)
    if err != nil || id <= 0 {
        renderError(w, http.StatusBadRequest, "Bad Request: invalid artist id", err)
        return
    }

    var artist *models.Artist
    for i := range models.Data.Artists {
        if models.Data.Artists[i].ID == id {
            artist = &models.Data.Artists[i]
            break
        }
    }
    if artist == nil {
        renderError(w, http.StatusNotFound, "Artist not found", nil)
        return
    }

    var location *models.Location
    for i := range models.Data.Locations {
        if models.Data.Locations[i].ID == id {
            location = &models.Data.Locations[i]
            break
        }
    }

    var dates *models.Date
    for i := range models.Data.Dates {
        if models.Data.Dates[i].ID == id {
            dates = &models.Data.Dates[i]
            break
        }
    }

    var relation *models.Relation
    for i := range models.Data.Relations {
        if models.Data.Relations[i].ID == id {
            relation = &models.Data.Relations[i]
            break
        }
    }

    view := models.ArtistDetailData{
        Artist:    *artist,
        Relation:  models.Relation{},
        Locations: nil,
        Dates:     nil,
    }

    if location != nil {
        view.Locations = location.Locations
    }
    if dates != nil {
        view.Dates = dates.Dates
    }
    if relation != nil {
        view.Relation = *relation
    }

    err = tmpl.ExecuteTemplate(w, "artist.html", view)
    if err != nil {
        renderError(w, http.StatusInternalServerError, "Internal Server Error", err)
    }
}
