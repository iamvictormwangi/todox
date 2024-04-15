package main

import (
	"math/rand"
	"time"
)

func generateRandomNumber() int {
	rand.Seed(time.Now().UnixNano())

	// Generate a random 4-digit number
	randomNumber := rand.Intn(9000) + 1000

	return randomNumber
}
