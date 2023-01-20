package presenters

import table "github.com/Rooseveltiny/go-table.git/src"

type TextPresenter struct{}

func (tp *TextPresenter) Present(table table.DefaultReportTable) string {
	return "text"
}
