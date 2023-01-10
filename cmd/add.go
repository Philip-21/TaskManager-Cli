package cmd

import (
	"fmt"
	"strings"
	"task-cli/database"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",                           //is the one-line usage message
	Short: "adds a task to your task list", // Short is the short description shown in the 'help' output
	Run: func(cmd *cobra.Command, args []string) {
		//takes in a slice and allows to merge the slice with whatever character we want
		task := strings.Join(args, "")
		_, err := database.Create(task)
		if err != nil {
			fmt.Println("Error in creating task:", err.Error())
			return
		}
		fmt.Printf("Added \"%s\" to your task list.\n", task) // \n opens a new line at the end
		//log.Println("Added \"%s\" to db ")
	},
}
