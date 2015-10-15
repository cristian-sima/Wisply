package action

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOperation(t *testing.T) {

	Convey("creating simple operation", t, func() {
		operation := &Operation{
			Action: NewAction(false, "Something"),
		}
		So(operation.Action.Content, ShouldEqual, "Something")
	})
}
