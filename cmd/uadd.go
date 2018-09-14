package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/donbattery/kkhcli/utils"
	"github.com/spf13/cobra"
)

// uaddCmd represents the uadd command
var uaddCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add new User to the KKHC server",
	Run:     uadd,
}

func init() {
	userCmd.AddCommand(uaddCmd)
}

func getInput(msg string) (string, error) {
	fmt.Print(msg)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(input, "\n"), nil
}

func uadd(cmd *cobra.Command, args []string) {
	type message struct {
		Command  string `json:"command"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var msg utils.Response
	username, err := getInput("Username :")
	if err != nil {
		log.Fatalf("Error reading input %v\n", err)
	}
	email, err := getInput("Email address :")
	if err != nil {
		log.Fatalf("Error reading input %v\n", err)
	}
	randomPass := utils.RandStringBytes(6)
	m := message{"addUser", username, email, randomPass}
	b, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("Error encoding message %v\n", err)
	}
	resp, err := utils.SendCommand("addUser", b)
	if err != nil {
		log.Fatalf("Error processing command\n%v\n", err)
	}
	if err := json.Unmarshal(resp, &msg); err != nil {
		log.Fatalf("Error adding new User %v\n", err)
	}
	fmt.Printf("\n%v\nPassword : %v\n", msg.Msg, randomPass)
}
