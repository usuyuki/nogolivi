package parser

import (
	"regexp"
	"strconv"

	"github.com/usuyuki/nogolivi/traceStruct"
)

// スタックトレースのパース
func parseStackTrace(trace []string) (stackTraces []traceStruct.StackTrace) {
	// 配列の長さが2の倍数でない場合はpanic
	if len(trace)%2 != 0 {
		panic("ParseStackTrace failed. Trace length is not even")
	}

	// 2行ずつでループ
	for i := 0; i < len(trace); i += 2 {
		// 1行目の処理(関数が取れる) ここはそのままでOK
		functionName := trace[i]

		// 2行目の処理(ファイル名と行数が取れる)
		//         /home/user/source_code/nogolivi/examples/go_living.go:19 +0x26 のような文字列をパースする
		// .goだけでなくruntime/asm_amd64.s:1650のようなパターンも有る
		re2 := regexp.MustCompile(`([^:\s]+):(\d+)`)
		match2 := re2.FindStringSubmatch(trace[i+1])

		if len(match2) != 3 {
			panic("ParseStackTrace failed. Cannot parse fileName or lineNumber")
		}
		fileName := match2[1]
		lineNumber, _ := strconv.Atoi(match2[2])

		stackTraces = append(stackTraces,
			traceStruct.StackTrace{
				FunctionName: functionName,
				FileName:     fileName,
				LineNumber:   lineNumber,
			},
		)
	}
	return
}
