package action

import (
	"fmt"

	database "github.com/cristian-sima/Wisply/models/database"
)

// - constructors

// NewAction creates a new action
func NewAction(isRunning bool, content string) *Action {
	return &Action{
		IsRunning: isRunning,
		Start:     getCurrentTimestamp(),
		Content:   content,
		result:    "normal",
	}
}

// NewOperation returns an operation from database specified by ID
// NOTE! It returns just the ID of both process and task.
// In order to get the entire object use NewProcess or NewTask
func NewOperation(ID int) *Operation {

	var (
		opID, proID, taskID int
		opStart, opEnd      int64
		opIsRunning         bool
		opContent           string
		result              string
	)

	fieldList := "`id`, `process`, `current_task`, `content`, `start`, `end`, `is_running`, `result`"
	sql := "SELECT " + fieldList + " FROM `operation` WHERE id= ?"
	query, err := database.Connection.Prepare(sql)

	query.QueryRow(ID).Scan(&opID, &proID, &taskID, &opContent, &opStart, &opEnd, &opIsRunning, &result)

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
			result:    result,
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
		taskID, opeID                            int
		taskStart, taskEnd                       int64
		taskIsRunning                            bool
		taskResult, taskContent, taskExplication string
	)

	fieldList := "`id`, `operation`, `content`, `start`, `end`, `is_running`, `result`, `explication`"
	sql := "SELECT " + fieldList + " FROM `task` WHERE id= ?"
	query, err := database.Connection.Prepare(sql)

	if err != nil {
		fmt.Println("It has been an error when tring to get the info about the task: ")
		fmt.Println(err)
	}

	query.QueryRow(ID).Scan(&taskID, &opeID, &taskContent, &taskStart, &taskEnd, &taskIsRunning, &taskResult, &taskExplication)

	task := &Task{
		Explication: taskExplication,
		Action: &Action{
			ID:        taskID,
			Start:     taskStart,
			End:       taskEnd,
			IsRunning: taskIsRunning,
			Content:   taskContent,
			result:    taskResult,
		},
		Operation: &Operation{
			Action: &Action{
				ID: taskID,
			},
		},
	}

	return task
}

// NewProcess returns the process by ID
// NOTE! It returns just the ID of the current operation.
// In order to get the entire object use NewOperation
func NewProcess(ID int) *Process {

	var (
		opeID            int
		proStart, proEnd int64
		proIsRunning     bool
		proContent       string
	)

	fieldList := "`current_operation`, `content`, `start`, `end`, `is_running`"
	sql := "SELECT " + fieldList + " FROM `process` WHERE id= ?"
	query, err := database.Connection.Prepare(sql)

	query.QueryRow(ID).Scan(&opeID, &proContent, &proStart, &proEnd, &proIsRunning)

	if err != nil {
		fmt.Println("It has been an error when tring to get the info about the process: ")
		fmt.Println(err)
	}

	process := &Process{
		Action: &Action{
			ID:        ID,
			Start:     proStart,
			End:       proEnd,
			IsRunning: proIsRunning,
			Content:   proContent,
		},
		currentOperation: &Operation{
			Action: &Action{
				ID: opeID,
			},
		},
	}

	return process
}
