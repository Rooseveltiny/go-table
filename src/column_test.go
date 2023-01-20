package table

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestColumn(t *testing.T) {
	Convey("test table column init", t, func() {
		nameOfCol := "test name for column"
		newCol := newTableColumn(nameOfCol)
		So(newCol.Title, ShouldEqual, nameOfCol)
	})
	Convey("test table columns struct init", t, func() {
		newColumns := newTableColumns()
		So(len(newColumns.tableColumns), ShouldEqual, 0)
		Convey("test adding new column", func() {
			newColumns.AddNewColumn(newTableColumn("Sum of revenue"))
			newColumns.AddNewColumn(newTableColumn("Manager name"))
			newColumns.AddNewColumn(newTableColumn("Sum of revenue"))
			So(newColumns.tableColumns[1].Title, ShouldEqual, "Manager name")
			Convey("test removing column", func() {
				columnToRemove := newColumns.tableColumns[1]
				newColumns.RemoveColumn(columnToRemove)
				So(len(newColumns.tableColumns), ShouldEqual, 2)
				err := newColumns.RemoveColumn(columnToRemove)
				So(err, ShouldNotBeNil)
			})
		})
	})
}
