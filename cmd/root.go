/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ranger",
	Short: "Ranger organizes files in a directory",
	Long:
`_ __ __ _ _ __   __ _  ___ _ __
| '__/ _  | '_ \ / _  |/ _ \ '__|
| | | (_| | | | | (_| |  __/ |
|_|  \__,_|_| |_|\__, |\___|_|
                  __/ |
                 |___/
Ranger is a commandline tool to organize files in a directory into subdirectories.
This can happen using different attributes and in varying levels of strictness.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


