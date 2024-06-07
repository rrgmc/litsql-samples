package main

import (
	"github.com/rrgmc/litsql-samples/util"
	"github.com/rrgmc/litsql/dialect/psql"
	"github.com/rrgmc/litsql/expr"
)

func sample() error {
	query := psql.SelectRawExpr(expr.Clause("select * from users where user_id = ?", 55))
	return util.PrintQuery(query, nil)
}
