// Package api contains the methods to export data as API requests or
// by downloading entire tables
package api

import "github.com/cristian-sima/Wisply/models/database"

var sensitiveTableList = []string{"account", "account_token", "account_search", "api_table_setting"}

// GetAllWisplyTables returns the list of all the tables which MAY BE downloaded
func GetAllWisplyTables() []string {
	var list []string
	sql := "SELECT `table_name` FROM `information_schema`.`tables` WHERE `table_schema`='wisply'"
	rows, _ := database.Connection.Query(sql)
	for rows.Next() {
		table := ""
		rows.Scan(&table)
		if IsAllowedTable(table) {
			list = append(list, table)
		}
	}
	return list
}

// AreValidDetails checks if the table is allowed and
// the description is valid
func AreValidDetails(table Table) bool {
	return !IsAllowedTable(table.Name) && isValidDescription(table.Description)
}

// InsertNewTable adds the table name on the list of the tables
// which can be downloaded
func InsertNewTable(table Table) error {
	sql := "INSERT INTO `api_table_setting` (`name`, `description`) VALUES (?, ?)"
	query, err := database.Connection.Prepare(sql)
	query.Exec(table.Name, table.Description)
	return err
}

// RemoveAllowedTable removes the table from the list of allowed tables
func RemoveAllowedTable(ID int) error {
	sql := "DELETE FROM `api_table_setting` WHERE `id`=? "
	query, err := database.Connection.Prepare(sql)
	query.Exec(ID)
	return err
}

// GetAllowedTables returns the list of allowed tables
func GetAllowedTables() []Table {
	var list []Table
	sql := "SELECT `id`, `name`, `description` FROM `api_table_setting`"
	rows, _ := database.Connection.Query(sql)
	for rows.Next() {
		table := Table{}
		rows.Scan(&table.ID, &table.Name, &table.Description)
		list = append(list, table)
	}
	return list
}

// NewTable creates a new table by ID
func NewTable(ID string) (*Table, error) {
	table := &Table{}
	fieldList := "`id`, `name`, `description`"
	sql := "SELECT " + fieldList + " FROM `api_table_setting` WHERE id=? "
	query, err := database.Connection.Prepare(sql)
	if err != nil {
		return table, err
	}
	query.QueryRow(ID).Scan(&table.ID, &table.Name, &table.Description)
	return table, nil
}

// ModifyDetails changes the details of the table
func ModifyDetails(table *Table, newDescription string) error {
	sql := "UPDATE `api_table_setting` SET `description`=? WHERE `id`=? "
	query, err1 := database.Connection.Prepare(sql)
	if err1 != nil {
		return err1
	}
	_, err2 := query.Exec(newDescription, table.ID)
	return err2
}

// GetWisplyTablesNamesNotAllowed returns the list of wisply tables which
// can be downloaded, but are not yet on the list
func GetWisplyTablesNamesNotAllowed() []string {
	var (
		list          []string
		allWisply     = GetAllWisplyTables()
		allowedTables = GetAllowedTables()
	)

	for _, wisplyTable := range allWisply {
		exists := false
		for _, allowedTable := range allowedTables {
			if allowedTable.Name == wisplyTable {
				exists = true
			}
		}
		if !exists {
			list = append(list, wisplyTable)
		}
	}
	return list
}

// IsAllowedTable checks if a table name is allowed to be downloaded
func IsAllowedTable(name string) bool {
	for _, rejectedName := range sensitiveTableList {
		if name == rejectedName {
			return false
		}
	}
	return true
}
