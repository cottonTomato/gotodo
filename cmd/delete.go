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

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a todo from the list",
	Long: `Delete a todo from the list.
Command should be used as:
    gotodo delete <task_id>`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db, err := datastore.InitDb()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		db.Exec("DELETE FROM tasks WHERE id = ?", args[0])

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Task: ", args[0], " deleted")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
