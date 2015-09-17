package auth

import (
	"errors"
	"strconv"

	wisply "github.com/cristian-sima/Wisply/models/database"
	"github.com/nu7hatch/gouuid"
)

// Account It represents an account
type Account struct {
	ID              int
	Name            string
	Password        string
	Email           string
	IsAdministrator bool
}

// ChangeType It changes the type of the account
func (account *Account) ChangeType(isAdministrator string) error {

	result := isValidAdminType(isAdministrator)
	if !result.IsValid {
		return errors.New("Not a valid type")
	}

	err := account.modifyStatus(isAdministrator)
	return err
}

func (account *Account) modifyStatus(isAdministrator string) error {

	stmt, err := wisply.Database.Prepare("UPDATE `account` SET administrator=? WHERE id=?")
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(isAdministrator, strconv.Itoa(account.ID))
	if err != nil {
		panic(err)
	}
	return err
}

// GenerateConnectionCookie It generates a new connection cookie
func (account *Account) GenerateConnectionCookie() *Cookie {

	var timestamp, value string
	temp, _ := uuid.NewV4()
	plain := temp.String()
	value = getSHA1Digest(plain)

	timestamp = getCurrentTimestamp()

	sql := "INSERT INTO `account_token` (`id`, `account`, `value`, `timestamp`) VALUES (?, ?, ?, ?)"
	query, err := wisply.Database.Prepare(sql)

	if err != nil {
		panic(err)
	}

	_, err = query.Exec("NULL", strconv.Itoa(account.ID), value, timestamp)

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

// Delete It delets the account
func (account *Account) Delete() error {

	sql := "DELETE from `account` WHERE id=?"

	query, _ := wisply.Database.Prepare(sql)
	_, err := query.Exec(strconv.Itoa(account.ID))

	if err != nil {
		panic(err)
	}
	return err
}
