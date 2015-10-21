package api

import "html/template"

// Table represents a wisply allowed table which can be downloaded
type Table struct {
	ID          int
	Name        string
	Description string
}

// GetDescription returns the description of the table as HTML code
// rather than a safe string
func (table *Table) GetDescription() template.HTML {
	return template.HTML([]byte(table.Description))
}
