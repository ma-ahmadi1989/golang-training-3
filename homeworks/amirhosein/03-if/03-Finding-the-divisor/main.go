package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Enter the first number:")
	fmt.Scan(&a)

	fmt.Print("Enter the two number:")
	fmt.Scan(&b)

	if a > b {
		fmt.Println("the first number is greater:", a)
		if a%b == 0 {
			fmt.Println("The second number divides the first numbe")
		} else {
			fmt.Println("The second number does not divide the first number")
		}
	} else {
		fmt.Println("the two number is greater:")
		if b%a == 0 {
			fmt.Println("The first number divides the second number")
		} else {
			fmt.Println("The first number does not divide the second number")
		}

	}
}
