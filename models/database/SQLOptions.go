package database

import (
	"errors"
	"strconv"
)

// Temp is used to store temp data before it is verified
type Temp struct {
	LimitMin string
	LimitMax string
	OrderBy  string
	Limit    int // used to limit the maximum number of resources
}

// SQLOptions represents a valid SQL option object
type SQLOptions struct {
	*Temp
}

// GetLimit returns the SQL limit option
func (options *SQLOptions) GetLimit() string {
	if options.LimitMin == "" && options.LimitMax == "" {
		return ""
	}
	return "LIMIT " + options.LimitMin + "," + options.LimitMax
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
	isValid := areValidLimits(options.LimitMin, options.LimitMax)

	if !isValid {
		return false
	}

	min, _ := strconv.Atoi(options.LimitMin)
	max, _ := strconv.Atoi(options.LimitMax)

	return (max-min <= options.Limit)
}
