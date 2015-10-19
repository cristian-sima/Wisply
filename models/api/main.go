package api

import (
	"fmt"
	"strings"

	"github.com/cristian-sima/Wisply/models/database"
)

var sensitiveTableList = []string{"account", "account_token", "api_table_settings"}

// GenerateTableFile creates the sql table
func GenerateTableFile(tableName string) {

	columns := ""

	sqlCol := `SELECT  GROUP_CONCAT(COLUMN_NAME SEPARATOR ',')
		FROM INFORMATION_SCHEMA.COLUMNS
		WHERE TABLE_SCHEMA='wisply' AND TABLE_NAME='` + tableName + `'`
	rows, err := database.Connection.Prepare(sqlCol)
	rows.QueryRow().Scan(&columns)

	if err != nil {
		fmt.Println(err)
	}

	elements := strings.Split(columns, ",")
	stringColumns := ""

	for _, element := range elements {
		stringColumns += `"` + element + `",`
	}
	stringColumns = stringColumns[:len(stringColumns)-1]

	sql := `SELECT ` + stringColumns + `
		UNION ALL
		SELECT ` + columns + `
		FROM ` + tableName + `
		INTO OUTFILE 'W:/go-workspace/src/github.com/cristian-sima/Wisply/cache/api/tables/` + tableName + `.csv'
		FIELDS TERMINATED BY ','
		ENCLOSED BY '"'
		LINES TERMINATED BY '\n';`

	query, err1 := database.Connection.Prepare(sql)
	if err1 != nil {
		fmt.Println("Error sql")
		panic(err1)
	}

	_, err2 := query.Exec()
	if err2 != nil {
		fmt.Println("Error sql")
		panic(err2)
	}
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

// AreValidDetails checks if the table is allowed and the description is valid
func AreValidDetails(table Table) bool {
	return !IsAllowedTable(table.Name) && isValidDescription(table.Description)
}

// InsertNewTable adds the table name on the list of the tables which can be downloaded
func InsertNewTable(table Table) error {
	sql := "INSERT INTO `api_table_settings` (`name`, `description`) VALUES (?, ?)"
	query, err := database.Connection.Prepare(sql)
	query.Exec(table.Name, table.Description)
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
	sql := "SELECT `id`, `name`, `description` FROM `api_table_settings`"
	rows, _ := database.Connection.Query(sql)
	for rows.Next() {
		table := Table{}
		rows.Scan(&table.ID, &table.Name, &table.Description)
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
