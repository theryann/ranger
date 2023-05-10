/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

// compileCmd represents the compile command
var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "move all files in this directory",
	Long: `moves all files from subdirectories into root directory of this command`,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd() // get CWD
		compile(cwd, cwd)
	},
}

func init() {
	rootCmd.AddCommand(compileCmd)

}
