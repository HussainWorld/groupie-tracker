package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ArtistsAPI   = "https://groupietrackers.herokuapp.com/api/artists"
	LocationsAPI = "https://groupietrackers.herokuapp.com/api/locations"
	DatesAPI     = "https://groupietrackers.herokuapp.com/api/dates"
	RelationAPI  = "https://groupietrackers.herokuapp.com/api/relation"
)

type LocationsResponse struct {
	Index []Location `json:"index"`
}
type RelationsResponse struct {
	Index []Relation `json:"index"`
}
type DatesResponse struct {
	Index []Date `json:"index"`
}

func LoadLocationFromAPI() error {
	resp, err := http.Get(LocationsAPI)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Create a wrapper struct to match the JSON
	var response LocationsResponse

	// Decode into the wrapper
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("failed to decode locations: %w", err)
	}

	// Assign the slice from the wrapper to the global locations variable
	Data.Locations = response.Index
	// fmt.Println(Data.Locations)
	return nil
}

func LoadRelationsFromAPI() error {
	resp, err := http.Get(RelationAPI)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Create a wrapper struct to match the JSON
	var response RelationsResponse

	// Decode into the wrapper
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("failed to decode locations: %w", err)
	}

	// Assign the slice from the wrapper to the global locations variable
	Data.Relations = response.Index
	// fmt.Println(Data.Locations)
	return nil
}

func LoadDatesFromAPI() error {
	resp, err := http.Get(DatesAPI)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Create a wrapper struct to match the JSON
	var response DatesResponse

	// Decode into the wrapper
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("failed to decode locations: %w", err)
	}

	// Assign the slice from the wrapper to the global locations variable
	Data.Dates = response.Index
	// fmt.Println(Data.Locations)
	return nil
}

func LoadArtistsFromAPI() error {
	resp, err := http.Get(ArtistsAPI)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Create a wrapper struct to match the JSON
	var response []Artist

	// Decode into the wrapper
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("failed to decode locations: %w", err)
	}

	// Assign the slice from the wrapper to the global locations variable
	Data.Artists = response
	RemoveInappropriatePic()
	return nil
}

func RemoveInappropriatePic() {
	for i := range Data.Artists {
		if Data.Artists[i].ID == 21 {
			Data.Artists[i].Image = "/assets/images/%E2%9D%8C.png"
		}
	}
}
