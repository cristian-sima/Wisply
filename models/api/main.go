package api

import (
	"fmt"
	"os"

	"github.com/cristian-sima/Wisply/models/database"
)

var sensitiveTableList = []string{"account", "account_token", "api_table_settings"}

// GenerateTableFile creates the sql table
func GenerateTableFile(tableName string) *os.File {

	path := "cache/api/table/" + tableName + ".csv"
	// generate
	sql := `SELECT id, name
		FROM account
		INTO OUTFILE '` + path + `'
		FIELDS TERMINATED BY ','
		ENCLOSED BY '"'
		LINES TERMINATED BY '\n';`
	query, err := database.Connection.Prepare(sql)
	query.Exec(tableName)

	if err != nil {
		fmt.Println("Error sql")
		fmt.Println(err)
	}

	// get the file
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	return file
}

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

// InsertNewTable adds the table name on the list of the tables which can be downloaded
func InsertNewTable(tableName string) error {
	sql := "INSERT INTO `api_table_settings` (`name`) VALUES (?)"
	query, err := database.Connection.Prepare(sql)
	query.Exec(tableName)
	return err
}

// RemoveAllowedTable removes the table from the list of allowed tables
func RemoveAllowedTable(ID int) error {
	sql := "DELETE FROM `api_table_settings` WHERE `id`=? "
	query, err := database.Connection.Prepare(sql)
	query.Exec(ID)
	return err
}

// GetAllowedTables returns the list of allowed tables
func GetAllowedTables() []Table {
	var list []Table
	sql := "SELECT `id`, `name` FROM `api_table_settings`"
	rows, _ := database.Connection.Query(sql)
	for rows.Next() {
		table := Table{}
		rows.Scan(&table.ID, &table.Name)
		list = append(list, table)
	}
	return list
}

// GetWisplyTablesNamesNotAllowed returns the list of wisply tables which can be downloaded, but are not yet on the list
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
