package parser

import (
	"strings"
	"sync"

	"github.com/usuyuki/nogolivi/traceStruct"
)

// スタックトレースをパースする
func Parse(data string) (goroutineStatuses []traceStruct.GoroutineStatus) {
	// goroutineごとに2つ改行があるのでそれを切れ目として判断

	// fmt.Println(data)
	goroutines := strings.Split(data, "\n\n")
	numGoroutines := len(goroutines)

	// goroutineごとにパース
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i, goroutine := range goroutines {

		// パースした最後は完全な文字列であることが保証できないためスキップ
		if i == numGoroutines-1 {
			break
		}

		wg.Add(1)
		// goroutineの中身を改行で分割
		go func(goroutine string) {
			defer wg.Done()

			// 各goroutineのトレース情報を改行で分割
			lines := strings.Split(goroutine, "\n")
			numLength := len(lines)

			// 1行目の処理(goroutineIDとStatusが取れる)
			id, status := parseIdStatus(lines[0])

			var parentGoroutine traceStruct.ParentGoroutine
			var stackTraces []traceStruct.StackTrace

			// rootのgoroutineは親情報がないのでその分岐
			// 2行目~後ろから3行目までの処理(スタックトレースが取れる)
			stackTraces = parseStackTrace(lines[1 : numLength-2])

			// goroutine1はmain goroutineなので親goroutineの情報が含まれないのでスキップ
			if id != 1 {
				// 後ろ2行目までの処理(スタックトレースが取れる)
				parentGoroutine = parseParentGoroutine(lines[numLength-2 : numLength])
			}

			mu.Lock()
			goroutineStatuses = append(goroutineStatuses,
				traceStruct.GoroutineStatus{
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
