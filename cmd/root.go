package cmd

import (
	"fmt"
	"os"

	"github.com/gugahoi/dogwatch/pkg/dognzb"

	"github.com/spf13/cobra"
)

var api string

func init() {
	RootCmd.PersistentFlags().StringVarP(&api, "api", "a", "", "dognzb apikey")
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(listCmd, addCmd, removeCmd)
	listCmd.AddCommand(listMoviesCmd, listTVCmd)
	addCmd.AddCommand(addMoviesCmd, addTVCmd)
	removeCmd.AddCommand(removeMoviesCmd, removeTVCmd)
}

// RootCmd is the entrypoint into app commands
var RootCmd = &cobra.Command{
	Use:   "dogwatch",
	Short: "dogwatch is a cli tool to interact with DogNZB's Watchlists",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return CheckAPI(cmd.Use, &api)
	},
}

// CheckAPI checks if the api has been provided through a flag
// or env variable for the commands that need it, which are
// all but the version and help commands.
func CheckAPI(cmdName string, api *string) error {
	allowedCmds := []string{"version", "help"}

	for _, cmd := range allowedCmds {
		if cmdName == cmd {
			return nil
		}
	}

	if *api != "" {
		return nil
	}

	*api = os.Getenv("DOGNZB_API")
	if *api != "" {
		return nil
	}

	return fmt.Errorf("missing required flag: -a, --apikey")
}

// Both tpye fn and func routine are on their own
// for testing purposes only.
type fn func(t dognzb.Type, id string) (*dognzb.AddRemoveQuery, error)

func routine(f fn, t dognzb.Type, id string, done chan<- string) {
	q, err := f(t, id)
	if err != nil {
		done <- fmt.Sprintf("failed to remove %v: %v", id, err) // nolint: gas
		return
	}
	done <- fmt.Sprintf("%v", q.Description) // nolint: gas
}
