package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func readUserConfig(configFile string) (string, string, string, string, error) {
	// Open the config file
	file, err := os.Open(configFile)
	if err != nil {
		log.Fatal("Failed to open config file")
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Variable to store the username
	var username string
	var email string
	var password string
	var theme string

	// Iterate over each line in the file
	for scanner.Scan() {
		line := scanner.Text()

		// Check if the line contains "username"
		if strings.Contains(line, "username") {
			// Split the line by "=" to get the username value
			parts := strings.Split(line, "=")
			if len(parts) == 2 {
				// Trim leading and trailing whitespace, as well as quotes
				username = strings.TrimSpace(strings.Trim(parts[1], `" `))
				break // Stop scanning after finding the username
			}
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal("Error:", err)
	}

	return username, password, email, theme, nil
}
