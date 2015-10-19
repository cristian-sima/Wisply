package api

import "html/template"

// Table represents a wisply allowed table which can be downloaded
type Table struct {
	ID          int
	Name        string
	Description string
}

func (table *Table) GetDescription() template.HTML {
	return template.HTML([]byte(table.Description))
}
