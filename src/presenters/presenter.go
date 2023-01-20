package presenters

import table "github.com/Rooseveltiny/go-table.git/src"

type Presenter[T string] interface {
	Present(table table.DefaultReportTable) T
}

type StringPresenter = Presenter[string]
