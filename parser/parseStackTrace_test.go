package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseStackTrace1(t *testing.T) {
	data := []string{
		"time.Sleep(0x3b9aca00)",
		"        /home/user/.asdf/installs/golang/1.21rc2/go/src/runtime/time.go:195 +0x125",
		"main.main.func1(0x2)",
		"        /home/user/source_code/nogolivi/examples/go_living.go:19 +0x26",
	}
	stackTraces := parseStackTrace(data)
	assert.Exactly(t, []StackTrace{
		{
			FunctionName: "time.Sleep(0x3b9aca00)",
			FileName:     "/home/user/.asdf/installs/golang/1.21rc2/go/src/runtime/time.go",
			LineNumber:   195,
		},
		{
			FunctionName: "main.main.func1(0x2)",
			FileName:     "/home/user/source_code/nogolivi/examples/go_living.go",
			LineNumber:   19,
		},
	}, stackTraces)
}

func TestParseStackTrace2(t *testing.T) {
	data := []string{
		"github.com/usuyuki/nogolivi/getter.GetTrace()",
		"        /home/user/source_code/nogolivi/getter/getTrace.go:15 +0x3c",
		"github.com/usuyuki/nogolivi.Trace()",
		"        /home/user/source_code/nogolivi/nogolivi.go:40 +0x213",
		"main.main()",
		"        /home/user/source_code/nogolivi/_examples/go_living.go:26 +0x14f",
	}
	stackTraces := parseStackTrace(data)
	assert.Exactly(t, []StackTrace{
		{
			FunctionName: "github.com/usuyuki/nogolivi/getter.GetTrace()",
			FileName:     "/home/user/source_code/nogolivi/getter/getTrace.go",
			LineNumber:   15,
		},
		{
			FunctionName: "github.com/usuyuki/nogolivi.Trace()",
			FileName:     "/home/user/source_code/nogolivi/nogolivi.go",
			LineNumber:   40,
		},
		{
			FunctionName: "main.main()",
			FileName:     "/home/user/source_code/nogolivi/_examples/go_living.go",
			LineNumber:   26,
		},
	}, stackTraces)
}

func TestParseStackTraceInvalidStringSecondLine(t *testing.T) {
	data := []string{
		"time.Sleep(0x3b9aca00)",
		"あああ",
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
