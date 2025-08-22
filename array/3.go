package main

import "fmt"

func main() {
    var arr [5]int
    sum := 0

    for i := 0; i < 5; i++ {
        fmt.Printf("Enter element %d: ", i+1)
        fmt.Scanln(&arr[i])
    }

    for _, value := range arr {
        sum += value
    }

    fmt.Println("Sum of the array elements is:", sum)
}
