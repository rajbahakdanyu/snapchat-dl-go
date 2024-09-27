package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "snapchat-dl",
	Short: "Download Snapchat videos",
	Long:  "A tool to download Snapchat videos",
	Args:  cobra.MinimumNArgs(1),
	Run:   runRoot,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  "Display the version of Snapchat-dl",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Snapchat-dl v%s\n", version)
	},
}
