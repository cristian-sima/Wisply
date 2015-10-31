package harvest

import (
	"errors"
	"strconv"

	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/repository"
	wisply "github.com/cristian-sima/Wisply/models/wisply/data"
)

// InsertFormatsTask represents a task that inserts the formats into database
type InsertFormatsTask struct {
	Tasker
	*Task
	repository    *repository.Repository
	formatsBuffer *database.SQLBuffer
}

// Insert saves the formats
func (task *InsertFormatsTask) Insert(formats []wisply.Formater) error {
	err := task.clearTable()
	if err != nil {
		task.hasProblems(err)
		return err
	}
	err = task.insertData(formats)
	if err != nil {
		task.hasProblems(err)
		return err
	}
	number := strconv.Itoa(len(formats))
	message := number + " formats inserted"
	task.Finish(message)
	return nil
}

func (task *InsertFormatsTask) hasProblems(err error) {
	task.ChangeResult("danger")
	task.Finish(err.Error())
}

func (task *InsertFormatsTask) clearTable() error {
	ID := task.repository.ID
	sql := "DELETE from `repository_format` WHERE repository=?"
	query, err := database.Connection.Prepare(sql)
	if err != nil {
		errorMessage := "<br />" + err.Error()
		message := "Error while trying to clear the `repository_format` table:" + errorMessage
		return errors.New(message)
	}
	query.Exec(strconv.Itoa(ID))
	return nil
}

func (task *InsertFormatsTask) insertData(formats []wisply.Formater) error {
	repositoryID := task.repository.ID
	for _, format := range formats {
		task.formatsBuffer.AddRow(repositoryID, format.GetSchema(), format.GetNamespace(), format.GetPrefix())
	}
	return task.formatsBuffer.Exec()
}

func newInsertFormatsTask(operationHarvest Operationer, repository *repository.Repository) *InsertFormatsTask {
	tableName := "repository_format"
	columns := "`repository`, `md_schema`, `namespace`, `prefix`"
	formatsBuffer := database.NewSQLBuffer(tableName, columns)
	return &InsertFormatsTask{
		Task: &Task{
			operation: operationHarvest,
			Task:      newTask(operationHarvest.GetOperation(), "Insert Formats"),
		},
		repository:    repository,
		formatsBuffer: formatsBuffer,
	}
}
