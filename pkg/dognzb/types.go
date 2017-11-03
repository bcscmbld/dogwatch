package dognzb

import (
	"fmt"
)

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
	TVdbID        string `xml:"tvdbid"`
	ImdbID        string `xml:"imdbid"`
	Plot          string `xml:"plot"`
	Actors        string `xml:"actors"`
	Genres        string `xml:"genres"`
	Year          string `xml:"year"`
	Runtime       int    `xml:"runtime"`
	Certification string `xml:"certification"`
	Trailer       string `xml:"trailer"`
	Poster        string `xml:"poster"`
	Network       string `xml:"network"`
	Status        string `xml:"status"`
}

// GetID return then appropriate id from the Item (tvdbid if an tv show, imdbid if a movie)
func (i *Item) GetID() string {
	if i.TVdbID != "" {
		return i.TVdbID
	}
	return fmt.Sprintf("tt%s", i.ImdbID)
}
