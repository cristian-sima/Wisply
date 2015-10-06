package auth

import (
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestToken(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Token value is encrypted", t, func() {
		value := "cat"
		token := Token{
			Value: value,
		}
		encrypted := token.Encrypt()
		digest := getSHA1Digest(value)
		So(encrypted, ShouldEqual, digest)
	})
}

func TestTokenValidation(t *testing.T) {

	oneDay := 60 * 60 * 24
	oneWeek := oneDay * 7
	maxAllowedPeriod := Settings["duration"].(int)

	Convey("Valid token from yesterday", t, func() {
		now, _ := strconv.Atoi(getCurrentTimestamp())
		timestamp := now - oneDay
		token := Token{
			Timestamp: timestamp,
		}
		So(token.IsValid(), ShouldBeTrue)
	})

	Convey("Valid token from last week", t, func() {
		now, _ := strconv.Atoi(getCurrentTimestamp())
		timestamp := now - oneWeek
		token := Token{
			Timestamp: timestamp,
		}
		So(token.IsValid(), ShouldBeTrue)
	})

	Convey("Valid token from last allowed period", t, func() {
		now, _ := strconv.Atoi(getCurrentTimestamp())
		timestamp := now - maxAllowedPeriod
		token := Token{
			Timestamp: timestamp,
		}
		So(token.IsValid(), ShouldBeTrue)
	})

	Convey("Expired token from the max allowed period + one second", t, func() {
		now, _ := strconv.Atoi(getCurrentTimestamp())
		timestamp := now - maxAllowedPeriod - 1
		token := Token{
			Timestamp: timestamp,
		}
		So(token.IsValid(), ShouldBeFalse)
	})
}
