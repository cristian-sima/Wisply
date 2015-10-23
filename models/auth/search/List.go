package searches

import "github.com/cristian-sima/Wisply/models/database"

// List encapsulates all the searches made by an account
type List struct {
	accountID int
}

// InsertAccessed adds a query that was accessed by the account
// It sets the timestamp
func (list List) InsertAccessed(query string) {
	list.insertQuery(query, true)
}

// InsertNotAccessed adds a query that was not accessed by account
// It sets the timestamp
func (list List) InsertNotAccessed(query string) {
	list.insertQuery(query, false)
}

// Insert adds the search into database
// It sets the timestamp
func (list List) insertQuery(query string, accessedBool bool) {
	var getAccessedString = func(accessed bool) string {
		if accessed {
			return "1"
		}
		return "0"
	}
	timestamp := getCurrentTimestamp()
	fieldList := "`query`, `timestamp`, `account`, `accessed`"
	questionList := "?, ?, ?, ?"
	accessedString := getAccessedString(accessedBool)
	sql := "INSERT INTO `account_searches` (" + fieldList + ") VALUES (" + questionList + ")"
	stmt, _ := database.Connection.Prepare(sql)
	stmt.Exec(query, timestamp, list.accountID, accessedString)
}

// Clear clears the history of a user
func (list List) Clear() {
	sql := "DELETE FROM `account_searches` WHERE account = ? "
	stmt, _ := database.Connection.Prepare(sql)
	stmt.Exec(list.accountID)
}

// GetAll returns the list of all
func (list List) GetAll() []Search {
	var allList []Search
	var getBoolFromString = func(accessed string) bool {
		if accessed == "1" {
			return true
		}
		return false
	}
	fieldList := "`id`, `query`, `timestamp`, `accessed`"
	whereClause := "WHERE `account` = ?"
	sql := "SELECT " + fieldList + "FROM `account_searches` " + whereClause
	rows, _ := database.Connection.Query(sql, list.accountID)
	for rows.Next() {
		accessedString := ""
		item := Search{}
		rows.Scan(&item.ID, &item.Query, &item.Timestamp, &accessedString)
		item.Accessed = getBoolFromString(accessedString)
		allList = append(allList, item)
	}
	return allList
}

// NewList loads a searches object for an account
func NewList(accountID int) List {
	return List{
		accountID: accountID,
	}
}
