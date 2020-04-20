package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "status command get current playback status ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("status")
	},
}
