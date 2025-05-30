package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Printf("Enter the first number:")
	fmt.Scan(&num1)

	fmt.Printf("Enter the two number2:")
	fmt.Scan(&num2)

	if num1 > num2 {
		fmt.Println("the first number is greater:", num1)
	} else {
		fmt.Println("the two number is greater:", num2)
	}
}
