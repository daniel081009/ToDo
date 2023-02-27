package cmd

import (
	"fmt"
	"strings"
	"todo/db"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get Todo List",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := db.GetToDo()
		if err != nil {
			fmt.Println(err)
			return
		}

		id, _ := cmd.Flags().GetString("id")
		if id != "" {
			data = map[string]db.Todo{id: data[id]}
		}

		name, _ := cmd.Flags().GetString("name")
		if name != "" {
			for key, value := range data {
				if !strings.Contains(value.Name, name) {
					delete(data, key)
				}
			}
		}
		detail, _ := cmd.Flags().GetString("detail")
		if detail != "" {
			for key, value := range data {
				if !strings.Contains(value.Name, detail) {
					delete(data, key)
				}
			}
		}
		category, _ := cmd.Flags().GetString("category")
		if category != "" {
			for key, value := range data {
				if !strings.Contains(value.Name, category) {
					delete(data, key)
				}
			}
		}
		all, _ := cmd.Flags().GetBool("all")
		if id != "" {
			fmt.Printf("ID: %s\nName: %s\nDetails: %s\nCategory: %s\nDeadline: %s\nCreate_at: %s\n", id, data[id].Name, data[id].Details, data[id].Category, data[id].Deadline.Format("2006-01-02"), data[id].Create_at.Format("2006-01-02"))
			return
		} else if name == "" && detail == "" && category == "" && !all {
			for key, value := range data {
				if value.Clear {
					delete(data, key)
				}
			}
		}

		db.Print(data)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringP("id", "i", "", "Get Todo by ID")
	getCmd.Flags().StringP("name", "n", "", "Get Todo by Name")
	getCmd.Flags().StringP("detail", "d", "", "Get Todo by Detail")
	getCmd.Flags().StringP("category", "c", "", "Get Todo by Category")
	getCmd.Flags().BoolP("all", "a", false, "all Todo Print")
}
