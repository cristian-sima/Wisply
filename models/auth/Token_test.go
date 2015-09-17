package auth

import (
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
