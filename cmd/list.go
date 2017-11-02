package cmd

import (
	"fmt"
	"os"

	"github.com/gugahoi/dogwatch/pkg/dognzb"
	"github.com/spf13/cobra"
)

func init() {
	listCmd.AddCommand(listMoviesCmd, listTVCmd)
}

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
	d := dognzb.New(api)
	items, err := d.List(t)
	if err != nil {
		return fmt.Errorf("Failed to list movies: %v\n ", err) // nolint: gas
	}

	for _, item := range items {
		// nolint: gas
		fmt.Fprintf(
			os.Stdout,
			"%s | %d | tt%d\n",
			item.Title,
			item.Year,
			item.ImdbID,
		)
	}
	return nil
}