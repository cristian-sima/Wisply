package wisply

import "strconv"

// Record is the wisply record
type Record struct {
	ID         int `json:"ID"`
	Identifier string
	Timestamp  string
	Keys       *RecordKeys
	Repository int
}

// GetIdentifier returns the unique identifier
func (record *Record) GetIdentifier() string {
	return record.Identifier
}

// GetTimestamp returns the timestamp of the record
func (record *Record) GetTimestamp() string {
	return record.Timestamp
}

// GetWisplyURL returns the local address of the resource within Wisply
func (record *Record) GetWisplyURL() string {
	return "/repository/" + strconv.Itoa(record.Repository) + "/resource/" + strconv.Itoa(record.ID)
}
