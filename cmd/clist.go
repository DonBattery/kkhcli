package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/donbattery/kkhcli/utils"
	"github.com/spf13/cobra"
)

// clistCmd represents the clist command
var clistCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List of collections",

	Run: func(cmd *cobra.Command, args []string) {
		var collections struct {
			Msg []utils.Collection `json:"msg"`
		}
		var msg utils.Response
		resp, err := utils.SendCommand("listCollections", nil)
		if err != nil {
			log.Fatalf("Error processing command\n%v\n", err)
		}
		if err := json.Unmarshal(resp, &msg); err == nil {
			log.Fatalf("Error getting collections %v\n", msg.Msg)
		}
		if err := json.Unmarshal(resp, &collections); err != nil {
			log.Fatalf("Error decoding collections %v\n", err)
		}
		fmt.Printf("\nList of collections :\n\n")
		for _, collection := range collections.Msg {
			fmt.Printf("%v\n", collection.Name)
		}
	},
}

func init() {
	collectionCmd.AddCommand(clistCmd)
}
