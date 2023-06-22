package getter

import (
	"fmt"
	"runtime"
)

func GetTrace() (trace string, isFull bool) {
	// そもそもgoroutineが1つ(main関数だけ)なら早期returnさせる

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

	fmt.Printf("%s\n", buf[:n])

	if n == bufLen {
		isFull = true
	}
	return
}
