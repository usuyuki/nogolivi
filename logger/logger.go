package logger

import (
	"fmt"

	"github.com/usuyuki/nogolivi/analyzer"
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

	fmt.Println("=== Check Started ===")
	analyzer.Analyze()

	// goroutineが1、つまりmainしか残ってない場合はOKだけど、2つ以上ある場合はgoroutineの詳細を出す

	fmt.Println("=== Check End ===")
}
