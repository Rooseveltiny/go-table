package table

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDefaultReportTable(t *testing.T) {
	Convey("test table init", t, func() {
		table := NewDefaultReportTable()
		So(table.tableColumns, ShouldNotBeNil)
		So(table.tableRows, ShouldNotBeNil)
		Convey("test add some columns", func() {
			table.AddColumn("Sum Column")
			table.AddColumn("Client Name")
			table.AddColumn("ProbProbTest-123")
			So(table.tableColumns.tableColumns[2].Title, ShouldResemble, "ProbProbTest-123")
			table.DeleteColumn(table.tableColumns.tableColumns[0])
			table.DeleteColumn(table.tableColumns.tableColumns[0])
			table.DeleteColumn(table.tableColumns.tableColumns[0])
			So(len(table.tableColumns.tableColumns), ShouldBeZeroValue)
		})
		Convey("test add some columns and rows and delete it", func() {
			table.AddColumn("Sum Column")
			table.AddColumn("Client Name")
			table.AddColumn("ProbProbTest-123")
			newRow := table.AddRow()
			So(newRow.RowValues[2].tableColumn.Title, ShouldResemble, "ProbProbTest-123")
			table.DeleteColumn(table.tableColumns.tableColumns[0])
			So(len(table.tableColumns.tableColumns), ShouldEqual, 2)
			So(newRow.RowValues[0].tableColumn.Title, ShouldResemble, "Client Name")
			So(len(newRow.RowValues), ShouldEqual, 2)
			row := table.AddRow()

			row.SetValue(table.tableColumns.tableColumns[0], 12)
			row.SetValue(table.tableColumns.tableColumns[1], "memo")
			So(row.RowValues[0].value, ShouldEqual, 12)
			So(row.RowValues[1].value, ShouldEqual, "memo")

			So(table.QuantityRows(), ShouldEqual, 2)

			table.AddRow()
			table.AddRow()
			table.AddRow()
			table.AddRow()
			So(table.QuantityRows(), ShouldEqual, 6)

			table.DeleteRowByIndex(0)
			So(table.QuantityRows(), ShouldEqual, 5)

			table.DeleteRow(table.tableRows.Rows[0])
			So(table.QuantityRows(), ShouldEqual, 4)

			table.ClearTable()
			So(table.QuantityRows(), ShouldEqual, 0)

		})
	})
}
