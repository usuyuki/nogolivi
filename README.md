# 🔥nogolivi (残り火)🔥

> 残り火は適切に消火して、 **no "go" livi**ng へ。

A tool for visualizing goroutines that are not terminated when the main function terminates.

main 関数終了時に終了していない goroutine の可視化を行うツールです。
<img src="logo.png">

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

https://go.dev/play/p/qSWvUEPKzrq?v=gotip

go 1.21 で goroutine を作成した goroutine の id が分かることで、どこでどの用に途中で止まっているかをより分かりやすく可視化できるのでは？

## 機能

main 関数で defer によって呼び出すことで、main 関数終了時に終わっていない goroutine を可視化します。

### 提供するもの

- 初学者にありがちな goroutine の終了を待たずに main 関数を終了させてしまうケースの検知

### 提供しないもの

- go tool trace にあるような高度な状態の確認

## Usage

```go
package main

import (
	"github.com/usuyuki/nogolivi" // add
)

func main() {
	defer nogolivi.Trace() // add

	// your code
}

```

## example

### example 1

When there is some goroutines in the middle of execution when the main function is executed.

```shell
go run example/go_living.go
```

### example 2

When there is no goroutine in the middle of execution when the main function is executed.

```shell
go run example/no_go_living.go
```
