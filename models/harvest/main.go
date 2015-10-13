package harvest

import (
	"fmt"
	"strconv"

	"github.com/cristian-sima/Wisply/models/action"
	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/harvest/remote"
	"github.com/cristian-sima/Wisply/models/repository"
)

// RecoverProcess gets the information about the current process and recovers it
func RecoverProcess(process *Process, controller Controller) *Process {
	process.controller = controller
	return process
}

// CreateProcess creates a new harvest process
func CreateProcess(ID string, controller Controller) *Process {

	process := &*action.CreateProcess("Harvest")

	harvestProcess := buildProcess(ID, controller, process)

	harvestID := insertHarvestProcess(harvestProcess)

	harvestProcess.HarvestID = harvestID

	return harvestProcess
}

func buildProcess(ID string, controller Controller, process *action.Process) *Process {

	local, _ := repository.NewRepository(ID)

	harvestProcess := &Process{
		Process:    process,
		remote:     getRemoteServer(local),
		controller: controller,
		repository: local,
	}

	return harvestProcess
}

func getRemoteServer(local *repository.Repository) remote.RepositoryInterface {

	var rem remote.RepositoryInterface

	switch local.Category {
	case "EPrints":
		{
			rem = remote.NewEPrints(local)
		}
	}

	return rem
}

func insertHarvestProcess(process *Process) int {

	var harvestID int

	columns := "(`process`, `repository`)"
	values := "(?, ?)"
	sql := "INSERT INTO `process_harvest` " + columns + " VALUES " + values

	query, err := database.Connection.Prepare(sql)

	if err != nil {
		fmt.Println("Error when creating the harvest process:")
		fmt.Println(sql)
		fmt.Println(err)
	}
	query.Exec(process.ID, process.GetRepository().ID)

	// find its ID
	sql = "SELECT `id` FROM `process_harvest` WHERE process=? AND repository=? ORDER by id LIMIT 0,1"
	query, err = database.Connection.Prepare(sql)
	query.QueryRow(process.ID, process.GetRepository().ID).Scan(&harvestID)

	if err != nil {
		fmt.Println("Error when selecting the harvest id:")
		fmt.Println(err)
	}

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

	return &Process{
		Formats:     formats,
		Records:     records,
		Collections: collections,
		Identifiers: identifiers,
		HarvestID:   harvestID,
		repository:  local,
		remote:      getRemoteServer(local),
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
	sql := "SELECT `token_" + name + "` FROM `process_harvest` WHERE id=? LIMIT 0,1"
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
