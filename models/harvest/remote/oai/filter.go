package oai

import (
	"regexp"

	"github.com/cristian-sima/Wisply/models/repository"
)

type filter struct {
	repository                repository.Filter
	rejectedRecordsIdentifier []string
}

// It checks if the keys of the records are the same as the filter one
func (filter *filter) isRecordAllowed(record Record) bool {
	if !filter.repository.IsActive() {
		return true
	}
	for _, key := range record.Keys.Identifiers {
		value := key
		regex := filter.repository.GetStructure().Harvest.Records.Reject.Identifier + "(?s)"
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
	if !filter.repository.IsActive() {
		return true
	}
	for _, element := range filter.rejectedRecordsIdentifier {
		if element == identifier.Identifier {
			return false
		}
	}
	return true
}

func newFilter(repFilter repository.Filter) filter {
	return filter{
		repository: repFilter,
	}
}
