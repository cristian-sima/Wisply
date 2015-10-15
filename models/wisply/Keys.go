package wisply

import wisplyFormat "github.com/cristian-sima/Wisply/models/wisply/format"

// RecordKeys holds all the keys for the record
type RecordKeys struct {
	data map[string][]string
}

// Add inserts a new value for a name
func (keys *RecordKeys) Add(name, value string) {
	if keys.data == nil {
		keys.data = make(map[string][]string)
	}
	_, exists := keys.data[name]
	if !exists {
		keys.data[name] = []string{value}
	} else {
		keys.data[name] = append(keys.data[name], value)
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
	_, exists := keys.data[name]
	if exists {
		ret = keys.data[name]
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
