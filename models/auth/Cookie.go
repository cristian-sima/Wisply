package auth

import (
	"strconv"
)

type Cookie struct {
	Token    *Token
	Account  *Account
	Path     string
	Duration int
}

func (cookie *Cookie) GetValue() string {
	return strconv.Itoa(cookie.Account.Id) + "::" + cookie.Token.Encrypt()
}
