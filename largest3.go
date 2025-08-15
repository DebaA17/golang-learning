package main

import (
	"fmt"
)

func main() {
	var num1, num2, num3 int

	fmt.Print("Enter the first number :")
	fmt.Scanln(&num1)

	fmt.Print("Enter the second number : ")
	fmt.Scanln(&num2)

	fmt.Print("Enter the third number : ")
	fmt.Scanln(&num3)

	if num1 >= num2 && num1 >= num3 {
		fmt.Printf("%d is max", num1)
	} else if num2 >= num1 && num2 >= num3 {
		fmt.Printf("%d is max", num2)
	} else {
		fmt.Printf("%d is max", num3)
	}

}
