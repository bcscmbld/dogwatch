package cmd

import (
	"fmt"
	"net/http"

	"github.com/gugahoi/dogwatch/pkg/dognzb"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add items to the watchlist",
}

var addMoviesCmd = &cobra.Command{
	Use:   "movies",
	Short: "Add movies to the watchlist",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		add(dognzb.Movies, args)
	},
}

var addTVCmd = &cobra.Command{
	Use:   "tv",
	Short: "Add tv shows to the watchlist",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		add(dognzb.TV, args)
	},
}

func add(t dognzb.Type, ids []string) {
	d := dognzb.New(api, &http.Client{})
	done := make(chan string, 1)

	for _, id := range ids {
		go routine(d.Add, t, id, done)
	}

	for i := 0; i < len(ids); i++ {
		fmt.Println(<-done)
	}
}
