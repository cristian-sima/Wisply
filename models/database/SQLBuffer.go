package database

import "errors"

var maxRowsPerExecution = 100

// SQLBuffer is a faster way to insert data into database
// It keeps a buffer of all the rows and executes the statement only once
//
//
// The Buffer has an internal way to prevent a buffer overflow. If the number of rows stored in the buffer exceeds the maximum number (maxRowsPerExecution), the buffer executes and clears the memory and saves the result (which is an error). In case the error is not nil, it will return it when the Exec is called.
//
// Inspired by this post
// http://stackoverflow.com/questions/21108084/golang-mysql-insert-multiple-data-at-once
//
type SQLBuffer struct {
	memory        []interface{}
	questions     string
	table         string
	columns       string
	internalError error
	numberOfRows  int
}

// AddRow adds a new row to the buffer
func (buffer *SQLBuffer) AddRow(values ...interface{}) {

	if buffer.internalError == nil {

		row := "("

		for range values {
			row = row + "?,"
		}

		// delete the last and append ),

		row = row[0:len(row)-1] + "), "

		buffer.questions += row
		buffer.memory = append(buffer.memory, values...)

		buffer.numberOfRows++

		if buffer.numberOfRows > maxRowsPerExecution {
			buffer.internalError = buffer.Exec()
		}
	}
}

// Exec executes the entire buffer
func (buffer *SQLBuffer) Exec() error {

	if buffer.internalError != nil {
		return buffer.internalError
	}

	if buffer.memory == nil {
		return nil
	}

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
		return errors.New("Problem executing the buffer for table: `" + buffer.table + "`:" + err.Error() + "<br /><br />The query was: <br /><br />" + sqlStr)
	}
	return nil
}

func (buffer *SQLBuffer) clear() {
	buffer.memory = nil
	buffer.questions = ""
	buffer.numberOfRows = 0
}

// NewSQLBuffer creates a new buffer
// columns should be like this: column1, column2, column3
func NewSQLBuffer(table, columns string) *SQLBuffer {
	return &SQLBuffer{
		table:   table,
		columns: columns,
	}
}
