/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// azcheckCmd represents the azcheck command
var azcheckCmd = &cobra.Command{
	Use:   "azcheck",
	Short: "Check if Azure CLI is installed and show its version",
	Long: "This command checks if the Azure CLI is available on the system, and prints out the version of Azure CLI if installed",
	Run: func(cmd *cobra.Command, args []string) {
		// Check if Azure CLI is installed
		checkCmd := exec.Command("az", "version")
		
		out, err := checkCmd.CombinedOutput()

		if err != nil {
			log.Printf("[ERROR] Failed to run 'az version': %v\n", err)
			log.Printf("Output: %s\n", string(out))
			return
		}

		fmt.Println("Azure CLI version info:")
		fmt.Println(string(out))
	},
}

var checkLoginCmd = &cobra.Command{
	Use: "account",
	Short: "Check if Azure CLI is logged in",
	Long: "This command checks if the Azure CLI is logged in, and prints out the account info if logged in",
	Run: func(cmd *cobra.Command, args []string) {
		// Check if Azure CLI is logged in
		checkLoginCmd := exec.Command("az", "account", "show")

		out, err := checkLoginCmd.CombinedOutput()

		if err != nil {
			log.Printf("[ERROR] Failed to run 'az account show': %v\n", err)
			log.Printf("Output: %s\n", string(out))
			return
		}

		fmt.Println("Azure CLI account info:")
		fmt.Println(string(out))
	},
}

func init() {
	rootCmd.AddCommand(azcheckCmd)
	azcheckCmd.AddCommand(checkLoginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// azcheckCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// azcheckCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
