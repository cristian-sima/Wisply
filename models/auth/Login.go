package auth

import (
	"errors"
	. "github.com/cristian-sima/Wisply/models/adapter"
	"golang.org/x/crypto/bcrypt"
)

type Login struct {
}

func (login *Login) Try(loginDetails map[string]interface{}) (WisplyError, error) {

	var genericMessage string = "There was a problem while login. We think the email or the password were not valid."

	problem := WisplyError{}

	goodLoginDetails := IsValidLogin(loginDetails)
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

	validPassword := login.checkPasswordIsCorrect(account.Password, loginDetails["password"].(string))

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
