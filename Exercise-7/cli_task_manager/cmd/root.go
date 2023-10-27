package cmd

import "github.com/spf13/cobra"

//By not adding a func to the variable it makes it list all other commands
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is CLI task mannager",
}
