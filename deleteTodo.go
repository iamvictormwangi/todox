package main

import (
	"bufio"
	"os"
	"strings"
)

func deleteTodo(fileName string, todoId string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	tempFilePath := fileName + ".tmp"
	tempFile, err := os.Create(tempFilePath)
	if err != nil {
		return err
	}
	defer tempFile.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, todoId) {
			_, err := tempFile.WriteString(line + "\n")
			if err != nil {
				return err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if err := os.Rename(tempFilePath, fileName); err != nil {
		return err
	}

	return nil
}
