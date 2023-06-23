package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdStatusRnning(t *testing.T) {
	data := "goroutine 1 [running]:"
	id, status := parseIdStatus(data)
	assert.Exactly(t, 1, id)
	assert.Exactly(t, "running", status)
}

func TestIdStatusSleeping(t *testing.T) {
	data := "goroutine 18 [sleep]:"
	id, status := parseIdStatus(data)
	assert.Exactly(t, 18, id)
	assert.Exactly(t, "sleep", status)
}

func TestInvalidString(t *testing.T) {
	data := "不適切な文字列"
	assert.Panics(t,
		func() { parseIdStatus(data) },
		"parseIdStatus failed. Cannot parse goroutineID and status",
	)
}
