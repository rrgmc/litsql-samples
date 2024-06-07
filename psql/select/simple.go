package main

import (
	"github.com/rrgmc/litsql-samples/util"
	"github.com/rrgmc/litsql/dialect/psql"
	"github.com/rrgmc/litsql/dialect/psql/sm"
)

func simple() error {
	query := psql.Select(
		sm.Columns("id", "name"),
		sm.From("users"),
		sm.WhereClause("id = ?", "John"),
	)

	return util.PrintQuery(query, nil)
}
