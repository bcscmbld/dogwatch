package main

import (
	"fmt"
	"os"

	"github.com/gugahoi/dogwatch/cmd"
)

var version string

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err) // nolint: gas
		os.Exit(1)
	}
}
