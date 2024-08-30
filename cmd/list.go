/*
Copyright © 2024 Abhinav Pandey
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/cottonTomato/gotodo/datastore"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jmoiron/sqlx"
	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

var all *bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos.",
	Long: `List all todos.
To list all todos use --all or -a flag.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db, err := datastore.InitDb()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var rowsData *sqlx.Rows
		if *all {
			rowsData, err = db.Queryx("SELECT * FROM tasks")
		} else {
			rowsData, err = db.Queryx("SELECT * FROM tasks WHERE done = 1")
		}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		time_format := "2006-01-02 15:04:05"
		rowsDisplayData := []table.Row{}
		for rowsData.Next() {
			var row datastore.Task
			_ = rowsData.StructScan(&row) // TODO: log this

			time, _ := time.Parse(time_format, row.Created_at) // TODO: log this
			done := row.Done != 0

			rowsDisplayData = append(rowsDisplayData, []any{
				row.Id,
				row.Description,
				timediff.TimeDiff(time),
				done,
			})
		}

		out_table := table.NewWriter()
		out_table.SetOutputMirror(os.Stdout)
		out_table.AppendHeader(table.Row{"Task ID", "Description", "Created", "Status"})
		out_table.AppendSeparator()
		out_table.AppendRows(rowsDisplayData)
		out_table.Render()
	},
}

func init() {
	all = listCmd.Flags().BoolP("all", "a", false, "List all todos")
	rootCmd.AddCommand(listCmd)
}
