package auth

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id            int
	Username      string
	Password      string
	Email         string
	Administrator bool
}

func NewUser() User {
	var user User
	orm := orm.NewOrm()

	orm.Raw("SELECT username, password, administrator FROM user").QueryRow(&user)

	return user
}
