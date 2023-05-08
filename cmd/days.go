/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// daysCmd represents the days command
var daysCmd = &cobra.Command{
	Use:   "days",
	Short: "order files by day",
	Long: `creates folder for each date up to the day and sorts files inside.
Note: only matches filenames between 2000 - 2050`,
	Run: func(cmd *cobra.Command, args []string) {
		organizeByDate(Days, FileName)
	},
}
// daysModTimeCmd represents the days subcommand to order by modify time
var daysModTimeCmd = &cobra.Command{
	Use:   "modtime",
	Short: "time source is modified time",
	Long:  `sortes by day and uses the last modified time of the operating system`,
	Run: func(cmd *cobra.Command, args []string) {
		organizeByDate(Days, ModifiedTime)
	},
}

func init() {
	rootCmd.AddCommand(daysCmd)
	daysCmd.AddCommand(daysModTimeCmd)
}
