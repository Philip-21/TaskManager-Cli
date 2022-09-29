package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",                           //is the one-line usage message
	Short: "adds a task to your task list", // Short is the short description shown in the 'help' output
	Run: func(cmd *cobra.Command, args []string) {
		//takes in a slice and allows to merge the slice with whatever character we want
		task := strings.Join(args, "")
		fmt.Printf("Added \"%s\" to your task list.\n", task) // \n opens a new line at the end
	},
}

//init() function sets off a piece of code to run before any other
// part of your package. This code will execute as soon as the
// package is imported, and can be used when you need your
//application to initialize in a specific state,
// such as when you have a specific configuration or set of
//resources with which your application needs to start

func init() {
	RootCmd.AddCommand(addCmd)
}
