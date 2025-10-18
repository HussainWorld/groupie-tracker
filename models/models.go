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

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type APIData struct {
	Artists   []Artist
	Locations []Location
	Dates     []Date
	Relations []Relation
}

// Shared global data container
var Data APIData

type IndexData struct {
	Artists []Artist
	Error   string
}

type ArtistDetailData struct {
	Artist    Artist
	Relation  Relation
	Locations []string
	Dates     []string
	Error     string
}

type SearchResult struct {
	Artists      []Artist
	Query        string
	ResultsCount int
}
