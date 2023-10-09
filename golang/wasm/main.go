//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"syscall/js"

	"github.com/FulecoRafa/merge_sort_wasm/golang/lib"
)

func wrapMergeSort() js.Func {
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
		ordered := lib.MergeSort(data)

        //Convert result back
		converted := make([]interface{}, len(ordered))
		for i := range ordered {
			converted[i] = ordered[i]
		}
		return js.ValueOf(converted)
	})
	return helperFunc
}

func main() {
	js.Global().Set("GoMergeSort", wrapMergeSort())
	<-make(chan struct{})
}
