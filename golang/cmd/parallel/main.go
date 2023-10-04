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
    readFile, err := os.Open(os.Args[1])
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

    outChan := make(chan []int, 1)
    start := time.Now()
    lib.ParallelMergeSort(data, outChan)
    result := <- outChan
    log.Println("Time solving merge sort: ", time.Since(start))
    close(outChan)

    for i, v := range result {
        if v != i {
            log.Panic("Result was unordered")
        }
    }
    log.Println("All good")
}
