package viewer

import "fmt"

func ShowNg(goroutineCount int, message []string) {
	fmt.Println("Number of remaining goroutines: ", goroutineCount)
	for _, value := range message {
		fmt.Println(value)
	}
}
