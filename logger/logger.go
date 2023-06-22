package logger

import (
	"fmt"
)

func Trace() {
	// panicで終了したときは表示しない
	if r := recover(); r != nil {
		fmt.Println("No display due to panic")
		return
	}

	fmt.Println("=== Check Started ===")
	fmt.Println("=== Check End ===")
}
