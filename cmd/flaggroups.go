// Houses functions to easily add flags to cobra commands.
package cmd

import (
	"github.com/spf13/cobra"
)

// Add to commands with filterable not output.
var jsonFilter string
func filterable(cmd *cobra.Command) {
	cmd.Flags().StringVar(&jsonFilter, "filter", "",
	`Filter by json (eg. {"title": "my little pony screenplay", "tag": "#bighitsforthefuture"}).
If this is given, all other filterflags or arguments are ignored.`)
}

// Add to commands that allow json output.
var jsonOut bool
func jsonable(cmd *cobra.Command) {
	cmd.Flags().BoolVar(&jsonOut, "json", false, "Return json output.")
}
