package harvest

import (
	"errors"
	"strconv"

	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/repository"
)

// InsertFormatsTask represents a task that inserts the formats into database
type InsertFormatsTask struct {
	Tasker
	*Task
	repository *repository.Repository
}

// Insert saves the formats
func (task *InsertFormatsTask) Insert(formats []Formater) error {

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
	task.Finish(number + " formats inserted")

	return nil
}

func (task *InsertFormatsTask) hasProblems(err error) {
	task.ChangeResult("danger")
	task.Finish(err.Error())
}

func (task *InsertFormatsTask) clearTable() error {

	ID := task.repository.ID

	sql := "DELETE from `repository_format` WHERE repository=?"
	query, err := database.Database.Prepare(sql)

	if err != nil {
		return errors.New("Error while trying to clear the `repository_format` table: <br />" + err.Error())
	}

	query.Exec(strconv.Itoa(ID))

	return nil
}

func (task *InsertFormatsTask) insertData(formats []Formater) error {

	ID := task.repository.ID

	for _, format := range formats {
		sqlColumns := "(`repository`, `md_schema`, `namespace`, `prefix`)"
		sqlValues := "(?, ?, ?, ?)"
		sql := "INSERT INTO `repository_format` " + sqlColumns + " VALUES " + sqlValues

		query, err := database.Database.Prepare(sql)

		if err != nil {
			return errors.New("Error while trying to insert into `repository_format` table: <br />" + err.Error())
		}
		query.Exec(ID, format.GetSchema(), format.GetNamespace(), format.GetPrefix())
	}
	return nil
}

func newInsertFormatsTask(operationHarvest Operationer, repository *repository.Repository) *InsertFormatsTask {
	return &InsertFormatsTask{
		Task: &Task{
			operation: operationHarvest,
			Task:      newTask(operationHarvest.GetOperation(), "Insert Formats"),
		},
		repository: repository,
	}
}
