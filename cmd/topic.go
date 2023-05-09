/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// topicCmd represents the topic command
var topicCmd = &cobra.Command{
	Use:   "topic",
	Short: "directorify files by topic (string in filename)",
	Long:  `specify a list of strings (topics). All files containing the given string (="topic") in the filename are stored in a directory respectively`,
	Run: func(cmd *cobra.Command, args []string) {
		organizeByTopic(args)
	},
}

func init() {
	rootCmd.AddCommand(topicCmd)

}
