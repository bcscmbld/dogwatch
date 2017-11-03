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
	Title       string      `xml:"title"`
	Description string      `xml:"description"`
	UUID        string      `xml:"uuid"`
	Items       []MovieItem `xml:"item"`
}

// MovieItem ...
type MovieItem struct {
	Title         string `xml:"title"`
	ID            int    `xml:"imdbid"`
	Plot          string `xml:"plot"`
	Actors        string `xml:"actors"`
	Genres        string `xml:"genres"`
	Year          int    `xml:"year"`
	Runtime       int    `xml:"runtime"`
	Certification string `xml:"certification"`
	Trailer       string `xml:"trailer"`
	Poster        string `xml:"poster"`
}

//TVItem ...
// type TVItem struct {
// 	Title   string `xml:"title"`
// 	ID      string `xml:"tvdbid"`
// 	Plot    string `xml:"plot"`
// 	Actors  string `xml:"actors"`
// 	Genres  string `xml:"genres"`
// 	Network string `xml:"network"`
// 	Status  string `xml:"status"`
// 	Trailer string `xml:"trailer"`
// 	Poster  string `xml:"poster"`
// }
