package oai

import (
	"encoding/json"
	"fmt"
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
	data     structure
	isActive bool
}

func (filter *filter) isRecordAllowed(record Record) bool {
	if !filter.isActive {
		return true
	}
	for _, key := range record.Keys.Identifiers {
		value := key
		regex := filter.data.Harvest.Records.Reject.Identifier + "(?s)"
		matched, _ := regexp.MatchString(regex, value)
		if matched {
			fmt.Println("Filter rejected: " + key)
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
