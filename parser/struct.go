package parser

// goroutineごとにこの構造体を持つ
type GoroutineStatus struct {
	// goroutineのID
	Id int
	// goroutineの状態( running, runnable, waiting, dead)
	// @todo ここ列挙型にしたい
	Status string
	// 親のgoroutineの情報
	Parent ParentGoroutine
	// goroutineのスタックトレース
	StackTrace []StackTrace
}

// スタックトレースの1つずつにこの構造体を持つ
type StackTrace struct {
	// 関数名
	FunctionName string
	// ファイル名
	FileName string
	// 行数
	LineNumber int
}

type ParentGoroutine struct {
	// goroutineのID
	Id int
	// 関数名
	FunctionName string
	// ファイル名
	FileName string
	// 行数
	LineNumber int
}

/*
スタックトレースの例 go 1.21rc1

goroutine 1 [running]:
github.com/usuyuki/nogolivi/analyzer.Analyze()
        /home/user/source_code/nogolivi/analyzer/analyzer.go:32 +0x12b
github.com/usuyuki/nogolivi.Trace()
        /home/user/source_code/nogolivi/nogolivi.go:22 +0x85
main.main()
        /home/user/source_code/nogolivi/example/go_living.go:26 +0x14f

goroutine 6 [runnable]:
main.main.func2()
        /home/user/source_code/nogolivi/example/go_living.go:18
runtime.goexit()
        /home/user/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650 +0x1
created by main.main in goroutine 1
        /home/user/source_code/nogolivi/example/go_living.go:18 +0x59

goroutine 7 [runnable]:
main.main.func2()
        /home/user/source_code/nogolivi/example/go_living.go:18
runtime.goexit()
        /home/user/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650 +0x1
created by main.main in goroutine 1
        /home/user/source_code/nogolivi/example/go_living.go:18 +0x59

goroutine 8 [runnable]:
main.main.func2()
        /home/user/source_code/nogolivi/example/go_living.go:18
runtime.goexit()
        /home/user/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650 +0x1
created by main.main in goroutine 1
        /home/user/source_code/nogolivi/example/go_living.go:18 +0x59


goroutine 9 [runnable]:
main.main.func2()
        /home/user/source_code/nogolivi/example/go_living.go:18
runtime.goexit()
        /home/user/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650 +0x1
created by main.main in goroutine 1
        /home/user/source_code/nogolivi/example/go_living.go:18 +0x59

goroutine 10 [runnable]:
main.main.func2()
        /home/user/source_code/nogolivi/example/go_living.go:18
runtime.goexit()

        /home/user/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650 +0x1
created by main.main in goroutine 1
        /home/user/source_code/nogolivi/example/go_living.go:18 +0x59

goroutine 11 [runnable]:
main.main.func2()
        /home/user/source_code/nogolivi/example/go_living.go:18
runtime.goexit()
        /home/user/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650 +0x1
created by main


*/
