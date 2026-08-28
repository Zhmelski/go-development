package fileutils

import (
	"fmt"
	"os"
)

// 'ReadFile' reads the contents of a file and returns it as a string
func ReadFile(filePath string) (string, error) {
	// Checking the existence of a file
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", fmt.Errorf("file does not exist: %s", filePath)
	}

	// Reading the contents of the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	return string(content), nil
}
