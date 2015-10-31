// Package api contains the methods to export data as API requests or
// by downloading entire tables
package api

import "github.com/cristian-sima/Wisply/models/database"

var rejectedTables = []string{"account", "account_token", "account_search", "api_table_setting"}

// GetAllTables returns the entire list of tables which are not restricted
// from being downloaded
func GetAllTables() []string {
	var list []string
	sql := "SELECT `table_name` FROM `information_schema`.`tables` WHERE `table_schema`='wisply'"
	rows, _ := database.Connection.Query(sql)
	for rows.Next() {
		table := ""
		rows.Scan(&table)
		if !IsRestrictedTable(table) {
			list = append(list, table)
		}
	}
	return list
}

// IsTableAllowedToDownload checks if a table name exists and if
// it is not on the reject list
func IsTableAllowedToDownload(name string) bool {
	allowedTables := GetAllTables()
	// check it is not rejected
	for _, allowedTable := range allowedTables {
		if name == allowedTable {
			return true
		}
	}
	return false
}

// AreValidDetails checks if the table is allowed and
// the description is valid
func AreValidDetails(table Table) bool {
	return IsTableAllowedToDownload(table.Name) && isValidDescription(table.Description)
}

// InsertNewTable adds the table name on the list of the tables
// which can be downloaded
func InsertNewTable(table Table) error {
	sql := "INSERT INTO `api_table_setting` (`name`, `description`) VALUES (?, ?)"
	query, err := database.Connection.Prepare(sql)
	query.Exec(table.Name, table.Description)
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

// GetWisplyTablesNamesNotAllowed returns the list of wisply tables which
// can be downloaded, but are not yet on the list
func GetWisplyTablesNamesNotAllowed() []string {
	var (
		list          []string
		allWisply     = GetAllTables()
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

// IsRestrictedTable checks if a table name is on the restricted list
func IsRestrictedTable(name string) bool {
	for _, rejectedName := range rejectedTables {
		if name == rejectedName {
			return true
		}
	}
	return false
}
