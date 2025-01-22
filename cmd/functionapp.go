/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"azurenv/internal/appconfig"
	"azurenv/internal/services"

	"github.com/spf13/cobra"
)

var functionAppConfig appconfig.AppConfig

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
		services.ListRemote(functionAppConfig)
	},
}

var applyFunctionappCmd = &cobra.Command{
	Use: "apply",
	Short: "Apply local .env settings to an Azure Functions",
	Long: "Read a local .env file and apply all key-value pairs to the specified Azure Functions",
	Run: func(cmd *cobra.Command, args []string) {
		services.Apply(functionAppConfig)
 	},
}


func init() {
	rootCmd.AddCommand(functionappCmd)
	functionappCmd.AddCommand(listRemoteFunctionCmd)
	functionappCmd.AddCommand(applyFunctionappCmd)

	// Flags for list-remote command
	listRemoteFunctionCmd .Flags().StringVarP(&functionAppConfig.AppName, "name", "n", "", "Name of the Azure App Functions")
	listRemoteFunctionCmd .Flags().StringVarP(&functionAppConfig.ResourceGroup, "resource-group", "g", "", "Resource group of the Azure App Service")
	listRemoteFunctionCmd .MarkFlagRequired("name")
	listRemoteFunctionCmd .MarkFlagRequired("resource-group")

	// Flags for apply command
	applyFunctionappCmd.Flags().StringVarP(&functionAppConfig.AppName, "name", "n", "", "Name of the Azure Functions")
	applyFunctionappCmd.Flags().StringVarP(&functionAppConfig.ResourceGroup, "resource-group", "g", "", "Resource group of the Azure Functions")
	applyFunctionappCmd.Flags().StringVarP(&functionAppConfig.FilePath, "file", "f", ".env", "Path to the .env file")
	applyFunctionappCmd.MarkFlagRequired("name")
	applyFunctionappCmd.MarkFlagRequired("resource-group")
}
