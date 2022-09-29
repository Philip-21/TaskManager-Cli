package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "done",
	Short: "marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("done")
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
