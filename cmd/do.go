package cmd

import (
	"fmt"
	"strconv"
	"task-cli/database"

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
			tasks, err := database.GetTasks()
			if err != nil {
				fmt.Println("Error in getting tasks :", err)
				return
			}
			for _, id := range ids {
				if id < 0 || id > len(tasks) {
					fmt.Println("invalid task number", id)
					continue
				}
				task := tasks[id-1]
				err := database.DeleteTask(task.Key)
				if err != nil {
					fmt.Printf("Failed to mark \"%d\" as completed. Error %s\n", id, err)
				} else {
					fmt.Printf("Marked \"%d\" as completed.\n", id)
				}
			}
		}
		//fmt.Println(ids)
	},
}
