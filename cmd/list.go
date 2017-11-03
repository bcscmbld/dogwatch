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
		var desc string
		if item.Year != "" {
			desc = item.Year
		} else {
			desc = item.Network
		}

		// nolint: gas
		fmt.Fprintf(
			os.Stdout,
			"%s | %s | %s\n",
			item.Title,
			desc,
			item.GetID(),
		)
	}
	return nil
}
