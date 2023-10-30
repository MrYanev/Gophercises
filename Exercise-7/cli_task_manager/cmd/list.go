package cmd

import (
	"fmt"
	"os"

	"github.com/MrYanev/Gophercises/Exercise-7/cli_task_manager/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong")
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks left")
		}

		fmt.Println("You have the following tasks")
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
