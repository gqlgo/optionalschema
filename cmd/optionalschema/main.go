package main

import (
	"flag"
	"github.com/gqlgo/gqlanalysis/multichecker"
	"github.com/gqlgo/optionalschema"
)

func main() {
	var excludes string
	flag.StringVar(&excludes, "excludes", "", "exclude GraphQL schema type name. it can specify multiple values separated by `,`")
	flag.Parse()

	multichecker.Main(
		optionalschema.Analyzer(excludes),
	)
}
