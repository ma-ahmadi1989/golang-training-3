package main

import (
	"fmt"
)

func main() {
	var numberN, inputNumber int

	fmt.Print("Enter the number of inputs (n): ")
	fmt.Scan(&numberN)

	sum := 0

	for i := 1; i <= numberN; i++ {
		fmt.Printf("Enter number %d: ", i)
		fmt.Scan(&inputNumber)
		sum += inputNumber
	}

	average := float64(sum) / float64(numberN)

	fmt.Printf("The average of the entered numbers is: %.2f", average)
}
