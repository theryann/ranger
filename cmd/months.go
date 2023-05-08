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
		organizeByDate(Months, FileName)
	},
}

// monthsModTimeCmd represents the months subcommand to order by modify time
var monthsModTimeCmd = &cobra.Command{
	Use:   "modtime",
	Short: "time source is modified time",
	Long:  `sortes by month and uses the last modified time of the operating system`,
	Run: func(cmd *cobra.Command, args []string) {
		organizeByDate(Months, ModifiedTime)
	},
}
func init() {
	rootCmd.AddCommand(monthsCmd)
	monthsCmd.AddCommand(monthsModTimeCmd)
}
