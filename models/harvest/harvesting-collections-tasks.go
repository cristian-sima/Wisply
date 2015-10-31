package harvest

import (
	"errors"
	"strconv"

	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/repository"
	wisply "github.com/cristian-sima/Wisply/models/wisply/data"
)

// InsertCollectionsTask represents a task that inserts the collections into database
type InsertCollectionsTask struct {
	Tasker
	*Task
	repository        *repository.Repository
	collectionsBuffer *database.SQLBuffer
}

// Insert clears the table and then inserts them
func (task *InsertCollectionsTask) Insert(collections []wisply.Collectioner) error {
	err := task.insertData(collections)
	if err != nil {
		task.hasProblems(err)
		return err
	}
	number := strconv.Itoa(len(collections))
	task.Finish(number + " collections inserted")
	return nil
}

func (task *InsertCollectionsTask) hasProblems(err error) {
	task.ChangeResult("danger")
	task.Finish(err.Error())
}

func (task *InsertCollectionsTask) clearTable() error {
	repositoryID := task.repository.ID
	sql := "DELETE from `repository_collection` WHERE repository=?"
	query, err := database.Connection.Prepare(sql)

	if err != nil {
		errorMessage := "<br />" + err.Error()
		message := "Error while trying to clear the `repository_collection` table:" + errorMessage
		return errors.New(message)
	}
	query.Exec(strconv.Itoa(repositoryID))
	return nil
}

func (task *InsertCollectionsTask) insertData(collections []wisply.Collectioner) error {
	repositoryID := task.repository.ID
	for _, collection := range collections {
		task.collectionsBuffer.AddRow(repositoryID, collection.GetPath(), collection.GetSpec(), collection.GetName())
	}
	return task.collectionsBuffer.Exec()
}

func newInsertCollectionsTask(operationHarvest Operationer, repository *repository.Repository) *InsertCollectionsTask {
	tableName := "repository_collection"
	columns := "`repository`, `path`, `spec`, `name`"
	collectionsBuffer := database.NewSQLBuffer(tableName, columns)
	return &InsertCollectionsTask{
		Task: &Task{
			operation: operationHarvest,
			Task:      newTask(operationHarvest.GetOperation(), "Insert Formats"),
		},
		collectionsBuffer: collectionsBuffer,
		repository:        repository,
	}
}
