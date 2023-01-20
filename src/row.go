package table

/*


 */

type TableRows struct {
	Rows []*Row
}

func (trs *TableRows) AddRow(columnsShouldHave []*tableColumn) *Row {
	newRow := newRow()
	for _, c := range columnsShouldHave {
		newRow.AddColumn(c)
	}
	trs.Rows = append(trs.Rows, newRow)
	return newRow
}

func (trs *TableRows) SetValueForRow(row Row, tableColumn tableColumn, value interface{}) {
	row.SetValue(&tableColumn, value)
}

func (trs *TableRows) AddColumnToRows(column *tableColumn) {
	for _, v := range trs.Rows {
		v.AddColumn(column)
	}
}

func (trs *TableRows) RemoveColumnOfRows(column *tableColumn) {
	for _, v := range trs.Rows {
		v.DeleteColumn(column)
	}
}

func (trs *TableRows) RemoveRow(row *Row) {
	for i, v := range trs.Rows {
		if v == row {
			trs.Rows = append(trs.Rows[:i], trs.Rows[i+1:]...)
		}
	}
}

func (trs *TableRows) RemoveRowByIndex(indx int) {
	for i := range trs.Rows {
		if i == indx {
			trs.Rows = append(trs.Rows[:i], trs.Rows[i+1:]...)
		}
	}
}

func (trs *TableRows) RemoveLastRow() {
	trs.removeLastRow()
}

func (trs *TableRows) RemoveAllRows() {
	for len(trs.Rows) != 0 {
		trs.removeLastRow()
	}
}

func (trs *TableRows) QuantityRows() int { return len(trs.Rows) }

func (trs *TableRows) removeLastRow() {
	if len(trs.Rows) != 0 {
		i := len(trs.Rows) - 1
		trs.Rows = append(trs.Rows[:i], trs.Rows[i+1:]...)
	}
}

func newTableRows() *TableRows {
	return &TableRows{}
}

/*


 */

type Row struct {
	RowValues []*RowValue
}

func (rvs *Row) DeleteColumn(column *tableColumn) {
	for i, v := range rvs.RowValues {
		if v.tableColumn == column {
			rvs.RowValues = append(rvs.RowValues[:i], rvs.RowValues[i+1:]...)
		}
	}
}

func (rvs *Row) AddColumn(column *tableColumn) {
	for _, v := range rvs.RowValues {
		if v.tableColumn == column {
			return
		}
	}
	rvs.RowValues = append(rvs.RowValues, newRowValue(column, nil))
}

func (rvs *Row) SetValue(column *tableColumn, value interface{}) {
	for _, v := range rvs.RowValues {
		if v.tableColumn == column {
			v.value = value // Could be done via setter in future
		}
	}
}

func newRow() *Row {
	return &Row{}
}

/*


 */

type RowValue struct {
	tableColumn *tableColumn
	value       interface{}
}

func newRowValue(tableColumn *tableColumn, value interface{}) *RowValue {
	return &RowValue{tableColumn: tableColumn, value: value}
}
