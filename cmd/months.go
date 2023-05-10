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
	Short: "directorify files by months",
	Long: `creates folder for each each date up to the month and sorts files inside.
Note: only matches filenames between 2000 - 2050`,
	Run: func(cmd *cobra.Command, args []string) {
		modTime, _  := cmd.Flags().GetBool("modtime")
		exifTime, _ := cmd.Flags().GetBool("exif")
		fileName, _ := cmd.Flags().GetBool("filename")

		if modTime {
			organizeByDate(Months, ModifiedTime)
		} else if exifTime {
			organizeByDate(Months, EXIF)
		} else if fileName {
			organizeByDate(Months, FileName)
		} else {
			organizeByDate(Months, FileName)
		}

	},
}

func init() {
	rootCmd.AddCommand(monthsCmd)
	monthsCmd.Aliases = []string {"month"}
	monthsCmd.PersistentFlags().BoolP("modtime",  "m" , false, "time source is modified time")
	monthsCmd.PersistentFlags().BoolP("exif",     "e" , false, "time source is exif data")
	monthsCmd.PersistentFlags().BoolP("filename", "f" , true,  "time source is filename")
}
