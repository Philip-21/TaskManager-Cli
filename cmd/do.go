package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// marks and acts on the items in our task list as an integer
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		//convert the args to int
		var ids []int //using a []int for doing multiple integers all at once
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to Parse the item:", arg)
			} else {
				//if its a valid int add it to our ids list
				ids = append(ids, id)
			}
		}
		fmt.Println(ids)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
