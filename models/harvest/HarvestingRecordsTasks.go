package harvest

import (
	"errors"
	"strconv"

	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/repository"
	wisply "github.com/cristian-sima/Wisply/models/wisply/data"
)

// InsertRecordsTask represents a task that inserts the records into database
type InsertRecordsTask struct {
	Tasker
	*Task
	repository       *repository.Repository
	keysBuffer       *database.SQLBuffer
	repositoryBuffer *database.SQLBuffer
}

// Insert clears the tables and inserts the records
func (task *InsertRecordsTask) Insert(records []wisply.Recorder) error {
	err := task.insertRecords(records)
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

// Clear deletes all the records
func (task *InsertRecordsTask) Clear() error {
	ID := task.repository.ID
	sql := "DELETE from `repository_resource` WHERE repository=?"
	query, err := database.Connection.Prepare(sql)
	if err != nil {
		message := "Error while trying to clear the `repository_resource` table: <br />" + err.Error()
		return errors.New(message)
	}
	query.Exec(ID)
	// clear keys
	sql = "DELETE from `resource_key` WHERE repository=?"
	query, err = database.Connection.Prepare(sql)
	if err != nil {
		message := "Error while trying to clear the `resource_key` table: <br />" + err.Error()
		return errors.New(message)
	}
	query.Exec(ID)
	finishMessage := "All the previous records and keys have been deleted"
	task.Finish(finishMessage)
	return nil
}

func (task *InsertRecordsTask) insertRecords(records []wisply.Recorder) error {
	for _, record := range records {
		err := task.insertRecord(record)
		if err != nil {
			return err
		}
	}
	// execute the buffers
	err := task.repositoryBuffer.Exec()
	if err != nil {
		return err
	}
	err = task.keysBuffer.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (task *InsertRecordsTask) insertRecord(record wisply.Recorder) error {
	ID := task.repository.ID
	task.repositoryBuffer.AddRow(ID, record.GetIdentifier(), record.GetTimestamp())
	return task.saveKeys(&record)
}

func (task *InsertRecordsTask) saveKeys(record *wisply.Recorder) error {
	return task.insertTitles(record)
}

// There is no way to compress them
func (task *InsertRecordsTask) insertTitles(record *wisply.Recorder) error {
	var keys = (*record).GetKeys()
	err := task.insertKeys(record, keys.GetTitles(), "title")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, keys.GetCreators(), "creator")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, keys.GetSubjects(), "subject")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, keys.GetDescriptions(), "description")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, keys.GetPublishers(), "publisher")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, keys.GetContributors(), "contributor")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, keys.GetDates(), "date")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, keys.GetTypes(), "type")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, keys.GetFormats(), "format")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, keys.GetIdentifiers(), "identifier")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, keys.GetSources(), "source")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, keys.GetLanguages(), "language")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, keys.GetRelations(), "relation")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, keys.GetCoverages(), "coverage")

	if err != nil {
		return err
	}
	err = task.insertKeys(record, keys.GetRights(), "right")

	if err != nil {
		return err
	}
	return nil
}

func (task *InsertRecordsTask) insertKeys(record *wisply.Recorder, keys []string, name string) error {
	ID := task.repository.ID
	for _, value := range keys {
		task.keysBuffer.AddRow(name, ID, value, (*record).GetIdentifier())
	}
	return nil
}

func newInsertRecordsTask(operationHarvest Operationer, repository *repository.Repository) *InsertRecordsTask {
	var createKeysBuffer = func() *database.SQLBuffer {
		columns := "`resource_key`, `repository`, `value`, `resource`"
		tableName := "resource_key"
		return database.NewSQLBuffer(tableName, columns)
	}
	var createRepositoryBuffer = func() *database.SQLBuffer {
		columns := "`repository`, `identifier`, `datestamp`"
		tableName := "repository_resource"
		return database.NewSQLBuffer(tableName, columns)
	}
	keysBuffer := createKeysBuffer()
	repositoryBuffer := createRepositoryBuffer()

	return &InsertRecordsTask{
		Task: &Task{
			operation: operationHarvest,
			Task:      newTask(operationHarvest.GetOperation(), "Insert Records"),
		},
		repository:       repository,
		keysBuffer:       keysBuffer,
		repositoryBuffer: repositoryBuffer,
	}
}

// UpdateNumberOfRecordsTask represents a task that updates the number of
// records for all the collections
type UpdateNumberOfRecordsTask struct {
	Tasker
	*Task
	repository       *repository.Repository
	keysBuffer       *database.SQLBuffer
	repositoryBuffer *database.SQLBuffer
}

// Perform gets the number of records for each collection
func (task *UpdateNumberOfRecordsTask) Perform() error {
	err := task.update()
	if err != nil {
		task.hasProblems(err)
		return err
	}
	task.Finish("The number of records has been updated")
	return err
}

func (task *UpdateNumberOfRecordsTask) update() error {

	space := " "

	selectWhereClause := "WHERE `identifier_set`.`setSpec` = `repository_collection`.`spec`"
	numberOfRecords := "SELECT COUNT(*) FROM `identifier_set`" + space + selectWhereClause
	setClause := "SET `repository_collection`.`numberOfRecords` = (" + numberOfRecords + ")"

	whereClause := "WHERE `repository_collection`.`repository` = ?"
	sql := "UPDATE `repository_collection` " + setClause + space + whereClause

	repositoryID := task.repository.ID
	_, err := database.Connection.Query(sql, repositoryID)
	return err
}

func (task *UpdateNumberOfRecordsTask) hasProblems(err error) {
	task.ChangeResult("danger")
	task.Finish(err.Error())
}

func newUpdateNumberOfRecordsTask(operationHarvest Operationer, repository *repository.Repository) *UpdateNumberOfRecordsTask {

	return &UpdateNumberOfRecordsTask{
		Task: &Task{
			operation: operationHarvest,
			Task:      newTask(operationHarvest.GetOperation(), "Update the number of records"),
		},
		repository: repository,
	}
}
