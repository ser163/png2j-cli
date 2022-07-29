/*
Copyright © 2022 Harry Liu <L3478830@163.com>

*/
package cmd

import (
	"fmt"
	"github.com/ser163/png2j-cli/lib"
	"github.com/spf13/cobra"
	"os"
)

var src string
var dst string
var width uint
var height uint
var quality uint8

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "png2j-cli",
	Short: "A simple command-line tool for converting PNG files to JPG files.",
	Long:  `A simple command-line tool for converting PNG files to JPG files.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if (src == "") || (dst == "") {
			fmt.Println("require source and output file")
			os.Exit(1)
		}
		if width == 0 && height == 0 {
			// default width and height
			err := lib.ConvertToOrgJpg(src, dst, quality)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			// width and height
			if width < 1 || height < 1 {
				fmt.Println("width and height must be greater than 0")
			}
			// resize
			err := lib.ConvertToJpg(src, dst, width, height, quality)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&src, "source", "s", "", "source PNG file")
	rootCmd.PersistentFlags().StringVarP(&dst, "output", "o", "", "output JPG file")
	rootCmd.PersistentFlags().UintVarP(&width, "width", "w", 0, "width of output JPG file")
	rootCmd.PersistentFlags().UintVarP(&height, "height", "e", 0, "height of output JPG file")
	rootCmd.PersistentFlags().Uint8VarP(&quality, "quality", "q", 100, "quality value: 1-100")
}
