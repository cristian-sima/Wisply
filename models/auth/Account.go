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
	plain := temp.String()
	value = getSHA1_digest(plain)

	timestamp = GetCurrentTimestamp()


	elementsInsert := []string{
		"NULL",
		strconv.Itoa(account.Id),
		value,
		timestamp,
	}
	_, err := Database.Raw("INSERT INTO `account_token` (`id`, `account`, `value`, `timestamp`) VALUES (?, ?, ?, ?)", elementsInsert).Exec()

	if err != nil {
		panic(err)
	}

	intTimestamp, _ := strconv.Atoi(timestamp)
	token := Token{
		Value:     plain,
		Timestamp: intTimestamp,
	}

	cookie := Cookie{
		Account:  account,
		Token:    &token,
		Duration: Settings["duration"].(int),
		Path:     Settings["path"].(string),
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
