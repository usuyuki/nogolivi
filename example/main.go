package main

import (
	"fmt"

	"github.com/usuyuki/nogolivi"
)

func main() {
	defer nogolivi.Trace()

	fmt.Println("色々な処理")
}
