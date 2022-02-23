// Houses functions to easily add flags to cobra commands.
package cmd

import (
	"github.com/spf13/cobra"
)

// Add to commands with filterable not output.
var (
	titleFilter string
	jsonFilter string
	tagFilter string
)
func filterable(cmd *cobra.Command) {
	listCmd.Flags().StringVar(&titleFilter, "title", "", "Filter results by title.")
	listCmd.Flags().StringVar(&tagFilter, "tag", "", "Filter results by tag.")
	listCmd.Flags().StringVar(&jsonFilter, "filter", "",
	`Filter by json (eg. {"title": "my little pony screenplay", "tag": "#bighitsforthefuture"}).
If this is given, all other filterflags like "title" are ignored.`)
}
