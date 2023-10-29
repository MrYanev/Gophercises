package main

import (
	"path/filepath"

	"github.com/MrYanev/Gophercises/Exercise-7/cli_task_manager/cmd"
	"github.com/MrYanev/Gophercises/Exercise-7/cli_task_manager/db"
	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	err := db.Init(dbPath)
	if err != nil {
		panic(err)
	}
	cmd.RootCmd.Execute()
}
