package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is CLI task mannager",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}
