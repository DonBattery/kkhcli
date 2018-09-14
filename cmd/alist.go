package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/donbattery/kkhcli/utils"
	"github.com/spf13/cobra"
)

// alistCmd represents the alist command
var alistCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List of Avatars",
	Run: func(cmd *cobra.Command, args []string) {
		var avatars struct {
			Msg []utils.Avatar `json:"msg"`
		}
		var msg utils.Response
		resp, err := utils.SendCommand("getAvatars", nil)
		if err != nil {
			log.Fatalf("Error processing command\n%v\n", err)
		}
		if err := json.Unmarshal(resp, &msg); err == nil {
			log.Fatalf("Error getting avatars %v\n", msg.Msg)
		}
		if err := json.Unmarshal(resp, &avatars); err != nil {
			log.Fatalf("Error decoding Avatars %v\n", err)
		}
		fmt.Printf("\nList of Avatars :\n\n")
		for _, avatar := range avatars.Msg {
			fmt.Printf("Filename : %v\tExtension : %v\n", avatar.NameOnDisc, avatar.Extenison)
		}
	},
}

func init() {
	avatarCmd.AddCommand(alistCmd)
}
