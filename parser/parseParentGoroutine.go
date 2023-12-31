package parser

import (
	"regexp"
	"strconv"

	"github.com/usuyuki/nogolivi/traceStruct"
)

// 親goroutine情報のパース
func parseParentGoroutine(parentLine []string) (parent traceStruct.ParentGoroutine) {
	if len(parentLine) != 2 {
		panic("ParseParent failed. Parent goroutine info is not 2 lines")
	}
	// 1行目で関数名とgoroutineIDが取れる
	// created by main.main in goroutine 1  的な内容
	re1 := regexp.MustCompile(`created by (.+) in goroutine (\d+)`)
	match1 := re1.FindStringSubmatch(parentLine[0])
	if len(match1) != 3 {
		panic("ParseParent failed. Cannot parse parent goroutine functionName or id")
	}
	functionName := match1[1]
	goroutineID, _ := strconv.Atoi(match1[2])

	// 2行目でファイル名と行数が取れる
	//         /home/user/source_code/nogolivi/examples/go_living.go:18 +0x59   的な内容
	re2 := regexp.MustCompile(`^\s*(.+):(\d+)\s`)
	match2 := re2.FindStringSubmatch(parentLine[1])
	if len(match2) != 3 {
		panic("ParseParent failed. Cannot parse parent goroutine fileName or lineNumber")
	}
	fileName := match2[1]
	lineNumber, _ := strconv.Atoi(match2[2])

	parent = traceStruct.ParentGoroutine{
		FunctionName: functionName,
		Id:           goroutineID,
		FileName:     fileName,
		LineNumber:   lineNumber,
	}

	return
}
