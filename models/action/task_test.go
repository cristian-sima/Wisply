package action

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTask(t *testing.T) {
	task := &Task{
		Action: NewAction(false, "Simple Task"),
	}
	Convey("creating simple task object", t, func() {
		So(task.Action.Content, ShouldEqual, "Simple Task")
	})
}
