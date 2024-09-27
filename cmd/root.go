package cmd

import (
	"fmt"
	"os"
)

func init() {
	rootCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Quiet mode (Only show errors)")
	rootCmd.AddCommand(versionCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
