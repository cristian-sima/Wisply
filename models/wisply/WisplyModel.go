package WisplyModel

import (
	orm "github.com/astaxie/beego/orm"
)

var Database orm.Ormer

type WisplyModel struct {
}

func InitDatabase() {
	Database = orm.NewOrm()
}
