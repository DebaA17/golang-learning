package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Print("Enter the number : ")
	fmt.Scanln(&num)

	if num%2 == 0 {
		fmt.Printf("%d is Even", num)
	} else {
		fmt.Printf("%d is Odd", num)
	}
}
