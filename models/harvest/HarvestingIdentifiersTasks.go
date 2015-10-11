package harvest

import (
	"errors"
	"strconv"

	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/harvest/wisply"
	"github.com/cristian-sima/Wisply/models/repository"
)

// InsertIdentifiersTask represents a task that inserts the identifiers into database
type InsertIdentifiersTask struct {
	Tasker
	*Task
	repository *repository.Repository
}

// Insert clears the tables and inserts the identifiers
func (task *InsertIdentifiersTask) Insert(identifiers []wisply.Identifier) error {

	err := task.clear()

	if err != nil {
		task.hasProblems(err)
		return err
	}
	err = task.insertIdentifiers(identifiers)

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

func (task *InsertIdentifiersTask) clear() error {

	ID := task.repository.ID

	sql := "DELETE from `identifier` WHERE repository=?"
	query, err := database.Connection.Prepare(sql)

	if err != nil {
		return errors.New("Error while trying to clear the `identifier` table: <br />" + err.Error())
	}

	query.Exec(ID)

	// clear sets

	sql = "DELETE from `identifier_set` WHERE repository=?"
	query, err = database.Connection.Prepare(sql)

	if err != nil {
		return errors.New("Error while trying to clear the `identifier_set` table: <br />" + err.Error())
	}

	query.Exec(ID)

	return nil
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
	sqlColumns := "(`repository`, `identifier`, `datestamp`)"
	sqlValues := "(?, ?, ?)"
	sql := "INSERT INTO `identifier` " + sqlColumns + " VALUES " + sqlValues

	query, err := database.Connection.Prepare(sql)

	if err != nil {
		return errors.New("Error while trying to insert into `identifier` table: <br />" + err.Error())
	}
	query.Exec(ID, identifier.GetIdentifier(), identifier.GetTimestamp())
	return nil
}

func (task *InsertIdentifiersTask) insertSets(identifier string, sets []string) error {
	for _, set := range sets {
		ID := task.repository.ID
		sqlColumns := "(`repository`, `identifier`, `setSpec`)"
		sqlValues := "(?, ?, ?)"
		sql := "INSERT INTO `identifier_set` " + sqlColumns + " VALUES " + sqlValues

		query, err := database.Connection.Prepare(sql)

		if err != nil {
			return errors.New("Error while trying to insert into `identifier_set` table: <br />" + err.Error())
		}
		query.Exec(ID, identifier, set)
	}
	return nil
}

func newInsertIdentifiersTask(operationHarvest Operationer, repository *repository.Repository) *InsertIdentifiersTask {
	return &InsertIdentifiersTask{
		Task: &Task{
			operation: operationHarvest,
			Task:      newTask(operationHarvest.GetOperation(), "Insert Identifiers"),
		},
		repository: repository,
	}
}
