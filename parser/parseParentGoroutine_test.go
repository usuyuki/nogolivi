package parser

import (
	"fmt"
	"testing"
)

func TestParseRunning(t *testing.T) {
	data := [2]string{
		"created by main.main in goroutine 1",
		"        /home/user/source_code/nogolivi/examples/go_living.go:18 +0x59",
	}
	parentGoroutine := parser.parseParentGoroutine(data)
	if parentGoroutine.Id != 1 {
		t.Errorf("id is not 1")
	}
	if parentGoroutine.FunctionName != "main.main" {
		t.Errorf("functionName is not main.main")
	}
	if parentGoroutine.FileName != "/home/user/source_code/nogolivi/examples/go_living.go" {
		t.Errorf("fileName is not /home/user/source_code/nogolivi/examples/go_living.go")
	}
	if parentGoroutine.LineNumber != 18 {
		t.Errorf("lineNumber is not 18")
	}
}

func TestInvalidStringFirstLine(t *testing.T) {
	defer func() {
		err := recover()
		fmt.Println(err)
		if err != "parseIdStatus failed. Cannot parse goroutineID and status" {
			t.Errorf("no proper panic")
		}
	}()
	data := [2]string{
		"てきとうな文字列",
		"        /home/user/source_code/nogolivi/examples/go_living.go:18 +0x59",
	}
	parser.parseParentGoroutine(data)
}

func TestInvalidStringSecondLine(t *testing.T) {
	defer func() {
		err := recover()
		fmt.Println(err)
		if err != "parseIdStatus failed. Cannot parse goroutineID and status" {
			t.Errorf("no proper panic")
		}
	}()
	data := [2]string{
		"created by main.main in goroutine 1",
		"てきとうな文字列",
	}
	parser.parseParentGoroutine(data)
}

func TestInvalidLenght(t *testing.T) {
	defer func() {
		err := recover()
		fmt.Println(err)
		if err != "parseIdStatus failed. Cannot parse goroutineID and status" {
			t.Errorf("no proper panic")
		}
	}()
	data := [1]string{
		"created by main.main in goroutine 1",
	}
	parser.parseParentGoroutine(data)
}
