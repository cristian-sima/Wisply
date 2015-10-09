package harvest

import (
	"errors"
	"strconv"

	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/repository"
)

// InsertRecordsTask represents a task that inserts the records into database
type InsertRecordsTask struct {
	Tasker
	*Task
	repository *repository.Repository
}

// Insert clears the tables and inserts the records
func (task *InsertRecordsTask) Insert(records []Recorder) error {

	err := task.clear()

	if err != nil {
		task.hasProblems(err)
		return err
	}
	err = task.insertRecords(records)

	if err != nil {
		task.hasProblems(err)
		return err
	}

	number := strconv.Itoa(len(records))
	task.Finish(number + " records inserted")

	return nil
}

func (task *InsertRecordsTask) hasProblems(err error) {
	task.ChangeResult("danger")
	task.Finish(err.Error())
}

func (task *InsertRecordsTask) clear() error {

	ID := task.repository.ID

	sql := "DELETE from `repository_resource` WHERE repository=?"
	query, err := database.Database.Prepare(sql)

	if err != nil {
		return errors.New("Error while trying to clear the `repository_resource` table: <br />" + err.Error())
	}

	query.Exec(ID)

	// clear keys

	sql = "DELETE from `resource_key` WHERE repository=?"
	query, err = database.Database.Prepare(sql)

	if err != nil {
		return errors.New("Error while trying to clear the `resource_key` table: <br />" + err.Error())
	}

	query.Exec(ID)

	return nil
}

func (task *InsertRecordsTask) insertRecords(records []Recorder) error {
	for _, record := range records {
		err := task.insertRecord(record)
		if err != nil {
			return err
		}
	}
	return nil
}

func (task *InsertRecordsTask) insertRecord(record Recorder) error {

	ID := task.repository.ID

	sqlColumns := "(`repository`, `identifier`, `datestamp`)"
	sqlValues := "(?, ?, ?)"
	sql := "INSERT INTO `repository_resource` " + sqlColumns + " VALUES " + sqlValues

	query, err := database.Database.Prepare(sql)

	if err != nil {
		return errors.New("Error while trying to insert into `repository_resource` table: <br />" + err.Error())
	}

	query.Exec(ID, record.GetIdentifier(), record.GetDatestamp())

	return task.saveKeys(&record)
}

func (task *InsertRecordsTask) saveKeys(record *Recorder) error {
	return task.insertTitles(record)
}

// There is no way to compress them
func (task *InsertRecordsTask) insertTitles(record *Recorder) error {
	var keys = (*record).GetKeys()
	err := task.insertKeys(record, keys.Titles, "title")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, keys.Creators, "creator")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, keys.Subjects, "subject")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, keys.Descriptions, "description")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, keys.Publishers, "publisher")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, keys.Contributors, "contributor")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, keys.Dates, "date")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, keys.Types, "type")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, keys.Formats, "format")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, keys.Identifiers, "identifier")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, keys.Sources, "source")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, keys.Languages, "language")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, keys.Relations, "relation")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, keys.Coverages, "coverage")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, keys.Rights, "right")

	if err != nil {
		return err
	}
	return nil
}

func (task *InsertRecordsTask) insertKeys(record *Recorder, keys []string, name string) error {

	ID := task.repository.ID

	for _, value := range keys {
		sqlColumns := "(`repository`, `resource`, `value`, `resource_key`)"
		sqlValues := "(?, ?, ?, ?)"
		sql := "INSERT INTO `resource_key` " + sqlColumns + " VALUES " + sqlValues

		query, err := database.Database.Prepare(sql)

		if err != nil {
			return errors.New("Error while inserting into `resource_key`: " + err.Error())
		}

		query.Exec(ID, (*record).GetIdentifier(), value, name)
	}
	return nil
}

func newInsertRecordsTask(operationHarvest Operationer, repository *repository.Repository) *InsertRecordsTask {
	return &InsertRecordsTask{
		Task: &Task{
			operation: operationHarvest,
			Task:      newTask(operationHarvest.GetOperation(), "Insert Records"),
		},
		repository: repository,
	}
}
