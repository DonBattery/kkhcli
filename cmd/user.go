package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:     "user",
	Aliases: []string{"u"},
	Short:   "Root of User command. Subcommands: list add reset",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
Root of the User command. Subcommands: list add reset
		
Usage: 
kkhcli user [options]

Example:
kkhcli user list

See more:
kkhcli user -h`)
	},
}

func init() {
	RootCmd.AddCommand(userCmd)
}
