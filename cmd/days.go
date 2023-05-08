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
		modTime, _  := cmd.Flags().GetBool("modtime")
		exifTime, _ := cmd.Flags().GetBool("exif")
		fileName, _ := cmd.Flags().GetBool("filename")

		if modTime {
			organizeByDate(Years, ModifiedTime)
		} else if exifTime {
			organizeByDate(Years, EXIF)
		} else if fileName {
			organizeByDate(Years, FileName)
		} else {
			organizeByDate(Years, FileName)
		}

	},
}

func init() {
	rootCmd.AddCommand(daysCmd)
	daysCmd.PersistentFlags().BoolP("modtime",  "m" , false, "time source is modified time")
	daysCmd.PersistentFlags().BoolP("exif",     "e" , false, "time source is exif data")
	daysCmd.PersistentFlags().BoolP("filename", "f" , true,  "time source is filename")
}
