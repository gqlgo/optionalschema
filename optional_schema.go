package optionalschema

import (
	"github.com/gqlgo/gqlanalysis"
	"strings"
)

var excludes []string

func isExcludeField(filedName string) bool {
	for _, exclude := range excludes {
		if exclude == filedName {
			return true
		}
	}
	return false
}

func Analyzer(excludes string) *gqlanalysis.Analyzer {
	return &gqlanalysis.Analyzer{
		Name: "optionalschema",
		Doc:  "optionalschema finds a optionalschema in your GraphQL query files",
		Run:  run(excludes),
	}
}

func run(excludesStr string) func(pass *gqlanalysis.Pass) (interface{}, error) {
	excludes = strings.Split(excludesStr, ",")

	return func(pass *gqlanalysis.Pass) (interface{}, error) {
		for _, t := range pass.Schema.Types {
			if t.BuiltIn {
				continue
			}
			for _, field := range t.Fields {
				if field != nil && field.Type != nil {
					if !field.Type.NonNull && !isExcludeField(field.Name) {
						if field.Position != nil {
							pass.Reportf(field.Position, "%s is optional", field.Name)
						}
					}
					for _, argument := range field.Arguments {
						if argument != nil && argument.Type != nil {
							if !argument.Type.NonNull && !isExcludeField(argument.Name) {
								pass.Reportf(argument.Position, "%s is optional", argument.Name)
							}
						}
					}
				}
			}
		}

		return nil, nil
	}
}
