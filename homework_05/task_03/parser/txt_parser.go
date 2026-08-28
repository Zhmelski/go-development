package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"homework_05/task_03/models"
)

// 'ParseDataFile' reads and parses a data file
func ParseTXT(fileName string) ([]models.Object, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %w", err)
	}
	defer file.Close()

	var objects []models.Object
	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())

		// Skipping empty lines
		if line == "" {
			continue
		}

		// Parse the string: name,x,y
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			return nil, fmt.Errorf("line %d: invalid format, expected 3 fields", lineNumber)
		}

		name := strings.TrimSpace(parts[0])

		// Parsing X
		var x float64
		n, err := fmt.Sscanf(strings.TrimSpace(parts[1]), "%f", &x)
		if err != nil || n != 1 {
			return nil, fmt.Errorf("line %d: invalid X coordinate: %w", lineNumber, err)
		}

		// Parsing Y
		var y float64
		n, err = fmt.Sscanf(strings.TrimSpace(parts[2]), "%f", &y)
		if err != nil || n != 1 {
			return nil, fmt.Errorf("line %d: invalid Y coordinate: %w", lineNumber, err)
		}

		objects = append(objects, models.Object{
			Name: name,
			X:    x,
			Y:    y,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error reading file: %w", err)
	}

	if len(objects) == 0 {
		return nil, fmt.Errorf("file is empty")
	}

	return objects, nil
}
