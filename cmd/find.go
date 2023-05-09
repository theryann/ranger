/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "find files with certain attributes",
	Long: `Find files in the current and child directories that have certain attributes`,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd() // get CWD
		find(cwd, EXIFLoc)
	},
}

func init() {
	rootCmd.AddCommand(findCmd)

}
