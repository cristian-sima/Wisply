package harvest

import (
	"errors"
	"strconv"

	"github.com/cristian-sima/Wisply/models/database"
	wisply "github.com/cristian-sima/Wisply/models/harvest/wisply"
	"github.com/cristian-sima/Wisply/models/repository"
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

	err := task.clearTable()

	if err != nil {
		task.hasProblems(err)
		return err
	}
	err = task.insertData(collections)

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
		return errors.New("Error while trying to clear the `repository_collection` table: <br />" + err.Error())
	}

	query.Exec(strconv.Itoa(repositoryID))

	return nil
}

func (task *InsertCollectionsTask) insertData(collections []wisply.Collectioner) error {
	repositoryID := task.repository.ID
	for _, collection := range collections {
		task.collectionsBuffer.AddRow(repositoryID, collection.GetName(), collection.GetSpec())
	}
	return task.collectionsBuffer.Exec()
}

func newInsertCollectionsTask(operationHarvest Operationer, repository *repository.Repository) *InsertCollectionsTask {
	collectionsBuffer := database.NewSQLBuffer("repository_collection", "`repository`, `name`, `spec`")
	return &InsertCollectionsTask{
		Task: &Task{
			operation: operationHarvest,
			Task:      newTask(operationHarvest.GetOperation(), "Insert Formats"),
		},
		collectionsBuffer: collectionsBuffer,
		repository:        repository,
	}
}
