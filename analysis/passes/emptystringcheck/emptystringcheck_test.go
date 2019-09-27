package emptystringcheck

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestEmptyStringCheck(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer)
}
