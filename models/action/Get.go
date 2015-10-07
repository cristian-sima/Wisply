package action

import (
	"fmt"

	wisply "github.com/cristian-sima/Wisply/models/database"
	repository "github.com/cristian-sima/Wisply/models/repository"
)

// GetAllProcesses returns a list with all available processes
func GetAllProcesses() []*Process {
	var list []*Process
	fieldList := "process.id, process.content, process.start, process.end, repository.id, repository.name"
	sql := "SELECT " + fieldList + " FROM `process` AS process JOIN `repository` AS `repository` ON repository.id = process.repository ORDER BY process.id DESC"

	rows, err := wisply.Database.Query(sql)
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var (
			ID, repID        int
			start, end       int64
			isRunning        bool
			content, repName string
			rep              *repository.Repository
		)
		rows.Scan(&ID, &content, &start, &end, &repID, &repName)

		rep = &repository.Repository{
			ID:   repID,
			Name: repName,
		}

		list = append(list, &Process{
			ID:         ID,
			Repository: rep,
			Action: &Action{
				Start:   start,
				End:     end,
				Content: content,
			},
		})
	}
	return list
}
