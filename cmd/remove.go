package cmd

import (
	"fmt"
	"net/http"

	"github.com/gugahoi/dogwatch/pkg/dognzb"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove items from the watchlist",
}

var removeMoviesCmd = &cobra.Command{
	Use:   "movies",
	Short: "remove movies from the watchlist",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		remove(dognzb.Movies, args)
	},
}

var removeTVCmd = &cobra.Command{
	Use:   "tv",
	Short: "remove tv shows from the watchlist",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		remove(dognzb.TV, args)
	},
}

func remove(t dognzb.Type, ids []string) {
	d := dognzb.New(api, &http.Client{})
	done := make(chan string, 1)

	for _, id := range ids {
		go routine(d.Remove, t, id, done)
	}

	for i := 0; i < len(ids); i++ {
		fmt.Println(<-done)
	}
}
