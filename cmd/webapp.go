/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"azurenv/internal/appconfig"
	"azurenv/internal/services"

	"github.com/spf13/cobra"
)

var webAppConfig appconfig.AppConfig

// webappCmd represents the webapp command
var webappCmd = &cobra.Command{
	Use:   "webapp",
	Short: "Manage Azure App Service app settings",
	Long:  "List and update app settings for Azure App Service",
}

var listRemoteWebappCmd = &cobra.Command{
	Use:   "list-remote",
	Short: "List remote settings from an Azure App Service",
	Run: func(cmd *cobra.Command, args []string) {
		services.ListRemote(webAppConfig)
	},
}

var applyWebappCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply local .env settings to an Azure App Service",
	Long:  "Read a local .env file and apply all key-value pairs to the specified Azure App Service",
	Run: func(cmd *cobra.Command, args []string) {
		services.Apply(webAppConfig)
	},
}

var diffWebappCmd = &cobra.Command{
	Use:   "diff",
	Short: "Show differences between local .env and Azure App Service app settings",
	Run: func(cmd *cobra.Command, args []string) {
		services.ShowDiff(webAppConfig)
	},
}

func init() {
	rootCmd.AddCommand(webappCmd)
	webappCmd.AddCommand(listRemoteWebappCmd)
	webappCmd.AddCommand(applyWebappCmd)
	webappCmd.AddCommand(diffWebappCmd)

	// Flags for list-remote command
	listRemoteWebappCmd.Flags().StringVarP(&webAppConfig.AppName, "name", "n", "", "Name of the Azure App Service")
	listRemoteWebappCmd.Flags().StringVarP(&webAppConfig.ResourceGroup, "resource-group", "g", "", "Resource group of the Azure App Service")
	listRemoteWebappCmd.MarkFlagRequired("name")
	listRemoteWebappCmd.MarkFlagRequired("resource-group")

	// Flags for apply command
	applyWebappCmd.Flags().StringVarP(&webAppConfig.AppName, "name", "n", "", "Name of the Azure App Service")
	applyWebappCmd.Flags().StringVarP(&webAppConfig.ResourceGroup, "resource-group", "g", "", "Resource group of the Azure App Service")
	applyWebappCmd.Flags().StringVarP(&webAppConfig.FilePath, "file", "f", ".env", "Path to the .env file")
	applyWebappCmd.MarkFlagRequired("name")
	applyWebappCmd.MarkFlagRequired("resource-group")

	// Flags for diff command
	diffWebappCmd.Flags().StringVarP(&webAppConfig.AppName, "name", "n", "", "Name of the Azure App Service")
	diffWebappCmd.Flags().StringVarP(&webAppConfig.ResourceGroup, "resource-group", "g", "", "Resource group of the Azure App Service")
	diffWebappCmd.Flags().StringVarP(&webAppConfig.FilePath, "file", "f", ".env", "Path to the .env file")
	diffWebappCmd.MarkFlagRequired("name")
	diffWebappCmd.MarkFlagRequired("resource-group")
}
