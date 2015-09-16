package auth

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
