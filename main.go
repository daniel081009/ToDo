package main

import (
	"todo/cmd"
	"todo/db"
)

func main() {
	_, err := db.GetToDo()
	if err != nil {
		db.AddToDo(db.Todo{
			Name: "temp",
		})
	}
	cmd.Execute()
}
