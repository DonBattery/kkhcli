package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"regexp"
	flag "github.com/ogier/pflag"
)

var (
	dev bool
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

	if dev {
		apiURL = "http://localhost:3099/admin"
		fmt.Println("\nDevelopment mode\napiURL set to http://localhost:3099/admin\n")
	}

	envPass = os.Getenv("ADMIN_PASSWORD")
	if envPass == "" {
		fmt.Println("\nNo ADMIN_PASSWORD found in the environment\n")
		envPass = getInput("Admin password : ")
		fmt.Println()
	}
	
	if seed {
		fmt.Println("\nSeeding DB ...\n")
		response := seedDB()
		fmt.Println(response.Msg, "\n")
	}
	
	if list {
		users := getUsers()
		fmt.Println("\nList of Users\n")
		for _, u := range users {
			fmt.Println(u.Username)
		}
	}
	
	if listCollections {
		collections := getCollections()
		fmt.Println("\nList of Collections\n")
		for _, coll := range collections {
			fmt.Println(coll.Name)
		}
	}
	
	if userToAdd != "" {
		fmt.Println("User to add :", userToAdd)
		password := getInput("Password :")
		if !isValidPassword(password) {
			exit("Invalid password !\nit must be et least 6 character long")
		} else {
			email := getInput("Email address :")
			if !isValidEmail(email) {
				exit("Invalid Email address")
			}
			response := addUser(userToAdd, password, email)
			fmt.Println(response.Msg, "\n")
		}
	}
		
	if userToReset != "" {
		fmt.Println("User to reset :", userToReset)
		password := getInput("New Password :")
		if !isValidPassword(password) {
			exit("Invalid password")
		} else {
			response := resetUser(userToReset, password)
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
	flag.BoolVarP(&dev, "dev", "d", false, "Development Mode")
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
	fmt.Println("All password must be et least 6 characters long\n")
	fmt.Println("All Email address should be in legal format")
	fmt.Println("for example: foo@bar.baz\n")
	fmt.Println("You can flush multiple Collections at once by separating them with a single comma ,")
	fmt.Println("for example: kkhcli -f Folder,Tag,Image\n")
	fmt.Println("You can chain boolean type options in a single flag")
	fmt.Println("for example: kkhcli -dlc\n")
	os.Exit(1)
}

func exit(msg string) {
	fmt.Println("\n Error :", msg, "\n")
	os.Exit(2)
}

func getInput(msg string) string {
	fmt.Print(msg)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSuffix(input, "\n")
}

func isValidPassword(pass string) bool {
	return regexp.MustCompile(`\w{6,}`).MatchString(pass)
}


func isValidEmail(email string) bool {
	return regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$").MatchString(email)
}
