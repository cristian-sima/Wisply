// Package data contains the functionality for exporting data
package data

const (
	numberOfValidHoursForFile = 3600 * 24
	pathToFolder              = "cache/developer/tables/"
)

var messages = map[string]string{
	"tableNotAllowed": "Wisply does not know this table :(",
}
