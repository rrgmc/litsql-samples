package util

import (
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/rrgmc/litsql/sq"
)

func PrintQuery(q sq.BuildQuery, params map[string]any, options ...sq.BuildQueryOption) error {
	squery, args, err := q.Build(options...)
	if err != nil {
		return err
	}
	fmt.Println(strings.Repeat("=", 15), "QUERY", strings.Repeat("=", 15))
	fmt.Println(squery)
	fmt.Println(strings.Repeat("-", 15), "QUERY ARGS", strings.Repeat("-", 15))
	spew.Dump(args)
	if params != nil {
		fmt.Println(strings.Repeat("+", 15), "PARSED ARGS", strings.Repeat("+", 15))
		parsedArgs, err := args.Parse(params)
		if err != nil {
			return err
		}
		spew.Dump(parsedArgs)
	}
	return nil
}
