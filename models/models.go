// package models

// // Artist data
// type Artist struct {
// 	ID           int      `json:"id"`
// 	Image        string   `json:"image"`
// 	Name         string   `json:"name"`
// 	Members      []string `json:"members"`
// 	CreationDate int      `json:"creationDate"`
// 	FirstAlbum   string   `json:"firstAlbum"`
// 	Locations    string   `json:"locations"`
// 	ConcertDates string   `json:"concertDates"`
// 	Relations    string   `json:"relations"`
// }

// type Location struct {
// 	ID        int      `json:"id"`
// 	Locations []string `json:"locations"`
// 	Dates     string   `json:"dates"`
// }

// type Date struct {
// 	ID    int      `json:"id"`
// 	Dates []string `json:"dates"`
// }

// type Relation struct {
// 	ID             int                 `json:"id"`
// 	DatesLocations map[string][]string `json:"datesLocations"`
// }

// // To combine everything for one artist
// type ArtistDetails struct {
// 	Artist
// 	LocationData Location
// 	DateData     Date
// 	RelationData Relation
// }

package models

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type LocationsResponse struct {
	Index []Location `json:"index"`
}
type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type DatesResponse struct {
	Index []Date `json:"index"`
}
type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type RelationResponse struct {
	Index []Relation `json:"index"`
}
type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
