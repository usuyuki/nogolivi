package testHelper

import "github.com/usuyuki/nogolivi/traceStruct"

func GetGoroutineStatusesStructExample() []traceStruct.GoroutineStatus {
	return []traceStruct.GoroutineStatus{
		{
			Id:     18,
			Status: "sleep",
			Parent: traceStruct.ParentGoroutine{
				Id:           1,
				FunctionName: "main.main",
				FileName:     "/home/user/source_code/nogolivi/_examples/go_living.go",
				LineNumber:   18,
			},
			StackTrace: []traceStruct.StackTrace{
				{
					FunctionName: "time.Sleep(0x3b9aca00)",
					FileName:     "/home/user/.asdf/installs/golang/1.21rc2/go/src/runtime/time.go",
					LineNumber:   195,
				},
				{
					FunctionName: "main.main.func1(0x1)",
					FileName:     "/home/user/source_code/nogolivi/_examples/go_living.go",
					LineNumber:   19,
				},
			},
		},
		{
			Id:     1,
			Status: "running",
			Parent: traceStruct.ParentGoroutine{
				Id:           0,
				FunctionName: "",
				FileName:     "",
				LineNumber:   0,
			},
			StackTrace: []traceStruct.StackTrace{
				{
					FunctionName: "github.com/usuyuki/nogolivi/getter.GetTrace()",
					FileName:     "/home/user/source_code/nogolivi/getter/getTrace.go",
					LineNumber:   15,
				},
				{
					FunctionName: "github.com/usuyuki/nogolivi.Trace()",
					FileName:     "/home/user/source_code/nogolivi/nogolivi.go",
					LineNumber:   40,
				},
			},
		},
	}
}
