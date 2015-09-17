package WisplyModel

import (
	orm "github.com/astaxie/beego/orm"
)

// It represents the connection to the database
var (
	Database orm.Ormer
)

// IsEmptyQuery checks if a query is empty or not
func IsEmptyQuery(sql string, elements []string) bool {
	var list orm.ParamsList
	num, err := Database.Raw(sql, elements).ValuesFlat(&list)
	if err == nil && num > 0 {
		return true
	}
	return false
}

// Model encapsulates the general operations for all models
type Model struct {
}
