package auth

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Username string
	Password string
	isAdmin  bool
}

func (user *User) isAdministrator() bool {
	return user.isAdmin
}

func NewUser() User {
	var user User
	orm := orm.NewOrm()

	orm.Raw("SELECT username, password, isAdmin FROM user").QueryRow(&user)

	return user
}
