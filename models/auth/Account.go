package auth

import (
	. "github.com/cristian-sima/Wisply/models/wisply"
	"github.com/nu7hatch/gouuid"
	"strconv"
)

type Account struct {
	Id            int
	Name          string
	Password      string
	Email         string
	Administrator bool
}

func (this *Account) IsAdministrator() bool {
	return this.Administrator
}

func (account *Account) ChangeType(isAdministrator string) error {

	// to do to check type

	stringElements := []string{
		isAdministrator,
		strconv.Itoa(account.Id),
	}
	_, err := Database.Raw("UPDATE `account` SET administrator=? WHERE id=?", stringElements).Exec()
	return err
}

func (account *Account) GenerateConnectionCookie() *Cookie {

	var timestamp, value string
	temp, _ := uuid.NewV4()
	value = temp.String()

	timestamp = GetCurrentTimestamp()

	elementsInsert := []string{
		"NULL",
		strconv.Itoa(account.Id),
		value,
		timestamp,
	}
	Database.Raw("INSERT INTO `account_login` (`id`, `account`, `token`, `timestamp`) VALUES (?, ?, ?, ?)", elementsInsert).Exec()

	intTimestamp, _ := strconv.Atoi(timestamp)
	token := Token{
		Value:     value,
		Timestamp: intTimestamp,
	}

	cookie := Cookie{
		Account:  account,
		Token:    &token,
		Duration: settings["duration"].(int),
		Path:     settings["path"].(string),
	}

	return &cookie
}

func (account *Account) Delete() error {
	elememts := []string{
		strconv.Itoa(account.Id),
	}
	_, err := Database.Raw("DELETE from `account` WHERE id=?", elememts).Exec()
	return err
}
