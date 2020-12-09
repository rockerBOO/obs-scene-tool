package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

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
