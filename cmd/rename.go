/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "directorify files by renaming them with a number",
	Long: `The files are renamed to a number that reflects their order.
This can be done using EXIF or modified time`,
	Run: func(cmd *cobra.Command, args []string) {
		modTime, _  := cmd.Flags().GetBool("modtime")
		exifTime, _ := cmd.Flags().GetBool("exif")

		if modTime {
			rename(ModifiedTime)
		} else if exifTime {
			rename(EXIF)
		} else {
			rename(EXIF)
		}
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)
	renameCmd.PersistentFlags().BoolP("exif",     "e" , true, "time source is exif data")
	renameCmd.PersistentFlags().BoolP("modtime",  "m" , false,  "time source is filename")


}
