package adapter

import (
	"strconv"

	validity "github.com/cristian-sima/validity"
)

// Validate It validates a set of data (parameter data) against the rules
func Validate(rawData map[string]interface{}, rules validity.ValidationRules) *validity.ValidationResults {
	return validity.ValidateMap(rawData, rules)
}

// WisplyError It represents the error message which is delivered by a model to a constructor
// It can contain a simple message in the Message field or it can encapsulets an entire validation result
type WisplyError struct {
	Data    map[string][]string
	Message string
}

// GetMessage It returns the message of the error message.
// In case it is not only a string message, it checks how many error there are.
func (err *WisplyError) GetMessage() string {
	number := len(err.Data)
	if number == 0 {
		return err.Message
	}
	tempMsg := "Your request was not successful. There were problems with "
	correctForm := ""
	if number == 1 {
		correctForm = "one field"
	} else {
		correctForm = strconv.Itoa(number) + " fields"
	}
	return tempMsg + correctForm + ":"
}
