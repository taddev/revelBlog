/*
	Command-line admin application to manage the database directly. This
	program has unfetered access to everything in the database so use it
	carefully.
*/
package main

import (
	"bufio"
	"code.google.com/p/go.crypto/bcrypt"
	"fmt"
	"github.com/robfig/revel"
	"github.com/taddevries/lazyboy"
	"github.com/taddevries/revelBlog/app/models"
	"os"
	"strings"
)

func AddUser(username, password, display string) {
	bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	stringPassword := string(bcryptPassword)
	demoUser := models.User{Id: username, Rev: "", Username: username, DisplayName: display, Password: stringPassword}

	lazyboy.Database.Insert(demoUser)
	fmt.Println("User Added..")
}

func loadApp(args []string) {
	mode := "dev"

	// Find and parse app.conf
	revel.Init(mode, args[0], "")
	revel.LoadMimeConfig()

	lazyboy.AppInit()
}

func printHelp() {
	fmt.Println("\nadmin your/app/path [command]\n")
	fmt.Println("List of possible commands:")
	fmt.Println("--------------------------")
	fmt.Println("adduser [Create a new user]")
	fmt.Println("\n")
}

func main() {
	args := os.Args[1:]

	if len(args) >= 2 {
		switch args[1] {
		case "adduser":

			fmt.Println("Loading configuration...")
			loadApp(args)

			var name, username, password string
			reader := bufio.NewReader(os.Stdin)

			fmt.Println("Adding a new user")
			fmt.Println("-----------------")

			fmt.Println("What is the user's real name? ")
			name, _ = reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Println("What should the username be? ")
			username, _ = reader.ReadString('\n')
			username = strings.TrimSpace(username)

			fmt.Println("Please enter a strong password: ")
			password, _ = reader.ReadString('\n')
			password = strings.TrimSpace(password)

			AddUser(username, password, name)

		case "test":
			fmt.Println("Loading Configuration..")
			loadApp(args)
			fmt.Println(lazyboy.DBUrl)
		default:
			printHelp()
		}
	} else {
		printHelp()
	}

}
