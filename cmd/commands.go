package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "0.1.0"

var rootCmd = &cobra.Command{
	Use:   "snapchat-dl",
	Short: "Snapchat-dl is a tool to download Snapchat videos",
	Long:  "Snapchat-dl is a tool to download Snapchat videos",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Snapchat-dl")
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Snapchat-dl",
	Long:  "All software has versions. This is Snapchat-dl's",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Snapchat-dl v" + version)
	},
}
