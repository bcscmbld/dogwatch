package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var api string

func init() {
	RootCmd.PersistentFlags().StringVarP(&api, "api", "a", "", "dognzb apikey")
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(listCmd, addCmd)
	listCmd.AddCommand(listMoviesCmd, listTVCmd)
	addCmd.AddCommand(addMoviesCmd, addTVCmd)
}

// RootCmd is the entrypoint into app commands
var RootCmd = &cobra.Command{
	Use:   "dogwatch",
	Short: "dogwatch is a cli tool to interact with DogNZB's Watchlists",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Name() == "version" {
			return nil
		}

		if api == "" {
			api = os.Getenv("DOGNZB_API")
			if api == "" {
				return fmt.Errorf("missing required flag: -a, --apikey")
			}
		}
		return nil
	},
}
