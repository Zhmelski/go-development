package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 'PersonData' contains data from the persons.txt file
type PersonData struct {
	FirstName string
	LastName  string
	LineNum   int
}

// 'ParsePersonsFile' reads and parses the persons.txt file
func ParsePersonsFile(filename string) ([]PersonData, []string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("file opening error: %w", err)
	}
	defer file.Close()

	var persons []PersonData
	var errors []string

	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())

		// Skipping empty lines
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) != 2 {
			errors = append(errors, fmt.Sprintf("Line %d: Invalid format - '%s'", lineNum, line))
			continue
		}

		persons = append(persons, PersonData{
			FirstName: parts[0],
			LastName:  parts[1],
			LineNum:   lineNum,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("file reading error: %w", err)
	}

	return persons, errors, nil
}

// 'WriteErrors' writes errors to the errors.txt file
func WriteErrors(errors []string) error {
	if len(errors) == 0 {
		return nil
	}

	file, err := os.Create("errors.txt")
	if err != nil {
		return fmt.Errorf("error creating errors.txt file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, errMsg := range errors {
		_, err := writer.WriteString(errMsg + "\n")
		if err != nil {
			return fmt.Errorf("error writing to errors.txt: %w", err)
		}
	}

	return writer.Flush()
}
