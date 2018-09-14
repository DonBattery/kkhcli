package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// avatarCmd represents the avatar command
var avatarCmd = &cobra.Command{
	Use:     "avatar",
	Aliases: []string{"a"},
	Short:   "Root command of avatars. Subcommands: list add seed",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
Root command of avatars. Subcommands: list add seed

Usage: 
kkhcli avatar [options]

Example:
kkhcli avatar list

See more:
kkhcli avatar -h`)
	},
}

func init() {
	RootCmd.AddCommand(avatarCmd)
}
