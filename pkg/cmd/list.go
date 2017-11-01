package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/gugahoi/dogwatch/pkg/dognzb"
)

// List is the list subcommand
type List struct {
	api *string
}

// Name is the name of the list command
func (cmd *List) Name() string {
	return "list"
}

// Usage is the usage instructions for the list command
func (cmd *List) Usage() {
	t := `%[1]s list (movie|tv) [-a apikey]
	* list movie or series in the watchlist
	-a dognzb apikey
	`
	fmt.Fprintf(os.Stderr, t, os.Args[0])
}

// DefineFlags are the flags the list command accepts
func (cmd *List) DefineFlags(fs *flag.FlagSet) {
	cmd.api = fs.String("a", "", "apikey")
}

// Run is what happens when the list cmd is called
func (cmd *List) Run() int {
	d := dognzb.New(*cmd.api)

	if *cmd.api == "" {
		fmt.Fprintf(os.Stderr, "Missing required parameter 'apikey'")
		return 32
	}

	items, err := d.List(dognzb.Movies)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to list: %v", err)
		return 32
	}

	for _, item := range items {
		fmt.Fprintf(os.Stdout, "%s (%d)", item.Title, item.Year)
	}
	return 0
}
