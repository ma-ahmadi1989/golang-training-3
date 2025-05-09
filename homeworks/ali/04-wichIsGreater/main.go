package main

import "fmt"

func main() {
	var numA float64
	var numB float64
	fmt.Print("enter number one: ")
	fmt.Scan(&numA)
	fmt.Print("enter nuber two: ")
	fmt.Scan(&numB)
	if numA > numB {
		fmt.Printf("%0.2f is greater than %0.2f", numA, numB)
	} else if numA == numB {
		fmt.Printf("they are same")
	} else {
		fmt.Printf("%.02f is greater than %.02f", numB, numA)
	}
}
