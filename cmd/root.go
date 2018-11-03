package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "kkhcli",
	Short: "KKHCLI is an admin tool for the KKHC Server",
	Long: `
KKHCLI is an admin tool for the KKHC Server

With thel help of KKHCLI you can list, add, remove
users, collections, avatars. And you can seed the KKHC database as well
`,
}
