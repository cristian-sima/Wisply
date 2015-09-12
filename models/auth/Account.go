package auth

import (
	"github.com/astaxie/beego/orm"
)

type Account struct {
	Id            int
	Name          string
	Password      string
	Email         string
	Administrator bool
}

func NewAccount(id string) Account {
	var account Account
	db := orm.NewOrm()

	db.Raw("SELECT id, name, password, administrator FROM account").QueryRow(&account)
	return account
}

func (this *Account) IsAdministrator() bool {
	return this.Administrator
}
