package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List called")
	},
}
