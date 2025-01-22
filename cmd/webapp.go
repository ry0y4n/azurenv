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
	filePathForWebapp string
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
		for i, s := range settings {
			fmt.Printf("[%d] %s = %s\n", i+1, s.Name, s.Value)
		}
	},
}

var applyWebappCmd = &cobra.Command{
	Use: "apply",
	Short: "Apply local .env settings to an Azure App Service",
	Long: "Read a local .env file and apply all key-value pairs to the specified Azure App Service",
	Run: func(cmd *cobra.Command, args []string) {
		// Load local .env file
		keys, envMap, err := loadEnvFile(filePathForWebapp)
		if err != nil {
			log.Fatalf("[ERROR] Failed to load %s: %v\n", filePathForWebapp, err)
		}
		
		if len(keys) == 0 {
			log.Println("[WARN] No environment variables found in", filePathForWebapp)
			return
		}

		// Prepare arguments
		var settingsArgs []string
		for _, k := range keys {
			v := envMap[k]
			if k == "" {
				continue
			}
			arg := fmt.Sprintf("%s=%s", k, v)
			settingsArgs = append(settingsArgs, arg)
		}

		// Build the Azure CLI command
		azArgs := []string{
			"webapp", "config", "appsettings", "set",
			"--name", webappName,
			"--resource-group", webappResourceGroup,
			"--settings",
		}
		azArgs = append(azArgs, settingsArgs...)

		cmdAz := exec.Command("az", azArgs...)
		out, err := cmdAz.Output()
		if err != nil {
			log.Fatalf("[ERROR] Failed to run az command:\n%v\nOutput:%s", err, string(out))
		}

		// Parse JSON output
		var updated []AppSettings
		if err := json.Unmarshal(out, &updated); err != nil {
			log.Fatalf("[ERROR] Failed to parse JSON output: %v\nOutput: %s", err, string(out))
		}

		// Display success message
		fmt.Printf("Successfully applied %d keys to App Service '%s':\n", len(settingsArgs), webappName)
		for i, pair := range settingsArgs {
			fmt.Printf("[%d] %s\n", i+1, pair)
		}
 	},
}

func init() {
	rootCmd.AddCommand(webappCmd)
	webappCmd.AddCommand(listRemoteWebappCmd)
	webappCmd.AddCommand(applyWebappCmd)

	// Flags for list-remote command
	listRemoteWebappCmd.Flags().StringVarP(&webappName, "name", "n", "", "Name of the Azure App Service")
	listRemoteWebappCmd.Flags().StringVarP(&webappResourceGroup, "resource-group", "g", "", "Resource group of the Azure App Service")
	listRemoteWebappCmd.MarkFlagRequired("name")
	listRemoteWebappCmd.MarkFlagRequired("resource-group")

	// Flags for apply command
	applyWebappCmd.Flags().StringVarP(&webappName, "name", "n", "", "Namew of the Azure App Service")
    applyWebappCmd.Flags().StringVarP(&webappResourceGroup, "resource-group", "g", "", "Resource group of the Azure App Service")
    applyWebappCmd.Flags().StringVarP(&filePathForWebapp, "file", "f", ".env", "Path to the .env file")
    applyWebappCmd.MarkFlagRequired("name")
    applyWebappCmd.MarkFlagRequired("resource-group")
}
