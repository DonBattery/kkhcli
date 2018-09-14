package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/donbattery/kkhcli/utils"
	"github.com/spf13/cobra"
)

// ulistCmd represents the ulist command
var ulistCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List the Users",
	Run: func(cmd *cobra.Command, args []string) {
		var users struct {
			Msg []utils.User `json:"msg"`
		}
		var msg utils.Response
		resp, err := utils.SendCommand("getUsersInfo", nil)
		if err != nil {
			log.Fatalf("Error processing command\n%v\n", err)
		}
		if err := json.Unmarshal(resp, &msg); err == nil {
			log.Fatalf("Error getting users %v\n", msg.Msg)
		}
		if err := json.Unmarshal(resp, &users); err != nil {
			log.Fatalf("Error decoding users %v\n", err)
		}
		fmt.Printf("\nList of users :\n\n")
		for _, user := range users.Msg {
			fmt.Printf("Username : %v\nEmail : %v\n\n", user.Username, user.Email)
		}
	},
}

func init() {
	userCmd.AddCommand(ulistCmd)
}
