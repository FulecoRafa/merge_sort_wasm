package lib

import (
	"golang.org/x/exp/constraints"
)

func MergeSort[T constraints.Ordered](data []T) (result []T) {
    size := len(data)
    if size <= 1 {
        return data
    }

    result = make([]T, 0, size)

    mid := size / 2
    left := data[:mid]
    right := data[mid:]
    left = MergeSort(left)
    right = MergeSort(right)

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

    return
}
