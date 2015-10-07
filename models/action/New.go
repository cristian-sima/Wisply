package action

import (
	"fmt"

	wisply "github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/repository"
)

// NewOperation returns an operation from database specified by ID
// NOTE! It returns just the ID of both process and task.
// In order to get the entire object use NewProcess or NewTask
func NewOperation(ID int) *Operation {

	var (
		opID, proID, taskID int
		opStart, opEnd      int64
		opIsRunning         bool
		opContent           string
	)

	fieldList := "`id`, `process`, `current_task`, `content`, `start`, `end`, `is_running`"
	sql := "SELECT " + fieldList + " FROM `operation` WHERE id= ?"
	query, err := wisply.Database.Prepare(sql)

	query.QueryRow(ID).Scan(&opID, &proID, &taskID, &opContent, &opStart, &opEnd, &opIsRunning)

	if err != nil {
		fmt.Println("It has been an error when tring to get the info about the operation: ")
		fmt.Println(err)
	}

	operation := &Operation{
		Action: &Action{
			ID:        opID,
			Start:     opStart,
			End:       opEnd,
			IsRunning: opIsRunning,
			Content:   opContent,
		},
		Process: &Process{
			Action: &Action{
				ID: proID,
			},
		},
		CurrentTask: &Task{
			Action: &Action{
				ID: taskID,
			},
		},
	}

	return operation
}

// NewTask returns a task from database specified by ID
// NOTE! It returns just the ID of the operation
// In order to get the entire object use NewOperation
func NewTask(ID int) *Task {

	var (
		taskID, opeID           int
		taskStart, taskEnd      int64
		taskIsRunning           bool
		taskStatus, taskContent string
	)

	fieldList := "`id`, `operation`, `content`, `start`, `end`, `is_running`, `status`"
	sql := "SELECT " + fieldList + " FROM `task` WHERE id= ?"
	query, err := wisply.Database.Prepare(sql)

	query.QueryRow(ID).Scan(&taskID, &opeID, &taskContent, &taskStart, &taskEnd, &taskIsRunning, &taskStatus)

	if err != nil {
		fmt.Println("It has been an error when tring to get the info about the task: ")
		fmt.Println(err)
	}

	task := &Task{
		Action: &Action{
			ID:        taskID,
			Start:     taskStart,
			End:       taskEnd,
			IsRunning: taskIsRunning,
			Content:   taskContent,
		},
		status: taskStatus,
		Operation: &Operation{
			Action: &Action{
				ID: taskID,
			},
		},
	}

	return task
}

// NewProcess returns the process by ID
// NOTE! It returns just the ID of the current task.
// In order to get the entire object use NewOperation
func NewProcess(ID int) *Process {

	var (
		opeID             int
		proStart, proEnd  int64
		proIsRunning      bool
		proContent, repID string
	)

	fieldList := "`current_operation`, `content`, `start`, `end`, `is_running`, `repository`"
	sql := "SELECT " + fieldList + " FROM `process` WHERE id= ?"
	query, err := wisply.Database.Prepare(sql)

	query.QueryRow(ID).Scan(&opeID, &proContent, &proStart, &proEnd, &proIsRunning, &repID)

	if err != nil {
		fmt.Println("It has been an error when tring to get the info about the process: ")
		fmt.Println(err)
	}

	repository, _ := repository.NewRepository(repID)

	process := &Process{
		Action: &Action{
			ID:        ID,
			Start:     proStart,
			End:       proEnd,
			IsRunning: proIsRunning,
			Content:   proContent,
		},
		Repository: repository,
		currentOperation: &Operation{
			Action: &Action{
				ID: opeID,
			},
		},
	}

	return process
}
