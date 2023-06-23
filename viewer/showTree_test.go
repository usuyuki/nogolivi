package viewer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/usuyuki/nogolivi/testHelper"
)

func TestShowTreeNoPanic(t *testing.T) {
	assert.NotPanics(t,
		func() { ShowTree(testHelper.GetGoroutineStatusesStructExample()) },
	)
}
