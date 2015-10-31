package api

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTables(t *testing.T) {

	Convey("Rejects a sensitive table", t, func() {
		isRejected := IsRestrictedTable("account")
		So(isRejected, ShouldBeTrue)
	})

	Convey("Accepts a table which is not on the list", t, func() {
		isRejected := IsRestrictedTable("task")
		So(isRejected, ShouldBeFalse)
	})
}
