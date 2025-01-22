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
	filePathForFunctionapp string
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
		for i, s := range settings {
			fmt.Printf("[%d] %s = %s\n", i+1, s.Name, s.Value)
		}
	},
}

var applyFunctionappCmd = &cobra.Command{
	Use: "apply",
	Short: "Apply local .env settings to an Azure Functions",
	Long: "Read a local .env file and apply all key-value pairs to the specified Azure Functions",
	Run: func(cmd *cobra.Command, args []string) {
		// Load local .env file
		keys, envMap, err := loadEnvFile(filePathForFunctionapp)
		if err != nil {
			log.Fatalf("[ERROR] Failed to load %s: %v\n", filePathForFunctionapp, err)
		}
		
		if len(keys) == 0 {
			log.Println("[WARN] No environment variables found in", filePathForFunctionapp)
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
			"functionapp", "config", "appsettings", "set",
			"--name", functionappName,
			"--resource-group", functionAppResourceGroup,
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
		fmt.Printf("Successfully applied %d keys to App Functions '%s':\n", len(settingsArgs), functionappName)
		for i, pair := range settingsArgs {
			fmt.Printf("[%d] %s\n", i+1, pair)
		}
 	},
}


func init() {
	rootCmd.AddCommand(functionappCmd)
	functionappCmd.AddCommand(listRemoteFunctionCmd)
	functionappCmd.AddCommand(applyFunctionappCmd)

	// Flags for list-remote command
	listRemoteFunctionCmd .Flags().StringVarP(&functionappName, "name", "n", "", "Name of the Azure App Functions")
	listRemoteFunctionCmd .Flags().StringVarP(&functionAppResourceGroup, "resource-group", "g", "", "Resource group of the Azure App Service")
	listRemoteFunctionCmd .MarkFlagRequired("name")
	listRemoteFunctionCmd .MarkFlagRequired("resource-group")

	// Flags for apply command
	applyFunctionappCmd.Flags().StringVarP(&functionappName, "name", "n", "", "Name of the Azure Functions")
	applyFunctionappCmd.Flags().StringVarP(&functionAppResourceGroup, "resource-group", "g", "", "Resource group of the Azure Functions")
	applyFunctionappCmd.Flags().StringVarP(&filePathForFunctionapp, "file", "f", ".env", "Path to the .env file")
	applyFunctionappCmd.MarkFlagRequired("name")
	applyFunctionappCmd.MarkFlagRequired("resource-group")
}
