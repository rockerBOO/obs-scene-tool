package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of obs-scene-tool",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("obs-scene-tool 0.1-beginboy <- its fake.")
	},
}
