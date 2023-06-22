package analyzer

import (
	"runtime"
	"strconv"
)

func Analyze() (goroutineCount int, message []string) {
	// そもそもgoroutineが1つ(main関数だけ)なら早期returnさせる
	goroutineCount = runtime.NumGoroutine()
	if goroutineCount == 1 {
		message = append(message, "\n🟢OK\n", "No living goroutines except main goroutine")
		return
	}

	bufLen := 2048

	// 見切れてる
	buf := make([]byte, bufLen)
	// buffがいっぱいになったか？は分かる→ 元のスライスの長さ=帰ってくる長さが同じ時
	// runtime.Traceを出して、それをパースする？
	// 途中で切れてたら and more的なものを出す
	// sync.Poolを使っていい感じにする

	/**
	 *  @see https://tip.golang.org/doc/go1.21#runtime
	 *  @see https://pkg.go.dev/runtime#Stack
	 */
	n := runtime.Stack(buf, true)

	// fmt.Printf("%s\n", buf[:n])
	if goroutineCount == 0 {
		panic("ParseError")
	}

	message = append(message, "\n❌NG\n", "Number of remaining goroutines excluding the main goroutine: "+strconv.Itoa(goroutineCount-1))
	if n == bufLen {
		message = append(message, "and more")
	}
	return
}
