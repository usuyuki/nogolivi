package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

// スタックトレースをパースする
func Parse(data string) (goroutineStatuses []GoroutineStatus) {
	// goroutineごとに2つ改行があるのでそれを切れ目として判断
	goroutines := strings.Split(data, "\n\n")

	// goroutineごとにパース
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, goroutine := range goroutines {

		wg.Add(1)
		// goroutineの中身を改行で分割
		go func(goroutine string) {
			defer wg.Done()

			length := len(goroutine)
			// 各goroutineのトレース情報を改行で分割
			lines := strings.Split(goroutine, "\n")
			fmt.Println("lines", len(lines))

			// 1行目の処理(goroutineIDとStatusが取れる)
			id, status := parseIdStatus(lines[0])

			var parentGoroutine ParentGoroutine
			var stackTraces []StackTrace

			// goroutine1はmain goroutineなので親goroutineの情報が含まれない
			if id != 1 {
				// 2行目~後ろから3行目までの処理(スタックトレースが取れる)
				stackTraces = parseStackTrace(lines[1 : length-2])

				// 後ろ2行目までの処理(スタックトレースが取れる)
				parentGoroutine = parseParent(lines[length-2 : length])
			}

			mu.Lock()
			goroutineStatuses = append(goroutineStatuses,
				GoroutineStatus{
					Id:         id,
					Status:     status,
					Parent:     parentGoroutine,
					StackTrace: stackTraces,
				},
			)
			mu.Unlock()
		}(goroutine)
	}
	wg.Wait()

	return
}

func parseIdStatus(firstLine string) (goroutineID int, goroutineStatus string) {
	re := regexp.MustCompile(`goroutine (\d+) \[(\w+)\]:`)
	match := re.FindStringSubmatch(firstLine)
	if len(match) > 0 {
		goroutineID, _ = strconv.Atoi(match[1])
		goroutineStatus = match[2]
	} else {
		panic("parseIdStatus failed. Cannot parse goroutineID and status")
	}
	return
}

// スタックトレースのパース
func parseStackTrace(trace []string) (stackTraces []StackTrace) {
	// 配列の長さが2の倍数でない場合はpanic
	if len(trace)%2 != 0 {
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

// 親goroutine情報のパース
func parseParent(parentLine []string) (parent ParentGoroutine) {
	if len(parentLine) != 2 {
		panic("parseParent failed. Parent goroutine info is not 2 lines")
	}
	// 1行目で関数名とgoroutineIDが取れる
	// created by main.main in goroutine 1  的な内容
	re1 := regexp.MustCompile(`created by (.+) in goroutine (\d+)`)
	match1 := re1.FindStringSubmatch(parentLine[0])
	if len(match1) != 3 {
		panic("parseParent failed. Cannot parse parent goroutine functionName or id")
	}
	functionName := match1[1]
	goroutineID, _ := strconv.Atoi(match1[2])

	// 2行目でファイル名と行数が取れる
	//         /home/user/source_code/nogolivi/examples/go_living.go:18 +0x59   的な内容
	re2 := regexp.MustCompile(`\t(.+) \((.+):(\d+)\)`)
	match2 := re2.FindStringSubmatch(parentLine[1])
	if len(match2) != 4 {
		panic("parseParent failed. Cannot parse parent goroutine fileName or lineNumber")
	}
	fileName := match2[2]
	lineNumber, _ := strconv.Atoi(match2[3])

	parent = ParentGoroutine{
		FunctionName: functionName,
		Id:           goroutineID,
		FileName:     fileName,
		LineNumber:   lineNumber,
	}

	return
}
