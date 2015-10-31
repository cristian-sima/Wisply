package action

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	now = getCurrentTimestamp()
)

func TestAction(t *testing.T) {

	action := NewAction(false, "Something")
	action.End = now + 20

	Convey("Get right start date", t, func() {
		expectedStart := time.Unix(now, 0).Format(DateFormat)
		So(action.GetStartDate(), ShouldEqual, expectedStart)
	})
	Convey("Get right end date", t, func() {
		expectedEnd := time.Unix(now+20, 0).Format(DateFormat)
		So(action.GetEndDate(), ShouldEqual, expectedEnd)
	})
	Convey("Get end date for a process which is not finished", t, func() {
		action.End = 0
		So(action.GetEndDate(), ShouldEqual, "Not yet")
	})
	Convey("right result", t, func() {
		action.ChangeResult("danger")
		So(action.GetResult(), ShouldEqual, "danger")
	})
	Convey("Complete example", t, func() {
		action2 := &Action{
			Start: now,
			End:   0,
		}
		So(action.GetEndDate(), ShouldEqual, "Not yet")

		action2.ChangeResult("warning")
		action2.Finish()
		So(action.IsRunning, ShouldBeFalse)
	})
}

func TestActionGetDuration(t *testing.T) {

	action := NewAction(true, "Something")
	action.End = now

	Convey("Get duration for seconds", t, func() {
		So(action.getDuration(action.Start, action.End), ShouldEqual, "0")
	})

	Convey("Get duration for seconds", t, func() {
		action.End = now + 20
		So(action.getDuration(action.Start, action.End), ShouldEqual, "20s")
	})

	Convey("Get duration for minutes", t, func() {
		action.End = now + 79
		So(action.getDuration(action.Start, action.End), ShouldEqual, "1m19s")
	})

	Convey("Get duration for hours", t, func() {
		action.End = now + 3600
		So(action.getDuration(action.Start, action.End), ShouldEqual, "1h0m0s")
	})

	Convey("Get complex duration", t, func() {
		action.End = now + 3661
		So(action.getDuration(action.Start, action.End), ShouldEqual, "1h1m1s")
	})
}

func TestActionNiceDuration(t *testing.T) {

	action := NewAction(true, "Something")
	action.End = now

	Convey("Get duration for seconds", t, func() {
		duration := action.getDuration(action.Start, action.End)
		So(action.getNiceDuration(duration), ShouldEqual, "0")
	})

	Convey("Get duration for seconds", t, func() {
		action.End = now + 20
		duration := action.getDuration(action.Start, action.End)
		So(action.getNiceDuration(duration), ShouldEqual, "20 sec")
	})

	Convey("Get duration for minutes", t, func() {
		action.End = now + 79
		duration := action.getDuration(action.Start, action.End)
		So(action.getNiceDuration(duration), ShouldEqual, "1 min, 19 sec")
	})

	Convey("Get duration for hours", t, func() {
		action.End = now + 3600
		duration := action.getDuration(action.Start, action.End)
		So(action.getNiceDuration(duration), ShouldEqual, "1h, 0 min, 0 sec")
	})

	Convey("Get complex duration", t, func() {
		action.End = now + 3661
		duration := action.getDuration(action.Start, action.End)
		So(action.getNiceDuration(duration), ShouldEqual, "1h, 1 min, 1 sec")
	})
}

func TestActionDuration(t *testing.T) {

	action := NewAction(false, "Something")
	action.End = now

	Convey("Get dots if the action is not finished", t, func() {
		action.End = 0
		So(action.GetDuration(), ShouldEqual, "...")
	})

	Convey("Get duration for seconds", t, func() {
		action.End = action.Start
		So(action.GetDuration(), ShouldEqual, "0")
	})

	Convey("Get duration for seconds", t, func() {
		action.End = now + 20
		So(action.GetDuration(), ShouldEqual, "20 sec")
	})

	Convey("Get duration for minutes", t, func() {
		action.End = now + 79
		So(action.GetDuration(), ShouldEqual, "1 min, 19 sec")
	})

	Convey("Get duration for hours", t, func() {
		action.End = now + 3600
		So(action.GetDuration(), ShouldEqual, "1h, 0 min, 0 sec")
	})

	Convey("Get complex duration", t, func() {
		action.End = now + 3661
		So(action.GetDuration(), ShouldEqual, "1h, 1 min, 1 sec")
	})
}

func TestActionChange(t *testing.T) {

	action := NewAction(true, "Something")
	action.End = now

	Convey("reject invalid result type", t, func() {
		badResult := "invalid"
		So(func() {
			action.ChangeResult(badResult)
		}, ShouldPanic)
	})
	Convey("accept normal result", t, func() {
		So(func() {
			action.ChangeResult("normal")
		}, ShouldNotPanic)
	})
	Convey("accept error result", t, func() {
		So(func() {
			action.ChangeResult("danger")
		}, ShouldNotPanic)
	})
	Convey("accept warning result", t, func() {
		So(func() {
			action.ChangeResult("warning")
		}, ShouldNotPanic)
	})
	Convey("accept success result", t, func() {
		So(func() {
			action.ChangeResult("success")
		}, ShouldNotPanic)
	})
}
