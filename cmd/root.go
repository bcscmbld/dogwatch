package cmd

import (
	"github.com/spf13/cobra"
)

var api string

func init() {
	RootCmd.PersistentFlags().StringVarP(&api, "api", "a", "", "dognzb apikey")
	RootCmd.AddCommand(listCmd)
}

// RootCmd is the entrypoint into app commands
var RootCmd = &cobra.Command{
	Use:   "dogwtach",
	Short: "dogwatch is a cli tool to interact with DogNZB's Watchlists",
	RunE:  nil,
}
