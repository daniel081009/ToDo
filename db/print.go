package db

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func Print(data map[string]Todo) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("ID", "Name", "Details", "Category", "Deadline", "Claer", "Create_at")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for key, data := range data {
		tbl.AddRow(key, data.Name, data.Details, data.Category, data.Deadline.Format("2006-01-02"), data.Clear, data.Create_at.Format("2006-01-02"))
	}

	tbl.Print()
}
