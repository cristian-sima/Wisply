package harvest

import (
	"errors"
	"strconv"

	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/repository"
	"github.com/cristian-sima/Wisply/models/wisply"
)

// InsertIdentifiersTask represents a task that inserts the identifiers into database
type InsertIdentifiersTask struct {
	Tasker
	*Task
	repository        *repository.Repository
	identifiersBuffer *database.SQLBuffer
	setsBuffer        *database.SQLBuffer
}

// Insert clears the tables and inserts the identifiers
func (task *InsertIdentifiersTask) Insert(identifiers []wisply.Identifier) error {
	err := task.insertIdentifiers(identifiers)
	if err != nil {
		task.hasProblems(err)
		return err
	}
	// execute buffers
	err = task.identifiersBuffer.Exec()
	if err != nil {
		task.hasProblems(err)
		return err
	}
	err = task.setsBuffer.Exec()
	if err != nil {
		task.hasProblems(err)
		return err
	}
	number := strconv.Itoa(len(identifiers))
	task.Finish(number + " identifiers inserted")
	return nil
}

func (task *InsertIdentifiersTask) hasProblems(err error) {
	task.ChangeResult("danger")
	task.Finish(err.Error())
}

// Clear deletes all identifiers and sets
func (task *InsertIdentifiersTask) Clear() error {
	ID := task.repository.ID
	sql := "DELETE from `identifier` WHERE repository=?"
	query, err := database.Connection.Prepare(sql)
	if err != nil {
		message := "Error while trying to clear the `identifier` table: <br />" + err.Error()
		return errors.New(message)
	}
	query.Exec(ID)
	// clear sets
	sql = "DELETE from `identifier_set` WHERE repository=?"
	query, err = database.Connection.Prepare(sql)
	if err != nil {
		message := "Error while trying to clear the `identifier_set` table: <br />" + err.Error()
		return errors.New(message)
	}
	finishMessage := "All the previous identifiers and sets have been deleted"
	task.Finish(finishMessage)
	_, err = query.Exec(ID)
	return err
}

func (task *InsertIdentifiersTask) insertIdentifiers(identifiers []wisply.Identifier) error {
	for _, identifier := range identifiers {
		err := task.insertIdentifier(identifier)
		if err != nil {
			return err
		}
	}
	return nil
}

func (task *InsertIdentifiersTask) insertIdentifier(identifier wisply.Identifier) error {
	err := task.insertData(identifier)
	if err != nil {
		return err
	}
	err = task.insertSets(identifier.GetIdentifier(), identifier.GetSpec())
	if err != nil {
		return err
	}
	return nil
}

func (task *InsertIdentifiersTask) insertData(identifier wisply.Identifier) error {
	ID := task.repository.ID
	task.identifiersBuffer.AddRow(identifier.GetIdentifier(), identifier.GetTimestamp(), ID)
	return nil
}

func (task *InsertIdentifiersTask) insertSets(identifier string, sets []string) error {
	for _, set := range sets {
		task.setsBuffer.AddRow(identifier, set, task.repository.ID)
	}
	return nil
}

func newInsertIdentifiersTask(operationHarvest Operationer, repository *repository.Repository) *InsertIdentifiersTask {
	var createIdentifiersBuffer = func() *database.SQLBuffer {
		columns := "`identifier`, `value`, `repository`"
		tableName := "identifier"
		return database.NewSQLBuffer(tableName, columns)
	}
	var createSetsBuffer = func() *database.SQLBuffer {
		columns := "`identifier`, `setSpec`, `repository`"
		tableName := "identifier_set"
		return database.NewSQLBuffer(tableName, columns)
	}
	identifiersBuffer := createIdentifiersBuffer()
	setsBuffer := createSetsBuffer()

	return &InsertIdentifiersTask{
		Task: &Task{
			operation: operationHarvest,
			Task:      newTask(operationHarvest.GetOperation(), "Insert Identifiers"),
		},
		repository:        repository,
		identifiersBuffer: identifiersBuffer,
		setsBuffer:        setsBuffer,
	}
}
