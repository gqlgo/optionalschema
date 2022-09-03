# optionalschema

[![pkg.go.dev][gopkg-badge]][gopkg]

`optionalschema` finds optional fields and arguments in your GraphQL schema files.

```graphql
type Query {
    field: String!
    optionalField: String # want "optionalField is optional"
}
```

## How to use

A runnable linter can be created with multichecker package.
You can create own linter with your favorite Analyzers.

```go
package main

import (
	"flag"
	"github.com/gqlgo/optionalschema"
	"github.com/gqlgo/gqlanalysis/multichecker"
)

func main() {
	var excludes string
	flag.StringVar(&excludes, "excludes", "", "exclude GraphQL schema type name. it can specify multiple values separated by `,`")
	flag.Parse()

	multichecker.Main(
		optionalschema.Analyzer(excludes),
	)
}
```

`optionalschema` provides a typical main function and you can install with `go install` command.

```sh
$ go install github.com/gqlgo/optionalschema/cmd/optionalschema@latest
```

The `optionalschema` command has a flag, `schema` which will be parsed and analyzed by optionalschema's Analyzer.

```sh
$ optionalschema -schema="server/graphql/schema/**/*.graphql" -excludes="field1,field2"
```

The default value of `schema` is "schema/*/**.graphql".

## Author

[![Appify Technologies, Inc.](appify-logo.png)](http://github.com/appify-technologies)

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/gqlgo/optionalschema
[gopkg-badge]: https://pkg.go.dev/badge/github.com/gqlgo/optionalschema?status.svg
