package adapter

import (
	validity "github.com/cristian-sima/validity"
	"strconv"
)

func Validate(rawData map[string]interface{}, rules validity.ValidationRules) *validity.ValidationResults {
	return validity.ValidateMap(rawData, rules)
}

type WisplyError struct {
	Data    map[string][]string
	Message string
}

func (err *WisplyError) GetMessage() string {
	number := len(err.Data)
	if number == 0 {
		return err.Message
	} else {
		tempMsg := "Your request was not successful. There were problems with "
		correctForm := ""
		if number == 1 {
			correctForm = "one field"
		} else {
			correctForm = strconv.Itoa(number) + " fields"
		}
		return tempMsg + correctForm + ":"
	}
}
