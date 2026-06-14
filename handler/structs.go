package handler

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Location struct {
	ID       int      `json:"id"`
	Location []string `json:"locations"`
}

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	ID            int                 `json:"id"`
	DatesLocation map[string][]string `json:"datesLocations"`
}

type ArtistPage struct {
	Artist    Artist
	Dates     []string
	Locations []string
	Relations map[string][]string
}

type LocationResponse struct {
	Index []Location `json:"index"`
}

type DatesResponse struct {
	Index []Date `json:"index"`
}

type RelationResponse struct {
	Index []Relation
}

type HomePage struct {
	Artist []Artist
}
