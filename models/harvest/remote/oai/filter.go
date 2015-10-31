package oai

import (
	"encoding/json"
	"regexp"
)

type structure struct {
	Harvest struct {
		Records struct {
			Reject struct {
				Identifier string `json:"identifier"`
			} `json:"reject"`
		} `json:"records"`
	} `json:"harvest"`
}

type filter struct {
	data                      structure
	isActive                  bool
	rejectedRecordsIdentifier []string
}

// It checks if the keys of the records are the same as the filter one
func (filter *filter) isRecordAllowed(record Record) bool {
	if !filter.isActive {
		return true
	}
	for _, key := range record.Keys.Identifiers {
		value := key
		regex := filter.data.Harvest.Records.Reject.Identifier + "(?s)"
		matched, _ := regexp.MatchString(regex, value)
		if matched {
			newList := append(filter.rejectedRecordsIdentifier, record.Identifier)
			filter.rejectedRecordsIdentifier = newList
			return false
		}
	}
	return true
}

// It checks if a record with that identifier is present in the rejected list
func (filter *filter) isIdentifierAllowed(identifier Identifier) bool {
	if !filter.isActive {
		return true
	}
	for _, element := range filter.rejectedRecordsIdentifier {
		if element == identifier.Identifier {
			return false
		}
	}
	return true
}

func newFilter(raw string) filter {
	var isActive bool
	filterStructure := structure{}
	err := json.Unmarshal([]byte(raw), &filterStructure)
	if err == nil {
		isActive = true
	}
	return filter{
		data:     filterStructure,
		isActive: isActive,
	}
}
