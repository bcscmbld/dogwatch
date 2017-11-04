package cmd

import (
	"fmt"
	"net/http"
	"os"

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
	done := make(chan int, len(ids))

	for _, id := range ids {
		go func(id string, done chan<- int) {
			q, err := d.Remove(t, id)
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to remove %v: %v\n", id, err) // nolint: gas
			} else {
				fmt.Fprintf(os.Stdout, "%v\n", q.Description) // nolint: gas
			}
			done <- 1
		}(id, done)
	}

	for i := 0; i < len(ids); i++ {
		<-done
	}
}
