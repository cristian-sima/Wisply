package auth

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/nu7hatch/gouuid"
	"golang.org/x/crypto/bcrypt"
	"time"
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

func (this *AuthModel) ValidateModifyUser(rawData map[string]interface{}) (map[string][]string, error) {

	validationResult := ValidateModify(rawData)

	if !validationResult.IsValid {
		return validationResult.Errors, errors.New("Validation invalid")
	}
	return nil, nil
}

func (model *AuthModel) DeleteUserById(id string) error {

	orm := orm.NewOrm()

	elememts := []string{id}

	_, err := orm.Raw("DELETE from `user` WHERE id=?", elememts).Exec()

	return err
}

func (model *AuthModel) GetUserById(rawIndex string) (*User, error) {

	var isValid bool

	orm := orm.NewOrm()
	user := new(User)

	isValid = ValidateIndex(rawIndex)

	if !isValid {
		return user, errors.New("Validation invalid")
	}
	error := orm.Raw("SELECT id, username, password, email, administrator FROM user WHERE id = ?", rawIndex).QueryRow(&user)
	return user, error
}

func (model *AuthModel) UpdateUserType(userId string, isAdministrator string) error {

	orm := orm.NewOrm()

	fmt.Println(isAdministrator)
	stringElements := []string{isAdministrator,
		userId}

	_, err := orm.Raw("UPDATE `user` SET administrator=? WHERE id=?", stringElements).Exec()

	return err
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
		username, unsafePassword, hashedPassword, email, isAdministrator string
	)
	orm := orm.NewOrm()

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

	_, err := orm.Raw("INSERT INTO `user` (`username`, `password`, `email`, `administrator`) VALUES (?, ?, ?, ?)", elements).Exec()

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

func (this *AuthModel) TryLoginUser(rawData map[string]interface{}) (*User, error) {

	var (
		passwordIsValid         bool = false
		plainPassword, username string
	)

	orm := orm.NewOrm()
	user := new(User)

	username = rawData["username"].(string)
	plainPassword = rawData["password"].(string)

	elements := []string{username}

	error := orm.Raw("SELECT id, username, password, email, administrator FROM user WHERE username = ?", elements).QueryRow(&user)

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

func (model *AuthModel) GenerateUserToken(userId string) {

	var timeStamp string

	db := orm.NewOrm()
	token, _ := uuid.NewV4()

	t := time.Now()
	timeStamp = t.Format(time.RFC850)

	elementsDelete := []string{
		userId,
	}

	elementsInsert := []string{
		userId,
		token.String(),
		timeStamp,
	}

	db.Raw("DELETE from user_login WHERE userId= ? ", elementsDelete).Exec()

	db.Raw("INSERT INTO `user_login` (`userId`, `token`, `timestamp`) VALUES (?, ?, ?)", elementsInsert).Exec()

}

func (model *AuthModel) GetAll() []User {

	orm := orm.NewOrm()

	var list []User

	orm.Raw("SELECT id, username, password, email, administrator FROM user").QueryRows(&list)

	return list
}

func Count() int {

	orm := orm.NewOrm()

	var number int

	orm.Raw("SELECT count(*) FROM user").QueryRow(&number)

	return number
}
