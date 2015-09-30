package auth

import (
	"fmt"
	"strconv"

	wisply "github.com/cristian-sima/Wisply/models/database"
)

// Token The token is a combination of a hashed unique random number and a Timestamp
// It is used to remember an account for a long period
// It is stored in the cookie "connection" along with the ID of the account
type Token struct {
	Value     string
	Timestamp int
}

// Encrypt It encrypts the value of the token
func (token *Token) Encrypt() string {
	return getSHA1Digest(token.Value)
}

// IsValid checks if the token has not expired
func (token *Token) IsValid() bool {
	var isValid bool
	now, _ := strconv.Atoi(getCurrentTimestamp())
	duration := Settings["duration"].(int)
	fmt.Println(strconv.Itoa(now - duration - token.Timestamp))
	if token.Timestamp+duration >= now {
		isValid = true
	}
	return isValid
}

// It creates a new token using a LoginCookie
func newTokenFromCookie(cookie *LoginCookie) (*Token, error) {
	token := &Token{}
	sql := "SELECT value, timestamp FROM account_token WHERE account=? AND value=?"
	query, err := wisply.Database.Prepare(sql)
	query.QueryRow(cookie.AccountID, cookie.Token).Scan(&token.Value, &token.Timestamp)
	return token, err
}
