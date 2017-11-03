package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gugahoi/dogwatch/pkg/dognzb"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List items from watchlist",
	RunE:  nil,
}

var listMoviesCmd = &cobra.Command{
	Use:   "movies",
	Short: "List movies from watchlist",
	RunE: func(cmd *cobra.Command, args []string) error {
		return list(dognzb.Movies)
	},
}

var listTVCmd = &cobra.Command{
	Use:   "tv",
	Short: "List tv shows from watchlist",
	RunE: func(cmd *cobra.Command, args []string) error {
		return list(dognzb.TV)
	},
}

func list(t dognzb.Type) error {
	d := dognzb.New(api, &http.Client{})
	items, err := d.List(t)
	if err != nil {
		return fmt.Errorf("Failed to list: %v\n ", err) // nolint: gas
	}

	for _, item := range items {
		// nolint: gas
		fmt.Fprintf(
			os.Stdout,
			"%s | %d | tt%d\n",
			item.Title,
			item.Year,
			item.ID,
		)
	}
	return nil
}
