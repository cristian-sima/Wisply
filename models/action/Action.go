package action

import (
	"fmt"
	"time"
)

// Actioner ... defines the set of methods to be implemented by an action
type Actioner interface {
	Go()
	finish()
	updateInDatabase()
}

// Action is the most basic type. It has a starting and ending timestamps and a content
type Action struct {
	Actioner
	ID        int
	Start     int64
	End       int64
	IsRunning bool
	Content   string
	result    string // it can be: error, warning, success, normal
}

// GetStartDate returns the start date of the action in a human readable form
func (action *Action) GetStartDate() string {
	return action.getDate(action.Start)
}

// GetEndDate returns the end date of the action in a human readable form
func (action *Action) GetEndDate() string {
	if action.IsRunning || action.End == 0 {
		return "Not yet"
	}
	return action.getDate(action.End)
}

// GetDuration returns the duration between the start and end or some dots if it did not finished
func (action *Action) GetDuration() string {
	if action.IsRunning || action.End == 0 {
		return "..."
	}
	var (
		startTime time.Time
		endTime   time.Time
		duration  string
	)
	startTime = time.Unix(action.Start, 0)
	endTime = time.Unix(action.End, 0)

	duration = endTime.Sub(startTime).String()

	return duration
}

func (action *Action) getDate(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(time.Stamp)
}

// ChangeResult checks if the result is valid and it changes it
func (action *Action) ChangeResult(result string) {
	if result != "error" &&
		result != "warning" &&
		result != "success" &&
		result != "normal" {
		fmt.Println("Task change result error. This result is not valid: " + result)
	}
	action.result = result
}

// Finish terminates normal an action
func (action *Action) Finish() {
	action.IsRunning = false
	action.End = getCurrentTimestamp()
}

// GetResult returns the result of the task
func (action *Action) GetResult() string {
	return action.result
}
