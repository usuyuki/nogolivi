package nogolivi

import (
	"fmt"
	"runtime"
	"strconv"

	"github.com/usuyuki/nogolivi/getter"
	"github.com/usuyuki/nogolivi/parser"
	"github.com/usuyuki/nogolivi/viewer"
)

func Trace() {
	defer func() {
		// main関数のpanicはハンドリングできないので、あくまでnogolivi内部のpanicハンドリング用
		if r := recover(); r != nil {
			fmt.Println("Cannot display due to Panic '", r.(string), "' inside of nogolivi package")
			fmt.Println("It would be helpful if you could report it to contact@usuyuki.net.")
			return
		}
	}()

	fmt.Println("\n=== 🔥 Nogolivi Check Started 🔥 ===")

	var message []string
	goroutineCount := runtime.NumGoroutine()

	if goroutineCount == 1 {
		/*
			goroutineがmain goroutineのみの場合
		*/
		message = append(message, "\n🟢OK\n", "No living goroutines except main goroutine")
		// goroutineの残りがない場合はruntime.Stackなどを呼び出さず終える
		// 結果の表示
		viewer.ShowNg(goroutineCount-1, message)
	} else {
		/*
			goroutineがmain goroutine以外もある場合
		*/
		message = append(message, "\n❌NG\n", "Number of remaining goroutines excluding the main goroutine: "+strconv.Itoa(goroutineCount-1))

		// スタックトレースの取得 文字列を返す
		trace, isFull := getter.GetTrace()

		// スタックトレースのパース 解析結果を返す
		parseResult := parser.Parse(trace)

		// 結果の表示
		viewer.ShowNg(goroutineCount-1, message)
		// スタックトレースの表示
		viewer.ShowTree(parseResult)

		// 結果の判定
		if isFull {
			message = append(message, "and more")
		}
	}

	// 結果表示

	fmt.Println("\n===  🔥  Nogolivi Check End   🔥 ===")
}
