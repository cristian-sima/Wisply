package auth

import (
	"errors"
	"strconv"

	wisply "github.com/cristian-sima/Wisply/models/wisply"
	"github.com/nu7hatch/gouuid"
)

// Account It represents an account
type Account struct {
	ID            int
	Name          string
	Password      string
	Email         string
	Administrator bool
}

// IsAdministrator It checks if it the account is an administrator
func (account *Account) IsAdministrator() bool {
	return account.Administrator
}

// ChangeType It changes the type of the account
func (account *Account) ChangeType(isAdministrator string) error {

	result := isValidAdminType(isAdministrator)
	if !result.IsValid {
		return errors.New("Not a valid type")
	}

	stringElements := []string{
		isAdministrator,
		strconv.Itoa(account.ID),
	}
	_, err := wisply.Database.Raw("UPDATE `account` SET administrator=? WHERE id=?", stringElements).Exec()
	return err
}

// GenerateConnectionCookie It generates a new connection cookie
func (account *Account) GenerateConnectionCookie() *Cookie {

	var timestamp, value string
	temp, _ := uuid.NewV4()
	plain := temp.String()
	value = getSHA1Digest(plain)

	timestamp = getCurrentTimestamp()

	elementsInsert := []string{
		"NULL",
		strconv.Itoa(account.ID),
		value,
		timestamp,
	}
	_, err := wisply.Database.Raw("INSERT INTO `account_token` (`id`, `account`, `value`, `timestamp`) VALUES (?, ?, ?, ?)", elementsInsert).Exec()

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
	elememts := []string{
		strconv.Itoa(account.ID),
	}
	_, err := wisply.Database.Raw("DELETE from `account` WHERE id=?", elememts).Exec()
	return err
}
