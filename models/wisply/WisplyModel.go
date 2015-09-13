package WisplyModel

import (
	"fmt"
	orm "github.com/astaxie/beego/orm"
)

var (
	Database      orm.Ormer
)

type WisplyModel struct {
}

func IsEmptyQuery(sql string, elements []string) bool {
	var list orm.ParamsList
	num, err := Database.Raw(sql, elements).ValuesFlat(&list)
	if err == nil && num > 0 {
		return true
	}
	return false
}

func InitDatabase() {
	Database = orm.NewOrm()
	fmt.Println("initializez baza de date")
}
