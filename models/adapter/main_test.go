package adapter

import (
	"testing"

	"local-projects/validity"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWisplyError(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Empty Error Message", t, func() {
		error := WisplyError{}
		msg := error.GetMessage()
		So(msg, ShouldEqual, "There was a problem")
	})

	Convey("Test Wisply Error", t, func() {
		error := WisplyError{}
		error.Message = "Simple error message"
		msg := error.GetMessage()
		Convey("The value should be the simple message", func() {
			So(msg, ShouldEqual, "Simple error message")
		})
	})

	Convey("Validation error message, one field", t, func() {
		error := WisplyError{}
		data := make(map[string]interface{})
		data["email"] = ""
		rules := validity.ValidationRules{
			"email": []string{"String", "email"},
		}
		result := Validate(data, rules)
		error.Data = result.Errors
		msg := error.GetMessage()
		Convey("The error should detect problems with one field", func() {
			So(msg, ShouldEqual, "Your request was not successful. There were problems with one field:")
		})
	})

	Convey("Validation error message, multiple fields", t, func() {
		error := WisplyError{}
		data := make(map[string]interface{})
		data["email"] = "iiii"
		data["name"] = ""
		data["password"] = ""
		rules := validity.ValidationRules{
			"name":     []string{"String", "between_inclusive:3,25"},
			"email":    []string{"String", "email"},
			"password": []string{"String", "between_inclusive:6,60"},
		}
		result := Validate(data, rules)
		error.Data = result.Errors
		msg := error.GetMessage()
		Convey("The error should detect multiple problems", func() {
			So(msg, ShouldEqual, "Your request was not successful. There were problems with 3 fields:")
		})
	})
}
