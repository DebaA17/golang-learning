package main

import "fmt"

func main() {
    var arr [5]int

    for i := 0; i < 5; i++ {
        fmt.Printf("Enter element %d: ", i+1)
        fmt.Scanln(&arr[i])
    }

    fmt.Println("Array elements are:", arr)
}

