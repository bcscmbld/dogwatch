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

// DogNZB is a struct to talk to dog's api
type DogNZB struct {
	api string
}

// New returns a new dognzb struct
func New(api string) *DogNZB {
	return &DogNZB{
		api: api,
	}
}

func (d *DogNZB) buildURL(verb string, t Type, id string) string {
	params := url.Values{
		"t":      []string{verb},
		"o":      []string{"json"},
		"apikey": []string{d.api},
	}

	return fmt.Sprintf("%s/watchlist?%s=%s&%s", dogNZBURL, t, id, params.Encode())
}

func (d *DogNZB) get(url string) (Query, error) {
	r, err := http.Get(url)
	if err != nil || r.StatusCode != http.StatusOK {
		return Query{}, fmt.Errorf("failed to list: %v", err)
	}

	b, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		return Query{}, fmt.Errorf("failed reading body: %v", err)
	}

	var q Query
	if err := xml.Unmarshal(b, &q); err != nil {
		return Query{}, err
	}

	return q, nil
}

// List lists the item in the appropriate watchlist (tv or movie)
func (d *DogNZB) List(t Type) ([]Item, error) {
	q, err := d.get(d.buildURL("list", t, ""))
	if err != nil {
		return nil, err
	}

	for _, item := range q.Channel.Items {
		fmt.Println(item.Title)
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
