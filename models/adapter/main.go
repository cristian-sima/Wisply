// Package adapter contains the functionality for adapting 3rd party modules to wisply
package adapter

import (
	"strconv"

	validity "local-projects/validity"
)

// Validate represents an adapter for validity package
// It validates a set of data (parameter data) against the rules
func Validate(rawData map[string]interface{}, rules validity.ValidationRules) *validity.ValidationResults {
	return validity.ValidateMap(rawData, rules)
}

// WisplyError represents the error message which is delivered by a model to a constructor
// It can contain a simple message in the Message field
// or it can encapsulets an entire validation result
type WisplyError struct {
	Data    map[string][]string
	Message string
}

// GetMessage returns the message of the error message.
// In case it is not only a string message, it checks how many error there are.
func (err *WisplyError) GetMessage() string {
	if len(err.Data) == 0 {
		return err.getSimpleMessage()
	}
	return err.getFullMessage()
}

func (err *WisplyError) getSimpleMessage() string {
	defaultMessage := "There was a problem"
	if err.Message == "" {
		return defaultMessage
	}
	return err.Message
}

func (err *WisplyError) getFullMessage() string {
	number := len(err.Data)
	tempMsg := "Your request was not successful. There were problems with "
	correctForm := ""
	if number == 1 {
		correctForm = "one field"
	} else {
		correctForm = strconv.Itoa(number) + " fields"
	}
	return tempMsg + correctForm + ":"
}
