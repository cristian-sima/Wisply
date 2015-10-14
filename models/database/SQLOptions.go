package database

import (
	"errors"
	"strconv"
)

// Temp is used to store temp data before it is verified
type Temp struct {
	LimitMin string
	Offset   string // the actual limit from the client
	OrderBy  string
	Limit    int // used to limit the maximum number of resources
}

// SQLOptions represents a valid SQL option object
type SQLOptions struct {
	*Temp
}

// GetLimit returns the SQL limit option
func (options *SQLOptions) GetLimit() string {
	if options.LimitMin == "" && options.Offset == "" {
		return ""
	}
	return "LIMIT " + options.LimitMin + "," + options.Offset
}

// GetOrder returns the Order By option
func (options *SQLOptions) GetOrder() string {
	return options.OrderBy
}

// NewSQLOptions validates and create a valid SQLOptions object
func NewSQLOptions(options Temp) (SQLOptions, error) {
	valid := SQLOptions{}
	if isValidOption(options) {
		valid.Temp = &options
		return valid, nil
	}
	return valid, errors.New("The SQL is not valid")
}

func isValidOption(options Temp) bool {
	return validateLimit(options)
}

func validateLimit(options Temp) bool {
	isValid := areValidLimits(options.LimitMin, options.Offset)

	if !isValid {
		return false
	}

	number, _ := strconv.Atoi(options.Offset)

	return (number <= options.Limit)
}
