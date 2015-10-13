package wisply

// TODO Keys

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
	var ret []string
	_, exists := keys.data[name]
	if exists {
		ret = keys.data[name]
	}
	return ret
}
