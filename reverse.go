package main

import (
	"fmt"
)

func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

func main() {
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)
	reversed := reverseString(input)
	fmt.Println("Reversed string:", reversed)
}
