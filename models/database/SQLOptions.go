package database

// SQLOptions can be passed to a SQL in order to change the parameters
type SQLOptions struct {
	Limit   string
	OrderBy string
}

// GetLimit returns the SQL limit option
func (options SQLOptions) GetLimit() string {
	return "LIMIT " + options.Limit
}

// GetOrder returns the Order By option
func (options SQLOptions) GetOrder() string {
	return "ORDER BY " + options.OrderBy
}
