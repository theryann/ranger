/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// typeCmd represents the type command
var typeCmd = &cobra.Command{
	Use:   "type",
	Short: "directify by media type",
	Long: `orders files in directories depending on the media type they are.
By adding extra flags you can specify i.e. videos only.
Yau can also chain types`,
	Run: func(cmd *cobra.Command, args []string) {
		media, _     := cmd.Flags().GetBool("media")
		image, _     := cmd.Flags().GetBool("image")
		video, _     := cmd.Flags().GetBool("video")
		audio, _     := cmd.Flags().GetBool("audio")
		extention, _ := cmd.Flags().GetBool("extention")

		if media { typify(All)
		} else if extention { typify(Extention)
		} else if image     { typify(Pictures)
		} else if audio     { typify(Music)
		} else if video     { typify(Videos) }

	},
}

func init() {
	rootCmd.AddCommand(typeCmd)
	typeCmd.PersistentFlags().BoolP("media",  "m" , false, "by all media formats")
	typeCmd.PersistentFlags().BoolP("image",  "i" , false, "image formates")
	typeCmd.PersistentFlags().BoolP("video",  "v" , false, "video formates")
	typeCmd.PersistentFlags().BoolP("audio",  "a" , false, "audio formates")
	typeCmd.PersistentFlags().BoolP("extention",  "e" , false, "by file extention")
}
