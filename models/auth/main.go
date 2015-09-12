package auth

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"github.com/astaxie/beego/orm"
	. "github.com/cristian-sima/Wisply/models/wisply"
	"github.com/nu7hatch/gouuid"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
	"time"
)

type AuthModel struct {
	WisplyModel
}

func (this *AuthModel) ValidateRegisterDetails(rawData map[string]interface{}) (map[string][]string, error) {
	validationResult := ValidateNewAccountDetails(rawData)
	if !validationResult.IsValid {
		return validationResult.Errors, errors.New("Validation invalid")
	}
	return nil, nil
}

func (this *AuthModel) ValidateModifyAccount(rawData map[string]interface{}) (map[string][]string, error) {
	validationResult := ValidateModify(rawData)
	if !validationResult.IsValid {
		return validationResult.Errors, errors.New("Validation invalid")
	}
	return nil, nil
}

func (model *AuthModel) DeleteAccountById(id string) error {
	elememts := []string{id}
	_, err := Database.Raw("DELETE from `account` WHERE id=?", elememts).Exec()
	return err
}

func (model *AuthModel) GetAccountById(rawIndex string) (*Account, error) {
	var isValid bool
	account := new(Account)
	isValid = ValidateIndex(rawIndex)
	if !isValid {
		return account, errors.New("Validation invalid")
	}
	error := Database.Raw("SELECT id, name, password, email, administrator FROM account WHERE id = ?", rawIndex).QueryRow(&account)
	return account, error
}

func (model *AuthModel) UpdateAccountType(accountId string, isAdministrator string) error {
	stringElements := []string{isAdministrator,
		accountId}
	_, err := Database.Raw("UPDATE `account` SET administrator=? WHERE id=?", stringElements).Exec()
	return err
}

func (model *AuthModel) CheckEmailExists(name string) bool {
	var maps []orm.Params
	num, err := Database.Raw("SELECT id, name FROM account WHERE email = ?", name).Values(&maps)
	if err == nil && num > 0 {
		return true
	}
	return false
}

func (model *AuthModel) CreateNewAccount(rawData map[string]interface{}) error {
	var name, unsafePassword, hashedPassword, email, isAdministrator string

	isAdministrator = "false"
	name = rawData["name"].(string)
	unsafePassword = rawData["password"].(string)
	email = rawData["email"].(string)
	hashedPassword = getHashedPassword(unsafePassword)

	elements := []string{
		name,
		hashedPassword,
		email,
		isAdministrator,
	}
	_, err := Database.Raw("INSERT INTO `account` (`name`, `password`, `email`, `administrator`) VALUES (?, ?, ?, ?)", elements).Exec()
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

func (model *AuthModel) TryLoginAccount(rawData map[string]interface{}) (*Account, error) {
	var (
		passwordIsValid      bool = false
		plainPassword, email string
	)
	account := new(Account)
	email = rawData["email"].(string)
	plainPassword = rawData["password"].(string)
	elements := []string{email}
	error := Database.Raw("SELECT id, name, password, email, administrator FROM account WHERE email = ?", elements).QueryRow(&account)
	if error != nil {
		return account, errors.New("Problem")
	}
	passwordIsValid = checkPasswordIsCorrect(account.Password, plainPassword)
	if !passwordIsValid {
		return account, errors.New("Problem")
	}
	// connect the account
	return account, nil
}

func checkPasswordIsCorrect(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err == nil {
		return true
	}
	return false
}

func (model *AuthModel) GetLoginCookie(accountId string) string {
	var cookie, plainToken, hashedToken string

	plainToken = model.generateAccoutLoginToken(accountId)
	hashedToken = model.getSHA1_digest(plainToken)

	cookie = model.encryptCookie(accountId, hashedToken)
	return cookie
}

func (model *AuthModel) getSHA1_digest(plainToken string) string {
	array := []byte(plainToken)
	hasher := sha1.New()
	hasher.Write(array)
	token := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return token
}

func (model *AuthModel) encryptCookie(accountId, hashedToken string) string {
	var cookieValue string
	cookieValue = accountId + "::" + hashedToken
	return cookieValue
}

func (model *AuthModel) decrypyCookie(cookieValue string) (string, string) {
	values := strings.Split(cookieValue, "::")
	return values[0], values[1]
}

func (model *AuthModel) generateAccoutLoginToken(accountId string) string {
	var timestamp, token string
	tokenObject, _ := uuid.NewV4()
	token = tokenObject.String()

	timestamp = model.GetCurrentTimeStamp()

	elementsInsert := []string{
		"NULL",
		accountId,
		timestamp,
	}
	Database.Raw("INSERT INTO `account_login` (`id`, `account`, `token`, `timestamp`) VALUES (?, ?, ?, ?)", elementsInsert).Exec()

	return token
}

func (model *AuthModel) IsConnectionValid (accountId string) bool {
	return model.checkTokenIsValid(accountId)
}

func (model *AuthModel) checkTokenIsValid(accountId string) bool {

	var (
		now, duration, login int
		isValid bool
	)

	now,_ = strconv.Atoi(model.GetCurrentTimeStamp())
	duration = beego.AppConfig.String("connectionDuration")

	login = model.getAccountLoginTokenTimestamp();

	isValid = (now <= (login + duration))
	return isValid
}

func (model *AuthModel) getAccountLoginTokenTimestamp() int {

		var token int
		type Temp struct {
			Timestamp string
	}
		temp := new(Temp)
		Database.Raw("SELECT timestamp FROM account_login").QueryRow(&temp)
		token,_  = strconv.Atoi(temp.Timestamp)
		return token
}

func (model *AuthModel) DeleteOldTokens(accountId string) {

		elementsDelete := []string{
			accountId,
		}
		Database.Raw("DELETE from account_login WHERE account= ? AND timestamp ", elementsDelete).Exec()

}

func (model *AuthModel) IsLoginValid(accountId string) bool {
	//var tokenValid bool
	//tokenValid = model.checkTimestampIsValid()
	return true
}

func (model *AuthModel) GetCurrentTimeStamp() string {
	var timestamp string
	unixTime := time.Now().Unix()
	timestamp = strconv.FormatInt(unixTime, 10)
	return timestamp
}

func (model *AuthModel) GetAll() []Account {
	var list []Account
	Database.Raw("SELECT id, name, password, email, administrator FROM account").QueryRows(&list)
	return list
}

func Count() int {
	var number int = 0
	Database.Raw("SELECT count(*) FROM account").QueryRow(&number)
	return number
}
