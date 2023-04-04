package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean up the project",
	Long:  "Clean up the project",
	Run: func(cmd *cobra.Command, args []string) {
		os.Remove("ark.db")
	},
}
