package dognzb

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	dogNZBURL = "https://api.dognzb.cr"
)

// Getter is a http.Get interface
type Getter interface {
	Get(url string) (*http.Response, error)
}

// DogNZB is a struct to talk to DogNZB's api
type DogNZB struct {
	api string
	h   Getter
}

// New returns a new dognzb struct
func New(api string, h Getter) *DogNZB {
	return &DogNZB{
		api: api,
		h:   h,
	}
}

func (d *DogNZB) buildURL(verb string, t Type, id string) string {
	params := url.Values{
		"t":       []string{verb},
		"o":       []string{"json"},
		"apikey":  []string{d.api},
		string(t): []string{id},
	}

	return fmt.Sprintf("%s/watchlist?%s", dogNZBURL, params.Encode())
}

func (d *DogNZB) get(url string) (Query, error) {
	r, err := d.h.Get(url)
	if err != nil {
		return Query{}, fmt.Errorf("failed to list: %v", err)
	}

	if r.StatusCode != http.StatusOK {
		return Query{}, fmt.Errorf("bad response: %v", r.StatusCode)
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close() // nolint: errcheck
	if err != nil {
		return Query{}, fmt.Errorf("failed reading body: %v", err)
	}

	var q Query
	if err := xml.Unmarshal(b, &q); err != nil {
		return Query{}, err
	}

	// if dognzb sent an error back, we should also error
	if q.ErrorCode != 0 {
		return Query{}, fmt.Errorf("%v", q.ErrorDesc)
	}

	return q, nil
}

// List lists the item in the appropriate watchlist (tv or movie)
func (d *DogNZB) List(t Type) ([]MovieItem, error) {
	q, err := d.get(d.buildURL("list", t, ""))
	if err != nil {
		return nil, err
	}

	return q.Channel.Items, nil
}

// Add adds an item to the appropriate watchlist (tv or movie)
func (d *DogNZB) Add(t Type, id string) error {
	q, err := d.get(d.buildURL("add", t, id))
	if err != nil {
		return err
	}

	fmt.Println(q)
	return nil
}

// Remove removes an item from the appropriate watchlist (tv or movie)
func (d *DogNZB) Remove(t Type, id string) error {
	q, err := d.get(d.buildURL("remove", t, id))
	if err != nil {
		return err
	}

	fmt.Println(q)
	return nil
}
