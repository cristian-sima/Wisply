package harvest

import (
	"errors"
	"strconv"

	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/harvest/wisply"
	"github.com/cristian-sima/Wisply/models/repository"
)

// InsertRecordsTask represents a task that inserts the records into database
type InsertRecordsTask struct {
	Tasker
	*Task
	repository *repository.Repository
}

// Insert clears the tables and inserts the records
func (task *InsertRecordsTask) Insert(records []wisply.Recorder) error {

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

func (task *InsertRecordsTask) insertRecords(records []wisply.Recorder) error {
	for _, record := range records {
		err := task.insertRecord(record)
		if err != nil {
			return err
		}
	}
	return nil
}

func (task *InsertRecordsTask) insertRecord(record wisply.Recorder) error {

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

func (task *InsertRecordsTask) saveKeys(record *wisply.Recorder) error {
	return task.insertTitles(record)
}

// There is no way to compress them
func (task *InsertRecordsTask) insertTitles(record *wisply.Recorder) error {
	var keys = (*record).GetKeys()
	err := task.insertKeys(record, (*keys).GetTitles(), "title")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, (*keys).GetCreators(), "creator")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, (*keys).GetSubjects(), "subject")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, (*keys).GetDescriptions(), "description")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, (*keys).GetPublishers(), "publisher")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, (*keys).GetContributors(), "contributor")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, (*keys).GetDates(), "date")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, (*keys).GetTypes(), "type")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, (*keys).GetFormats(), "format")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, (*keys).GetIdentifiers(), "identifier")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, (*keys).GetSources(), "source")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, (*keys).GetLanguages(), "language")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, (*keys).GetRelations(), "relation")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, (*keys).GetCoverages(), "coverage")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, (*keys).GetRights(), "right")

	if err != nil {
		return err
	}
	return nil
}

func (task *InsertRecordsTask) insertKeys(record *wisply.Recorder, keys []string, name string) error {

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