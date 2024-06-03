package main

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql-samples/util"
	"github.com/rrgmc/litsql/dialect/psql"
	"github.com/rrgmc/litsql/dialect/psql/sm"
	"github.com/rrgmc/litsql/dialect/psql/tag"
	"github.com/rrgmc/litsql/sq"
	"github.com/rrgmc/litsql/sq/clause"
)

func customClause() error {
	query := psql.Select(
		sm.Columns("id", "name"),
		sm.From("users"),
		sm.WhereC("id = ?", "John"),
		MyClause(),
	)

	return util.PrintQuery(query, nil)
}

func MyClause() psql.SelectMod {
	return sq.QueryModFunc[tag.SelectTag](func(a litsql.QueryBuilder) {
		a.AddQueryClause(&myClause{})
	})
}

type myClause struct {
}

func (m *myClause) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) (args []any, err error) {
	w.AddSeparator(true)
	w.Write("myclause")
	return nil, nil
}

func (m *myClause) ClauseID() string {
	return "a9d51c89-c808-4200-8f23-b79d51fc337b"
}

func (m *myClause) ClauseOrder() int {
	return clause.OrderWhere + 10
}
