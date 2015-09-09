package auth

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

type AuthModel struct {
}

func (this *AuthModel) ValidateRegisterDetails(rawData map[string]interface{}) (map[string][]string, error) {

	validationResult := ValidateNewUserDetails(rawData)

	if !validationResult.IsValid {
		return validationResult.Errors, errors.New("Validation invalid")
	}
	return nil, nil
}

func (this *AuthModel) CheckUsernameExists(username string) bool {

	var maps []orm.Params

	db := orm.NewOrm()

	num, err := db.Raw("SELECT id, username FROM user WHERE username = ?", username).Values(&maps)

	if err == nil && num > 0 {
		return true
	}
	return false
}

func (this *AuthModel) CreateNewUser(rawData map[string]interface{}) error {

	var (
		username, unsafePassword, hashedPassword, email string
		isAdmin                                         string = "0"
	)
	orm := orm.NewOrm()

	username = rawData["username"].(string)
	unsafePassword = rawData["password"].(string)
	email = rawData["email"].(string)

	hashedPassword = getHashedPassword(unsafePassword)

	elements := []string{
		username,
		hashedPassword,
		email,
		isAdmin,
	}

	_, err := orm.Raw("INSERT INTO `user` (`username`, `password`, `email`, `isAdmin`) VALUES (?, ?, ?, ?)", elements).Exec()

	return err
}

func getHashedPassword(plainPassword string) string {
	passwordArray := []byte(plainPassword)
	// Hashing the password with the cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordArray, 10)
	if err != nil {
		panic(err)
	}
	// concat the slice to string
	return string(hashedPassword[:])
}

func (this *AuthModel) ValidateLoginDetails(rawData map[string]interface{}) (map[string][]string, error) {

	validationResult := ValidateLoginDetails(rawData)

	if !validationResult.IsValid {
		return validationResult.Errors, errors.New("Validation invalid")
	}
	return nil, nil
}

func (this *AuthModel) TryLoginUser(rawData map[string]interface{}) error {

	var (
		passwordIsValid         bool = false
		plainPassword, username string
	)

	orm := orm.NewOrm()
	user := new(User)

	username = rawData["username"].(string)
	plainPassword = rawData["password"].(string)

	elements := []string{username}

	error := orm.Raw("SELECT id, username, password, email, isAdmin FROM user WHERE username = ?", elements).QueryRow(&user)

	if error != nil {
		return errors.New("Problem")
	}

	fmt.Println(user.Password)
	fmt.Println(plainPassword)
	passwordIsValid = checkPasswordIsCorrect(user.Password, plainPassword)

	if !passwordIsValid {
		return errors.New("Problem")
	}

	// connect the user
	return nil
}

func checkPasswordIsCorrect(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err == nil {
		return true
	}
	return false
}

type User struct {
	Id       int
	Username string
	Password string
	Email    string
	isAdmin  bool
}

func (user *User) isAdministrator() bool {
	return user.isAdmin
}

func NewUser() User {
	var user User
	orm := orm.NewOrm()

	orm.Raw("SELECT username, password, isAdmin FROM user").QueryRow(&user)

	return user
}
