package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	flag "github.com/ogier/pflag"
)

var (
	list bool
	seed bool
	userToAdd string
	userToReset string
	userToEnable string
	userToDisable string
	apiURL string
	envPass string
)

type Response struct {
	Msg string `json:"msg"`
}

func main() {

	flag.Parse()
	if flag.NFlag() == 0 {
		printUsage()
	}

	envPass = os.Getenv("ADMIN_PASSWORD")
	if envPass == "" {
		fmt.Println("\nNo ADMIN_PASSWORD found in the environment\n")
		fmt.Print("Admin password : ")
		reader := bufio.NewReader(os.Stdin)
		envPass, _ = reader.ReadString('\n')
		envPass = strings.TrimSuffix(envPass, "\n")
		fmt.Println()
	}
	
	if seed {
		fmt.Println("\nSeeding DB ...\n")
		response:= Response{}
		response = seedDB()
		fmt.Println(response.Msg, "\n")
	}

	if list {
		users := getUsers(envPass)
		fmt.Println("\nList of Users\n")
		for _, u := range users {
			fmt.Println("Username: ", u.Username)
			fmt.Println("Enabled: ", u.Enabled, "\n")
		}
	}

	if userToAdd != "" {
		fmt.Println("User to add :", userToAdd)
		fmt.Print("Password : ")
		reader := bufio.NewReader(os.Stdin)
		password, _ := reader.ReadString('\n')
		password = strings.TrimSuffix(password, "\n")
		if password == "" {
			exit("No Password provided !")
		} else {
			response:= Response{}
			response = addUser(envPass, userToAdd, password)
			fmt.Println(response.Msg, "\n")
		}
	}

	if userToReset != "" {
		fmt.Println("User to reset :", userToReset)
		fmt.Print("New Password : ")
		reader := bufio.NewReader(os.Stdin)
		password, _ := reader.ReadString('\n')
		password = strings.TrimSuffix(password, "\n")
		if password == "" {
			exit("No Password provided !")
		} else {
			response:= Response{}
			response = resetUser(envPass, userToReset, password)
			fmt.Println(response.Msg, "\n")
		}
	}

}
	
func init() {
	flag.BoolVarP(&list, "list", "l", false, "List users")
	flag.BoolVarP(&seed, "seed", "s", false, "Seed database")
	flag.StringVarP(&userToAdd, "add", "a", "", "Add user")
	flag.StringVarP(&userToReset, "reset", "r", "", "Reset password")
	flag.StringVarP(&userToDisable, "disable", "d", "", "Disable user")
	flag.StringVarP(&userToEnable, "enable", "e", "", "Enable user")
	flag.StringVarP(&apiURL, "url", "u", "https://kkhc.eu/admin", "WEB URL")
}
	
func printUsage() {
	fmt.Println("\n . o O KKHCLI O o .\n")
	fmt.Printf("Usage: %s [option] [username]\n\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	fmt.Println()
	os.Exit(1)
}

func exit(msg string) {
	fmt.Println("\n Error :", msg, "\n")

	os.Exit(2)
} 
