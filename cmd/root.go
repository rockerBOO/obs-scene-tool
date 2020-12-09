package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "obs-scene-tool",
	Short: "Tool to help manage scenes in OBS",
	Long:  `Tool to help manage scenes in OBS | https://github.com/rockerBOO/obs-scene-tool`,
	Run: func(cmd *cobra.Command, args []string) {
		// do stuff here
		fmt.Println("Congratulations")
	},
}

var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge 2 scenes json files",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Merge not working yet :)")
	},
}

func init() {
	rootCmd.AddCommand(mergeCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
