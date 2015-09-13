package auth

import (
	"errors"
	. "github.com/cristian-sima/Wisply/models/adapter"
	. "github.com/cristian-sima/Wisply/models/wisply"
	"golang.org/x/crypto/bcrypt"
)

type Register struct {
}

func (register *Register) Try(userDetails map[string]interface{}) (WisplyError, error) {
	var problem = WisplyError{}

	registerDetails := IsValidRegister(userDetails)
	if !registerDetails.IsValid {
		problem.Data = registerDetails.Errors
		return problem, errors.New("Error")
	}

	email := userDetails["email"].(string)
	emailUsed := register.checkEmailExists(email)
	if emailUsed {
		problem.Message = "Hmmm, the email " + email + " is already used."
		return problem, errors.New("Error")
	}

	register.CreateNewAccount(userDetails)

	return problem, nil
}

func (register *Register) CreateNewAccount(details map[string]interface{}) error {
	var name, unsafePassword, hashedPassword, email, isAdministrator string

	isAdministrator = "false"
	name = details["name"].(string)
	unsafePassword = details["password"].(string)
	email = details["email"].(string)
	hashedPassword = register.getHashedPassword(unsafePassword)

	elements := []string{
		name,
		hashedPassword,
		email,
		isAdministrator,
	}
	_, err := Database.Raw("INSERT INTO `account` (`name`, `password`, `email`, `administrator`) VALUES (?, ?, ?, ?)", elements).Exec()
	return err
}

func (register *Register) getHashedPassword(plainPassword string) string {
	passwordArray := []byte(plainPassword)
	// Hashing the password with the cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordArray, 10)
	if err != nil {
		panic(err)
	}
	// concat the slice to string
	return string(hashedPassword[:])
}

func (register *Register) checkEmailExists(email string) bool {

	sql := "SELECT id FROM account WHERE email = ?"
	elements := []string{
		email,
	}
	return IsEmptyQuery(sql, elements)
}