/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"fmt"

	"github.com/spf13/cobra"
)

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "find files with certain attributes",
	Long: `Find files in the current and child directories that have certain attributes`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Usage()
		}
	},
}
var findExifLocCmd = &cobra.Command{
	Use:   "loc",
	Short: "find files with EXIF Location",
	Long: `Find files that have an exif location tag`,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd() // get CWD
		find(cwd, EXIFLoc, "") // "" is a string to search for. Not relevant when searching for exif data
	},
}

var findNameCmd = &cobra.Command{
	Use:   "name",
	Short: "find files with name",
	Long: `Find files that the provided name in their filename`,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd() // get CWD
		if len(args) > 0 {
			find(cwd, Name, args[0])
		} else {
			fmt.Println("please specify string to look for")
		}
	},
}

func init() {
	rootCmd.AddCommand(findCmd)
	findCmd.AddCommand(findExifLocCmd)
	findCmd.AddCommand(findNameCmd)

}
