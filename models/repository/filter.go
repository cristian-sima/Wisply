package repository

import "encoding/json"

// Structure represents the filter's structure
type Structure struct {
	Harvest struct {
		Records struct {
			Reject struct {
				Identifier string `json:"identifier"`
			} `json:"reject"`
		} `json:"records"`
	} `json:"harvest"`
	Analyse struct {
		Reject []string `json:"Reject"`
	} `json:"analyse"`
}

// Filter represents a JSON object used to reject or allow fields
type Filter struct {
	structure Structure
	isActive  bool
}

// IsActive checks if the filter is enabled or not
func (filter Filter) IsActive() bool {
	return filter.isActive
}

// ToString returns the JSON string of the filter
func (filter Filter) ToString() string {
	rawString, err := json.Marshal(filter.structure)
	if err != nil {
		return "{}"
	}
	return string(rawString)
}

// GetStructure returns the structure of the filter
func (filter Filter) GetStructure() Structure {
	return filter.structure
}

func newFilter(raw string) Filter {
	var isActive bool
	filterStructure := Structure{}
	err := json.Unmarshal([]byte(raw), &filterStructure)
	if err == nil {
		isActive = true
	}
	return Filter{
		structure: filterStructure,
		isActive:  isActive,
	}
}
