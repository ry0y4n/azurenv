/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// variable for '--file' flag
var envFile string

// envCmd represents the env command
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Manage local .env file",
	Long: "Subcommand for reading or handling local environment files.",
}

var listCmd = &cobra.Command{
	Use: "list",
	Short: "List environment variables from a .env file",
	Long: `Read a .env file and display its key-value pairs.
You can specify the path with --file (default is ".env").`,
	Run: func(cmd *cobra.Command, args []string) {
		keys, envMap, err := loadEnvFile(envFile)
		if err != nil {
			fmt.Println("[ERROR]", err)
			os.Exit(1)
		}

		// Display key-value pairs
		fmt.Printf("Loaded environment variables from %s:\n", envFile)
		for _, k := range keys {
			fmt.Printf("%s=%s\n", k, envMap[k])
		}
	},
}

func init() {
	rootCmd.AddCommand(envCmd)
	envCmd.AddCommand(listCmd)

	listCmd.Flags().StringVarP(&envFile, "file", "f", ".env", "Path to the .env file")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// envCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// envCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func loadEnvFile(filename string) ([]string, map[string]string, error) {
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
			value = value[1: len(value)-1]
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
