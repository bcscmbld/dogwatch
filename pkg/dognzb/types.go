package dognzb

import (
	"encoding/xml"
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

// ListQuery is a struct that maps to the XML return format of the "list" call
type ListQuery struct {
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

// AddQuery is the struct that represents the return xml format of the "add" call
type AddQuery struct {
	ErrorCode   string
	ErrorDesc   string
	Code        string
	Description string
}

// UnmarshalXML is a custom way to convert the xml to struct.
// As DogNZB returns the same attrs with different xml names (uuid and error) it becomes hard to figure out
// if the requests really failed or not. Ideally the http status code would be appropriate but it is always 200.
// It would be nice to know if there is a better way to do this...
func (aq *AddQuery) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	switch start.Name.Local {
	case "uuid":
		for _, attr := range start.Attr {
			if attr.Name.Local == "code" {
				aq.Code = attr.Value
			}
			if attr.Name.Local == "description" {
				aq.Description = attr.Value
			}
		}
	case "error":
		for _, attr := range start.Attr {
			if attr.Name.Local == "code" {
				aq.ErrorCode = attr.Value
			}
			if attr.Name.Local == "description" {
				aq.ErrorDesc = attr.Value
			}
		}
	}

	// this seems oddly useless, need to find a better way to achieve it.
	for {
		t, _ := d.Token()
		switch tt := t.(type) {
		case xml.EndElement:
			if tt == start.End() {
				return nil
			}
		}
	}
}
