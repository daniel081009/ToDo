package cmd

import (
	"fmt"
	"todo/db"

	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "TODO Clear",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("ID not provided")
			return
		}
		todoID := args[0]
		if db.TodoClear(todoID) {
			fmt.Println("Todo Clear!")
		} else {
			fmt.Println("Todo not found")
		}
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
