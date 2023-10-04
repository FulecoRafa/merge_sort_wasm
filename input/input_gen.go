package main

import (
    "fmt"
    "math/rand"
)

func main() {
    // Generate a list of random numbers from 0 to 1000.
    numbers := make([]int, 1000)
    for i := range numbers {
        numbers[i] = i
    }

    // Shuffle the list of numbers.
    for i := 0; i < len(numbers) - 1; i++ {
        j := rand.Intn(len(numbers) - i)
        numbers[i], numbers[j] = numbers[j], numbers[i]
    }

    // Print the list of numbers.
    for _, number := range numbers {
        fmt.Println(number)
    }
}
