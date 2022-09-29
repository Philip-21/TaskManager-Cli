package cmd

import "github.com/spf13/cobra"

//this  handles all tasks
var RootCmd = &cobra.Command{
	Use:   "task-cli",                                    //is the one-line usage message
	Short: "task-cli is a CLI task manager built in go ", // Short is the short description shown in the 'help' output
}
