package main

import "fmt"

func main(){

	var numberA, numberB int
	
	fmt.Print("Enter the base (A): ")
	fmt.Scan(&numberA)

	fmt.Print("Enter the exponent (B): ")
	fmt.Scan(&numberB)

	result := 1
    for i := 1; i <= numberB; i++ {
        result *= numberA
    }

    fmt.Printf("%d raised to the power of %d is %d", numberA, numberB, result)
}