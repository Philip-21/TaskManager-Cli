package main

import (
	"fmt"
	"os"
	"path/filepath"
	"task-cli/cmd"
	"task-cli/database"

	homedir "github.com/mitchellh/go-homedir" //detects the user's home directory without the use of cgo, so the library can be used in cross-compilation environments
)

func main() {
	home, _ := homedir.Dir()
	dbpath := filepath.Join(home, "task.db")

	database.ConnectDB(dbpath)
	displayError(cmd.RootCmd.Execute())
}

func displayError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
