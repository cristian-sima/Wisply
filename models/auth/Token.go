package auth

import (
	"strconv"

	database "github.com/cristian-sima/Wisply/models/database"
	"github.com/nu7hatch/gouuid"
)

// Token The token is a combination of a hashed unique random number and a Timestamp
// It is used to remember an account for a long period
// It is stored in the cookie "connection" along with the ID of the account
type Token struct {
	Value     string
	Timestamp int
	Account   *Account
}

// Encrypt It encrypts the value of the token
func (token *Token) Encrypt() string {
	return getSHA1Digest(token.Value)
}

// Insert inserts the token in the database
func (token *Token) Insert() {
	sql := "INSERT INTO `account_token` (`id`, `account`, `value`, `timestamp`) VALUES (?, ?, ?, ?)"
	query, _ := database.Connection.Prepare(sql)
	query.Exec("NULL", strconv.Itoa(token.Account.ID), token.Encrypt(), token.Timestamp)
}

// IsValid checks if the token has not expired
func (token *Token) IsValid() bool {
	var isValid bool
	now, _ := strconv.Atoi(getCurrentTimestamp())
	duration := Settings["duration"].(int)
	if token.Timestamp+duration >= now {
		isValid = true
	}
	return isValid
}

// It creates a new token using a LoginCookie
func newTokenFromCookie(cookie *LoginCookie) (*Token, error) {
	token := &Token{}
	sql := "SELECT value, timestamp FROM account_token WHERE account=? AND value=?"
	query, err := database.Connection.Prepare(sql)
	query.QueryRow(cookie.AccountID, cookie.Token).Scan(&token.Value, &token.Timestamp)
	return token, err
}

func generateToken(account *Account) *Token {
	randomNumber, _ := uuid.NewV4()
	plainValue := randomNumber.String()
	timestamp := getCurrentTimestamp()
	intTimestamp, _ := strconv.Atoi(timestamp)
	token := Token{
		Value:     plainValue,
		Timestamp: intTimestamp,
		Account:   account,
	}
	return &token
}
