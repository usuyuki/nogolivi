# 🔥nogolivi (残り火)🔥

> 残り火は適切に消火して、 **no "go" livi**ng へ。

A tool for visualizing goroutines that are not terminated when the main function terminates.

main 関数終了時に終了していない goroutine の可視化を行うツールです。
<img src="logo.png">

## Attention

Go 1.21 の runtime.Stack の仕様変更に依存しているため、現在は go1.21RC2 環境での動作としています。

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
