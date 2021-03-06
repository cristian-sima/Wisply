package auth

import (
	"errors"
	"strconv"
	"strings"

	"github.com/cristian-sima/Wisply/models/database"
)

// Cookie It represents a connection cookie
// It is used to remember the user for a long period of time
type Cookie struct {
	Token    *Token
	Account  *Account
	Path     string
	Duration int
}

// GetValue It returns the value of the Cookie
// The value is in this format: ID::EncryptedToken
func (cookie *Cookie) GetValue() string {
	return strconv.Itoa(cookie.Account.ID) + "::" + cookie.Token.Encrypt()
}

// LoginCookie represents a cookie for login
// It is used to remember the user for a long period of time
type LoginCookie struct {
	Token     string
	AccountID string
}

// IsGood creates a Token from the id and hashed token and calls its IsValid method
func (cookie *LoginCookie) IsGood() bool {
	if database.Connection == nil {
		return false
	}
	token, err := newTokenFromCookie(cookie)
	if err != nil {
		return false
	}
	isValid := token.IsValid()
	return isValid
}

func newLoginCookie(plainCookie string) (*LoginCookie, error) {
	var cookie *LoginCookie

	if !hasGoodFormat(plainCookie) {
		message := "The cookie does not have a valid format"
		return cookie, errors.New(message)
	}
	elements := getCookieElements(plainCookie)

	cookie = &LoginCookie{
		AccountID: elements[0],
		Token:     elements[1],
	}

	return cookie, nil
}

func hasGoodFormat(plainCookie string) bool {
	separator := Settings["separatorCookie"].(string)
	containsSeparator := strings.Contains(plainCookie, separator)
	if !containsSeparator {
		return false
	}
	elements := getCookieElements(plainCookie)
	if len(elements) != 2 {
		return false
	}
	ID := elements[0]
	token := elements[1]
	if len(ID) == 0 || len(token) == 0 {
		return false
	}
	if _, err := strconv.Atoi(ID); err != nil {
		return false
	}
	return true
}

func getCookieElements(plainCookie string) []string {
	var elements []string
	separator := Settings["separatorCookie"].(string)
	elements = strings.Split(plainCookie, separator)
	return elements
}
