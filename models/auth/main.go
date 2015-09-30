package auth

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"strconv"
	"time"

	wisply "github.com/cristian-sima/Wisply/models/database"
)

// Settings The authentication's settings
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

// ReconnectUsingCookie It tries to reconnect the user using the value from the connection cookiePath
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

	query, _ := wisply.Database.Prepare("DELETE from account_token WHERE timestamp < ?")
	query.Exec(strconv.Itoa(diff))
}

// GetAllAccounts It returns an array of Account with all the accounts
func (model *Model) GetAllAccounts() []Account {
	var list []Account
	sql := "SELECT id, name, password, email, administrator FROM account"
	rows, _ := wisply.Database.Query(sql)
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
	query, _ := wisply.Database.Prepare("SELECT count(*) FROM account")
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

	sql := "SELECT id, name, password, email, administrator FROM account WHERE email = ? "
	query, err := wisply.Database.Prepare(sql)
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
	account := new(Account)

	sql := "SELECT id, name, password, email, administrator FROM account WHERE id= ?"
	query, err := wisply.Database.Prepare(sql)
	query.QueryRow(ID).Scan(&account.ID, &account.Name, &account.Password, &account.Email, &account.IsAdministrator)

	if err != nil {
		return account, errors.New("Error")
	}

	return account, nil
}

func getSHA1Digest(plainToken string) string {
	array := []byte(plainToken)
	hasher := sha1.New()
	hasher.Write(array)
	token := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return token
}
