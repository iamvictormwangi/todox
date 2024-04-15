package main

import "fmt"

func showErrorMessage() {
	fmt.Println("Commands")

	fmt.Println("  add            Create a new todo")
	fmt.Println("  check          Check todos")
	fmt.Println("  complete       Complete todo")
	fmt.Println("  incomplete     Incomplete todo")
	fmt.Println("  delete         Delete todo")

	fmt.Println("Options")

	fmt.Println("  --note         Manipulate note instead")
}
