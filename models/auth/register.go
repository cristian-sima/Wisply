package auth

import (
	"errors"

	"github.com/cristian-sima/Wisply/models/adapter"
	"github.com/cristian-sima/Wisply/models/database"
	"golang.org/x/crypto/bcrypt"
)

// Register It manages the operations for creating a new account
type Register struct {
}

// Try It tries to create a new account
func (register *Register) Try(userDetails map[string]interface{}) (adapter.WisplyError, error) {
	var problem = adapter.WisplyError{}

	result := isValidRegister(userDetails)
	if !result.IsValid {
		problem.Data = result
		return problem, errors.New("Error")
	}

	email := userDetails["email"].(string)
	emailUsed := register.checkEmailExists(email)
	if emailUsed {
		problem.Message = "Hmmm, the email " + email + " is already used."
		return problem, errors.New("Error")
	}

	register.createNewAccount(userDetails)

	return problem, nil
}

// It creates the new account
func (register *Register) createNewAccount(details map[string]interface{}) error {
	var name, unsafePassword, hashedPassword, email, isAdministrator string

	isAdministrator = "false"
	name = details["name"].(string)
	unsafePassword = details["password"].(string)
	email = details["email"].(string)
	hashedPassword = register.getHashedPassword(unsafePassword)

	fieldList := "`name`, `password`, `email`, `administrator`"
	questionList := "?, ?, ?, ?"
	sql := "INSERT INTO `account` (" + fieldList + ") VALUES (" + questionList + ")"
	query, err := database.Connection.Prepare(sql)
	query.Exec(name, hashedPassword, email, isAdministrator)
	return err
}

func (register *Register) getHashedPassword(plainPassword string) string {
	passwordArray := []byte(plainPassword)
	// Hashing the password with the cost of 10
	hashedPassword, _ := bcrypt.GenerateFromPassword(passwordArray, 10)
	return string(hashedPassword[:])
}

func (register *Register) checkEmailExists(email string) bool {
	var (
		exists bool
		id     int
	)
	sql := "SELECT id FROM account WHERE email = ?"
	query, _ := database.Connection.Prepare(sql)
	query.QueryRow(email).Scan(&id)

	if id != 0 {
		exists = true
	}
	return exists
}
