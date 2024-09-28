package cmd

import (
	"fmt"
	"os"
)

func init() {
	rootCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Quiet mode (Only show errors)")
	rootCmd.Flags().Uint16VarP(&maxStoryNum, "limit-story", "l", 0, "Set maximum number of stories to download.")
	rootCmd.AddCommand(versionCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
