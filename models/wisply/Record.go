package wisply

// Record is the wisply record
type Record struct {
	ID         int `json:"ID"`
	identifier string
	timestamp  string
	Keys       *RecordKeys
}

// GetIdentifier returns the unique identifier
func (record *Record) GetIdentifier() string {
	return record.identifier
}

// GetTimestamp returns the timestamp of the record
func (record *Record) GetTimestamp() string {
	return record.timestamp
}
