package auth

import (
	"github.com/astaxie/beego/orm"
)

type Account struct {
	Id            int
	Username      string
	Password      string
	Email         string
	Administrator bool
}

func NewAccount(id string) Account {
	var account Account
	db := orm.NewOrm()

	db.Raw("SELECT id, username, password, administrator FROM account").QueryRow(&account)
	return account
}

func (this *Account) IsAdministrator() bool {
	return this.Administrator
}
