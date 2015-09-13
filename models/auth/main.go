package auth

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	. "github.com/cristian-sima/Wisply/models/wisply"
	"strconv"
	"strings"
	"time"
)

var Settings = map[string]interface{}{
	"duration":        (60 * 60 * 24 * 7), // one week
	"path":            "/",
	"separatorCookie": "::",
	"cookieName":      "connection",
	"cookiePath":      "/",
}

type AuthModel struct {
	WisplyModel
}

func ReConnect(plainCookie string) (string, error) {
	var accountId string = ""

	elements, err := splitCookie(plainCookie)
	if err != nil {
		return accountId, errors.New("The cookie has an invalid format")
	}

	accountId = elements[0]
	hashedToken := elements[1]

	validToken := isTokenValid(accountId, hashedToken)
	if !validToken {
		return accountId, errors.New("The token is not valid")
	}

	deleteOldTokens()

	return accountId, nil
}

func isTokenValid(accountId, hashedToken string) bool {

	elements := []string{
		accountId,
		hashedToken,
	}
	token := Token{}

	if Database == nil {
		return false
	}

	sql := "SELECT value, timestamp FROM account_token WHERE account=? AND value=?"
	err := Database.Raw(sql, elements).QueryRow(&token)

	if err != nil {
		return false
	}

	now, _ := strconv.Atoi(GetCurrentTimestamp())
	duration := Settings["duration"].(int)

	fmt.Println(elements)
	fmt.Println(token)

	fmt.Println("now: ")
	fmt.Println(now)

	fmt.Println("duration: ")
	fmt.Println(duration)

	fmt.Println("value: ")
	fmt.Println(token.Timestamp)

	isValid := (now <= (token.Timestamp + duration))
	return isValid
}

func deleteOldTokens() {
	var (
		now, duration, diff int
	)
	now, _ = strconv.Atoi(GetCurrentTimestamp())
	duration = Settings["duration"].(int)
	diff = now - duration

	elementsDelete := []string{
		strconv.Itoa(diff),
	}
	Database.Raw("DELETE from account_token WHERE timestamp < ?", elementsDelete).Exec()
}

func (model *AuthModel) GetAllAccounts() []Account {
	var list []Account
	Database.Raw("SELECT id, name, password, email, administrator FROM account").QueryRows(&list)
	return list
}

// --------------------

func CountAccounts() int {
	var number int = 0
	Database.Raw("SELECT count(*) FROM account").QueryRow(&number)
	return number
}

func GetCurrentTimestamp() string {
	var timestamp string
	unixTime := time.Now().Unix()
	timestamp = strconv.FormatInt(unixTime, 10)
	return timestamp
}

func splitCookie(cookieValue string) ([]string, error) {
	var (
		toReturn  []string
		separator string = Settings["separatorCookie"].(string)
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

func GetAccountByEmail(email string) (*Account, error) {

	account := Account{}

	id := IsValidEmail(email)
	if !id.IsValid {
		return &account, errors.New("The id is not valid")
	}

	elements := []string{
		email,
	}

	err := Database.Raw("SELECT id, name, password, email, administrator FROM account WHERE email = ? ", elements).QueryRow(&account)

	if err != nil {
		fmt.Println(err)
		return &account, errors.New("No such account")
	}

	return &account, nil
}

func NewAccount(rawId string) (*Account, error) {

	id := IsValidId(rawId)
	if !id.IsValid {
		return nil, errors.New("The id is not valid")
	}

	account := new(Account)

	elements := []string{
		rawId,
	}
	err := Database.Raw("SELECT id, name, password, email, administrator FROM account WHERE id= ?", elements).QueryRow(&account)

	if err != nil {
		return account, errors.New("Error")
	}

	return account, nil
}

func getSHA1_digest(plainToken string) string {
	array := []byte(plainToken)
	hasher := sha1.New()
	hasher.Write(array)
	token := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return token
}
