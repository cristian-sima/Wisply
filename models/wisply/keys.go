package wisply

import (
	"strings"

	wisplyFormat "github.com/cristian-sima/Wisply/models/wisply/format"
)

// RecordKeys holds all the keys for the record
type RecordKeys struct {
	Data map[string][]string
}

// Add inserts a new value for a name
func (keys *RecordKeys) Add(name, value string) {
	if keys.Data == nil {
		keys.Data = make(map[string][]string)
	}
	_, exists := keys.Data[name]
	if !exists {
		keys.Data[name] = []string{value}
	} else {
		keys.Data[name] = append(keys.Data[name], value)
	}
}

// Get returns all the keys for a name
func (keys *RecordKeys) Get(name string) []string {
	return keys.get(name, false)
}

// Nice returns all the keys for a name and processes them
func (keys *RecordKeys) Nice(name string) []string {
	return keys.get(name, true)
}

// Get returns all the keys for a name
func (keys *RecordKeys) get(name string, processThem bool) []string {
	var ret []string
	_, exists := keys.Data[name]
	if exists {
		ret = keys.Data[name]
	}
	if processThem {
		return keys.processKeys(name, ret)
	}
	return ret
}

func (keys *RecordKeys) processKeys(name string, values []string) []string {
	switch name {
	case "format":
		return wisplyFormat.ConvertFormats(values)
	}
	return values
}

// ProcessFormats returns the formats in a wisply form
func (keys *RecordKeys) ProcessFormats() map[string]int {
	values := keys.get("format", false)
	return Compress(wisplyFormat.ConvertFormats(values))
}

// GetTitle returns all the title of the string as a string
func (keys *RecordKeys) GetTitle() string {
	values := keys.get("title", false)
	title := ""
	for _, currentTitle := range values {
		if len(title) != 0 {
			title += ". "
		}
		title += currentTitle
	}
	return title
}

// GetURL returns the first relation which starts with http
func (keys *RecordKeys) GetURL() string {
	values := keys.get("relation", false)
	for _, relation := range values {
		if strings.HasPrefix(relation, "http") {
			return relation
		}
	}
	return ""
}

// Compress compacts and counts the values
func Compress(values []string) map[string]int {
	compressed := make(map[string]int)
	for _, value := range values {
		_, exists := compressed[value]
		if !exists {
			compressed[value] = 1
		} else {
			compressed[value]++
		}
	}
	return compressed
}
