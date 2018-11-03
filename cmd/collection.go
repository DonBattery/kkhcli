package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// collectionCmd represents the collection command
var collectionCmd = &cobra.Command{
	Use:     "collection",
	Aliases: []string{"coll", "col", "c"},
	Short:   "Root command of Collections. Subcommands: list, flush, seed",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
Root command of Collections. Subcommands: list, flush, seed

Usage: 
kkhcli collection [options]

Example:
kkhcli collection list

See more:
kkhcli collection -h`)
	},
}

func init() {
	RootCmd.AddCommand(collectionCmd)
}
