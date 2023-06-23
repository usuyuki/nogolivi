# 🔥nogolivi (残り火)🔥

[![Go Reference](https://pkg.go.dev/badge/github.com/usuyuki/nogolivi.svg)](https://pkg.go.dev/github.com/usuyuki/nogolivi)

Incomplete goroutine visualization library for beginners  
(初学者向けの未完了 goroutine 可視化ライブラリ)

<img width="300" alt="SCR-20230502-nedr" src="logo.png">

> 残り火は適切に消火して、 **no "go" livi**ng へ。

## 概要

A library that easily completes the visualization of goroutines that are not terminated when the main function terminates.  
You can visualize unterminated goroutines by adding just two lines to your existing main function.

main 関数終了時に終了していない goroutine の可視化を簡単完結に行うライブラリです。  
既存の main 関数に 2 行追加するだけで終了していない goroutine を可視化できます。

```go
package main

import (
	"github.com/usuyuki/nogolivi" // add
)

func main() {
	defer nogolivi.Trace() // add

	// your code here
}

```

## Attention

Go 1.21 の runtime.Stack の仕様変更に依存しているため、現在は go1.21RC2 環境での動作としています。

### 詳細

Go 1.21 から runtime.Stack においてスタック トレース内の各 goroutine を作成した goroutine の ID が含まれるようになったことを利用.

https://tip.golang.org/doc/go1.21#runtime

go 1.20

```
goroutine 18 [runnable]:
main.main.func3()
	/tmp/sandbox1046556636/prog.go:25
runtime.goexit()
	/usr/local/go-faketime/src/runtime/asm_amd64.s:1598 +0x1
created by main.main
	/tmp/sandbox1046556636/prog.go:25 +0x71

```

go 1.21

```
goroutine 8 [runnable]:
main.main.func3()
	/tmp/sandbox910302531/prog.go:25
runtime.goexit()
	/usr/local/go-faketime/src/runtime/asm_amd64.s:1650 +0x1
created by main.main in goroutine 1
	/tmp/sandbox910302531/prog.go:25 +0x6a
```

https://go.dev/play/p/kh76htbRETg?v=gotip

go 1.21 で goroutine を作成した goroutine の id が分かることで、どこでどの用に途中で止まっているかをより分かりやすく可視化できるのでは？

## 機能

main 関数で defer によって呼び出すことで、main 関数終了時に終わっていない goroutine を可視化します。

### 提供するもの

- 初学者にありがちな goroutine の終了を待たずに main 関数を終了させてしまうケースの検知

### 提供しないもの

- go tool trace にあるような高度な状態の確認

## Test

単体テスト、結合テストともに対象と同じパッケージにおいています。

```shell
go test -shuffle=on -v -p=1 ./...

or

make test

```

## Usage

## example

### example 1

When there is some goroutines in the middle of execution when the main function is executed.

```shell
go run _examples/go_living.go

or

make ex1
```

```
go run _examples/go_living.go
Sum: 0

=== 🔥 Nogolivi Check Started 🔥 ===

❌NG

Number of remaining goroutines excluding the main goroutine: 5

┌* Main Goroutine: 1
├─┬ Goroutine 21: runnable
│ ├── main.main.func2() (/home/naofumi/source_code/nogolivi/_examples/go_living.go:18)
│ ├── runtime.goexit() (/home/naofumi/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650)
│ └── Parent: main.main (/home/naofumi/source_code/nogolivi/_examples/go_living.go:18)
├─┬ Goroutine 20: runnable
│ ├── main.main.func2() (/home/naofumi/source_code/nogolivi/_examples/go_living.go:18)
│ ├── runtime.goexit() (/home/naofumi/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650)
│ └── Parent: main.main (/home/naofumi/source_code/nogolivi/_examples/go_living.go:18)
└─┬ Goroutine 18: runnable
  ├── main.main.func2() (/home/naofumi/source_code/nogolivi/_examples/go_living.go:18)
  ├── runtime.goexit() (/home/naofumi/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650)
  └── Parent: main.main (/home/naofumi/source_code/nogolivi/_examples/go_living.go:18)

===  🔥  Nogolivi Check End   🔥 ===
```

### example 2

When there is a recursive goroutine in the middle of executing the main function.

```shell
go run _examples/go_living_recursively.go

or

make ex2
```

```
go run _examples/go_living_recursively.go

Sum: 0

=== 🔥 Nogolivi Check Started 🔥 ===

❌NG

Number of remaining goroutines excluding the main goroutine: 5
and more

┌* Main Goroutine: 1
├─┬ Goroutine 35: runnable
│ ├── main.main.func1.2() (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)
│ ├── runtime.goexit() (/home/naofumi/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650)
│ └── Parent: main.main.func1 (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)

├─┬ Goroutine 23: runnable
│ ├── main.main.func1.2() (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)
│ ├── runtime.goexit() (/home/naofumi/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650)
│ └── Parent: main.main.func1 (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)
├─┬ Goroutine 24: runnable
│ ├── main.main.func1.2() (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)
│ ├── runtime.goexit() (/home/naofumi/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650)
│ └── Parent: main.main.func1 (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)
├─┬ Goroutine 25: runnable
│ ├── main.main.func1.2() (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)
│ ├── runtime.goexit() (/home/naofumi/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650)
│ └── Parent: main.main.func1 (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)
├─┬ Goroutine 26: runnable
│ ├── main.main.func1.2() (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)
│ ├── runtime.goexit() (/home/naofumi/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650)
│ └── Parent: main.main.func1 (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)
├─┬ Goroutine 22: runnable
│ ├── main.main.func1.2() (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)
│ ├── runtime.goexit() (/home/naofumi/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650)
│ └── Parent: main.main.func1 (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)
├─┬ Goroutine 27: runnable
│ ├── main.main.func1.2() (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)
│ ├── runtime.goexit() (/home/naofumi/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650)
│ └── Parent: main.main.func1 (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)
├─┬ Goroutine 19: runnable
│ ├── main.main.func1.2() (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)
│ ├── runtime.goexit() (/home/naofumi/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650)
│ └── Parent: main.main.func1 (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)
├─┬ Goroutine 32: runnable
│ ├── main.main.func1.2() (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)
│ ├── runtime.goexit() (/home/naofumi/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650)
│ └── Parent: main.main.func1 (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)

├─┬ Goroutine 31: runnable
│ ├── main.main.func1.2() (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)
│ ├── runtime.goexit() (/home/naofumi/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650)
│ └── Parent: main.main.func1 (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)
├─┬ Goroutine 18: runnable
│ ├── main.main.func1.2() (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)

│ ├── runtime.goexit() (/home/naofumi/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650)
│ └── Parent: main.main.func1 (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)
├─┬ Goroutine 21: runnable
│ ├── main.main.func1.2() (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)
│ ├── runtime.goexit() (/home/naofumi/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650)
│ └── Parent: main.main.func1 (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)
├─┬ Goroutine 33: runnable
│ ├── main.main.func1.2() (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)
│ ├── runtime.goexit() (/home/naofumi/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650)
│ └── Parent: main.main.func1 (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)
├─┬ Goroutine 20: runnable
│ ├── main.main.func1.2() (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)
│ ├── runtime.goexit() (/home/naofumi/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650)
│ └── Parent: main.main.func1 (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)
├─┬ Goroutine 30: runnable
│ ├── main.main.func1.2() (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)
│ ├── runtime.goexit() (/home/naofumi/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650)

│ └── Parent: main.main.func1 (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)
├─┬ Goroutine 10: runnable

│ ├── main.main.func2() (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:18)
│ ├── runtime.goexit() (/home/naofumi/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650)
│ └── Parent: main.main (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:18)
└─┬ Goroutine 17: runnable
  ├── main.main.func1.2() (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)
  ├── runtime.goexit() (/home/naofumi/.asdf/installs/golang/1.21rc2/go/src/runtime/asm_amd64.s:1650)
  └── Parent: main.main.func1 (/home/naofumi/source_code/nogolivi/_examples/go_living_recursively.go:20)


===  🔥  Nogolivi Check End   🔥 ===
```

### example 3

When there is no goroutine in the middle of execution when the main function is executed.

```shell
go run _examples/no_go_living.go

or

make ex3
```

```
Sum: 15

=== 🔥 Nogolivi Check Started 🔥 ===

🟢OK

No living goroutines except main goroutine

===  🔥  Nogolivi Check End   🔥 ===

```
