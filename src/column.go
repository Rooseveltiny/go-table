package table

import "fmt"

type TableColumns struct {
	tableColumns []*tableColumn
}

func (tcs *TableColumns) AddNewColumn(column *tableColumn) {
	tcs.tableColumns = append(tcs.tableColumns, column)
}

func (tcs *TableColumns) RemoveColumn(column *tableColumn) error {
	for i, v := range tcs.tableColumns {
		if v == column {
			tcs.removeColumnByIndex(i)
			return nil
		}
	}
	return fmt.Errorf("could not find such column")
}

func (tcs *TableColumns) removeColumnByIndex(i int) {
	tcs.tableColumns = append(tcs.tableColumns[:i], tcs.tableColumns[i+1:]...)
}

func newTableColumns() *TableColumns {
	return &TableColumns{}
}

type tableColumn struct {
	Title string
}

func newTableColumn(title string) *tableColumn {
	return &tableColumn{Title: title}
}
