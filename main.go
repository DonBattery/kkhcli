package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"

	"github.com/donbattery/kkhcli/cmd"
	"github.com/donbattery/kkhcli/utils"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatalf("Error during execution : %v\n", err)
	}
}

func init() {
	// Load config from cli.yaml and ENV
	if err := utils.ReadConfig("./", "kkhcli", map[string]interface{}{
		"adminPass": "",
		"apiURL":    "https://kkhc.eu/admin",
	}); err != nil {
		fmt.Printf("Cannot set configuration %v\n", err)
	}
	// Setup global flags
	cmd.RootCmd.PersistentFlags().StringP("adminPass", "p", "", "Admin Password")
	cmd.RootCmd.PersistentFlags().StringP("apiURL", "u", "", "API URL")
	if err := viper.BindPFlag("adminPass", cmd.RootCmd.PersistentFlags().Lookup("adminPass")); err != nil {
		fmt.Printf("Cannot bind flag adminPass %v\n", err)
	}
	if err := viper.BindPFlag("apiURL", cmd.RootCmd.PersistentFlags().Lookup("apiURL")); err != nil {
		fmt.Printf("Cannot bind flag apiURL %v\n", err)
	}
}
