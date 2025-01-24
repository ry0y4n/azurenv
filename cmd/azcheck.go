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

type AccountInfo struct {
	EnvironmentName     string          `json:"environmentName"`
	HomeTenantId        string          `json:"homeTenantId"`
	ID                  string          `json:"id"` // Subscription ID
	IsDefault           bool            `json:"isDefault"`
	ManagedByTenants    []ManagedTenant `json:"managedByTenants"`
	Name                string          `json:"name"` // Subscription Name
	State               string          `json:"state"`
	TenantDefaultDomain string          `json:"tenantDefaultDomain"`
	TenantDisplayName   string          `json:"tenantDisplayName"`
	TenantID            string          `json:"tenantId"`
	User                AccountUser     `json:"user"`
}

type ManagedTenant struct {
	TenantID string `json:"tenantId"`
}

type AccountUser struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// azcheckCmd represents the azcheck command
var azcheckCmd = &cobra.Command{
	Use:   "azcheck",
	Short: "Check Azure account info",
	Long: `Check whether Azure CLI is installed and if you are logged in to Azure.
It shows the current account info (subscription, tenant, etc.). 
If you're not logged in, it suggests 'az login' and subscription set commands.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Check if Azure CLI is logged in
		checkLoginCmd := exec.Command("az", "account", "show")
		out, err := checkLoginCmd.CombinedOutput()

		if err != nil {
			log.Printf("[ERROR] Failed to run 'az account show': %v\n", err)
			log.Printf("Output: %s\n", string(out))

			fmt.Println("\nIt seems you are not logged in or an error occurred.")
			fmt.Println("Please run: az login")
			fmt.Println("Then set the desired subscription with:")
			fmt.Println("  az account set --subscription <SUBSCRIPTION_ID>")

			return nil
		}

		// Parse JSON output
		var account AccountInfo
		if parseErr := json.Unmarshal(out, &account); parseErr != nil {
			return fmt.Errorf("[ERROR] Failed to parse 'az account show' JSON: %w\nRaw Output:\n%s", parseErr, string(out))
		}

		// Display account info
		fmt.Println("You are logged in to Azure!")
		fmt.Printf("\nSubscription Name: %s\n", account.Name)
		fmt.Printf("Subscription ID:   %s\n", account.ID)
		fmt.Printf("Tenant Name:       %s\n", account.TenantDisplayName)
		fmt.Printf("Tenant ID:         %s\n", account.TenantID)
		fmt.Printf("User:              %s (%s)\n", account.User.Name, account.User.Type)
		fmt.Printf("State:             %s\n", account.State)

		fmt.Println("\nIf you want to switch subscription, run:")
		fmt.Println("  az account set --subscription <SUBSCRIPTION_ID>")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(azcheckCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// azcheckCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// azcheckCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
