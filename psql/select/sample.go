package main

import (
	"github.com/rrgmc/litsql-samples/util"
	"github.com/rrgmc/litsql/dialect/psql"
	"github.com/rrgmc/litsql/dialect/psql/sm"
)

func sample() error {
	type userFilter struct {
		Name string
	}
	filter := userFilter{
		Name: "john",
	}
	query := psql.Select(
		sm.Columns("id", "name"),
		sm.From("users"),
	)
	if filter.Name != "" {
		query.Apply(
			sm.WhereC("name = ?", filter.Name),
		)
	}
	return util.PrintQuery(query, nil)
}
