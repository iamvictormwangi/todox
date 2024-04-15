package main

import (
	"os"
	"strconv"
)

func addTodo(fileName string, data string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	var random = strconv.Itoa(generateRandomNumber())

	if _, err := file.WriteString("[" + random + "]" + "[i]" + data + "\n"); err != nil {
		return err
	}
	return nil
}
