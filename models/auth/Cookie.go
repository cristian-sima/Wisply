package auth

import (
	"strconv"
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
