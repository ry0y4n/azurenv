/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "azurenv",
	Short: "Manage environment variables for Azure Appservice / Azure Functions",
	Long:  "azurenv is a CLI tool to manage environment variables both locally (.env) and on Azure Service / Azure Functions).",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("")
		fmt.Println(`🚀 Welcome to azurenv CLI! 🌍

_____  ______________ ________________________ ___________   ____
/  _  \ \____    /    |   \______   \_   _____/ \      \   \ /   /
/  /_\  \  /     /|    |   /|       _/|    __)_  /   |   \   Y   / 
/    |    \/     /_|    |  / |    |   \|        \/    |    \     /  
\____|__  /_______ \______/  |____|_  /_______  /\____|__  /\___/   
		\/        \/                \/        \/         \/         
		`)
		fmt.Println("For detailed instructions, please vist:")
		fmt.Println("  https://github.com/ry0y4n/azurenv")
		fmt.Println("\nYou can also run 'azurenv --help' to see available commands and options.")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
