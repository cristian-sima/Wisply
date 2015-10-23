// Package searches contains the objects which manage the accounts' searches
package searches

import (
	"strconv"
	"time"
)

func getCurrentTimestamp() string {
	var timestamp string
	unixTime := time.Now().Unix()
	timestamp = strconv.FormatInt(unixTime, 10)
	return timestamp
}
