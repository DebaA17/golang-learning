package main

import "fmt"

func main() {
	var principal, rate, time float64

	fmt.Print("Enter principal amount: ")
	fmt.Scanln(&principal)

	fmt.Print("Enter rate of interest (in %): ")
	fmt.Scanln(&rate)

	fmt.Print("Enter time (in years): ")
	fmt.Scanln(&time)

	simpleInterest := (principal * rate * time) / 100

	fmt.Printf("Simple Interest = %f\n", simpleInterest)
}
