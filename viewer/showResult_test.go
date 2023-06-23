package viewer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShowResultNoPanic(t *testing.T) {
	data := []string{
		"hoge",
		"fuga",
	}
	assert.NotPanics(t,
		func() { ShowResult(data) },
	)
}
