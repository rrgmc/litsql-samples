//nolint:unused
package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rrgmc/litsql/dialect/psql"
	"github.com/rrgmc/litsql/dialect/psql/sm"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/sq"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("pgx", "postgres://postgres:sakila@localhost:5455/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = runPrepared(ctx, db)
	if err != nil {
		panic(err)
	}
}

func run(ctx context.Context, db *sql.DB) error {
	query := psql.Select(
		sm.Columns("film_id", "title", "length"),
		sm.From("film"),
		sm.WhereClause("length > ?", sq.NamedArg("length")),
		sm.Limit(10),
	)

	squery, params, err := query.Build()
	if err != nil {
		return err
	}

	fmt.Println(squery)

	args, err := sq.ParseArgs(params, sq.MapArgValues{
		"length": 100,
	})
	if err != nil {
		return err
	}

	rows, err := db.QueryContext(ctx, squery, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id, length int
		var title string
		if err := rows.Scan(&id, &title, &length); err != nil {
			return err
		}
		fmt.Println(id, title, length)
	}

	if rows.Err() != nil {
		return rows.Err()
	}

	return nil
}

func runPrepared(ctx context.Context, db *sql.DB) error {
	query := psql.Select(
		sm.Columns("film_id", "title", "length"),
		sm.From("film"),
		sm.WhereClause("length > ?", sq.NamedArg("length")),
		sm.LimitExpr(expr.ArgNamed("limit")),
	)

	queryStr, args, err := query.Build()
	if err != nil {
		return err
	}

	prepq, err := db.PrepareContext(ctx, queryStr)
	if err != nil {
		return err
	}

	pargs, err := sq.ParseArgs(args, sq.MapArgValues{
		"length": 100,
		"limit":  10,
	})
	if err != nil {
		return err
	}

	rows, err := prepq.QueryContext(ctx, pargs...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id, length int
		var title string
		if err := rows.Scan(&id, &title, &length); err != nil {
			return err
		}
		fmt.Println(id, title, length)
	}

	if rows.Err() != nil {
		return rows.Err()
	}

	return nil
}

func runRaw(ctx context.Context, db *sql.DB) error {
	rows, err := db.QueryContext(ctx, "select film_id, title from film")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var title string
		if err := rows.Scan(&id, &title); err != nil {
			return err
		}
		fmt.Println(id, title)
	}

	return nil
}
