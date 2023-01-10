package cmd

import "github.com/spf13/cobra"

//this  handles all tasks
var RootCmd = &cobra.Command{
	Use:   "task-cli",                                    //is the one-line usage message
	Short: "task-cli is a CLI task manager built in go ", // Short is the short description shown in the 'help' output
}

//init() function sets off a piece of code to run before any other
// part of your package. This code will execute as soon as the
// package is imported, and can be used when you need your
//application to initialize in a specific state,
// such as when you have a specific configuration or set of
//resources with which your application needs to start

func init() {
	RootCmd.AddCommand(addCmd)
	RootCmd.AddCommand(doCmd)
	RootCmd.AddCommand(listCmd)
}
