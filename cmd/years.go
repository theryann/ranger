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
	Short: "directorify files by year",
	Long: `creates folder for each date up to the year and sorts files inside.
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
	rootCmd.AddCommand(yearsCmd)
	yearsCmd.PersistentFlags().BoolP("modtime",  "m" , false, "time source is modified time")
	yearsCmd.PersistentFlags().BoolP("exif",     "e" , false, "time source is exif data")
	yearsCmd.PersistentFlags().BoolP("filename", "f" , true,  "time source is filename")
}
