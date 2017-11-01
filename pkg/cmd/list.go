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
	t := `
	dogwatch list (movie|tv) [-a apikey]
		* list movie or series in the watchlist
		-a dognzb apikey`
	fmt.Fprintf(os.Stderr, t) // nolint: gas
}

// DefineFlags are the flags the list command accepts
func (cmd *List) DefineFlags(fs *flag.FlagSet) {
	cmd.api = fs.String("a", "", "apikey")
}

// Run is what happens when the list cmd is called
func (cmd *List) Run() int {
	d := dognzb.New(*cmd.api)

	if *cmd.api == "" {
		fmt.Fprintf(os.Stderr, "Missing required parameter 'apikey'\n") // nolint: gas
		return 32
	}

	items, err := d.List(dognzb.Movies)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to list: %v\n ", err) // nolint: gas
		return 32
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
	return 0
}
