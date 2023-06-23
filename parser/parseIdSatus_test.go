package parser

import (
	"fmt"
	"testing"
)

func Testnning(t *testing.T) {
	data := "goroutine 1 [running]:"
	id, status := parseIdStatus(data)
	if id != 1 {
		t.Errorf("id is not 1")
	}
	if status != "running" {
		t.Errorf("status is not running")
	}
}

func TestParseSleep(t *testing.T) {
	data := "goroutine 18 [sleep]:"
	id, status := parseIdStatus(data)
	if id != 18 {
		t.Errorf("id is not 18")
	}
	if status != "sleep" {
		t.Errorf("status is not sleep")
	}
}

func TestInvalidString(t *testing.T) {
	defer func() {
		err := recover()
		fmt.Println(err)
		if err != "parseIdStatus failed. Cannot parse goroutineID and status" {
			t.Errorf("no proper panic")
		}
	}()
	data := "不適切な文字列"
	parseIdStatus(data)
}
