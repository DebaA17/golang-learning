package main

import (
    "fmt"
)

func main() {
    var fahrenheit float64

    fmt.Print("Enter temperature in Fahrenheit: ")
    _, err := fmt.Scanf("%f", &fahrenheit)
    if err != nil {
        fmt.Println("Invalid input:", err)
        return
    }

    celsius := (fahrenheit - 32) * 5 / 9
    fmt.Printf("%.f°F is %.f°C\n", fahrenheit, celsius)
}

