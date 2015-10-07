package action

import (
	"fmt"
	"strconv"

	wisply "github.com/cristian-sima/Wisply/models/database"
	repository "github.com/cristian-sima/Wisply/models/repository"
)

// GetAllProcesses returns a list with all available processes
func GetAllProcesses() []*Process {
	var list []*Process
	processFields := "process.id, process.content, process.start, process.end, process.is_running"
	repositoryFields := "repository.id, repository.name"
	fieldList := processFields + ", " + repositoryFields
	sql := "SELECT " + fieldList + " FROM `process` AS process JOIN `repository` AS `repository` ON repository.id = process.repository ORDER BY process.id DESC"

	rows, err := wisply.Database.Query(sql)
	if err != nil {
		fmt.Println("Problem when getting all the processes: ")
		fmt.Println(err)
	}

	var (
		ID, repID        int
		start, end       int64
		isRunningString  string
		isRunning        bool
		content, repName string
		rep              *repository.Repository
	)

	for rows.Next() {
		rows.Scan(&ID, &content, &start, &end, &isRunningString, &repID, &repName)

		rep = &repository.Repository{
			ID:   repID,
			Name: repName,
		}

		isRunning, err = strconv.ParseBool(isRunningString)

		if err != nil {
			fmt.Println(err)
		}

		list = append(list, &Process{
			ID:         ID,
			Repository: rep,
			Action: &Action{
				IsRunning: isRunning,
				Start:     start,
				End:       end,
				Content:   content,
			},
		})
	}
	return list
}
