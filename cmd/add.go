package cmd

import (
	"fmt"
	"net/http"
	"os"

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
	done := make(chan int, len(ids))

	for _, id := range ids {
		go func(id string, done chan<- int) {
			err := d.Add(t, id)
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to add %v: %v\n", id, err) // nolint: gas
			} else {
				fmt.Fprintf(os.Stdout, "added %v\n", id) // nolint: gas
			}
			done <- 1
		}(id, done)
	}

	for i := 0; i < len(ids); i++ {
		<-done
	}
}
