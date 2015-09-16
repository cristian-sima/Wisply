package auth

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"strconv"
	"strings"
	"time"

	wisply "github.com/cristian-sima/Wisply/models/wisply"
)

// Settings The authentication's settings
var Settings = map[string]interface{}{
	"duration":        (60 * 60 * 24 * 7), // one week
	"path":            "/",
	"separatorCookie": "::",
	"cookieName":      "connection",
	"cookiePath":      "/",
}

// Model It encapsulates all the main operations for authentication
type Model struct {
	wisply.WisplyModel
}

// ReConnect It tries to reconnect the user using the value from the connection cookiePath
// The value is splited in 2 values: ID and hashedToken
// Then the values are verified in the database
// Finally, it is checked if the token is still valid
func ReConnect(plainCookie string) (string, error) {
	ID := ""

	elements, err := splitCookie(plainCookie)
	if err != nil {
		return ID, errors.New("The cookie has an invalid format")
	}

	ID = elements[0]
	hashedToken := elements[1]

	validToken := isTokenValid(ID, hashedToken)
	if !validToken {
		return ID, errors.New("The token is not valid")
	}

	deleteOldTokens()

	return ID, nil
}

func isTokenValid(ID, hashedToken string) bool {

	elements := []string{
		ID,
		hashedToken,
	}
	token := Token{}

	if wisply.Database == nil {
		return false
	}

	sql := "SELECT value, timestamp FROM account_token WHERE account=? AND value=?"
	err := wisply.Database.Raw(sql, elements).QueryRow(&token)

	if err != nil {
		return false
	}

	now, _ := strconv.Atoi(getCurrentTimestamp())
	duration := Settings["duration"].(int)

	isValid := (now <= (token.Timestamp + duration))
	return isValid
}

func deleteOldTokens() {
	var (
		now, duration, diff int
	)
	now, _ = strconv.Atoi(getCurrentTimestamp())
	duration = Settings["duration"].(int)
	diff = now - duration

	elementsDelete := []string{
		strconv.Itoa(diff),
	}
	wisply.Database.Raw("DELETE from account_token WHERE timestamp < ?", elementsDelete).Exec()
}

// GetAllAccounts It returns an array of Account with all the accounts
func (model *Model) GetAllAccounts() []Account {
	var list []Account
	wisply.Database.Raw("SELECT id, name, password, email, administrator FROM account").QueryRows(&list)
	return list
}

// CountAccounts It returns the number of accounts
func CountAccounts() int {
	number := 0
	wisply.Database.Raw("SELECT count(*) FROM account").QueryRow(&number)
	return number
}

func getCurrentTimestamp() string {
	var timestamp string
	unixTime := time.Now().Unix()
	timestamp = strconv.FormatInt(unixTime, 10)
	return timestamp
}

func splitCookie(cookieValue string) ([]string, error) {
	var (
		toReturn  []string
		separator = Settings["separatorCookie"].(string)
	)

	validFormat := strings.Contains(cookieValue, separator)

	if !validFormat {
		return toReturn, errors.New("Not a valid format")
	}
	toReturn = strings.Split(cookieValue, "::")
	if len(toReturn) != 2 {
		return toReturn, errors.New("Not a valid format")
	}
	return toReturn, nil
}

// GetAccountByEmail It searches and returns the account with that email
func GetAccountByEmail(email string) (*Account, error) {

	account := Account{}

	id := isValidEmail(email)
	if !id.IsValid {
		return &account, errors.New("The id is not valid")
	}

	elements := []string{
		email,
	}

	err := wisply.Database.Raw("SELECT id, name, password, email, administrator FROM account WHERE email = ? ", elements).QueryRow(&account)

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

	elements := []string{
		ID,
	}
	err := wisply.Database.Raw("SELECT id, name, password, email, administrator FROM account WHERE id= ?", elements).QueryRow(&account)

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
