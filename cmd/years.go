/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	// "fmt"

	"github.com/spf13/cobra"
)

// yearsCmd represents the years command
var yearsCmd = &cobra.Command{
	Use:   "years",
	Short: "order files by year",
	Long: `creates folder for each found year and sorts files inside`,
	Run: func(cmd *cobra.Command, args []string) {
		organize(Years)
	},
}

func init() {
	rootCmd.AddCommand(yearsCmd)
}
