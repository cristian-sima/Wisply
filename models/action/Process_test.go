package action

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestProcess(t *testing.T) {

	Convey("creating simple process object", t, func() {
		process := &Process{
			Action: NewAction(false, "Simple Process"),
		}
		So(process.Action.Content, ShouldEqual, "Simple Process")
	})
}
