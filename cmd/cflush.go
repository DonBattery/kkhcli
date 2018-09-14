package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/donbattery/kkhcli/utils"
	"github.com/spf13/cobra"
)

// cflushCmd represents the cflush command
var cflushCmd = &cobra.Command{
	Use:     "flush",
	Short:   "Flush collection(s) by name",
	Aliases: []string{"f"},
	Long: `
Flush collection(s) by name
Usage:
kkhcli collection flush [...collectionName]

Example:
kkhcli user reset CommentFlow Avatar Image

See more:
kkhcli collection flush -h
`,
	Run: func(cmd *cobra.Command, args []string) {
		var msg utils.Response
		type message struct {
			Command    string `json:"command"`
			Collection string `json:"collection"`
		}
		if len(args) == 0 {
			cmd.Help()
		}
		for _, collection := range args {
			m := message{"flushDbCollection", collection}
			b, err := json.Marshal(m)
			if err != nil {
				log.Fatalf("Error encoding command %v\n", err)
			}
			resp, err := utils.SendCommand("flushDbCollection", b)
			if err != nil {
				log.Fatalf("Error processing command\n%v\n", err)
			}
			if err := json.Unmarshal(resp, &msg); err != nil {
				log.Fatalf("Error decoding message %v\n", err)
			}
			fmt.Printf("%v\n", msg.Msg)
		}
	},
}

func init() {
	collectionCmd.AddCommand(cflushCmd)
}
