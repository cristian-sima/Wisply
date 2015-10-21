package auth

import (
	"errors"

	adapter "github.com/cristian-sima/Wisply/models/adapter"
	"golang.org/x/crypto/bcrypt"
)

// Login It manages the login operations
type Login struct {
}

// Try It tries to log in the username
func (login *Login) Try(loginDetails map[string]interface{}) (adapter.WisplyError, error) {

	deleteOldTokens()

	thinkMessage := "We think the email or the password were not valid."
	genericMessage := "There was a problem while login. " + thinkMessage

	problem := adapter.WisplyError{}

	goodLoginDetails := isValidLogin(loginDetails)
	if !goodLoginDetails.IsValid {
		problem.Data = goodLoginDetails.Errors
		return problem, errors.New("Error")
	}

	email := loginDetails["email"].(string)

	account, err := GetAccountByEmail(email)
	if err != nil {
		problem.Message = genericMessage
		return problem, errors.New(genericMessage)
	}

	passwordString := loginDetails["password"].(string)
	validPassword := login.checkPasswordIsCorrect(account.Password, passwordString)

	if !validPassword {
		problem.Message = genericMessage
		return problem, errors.New(genericMessage)
	}

	return problem, nil
}

func (login *Login) checkPasswordIsCorrect(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err == nil {
		return true
	}
	return false
}
