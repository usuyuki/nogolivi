package analyzer

import (
	"fmt"
	"runtime"
)

func Analyze() {
	fmt.Println("Analyze")

	// 見切れてる
	buf := make([]byte, 2048)
	n := runtime.Stack(buf, true)

	// bufに格納された結果をどうパースする？ 正規表現？

	fmt.Printf("%s\n", buf[:n])
}
