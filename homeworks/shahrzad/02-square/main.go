package main

import "fmt"

func main() {

	var squareSide float32

	fmt.Print("Enter the desired square side size: ")
	fmt.Scan(&squareSide)

	area := squareSide*squareSide
	perimeter := squareSide*4

	fmt.Printf("Area of the square: %.2f\n", area)
	fmt.Printf("Perimeter of the square: %.2f\n", perimeter)
}