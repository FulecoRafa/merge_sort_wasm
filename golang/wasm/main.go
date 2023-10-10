//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"syscall/js"

	"github.com/FulecoRafa/merge_sort_wasm/golang/lib"
)

func wrapMerge(mergeFunc func([]int) []int) js.Func {
	helperFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// Check if arguments are correct
		if len(args) != 1 {
			return "invalid params"
		}

		arg := args[0]
		if arg.Type() != js.TypeObject {
			fmt.Println("the first argument should be an array")
			return nil
		}

		// Convert arguments
		data := make([]int, arg.Length())
		for i := range data {
			item := arg.Index(i)
			if item.Type() != js.TypeNumber {
				fmt.Println("Arg ", i, " is not a number")
				return nil
			}
			data[i] = item.Int()
		}

		// Lib execution
		ordered := mergeFunc(data)

		//Convert result back
		converted := make([]interface{}, len(ordered))
		for i := range ordered {
			converted[i] = ordered[i]
		}
		return js.ValueOf(converted)
	})
	return helperFunc
}

func wrapMergeSort() js.Func {
	return wrapMerge(lib.MergeSort[int])
}

func wrapMergeSortParallel() js.Func {
	return wrapMerge(func(data []int) []int {
		ch := make(chan []int, 1)
		lib.ParallelMergeSort[int](data, ch)
		result := <-ch
		return result
	})

}

func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func wrapFibonacci() js.Func {
    return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        return fibonacci(args[0].Int())
    })
}

func main() {
	js.Global().Set("GoMergeSort", wrapMergeSort())
	js.Global().Set("GoParallelMergeSort", wrapMergeSortParallel())
	js.Global().Set("GoFibonacci", wrapFibonacci())
	<-make(chan struct{})
}
