package util

import (
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/sq"
)

func PrintQuery(q litsql.BuildQuery, params map[string]any, writerOptions ...litsql.WriterOption) error {
	squery, args, err := q.Build(writerOptions...)
	if err != nil {
		return err
	}
	fmt.Println(strings.Repeat("=", 15), "QUERY", strings.Repeat("=", 15))
	fmt.Println(squery)
	fmt.Println(strings.Repeat("-", 15), "QUERY ARGS", strings.Repeat("-", 15))
	spew.Dump(args)
	fmt.Println(strings.Repeat("+", 15), "PARSED ARGS", strings.Repeat("+", 15))
	parsedArgs, err := sq.ParseArgs(args, sq.MapArgValues(params))
	if err != nil {
		return err
	}
	spew.Dump(parsedArgs)
	return nil
}
