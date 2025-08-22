package main

import (
	"fmt"
)

func main(){
	var arr [5] int 
	

	for i :=0;i<5;i++ {
		fmt.Printf("Enter element %d: ",i+1)
		fmt.Scanln(&arr[i])
	}

	max := arr[0]
	for _, value := range arr {
        if value > max {
            max = value
        }
    }

    fmt.Println("Maximum value in the array is:", max)
}
