package parser

import (
	"fmt"
	"regexp"
	"strconv"
)

// スタックトレースのパース
func parseStackTrace(trace []string) (stackTraces []StackTrace) {
	// 配列の長さが2の倍数でない場合はpanic
	if len(trace)%2 != 0 {
		fmt.Println(trace)
		panic("parseStackTrace failed. Trace length is not even")
	}

	// 2行ずつでループ
	for i := 0; i < len(trace); i += 2 {
		// 1行目の処理(関数が取れる)
		functionName := trace[i]
		// 2行目の処理(ファイル名と行数が取れる)
		re := regexp.MustCompile(`\t(.+) \((.+):(\d+)\)`)
		match := re.FindStringSubmatch(trace[i+1])

		if len(match) != 4 {
			fmt.Println(trace, i, trace[i], match)
			panic("parseStackTrace failed. Cannot parse fileName or lineNumber")
		}
		fileName := match[2]
		lineNumber, _ := strconv.Atoi(match[3])

		stackTraces = append(stackTraces,
			StackTrace{
				FunctionName: functionName,
				FileName:     fileName,
				LineNumber:   lineNumber,
			},
		)
	}
	return
}
