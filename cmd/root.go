package cmd

import (
	"fmt"
	"todo/db"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "td",
	Short: "Create,Read,Edit,Delete CLI application for Todo",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := db.GetToDo()
		if err != nil {
			fmt.Println(err)
			return
		}
		for key, value := range data {
			if value.Clear {
				delete(data, key)
			}
		}
		db.Print(data)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
