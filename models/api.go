package models

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

const (
    ArtistsAPI   = "https://groupietrackers.herokuapp.com/api/artists"
    LocationsAPI = "https://groupietrackers.herokuapp.com/api/locations"
    DatesAPI     = "https://groupietrackers.herokuapp.com/api/dates"
    RelationAPI  = "https://groupietrackers.herokuapp.com/api/relation"
)


// FetchAPIData populates the global Data variable with API results
func FetchAPIData() error {
    client := &http.Client{Timeout: 10 * time.Second}

    // Fetch artists
    artists, err := fetchJSON[[]Artist](client, ArtistsAPI)
    if err != nil {
        return fmt.Errorf("error fetching artists: %w", err)
    }
    Data.Artists = artists

    // Fetch locations
    var locData struct {
        Index []Location `json:"index"`
    }
    if err := fetchJSONStruct(client, LocationsAPI, &locData); err != nil {
        return fmt.Errorf("error fetching locations: %w", err)
    }
    Data.Locations = locData.Index

    // Fetch dates
    var dateData struct {
        Index []Date `json:"index"`
    }
    if err := fetchJSONStruct(client, DatesAPI, &dateData); err != nil {
        return fmt.Errorf("error fetching dates: %w", err)
    }
    Data.Dates = dateData.Index

    // Fetch relations
    var relData struct {
        Index []Relation `json:"index"`
    }
    if err := fetchJSONStruct(client, RelationAPI, &relData); err != nil {
        return fmt.Errorf("error fetching relations: %w", err)
    }
    Data.Relations = relData.Index

    return nil
}

// Generic JSON fetcher for typed responses
func fetchJSON[T any](client *http.Client, url string) (T, error) {
    var result T
    resp, err := client.Get(url)
    if err != nil {
        return result, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return result, fmt.Errorf("API returned status: %d", resp.StatusCode)
    }

    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return result, err
    }

    return result, nil
}

// JSON fetcher for custom struct targets
func fetchJSONStruct(client *http.Client, url string, v interface{}) error {
    resp, err := client.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("API returned status: %d", resp.StatusCode)
    }

    return json.NewDecoder(resp.Body).Decode(v)
}
