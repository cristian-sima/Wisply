package auth

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/cristian-sima/Wisply/models/database"
)

// Settings is the object which holds the authentication's settings
var Settings = map[string]interface{}{
	"duration":        (60 * 60 * 24 * 7 * 4), // one month
	"path":            "/",
	"separatorCookie": "::",
	"cookieName":      "connection",
	"cookiePath":      "/",
}

// Model It encapsulates all the main operations for authentication
type Model struct {
}

// ReconnectUsingCookie It tries to reconnect the user using the value
// from the connection cookiePath
// The value is splited in 2 values: ID and hashedToken
// Then the values are verified in the database
// Finally, it is checked if the token is still valid
func ReconnectUsingCookie(plainCookie string) (string, error) {
	cookie, err := newLoginCookie(plainCookie)
	if err != nil {
		return "", errors.New("The cookie has an invalid format")
	}
	validToken := cookie.IsGood()
	if !validToken {
		return "", errors.New("The token is not valid")
	}
	deleteOldTokens()
	return cookie.AccountID, nil
}

func deleteOldTokens() {
	now, _ := strconv.Atoi(getCurrentTimestamp())
	duration := Settings["duration"].(int)
	diff := now - duration
	sql := "DELETE from account_token WHERE timestamp < ?"
	query, _ := database.Connection.Prepare(sql)
	query.Exec(strconv.Itoa(diff))
}

// GetAllAccounts It returns an array of Account with all the accounts
func (model *Model) GetAllAccounts() []Account {
	var list []Account
	sql := "SELECT id, name, password, email, administrator FROM account"
	rows, _ := database.Connection.Query(sql)
	for rows.Next() {
		account := Account{}
		rows.Scan(&account.ID, &account.Name, &account.Password, &account.Email, &account.IsAdministrator)
		list = append(list, account)
	}
	return list
}

// CountAccounts It returns the number of accounts
func CountAccounts() int {
	number := 0
	sql := "SELECT count(*) FROM account"
	query, _ := database.Connection.Prepare(sql)
	query.QueryRow().Scan(&number)
	return number
}

func getCurrentTimestamp() string {
	var timestamp string
	unixTime := time.Now().Unix()
	timestamp = strconv.FormatInt(unixTime, 10)
	return timestamp
}

// GetAccountByEmail It searches and returns the account with that email
func GetAccountByEmail(email string) (*Account, error) {
	account := Account{}
	id := isValidEmail(email)
	if !id.IsValid {
		return &account, errors.New("The id is not valid")
	}
	fieldList := "id, name, password, email, administrator"
	sql := "SELECT " + fieldList + " FROM account WHERE email = ? "
	query, err := database.Connection.Prepare(sql)
	query.QueryRow(email).Scan(&account.ID, &account.Name, &account.Password, &account.Email, &account.IsAdministrator)

	if err != nil {
		return &account, errors.New("No such account")
	}

	return &account, nil
}

// NewAccount It creates an account using the ID
func NewAccount(ID string) (*Account, error) {

	result := isValidID(ID)
	if !result.IsValid {
		return nil, errors.New("The id is not valid")
	}
	account := &Account{}
	fieldList := "id, name, password, email, administrator"
	sql := "SELECT " + fieldList + " FROM account WHERE id= ?"
	query, err := database.Connection.Prepare(sql)
	query.QueryRow(ID).Scan(&account.ID, &account.Name, &account.Password, &account.Email, &account.IsAdministrator)

	if err != nil {
		return account, errors.New("Error")
	}

	return account, nil
}

// Returns the SHA1 of a string in a string format
func getSHA1Digest(plainToken string) string {
	array := []byte(plainToken)
	hasher := sha1.New()
	hasher.Write(array)
	token := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return token
}

// VerifyAccount checks if the password for an account is good or not
func VerifyAccount(account *Account, plainPassword string) bool {
	result := checkPasswordIsCorrect(account.Password, plainPassword)
	return result
}

func checkPasswordIsCorrect(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err == nil {
		return true
	}
	return false
}
