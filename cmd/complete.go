/*
Copyright © 2024 Abhinav Pandey
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/cottonTomato/gotodo/datastore"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark todo as complete",
	Long: `Mark todo as complete.
Command should be used as:
    gotodo complete <task_id>`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db, err := datastore.InitDb()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		db.Exec("UPDATE tasks SET done = 1 WHERE id = ?", args[0])

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Task: ", args[0], " set to done")
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
