package database

import "errors"

// SQLBuffer is a fater way to insert data into database
// It keeps a buffer of all the rows and executes the statement only once
//
// Inspired by this post
// http://stackoverflow.com/questions/21108084/golang-mysql-insert-multiple-data-at-once
//
type SQLBuffer struct {
	memory    []interface{}
	questions string
	table     string
	columns   string
}

// AddRow adds a new row to the buffer
func (buffer *SQLBuffer) AddRow(values ...interface{}) {

	row := ""

	row = "("

	for range values {
		row = row + "?,"
	}
	// delete the last ,
	row = row[0 : len(row)-1]

	row += "), "

	buffer.questions += row
	buffer.memory = append(buffer.memory, values...)

}

// Exec executes the entire buffer
func (buffer *SQLBuffer) Exec() error {

	sqlStr := "INSERT INTO `" + buffer.table + "` (" + buffer.columns + ")" + " VALUES "
	sqlStr += buffer.questions

	//trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-2]
	//prepare the statement

	stmt, err := Connection.Prepare(sqlStr)

	if err != nil {
		return err
	}

	//format all vals at once
	_, err = stmt.Exec(buffer.memory...)

	buffer.clear()

	if err != nil {
		return errors.New("Problem executing the buffer for table `" + buffer.table + "`:" + err.Error())
	}
	return nil
}

func (buffer *SQLBuffer) clear() {
	buffer.memory = nil
	buffer.questions = ""
}

// NewSQLBuffer creates a new buffer
// columns should be like this: column1, column2, column3
func NewSQLBuffer(table, columns string) *SQLBuffer {
	return &SQLBuffer{
		table:   table,
		columns: columns,
	}
}
