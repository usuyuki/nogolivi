package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/usuyuki/nogolivi/testHelper"
)

func TestParse(t *testing.T) {
	traceString := testHelper.GetTraceStringExample()
	parseResult := Parse(traceString)
	assert.ElementsMatch(t, testHelper.GetGoroutineStatusesStructExample(), parseResult)
}
