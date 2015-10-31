package auth

import (
	"errors"
	"fmt"

	adapter "github.com/cristian-sima/Wisply/models/adapter"
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
	result := isValidLogin(loginDetails)
	if !result.IsValid {
		problem.Data = result
		fmt.Println(result.Errors)
		return problem, errors.New("Error")
	}

	email := loginDetails["email"].(string)

	account, err := GetAccountByEmail(email)
	if err != nil {
		problem.Message = genericMessage
		return problem, errors.New(genericMessage)
	}

	passwordString := loginDetails["password"].(string)
	validPassword := checkPasswordIsCorrect(account.Password, passwordString)

	if !validPassword {
		problem.Message = genericMessage
		return problem, errors.New(genericMessage)
	}

	return problem, nil
}
