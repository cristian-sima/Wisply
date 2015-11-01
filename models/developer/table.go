package developer

import (
	"html/template"

	"github.com/cristian-sima/Wisply/models/database"
)

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

// RemoveAllowedTable removes the table from the list of allowed tables
func (table *Table) Delete() error {
	sql := "DELETE FROM `download_table` WHERE `id`=? "
	query, err := database.Connection.Prepare(sql)
	query.Exec(table.ID)
	return err
}

// NewTable creates a new table by ID
func NewTable(ID string) (*Table, error) {
	table := &Table{}
	fieldList := "`id`, `name`, `description`"
	sql := "SELECT " + fieldList + " FROM `download_table` WHERE id=? "
	query, err := database.Connection.Prepare(sql)
	if err != nil {
		return table, err
	}
	query.QueryRow(ID).Scan(&table.ID, &table.Name, &table.Description)
	return table, nil
}

// ModifyDetails changes the details of the table
func ModifyDetails(table *Table, newDescription string) error {
	sql := "UPDATE `download_table` SET `description`=? WHERE `id`=? "
	query, err1 := database.Connection.Prepare(sql)
	if err1 != nil {
		return err1
	}
	_, err2 := query.Exec(newDescription, table.ID)
	return err2
}
