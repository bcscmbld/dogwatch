package main

import (
	"github.com/gugahoi/dogwatch/pkg/dognzb"
)

func main() {
	dogNZBAPI := ""

	d := dognzb.New(dogNZBAPI)

	// List movies in watchlist
	d.List(dognzb.Movies)

	// List shows in watchlist
	d.List(dognzb.TV)
}
