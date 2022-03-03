package service

import (
	"bufio"
	"errors"
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

func WriteFile(path string, contents string) error {
	if _, err := os.Stat(path); err == nil {
		f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Error: %s", err)
			fmt.Printf("ree1")
		}
		defer f.Close()
		if _, err := f.WriteString(contents + "\n"); err != nil {
			fmt.Printf("Error: %s", err)

		}
		return nil

	} else if errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(path)
		if err != nil {
			fmt.Printf("Error: %s", err)
			fmt.Printf("ree")
		}

		defer file.Close()
		w := bufio.NewWriter(file)
		fmt.Fprintln(w, contents)

		return w.Flush()
	} else {
		fmt.Printf("Schrodingers error... How did you do this?")
		return nil
	}

}
