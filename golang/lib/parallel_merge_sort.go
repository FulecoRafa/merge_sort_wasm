package lib

import (
	"golang.org/x/exp/constraints"
)

func ParallelMergeSort[T constraints.Ordered](data []T, outChan chan<- []T) {
    size := len(data)
    if size <= 1 {
        outChan <- data
        return
    }

    result := make([]T, 0, size)

    mid := size / 2
    left := data[:mid]
    right := data[mid:]
    splitChan := make(chan []T, 2)
    go ParallelMergeSort(left, splitChan)
    go ParallelMergeSort(right, splitChan)

    for i := 0; i < 2; i++ {
        received := <- splitChan
        if i == 0 {
            left = received
        } else {
            right = received
        }
    }
    close(splitChan)

    // Indexes to both arrays
    l, r := 0, 0
    for l < len(left) && r < len(right) {
        if left[l] < right[r] {
            result = append(result, left[l])
            l++
        } else {
            result = append(result, right[r])
            r++
        }
    }

    for ; l < len(left); l++ {
        result = append(result, left[l])
    }

    for ; r < len(right); r++ {
        result = append(result, right[r])
    }

    outChan <- result
    return
}
