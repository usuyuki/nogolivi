package nogolivi

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

	fmt.Println("\n=== Check Started ===")

	// 解析
	goroutineCount, message := analyzer.Analyze()

	// 結果表示
	fmt.Println("Number of remaining goroutines: ", goroutineCount)
	for _, value := range message {
		fmt.Println(value)
	}

	fmt.Println("\n===   Check End   ===")
}
