package main

import (
	"github.com/rrgmc/litsql-samples/util"
	"github.com/rrgmc/litsql/dialect/psql"
	"github.com/rrgmc/litsql/dialect/psql/sm"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/sq"
)

func all() error {
	query := psql.Select(
		sm.With("c", "id", "data").
			As(psql.Select(
				sm.Columns("id"),
				sm.From("test1"),
				sm.LeftJoin("test2").Using("hhh", "jjjj"),
			)),
		sm.With("h", "id2", "data2").As(psql.Select(
			sm.Columns("id2"),
			sm.From("test12"),
		)),
		sm.Distinct("b", "c"),
		sm.Columns("a", "b", "c"),
		sm.Columns("d", "e", "f"),
		sm.FromQuery(psql.Select(
			sm.From("xt"),
			sm.Where("x = 12"),
			sm.WhereClause("h = ?", sq.NamedArg("x", sq.WithDefaultValue(9122))),
			// sm.WhereClause("j = ?", sql.Named("x", 444)),
			// sm.WhereClause("j = ?", sqlb.DBNamedArg("x")),
		)),
		sm.InnerJoin("device").As("x").On("d.x = d.y").On("abc = def"),
		sm.InnerJoin("double").As("h").On("h.j = x.t"),
		sm.Where("j = 5 AND k = 12"),
		sm.WhereClause("j IN ?", expr.InP([]any{sq.NamedArg("x"), 2, 3})),
		sm.WhereClause("h IN ?", psql.Select(
			sm.From("xxx"),
		)),
		sm.WhereClause("j = ? AND k = ?", sq.ArgFunc(func() (any, error) {
			return "99", nil
		}), sq.NamedArg("second")),
		sm.WhereExpr(
			expr.Or(
				"a = 5 AND b = 12",
				"t = 5 AND s = 12",
			),
		),
		sm.WhereExpr(
			expr.OrExpr(
				expr.Paren("a = 5 AND b = 12"),
				expr.Paren("t = 5 AND s = 12"),
			),
		),
		sm.Apply(func(a psql.SelectModApply) {
			a.Apply(
				sm.Columns("ii", "bb"),
				sm.Where("x = 15"),
			)
		}),
		sm.Window("abc").PartitionBy("depname").OrderBy("salary").From("uuu").Frame("GROUPS"),
		sm.Window("xyz").PartitionBy("tutor").OrderBy("body"),
		sm.GroupBy("a", "b").Distinct(),
		sm.Having("b > 12"),
		sm.OrderBy("b DESC", "c"),
		sm.OffsetArg(10),
		sm.LimitArg(99),
		sm.Union(psql.Select(
			sm.Columns("a", "b", "c"),
			sm.From("ttt111"),
		)),
		sm.Union(psql.Select(
			sm.Columns("t", "a", "x"),
			sm.From("uuuu8888"),
		)),
	)
	return util.PrintQuery(query, sq.MapArgValues{
		"x":      56,
		"second": 2,
	})
}
