package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	traceString := `goroutine 1 [running]:
github.com/usuyuki/nogolivi/getter.GetTrace()
        /home/user/source_code/nogolivi/getter/getTrace.go:15 +0x3c
github.com/usuyuki/nogolivi.Trace()
        /home/user/source_code/nogolivi/nogolivi.go:40 +0x213
main.main()
        /home/user/source_code/nogolivi/_examples/go_living.go:26 +0x14f

goroutine 18 [sleep]:
time.Sleep(0x3b9aca00)
        /home/user/.asdf/installs/golang/1.21rc2/go/src/runtime/time.go:195 +0x125
main.main.func1(0x1)
        /home/user/source_code/nogolivi/_examples/go_living.go:19 +0x26
created by main.main in goroutine 1
        /home/user/source_code/nogolivi/_examples/go_living.go:18 +0x59

goroutine 118 [running]:
github.com/usuyuki/nogolivi/parser.parseStackTrace({0xc000218010, 0x6, 0x4dc250?})
        /home/user/source_code/nogolivi/parser/parseStackTrace.go:23 +0x2da
github.com/usuyuki/nogolivi/parser.Parse.func1({0xc00018a800?, 0x0?})
        /home/user/source_code/nogolivi/parser/parse.go:46 +0x254
created by github.com/usuyuki/nogolivi`
	parseResult := Parse(traceString)
	assert.ElementsMatch(t, []GoroutineStatus{
		{
			Id:     18,
			Status: "sleep",
			Parent: ParentGoroutine{
				Id:           1,
				FunctionName: "main.main",
				FileName:     "/home/user/source_code/nogolivi/_examples/go_living.go",
				LineNumber:   18,
			},
			StackTrace: []StackTrace{
				{
					FunctionName: "time.Sleep(0x3b9aca00)",
					FileName:     "/home/user/.asdf/installs/golang/1.21rc2/go/src/runtime/time.go",
					LineNumber:   195,
				},
				{
					FunctionName: "main.main.func1(0x1)",
					FileName:     "/home/user/source_code/nogolivi/_examples/go_living.go",
					LineNumber:   19,
				},
			},
		},
		{
			Id:     1,
			Status: "running",
			Parent: ParentGoroutine{
				Id:           0,
				FunctionName: "",
				FileName:     "",
				LineNumber:   0,
			},
			StackTrace: []StackTrace{
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
			},
		},
	}, parseResult)
}
