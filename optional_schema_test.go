package optionalschema_test

import (
	"testing"

	"github.com/gqlgo/gqlanalysis/analysistest"
	"github.com/gqlgo/optionalschema"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData(t)
	analysistest.Run(t, testdata, optionalschema.Analyzer("excludeOptionalTypeField"), "a")
}
