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

// Example:
// process = newProcess...
// operation := process.CreateOperation("Testing")
// process.ChangeCurrentOperation(operation)
// task1 := operation.CreateTask("Request 1")
// time.Sleep(time.Second * 1)
// task1.Finish("It was ok")
// task2 := operation.CreateTask("Request 2")
// time.Sleep(time.Second * 1)
// task3 := operation.CreateTask("Request 3")
// time.Sleep(time.Second * 1)
// task4 := operation.CreateTask("Request 4")
// task2.CustomFinish("danger", "Timeout expired")
// time.Sleep(time.Second * 1)
// task3.CustomFinish("warning", "Ignored")
// time.Sleep(time.Second * 20)
// task4.CustomFinish("success", "Success")
// operation.Finish()
// process.Finish()
