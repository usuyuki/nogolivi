package getter

import (
	"runtime"
)

func GetTrace() (trace string, isFull bool) {
	bufLen := 2048
	buf := make([]byte, bufLen)

	/**
	 *  @see https://tip.golang.org/doc/go1.21#runtime
	 *  @see https://pkg.go.dev/runtime#Stack
	 */
	n := runtime.Stack(buf, true)
	// fmt.Printf("%s\n", buf[:n])

	if n == bufLen {
		isFull = true
	}
	trace = string(buf[:n])

	return
}
