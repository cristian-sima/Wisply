package auth

import (
	"errors"
	"github.com/astaxie/beego/orm"
	. "github.com/cristian-sima/Wisply/models/wisply"
	"github.com/nu7hatch/gouuid"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type AuthModel struct {
	WisplyModel
}

func (this *AuthModel) ValidateRegisterDetails(rawData map[string]interface{}) (map[string][]string, error) {
	validationResult := ValidateNewUserDetails(rawData)
	if !validationResult.IsValid {
		return validationResult.Errors, errors.New("Validation invalid")
	}
	return nil, nil
}

func (this *AuthModel) ValidateModifyUser(rawData map[string]interface{}) (map[string][]string, error) {
	validationResult := ValidateModify(rawData)
	if !validationResult.IsValid {
		return validationResult.Errors, errors.New("Validation invalid")
	}
	return nil, nil
}

func (model *AuthModel) DeleteUserById(id string) error {
	elememts := []string{id}
	_, err := Database.Raw("DELETE from `user` WHERE id=?", elememts).Exec()
	return err
}

func (model *AuthModel) GetUserById(rawIndex string) (*User, error) {
	var isValid bool
	user := new(User)
	isValid = ValidateIndex(rawIndex)
	if !isValid {
		return user, errors.New("Validation invalid")
	}
	error := Database.Raw("SELECT id, username, password, email, administrator FROM user WHERE id = ?", rawIndex).QueryRow(&user)
	return user, error
}

func (model *AuthModel) UpdateUserType(userId string, isAdministrator string) error {
	stringElements := []string{isAdministrator,
		userId}
	_, err := Database.Raw("UPDATE `user` SET administrator=? WHERE id=?", stringElements).Exec()
	return err
}

func (model *AuthModel) CheckUsernameExists(username string) bool {
	var maps []orm.Params
	num, err := Database.Raw("SELECT id, username FROM user WHERE username = ?", username).Values(&maps)
	if err == nil && num > 0 {
		return true
	}
	return false
}

func (model *AuthModel) CreateNewUser(rawData map[string]interface{}) error {
	var username, unsafePassword, hashedPassword, email, isAdministrator string

	isAdministrator = "false"
	username = rawData["username"].(string)
	unsafePassword = rawData["password"].(string)
	email = rawData["email"].(string)
	hashedPassword = getHashedPassword(unsafePassword)

	elements := []string{
		username,
		hashedPassword,
		email,
		isAdministrator,
	}
	_, err := Database.Raw("INSERT INTO `user` (`username`, `password`, `email`, `administrator`) VALUES (?, ?, ?, ?)", elements).Exec()
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

func (model *AuthModel) TryLoginUser(rawData map[string]interface{}) (*User, error) {
	var (
		passwordIsValid         bool = false
		plainPassword, username string
	)
	user := new(User)
	username = rawData["username"].(string)
	plainPassword = rawData["password"].(string)
	elements := []string{username}
	error := Database.Raw("SELECT id, username, password, email, administrator FROM user WHERE username = ?", elements).QueryRow(&user)
	if error != nil {
		return user, errors.New("Problem")
	}
	passwordIsValid = checkPasswordIsCorrect(user.Password, plainPassword)
	if !passwordIsValid {
		return user, errors.New("Problem")
	}
	// connect the user
	return user, nil
}

func checkPasswordIsCorrect(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err == nil {
		return true
	}
	return false
}

func (model *AuthModel) UpdateUserLoginToken(userId string) {
	var timeStamp string
	token, _ := uuid.NewV4()
	timeStamp = model.GetCurrentTimeStamp()
	elementsDelete := []string{
		userId,
	}
	elementsInsert := []string{
		"NULL",
		userId,
		token.String(),
		timeStamp,
	}
	Database.Raw("DELETE from user_login WHERE user= ? ", elementsDelete).Exec()
	_, err := Database.Raw("INSERT INTO `user_login` (`id`, `user`, `token`, `timestamp`) VALUES (?, ?, ?, ?)", elementsInsert).Exec()
	if err != nil {
		panic(err)
	}
}

func (model *AuthModel) GetCurrentTimeStamp() string {
	var timestamp string
	unixTime := time.Now().Unix()
	timestamp = strconv.FormatInt(unixTime, 10)
	return timestamp
}

func (model *AuthModel) GetAll() []User {
	var list []User
	Database.Raw("SELECT id, username, password, email, administrator FROM user").QueryRows(&list)
	return list
}

func Count() int {
	var number int = 0
	Database.Raw("SELECT count(*) FROM user").QueryRow(&number)
	return number
}
