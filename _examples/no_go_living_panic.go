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

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(num int) {
			atomic.AddInt64(&sum, int64(num))
			panic("hoge")
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("Sum:", sum)
}
