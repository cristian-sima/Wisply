package auth

type Token struct {
	Value     string
	Timestamp int
}

func (token *Token) Encrypt() string {
	return getSHA1_digest(token.Value)
}
