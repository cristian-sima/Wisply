package api

import (
	"fmt"
	"strings"

	"github.com/cristian-sima/Wisply/models/database"
)

type csvFile struct {
	*file
}

func (file *csvFile) generate() {

	columns := file.getRawColumns()
	headerColumns := file.getHeaderColumns(columns)
	fullPath := file.getFullPath()

	sql := `SELECT ` + headerColumns + `
	UNION ALL
	SELECT ` + columns + `
	FROM ` + file.name + `
	INTO OUTFILE '` + fullPath + `'
	FIELDS TERMINATED BY ','
	ENCLOSED BY '"'
	LINES TERMINATED BY '\n';`

	query, err1 := database.Connection.Prepare(sql)
	if err1 != nil {
		fmt.Println("Error 1 generating the file from SQL: ")
		panic(err1)
	}

	_, err2 := query.Exec()
	if err2 != nil {
		fmt.Println("Error 2 generating the file from SQL: ")
		panic(err2)
	}
}

func (file *csvFile) getHeaderColumns(rawColumns string) string {
	elements := strings.Split(rawColumns, ",")
	stringColumns := ""
	for _, element := range elements {
		stringColumns += `"` + element + `",`
	}
	return stringColumns[:len(stringColumns)-1]
}

func (file *csvFile) getRawColumns() string {
	columns := ""
	sqlCol := `SELECT  GROUP_CONCAT(COLUMN_NAME SEPARATOR ',')
	FROM INFORMATION_SCHEMA.COLUMNS
	WHERE TABLE_SCHEMA='wisply' AND TABLE_NAME='` + file.name + `'`
	rows, err := database.Connection.Prepare(sqlCol)
	rows.QueryRow().Scan(&columns)
	if err != nil {
		fmt.Println("Error 0 getting the file columns: ")
		panic(err)
	}
	return columns
}

func createCSVFile(tableName string) *csvFile {
	return &csvFile{
		file: createFile(tableName, "csv"),
	}
}
