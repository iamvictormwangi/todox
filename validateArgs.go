package main

import (
	"fmt"
	"log"
	"os"
)

func validateArgs() {

	fmt.Println("Validating")

	for _, arg := range os.Args {
		if arg == "--note" {
			if arg == "--todo" {
				log.Fatal("You cannot pass --note and --todo in the same prompt")
				showErrorMessage()
				return
			}
		}

		if arg == "--todo" {
			if arg == "--note" {
				log.Fatal("You cannot pass --note and --todo in the same prompt")
				showErrorMessage()
				return
			}
		}
	}

	fmt.Println("WTF")

	if len(os.Args) < 2 {
		log.Fatal("Too short")
		return
	}

}
