package searches

import "time"

// Search is a typical search performed by a user
type Search struct {
	ID        int
	Accessed  bool
	Query     string
	Timestamp int64
}

// GetDate transforms the timestamp into a date
func (search Search) GetDate() string {
	return time.Unix(search.Timestamp, 0).Format(dateFormat)
}
