package harvest

import (
	"errors"
	"strconv"

	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/wisply"
	"github.com/cristian-sima/Wisply/models/repository"
)

// InsertIdentificationTask represents a task that inserts the identification into database
type InsertIdentificationTask struct {
	Tasker
	*Task
	repository *repository.Repository
}

// Insert inserts the identification details
func (task *InsertIdentificationTask) Insert(identification *wisply.Identificationer) error {
	err := task.clearTables()

	if err != nil {
		task.ChangeResult("danger")
		task.Finish(err.Error())
		return err
	}

	return task.insertData(identification)
}

func (task *InsertIdentificationTask) clearTables() error {

	ID := task.repository.ID

	sql := "DELETE from `repository_identification` WHERE repository=?"
	query, err1 := database.Connection.Prepare(sql)

	if err1 != nil {
		return errors.New("Error while trying to clear the repository_identification table: <br />" + err1.Error())
	}

	query.Exec(strconv.Itoa(ID))

	sql2 := "DELETE from `repository_email` WHERE repository=?"
	query2, err2 := database.Connection.Prepare(sql2)

	if err2 != nil {
		return errors.New("Error while trying to clear the repository_email table: <br />" + err2.Error())
	}
	query2.Exec(strconv.Itoa(ID))

	return nil
}

func (task *InsertIdentificationTask) insertData(identification *wisply.Identificationer) error {
	err1 := task.insertDetails(identification)
	if err1 != nil {
		task.ChangeResult("danger")
		task.Finish(err1.Error())
		return err1
	}
	err2 := task.insertEmails((*identification).GetAdminEmails())
	if err2 != nil {
		task.ChangeResult("danger")
		task.Finish(err2.Error())
		return err1
	}
	task.Finish("Success")
	return nil

}

func (task *InsertIdentificationTask) insertEmails(emails []string) error {

	ID := strconv.Itoa(task.repository.ID)

	for _, email := range emails {

		sqlColumns := "(`repository`, `email`)"
		sqlValues := "(?, ?)"
		sql := "INSERT INTO `repository_email` " + sqlColumns + " VALUES " + sqlValues

		query, err := database.Connection.Prepare(sql)
		if err != nil {
			return errors.New("There was problem while trying to insert the email addresses: " + err.Error())
		}
		query.Exec(ID, email)
	}
	return nil
}

func (task *InsertIdentificationTask) insertDetails(identification *wisply.Identificationer) error {

	modifyName := "UPDATE `repository` SET `name`=? WHERE `id` = ?"

	query1, err1 := database.Connection.Prepare(modifyName)

	if err1 != nil {
		return errors.New("Eror while trying to insert into `repository_identification`: <br />" + err1.Error())
	}
	query1.Exec((*identification).GetName(), task.repository.ID)

	sqlColumns := "(`repository`, `protocol_version`, `earliest_datestamp`, `delete_policy`, `granularity`)"
	sqlValues := "(?, ?, ?, ?, ?)"
	idenSQL := "INSERT INTO `repository_identification` " + sqlColumns + " VALUES " + sqlValues

	query2, err2 := database.Connection.Prepare(idenSQL)

	if err2 != nil {
		return errors.New("Eror while trying to insert into `repository_identification`: <br />" + err2.Error())
	}
	query2.Exec(task.repository.ID, (*identification).GetProtocol(), (*identification).GetEarliestDatestamp(), (*identification).GetDeletedRecord(), (*identification).GetGranularity())

	return nil
}

func newInsertIdentificationTask(operationHarvest Operationer, repository *repository.Repository) *InsertIdentificationTask {
	return &InsertIdentificationTask{
		Task: &Task{
			operation: operationHarvest,
			Task:      newTask(operationHarvest.GetOperation(), "Insert Identification details"),
		},
		repository: repository,
	}
}
