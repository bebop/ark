package cmd

import (
	"github.com/spf13/cobra"
)

var localCmd = &cobra.Command{
	Use:   "local",
	Short: "Builds a small test database for local development.",
	Run: func(cmd *cobra.Command, args []string) {
		// local()
	},
}
