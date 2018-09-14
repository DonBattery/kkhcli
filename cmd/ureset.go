package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/donbattery/kkhcli/utils"

	"github.com/spf13/cobra"
)

// UresetCmd represents the ureset command
var UresetCmd = &cobra.Command{
	Use:     "reset",
	Aliases: []string{"res", "r"},
	Short:   "Resets given user's password",
	Long: `
Reset given user's password
Usage:
kkhcli user reset [...emailaddress]

Example:
kkhcli user reset hogyne@gmail.com tegedis@freemail.hu

`,
	Run: ureset,
}

func ureset(cmd *cobra.Command, args []string) {
	type message struct {
		Command  string `json:"command"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var msg utils.Response
	if len(args) == 0 {
		cmd.Help()
	}
	for _, email := range args {
		randomPass := utils.RandStringBytes(6)
		m := message{"resetUserPassword", email, randomPass}
		b, err := json.Marshal(m)
		if err != nil {
			log.Fatalf("Error encoding command %v\n", err)
		}
		resp, err := utils.SendCommand("resetUserPassword", b)
		if err != nil {
			log.Fatalf("Error processing command\n%v\n", err)
		}
		if err := json.Unmarshal(resp, &msg); err != nil {
			log.Fatalf("Decoding message %v\n", err)
		}
		fmt.Printf("%v\nNew password: %v\n", msg.Msg, randomPass)
	}
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	userCmd.AddCommand(UresetCmd)
}
