package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseStackTraceRunning(t *testing.T) {
	data := []string{
		"time.Sleep(0x3b9aca00)",
		"        /home/user/.asdf/installs/golang/1.21rc2/go/src/runtime/time.go:195 +0x125",
		"main.main.func1(0x2)",
		"        /home/user/source_code/nogolivi/examples/go_living.go:19 +0x26",
	}
	stackTraces := parseStackTrace(data)
	assert.Exactly(t, []StackTrace{
		{
			FunctionName: "time.Sleep",
			FileName:     "/home/user/.asdf/installs/golang/1.21rc2/go/src/runtime/time.go",
			LineNumber:   195,
		},
		{
			FunctionName: "main.main.func1",
			FileName:     "/home/user/source_code/nogolivi/examples/go_living.go",
			LineNumber:   19,
		},
	}, stackTraces)
}

func TestParseStackTraceInvalidStringFirstLine(t *testing.T) {
	data := []string{
		"不適切な文字列",
		"        /home/user/.asdf/installs/golang/1.21rc2/go/src/runtime/time.go:195 +0x125",
		"main.main.func1(0x2)",
		"        /home/user/source_code/nogolivi/examples/go_living.go:19 +0x26",
	}
	assert.Panics(t,
		func() { parseStackTrace(data) },
		"ParseStackTrace failed. Cannot parse functionName",
	)
}

func TestParseStackTraceInvalidStringLastLine(t *testing.T) {
	data := []string{
		"time.Sleep(0x3b9aca00)",
		"        /home/user/.asdf/installs/golang/1.21rc2/go/src/runtime/time.go:195 +0x125",
		"main.main.func1(0x2)",
		"不適切な文字列",
	}
	assert.Panics(t,
		func() { parseStackTrace(data) },
		"ParseStackTrace failed. Cannot parse fileName or lineNumber",
	)
}

func TestParseStackTraceInvalidLenght(t *testing.T) {
	data := []string{
		"time.Sleep(0x3b9aca00)",
	}

	assert.Panics(t,
		func() { parseStackTrace(data) },
		"ParseStackTrace failed. Trace length is not even",
	)
}
