package parser

// goroutineごとにこの構造体を持つ
type GoroutineStatus struct {
	// goroutineのID
	id int
	// goroutineの状態( running, runnable, waiting, dead)
	// @todo ここ列挙型にしたい
	status string
	// 親のgoroutineのID
	parentId int
	// goroutineのスタックトレース
	stackTrace []stackTrace
}

// スタックトレースの1つずつにこの構造体を持つ
type stackTrace struct {
	// 関数名
	functionName string
	// ファイル名
	fileName string
	// 行数
	lineNumber int
}

/*
スタックトレースの例 go 1.21rc1

goroutine 1 [running]:
github.com/usuyuki/nogolivi/analyzer.Analyze()
        /home/naofumi/source_code/nogolivi/analyzer/analyzer.go:32 +0x12b
github.com/usuyuki/nogolivi.Trace()
        /home/naofumi/source_code/nogolivi/nogolivi.go:22 +0x85
main.main()
        /home/naofumi/source_code/nogolivi/example/go_living.go:26 +0x14f

goroutine 6 [runnable]:
main.main.func2()
        /home/naofumi/source_code/nogolivi/example/go_living.go:18
runtime.goexit()
        /home/naofumi/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650 +0x1
created by main.main in goroutine 1
        /home/naofumi/source_code/nogolivi/example/go_living.go:18 +0x59

goroutine 7 [runnable]:
main.main.func2()
        /home/naofumi/source_code/nogolivi/example/go_living.go:18
runtime.goexit()
        /home/naofumi/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650 +0x1
created by main.main in goroutine 1
        /home/naofumi/source_code/nogolivi/example/go_living.go:18 +0x59

goroutine 8 [runnable]:
main.main.func2()
        /home/naofumi/source_code/nogolivi/example/go_living.go:18
runtime.goexit()
        /home/naofumi/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650 +0x1
created by main.main in goroutine 1
        /home/naofumi/source_code/nogolivi/example/go_living.go:18 +0x59


goroutine 9 [runnable]:
main.main.func2()
        /home/naofumi/source_code/nogolivi/example/go_living.go:18
runtime.goexit()
        /home/naofumi/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650 +0x1
created by main.main in goroutine 1
        /home/naofumi/source_code/nogolivi/example/go_living.go:18 +0x59

goroutine 10 [runnable]:
main.main.func2()
        /home/naofumi/source_code/nogolivi/example/go_living.go:18
runtime.goexit()

        /home/naofumi/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650 +0x1
created by main.main in goroutine 1
        /home/naofumi/source_code/nogolivi/example/go_living.go:18 +0x59

goroutine 11 [runnable]:
main.main.func2()
        /home/naofumi/source_code/nogolivi/example/go_living.go:18
runtime.goexit()
        /home/naofumi/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650 +0x1
created by main


*/
