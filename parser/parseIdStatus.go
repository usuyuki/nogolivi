package parser

import (
	"regexp"
	"strconv"
)

// idとstatusを取得
func parseIdStatus(firstLine string) (goroutineID int, goroutineStatus string) {
	re := regexp.MustCompile(`goroutine (\d+) \[(\w+)\]:`)
	match := re.FindStringSubmatch(firstLine)
	if len(match) > 0 {
		goroutineID, _ = strconv.Atoi(match[1])
		goroutineStatus = match[2]
	} else {
		panic("parseIdStatus failed. Cannot parse goroutineID and status")
	}
	return
}
