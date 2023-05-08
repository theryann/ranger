/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// yearsCmd represents the years command
var yearsCmd = &cobra.Command{
	Use:   "years",
	Short: "order files by year",
	Long: `creates folder for each date up to the year and sorts files inside.
Note: only matches filenames between 2000 - 2050`,
	Run: func(cmd *cobra.Command, args []string) {
		organizeByDate(Years, FileName)
	},
}

// yearsModTimeCmd represents the years subcommand to order by modify time
var yearsModTimeCmd = &cobra.Command{
	Use:   "modtime",
	Short: "time source is modified time",
	Long:  `sortes by year and uses the last modified time of the operating system`,
	Run: func(cmd *cobra.Command, args []string) {
		organizeByDate(Years, ModifiedTime)
	},
}

func init() {
	rootCmd.AddCommand(yearsCmd)
	yearsCmd.AddCommand(yearsModTimeCmd)
}
