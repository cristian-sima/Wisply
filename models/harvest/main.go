package harvest

import (
	"fmt"
	"strconv"

	"github.com/cristian-sima/Wisply/models/action"
	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/harvest/remote"
	"github.com/cristian-sima/Wisply/models/repository"
)

// CreateProcess creates a new harvest process
func CreateProcess(repositoryID string, controller Controller) *Process {

	process := &*action.CreateProcess("Harvest")

	harvestProcess := buildProcess(repositoryID, process)
	harvestProcess.controller = controller

	harvestID := insertHarvestProcess(harvestProcess)

	harvestProcess.HarvestID = harvestID

	return harvestProcess
}

func buildProcess(ID string, process *action.Process) *Process {

	local, _ := repository.NewRepository(ID)
	remoteServer, _ := remote.New(local)

	harvestProcess := &Process{
		Process:    process,
		remote:     remoteServer,
		repository: local,
	}

	return harvestProcess
}

func insertHarvestProcess(process *Process) int {

	var (
		harvestID                                  int
		formats, collections, records, identifiers int
	)

	if process.GetRepository().LastProcess != 0 {
		last := NewProcess(process.GetRepository().LastProcess)
		formats, collections, records, identifiers = last.GetStatistics()
	}

	columns := "(`process`, `repository`, `formats`, `collections`, `records`, `identifiers`)"
	values := "(?, ?, ?, ?, ?, ?)"
	sql := "INSERT INTO `process_harvest` " + columns + " VALUES " + values

	query, err := database.Connection.Prepare(sql)

	if err != nil {
		fmt.Println("Error when creating the harvest process:")
		fmt.Println(sql)
		fmt.Println(err)
	}
	query.Exec(process.ID, process.GetRepository().ID, formats, collections, records, identifiers)

	// find its ID
	sql = "SELECT `id` FROM `process_harvest` WHERE process=? AND repository=? ORDER by id LIMIT 0,1"
	query, err = database.Connection.Prepare(sql)
	query.QueryRow(process.ID, process.GetRepository().ID).Scan(&harvestID)

	if err != nil {
		fmt.Println("Error when selecting the harvest id:")
		fmt.Println(err)
	}

	process.Formats = formats
	process.Collections = collections
	process.Records = records
	process.Identifiers = identifiers

	return harvestID
}

// NewProcessByID loads the harvest process from database by its own ID
func NewProcessByID(harvestProcessID int) *Process {
	var processID int
	sql := "SELECT `process` FROM `process_harvest` WHERE id=?"
	query, err := database.Connection.Prepare(sql)

	if err != nil {
		fmt.Println("Error 123 when selecting the ID of repository from harvest process:")
		fmt.Println(err)
	}
	query.QueryRow(harvestProcessID).Scan(&processID)

	return NewProcess(processID)
}

// NewProcess selects from database and creates a harvest.Process by ID
// NOTE! It returns only the Repository
func NewProcess(processID int) *Process {

	var (
		repID, harvestID                           int
		local                                      *repository.Repository
		formats, collections, records, identifiers int
	)

	sql := "SELECT `id`, `repository`, `formats`, `collections`, `records`, `identifiers` FROM `process_harvest` WHERE process=?"
	query, err := database.Connection.Prepare(sql)

	if err != nil {
		fmt.Println("Error when selecting the ID of repository from harvest process:")
		fmt.Println(err)
	}
	query.QueryRow(processID).Scan(&harvestID, &repID, &formats, &collections, &records, &identifiers)

	local, err2 := repository.NewRepository(strconv.Itoa(repID))

	if err2 != nil {
		fmt.Println(err2)
	}

	remoteServer, _ := remote.New(local)

	return &Process{
		Formats:     formats,
		Records:     records,
		Collections: collections,
		Identifiers: identifiers,
		HarvestID:   harvestID,
		repository:  local,
		remote:      remoteServer,
		Process:     &*action.NewProcess(processID),
	}
}

// GetProcessesByRepository returns the processes of for the repository
func GetProcessesByRepository(repositoryID int) []*Process {

	var (
		list                                       []*Process
		processID, harvestID                       int
		repID                                      string
		formats, collections, records, identifiers int
	)

	repID = strconv.Itoa(repositoryID)

	sql := "SELECT `id`, `process`, `formats`, `collections`, `records`, `identifiers` FROM `process_harvest` WHERE `repository` = ? ORDER BY process DESC"
	rows, err := database.Connection.Query(sql, repositoryID)

	if err != nil {
		fmt.Println("Error while selecting the processes by repository: ")
		fmt.Println(repositoryID)
	}

	for rows.Next() {
		rows.Scan(&harvestID, &processID, &formats, &collections, &records, &identifiers)
		rep, _ := repository.NewRepository(repID)
		process := Process{
			Formats:     formats,
			Records:     records,
			Collections: collections,
			Identifiers: identifiers,
			HarvestID:   harvestID,
			repository:  rep,
			Process:     action.NewProcess(processID),
		}
		list = append(list, &process)
	}
	return list
}

// GetProcessToken returns the token for a process by the name
func GetProcessToken(ID int, name string) string {
	var token string
	sql := "SELECT `token_" + name + "` FROM `process_harvest` WHERE process=? LIMIT 0,1"
	query, err := database.Connection.Prepare(sql)
	query.QueryRow(ID).Scan(&token)
	if err != nil {
		fmt.Println("Error when selecting the token for " + name + " inside the harvesting process:")
		fmt.Println(err)
	}
	return token
}

// DeleteProcess deletes a process
func DeleteProcess(processID int) {
	sql := "DELETE FROM `process_harvest` WHERE process=?"
	query, err := database.Connection.Prepare(sql)

	if err != nil {
		fmt.Println("Error when deleting the harvest process:")
		fmt.Println(err)
	}
	query.Exec(processID)
}

// RecoverProcess gets the information about the current process and recovers it
func RecoverProcess(process *Process, controller Controller) *Process {
	process.controller = controller
	return process
}
