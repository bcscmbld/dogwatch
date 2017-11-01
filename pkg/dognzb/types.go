package dognzb

// Type ...
type Type string

const (
	// TV is the type of tv for the watchlist features
	TV Type = "tvdbid"
	// Movies is the type of movie for the watchlist features
	Movies Type = "imdbid"
)

// Query ...
type Query struct {
	Channel   Channel `xml:"channel"`
	Version   string  `xml:"version,attr"`
	ErrorCode int     `xml:"code,attr"`
	ErrorDesc string  `xml:"description,attr"`
}

// Channel ...
type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	UUID        string `xml:"uuid"`
	Items       []Item `xml:"item"`
}

// Item ...
type Item struct {
	Title         string `xml:"title"`
	ImdbID        int    `xml:"imdbid"`
	Plot          string `xml:"plot"`
	Actors        string `xml:"actors"`
	Genres        string `xml:"genres"`
	Year          int    `xml:"year"`
	Runtime       int    `xml:"runtime"`
	Certification string `xml:"certification"`
	Trailer       string `xml:"trailer"`
	Poster        string `xml:"poster"`
}
