package auth

import (
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var separator = Settings["separatorCookie"].(string)

func TestCookie(t *testing.T) {

	tokenValue := "This is a test 1234"
	hashedValue := getSHA1Digest(tokenValue)

	account := &Account{
		ID: 4,
	}
	token := &Token{
		Account: account,
		Value:   tokenValue,
	}
	Convey("Get cookie value", t, func() {
		cookie := Cookie{
			Token:   token,
			Account: account,
		}
		So(cookie.GetValue(), ShouldEqual, strconv.Itoa(account.ID)+separator+hashedValue)
	})
}

func TestRejectLoginCookie(t *testing.T) {
	Convey("Empty login cookie", t, func() {
		plainCookie := ""
		cookieIsValid := hasGoodFormat(plainCookie)
		So(cookieIsValid, ShouldBeFalse)
	})
	Convey("Only ID provided", t, func() {
		plainCookie := "120" + separator
		cookieIsValid := hasGoodFormat(plainCookie)
		So(cookieIsValid, ShouldBeFalse)
	})
	Convey("Only token provided", t, func() {
		plainCookie := separator + "CT7htExPXZKmbeTrc_1WyuMQg3w="
		cookieIsValid := hasGoodFormat(plainCookie)
		So(cookieIsValid, ShouldBeFalse)
	})
	Convey("Wrong separator", t, func() {
		plainCookie := "123:CT7htExPXZKmbeTrc_1WyuMQg3w="
		cookieIsValid := hasGoodFormat(plainCookie)
		So(cookieIsValid, ShouldBeFalse)
	})
	Convey("Only separator", t, func() {
		plainCookie := "::"
		cookieIsValid := hasGoodFormat(plainCookie)
		So(cookieIsValid, ShouldBeFalse)
	})
	Convey("Not a numeric ID", t, func() {
		plainCookie := "-::CT7htExPXZKmbeTrc_1WyuMQg3w="
		cookieIsValid := hasGoodFormat(plainCookie)
		So(cookieIsValid, ShouldBeFalse)
	})
}

func TestAcceptsLoginCookie(t *testing.T) {
	Convey("Accept valid cookie", t, func() {
		plainCookie := "96" + separator + "TCT7htExPXZKmabeTrc_1WyuMQg3w="
		cookieIsValid := hasGoodFormat(plainCookie)
		So(cookieIsValid, ShouldBeTrue)
	})
	Convey("Get login cookie elements", t, func() {
		plainCookie := "19" + separator + "CT7htExPXZKmbeTrc_1WyuMQg3w="
		elements := getCookieElements(plainCookie)
		So(elements[0], ShouldEqual, "19")
		So(elements[1], ShouldEqual, "CT7htExPXZKmbeTrc_1WyuMQg3w=")
	})
}

func TestLoginCookieElements(t *testing.T) {
	Convey("Get the values of login cookie from plain text", t, func() {
		plainCookie := "189" + separator + "CT7htExPXZKmbeTrc_1WyuMQg3w="
		cookie, err := newLoginCookie(plainCookie)
		So(err, ShouldBeNil)
		ID := cookie.AccountID
		Token := cookie.Token
		So(ID, ShouldEqual, "189")
		So(Token, ShouldEqual, "CT7htExPXZKmbeTrc_1WyuMQg3w=")
	})
	Convey("Return error for a invalid cookie", t, func() {
		plainCookie := "18a9:CT7htExPXZKmbeTrc_1WyuMQg3w="
		_, err := newLoginCookie(plainCookie)
		So(err, ShouldNotBeNil)
	})
}
