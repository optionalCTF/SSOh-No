package service

import (
	"bufio"
	"fmt"
	"os"
)

// File read from https://stackoverflow.com/a/18479916
// Main purpose to be called for mass user enum/password spraying

func ReadFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	scanErr := scanner.Err()
	return lines, scanErr
}
