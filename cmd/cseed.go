package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/donbattery/kkhcli/utils"
	"github.com/spf13/cobra"
)

// cseedCmd represents the cseed command
var cseedCmd = &cobra.Command{
	Use:     "seed",
	Aliases: []string{"s"},
	Short:   "Seed the database with base folderstructure",
	Run: func(cmd *cobra.Command, args []string) {
		var msg utils.Response
		resp, err := utils.SendCommand("seedDbWithDriveFiles", nil)
		if err != nil {
			log.Fatalf("Error processing command\n%v\n", err)
		}
		if err := json.Unmarshal(resp, &msg); err != nil {
			log.Fatalf("Error decodig message %v\n", err)
		}
		fmt.Println(msg.Msg)
	},
}

func init() {
	collectionCmd.AddCommand(cseedCmd)
}
