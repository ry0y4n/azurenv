/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"azurenv/internal/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// variable for '--file' flag
var envFile string

// envCmd represents the env command
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Manage local .env file",
	Long:  "Subcommand for reading or handling local environment files.",
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List environment variables from a .env file",
	Long: `Read a .env file and display its key-value pairs.
You can specify the path with --file (default is ".env").`,
	Run: func(cmd *cobra.Command, args []string) {
		keys, envMap, err := utils.LoadEnvFile(envFile)
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
}
