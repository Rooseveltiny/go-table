package table

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRowStructs(t *testing.T) {
	Convey("test table rows init", t, func() {
		newTRs := newTableRows()
		So(len(newTRs.Rows), ShouldEqual, 0)
	})
	Convey("test row values init", t, func() {
		newRVs := newRow()
		So(len(newRVs.RowValues), ShouldEqual, 0)
	})
	Convey("test row value init", t, func() {
		newRV := newRowValue(newTableColumn("column name"), 12)
		So(newRV.value, ShouldEqual, 12)
		So(newRV.tableColumn.Title, ShouldEqual, "column name")
	})
	Convey("test adding new row values", t, func() {
		newRVs := newRow()
		colToAdd := newTableColumn("col test name")
		newRVs.AddColumn(colToAdd)
		newRVs.AddColumn(colToAdd)
		newRVs.AddColumn(colToAdd)
		So(len(newRVs.RowValues), ShouldEqual, 1)
		anotherCol := newTableColumn("col test name")
		newRVs.AddColumn(anotherCol)
		newRVs.AddColumn(anotherCol)
		So(len(newRVs.RowValues), ShouldEqual, 2)
		anotherCol1 := newTableColumn("col test name")
		anotherCol2 := newTableColumn("col test name")
		newRVs.AddColumn(anotherCol1)
		newRVs.AddColumn(anotherCol2)
		So(len(newRVs.RowValues), ShouldEqual, 4)
		newRVs.DeleteColumn(anotherCol1)
		So(len(newRVs.RowValues), ShouldEqual, 3)
		newRVs.DeleteColumn(anotherCol1)
		So(len(newRVs.RowValues), ShouldEqual, 3)
		newRVs.SetValue(anotherCol2, "test test test")
		So(newRVs.RowValues[2].value, ShouldResemble, "test test test")
		newRVs.SetValue(anotherCol2, 12)
		So(newRVs.RowValues[2].value, ShouldResemble, 12)
		newRVs.SetValue(colToAdd, true)
		So(newRVs.RowValues[0].value, ShouldResemble, true)
	})
	Convey("test manipulating with TableRows", t, func() {
		TRs := newTableRows()
		TRs.AddRow([]*tableColumn{})
		TRs.AddRow([]*tableColumn{})
		testRow := TRs.AddRow([]*tableColumn{})
		testRow1 := TRs.AddRow([]*tableColumn{})
		col := newTableColumn("sum")
		TRs.AddColumnToRows(col)
		col1 := newTableColumn("name")
		TRs.AddColumnToRows(col1)
		col2 := newTableColumn("client")
		TRs.AddColumnToRows(col2)
		col3 := newTableColumn("test")
		TRs.AddColumnToRows(col3)
		col4 := newTableColumn("example")
		TRs.AddColumnToRows(col4)
		So(TRs.Rows[0].RowValues[1].tableColumn.Title, ShouldResemble, "name")
		So(TRs.Rows[1].RowValues[1].tableColumn.Title, ShouldResemble, "name")
		So(TRs.Rows[2].RowValues[1].tableColumn.Title, ShouldResemble, "name")
		So(TRs.Rows[3].RowValues[1].tableColumn.Title, ShouldResemble, "name")
		TRs.RemoveColumnOfRows(col1)
		So(TRs.Rows[0].RowValues[1].tableColumn.Title, ShouldResemble, "client")
		So(TRs.Rows[1].RowValues[1].tableColumn.Title, ShouldResemble, "client")
		So(TRs.Rows[2].RowValues[1].tableColumn.Title, ShouldResemble, "client")
		So(TRs.Rows[2].RowValues[1].tableColumn.Title, ShouldResemble, "client")
		testRow.SetValue(col3, 757)
		So(TRs.Rows[2].RowValues[2].value, ShouldResemble, 757)
		So(TRs.QuantityRows(), ShouldEqual, 4)
		TRs.RemoveRow(testRow1)
		So(TRs.QuantityRows(), ShouldEqual, 3)
		TRs.RemoveLastRow()
		So(TRs.QuantityRows(), ShouldEqual, 2)
		TRs.RemoveAllRows()
		So(TRs.QuantityRows(), ShouldEqual, 0)
	})
}