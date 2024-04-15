package main

import (
	"os"
	"strings"
)

func getTodos(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var todos []string
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if line != "" {
			todos = append(todos, line)
		}
	}
	return todos, nil
}
