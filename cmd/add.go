/*
Copyright © 2024 Abhinav Pandey
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/cottonTomato/gotodo/datastore"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add entry to todo list",
	Long: `Add entry to todo list.
Add any number of sentences after this command.
Empty tasks are not allowed.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db, err := datastore.InitDb()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		_, err = db.Exec("INSERT INTO tasks (description) VALUES (?)", strings.Join(args[:], " "))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Task added!")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
