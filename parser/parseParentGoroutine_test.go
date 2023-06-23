package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParentGoroutineSuccess(t *testing.T) {
	data := []string{
		"created by main.main in goroutine 1",
		"        /home/user/source_code/nogolivi/examples/go_living.go:18 +0x59",
	}
	parentGoroutine := parseParentGoroutine(data)
	assert.Exactly(t, ParentGoroutine{
		FunctionName: "main.main",
		Id:           1,
		FileName:     "/home/user/source_code/nogolivi/examples/go_living.go",
		LineNumber:   18,
	}, parentGoroutine)
}

func TestParentGoroutineInvalidStringFirstLine(t *testing.T) {
	data := []string{
		"てきとうな文字列",
		"        /home/user/source_code/nogolivi/examples/go_living.go:18 +0x59",
	}
	assert.Panics(t,
		func() { parseParentGoroutine(data) },
		"ParseParent failed. Cannot parse parent goroutine fileName or lineNumber",
	)
}

func TestParentGoroutineInvalidStringSecondLine(t *testing.T) {
	data := []string{
		"created by main.main in goroutine 1",
		"てきとうな文字列",
	}
	assert.Panics(t,
		func() { parseParentGoroutine(data) },
		"ParseParent failed. Cannot parse parent goroutine fileName or lineNumber",
	)
}

func TestParentGoroutineInvalidLenght(t *testing.T) {
	data := []string{
		"てきとうな文字列",
	}

	assert.Panics(t,
		func() { parseParentGoroutine(data) },
		"ParseParent failed. Parent goroutine info is not 2 lines ",
	)
}
