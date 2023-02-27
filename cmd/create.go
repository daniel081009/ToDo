package cmd

import (
	"fmt"
	"time"
	"todo/db"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "add",
	Short: "get Todo List",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Println("Name not provided")
			return
		}
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
		db.AddToDo(db.Todo{
			Name:      name,
			Details:   details,
			Category:  category,
			Deadline:  deadtime,
			Create_at: time.Now(),
		})
		fmt.Println("Todo Created")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "", "Create Todo by name(Required)")
	createCmd.Flags().StringP("details", "d", "", "Create Todo by details(Required)")
	createCmd.Flags().StringP("category", "c", "", "Create Todo by category(Required)")
	createCmd.Flags().StringP("deadline", "l", "", "Create Todo by category(Required)")
}
