package developer

import (
	"fmt"
	"strings"

	"github.com/cristian-sima/Wisply/models/database"
)

type csvFile struct {
	*file
	columns       string
	headerColumns string
}

func (file *csvFile) generate() {

	file.columns = file.getRawColumns()
	file.headerColumns = file.getHeaderColumns()

	fullPath := file.getFullPath()

	err := file.execute(fullPath)

	if err != nil {
		tempPath := "/tmp/" + file.getFileName()
		// unix does not allow mysql to write everywhere
		// but it allows it to write to "/tmp"
		// so we create a file there and then copy it to the good location
		deleteFile(tempPath)
		file.execute(tempPath)
		copyFile(fullPath, tempPath)
	}
}

func (file *csvFile) execute(path string) error {
	sql := `SELECT ` + file.headerColumns + `
  UNION ALL
  SELECT ` + file.columns + `
  FROM ` + file.name + `
  INTO OUTFILE '` + path + `'
  FIELDS TERMINATED BY ','
  ENCLOSED BY '"'
  LINES TERMINATED BY '\n';`
	query, err1 := database.Connection.Prepare(sql)
	if err1 != nil {
		fmt.Println("Error #1 generating the file from SQL: ")
		fmt.Println(err1)
		return err1
	}
	_, err2 := query.Exec()
	if err2 != nil {
		fmt.Println("Error #2 generating the file from SQL: ")
		fmt.Println(err2)
		return err2
	}
	return nil
}

func (file *csvFile) getHeaderColumns() string {
	elements := strings.Split(file.columns, ",")
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
