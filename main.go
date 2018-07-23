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
	listCollections bool
	seed bool
	userToAdd string
	userToReset string
	collectionToFlush string
	apiURL string
	envPass string
)

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
		response := seedDB()
		fmt.Println(response.Msg, "\n")
	}
	
	if list {
		users := getUsers(envPass)
		fmt.Println("\nList of Users\n")
		for _, u := range users {
			fmt.Println(u.Username)
		}
	}
	
	if listCollections {
		collections := getCollections(envPass)
		fmt.Println("\nList of Collections\n")
		for _, coll := range collections {
			fmt.Println(coll.Name)
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
			response := addUser(envPass, userToAdd, password)
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
			response := resetUser(envPass, userToReset, password)
			fmt.Println(response.Msg, "\n")
		}
	}
			
	if collectionToFlush != "" {
		collections := strings.Split(collectionToFlush, ",")
		fmt.Println("\nFlushing collection(s)", collectionToFlush)
		for _, coll := range collections {
			response := flushColllection(coll)
			fmt.Println(response.Msg)
		}			
	}			
}
		
func init() {
	flag.BoolVarP(&list, "list", "l", false, "List users")
	flag.BoolVarP(&seed, "seed", "s", false, "Seed database")
	flag.BoolVarP(&listCollections, "collections", "c", false, "List collections")
	flag.StringVarP(&userToAdd, "add", "a", "", "Add user")
	flag.StringVarP(&userToReset, "reset", "r", "", "Reset password")
	flag.StringVarP(&collectionToFlush, "flush", "f", "", "Flush collection")
	flag.StringVarP(&apiURL, "url", "u", "https://kkhc.eu/admin", "WEB URL")
}
	
func printUsage() {
	fmt.Println("\n . o O KKHCLI O o .\n")
	fmt.Printf("Usage: %s [option] [collection or username]\n\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	fmt.Println()
	os.Exit(1)
}

func exit(msg string) {
	fmt.Println("\n Error :", msg, "\n")
	os.Exit(2)
} 
