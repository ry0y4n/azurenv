/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version of azurenv CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("azurenv %s", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
