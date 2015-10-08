package action

import (
	"strings"
	"time"
)

const (
	// DateFormat is the format for the dates
	DateFormat = time.Stamp
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
	var duration string
	if action.IsRunning || action.End == 0 {
		return "..."
	}
	duration = action.getDuration(action.Start, action.End)

	// make it nice
	duration = action.getNiceDuration(duration)

	return duration
}

func (action *Action) getDuration(start, end int64) string {
	var (
		startTime time.Time
		endTime   time.Time
		duration  string
	)
	startTime = time.Unix(start, 0)
	endTime = time.Unix(end, 0)

	duration = endTime.Sub(startTime).String()
	return duration
}

func (action *Action) getNiceDuration(duration string) string {

	duration = strings.Replace(duration, "h", "h, ", -1)
	duration = strings.Replace(duration, "m", " min, ", -1)
	duration = strings.Replace(duration, "s", " sec", -1)

	return duration
}

func (action *Action) getDate(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(DateFormat)
}

// ChangeResult checks if the result is valid and it changes it
func (action *Action) ChangeResult(result string) {
	if !action.isValidResult(result) {
		panic("Action change result error. This result is not valid: " + result)
	}
	action.result = result
}

func (action *Action) isValidResult(result string) bool {
	return result == "danger" ||
		result == "warning" ||
		result == "success" ||
		result == "normal"
}

// Finish terminates normal an action
func (action *Action) Finish() {
	action.IsRunning = false
	action.End = getCurrentTimestamp()
}

// GetResult returns the result of the action
func (action *Action) GetResult() string {
	return action.result
}
