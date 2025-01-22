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
	webappName string
	webappResourceGroup string
)

// webappCmd represents the webapp command
var webappCmd = &cobra.Command{
	Use:   "webapp",
	Short: "Manage Azure App Service app settings",
	Long: "List and update app settings for Azure App Service",
}

var listRemoteWebappCmd = &cobra.Command{
	Use: "list-remote",
	Short: "List remote settings from an Azure App Service",
	Run: func(cmd *cobra.Command, args []string) {
		azCmd := exec.Command("az", "webapp", "config", "appsettings", "list",
			"--name", webappName,
			"--resource-group", webappResourceGroup,
		)

		out, err := azCmd.Output()
		if err != nil {
			log.Fatalf("[ERROR] Failed to run 'az webapp config appsettings list: %v\n", err)
		}

		// Parse JSON output
		var settings []AppSettings
		if err := json.Unmarshal(out, &settings); err != nil {
			log.Fatalf("[ERROR] Failed to parse JSON: %v\n", err)
		}

		// Display key-value pairs
		fmt.Printf("App Settings for App Service '%s':\n", webappName)
		for _, s := range settings {
			fmt.Printf("- %s = %s\n", s.Name, s.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(webappCmd)
	webappCmd.AddCommand(listRemoteWebappCmd)

	// Flags for list-remote command
	listRemoteWebappCmd.Flags().StringVarP(&webappName, "name", "n", "", "Name of the Azure App Service")
	listRemoteWebappCmd.Flags().StringVarP(&webappResourceGroup, "resource-group", "g", "", "Resource group of the Azure App Service")

	listRemoteWebappCmd.MarkFlagRequired("name")
	listRemoteWebappCmd.MarkFlagRequired("resource-group")
}
