package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "snapshot"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "dogwatch version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s %s", RootCmd.Use, version)
	},
}
