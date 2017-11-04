package main

import (
	"fmt"
	"os"

	"github.com/gugahoi/dogwatch/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err) // nolint: gas
		os.Exit(1)
	}
}
