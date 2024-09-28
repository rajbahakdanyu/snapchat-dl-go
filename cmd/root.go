package cmd

import (
	"fmt"
	"os"
)

func init() {
	rootCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Do not print anything except errors to the console.")
	rootCmd.Flags().Uint16VarP(&maxStoryNum, "limit-story", "l", 0, "Set maximum number of stories to download.")
	rootCmd.Flags().Uint16VarP(&sleepInterval, "sleep-interval", "n", 1, "Sleep between downloads in seconds. (Default: 1s)")
	rootCmd.AddCommand(versionCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
