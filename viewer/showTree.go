package viewer

import (
	"fmt"
	"os"

	pt "github.com/bayashi/go-proptree"
	"github.com/usuyuki/nogolivi/parser"
)

func ShowTree(parseResult []parser.GoroutineStatus) {
	tree := pt.Node("Main Goroutine").Icon("*").Tag("1")

	// ツリーの各要素を追加
	for _, goroutine := range parseResult {
		// 親goroutineの情報を追加(mainゴルーチンは除外)
		if goroutine.Id == 1 {
			break
		}
		node := pt.Node(fmt.Sprintf("Goroutine %d", goroutine.Id)).Tag(goroutine.Status)

		// スタックトレースを追加
		for _, stackTrace := range goroutine.StackTrace {
			stackNode := pt.Node(fmt.Sprintf("%s (%s:%d)", stackTrace.FunctionName, stackTrace.FileName, stackTrace.LineNumber))
			node.Append(stackNode)
		}

		parentNode := pt.Node(fmt.Sprintf("Parent: %s (%s:%d)", goroutine.Parent.FunctionName, goroutine.Parent.FileName, goroutine.Parent.LineNumber))
		node.Append(parentNode)

		tree.Append(node)
	}

	tree.RenderAsText(os.Stdout)
	return
}
