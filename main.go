package main

import (
	"task-cli/cmd"
	"task-cli/database"
)

func main() {
	cmd.RootCmd.Execute()
	database.ConnectDB()
}
