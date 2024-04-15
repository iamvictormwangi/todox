package main

import (
	"fmt"
	"log"
	// "os"
	"os/user"
	"path/filepath"
)

func main() {

	fmt.Println("inside main")

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
		return
	}

	fileName := filepath.Join(usr.HomeDir, configFile)

	configs, err := readUserConfig(fileName)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("reached here", configs)

}

/*
func main() {

	validateArgs()

	fmt.Println("WTF 1")

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	fileName := filepath.Join(usr.HomeDir, ".todos")

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: go run your_script.go <string>")
		return
	}

	if len(args) > 1 {

		if os.Args[1] == "add" {
			fmt.Println("Adding a todo")

			if len(args) < 2 {
				fmt.Println("Usage: go run your_script.go <string>")
				return
			}

			inputString := args[2]
			fmt.Println("You entered:", inputString)

			err := addTodo(fileName, inputString)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("String appended to", fileName)
			sendNotification("Todo was added")
		}

		if os.Args[1] == "check" {
			fmt.Println("Checking todos")

			todos, err := getTodos(fileName)

			if err != nil {
				log.Fatal(err)
			}

			for _, todo := range todos {
				fmt.Println("-", todo)
			}
		}

		if os.Args[1] == "complete" {
			fmt.Println("Completing a todo ...")
		}

		if os.Args[1] == "incomplete" {
			fmt.Println("Incompleting a todo ...")
		}

		if os.Args[1] == "delete" {
			fmt.Println("Deleting todo ...")
		}
	}
}
*/
