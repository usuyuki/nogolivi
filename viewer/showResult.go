package viewer

import "fmt"

func ShowResult(message []string) {
	for _, value := range message {
		fmt.Println(value)
	}
}
