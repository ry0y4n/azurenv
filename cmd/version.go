/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version of azurenv CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("azurenv CLI v0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
