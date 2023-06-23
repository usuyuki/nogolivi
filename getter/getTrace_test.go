package getter

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetTraceBig(t *testing.T) {
	defer func() {
		trace, isFull := GetTrace()
		assert.True(t, isFull)
		// fmt.Println(trace)
		// 複数ある時、created by が絶対に出るので
		assert.Contains(t, trace, "created by")
	}()

	var sum int64

	for i := 1; i <= 100; i++ {
		// wg.Add(1)  ←本来はwaitGroupなどを使って待つことが必要
		go func(num int) {
			time.Sleep(1 * time.Second)
			atomic.AddInt64(&sum, int64(num))
			// wg.Done()
		}(i)
	}

	assert.Exactly(t, int64(0), sum)
}

// 本ライブラリのgoroutineも入るので、実際の挙動と違って再現が難しい
// func TestGetTraceSmall(t *testing.T) {
// 	defer func() {
// 		trace, isFull := GetTrace()
// 		fmt.Println(trace)
// 		assert.False(t, isFull)
// 		// fmt.Println(trace)
// 		assert.Contains(t, trace, "goroutine 1")
// 	}()
// 	var sum int64
// 	var wg sync.WaitGroup
//
// 	for i := 1; i <= 100; i++ {
// 		wg.Add(1)
// 		go func(num int) {
// 			atomic.AddInt64(&sum, int64(num))
// 			wg.Done()
// 		}(i)
// 	}
//
// 	wg.Wait()
// 	assert.Exactly(t, int64(5050), sum)
// }
