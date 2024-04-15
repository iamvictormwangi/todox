package main

import (
	"fmt"
	"os/exec"
)

func sendNotification(notification string) {
	// Command to list files in the current directory
	cmd := exec.Command("notify-send", notification)

	// Run the command and capture its output
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the output
	fmt.Println(string(output))
}
