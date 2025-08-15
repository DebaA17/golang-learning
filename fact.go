package main

import (
	"fmt"
)

func main() {
	var n int
	var fact int = 1

	fmt.Print("Enter the number: ")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		fact *= i
	}

	fmt.Printf("Factorial of %d is %d\n", n, fact)
}
