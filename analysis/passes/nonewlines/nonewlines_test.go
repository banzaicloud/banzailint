package nonewlines

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestNoNewlines(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer)
}
