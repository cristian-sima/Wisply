package wisply

// Record is the wisply record
type Record struct {
	ID         int `json:"ID"`
	Identifier string
	Timestamp  string
	Keys       *RecordKeys
}

// GetIdentifier returns the unique identifier
func (record *Record) GetIdentifier() string {
	return record.Identifier
}

// GetTimestamp returns the timestamp of the record
func (record *Record) GetTimestamp() string {
	return record.Timestamp
}
