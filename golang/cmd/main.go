package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/FulecoRafa/merge_sort_wasm/golang/lib"
)

func main() {
    readFile, err := os.Open("../../input/input.txt")
    if err != nil {
        log.Panic(err)
    }

    var data []int

    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)

    for fileScanner.Scan() {
        num, err := strconv.Atoi(fileScanner.Text())
        if err != nil {
            log.Panic(err)
        }
        data = append(data, num)
    }

    start := time.Now()
    result := lib.MergeSort(data)
    log.Println("Time solving merge sort: ", time.Since(start))

    for i, v := range result {
        if v != i {
            log.Panic("Result was unordered")
        }
    }
    log.Println("All good")
}
