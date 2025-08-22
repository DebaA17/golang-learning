package main

import (
	"fmt"
)


func main(){
	var a, b int 
	fmt.Print("Enter the first number : ")
	fmt.Scanln(&a)

	fmt.Print("Enter the second number : ")
	fmt.Scanln(&b)

	fmt.Println("Before swap a = ",a, ", b = ",b)
	

	a, b = b, a

	fmt.Println("After swap a = ",a, ", b = ",b)
}
