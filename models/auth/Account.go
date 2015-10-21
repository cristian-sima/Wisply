// Package auth contains all the functions for creating,
// login and logout an account
package auth

import (
	"errors"
	"strconv"

	"github.com/cristian-sima/Wisply/models/database"
)

// Account represents an account
type Account struct {
	ID              int
	Name            string
	Password        string
	Email           string
	IsAdministrator bool
}

// ChangeType changes the type of the account
func (account *Account) ChangeType(isAdministrator string) error {

	result := isValidAdminType(isAdministrator)
	if !result.IsValid {
		return errors.New("Not a valid type")
	}

	err := account.modifyStatus(isAdministrator)
	return err
}

// It modifies the status of the user
func (account *Account) modifyStatus(isAdministrator string) error {
	sql := "UPDATE `account` SET administrator=? WHERE id=?"
	stmt, err := database.Connection.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(isAdministrator, strconv.Itoa(account.ID))
	if err != nil {
		return err
	}
	return err
}

// GenerateConnectionCookie generates a new "connection" cookie
// This cookie is used to remember the coonection of the account
func (account *Account) GenerateConnectionCookie() *Cookie {
	token := account.createNewToken()
	cookie := Cookie{
		Account:  account,
		Token:    token,
		Duration: Settings["duration"].(int),
		Path:     Settings["path"].(string),
	}
	return &cookie
}

func (account *Account) createNewToken() *Token {
	token := generateToken(account)
	token.Insert()
	return token
}

// Delete removes the account from the database
func (account *Account) Delete() error {
	sql := "DELETE from `account` WHERE id=?"
	query, _ := database.Connection.Prepare(sql)
	_, err := query.Exec(strconv.Itoa(account.ID))

	if err != nil {
		panic(err)
	}
	return err
}
