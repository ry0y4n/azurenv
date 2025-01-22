package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func LoadEnvFile(filename string) ([]string, map[string]string, error) {
	// Check if file exists
	info, err := os.Stat(filename)
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil, fmt.Errorf("%s does not exist", filename)
	} else if err != nil {
		return nil, nil, fmt.Errorf("Could not stat %s: %w", filename, err)
	}

	// Check if file is a directory
	if info.IsDir() {
		return nil, nil, fmt.Errorf("%s is a directory, not a file", filename)
	}

	// Open file
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("Could not open %s: %w", filename, err)
	}
	defer file.Close()

	// key-value map for retain order of keys
	keys := make([]string, 0)
	envMap := make(map[string]string)

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)

		// Skip empty lines and comments
		if trimmedLine == "" || strings.HasPrefix(trimmedLine, "#") {
			continue
		}

		// Split line into key and value
		splitIndex := strings.Index(trimmedLine, "=")
		if splitIndex == -1 {
			fmt.Println("Warning: invalid line (no '='):", trimmedLine)
			continue
		}

		key := strings.TrimSpace(trimmedLine[:splitIndex])
		value := strings.TrimSpace(trimmedLine[splitIndex+1:])

		// Remove surrounding quotes if present (' or ")
		if len(value) > 1 && ((value[0] == '"' && value[len(value)-1] == '"') || (value[0] == '\'' && value[len(value)-1] == '\'')) {
			value = value[1 : len(value)-1]
		}

		// Add mapping
		envMap[key] = value
		keys = append(keys, key)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("Error while reading %s: %w", filename, err)
	}

	return keys, envMap, nil
}
