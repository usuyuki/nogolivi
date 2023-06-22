package main

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/usuyuki/nogolivi"
)

func main() {
	defer nogolivi.Trace()

	var sum int64
	var wg sync.WaitGroup

	for i := 1; i <= 100; i++ {
		go func(num int) {
			atomic.AddInt64(&sum, int64(num))
		}(i)
	}

	wg.Wait()
	fmt.Println("Sum:", sum)
}
