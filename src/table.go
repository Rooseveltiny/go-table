package table

/*
	DefaultReportTable is a light useful data structure to use it while making some
	reports or making some data representations
*/

type DefaultReportTable struct {
	tableRows    *TableRows
	tableColumns *TableColumns
}

func (tb *DefaultReportTable) AddColumn(title string) {
	newCol := newTableColumn(title)
	tb.tableColumns.AddNewColumn(newCol)
	tb.tableRows.AddColumnToRows(newCol)
}

func (tb *DefaultReportTable) DeleteColumn(column *tableColumn) {
	tb.tableColumns.RemoveColumn(column)
	tb.tableRows.RemoveColumnOfRows(column)
}

func (tb *DefaultReportTable) AddRow() *Row {
	return tb.tableRows.AddRow(tb.tableColumns.tableColumns)
}

func (tb *DefaultReportTable) QuantityRows() int { return tb.tableRows.QuantityRows() }

func (tb *DefaultReportTable) DeleteRowByIndex(i int) { tb.tableRows.RemoveRowByIndex(i) }

func (tb *DefaultReportTable) DeleteRow(row *Row) { tb.tableRows.RemoveRow(row) }

func (tb *DefaultReportTable) ClearTable() { tb.tableRows.RemoveAllRows() }

func NewDefaultReportTable() *DefaultReportTable {
	return &DefaultReportTable{tableRows: newTableRows(), tableColumns: newTableColumns()}
}
