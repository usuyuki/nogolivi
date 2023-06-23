package main

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/usuyuki/nogolivi"
)

func main() {
	defer nogolivi.Trace()

	var sum int64

	for i := 1; i <= 5; i++ {
		// wg.Add(1)  ←本来はwaitGroupなどを使って待つことが必要
		go func(num int) {
			for j := 1; j <= 100; j++ {
				go func(num int) {
					time.Sleep(1 * time.Second)
					atomic.AddInt64(&sum, int64(num))
					// wg.Done()
				}(j)
			}
		}(i)
	}

	fmt.Println("Sum:", sum)
}
