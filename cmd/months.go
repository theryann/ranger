/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// monthsCmd represents the months command
var monthsCmd = &cobra.Command{
	Use:   "months",
	Short: "order files by months",
	Long: `creates folder for each each date up to the month and sorts files inside.
Note: only matches filenames between 2000 - 2050`,
	Run: func(cmd *cobra.Command, args []string) {
		organizeByDate(Months)
	},
}

func init() {
	rootCmd.AddCommand(monthsCmd)
}
