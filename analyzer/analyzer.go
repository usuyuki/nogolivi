package analyzer

import (
	"fmt"
	"regexp"
	"runtime"
)

func Analyze() {
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

	regex := regexp.MustCompile(`goroutine \d+ \[.*\]:\n`)

	// mainのgoroutineが存在するため最低1はある→nil判定は不要
	matches := regex.FindAllStringIndex(string(buf[:n]), -1)
	goroutineCount := len(matches)

	// fmt.Printf("%s\n", buf[:n])
	if goroutineCount == 0 {
		panic("ParseError")
	}

	if goroutineCount == 1 {
		fmt.Println("\n🟢OK\n")
		fmt.Println("No living goroutines except main goroutine")
	} else {
		fmt.Println("\n❌NG\n")
		fmt.Printf("Number of remaining goroutines excluding the main goroutine: %d ", goroutineCount-1)
		if n == bufLen {
			fmt.Println("and more")
		} else {
			// 改行調整用
			fmt.Println("")
		}
	}
	return
}
