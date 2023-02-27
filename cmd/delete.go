package cmd

import (
	"fmt"
	"todo/db"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "TODO Clear",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("ID not provided")
			return
		}
		for _, todoID := range args {
			db.DeleteTodo(todoID)
			fmt.Printf("Todo %s Delete\n", todoID)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
