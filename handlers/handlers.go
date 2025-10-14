package handlers

import (
    "groupie-tracker/models"
    "html/template"
    "log"
    "net/http"
    "sort"
    "strconv"
    "strings"
)

var templates *template.Template

func InitTemplates() {
    var err error
    templates, err = template.ParseGlob("templates/*.html")
    if err != nil {
        log.Fatal("Error loading templates:", err)
    }
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        renderError(w, "Page not found", http.StatusNotFound)
        return
    }

    data := models.IndexData{
        Artists: models.Data.Artists,
    }

    if err := templates.ExecuteTemplate(w, "index.html", data); err != nil {
        log.Printf("Template error: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/artist/")
    id, err := strconv.Atoi(idStr)
    if err != nil || id < 1 || id > len(models.Data.Artists) {
        renderError(w, "Invalid artist ID", http.StatusBadRequest)
        return
    }

    artist := models.Data.Artists[id-1]
    var relation models.Relation
    var locations []string
    var dates []string

    for _, rel := range models.Data.Relations {
        if rel.ID == id {
            relation = rel
            break
        }
    }

    locationMap := make(map[string]bool)
    for loc, concertDates := range relation.DatesLocations {
        formattedLoc := formatLocation(loc)
        locationMap[formattedLoc] = true
        dates = append(dates, concertDates...)
    }

    for loc := range locationMap {
        locations = append(locations, loc)
    }
    sort.Strings(locations)
    sort.Strings(dates)

    data := models.ArtistDetailData{
        Artist:    artist,
        Relation:  relation,
        Locations: locations,
        Dates:     dates,
    }

    if err := templates.ExecuteTemplate(w, "artist.html", data); err != nil {
        log.Printf("Template error: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
    query := strings.TrimSpace(r.URL.Query().Get("q"))
    if query == "" {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    results := searchArtists(query)

    data := models.SearchResult{
        Artists:      results,
        Query:        query,
        ResultsCount: len(results),
    }

    if err := templates.ExecuteTemplate(w, "search.html", data); err != nil {
        log.Printf("Template error: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}

func FilterHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    if err := r.ParseForm(); err != nil {
        renderError(w, "Invalid form data", http.StatusBadRequest)
        return
    }

    minYear := r.FormValue("minYear")
    maxYear := r.FormValue("maxYear")
    minMembers := r.FormValue("minMembers")
    maxMembers := r.FormValue("maxMembers")
    location := strings.TrimSpace(r.FormValue("location"))

    filtered := filterArtists(minYear, maxYear, minMembers, maxMembers, location)

    data := models.IndexData{
        Artists: filtered,
    }

    if err := templates.ExecuteTemplate(w, "index.html", data); err != nil {
        log.Printf("Template error: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}

func searchArtists(query string) []models.Artist {
    query = strings.ToLower(query)
    var results []models.Artist

    for _, artist := range models.Data.Artists {
        if strings.Contains(strings.ToLower(artist.Name), query) {
            results = append(results, artist)
            continue
        }
        for _, member := range artist.Members {
            if strings.Contains(strings.ToLower(member), query) {
                results = append(results, artist)
                break
            }
        }
        if strings.Contains(strconv.Itoa(artist.CreationDate), query) {
            results = append(results, artist)
            continue
        }
        if strings.Contains(strings.ToLower(artist.FirstAlbum), query) {
            results = append(results, artist)
        }
    }

    return results
}

func filterArtists(minYear, maxYear, minMembers, maxMembers, location string) []models.Artist {
    var filtered []models.Artist

    for _, artist := range models.Data.Artists {
        if minYear != "" {
            min, _ := strconv.Atoi(minYear)
            if artist.CreationDate < min {
                continue
            }
        }
        if maxYear != "" {
            max, _ := strconv.Atoi(maxYear)
            if artist.CreationDate > max {
                continue
            }
        }

        memberCount := len(artist.Members)
        if minMembers != "" {
            min, _ := strconv.Atoi(minMembers)
            if memberCount < min {
                continue
            }
        }
        if maxMembers != "" {
            max, _ := strconv.Atoi(maxMembers)
            if memberCount > max {
                continue
            }
        }

        if location != "" {
            hasLocation := false
            for _, rel := range models.Data.Relations {
                if rel.ID == artist.ID {
                    for loc := range rel.DatesLocations {
                        if strings.Contains(strings.ToLower(loc), strings.ToLower(location)) {
                            hasLocation = true
                            break
                        }
                    }
                    break
                }
            }
            if !hasLocation {
                continue
            }
        }

        filtered = append(filtered, artist)
    }

    return filtered
}

func formatLocation(loc string) string {
    loc = strings.ReplaceAll(loc, "-", " ")
    loc = strings.ReplaceAll(loc, "_", " ")
    words := strings.Fields(loc)
    for i, word := range words {
        words[i] = strings.Title(strings.ToLower(word))
    }
    return strings.Join(words, " ")
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
    w.WriteHeader(statusCode)
    data := struct {
        Error      string
        StatusCode int
    }{
        Error:      message,
        StatusCode: statusCode,
    }
    if err := templates.ExecuteTemplate(w, "error.html", data); err != nil {
        log.Printf("Error template error: %v", err)
        http.Error(w, message, statusCode)
    }
}
