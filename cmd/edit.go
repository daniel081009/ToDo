package cmd

import (
	"fmt"
	"time"
	"todo/db"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "add",
	Short: "get Todo List",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			fmt.Println("ID not provided")
			return
		}
		name, _ := cmd.Flags().GetString("name")
		details, _ := cmd.Flags().GetString("details")
		category, _ := cmd.Flags().GetString("category")
		deadline, _ := cmd.Flags().GetString("deadline")

		deadtime := time.Now().Add(time.Hour * 24 * 7)

		if deadline != "" {
			deadtime, err = time.Parse("2006-01-02", deadline)
			if err != nil {
				fmt.Println("Deadline not provided")
				return
			}
		}
		db.UpdateTodo(id, db.Todo{
			Name:      name,
			Details:   details,
			Category:  category,
			Deadline:  deadtime,
			Create_at: time.Now(),
		})
		fmt.Println("Todo Editd")
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().StringP("id", "i", "", "Edit Todo by id(Required)")
	editCmd.Flags().StringP("name", "n", "", "Edit Todo by name")
	editCmd.Flags().StringP("details", "d", "", "Edit Todo by details")
	editCmd.Flags().StringP("category", "c", "", "Edit Todo by category")
	editCmd.Flags().StringP("deadline", "l", "", "Edit Todo by category")
}
