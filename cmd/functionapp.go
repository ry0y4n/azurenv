/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	functionappName string
	functionAppResourceGroup string
)

// functionappCmd represents the functionapp command
var functionappCmd = &cobra.Command{
	Use:   "functionapp",
	Short: "Manage Azure Azure Functions app settings",
	Long: "List and update app settings for Azure Functions",
}

var listRemoteFunctionCmd  = &cobra.Command{
	Use: "list-remote",
	Short: "List remote settings from an Azure Functions",
	Run: func(cmd *cobra.Command, args []string) {
		azCmd := exec.Command("az", "functionapp", "config", "appsettings", "list",
			"--name", functionappName,
			"--resource-group", functionAppResourceGroup,
		)

		out, err := azCmd.Output()
		if err != nil {
			log.Fatalf("[ERROR] Failed to run 'az functionapp config appsettings list: %v\n", err)
		}

		// Parse JSON output
		var settings []AppSettings
		if err := json.Unmarshal(out, &settings); err != nil {
			log.Fatalf("[ERROR] Failed to parse JSON: %v\n", err)
		}

		// Display key-value pairs
		fmt.Printf("App Settings for Azure Functions '%s':\n", functionappName)
		for _, s := range settings {
			fmt.Printf("- %s = %s\n", s.Name, s.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(functionappCmd)
	functionappCmd.AddCommand(listRemoteFunctionCmd )

	// Flags for list-remote command
	listRemoteFunctionCmd .Flags().StringVarP(&functionappName, "name", "n", "", "Name of the Azure App Service")
	listRemoteFunctionCmd .Flags().StringVarP(&functionAppResourceGroup, "resource-group", "g", "", "Resource group of the Azure App Service")

	listRemoteFunctionCmd .MarkFlagRequired("name")
	listRemoteFunctionCmd .MarkFlagRequired("resource-group")
}
