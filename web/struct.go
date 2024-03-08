package groupie

type Groupie struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    Location `json:"-"`
	ConcertDates Date     `json:"-"`
	Relations    Relation `json:"-"`
	PrevPage     int
	NextPage     int
}

type Location struct {
	Locations []string `json:"locations"`
}

type Date struct {
	Dates []string `json:"dates"`
}

type Relation struct {
	DatesLocation map[string][]string `json:"datesLocations"`
}

type Err struct {
	StatusCode int
	StatusText string
}
